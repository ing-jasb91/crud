package main

import (
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/lib/pq"
)

// Estudiante Struct that's matched with table gocrud.estudiantes
type Estudiante struct {
	ID        int
	Name      string
	Age       int16
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

	intNull := sql.NullInt64{}
	strNull := sql.NullString{}

	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if e.Age == 0 {
		intNull.Valid = false
	} else {
		intNull.Valid = true
		intNull.Int64 = int64(e.Age)
	}

	if e.Name == "" {
		strNull.Valid = false
	} else {
		strNull.Valid = true
		strNull.String = e.Name
	}

	r, err := stmt.Exec(strNull, intNull, e.Active)
	if err != nil {
		return err
	}

	i, _ := r.RowsAffected()
	if i < 1 {
		return errors.New("Error QP00N: Se esperaba al menos una registro afectado")
	}
	log.Println("Se registro un cambio de", i, "registro.")

	return nil

}

// ConsultarDB busca la información de la tabla estudiantes
func ConsultarDB() (estudiante []Estudiante, err error) {
	q := `SELECT id, name, age, active, "createAt", "updateAt"
			FROM estudiantes`

	timeNull := pq.NullTime{}
	intNull := sql.NullInt64{}
	strNull := sql.NullString{}
	boolNull := sql.NullBool{}

	db := getConnection()
	defer db.Close()

	rows, err := db.Query(q)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		e := Estudiante{}
		err = rows.Scan(
			&e.ID,
			&strNull,
			&intNull,
			&boolNull,
			&e.CreatedAt,
			&timeNull,
		)
		if err != nil {
			return
		}

		e.Name = strNull.String
		e.Age = int16(intNull.Int64)
		e.Active = boolNull.Bool

		e.UpdatedAt = timeNull.Time

		estudiante = append(estudiante, e)
	}
	return estudiante, nil
}

// ActualizarDB actualiza registro de la tabla estudiantes
func ActualizarDB(e Estudiante) error {
	q := `UPDATE estudiantes
		SET name = $1, age =$2, active = $3, "updateAt" = now()
		WHERE id = $4`

	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(q)

	if err != nil {
		return err
	}
	defer stmt.Close()

	r, err := stmt.Exec(e.Name, e.Age, e.Active, e.ID)
	if err != nil {
		return err
	}
	i, _ := r.RowsAffected()
	if i < 1 {
		return errors.New("Error QP00N: Se esperaba al menos una registro afectado")
	}
	return nil
}

// DeleteDB elimina un registro de la tabla estudiantes
func DeleteDB(id int) error {
	q := `DELETE FROM estudiantes WHERE id = $1`

	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(q)

	if err != nil {
		return err
	}
	defer stmt.Close()

	r, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	i, _ := r.RowsAffected()
	if i < 1 {
		return errors.New("Error QP00N: Se esperaba al menos una registro afectado")
	}
	return nil

}
