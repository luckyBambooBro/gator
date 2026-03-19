package main

import (
	"fmt"
	"github.com/luckyBambooBro/gator/internal/config" 
	"log"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	err = cfg.SetUser("User1")
	if err != nil {
		log.Fatal(err)
	}
	//reread and print to terminal
	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	fmt.Print(*cfg)
}