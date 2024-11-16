package main

import (
	"fmt"
	"log"
	"os"

	"github.com/samersawan/gator/internal/config"
	"github.com/samersawan/gator/internal/handlers"
)

func main() {

	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	state := handlers.State{Cfg: &cfg}
	commands := handlers.Commands{Cmds: make(map[string]func(*handlers.State, handlers.Command) error)}

	commands.Register("login", handlers.HandlerLogin)

	args := os.Args
	if len(args) < 2 {
		fmt.Println("Not enough arguments!")
		os.Exit(1)
	}
	args = args[1:]
	cmd := handlers.Command{Name: args[0], Args: args[1:]}
	commands.Run(&state, cmd)

}
