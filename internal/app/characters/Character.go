package characters

import (
	"encoding/binary"
	"github.com/mcred/savetools"
)

// Character : struct with editable game attributes
type Character struct {
	NameID              savetools.Attribute
	CharID              savetools.Attribute
	HP                  savetools.Attribute
	MaxHP               savetools.Attribute
	MP                  savetools.Attribute
	MaxMP               savetools.Attribute
	BasePower           savetools.Attribute
	BaseStamina         savetools.Attribute
	BaseSpeed           savetools.Attribute
	BaseMagic           savetools.Attribute
	BaseHit             savetools.Attribute
	BaseEvade           savetools.Attribute
	BaseMagicDefense    savetools.Attribute
	Level               savetools.Attribute
	Experience          savetools.Attribute
	Helmet              savetools.Attribute
	Armor               savetools.Attribute
	Weapon              savetools.Attribute
	Relic               savetools.Attribute
	ExpToLevel          savetools.Attribute
	CurrentPower        savetools.Attribute
	CurrentStamina      savetools.Attribute
	CurrentSpeed        savetools.Attribute
	CurrentMagic        savetools.Attribute
	CurrentHit          savetools.Attribute
	CurrentEvade        savetools.Attribute
	CurrentMagicDefense savetools.Attribute
	CurrentAttack       savetools.Attribute
	CurrentDefense      savetools.Attribute
	CurrentMaxHP        savetools.Attribute
}

func gen8Bit(l int) savetools.Attribute {
	return savetools.Attribute{
		Location:   l,
		Bits:       8,
		Endianness: binary.LittleEndian,
	}
}

func gen16Bit(l int) savetools.Attribute {
	return savetools.Attribute{
		Location:   l,
		Bits:       16,
		Endianness: binary.LittleEndian,
	}
}

// CreateCharacter : Character based off a root position in the save file
func CreateCharacter(root int) Character {
	return Character{
		NameID:           gen8Bit(root),
		CharID:           gen8Bit(root + 0x01),
		HP:               gen16Bit(root + 0x03),
		MaxHP:            gen16Bit(root + 0x05),
		MP:               gen16Bit(root + 0x07),
		MaxMP:            gen16Bit(root + 0x09),
		BasePower:        gen8Bit(root + 0x0B),
		BaseStamina:      gen8Bit(root + 0x0C),
		BaseSpeed:        gen8Bit(root + 0x0D),
		BaseMagic:        gen8Bit(root + 0x0E),
		BaseHit:          gen8Bit(root + 0x0F),
		BaseEvade:        gen8Bit(root + 0x10),
		BaseMagicDefense: gen8Bit(root + 0x11),
		Level:            gen8Bit(root + 0x12),
		//Experience:          gen8Bit(root + 0x13, 3},
		Helmet:              gen8Bit(root + 0x27),
		Armor:               gen8Bit(root + 0x28),
		Weapon:              gen8Bit(root + 0x29),
		Relic:               gen8Bit(root + 0x2A),
		ExpToLevel:          gen16Bit(root + 0x2B),
		CurrentPower:        gen8Bit(root + 0x36),
		CurrentStamina:      gen8Bit(root + 0x37),
		CurrentSpeed:        gen8Bit(root + 0x38),
		CurrentMagic:        gen8Bit(root + 0x39),
		CurrentHit:          gen8Bit(root + 0x3A),
		CurrentEvade:        gen8Bit(root + 0x3B),
		CurrentMagicDefense: gen8Bit(root + 0x3C),
		CurrentAttack:       gen8Bit(root + 0x3D),
		CurrentDefense:      gen8Bit(root + 0x3E),
		CurrentMaxHP:        gen16Bit(root + 0x3F),
	}
}
