package party

import "ChronoTrigger/internal/app/storage"

type Party struct{
	//000 | 256 | Item Type
	//100 | 256 | Item Count
	//200 | 560 | Characters (Crono, Marle, Lucca, Robo, Frog, Ayla, Magus, Epoch)
	Member1 storage.Attribute
	Member2 storage.Attribute
	Member3 storage.Attribute
	SaveCount storage.Attribute
	Gold storage.Attribute
	//5B0 | 48  | Names (8 slots, 6 byte C-style strings)
	MilliSeconds storage.Attribute
	Seconds storage.Attribute
	Minutes storage.Attribute
	Hours storage.Attribute
	World storage.Attribute
	PosX storage.Attribute
	PosY storage.Attribute
}

func GetParty() Party {
	return Party{
		Member1:      storage.Attribute{0x580, 1},
		Member2:      storage.Attribute{0x581, 1},
		Member3:      storage.Attribute{0x582, 1},
		SaveCount:    storage.Attribute{0x59C, 1},
		Gold:         storage.Attribute{0x5E0, 2}, //maybe 24 bit
		MilliSeconds: storage.Attribute{0x5E3, 1},
		Seconds:      storage.Attribute{0x5E4, 1},
		Minutes:      storage.Attribute{0x5E5, 1},
		Hours:        storage.Attribute{0x5E6, 1},
		World:        storage.Attribute{0x5F3, 2},
		PosX:         storage.Attribute{0x5F5, 1},
		PosY:         storage.Attribute{0x5F6, 1},
	}
}