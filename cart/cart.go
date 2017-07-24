package cart

import "sync"

type Item struct {
	ItemID   int64
	Quantity int64
}

type Cart struct {

	UserID uint64
	Items  []*Item
}

type Manager struct {
	// Todo
	m           sync.RWMutex
	sessionCart map[SessionID]*Cart
	userCart    map[UserID]*Cart
}

func New() *Manager {
	// Todo

	return &Manager{sessionCart: make(map[SessionID]*Cart),
		userCart: make(map[UserID]*Cart),
	}
}

func (m *Manager) createCart(userID int64) *Cart {
	cart := &Cart{
		UserID: userID,
	}

	m.userCart[userID] = cart

	return cart
}

func (m *Manager) createSessionCart(sessionID string) *Cart {
	cart := &Cart{}
	m.sessionCart[sessionID] = cart

	return cart

}

func (c *Cart) addItem(itemID int64, quantity int) *Cart {
	i := findIndex(itemID, c.Items)


	if i == -1 {
		item := &Item{
			ItemID:    itemID,
			Quantity:   quantity,
		}

		c.Items = append(c.Items, item)
	} else {
		c.Items[i].Quantity += quantity

	// ...
}

func findIndex(itemID int64, items []Item) int {
	// Iterate over the items and check each one
	for i, item := range items {
		if item.ItemID == itemID {
			return i
		}
	}

	// None found, just return -1
	return -1
}


func  MergeCart(dst, src *Cart) *Cart {


	for _, item := range src.Items {
		dst.AddItem(itme.ItemID, item.Quantity)
	}

}


}


func (m *Manager) DeleteCart(userID int64) {
	deleteCart := &Manager {
									userCart: make(map[userID]*Cart),
						}
	delete(map[userID]*Cart, userID)

}


func (m *Manager) DeleteSessionCart(sessionID string) {

	delete(m.userCart, sessionID)

}



