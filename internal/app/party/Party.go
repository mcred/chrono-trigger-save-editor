package party

import (
	"ChronoTrigger/internal/app/inventory"
	"encoding/binary"
	"github.com/mcred/savetools"
)

// Item : struct party inventory item structure
type Item struct {
	Name  string
	Count int
}

// Party : struct with editable game attributes
type Party struct {
	Inventory []Item
	//000 | 256 | Item Type
	//100 | 256 | Item Count
	Member1   savetools.Attribute
	Member2   savetools.Attribute
	Member3   savetools.Attribute
	SaveCount savetools.Attribute
	Gold      savetools.Attribute
	//5B0 | 48  | Names (8 slots, 6 byte C-style strings)
	MilliSeconds savetools.Attribute
	Seconds      savetools.Attribute
	Minutes      savetools.Attribute
	Hours        savetools.Attribute
	World        savetools.Attribute
	PosX         savetools.Attribute
	PosY         savetools.Attribute
}

// GetParty : Party locations of party attributes in the save file
func GetParty() Party {
	return Party{
		Inventory:    []Item{},
		Member1:      savetools.Attribute{0x580, 8, binary.LittleEndian},
		Member2:      savetools.Attribute{0x581, 8, binary.LittleEndian},
		Member3:      savetools.Attribute{0x582, 8, binary.LittleEndian},
		SaveCount:    savetools.Attribute{0x59C, 8, binary.LittleEndian},
		Gold:         savetools.Attribute{0x5E0, 8, binary.LittleEndian},
		MilliSeconds: savetools.Attribute{0x5E3, 8, binary.LittleEndian},
		Seconds:      savetools.Attribute{0x5E4, 8, binary.LittleEndian},
		Minutes:      savetools.Attribute{0x5E5, 8, binary.LittleEndian},
		Hours:        savetools.Attribute{0x5E6, 8, binary.LittleEndian},
		World:        savetools.Attribute{0x5F3, 8, binary.LittleEndian},
		PosX:         savetools.Attribute{0x5F5, 8, binary.LittleEndian},
		PosY:         savetools.Attribute{0x5F6, 8, binary.LittleEndian},
	}
}

// LoadInventory : Load Inventory for Party
func (p *Party) LoadInventory(c savetools.Card) {
	for i := 0x0; i <= 0xFF; i++ {
		v := c.GetValue(savetools.Attribute{i, 8, binary.LittleEndian})
		if v != 0 {
			item := Item{
				inventory.AllItems().GetValByID(v),
				c.GetValue(savetools.Attribute{i + 0x100, 8, binary.LittleEndian}),
			}
			p.Inventory = append(p.Inventory, item)
		}
	}
}
