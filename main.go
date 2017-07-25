package main

import (
	"log"
	
	"github.com/nytro04/nytroshop/session"
	"github.com/nytro04/nytroshop/database"
	_ "github.com/nytro04/nytroshop/users"
	"time"
	"math/rand"
)

func main() {
	rand.Seed(time.Now().Unix())
	
	log.Println("Starting NytroShop")
	
	log.Println("Creating session manager")
	sessionStore := session.NewSessionStore()
	
	log.Println("Connecting to database")
	db, err := database.New("postgres://postgres:superM@n04@localhost/shopdb")
	if err != nil {
		log.Fatalf("Error while connecting to the database: %s\n", err)
	}
	defer db.Close()
	
	_, _ = db, sessionStore
//	userManager := users.NewManager(db)
//	sessionManager := sessions.NewManager(db)
//	cartManager := cart.cartNewManager(db, userManager, sessionManager)
	//frontend.Start(db, userManager, sessionManager, cartManager)
	

}
