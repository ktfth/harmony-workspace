package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func handleIndex(ctx context.Context, app *app) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var logger = app.Logger(ctx)
		logger.Info("index request received")

		if _, err := fmt.Fprint(w, indexHtml); err != nil {
			logger.Error("error writing index.html to response writer", "err", err)
		}
	}
}

// handleBin handles the "/bin" endpoint and its sub-methods.
// It receives a bin request and delegates the handling to the appropriate sub-method based on the HTTP method.
// The sub-method is extracted from the URL path.
// If the HTTP method is not allowed, it returns a "Method not allowed" error.
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

// handlePost handles the POST requests based on the subMethod parameter.
// It calls the appropriate handler function based on the subMethod value.
// If subMethod is "auth", it calls AuthHandler.
// If subMethod is "register", it calls RegisterHandler.
// Otherwise, it calls verifyAndCreatePrompt.
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

// handleGet handles the GET request and verifies the prompt before retrieving it.
// It takes the context, app instance, prompter, response writer, and request as parameters.
func handleGet(ctx context.Context, app *app, p Prompter, w http.ResponseWriter, r *http.Request) {
	verifyAndRetrievePrompt(ctx, app, p, w, r)
}

// verifyAndCreatePrompt is a function that verifies the token trigger and creates a prompt.
// It takes a context.Context, *app, Prompter, http.ResponseWriter, and *http.Request as parameters.
// If the token trigger is not verified, it returns an unauthorized error.
// Otherwise, it calls the CreatePrompt function.
func verifyAndCreatePrompt(ctx context.Context, app *app, p Prompter, w http.ResponseWriter, r *http.Request) {
	verifyErr := verifyTokenTrigger(ctx, app, w, r)

	if verifyErr != nil {
		http.Error(w, verifyErr.Error(), http.StatusUnauthorized)
	} else {
		CreatePrompt(ctx, app, p, w, r)
	}
}

// verifyAndRetrievePrompt is a function that verifies the token trigger and retrieves the prompt.
// It takes the context, app, Prompter, http.ResponseWriter, and http.Request as parameters.
// If the token trigger is not verified, it returns an unauthorized error.
// Otherwise, it calls the RetrievePrompt function.
func verifyAndRetrievePrompt(ctx context.Context, app *app, p Prompter, w http.ResponseWriter, r *http.Request) {
	verifyErr := verifyTokenTrigger(ctx, app, w, r)

	if verifyErr != nil {
		http.Error(w, verifyErr.Error(), http.StatusUnauthorized)
	} else {
		RetrievePrompt(ctx, app, p, w, r)
	}
}

// verifyTokenTrigger is a function that verifies the token in the request header.
// It takes the context, app, response writer, and request as parameters.
// It returns an error if the token is missing or invalid.
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

// RegisterHandler handles the registration request.
// It decodes the user information from the request body,
// performs field validation on the username and password,
// registers the user using the provided Prompter,
// and returns the registered bin in JSON format.
// If any error occurs during the process, it returns an appropriate HTTP error response.
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

// AuthHandler handles the authentication logic for the application.
// It receives a context, an app instance, a Prompter, a http.ResponseWriter, and a http.Request.
// If the request method is POST, it decodes the JSON body into a User struct.
// It then fetches the registered user using the Prompter.Fetch method.
// If the credentials are invalid, it returns a 401 Unauthorized response.
// If the credentials are valid, it creates a token using the createToken function and returns a 200 OK response with the token.
// If there is an error creating the token, it returns a 500 Internal Server Error response.
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

// CreatePrompt is a handler function that creates a prompt.
// It takes a context.Context, *app, Prompter, http.ResponseWriter, and *http.Request as parameters.
// It decodes the request body into a Prompt struct, validates the prompt text, and creates a bin using the Prompter.
// Finally, it writes the bin result in JSON format to the http.ResponseWriter.
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

// RetrievePrompt retrieves a prompt based on the provided ID.
// It takes a context.Context, an *app, a Prompter, an http.ResponseWriter, and an *http.Request as parameters.
// The function first extracts the ID from the request URL and converts it to an integer.
// If the ID is not a valid integer, it returns a HTTP 400 Bad Request error.
// It then retrieves the prompt using the Prompter's Retrieve method.
// If there is an error retrieving the prompt, it returns a HTTP 500 Internal Server Error.
// Finally, it marshals the prompt result into JSON format and writes it to the response writer.
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
