package main

import (
	"fmt"
	"github.com/example/db"
	"github.com/example/handler"
	"github.com/example/model"
	"github.com/example/repository"
	"github.com/example/router"
	"github.com/labstack/echo/v4"
	"gopkg.in/yaml.v2"
	"os"
	"time"
)

type Log struct {
	ID        int
	Message   string
	CreatedAt time.Time
}

func main() {
	// database connect
	var cfg model.Config
	loadConfig(&cfg)
	var sql = new(db.SQL)
	sql.Connect(cfg)
	defer sql.Close()

	e := echo.New()
	userHandler := handler.UserHandler{
		UserRepo: repository.NewUserRepo(sql),
	}

	api := router.API{
		Echo:        e,
		UserHandler: userHandler,
	}
	api.SetupRouter()
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", cfg.Server.Port)))

}

func loadConfig(cfg *model.Config) {
	f, err := os.Open("../env.dev.yml")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
}
