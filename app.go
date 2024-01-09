package primekit

import (
	"errors"
	"github.com/primebit/primekit/config"
)

var (
	ErrorAllreadyStarted = errors.New("application allready startet")
	instance             *Application
)

func App() *Application {
	if instance == nil {
		instance = &Application{
			configPath: "./config/",
		}
	}

	return instance
}

type Application struct {
	isStarted bool

	configPath string
}

func (app *Application) Run() error {
	if app.isStarted {
		return ErrorAllreadyStarted
	}

	err := config.Init(app.configPath)
	if err != nil {
		return err
	}

	// start sequence

	app.isStarted = true
	return nil
}

func (app *Application) WithConfigPath(path string) *Application {
	app.configPath = path
	return app
}

func (app *Application) WithMetrics() *Application {
	return app
}
