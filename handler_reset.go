package main

import (
	"context"
	"fmt"
)
func handlerReset(s *State, cmd Command) error{
	err := s.db.DeleteUser(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't reset: %w", err)
	}
	fmt.Println("Users have been reset")  
	return nil
	}