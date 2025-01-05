package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/eseferr/blog-aggregator/internal/database"
	"github.com/google/uuid"
)
func handlerRegister(s *State, cmd Command) error{
	if len(cmd.Args) == 0 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	params := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
	}
	_, err := s.db.GetUser(context.Background(), cmd.Args[0])
	if err == nil {
        return fmt.Errorf("user %s already exists", cmd.Args[0])
    } else if err != sql.ErrNoRows {
        return fmt.Errorf("error checking for existing user: %w", err)
    }

	user, err := s.db.CreateUser(context.Background(),params)
	if err != nil {
        return fmt.Errorf("couldn't create user: %w", err)
    }
	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}
	fmt.Printf("User %s created successfully\n", user.Name)
	fmt.Println("User has been set")  
	return nil
	}