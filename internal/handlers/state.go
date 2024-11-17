package handlers

import (
	"github.com/samersawan/gator/internal/config"
	"github.com/samersawan/gator/internal/database"
)

type State struct {
	Cfg *config.Config
	Db  *database.Queries
}
