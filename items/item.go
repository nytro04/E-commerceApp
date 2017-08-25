package items

//item properties
type Item struct {
	ID           uint64
	Name         string
	PriceInCents int
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