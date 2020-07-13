package characters

import "ChronoTrigger/internal/app/storage"

// Character : struct with editable game attributes
type Character struct {
	NameID              storage.Attribute
	CharID              storage.Attribute
	HP                  storage.Attribute
	MaxHP               storage.Attribute
	MP                  storage.Attribute
	MaxMP            storage.Attribute
	BasePower        storage.Attribute
	BaseStamina      storage.Attribute
	BaseSpeed        storage.Attribute
	BaseMagic        storage.Attribute
	BaseHit          storage.Attribute
	BaseEvade        storage.Attribute
	BaseMagicDefense storage.Attribute
	Level            storage.Attribute
	Experience       storage.Attribute
	Helmet           storage.Attribute
	Armor            storage.Attribute
	Weapon           storage.Attribute
	Equip4           storage.Attribute
	ExpToLevel       storage.Attribute
	CurrentPower     storage.Attribute
	CurrentStamina   storage.Attribute
	CurrentSpeed     storage.Attribute
	CurrentMagic     storage.Attribute
	CurrentHit       storage.Attribute
	CurrentEvade     storage.Attribute
	CurrentMagicDefense storage.Attribute
	CurrentAttack       storage.Attribute
	CurrentDefense      storage.Attribute
	CurrentMaxHP        storage.Attribute
}

// CreateCharacter : based off a root position in the save file
func CreateCharacter(root int) Character {
	return Character{
		NameID:              storage.Attribute{root, false},
		CharID:              storage.Attribute{root + 0x01, false},
		HP:                  storage.Attribute{root + 0x03, true},
		MaxHP:               storage.Attribute{root + 0x05, true},
		MP:                  storage.Attribute{root + 0x07, true},
		MaxMP:            storage.Attribute{root + 0x09, true},
		BasePower:        storage.Attribute{root + 0x0B, false},
		BaseStamina:      storage.Attribute{root + 0x0C, false},
		BaseSpeed:        storage.Attribute{root + 0x0D, false},
		BaseMagic:        storage.Attribute{root + 0x0E, false},
		BaseHit:          storage.Attribute{root + 0x0F, false},
		BaseEvade:        storage.Attribute{root + 0x10, false},
		BaseMagicDefense: storage.Attribute{root + 0x11, false},
		Level:            storage.Attribute{root + 0x12, false},
		Experience:       storage.Attribute{root + 0x13, true}, //Maybe 24Bit
		Helmet:           storage.Attribute{root + 0x27, false},
		Armor:            storage.Attribute{root + 0x28, false},
		Weapon:           storage.Attribute{root + 0x29, false},
		Equip4:           storage.Attribute{root + 0x2A, false},
		ExpToLevel:       storage.Attribute{root + 0x2B, true},
		CurrentPower:     storage.Attribute{root + 0x36, false},
		CurrentStamina:   storage.Attribute{root + 0x37, false},
		CurrentSpeed:     storage.Attribute{root + 0x38, false},
		CurrentMagic:     storage.Attribute{root + 0x39, false},
		CurrentHit:       storage.Attribute{root + 0x3A, false},
		CurrentEvade:     storage.Attribute{root + 0x3B, false},
		CurrentMagicDefense: storage.Attribute{root + 0x3C, false},
		CurrentAttack:       storage.Attribute{root + 0x3D, false},
		CurrentDefense:      storage.Attribute{root + 0x3E, false},
		CurrentMaxHP:        storage.Attribute{root + 0x3F, true},
	}
}
