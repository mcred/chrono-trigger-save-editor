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
		NameID:              storage.Attribute{root, 8},
		CharID:              storage.Attribute{root + 0x01, 8},
		HP:                  storage.Attribute{root + 0x03, 16},
		MaxHP:               storage.Attribute{root + 0x05, 16},
		MP:                  storage.Attribute{root + 0x07, 16},
		MaxMP:               storage.Attribute{root + 0x09, 16},
		BasePower:           storage.Attribute{root + 0x0B, 8},
		BaseStamina:         storage.Attribute{root + 0x0C, 8},
		BaseSpeed:           storage.Attribute{root + 0x0D, 8},
		BaseMagic:           storage.Attribute{root + 0x0E, 8},
		BaseHit:             storage.Attribute{root + 0x0F, 8},
		BaseEvade:           storage.Attribute{root + 0x10, 8},
		BaseMagicDefense:    storage.Attribute{root + 0x11, 8},
		Level:               storage.Attribute{root + 0x12, 8},
		Experience:          storage.Attribute{root + 0x13, 16}, //Maybe 24Bit
		Helmet:              storage.Attribute{root + 0x27, 8},
		Armor:               storage.Attribute{root + 0x28, 8},
		Weapon:              storage.Attribute{root + 0x29, 8},
		Relic:               storage.Attribute{root + 0x2A, 8},
		ExpToLevel:          storage.Attribute{root + 0x2B, 16},
		CurrentPower:        storage.Attribute{root + 0x36, 8},
		CurrentStamina:      storage.Attribute{root + 0x37, 8},
		CurrentSpeed:        storage.Attribute{root + 0x38, 8},
		CurrentMagic:        storage.Attribute{root + 0x39, 8},
		CurrentHit:          storage.Attribute{root + 0x3A, 8},
		CurrentEvade:        storage.Attribute{root + 0x3B, 8},
		CurrentMagicDefense: storage.Attribute{root + 0x3C, 8},
		CurrentAttack:       storage.Attribute{root + 0x3D, 8},
		CurrentDefense:      storage.Attribute{root + 0x3E, 8},
		CurrentMaxHP:        storage.Attribute{root + 0x3F, 16},
	}
}
