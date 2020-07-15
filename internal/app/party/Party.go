package party

import (
	"ChronoTrigger/internal/app/inventory"
	"ChronoTrigger/internal/app/storage"
)

// Item : struct party inventory item structure
type Item struct {
	Name  string
	Count uint
}

// Party : struct with editable game attributes
type Party struct {
	Inventory []Item
	//000 | 256 | Item Type
	//100 | 256 | Item Count
	Member1   storage.Attribute
	Member2   storage.Attribute
	Member3   storage.Attribute
	SaveCount storage.Attribute
	Gold      storage.Attribute
	//5B0 | 48  | Names (8 slots, 6 byte C-style strings)
	MilliSeconds storage.Attribute
	Seconds      storage.Attribute
	Minutes      storage.Attribute
	Hours        storage.Attribute
	World        storage.Attribute
	PosX         storage.Attribute
	PosY         storage.Attribute
}

// GetParty : Party locations of party attributes in the save file
func GetParty() Party {
	return Party{
		Inventory:    []Item{},
		Member1:      storage.Attribute{0x580, 1},
		Member2:      storage.Attribute{0x581, 1},
		Member3:      storage.Attribute{0x582, 1},
		SaveCount:    storage.Attribute{0x59C, 1},
		Gold:         storage.Attribute{0x5E0, 3},
		MilliSeconds: storage.Attribute{0x5E3, 1},
		Seconds:      storage.Attribute{0x5E4, 1},
		Minutes:      storage.Attribute{0x5E5, 1},
		Hours:        storage.Attribute{0x5E6, 1},
		World:        storage.Attribute{0x5F3, 2},
		PosX:         storage.Attribute{0x5F5, 1},
		PosY:         storage.Attribute{0x5F6, 1},
	}
}

func (p *Party) LoadInventory(g storage.Game) {
	for i := 0x0; i <= 0xFF; i++ {
		v := g.GetValue(storage.Attribute{i, 1})
		if v != 0 {
			item := Item{
				inventory.AllItems().GetValByID(v),
				g.GetValue(storage.Attribute{i + 0x100, 1}),
			}
			p.Inventory = append(p.Inventory, item)
		}
	}
}
