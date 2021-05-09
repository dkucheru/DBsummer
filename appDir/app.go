package appDir

import (
	"DBsummer/apiDir"
	"DBsummer/configDir"
	"DBsummer/dataBaseDir"
	"log"
	"os"
)

type App struct {
	Repository dataBaseDir.Repository
	Api        *apiDir.Rest
}

func New() (*App, error) {
	var err error
	app := &App{}

	app.Repository, err = dataBaseDir.NewDatabaseRepository(log.New(os.Stdout, "", log.Ltime), configDir.DBConfig{
		UserName: "postgres",
		Password: "burritos12345678",
		Host:     "127.0.0.1",
		Port:     5432,
		DBName:   "DBproject",
		SSL:      "?sslmode=disable",
	})
	if err != nil {
		return nil, err
	}

	app.Api = apiDir.New(":8080")
	return app, nil
}

func (a *App) Run() error {
	return a.Api.Listen()
}
