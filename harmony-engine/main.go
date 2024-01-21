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

func handleBin(ctx context.Context, app *app) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var logger = app.Logger(ctx)
		logger.Info("bin request received")

		var subMethod = strings.TrimPrefix(r.URL.Path, "/bin/")
		var p Prompter = app.prompter.Get()

		switch r.Method {
		case "POST":
			handlePost(ctx, app, p, w, r, subMethod)
		case "GET":
			handleGet(ctx, app, p, w, r)
		default:
			logger.Warn("bin request is not a POST or a GET")
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func handlePost(ctx context.Context, app *app, p Prompter, w http.ResponseWriter, r *http.Request, subMethod string) {
	if subMethod == "auth" {
		AuthHandler(ctx, app, p, w, r)
		return
	}

	if subMethod == "register" {
		RegisterHandler(ctx, app, p, w, r)
		return
	}

	verifyAndCreatePrompt(ctx, app, p, w, r)
}

func handleGet(ctx context.Context, app *app, p Prompter, w http.ResponseWriter, r *http.Request) {
	verifyAndRetrievePrompt(ctx, app, p, w, r)
}

func verifyAndCreatePrompt(ctx context.Context, app *app, p Prompter, w http.ResponseWriter, r *http.Request) {
	verifyErr := verifyTokenTrigger(ctx, app, w, r)

	if verifyErr != nil {
		http.Error(w, verifyErr.Error(), http.StatusUnauthorized)
	} else {
		CreatePrompt(ctx, app, p, w, r)
	}
}

func verifyAndRetrievePrompt(ctx context.Context, app *app, p Prompter, w http.ResponseWriter, r *http.Request) {
	verifyErr := verifyTokenTrigger(ctx, app, w, r)

	if verifyErr != nil {
		http.Error(w, verifyErr.Error(), http.StatusUnauthorized)
	} else {
		RetrievePrompt(ctx, app, p, w, r)
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
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	const fieldValidation = "required,min=3,max=255"

	errsUsername := validate.Var(user.Username, fieldValidation)
	errsPassword := validate.Var(user.Password, fieldValidation)

	if errsUsername != nil {
		http.Error(w, errsUsername.Error(), http.StatusBadRequest)
		return
	}

	if errsPassword != nil {
		http.Error(w, errsPassword.Error(), http.StatusBadRequest)
		return
	}

	logger.Info("bin request is a POST with a user and validations")

	var bin, err = p.Register(ctx, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	const fieldValidation = "required,min=3,max=255"

	errsPromptText := validate.Var(prompt.Text, fieldValidation)

	if errsPromptText != nil {
		http.Error(w, errsPromptText.Error(), http.StatusBadRequest)
		return
	}

	logger.Info("bin request is a POST with a prompt and validations")

	var bin, err = p.Bin(ctx, prompt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
		http.Error(w, errId.Error(), http.StatusBadRequest)
		return
	}

	logger.Info("bin request is a GET with an id")

	var prompt = &Prompt{
		Id: idInt,
	}
	var promptResult, errPrompt = p.Retrieve(ctx, prompt)
	if errPrompt != nil {
		http.Error(w, errPrompt.Error(), http.StatusInternalServerError)
		return
	}

	logger.Info("bin request is a GET with an id and a prompt")

	var promptResultInJson, _ = json.MarshalIndent(promptResult, "", "  ")
	var promptOut = string(promptResultInJson)

	logger.Info("bin request is a GET with an id and a prompt and a json")

	fmt.Fprintf(w, "%v", promptOut)

	return
}
