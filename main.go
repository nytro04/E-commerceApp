package main

import (
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/nytro04/nytroshop/database"
	"github.com/nytro04/nytroshop/items"
	"github.com/nytro04/nytroshop/session"
	"github.com/nytro04/nytroshop/users"
)

type itemPage struct {
	Title string
	Items []*items.Item
}

func main() {
	rand.Seed(time.Now().Unix())

	log.Println("Starting NytroShop")

	log.Println("Creating session manager")
	sessionStore := session.NewSessionStore()

	log.Println("Connecting to database")
	db, err := database.New("user=nytro password=superman host=localhost dbname=shop sslmode=disable")
	if err != nil {
		log.Fatalf("Error while connecting to the database: %s\n", err)
	}
	defer db.Close()

	_, err = users.CreateUser(db, "nytro", "password123", "")
	if err != nil {
		log.Fatal(err)
	}

	_, err = items.CreateItem(db, "item1", 50)
	if err != nil {
		log.Fatal(err)
	}

	var templates *template.Template

	templates, err = template.ParseGlob("tmpl/*")
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static/"))))

	http.HandleFunc("/", sessionStore.Wrap(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	}))

	http.HandleFunc("/items/all", func(w http.ResponseWriter, _ *http.Request) {
		data := &itemPage{
			Title: "All Items",
		}
		data.Items, err = db.GetAllItems()
		if err != nil {
			log.Fatal(err)
		}

		templates.ExecuteTemplate(w, "items_list.gohtml", data)
	})

	log.Fatalln(http.ListenAndServe(":8007", nil))
}

//	userManager := users.NewManager(db)
//	sessionManager := sessions.NewManager(db)
//	cartManager := cart.cartNewManager(db, userManager, sessionManager)
//frontend.Start(db, userManager, sessionManager, cartManager)
