package main

import (
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 { //or should i have if len(cmd.Args) != 1 ... (since we want just one argument for username?)
		return fmt.Errorf("usage: <%s> username error. Please provide one username", cmd.Name)
	}
	if err := s.cfg.SetUser(cmd.Args[0]); err != nil {
		return err
	}
	fmt.Println("username switched successfully")
	return nil
}