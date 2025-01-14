package main

import (
	"context"

	"github.com/eseferr/blog-aggregator/internal/database"
)

// Middleware function
func middlewareLoggedIn(handler func(s *State, cmd Command, user database.User) error) func(*State, Command) error {
    return func(s *State, cmd Command) error {
        user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
        if err != nil {
            return err
        }
        return handler(s, cmd, user)
    }
}