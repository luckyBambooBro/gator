package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/luckyBambooBro/gator/internal/config"
	"github.com/luckyBambooBro/gator/internal/database"

	_ "github.com/lib/pq"
)

const contextTimeout = 5 * time.Second

type state struct {
	db  *database.Queries
	cfg *config.Config
	timeout time.Duration
}

func main() {
	//first read of .gatorconfig.json
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	//db returns a *sql.DB - this is the postgres connection to a database
	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	//database.New() takes a DBTX interface, *sql.DB fits this interface
	//returns *database.Queries which has the methods for Go-SQL code
	dbQueries := database.New(db)

	// the following line is useful to have, but it made me fail lesson ch3l3
	// fmt.Printf("Read config: %+v\n", cfg)

	//save read contents to state
	programState := &state{
		db:  dbQueries,
		cfg: &cfg,
		timeout: contextTimeout,
	}

	//Create a new instance of the commands struct with
	// an initialized map of handler functions.

	cmds := &commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	//register functions here
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerListUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", handlerAddFeed)
	cmds.register("feeds", handlerListFeeds)
	cmds.register("follow", handlerFollow)
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
