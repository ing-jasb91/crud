package main

import "log"

func main() {
	// e := Estudiante{
	// 	// Name: "Jhonny",
	// 	// Age:    28,
	// 	// Active: true,
	// }

	// err := CrearEstudiante(e)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("Registro del estudiante %s creado exitosamente.\n", e.Name)

	// es, err := ConsultarDB()
	// if err != nil {
	// 	log.Println(err)
	// }
	// for _, v := range es {
	// 	fmt.Println(v.Name)
	// }
	// e := Estudiante{
	// 	ID:     1,
	// 	Name:   "Luisa",
	// 	Age:    52,
	// 	Active: false,
	// }

	// err := ActualizarDB(e)
	// if err != nil {
	// 	log.Println(err)
	// }

	// log.Println("Registro ID", e.ID, "actualizado correctamente")
	e := Estudiante{
		ID: 5,
	}

	err := DeleteDB(e.ID)
	if err != nil {
		log.Println(err)
	}

	log.Println("Registro ID", e.ID, "eliminado correctamente")
}
