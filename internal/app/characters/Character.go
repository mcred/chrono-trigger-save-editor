package characters

import "ChronoTrigger/internal/app/storage"

// Character : struct with editable game attributes
type Character struct {
	NameID              storage.Attribute
	CharID              storage.Attribute
	HP                  storage.Attribute
	MaxHP               storage.Attribute
	MP                  storage.Attribute
	MaxMP               storage.Attribute
	BasePower           storage.Attribute
	BaseStamina         storage.Attribute
	BaseSpeed           storage.Attribute
	BaseMagic           storage.Attribute
	BaseHit             storage.Attribute
	BaseEvade           storage.Attribute
	BaseMagicDefense    storage.Attribute
	Level               storage.Attribute
	Experience          storage.Attribute
	Helmet              storage.Attribute
	Armor               storage.Attribute
	Weapon              storage.Attribute
	Relic               storage.Attribute
	ExpToLevel          storage.Attribute
	CurrentPower        storage.Attribute
	CurrentStamina      storage.Attribute
	CurrentSpeed        storage.Attribute
	CurrentMagic        storage.Attribute
	CurrentHit          storage.Attribute
	CurrentEvade        storage.Attribute
	CurrentMagicDefense storage.Attribute
	CurrentAttack       storage.Attribute
	CurrentDefense      storage.Attribute
	CurrentMaxHP        storage.Attribute
}

// CreateCharacter : Character based off a root position in the save file
func CreateCharacter(root int) Character {
	return Character{
		NameID:              storage.Attribute{root, 1},
		CharID:              storage.Attribute{root + 0x01, 1},
		HP:                  storage.Attribute{root + 0x03, 2},
		MaxHP:               storage.Attribute{root + 0x05, 2},
		MP:                  storage.Attribute{root + 0x07, 2},
		MaxMP:               storage.Attribute{root + 0x09, 2},
		BasePower:           storage.Attribute{root + 0x0B, 1},
		BaseStamina:         storage.Attribute{root + 0x0C, 1},
		BaseSpeed:           storage.Attribute{root + 0x0D, 1},
		BaseMagic:           storage.Attribute{root + 0x0E, 1},
		BaseHit:             storage.Attribute{root + 0x0F, 1},
		BaseEvade:           storage.Attribute{root + 0x10, 1},
		BaseMagicDefense:    storage.Attribute{root + 0x11, 1},
		Level:               storage.Attribute{root + 0x12, 1},
		Experience:          storage.Attribute{root + 0x13, 3},
		Helmet:              storage.Attribute{root + 0x27, 1},
		Armor:               storage.Attribute{root + 0x28, 1},
		Weapon:              storage.Attribute{root + 0x29, 1},
		Relic:               storage.Attribute{root + 0x2A, 1},
		ExpToLevel:          storage.Attribute{root + 0x2B, 2},
		CurrentPower:        storage.Attribute{root + 0x36, 1},
		CurrentStamina:      storage.Attribute{root + 0x37, 1},
		CurrentSpeed:        storage.Attribute{root + 0x38, 1},
		CurrentMagic:        storage.Attribute{root + 0x39, 1},
		CurrentHit:          storage.Attribute{root + 0x3A, 1},
		CurrentEvade:        storage.Attribute{root + 0x3B, 1},
		CurrentMagicDefense: storage.Attribute{root + 0x3C, 1},
		CurrentAttack:       storage.Attribute{root + 0x3D, 1},
		CurrentDefense:      storage.Attribute{root + 0x3E, 1},
		CurrentMaxHP:        storage.Attribute{root + 0x3F, 2},
	}
}
