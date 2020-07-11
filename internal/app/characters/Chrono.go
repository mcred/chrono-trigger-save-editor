package characters

type Attribute struct {
	Location int
	Bytes int
	Reversed bool
}

type Character struct {
	HP Attribute
	HPMax Attribute
	MP Attribute
	MPMax Attribute
}

func Chrono() Character {
	return Character{
		HP:Attribute{0x203, 2, true},
		HPMax:Attribute{0x205, 2, true},
		MP:Attribute{0x207, 2, true},
		MPMax:Attribute{0x209, 2, true},
	}
}