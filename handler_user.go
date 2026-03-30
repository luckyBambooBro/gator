package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/luckyBambooBro/gator/internal/database"

	"github.com/google/uuid"
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
	if len(c.Args) == 0 {
				return errors.New("name not provided for register command")
	} 
	name := c.Args[0]

	//create new user parameters
	newUserParams := database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: name,
	}
	
	//create new user
	newUser, err := s.db.CreateUser(context.Background(), newUserParams)
	if err != nil {
		fmt.Println(err)
		fmt.Printf("User with name \"%s\" already exists\n", name)
		os.Exit(1)
	}

	//set user name in config
	s.cfg.SetUser(newUser.Name)
	fmt.Printf("User %+s successfully created", newUser.Name)
	fmt.Println(newUser)

	return nil

}