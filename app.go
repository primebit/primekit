package primekit

import "errors"

var (
	ErrorAllreadyStarted = errors.New("application allready startet")
	instance             *Application
)

func App() *Application {
	if instance == nil {
		instance = &Application{}
	}

	return instance
}

type Application struct {
	isStarted bool
}

func (app *Application) Run() error {
	if app.isStarted {
		return ErrorAllreadyStarted
	}

	// start sequence

	app.isStarted = true
	return nil
}
