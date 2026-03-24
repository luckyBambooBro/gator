package main

import (
	"errors"
	"fmt"

	"github.com/luckyBambooBro/gator/internal/config"	
)

type state struct {
	currentState *config.Config
}

type command struct {
	name string
	arguments []string
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 { //or should i have if len(cmd.arguments) != 1 ... (since we want just one argument for username?)
		return errors.New("username not provided")
	}
	if err := s.currentState.SetUser(cmd.arguments[0]); err != nil {
		return err
	}
	fmt.Printf("username set to %s\n", cmd.arguments[0])
	return nil
}

type commands struct {
	handler map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	if _, ok := c.handler[cmd.name]; !ok {
		return fmt.Errorf("%s command does not exist", cmd.name)
	}
	return c.handler[cmd.name](s, cmd)
}

func (c *commands) register (name string, f func(*state, command) error) {
	// This method registers a new handler function for a command name.
	if c.handler == nil {
		c.handler = make(map[string]func(*state, command) error)
	}
	c.handler[name] = f
}