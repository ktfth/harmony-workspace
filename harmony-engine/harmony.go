package main

import (
	"context"
	"database/sql"

	"github.com/ServiceWeaver/weaver"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type Prompter interface {
	// Registering and fetching users
	Register(context.Context, *User) (int64, error)
	Fetch(context.Context, *User) (*User, error)

	// Prompts
	Bin(context.Context, *Prompt) (int64, error)
	Retrieve(context.Context, *Prompt) (*Prompt, error)
}

type prompter struct {
	weaver.Implements[Prompter]
	weaver.WithConfig[config]

	db *sql.DB
}

type config struct {
	Driver string
	Source string
}

func (p *prompter) Init(_ context.Context) error {
	db, err := sql.Open(p.Config().Driver, p.Config().Source)
	p.db = db
	return err
}

func (p *prompter) Register(ctx context.Context, user *User) (int64, error) {
	var logger = p.Logger(ctx)

	logger.Info("registering user")

	var q = "SELECT COUNT(1) FROM users WHERE username = ?;"
	var counter int

	if errCount := p.db.QueryRowContext(ctx, q, user.Username).Scan(&counter); errCount != nil {
		return 0, errCount
	}

	if counter > 0 {
		return 0, nil
	}

	q = "INSERT INTO users (username, password) VALUES (?, ?);"
	var hash, errHash = bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if errHash != nil {
		return 0, errHash
	}

	var result, errExec = p.db.ExecContext(ctx, q, user.Username, hash)

	if errExec != nil {
		return 0, errExec
	}

	var id, errLastInsert = result.LastInsertId()

	if errLastInsert != nil {
		return 0, errLastInsert
	}

	logger.Info("user registered")

	return id, nil
}

func (p *prompter) Fetch(ctx context.Context, user *User) (*User, error) {
	var logger = p.Logger(ctx)

	logger.Info("fetching user")

	var q = "SELECT id, username, password FROM users WHERE username = ?;"
	var data User

	if err := p.db.QueryRowContext(ctx, q, user.Username).Scan(&data.Id, &data.Username, &data.Password); err != nil {
		return nil, err
	}

	logger.Info("user fetched")

	return &data, nil
}

func (p *prompter) Bin(ctx context.Context, prompt *Prompt) (int64, error) {
	var counter int
	var q = "SELECT COUNT(1) FROM prompts WHERE text = ?;"

	if errCount := p.db.QueryRowContext(ctx, q, prompt.Text).Scan(&counter); errCount != nil {
		return 0, errCount
	}

	if counter > 0 {
		return 0, nil
	}

	q = "INSERT INTO prompts (text, model, tags) VALUES (?, ?, ?);"
	var result, errExec = p.db.ExecContext(ctx, q, prompt.Text, prompt.Model, prompt.Tags)

	if errExec != nil {
		return 0, errExec
	}

	var id, errLastInsert = result.LastInsertId()

	if errLastInsert != nil {
		return 0, errLastInsert
	}

	return id, nil
}

func (p *prompter) Retrieve(ctx context.Context, prompt *Prompt) (*Prompt, error) {
	var logger = p.Logger(ctx)

	logger.Info("retrieving prompt")

	var q = "SELECT id, text, model, tags FROM prompts WHERE id = ?;"
	var data Prompt

	if err := p.db.QueryRowContext(ctx, q, prompt.Id).Scan(&data.Id, &data.Text, &data.Model, &data.Tags); err != nil {
		return nil, err
	}

	logger.Info("prompt retrieved")

	return &data, nil
}
