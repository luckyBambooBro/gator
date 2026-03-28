package main

import (
	"errors"
	"fmt"

	"github.com/luckyBambooBro/gator/internal/database"
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

/*
Ensure that a name was passed in the args.
Create a new user in the database. It should have access to the CreateUser query through the state -> db struct.
Pass context.Background() to the query to create an empty Context argument.
Use the uuid.New() function to generate a new UUID for the user.
created_at and updated_at should be the current time.
Use the provided name.
Exit with code 1 if a user with that name already exists.
Set the current user in the config to the given name.
Print a message that the user was created, and log the user's data to the console for your own debugging.
*/
func handlerRegister(s *state,c command) error {
	name := c.Args[0]
	if name == "" {
		return errors.New("name not provided for register command")
	}
}