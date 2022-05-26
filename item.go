package main

type Item struct {
	name   string
	weight float64 // in kg...
}

func NewItem(name string, weight float64) *Item {
	return &Item{name: name, weight: weight}
}

// Add put item together, if name is the same
func (i *Item) Add(ni *Item) bool {
	if i.name == ni.name {
		i.weight += ni.weight
		return true
	}
	return false
}

// Take some of the item out of the existing item
func (i *Item) Take(w float64) *Item {
	if w > 0 && w <= i.weight {
		i.weight -= w
		return NewItem(i.name, w)
	}
	return NewItem(i.name, 0)
}

func (i *Item) Empty() bool {
	return i.weight == 0
}
