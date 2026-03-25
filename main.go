package main

import (
	"fmt"
	"log"
	"os"

	"github.com/luckyBambooBro/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	//first read of .gatorconfig.json
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)

	//save read contents to state
	programState := &state{
		cfg: &cfg,
	}

	//Create a new instance of the commands struct with 
	// an initialized map of handler functions.

	cmds := &commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	//register functions here
	cmds.register("login", handlerLogin)

	//obtain args passed in by user in CLI
	args := os.Args
	if len(args) < 2 {
		fmt.Println("cli error: command/args not supplied")
		os.Exit(1)
	}
	//create a command struct to hold the command name and its arguments
	cmd := command{
		Name: args[1],
		Args: args[2:],
	}

	//run the function the user has entered
	if err = cmds.run(programState, cmd); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

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