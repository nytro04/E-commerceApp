package database

import (
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/nytro04/nytroshopitems"
	"github.com/nytro04/nytroshop/users"
)

type DB interface {
	items.ItemDB
	users.UserDB
	Close() error
}

type postgresDB struct {
	db *sql.DB
}

func New(conn string) (DB, error) {
	// connecting to postgres db
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	// checking if connection is alive
	if err := db.Ping(); err != nil {
		return nil, err
	}

	pg := &postgresDB{db: db}
	err = pg.ensureTables()

	return pg, err
}

func (db *postgresDB) ensureTables() error {
	_, err := db.db.Exec("CREATE TABLE IF NOT EXISTS items (id SERIAL, name TEXT, price NUMERIC);")
	if err != nil {
		return err
	}

	_, err = db.db.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL, name TEXT, hash TEXT, address TEXT);")
	return err
}

func (db *postgresDB) Close() error {
	return db.db.Close()
}

func (db *postgresDB) CreateItem(item *items.Item) (int64, error) {
	result, err := db.db.Prepare("INSERT INTO items (name, price) VALUES ($1, $2) RETURNING id;")
	if err != nil {
		return 0, err
	}
	defer result.Close()

	var lastInsertId int64

	err = result.QueryRow(item.Name, item.PriceInCents).Scan(&lastInsertId)
	if err != nil {
		return 0, err
	}

	return lastInsertId, err
}

func (db *postgresDB) GetItemByID(id int64) (*items.Item, error) {
	row := db.db.QueryRow("SELECT id, name, price FROM items WHERE id= $1;", id)

	i := new(items.Item)
	err := row.Scan(&i.ID, &i.Name, &i.PriceInCents)
	return i, err
}

func (db *postgresDB) GetItemsByName(name string) ([]*items.Item, error) {
	rows, err := db.db.Query("SELECT id, name, price FROM items WHERE name= $1;", name)
	if err != nil {
		return nil, err
	}
	var itemsSlice []*items.Item
	for rows.Next() {
		i := new(items.Item)
		//check error here...
		err = rows.Scan(&i.ID, &i.Name, &i.PriceInCents)
		itemsSlice = append(itemsSlice, i)
	}
	return itemsSlice, err
}

func (db *postgresDB) GetAllItems() ([]*items.Item, error) {
	rows, err := db.db.Query("SELECT id, name, price FROM items")
	if err != nil {
		return nil, err
	}
	var allItemsSlice []*items.Item
	for rows.Next() {
		i := new(items.Item)
		err = rows.Scan(&i.ID, &i.Name, &i.PriceInCents)
		allItemsSlice = append(allItemsSlice, i)
	}
	return allItemsSlice, err
}

func (db *postgresDB) UpdateItem(item *items.Item) error {
	_, err := db.db.Exec("UPDATE items SET name = $2, price = $3 WHERE id = $1;", item.ID, item.Name, item.PriceInCents)
	return err

}

func (db *postgresDB) RemoveItem(id int64) error {
	_, err := db.db.Exec("DELETE FROM items WHERE id=$1;", id)
	return err
}

//user db methods
func (db *postgresDB) CreateUser(user *users.User) (int64, error) {
	result, err := db.db.Prepare("INSERT INTO users (name, hash, address) VALUES ($1, $2, $3) RETURNING id;")
	if err != nil {
		return 0, err
	}
	defer result.Close()

	var lastInsertId int64
	err = result.QueryRow(user.Name, user.Hash, user.Address).Scan(&lastInsertId)
	if err != nil {
		return 0, err
	}
	return lastInsertId, err
}

func (db *postgresDB) GetUserByID(id int64) (*users.User, error) {
	row := db.db.QueryRow("SELECT id, name, hash, address FROM users WHERE id=$1;", id)

	i := new(users.User)
	err := row.Scan(&i.ID, &i.Name, &i.Hash, &i.Address)
	return i, err
}

func (db *postgresDB) GetUserByName(name string) (*users.User, error) {
	row := db.db.QueryRow("SELECT id, name, hash, address FROM users WHERE name=$1;", name)

	i := new(users.User)
	err := row.Scan(&i.ID, &i.Name, &i.Hash, &i.Address)
	return i, err
}

func (db *postgresDB) GetAllUsers() ([]*users.User, error) {
	rows, err := db.db.Query("SELECT id, name, hash, address FROM users")
	if err != nil {
		return nil, err
	}
	var allUsersSlice []*users.User
	for rows.Next() {
		i := new(users.User)
		err = rows.Scan(&i.ID, &i.Name, &i.Hash, &i.Address)
		allUsersSlice = append(allUsersSlice, i)
	}
	return allUsersSlice, err
}

func (db *postgresDB) UpdateUser(user *users.User) error {
	_, err := db.db.Exec("UPDATE users SET name = $2, hash = $3, address = $4 WHERE id = $1;", user.ID, user.Name, user.Hash, user.Address)
	return err
}

func (db *postgresDB) RemoveUser(id int64) error {
	_, err := db.db.Exec("DELETE FROM users WHERE id=$1;", id)
	return err
}
