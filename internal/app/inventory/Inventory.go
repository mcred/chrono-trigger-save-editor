package inventory

// Item : Basic struct for Inventory Item
type Item struct {
	ID   int
	Name string
}

// Inventory : splice of Inventory Items
type Inventory []Item

// GetValByID : Get Inventory Item Name by ID
func (i Inventory) GetValByID(id int) string {
	for _, item := range i {
		if item.ID == id {
			return item.Name
		}
	}
	return "Item Not Found"
}

// GetIDByVal : Get Inventory Item ID by Name
func (i Inventory) GetIDByVal(val string) int {
	for _, item := range i {
		if item.Name == val {
			return item.ID
		}
	}
	return 0
}

// GetVals : Get splice of Inventory Item Names
func (i Inventory) GetVals() []string {
	var v []string
	for _, item := range i {
		v = append(v, item.Name)
	}
	return v
}

// AllItems : Inventory list of all Items
func AllItems() Inventory {
	var items []Item
	for _, item := range Swords() {
		items = append(items, item)
	}
	for _, item := range Bows() {
		items = append(items, item)
	}
	for _, item := range Guns() {
		items = append(items, item)
	}
	for _, item := range Arms() {
		items = append(items, item)
	}
	for _, item := range Blades() {
		items = append(items, item)
	}
	for _, item := range Fists() {
		items = append(items, item)
	}
	for _, item := range Scythes() {
		items = append(items, item)
	}
	for _, item := range Helmets() {
		items = append(items, item)
	}
	for _, item := range Arms() {
		items = append(items, item)
	}
	for _, item := range Relics() {
		items = append(items, item)
	}
	for _, item := range Consumables() {
		items = append(items, item)
	}
	for _, item := range Armors() {
		items = append(items, item)
	}
	return items
}
