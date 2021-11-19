package env

import (
	"log"
	"os"
)

type Database struct {
	Name     string
	User     string
	Password string
	Driver   string
	Host     string
}

func (d *Database) Set() {
	d.Name = os.Getenv("POSTGRES_DB")
	d.User = os.Getenv("POSTGRES_USER")
	d.Password = os.Getenv("POSTGRES_PASSWORD")
	d.Driver = os.Getenv("POSTGRES_DRIVER")
	d.Host = os.Getenv("POSTGRES_HOST")
	if d.Name == "" || d.Password == "" || d.User == "" || d.Driver == "" || d.Host == "" {
		log.Fatalf("Wrong credentials")
	}
}
