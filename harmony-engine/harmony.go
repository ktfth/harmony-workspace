package main

import (
	"context"
	"database/sql"

	"github.com/ServiceWeaver/weaver"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type Prompter interface {
	// Register registers a user and returns the user ID.
	Register(context.Context, *User) (*IUserResult, error)

	// Fetch fetches a user based on the provided user object.
	Fetch(context.Context, *User) (*User, error)

	// Bin creates a new prompt and returns the prompt ID.
	Bin(context.Context, *Prompt) (*IPromptResult, error)

	// Retrieve retrieves a prompt based on the provided prompt object.
	Retrieve(context.Context, *Prompt) (*Prompt, error)

	List(context.Context) (*IPromptListResult, error)
}

// prompter represents a prompter implementation that implements the Prompter interface.
// It also includes a configuration from the config package and a database connection from the sql package.
type prompter struct {
	weaver.Implements[Prompter]
	weaver.WithConfig[config]

	db *sql.DB
}

// config represents the configuration for the Harmony engine.
type config struct {
	Driver string
	Source string
}

// Init initializes the prompter.
// It opens a SQL database connection using the driver and source specified in the prompter's configuration.
// The opened database connection is stored in the prompter's 'db' field.
// If an error occurs during the database connection setup, it is returned.
func (p *prompter) Init(_ context.Context) error {
	db, err := sql.Open(p.Config().Driver, p.Config().Source)
	p.db = db
	return err
}

// Register registers a new user in the system.
// It checks if the username already exists in the database.
// If the username exists, it returns 0 and no error.
// If the username does not exist, it inserts the user into the database and returns the generated user ID.
// If any error occurs during the registration process, it returns 0 and the corresponding error.
func (p *prompter) Register(ctx context.Context, user *User) (*IUserResult, error) {
	var logger = p.Logger(ctx)

	logger.Info("registering user")

	var q = "SELECT COUNT(1) FROM users WHERE username = ?;"
	var counter int

	if errCount := p.db.QueryRowContext(ctx, q, user.Username).Scan(&counter); errCount != nil {
		return &IUserResult{
			Id: 0,
		}, errCount
	}

	if counter > 0 {
		return &IUserResult{
			Id: 0,
		}, nil
	}

	q = "INSERT INTO users (username, password) VALUES (?, ?);"
	var hash, errHash = bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if errHash != nil {
		return &IUserResult{
			Id: 0,
		}, errHash
	}

	var result, errExec = p.db.ExecContext(ctx, q, user.Username, hash)

	if errExec != nil {
		return &IUserResult{
			Id: 0,
		}, errExec
	}

	var id, errLastInsert = result.LastInsertId()

	if errLastInsert != nil {
		return &IUserResult{
			Id: 0,
		}, errLastInsert
	}

	logger.Info("user registered")

	return &IUserResult{
		Id: id,
	}, nil
}

// Fetch retrieves a user from the database based on the provided username.
// It returns the fetched user and any error encountered during the retrieval process.
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

// Bin inserts a prompt into the database if it does not already exist.
// It checks if the prompt text already exists in the "prompts" table.
// If the prompt text exists, it returns 0 and no error.
// If the prompt text does not exist, it inserts the prompt into the "prompts" table and returns the last inserted ID.
// If there is an error during the database operations, it returns the error.
func (p *prompter) Bin(ctx context.Context, prompt *Prompt) (*IPromptResult, error) {
	var counter int
	var q = "SELECT COUNT(1) FROM prompts WHERE text = ?;"

	if errCount := p.db.QueryRowContext(ctx, q, prompt.Text).Scan(&counter); errCount != nil {
		return &IPromptResult{
			Id: 0,
		}, errCount
	}

	if counter > 0 {
		return &IPromptResult{
			Id: 0,
		}, nil
	}

	q = "INSERT INTO prompts (text, model, tags) VALUES (?, ?, ?);"
	var result, errExec = p.db.ExecContext(ctx, q, prompt.Text, prompt.Model, prompt.Tags)

	if errExec != nil {
		return &IPromptResult{
			Id: 0,
		}, errExec
	}

	var id, errLastInsert = result.LastInsertId()

	if errLastInsert != nil {
		return &IPromptResult{
			Id: 0,
		}, errLastInsert
	}

	return &IPromptResult{
		Id: id,
	}, nil
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

func (p *prompter) List(ctx context.Context) (*IPromptListResult, error) {
	var logger = p.Logger(ctx)

	logger.Info("listing prompts")

	var q = "SELECT id, text, model, tags FROM prompts;"
	var data []Prompt

	rows, err := p.db.QueryContext(ctx, q)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var prompt Prompt
		if err := rows.Scan(&prompt.Id, &prompt.Text, &prompt.Model, &prompt.Tags); err != nil {
			return nil, err
		}
		data = append(data, prompt)
	}

	logger.Info("prompts listed")

	return &IPromptListResult{
		Prompts: data,
	}, nil
}
