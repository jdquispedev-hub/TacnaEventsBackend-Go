package db

import (
    "context"
    "fmt"
    "github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func ConnectDB() (*pgxpool.Pool, error) {
    conn, err := pgxpool.New(context.Background(),
        "postgres://postgres:jesus@localhost:5432/tacna_events_go")
    if err != nil {
        return nil, err
    }

    // Test connection
    if err := conn.Ping(context.Background()); err != nil {
        return nil, err
    }

    fmt.Println("✅ Pool de conexiones PostgreSQL creado")
    return conn, nil
}