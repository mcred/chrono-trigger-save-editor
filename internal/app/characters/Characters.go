package characters

import "ChronoTrigger/internal/app/inventory"

// Crono : Character return Crono data
func Crono() Character {
	return CreateCharacter(0x200, inventory.Swords())
}

// Marle : Character return Marle data
func Marle() Character {
	return CreateCharacter(0x250, inventory.Bows())
}

// Lucca : Character return Lucca data
func Lucca() Character {
	return CreateCharacter(0x2A0, inventory.Guns())
}

// Robo : Character return Robo data
func Robo() Character {
	return CreateCharacter(0x2F0, inventory.Arms())
}

// Frog : Character return Frog data
func Frog() Character {
	return CreateCharacter(0x340, inventory.Blades())
}

// Ayla : Character return Ayla data
func Ayla() Character {
	return CreateCharacter(0x390, inventory.Fists())
}

// Magus : Character return Magus data
func Magus() Character {
	return CreateCharacter(0x3E0, inventory.Scythes())
}

// Epoch : Character return Epoch data
func Epoch() Character {
	return CreateCharacter(0x430, []inventory.Item{})
}
