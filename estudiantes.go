package main

import (
	"errors"
	"time"
)

// Estudiante Struct that's matched with table gocrud.estudiantes
type Estudiante struct {
	ID        int
	Name      string
	Age       uint16
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

// CrearEstudiante registra un estudiante en la DB
func CrearEstudiante(e Estudiante) error {
	q := `INSERT INTO
				estudiantes (name, age, active)
				VALUES ($1, $2, $3)
				`
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	r, err := stmt.Exec(e.Name, e.Age, e.Active)
	if err != nil {
		return err
	}

	i, _ := r.RowsAffected()
	if i != 1 {
		return errors.New("Error QP00N: Se esperaba al menos una registro afectado")
	}

	return nil

}
