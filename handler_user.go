package main

import (
	"fmt"
)
func handlerLogin(s *State, cmd Command) error{
	if len(cmd.Args) == 0 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	err := s.cfg.SetUser(cmd.Args[0]) 
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}
	fmt.Println("User has been set")  
	return nil
	}