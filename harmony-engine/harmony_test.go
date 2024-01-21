package main

import (
	"context"
	"testing"

	"github.com/ServiceWeaver/weaver/weavertest"
)

func TestAdd(t *testing.T) {
	runner := weavertest.Local // A runner that runs components in a single process
	runner.Config = `
	[serviceweaver]
	binary = ".\\harmony-engine.exe"

	[single]
	listeners.bin = {address = "localhost:12345"}

	[multi]
	listeners.bin = {address = "localhost:12345"}

	["harmony-engine/Prompter"]
	Driver = "mysql"
	Source = "root:@tcp(localhost:3306)/harmony"
	`
	runner.Test(t, func(t *testing.T, p Prompter) {
		ctx := context.Background()
		prompt := &Prompt{
			Text:  "What is the meaning of life?",
			Model: "GPT-2",
			Tags:  "philosophy, life",
		}
		got, err := p.Bin(ctx, prompt)
		if err != nil {
			t.Fatal(err)
		}
		t.Log("got", got)
		if got > 0 {
			t.Fatalf("got %q > 0", got)
		}
	})
}
