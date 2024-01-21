package main

import (
	"context"
	_ "embed"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ServiceWeaver/weaver"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

var validate *validator.Validate

var secretKey []byte

//go:embed index.html
var indexHtml string

func main() {
	dotEnvErr := godotenv.Load()

	if dotEnvErr != nil {
		log.Fatal("Error loading .env file")
	}

	validate = validator.New(validator.WithRequiredStructEnabled())
	secretKey = []byte(os.Getenv("SECRET_KEY"))

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

	http.HandleFunc("/", handleIndex(ctx, app))

	http.HandleFunc("/bin/", handleBin(ctx, app))

	return http.Serve(app.bin, nil)
}
