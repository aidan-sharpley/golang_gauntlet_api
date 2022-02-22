package main

import "github.com/gin-gonic/gin"

// MAIN
type Server struct {
	router *gin.Engine
	ships  *[]ShipCrew
}

// Client Types
type AstroList struct {
	People []Astros `json:"people"`
}
type Astros struct {
	Name  string `json:"name"`
	Craft string `json:"craft"`
}

// API Types
type ShipCrew struct {
	Name string
	Crew []string
}
