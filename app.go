package main

import (
	"context"
	"fmt"
	"res/models"
)

// App struct
type App struct {
	ctx context.Context
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

var currentExperiment = models.Experiment{
	Name: "Experiment of Smart Man",
}

var i = 0

func (a *App) GetCurrentExperiment() *models.Experiment {
	i++
	currentExperiment.Name = fmt.Sprintf("Experiment #%d", i)
	return &currentExperiment
}

func (a *App) SaveCurrentExperiment(experiment *models.Experiment) bool {
	currentExperiment = *experiment
	return true
}
