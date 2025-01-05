package main

import (
	"context"
	"fmt"
)
func handlerGetUsers(s *State, cmd Command) error{
	users,err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't reset: %w", err)
	}
	for _, item := range users {
		if item == s.cfg.CurrentUserName {
			fmt.Println(item + " (current)")  
		}
        fmt.Println(item)  
    }
	return nil
	}