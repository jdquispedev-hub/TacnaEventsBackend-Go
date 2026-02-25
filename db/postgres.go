package db

import (
    "context"
    "fmt"
    "github.com/jackc/pgx/v5"
    )


func ConnectDB() (*pgx.Conn, error) {
    conn, err := pgx.Connect(context.Background(),
        "postgres://postgres:jesus@localhost:5432/Tacna_events")
    if err != nil {
        return nil, err
    }

    fmt.Println("✅ Conectado a PostgreSQL")
    return conn, nil
}