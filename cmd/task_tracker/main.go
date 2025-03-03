package main

import (
	"github.com/Ma-Leal/to-do-list/internal/infra/files"
	"github.com/Ma-Leal/to-do-list/internal/infra/handlers/cli"
	"github.com/Ma-Leal/to-do-list/internal/usecase"
)

func main() {
	repo := files.NewTaskRepositoryFile("./task-tracker.json")
	taskUseCase := usecase.NewTaskUseCase(repo)
	handlerCLI := cli.NewTaskHandler(*taskUseCase)
	handlerCLI.Run()
}
