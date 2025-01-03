package main

import (
	"errors"
)
type Command struct {
	Name string
	Args []string
}
type Commands struct {
	commandHandlers map[string]func(*State, Command) error
}

func (c *Commands) register(name string, f func(*State, Command) error){
	c.commandHandlers[name] = f
}
func (c *Commands) run(s *State, cmd Command) error{
	handler, ok := c.commandHandlers[cmd.Name]
	if !ok{
		return errors.New("Command not found")
	}	
	return handler(s,cmd)
}