package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/samersawan/gator/internal/config"
	"github.com/samersawan/gator/internal/database"
	"github.com/samersawan/gator/internal/handlers"
)

func main() {

	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)
	state := handlers.State{Cfg: &cfg, Db: dbQueries}

	commands := handlers.Commands{Cmds: make(map[string]func(*handlers.State, handlers.Command) error)}

	commands.Register("login", handlers.HandlerLogin)
	commands.Register("register", handlers.HandlerRegister)
	commands.Register("reset", handlers.HandlerReset)
	commands.Register("users", handlers.HandlerGetUsers)

	args := os.Args
	if len(args) < 2 {
		fmt.Println("Not enough arguments!")
		os.Exit(1)
	}
	args = args[1:]
	cmd := handlers.Command{Name: args[0], Args: args[1:]}
	err = commands.Run(&state, cmd)
	if err != nil {
		log.Fatal(err)
	}

}
