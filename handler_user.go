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
	name := cmd.Args[0]
	//check the user exists (has been registered) in database before allowing login
	if _, err := s.db.GetUser(context.Background(), name); err != nil {
		fmt.Printf("username %q does not exist in database\n", name)
		os.Exit(1)
	}

	//set the user
	if err := s.cfg.SetUser(name); err != nil {
		return err
	}
	fmt.Println("username switched successfully")
	return nil
}

func handlerRegister(s *state, c command) error {
	if len(c.Args) == 0 {
		return errors.New("name not provided for register command")
	}
	name := c.Args[0]

	//create new user parameters
	newUserParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
	}

	//create new user
	newUser, err := s.db.CreateUser(context.Background(), newUserParams)
	if err != nil {
		fmt.Println(err)
		fmt.Printf("User with name %q already exists\n", name)
		os.Exit(1)
	}

	//set user name in config
	s.cfg.SetUser(newUser.Name)
	fmt.Printf("User %q successfully created\n", newUser.Name)
	fmt.Println(newUser)

	return nil

}

func handlerReset(s *state, c command) error {
	if err := s.db.DeleteUsers(context.Background()); err != nil {
		return err
	}
	fmt.Println("Successfully reset all users")
	return nil
}

func handlerListUsers(s *state, c command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error obtaining users: %w", err)

	}

	currentUser := s.cfg.CurrentUserName //up to here
	for _, user := range users {
		if user.Name == currentUser {
			fmt.Println(user.Name, "(current)")
			continue
		}
		fmt.Println(user.Name)
	}
	return nil
}

