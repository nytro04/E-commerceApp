package cart

import "sync"

type Item struct {
	ItemID   int64
	Quantity int64
}

type Cart struct {
	UserID int64
	Items  []*Item
}

type Manager struct {
	// Todo
	m           sync.RWMutex
	sessionCart map[string]*Cart
	userCart    map[int64]*Cart
}

func New() *Manager {
	return &Manager{sessionCart: make(map[string]*Cart),
		userCart: make(map[int64]*Cart),
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

func (c *Cart) addItem(itemID int64, quantity int64) {
	i := findIndex(itemID, c.Items)

	if i == -1 {
		item := &Item{
			ItemID:   itemID,
			Quantity: quantity,
		}

		c.Items = append(c.Items, item)
	} else {
		c.Items[i].Quantity += quantity
	}
	// ...
}

func findIndex(itemID int64, items []*Item) int {
	// Iterate over the items and check each one
	for i, item := range items {
		if item.ItemID == itemID {
			return i
		}
	}

	// None found, just return -1
	return -1
}

func MergeCarts(dst, src *Cart) {
	for _, item := range src.Items {
		dst.addItem(item.ItemID, item.Quantity)
	}

}

func (m *Manager) DeleteCart(userID int64) {
	delete(m.userCart, userID)
}

func (m *Manager) DeleteSessionCart(sessionID string) {
	delete(m.sessionCart, sessionID)
}
