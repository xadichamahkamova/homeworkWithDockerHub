package main

import (
	"fmt"
	"log"
	"service/config"
	api "service/internal/http"
	conn "service/internal/pkg"
	taskRepo "service/internal/repository"
	taskService "service/internal/service"
)

func main () {

	cfg, err := config.Load(".")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Configuration loaded", cfg)

	mongo, err := conn.NewConnection(*cfg)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB")

	in := taskRepo.NewTaskRepo(mongo)
	service := taskService.NewTaskService(in)
	r := api.NewGin(service)
	addr := fmt.Sprintf(":%s", cfg.Service.Port)
	r.Run(addr)
}