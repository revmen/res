package main

import (
	"context"
	"res/models"
)

// App struct
type App struct {
	ctx        context.Context
	Experiment *models.Experiment
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) NewExperiment() *models.Experiment {
	a.Experiment = models.NewExperiment()
	return a.Experiment
}
