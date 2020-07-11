package characters

import (
	"ChronoTrigger/internal/app/storage"
)

type Character struct {
	HP storage.Attribute
	//HPMax storage.Attribute
	//MP storage.Attribute
	//MPMax storage.Attribute
	//ATK storage.Attribute
	//DEF storage.Attribute
	//PWR storage.Attribute
	//SPD storage.Attribute
	//HIT storage.Attribute
	//EV storage.Attribute
	//MAG storage.Attribute
	//STAM storage.Attribute
	//MDEF storage.Attribute
	//EXP storage.Attribute
}

func Chrono() Character {
	return Character{
		HP:    storage.Attribute{0x203, true},
		//HPMax: storage.Attribute{0x205, true},
		//MP: storage.Attribute{0x207, true},
		//MPMax: storage.Attribute{0x209, true},
		//ATK: storage.Attribute{},
		//DEF: storage.Attribute{},
		//PWR: storage.Attribute},
		//SPD: storage.Attribute{},
		//HIT: storage.Attribute{},
		//EV: storage.Attribute{},
		//MAG: storage.Attribute{},
		//STAM: storage.Attribute{},
		//MDEF: storage.Attribute{},
		//EXP: storage.Attribute{},
	}
}
func Marle() Character  {
	return Character{
		HP: storage.Attribute{0x253, true},
		//HPMax: storage.Attribute{},
		//MP: storage.Attribute{},
		//MPMax: storage.Attribute{},
		//ATK: storage.Attribute{},
		//DEF: storage.Attribute{},
		//PWR: storage.Attribute{},
		//SPD: storage.Attribute{},
		//HIT: storage.Attribute{},
		//EV: storage.Attribute{},
		//MAG: storage.Attribute{},
		//STAM: storage.Attribute{},
		//MDEF: storage.Attribute{},
		//EXP: storage.Attribute{},
	}
}
func Luca() Character {
	return Character{
		HP: storage.Attribute{0x2A3, true},
		//HPMax: storage.Attribute{},
		//MP: storage.Attribute{},
		//MPMax: storage.Attribute{},
		//ATK: storage.Attribute{},
		//DEF: storage.Attribute{},
		//PWR: storage.Attribute{},
		//SPD: storage.Attribute{},
		//HIT: storage.Attribute{},
		//EV: storage.Attribute{},
		//MAG: storage.Attribute{},
		//STAM: storage.Attribute{},
		//MDEF: storage.Attribute{},
		//EXP: storage.Attribute{},
	}
}
func Frog() Character {
	return Character{
		HP: storage.Attribute{0x2F3, true},
		//HPMax: storage.Attribute{},
		//MP: storage.Attribute{},
		//MPMax: storage.Attribute{},
		//ATK: storage.Attribute{},
		//DEF: storage.Attribute{},
		//PWR: storage.Attribute{},
		//SPD: storage.Attribute{},
		//HIT: storage.Attribute{},
		//EV: storage.Attribute{},
		//MAG: storage.Attribute{},
		//STAM: storage.Attribute{},
		//MDEF: storage.Attribute{},
		//EXP: storage.Attribute{},
	}
}
func Robo() Character {
	return Character{
		HP: storage.Attribute{0x343, true},
		//HPMax: storage.Attribute{},
		//MP: storage.Attribute{},
		//MPMax: storage.Attribute{},
		//ATK: storage.Attribute{},
		//DEF: storage.Attribute{},
		//PWR: storage.Attribute{},
		//SPD: storage.Attribute{},
		//HIT: storage.Attribute{},
		//EV: storage.Attribute{},
		//MAG: storage.Attribute{},
		//STAM: storage.Attribute{},
		//MDEF: storage.Attribute{},
		//EXP: storage.Attribute{},
	}
}