# Ejemplo de CRUD con Go y Postgresql

Se toma este ejemplo de tres videos de YouTube del instructor Alexys Lozada, en el que explica como escribir código para una aplicación CRUD con la base de datos relacional PostgreSQL.

```url
Video 1: <https://youtu.be/grM6FjvfMP4>
Video 2: <https://youtu.be/Zvm7bg-PxLE>
Video 3: <https://youtu.be/-3O1n5AgphE>
```

Se utilizó un struct Estudiantes de Golang para preparar el esquema en la base de datos, y a partir de allí se contruyeron las sentencias Create, Insert, Select, Update y Delete.

```go
type Estudiante struct {
    ID        int
    Name      string
    Age       int16
    Active    bool
    CreatedAt time.Time
    UpdatedAt time.Time
}
```
