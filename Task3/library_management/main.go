package main

import (
	"library_management/controllers"
	"library_management/services"
)

func main() {
	libraryService := services.NewLibrary()
	controller := controllers.NewLibraryController(libraryService)
	controller.Run()
}