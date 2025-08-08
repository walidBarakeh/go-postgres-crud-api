package main

import (
    "flag"
    "fmt"
    "log"
    "os"

    "github.com/golang-migrate/migrate/v4"
    _ "github.com/golang-migrate/migrate/v4/database/postgres"
    _ "github.com/golang-migrate/migrate/v4/source/file"
    "github.com/joho/godotenv"
)

func main() {
    // Load environment variables
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found")
    }

    var action = flag.String("action", "", "Migration action: up, down, force, version")
    var steps = flag.Int("steps", 0, "Number of steps for up/down migration")
    var version = flag.Int("version", 0, "Version number for force migration")
    flag.Parse()

    if *action == "" {
        log.Fatal("Please specify an action: -action=up|down|force|version")
    }

    dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_NAME"),
        os.Getenv("DB_SSL_MODE"),
    )

    m, err := migrate.New("file://migrations", dbURL)
    if err != nil {
        log.Fatal(err)
    }
    defer m.Close()

    switch *action {
    case "up":
        if *steps > 0 {
            err = m.Steps(*steps)
        } else {
            err = m.Up()
        }
    case "down":
        if *steps > 0 {
            err = m.Steps(-*steps)
        } else {
            err = m.Down()
        }
    case "force":
        if *version == 0 {
            log.Fatal("Please specify version: -version=N")
        }
        err = m.Force(*version)
    case "version":
        ver, dirty, verErr := m.Version()
        if verErr != nil {
            log.Fatal(verErr)
        }
        fmt.Printf("Version: %d, Dirty: %t\n", ver, dirty)
        return
    default:
        log.Fatal("Invalid action. Use: up, down, force, or version")
    }

    if err != nil && err != migrate.ErrNoChange {
        log.Fatal(err)
    }

    if err == migrate.ErrNoChange {
        log.Println("No migrations to apply")
    } else {
        log.Println("Migration completed successfully")
    }
}