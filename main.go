package main

import (
	"fmt"
	"log"
	"os"

	"github.com/luckyBambooBro/gator/internal/config"
)

func main() {
	//first read of .gatorconfig.json
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)

	//save read contents to state
	s := &state{
		currentState: &cfg,
	}

	//Create a new instance of the commands struct with 
	// an initialized map of handler functions.

	cmds := &commands{
		handler: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)

	//obtain args passed in by user in CLI
	args := os.Args
	if len(args) < 2 {
		fmt.Println("cli usage: \"gator command args...\". command not supplied")
		os.Exit(1)
	}

	//create a command struct to hold the command name and its arguments
	cmd := command{
		name: args[1],
		arguments: args[2:],
	}

	cmds.run(s, cmd)
}
/*
	err = cfg.SetUser("User12333")
	if err != nil {
		log.Fatal(err)
	}
	//reread and print to terminal
	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)
}
*/