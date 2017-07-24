package database

import "database/sql"

type ItemDB interface {
	CreateItem(item items.Item) (int64, error)
	GetItemByID(id int64) (*items.Item, error)
	GetItemsByName(name string) ([]*items.Item, error)
	GetAllItems() ([]*items.Item, error)
	UpdateItem(item *items.Item) error
	RemoveItem(id int64) error
}

type UserDB interface {
	CreateUser(user *users.User) (int64, error)
	GetUserByID(id int64) (*users.User, error)
	GetUserByName(name string) (*users.User, error)
	GetAllUsers() ([]*users.Users, error)
	UpdateUser(user *users.User) error
	RemoveUser(id int64) error
}

type DB interface {
	ItemDB
	UserDB
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

	pg = &postgresDB{db: db}
	err = pg.ensureTables()

	return pg, err
}

func (db *postgreDB) ensureTables() error {
	_, err := db.db.Exec("CREATE TABLE IF NOT EXISTS items (id SERIAL, name TEXT, price NUMERIC);")
	if err != nil {
		return err
	}

	_, err = db.db.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL, name TEXT, hash TEXT, address TEXT);")
	return err
}

func (db *postgreDB) CreateItem(item items.Item) (int64, error) {
	result, err := db.db.Exec("INSERT INTO items (name, price) VALUES ($1, $2);", item.Name, item.Price)
	if err != nil {
		return 0, err
	}
	return resule.LastInsertId()
}

func (db *postgreDB) GetItemByID(id int64) (*items.Item, error) {
	row, err := db.db.QueryRow("SELECT id, name, price FROM items WHERE id= $1;", id)
	if err != nil {
		return nil, err
	}
	i := new(items.Item)
	err = row.Scan(&i.ID, &i.Name, &i.PriceInCents)
	return i, err
}

func (db *postgreDB) GetItemsByName(name string) ([]*items.Item, error) {
	rows, err := db.db.Query("SELECT id, name, price FROM items WHERE name= $1;", name)
	if err != nil {
		return nil, err
	}
	var itemsSlice *items.Item
	for rows.Next() {
		i := new(items.Item)
		err = rows.Scan(&i.ID, &i.Name, &i.PriceInCents)
		itemsSlice = append(itemsSlice, i)
	}
	return itemsSlice, err
}

func (db *postgreDB) GetAllItems() ([]*items.Item, error) {
	rows, err := db.db.Query("SELECT id, name, price FROM items")
	if err != nil {
		return nil, err
	}
	var allItemsSlice []*items.Item
	for rows.Next() {
		i := new(items.Item)
		err = rows.Scan(&i.ID, &i.Name, &i.PriceInCents)
		allItemsSlice = append(allItemSlice, i)
	}
	return allItemsSlice, err
}

func (db *postgreDB) UpdateItem(item *items.Item) error {
	_, err := db.db.Exec("UPDATE items SET name = $2, price = $3 WHERE id = $1;", item.ID, item.Name, item.PriceInCents)
	return err

}

func (db *postgreDB) RemoveItem(id int64) error {
	_, err := db.db.Exec("DELETE FROM items WHERE id=$1;", id)
	return err
}

//user db methods
func (db *postgresDB) CreateUser(user *users.User) (int64, error) {
	result, err := db.db.Exec("INSERT INTO users (name, hash, address) VALUES ($1, $2, $3);", user.Name, user.Hash, user.Address)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId(), nil
}

func (db *postgresDB) GetUserByID(id int64) (*users.User, error) {
	row, err := db.db.QueryRow("SELECT id, name, hash, address FROM users WHERE id=$1;", id)
	if err != nil {
		return nil, err
	}
	i := new(users.User)
	err = row.Scan(&i.ID, &i.Name, &i.Hash, &i.Address)
	return i, err
}

func (db *postgresDB) GetUserByName(name string) (*users.User, error) {
	row, err := db.db.QueryRow("SELECT id, name, hash, address FROM users WHERE name=$1;", name)
	if err != nil {
		return nil, err
	}
	i := new(users.User)
	err = row.Scan(&i.ID, &i.Name, &i.Hash, &i.Address)
	return i, err
}

func (db *postgresDB) GetAllUsers() ([]*users.Users, error) {
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
