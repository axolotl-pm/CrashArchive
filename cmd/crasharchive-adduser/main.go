package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/pmmp/CrashArchive/app"
	"github.com/pmmp/CrashArchive/app/database"
	"github.com/pmmp/CrashArchive/app/user"
)

func main() {
	log.SetFlags(log.Lshortfile)

	configPath := flag.String("c", "./config/config.json", "path to `config.json`")
	username := flag.String("u", "", "username of new user to add")
	password := flag.String("p", "", "password of new user")
	flag.Parse()

	var err error
	config, err := app.LoadConfig(*configPath)
	if err != nil {
		log.Printf("unable to load config: %v", err)
		os.Exit(1)
	}
	if *username == "" || *password == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	db, err := database.Connect(config.Database, 5*time.Second)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	db.AddUser(*username, []byte(*password), user.Admin)
	log.Printf("successfully added new user to database")
}
