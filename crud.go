package main

import (
	"fmt"
	"log"
)

func main() {
	e := Estudiante{
		Name:   "Jhon",
		Age:    28,
		Active: true,
	}

	err := CrearEstudiante(e)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Registro del estudiante %s creado exitosamente", e.Name)
}
