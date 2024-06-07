package db

import (
	"database/sql"
	"fmt"
	"log"
	"os/exec"
	"todolist/config"

	_ "github.com/lib/pq"
)

func createDatabase(config *config.Config) error {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", config.Database.Host, config.Database.Port, config.Database.User, config.Database.Password)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	defer db.Close()

	var exists bool
	err = db.QueryRow("SELECT EXISTS (SELECT 1 FROM pg_database WHERE datname = $1)", config.Database.Name).Scan(&exists)
	if err != nil {
		return err
	}
	if exists {
		return nil
	}

	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s", config.Database.Name))
	if err != nil {
		return err
	}

	return nil
}

func runMigrations(config *config.Config) error {
	migrateCmd := exec.Command("migrate", "-path", "migrations", "-database", fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", config.Database.User, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Name), "up")
	migrateCmd.Stdout = log.Writer()
	migrateCmd.Stderr = log.Writer()
	return migrateCmd.Run()
}

func InitDB() (*sql.DB, *config.Config, error) {
	cfg, err := config.ReadConfig("config.yaml")
	if err != nil {
		return nil, nil, fmt.Errorf("could not read config: %w", err)
	}

	err = createDatabase(cfg)
	if err != nil {
		return nil, nil, fmt.Errorf("could not create database: %w", err)
	}

	dbConnStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.Name)
	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		return nil, nil, fmt.Errorf("could not open database: %w", err)
	}

	err = runMigrations(cfg)
	if err != nil {
		return nil, nil, fmt.Errorf("could not run migrations: %w", err)
	}

	return db, cfg, nil
}

func dropTable(db *sql.DB) error {
	_, err := db.Exec("DROP TABLE IF EXISTS tasks;")
	return err
}

func dropDatabase(config *config.Config) error {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", config.Database.Host, config.Database.Port, config.Database.User, config.Database.Password)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s", config.Database.Name))
	return err
}

func CloseDB(db *sql.DB, config *config.Config) {
	err := dropTable(db)
	if err != nil {
		log.Fatalf("could not drop table: %v", err)
	}

	db.Close()

	err = dropDatabase(config)
	if err != nil {
		log.Fatalf("could not drop database: %v", err)
	}
}
