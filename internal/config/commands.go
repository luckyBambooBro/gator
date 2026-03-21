package config

type state struct {
	currentState *Config
}

type command struct {
	name string
	arguments []string
}

func handlerLogin(s *state, cmd command) error {
	
}
