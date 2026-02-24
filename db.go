package main

import (
    "context"
    "fmt"
    "github.com/jackc/pgx/v5"
)

func ConnectDB() (*pgx.Conn, error) {
    conn, err := pgx.Connect(context.Background(),
        "postgres://eventsuser:pg123456@localhost:5432/tacna_events")
    if err != nil {
        return nil, err
    }

    fmt.Println("✅ Conectado a PostgreSQL")
    return conn, nil
}