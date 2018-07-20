package items

//item properties
type Item struct {
	ID           int64
	Name         string
	PriceInCents int64
	Description  string
	Image		string
}

// item db interface
type ItemDB interface {
	CreateItem(*Item) (int64, error)
	GetItemByID(id int64) (*Item, error)
	GetItemsByName(name string) ([]*Item, error)
	GetAllItems() ([]*Item, error)
	UpdateItem(item *Item) error
	RemoveItem(id int64) error
}

//create item func
func CreateItem(db ItemDB, name string, price int64, description, image string,) (*Item, error) {
	var err error
	item := Item{
		Name:         name,
		PriceInCents: price,
		Description:	description,
		Image:		image,
	}

	item.ID, err = db.CreateItem(&item)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

type ByID []*Item

func (by ByID) Len() int {
	return len(by)
}

func (by ByID) Swap(a, b int) {
	by[a], by[b] = by[b], by[a]
}

func (by ByID) Less(a, b int) bool {
	return by[a].ID < by[b].ID
}

type ByName []*Item

func (n ByName) Len() int {
	return len(n)
}

func (n ByName) Swap(a, b int) {
	n[a], n[b] = n[b], n[a]
}

func (n ByName) Less(a, b int) bool {
	return n[a].Name < n[b].Name
}

type ByPrice []*Item

func (p ByPrice) Len() int {
	return len(p)
}

func (p ByPrice) Swap(a, b int) {
	p[a], p[b] = p[b], p[a]
}

func (p ByPrice) Less(a, b int) bool {
	return p[a].PriceInCents < p[b].PriceInCents
}
