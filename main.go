package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/nytro04/nytroshop/database"
	"github.com/nytro04/nytroshop/handlers"
)

var Templates *template.Template

func main() {

	var port int
	flag.IntVar(&port, "p", 8888, "HTTP port")
	flag.Parse()

	rand.Seed(time.Now().Unix())

	log.Println("Starting NytroShop")
	log.Println("Creating session manager")
	// sessionStore := session.NewSessionStore()

	log.Println("Connecting to database")
	db, err := database.New("user=francisbadasu host=localhost dbname=shoppingcart sslmode=disable")
	if err != nil {
		log.Fatalf("Error while connecting to the database: %s\n", err)
	}
	defer db.Close()
	fmt.Println("database ok!!")

	// Loading templates
	templ, err := Templates.ParseGlob("template/*.gohtml")
	if err != nil {
		log.Fatalf("Error parsing templates: %s\n", err)
	}

	// Initialize handlers
	fmt.Println("Initializing handlers...")
	h := handlers.New(db, templ)

	// Set up individuals routes
	fmt.Println("Setting up routes")
	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./templates"))))

	mux.HandleFunc("/", h.ShowAllItems)

	log.Printf("staring the server on port %d\n", port)
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%d", port), mux))


	//	userManager := users.NewManager(db)
	//	sessionManager := sessions.NewManager(db)
	//	cartManager := cart.cartNewManager(db, userManager, sessionManager)
	//frontend.Start(db, userManager, sessionManager, cartManager)
}
