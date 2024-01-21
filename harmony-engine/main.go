package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ServiceWeaver/weaver"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var validate *validator.Validate

var secretKey = []byte("super-secret")

func main() {
	validate = validator.New(validator.WithRequiredStructEnabled())

	if err := weaver.Run(context.Background(), serve); err != nil {
		log.Fatal(err)
	}
}

type app struct {
	weaver.Implements[weaver.Main]
	prompter weaver.Ref[Prompter]
	bin      weaver.Listener
}

func serve(ctx context.Context, app *app) error {
	fmt.Printf("bin listener available on %v\n", app.bin)

	http.HandleFunc("/bin/", handleBin(ctx, app))

	return http.Serve(app.bin, nil)
}

// handleBin handles the HTTP requests for the "/bin" endpoint.
// It takes a context, app instance, and returns an http.HandlerFunc.
// The returned handler function processes the incoming requests and performs the necessary actions based on the request method and sub-method.
func handleBin(ctx context.Context, app *app) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var logger = app.Logger(ctx)

		logger.Info("bin request received")

		var subMethod = strings.TrimPrefix(r.URL.Path, "/bin/")

		var p Prompter = app.prompter.Get()

		if subMethod == "auth" && r.Method == "POST" {
			AuthHandler(ctx, app, p, w, r)
			return
		}

		if r.Method == "POST" {
			if subMethod == "register" {
				RegisterHandler(ctx, app, p, w, r)
				return
			}

			verifyErr := verifyTokenTrigger(ctx, app, w, r)

			if verifyErr != nil {
				fmt.Fprintf(w, "%v", verifyErr, http.StatusUnauthorized)
			} else {
				CreatePrompt(ctx, app, p, w, r)
			}
		} else if r.Method == "GET" {
			verifyErr := verifyTokenTrigger(ctx, app, w, r)

			if verifyErr != nil {
				fmt.Fprintf(w, "%v", verifyErr, http.StatusUnauthorized)
			} else {
				RetrievePrompt(ctx, app, p, w, r)
			}
		} else {
			logger.Warn("bin request is not a POST or a GET")

			fmt.Fprintf(w, "%v", "Method not allowed", http.StatusMethodNotAllowed)

			return
		}
	}
}

func createToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func verifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}

func verifyTokenTrigger(ctx context.Context, app *app, w http.ResponseWriter, r *http.Request) error {
	var logger = app.Logger(ctx)

	logger.Info("verify token")

	w.Header().Set("Content-Type", "application/json")
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Missing authorization header")
		return fmt.Errorf("Missing authorization header")
	}
	tokenString = tokenString[len("Bearer "):]

	logger.Info("verify token with token string")

	err := verifyToken(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Invalid token")
		return err
	}

	logger.Info("verify token with token string and no error")

	return nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func RegisterHandler(ctx context.Context, app *app, p Prompter, w http.ResponseWriter, r *http.Request) {
	var logger = app.Logger(ctx)

	logger.Info("bin request is a POST")

	var user = &User{}
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		fmt.Fprintf(w, "%v", err, http.StatusBadRequest)
		return
	}

	const fieldValidation = "required,min=3,max=255"

	errsUsername := validate.Var(user.Username, fieldValidation)
	errsPassword := validate.Var(user.Password, fieldValidation)

	if errsUsername != nil {
		fmt.Fprintf(w, "%v", errsUsername, http.StatusBadRequest)
		return
	}

	if errsPassword != nil {
		fmt.Fprintf(w, "%v", errsPassword, http.StatusBadRequest)
		return
	}

	logger.Info("bin request is a POST with a user and validations")

	var bin, err = p.Register(ctx, user)
	if err != nil {
		fmt.Fprintf(w, "%v", err, http.StatusInternalServerError)
		return
	}

	logger.Info("bin request is a POST with a user and a bin")

	var binResultInJson, _ = json.MarshalIndent(bin, "", "  ")
	var binOut = string(binResultInJson)

	logger.Info("bin request is a POST with a user and a bin and a json")

	fmt.Fprintf(w, "%v", binOut)

	return
}

func AuthHandler(ctx context.Context, app *app, p Prompter, w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/json")

		var u User
		json.NewDecoder(r.Body).Decode(&u)
		fmt.Printf("The user request value %v", u)

		registeredUser, registeredErr := p.Fetch(ctx, &u)

		if registeredErr != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Invalid credentials")
			return
		}

		if u.Username == registeredUser.Username && CheckPasswordHash(u.Password, registeredUser.Password) {
			tokenString, err := createToken(u.Username)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Errorf("No username found")
			}
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, tokenString)
			return
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Invalid credentials")
		}
	}
}

func CreatePrompt(ctx context.Context, app *app, p Prompter, w http.ResponseWriter, r *http.Request) {
	var logger = app.Logger(ctx)

	logger.Info("bin request is a POST")

	var prompt = &Prompt{}
	if err := json.NewDecoder(r.Body).Decode(prompt); err != nil {
		fmt.Fprintf(w, "%v", err, http.StatusBadRequest)
		return
	}

	const fieldValidation = "required,min=3,max=255"

	errsPromptText := validate.Var(prompt.Text, fieldValidation)

	if errsPromptText != nil {
		fmt.Fprintf(w, "%v", errsPromptText, http.StatusBadRequest)
		return
	}

	logger.Info("bin request is a POST with a prompt and validations")

	var bin, err = p.Bin(ctx, prompt)
	if err != nil {
		fmt.Fprintf(w, "%v", err, http.StatusInternalServerError)
		return
	}

	logger.Info("bin request is a POST with a prompt and a bin")

	var binResultInJson, _ = json.MarshalIndent(bin, "", "  ")
	var binOut = string(binResultInJson)

	logger.Info("bin request is a POST with a prompt and a bin and a json")

	fmt.Fprintf(w, "%v", binOut)

	return
}

func RetrievePrompt(ctx context.Context, app *app, p Prompter, w http.ResponseWriter, r *http.Request) {
	var logger = app.Logger(ctx)

	logger.Info("bin request is a GET")

	id := strings.TrimPrefix(r.URL.Path, "/bin/")

	var idInt, errId = strconv.Atoi(id)

	if errId != nil {
		fmt.Fprintf(w, "%v", errId, http.StatusBadRequest)
		return
	}

	logger.Info("bin request is a GET with an id")

	var prompt = &Prompt{
		Id: idInt,
	}
	var promptResult, errPrompt = p.Retrieve(ctx, prompt)
	if errPrompt != nil {
		fmt.Fprintf(w, "%v", errPrompt, http.StatusInternalServerError)
		return
	}

	logger.Info("bin request is a GET with an id and a prompt")

	var promptResultInJson, _ = json.MarshalIndent(promptResult, "", "  ")
	var promptOut = string(promptResultInJson)

	logger.Info("bin request is a GET with an id and a prompt and a json")

	fmt.Fprintf(w, "%v", promptOut)

	return
}
