package handlers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/nytro04/nytroshop/database"

	"github.com/nytro04/nytroshop/items"
)

type Handlers struct {
	db   database.DB
	tmpl *template.Template
}

func New(db database.DB, tmpl *template.Template) *Handlers {
	return &Handlers{
		db:   db,
		tmpl: tmpl,
	}
}

func (h Handlers) ShowAllItems(w http.ResponseWriter, r *http.Request) {

	allItems, err := h.db.GetAllItems()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html;  charset=utf-8")
	h.tmpl.ExecuteTemplate(w, "index.gohtml", allItems)
}

func (h Handlers) CreateNewItem(w http.ResponseWriter, r *http.Request) {
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
	if item.ID, err = h.db.CreateItem(item); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	h.tmpl.ExecuteTemplate(w, "new.gohtml", item)
}
