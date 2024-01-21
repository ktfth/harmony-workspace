package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ServiceWeaver/weaver"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

var secretKey = []byte(os.Getenv("SECRET_KEY"))

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
