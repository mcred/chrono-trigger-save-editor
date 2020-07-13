package inventory

// Swords : Inventory list of available Swords for Crono
func Swords() Inventory {
	return []Item{
		{0x01, "Wood Sword"},
		{0x02, "Iron Blade"},
		{0x03, "Steel Saber"},
		{0x04, "Lode Sword"},
		{0x05, "Red Katana"},
		{0x06, "Flint Edge"},
		{0x07, "Dark saber"},
		{0x08, "Aeon Blade"},
		{0x09, "Demon Edge"},
		{0x0A, "AlloyBlade"},
		{0x0B, "star Sword"},
		{0x0C, "Vedic Blade"},
		{0x0D, "Kali Blade"},
		{0x0E, "Shiva Sword"},
		{0x0F, "Bolt Sword"},
		{0x10, "Slasher"},
		{0x4F, "Mop"},
		{0x53, "Swallow"},
		{0x54, "Slasher 2"},
		{0x55, "Rainbow"},
	}
}

// Bows : Inventory list of available Bows for Marle
func Bows() Inventory {
	return []Item{
		{0x11, "Bronze Bow"},
		{0x12, "Iron Bow"},
		{0x13, "Lode Bow"},
		{0x14, "Robin Bow"},
		{0x15, "Sage Bow"},
		{0x16, "Dream Bow"},
		{0x17, "CometArrow"},
		{0x18, "SonicArrow"},
		{0x19, "Valkerye"},
		{0x1A, "Siren"},
	}
}

// Guns : Inventory list of available Guns for Lucca
func Guns() Inventory {
	return []Item{
		{0x1F, "Air Gun"},
		{0x20, "Dart Gun"},
		{0x21, "Auto Gun"},
		{0x22, "PicoMagnum"},
		{0x23, "Plasma Gun"},
		{0x24, "Ruby Gun"},
		{0x25, "Dream Gun"},
		{0x26, "Megablast"},
		{0x27, "Shock Wave"},
		{0x28, "Wonder Shot"},
		{0x29, "Graedus"},
	}
}

// Arms : Inventory list of available Arms for Robo
func Arms() Inventory {
	return []Item{
		{0x2E, "Tin Arm"},
		{0x2F, "Hammer Arm"},
		{0x30, "Mirage Hand"},
		{0x31, "Stone Arm"},
		{0x32, "Doom Finger"},
		{0x33, "Magma Hand"},
		{0x34, "MegatonArm"},
		{0x35, "Big hand"},
		{0x36, "Kaiser Arm"},
		{0x37, "Giga Arm"},
		{0x38, "Terra Arm"},
		{0x39, "Crisis Arm"},
	}
}

// Blades : Inventory list of available Blades for Frog
func Blades() Inventory {
	return []Item{
		{0x3B, "Bronze Edge"},
		{0x3C, "Iron Sword"},
		{0x3D, "Masamune"},
		{0x3E, "Flash Blade"},
		{0x3F, "Pearl Edge"},
		{0x40, "Rune Blade"},
		{0x41, "Brave Sword"},
		{0x42, "Masamune"},
		{0x43, "Demon Hit"},
		{0x50, "Bent Sword"},
		{0x51, "Bent Hilt"},
		{0x52, "Masamune"},
	}
}

// Fists : Inventory list of available Fists for Ayla
func Fists() Inventory {
	return []Item{
		{0x44, "Fist (I)"},
		{0x45, "Fist (II)"},
		{0x46, "Fist (III)"},
		{0x47, "Iron Fist"},
		{0x48, "Bronze Fist"},
	}
}

// Scythes : Inventory list of Scythes Blades for Magus
func Scythes() Inventory {
	return []Item{
		{0x4B, "DarkScythe"},
		{0x4C, "Hurricane"},
		{0x4D, "StarScythe"},
		{0x4E, "DoomSickle"},
	}
}
