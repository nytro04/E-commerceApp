package handler

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nytro04/nytroshop/database"
	"github.com/nytro04/nytroshop/items"
)

// Handler holds a reference to the database
// and the parsed HTML template in memory
type Handler struct {
	DB          database.DB
	indexTmpl   *template.Template
	newTmpl     *template.Template
	datailsTmpl *template.Template
}

// New creates the HTTP handler
func New(db database.DB) http.Handler {
	h := Handler{
		DB:          db,
		indexTmpl:   template.Must(template.ParseFiles("template/index.html")),
		// newTmpl:     template.Must(template.ParseFiles("template/new.html")),
		// datailsTmpl: template.Must(template.ParseFiles("template/new.html")),
	}

	router := mux.NewRouter()
	router.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("template/assets"))))

	router.HandleFunc("/", h.ShowAllItems).Methods("GET")
	// router.HandleFunc("/allItems/{id}", h.ShowItem).Methods("GET")
	// router.HandleFunc("/checkout", h.CheckOut).Methods("GET")

	return router
}

//
func (h *Handler) ShowAllItems(w http.ResponseWriter, r *http.Request) {

	allItems, err := h.DB.GetAllItems()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	h.indexTmpl.Execute(w, allItems)
}

func (h *Handler) CreateNewItem(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	price, _ := strconv.ParseInt(r.PostFormValue("price"), 10, 64)

	item := &items.Item{
		Name:         r.PostFormValue("name"),
		PriceInCents: price,
		Description:  r.PostFormValue("desc"),
		Image:        r.PostFormValue("img"),
	}

	var err error
	if item.ID, err = h.DB.CreateItem(item); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	h.newTmpl.Execute(w, item)
}
