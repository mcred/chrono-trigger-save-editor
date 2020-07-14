package app

import (
	"ChronoTrigger/internal/app/characters"
	"ChronoTrigger/internal/app/inventory"
	"ChronoTrigger/internal/app/party"
	"ChronoTrigger/internal/app/storage"
	"fmt"
)

func showCharacterAttributes(character characters.Character, g storage.Game) {
	fmt.Print("Name ID: ")
	fmt.Println(g.GetValue(character.NameID))
	fmt.Print("Character ID: ")
	fmt.Println(g.GetValue(character.CharID))
	fmt.Print("Current HP: ")
	fmt.Println(g.GetValue(character.HP))
	fmt.Print("Max Hp: ")
	fmt.Println(g.GetValue(character.MaxHP))
	fmt.Print("Current MP: ")
	fmt.Println(g.GetValue(character.MP))
	fmt.Print("Base Max HP:")
	fmt.Println(g.GetValue(character.MaxMP))
	fmt.Print("Base Power: ")
	fmt.Println(g.GetValue(character.BasePower))
	fmt.Print("Base Stamina: ")
	fmt.Println(g.GetValue(character.BaseStamina))
	fmt.Print("Base Speed: ")
	fmt.Println(g.GetValue(character.BaseSpeed))
	fmt.Print("Base Magic: ")
	fmt.Println(g.GetValue(character.BaseMagic))
	fmt.Print("Base Hit: ")
	fmt.Println(g.GetValue(character.BaseHit))
	fmt.Print("Base Evade: ")
	fmt.Println(g.GetValue(character.BaseEvade))
	fmt.Print("Base Magic Defence: ")
	fmt.Println(g.GetValue(character.BaseMagicDefense))
	fmt.Print("Level: ")
	fmt.Println(g.GetValue(character.Level))
	fmt.Print("Experience: ")
	fmt.Println(g.GetValue(character.Experience))
	fmt.Print("Helmet: ")
	fmt.Println(g.GetValue(character.Helmet))
	fmt.Print("Armors: ")
	fmt.Println(g.GetValue(character.Armor))
	fmt.Print("Weapon: ")
	fmt.Println(g.GetValue(character.Weapon))
	fmt.Print("Relic: ")
	fmt.Println(g.GetValue(character.Relic))
	fmt.Print("Exp To Level: ")
	fmt.Println(g.GetValue(character.ExpToLevel))
	fmt.Print("Current Power: ")
	fmt.Println(g.GetValue(character.CurrentPower))
	fmt.Print("Current Stamina: ")
	fmt.Println(g.GetValue(character.CurrentStamina))
	fmt.Print("Current Speed: ")
	fmt.Println(g.GetValue(character.CurrentSpeed))
	fmt.Print("Current Magic: ")
	fmt.Println(g.GetValue(character.CurrentMagic))
	fmt.Print("Current Hit: ")
	fmt.Println(g.GetValue(character.CurrentHit))
	fmt.Print("Current Evade: ")
	fmt.Println(g.GetValue(character.CurrentEvade))
	fmt.Print("Current Magic Defence: ")
	fmt.Println(g.GetValue(character.CurrentMagicDefense))
	fmt.Print("Current Attack: ")
	fmt.Println(g.GetValue(character.CurrentAttack))
	fmt.Print("Current Defense: ")
	fmt.Println(g.GetValue(character.CurrentDefense))
	fmt.Print("Current Max HP: ")
	fmt.Println(g.GetValue(character.CurrentMaxHP))
}

func showPartyAttributes(p party.Party, g storage.Game){
	fmt.Print("Member1: ")
	fmt.Println(g.GetValue(p.Member1))
	fmt.Print("Member2: ")
	fmt.Println(g.GetValue(p.Member2))
	fmt.Print("Member3: ")
	fmt.Println(g.GetValue(p.Member3))
	fmt.Print("SaveCount: ")
	fmt.Println(g.GetValue(p.SaveCount))
	fmt.Print("Gold: ")
	fmt.Println(g.GetValue(p.Gold))
	fmt.Print("MilliSeconds: ")
	fmt.Println(g.GetValue(p.MilliSeconds))
	fmt.Print("Seconds: ")
	fmt.Println(g.GetValue(p.Seconds))
	fmt.Print("Minutes: ")
	fmt.Println(g.GetValue(p.Minutes))
	fmt.Print("Hours: ")
	fmt.Println(g.GetValue(p.Hours))
	fmt.Print("World: ")
	fmt.Println(g.GetValue(p.World))
	fmt.Print("PosX: ")
	fmt.Println(g.GetValue(p.PosX))
	fmt.Print("PosY: ")
	fmt.Println(g.GetValue(p.PosY))
}

func Run() {
	path := "./tests/files/save1.sav"
	GameData := storage.Open(path)

	for i := 0x0; i <= 0xFF; i++ {
		v := GameData.GetValue(storage.Attribute{i, false})
		if v != 0{
			fmt.Print(inventory.AllItems().GetValByID(v))
			fmt.Print(" ")
			fmt.Println(GameData.GetValue(storage.Attribute{i + 0x100, false}))
		}
	}

	//showPartyAttributes(party.GetParty(), GameData)
	fmt.Println(decodeName(characters.Crono(), GameData))
	fmt.Println(decodeName(characters.Marle(), GameData))
	fmt.Println(decodeName(characters.Lucca(), GameData))
	fmt.Println(decodeName(characters.Frog(), GameData))
	fmt.Println(decodeName(characters.Robo(), GameData))
	fmt.Println(decodeName(characters.Ayla(), GameData))
	fmt.Println(decodeName(characters.Magus(), GameData))
}

func decodeName(c characters.Character, g storage.Game) string {
	id := g.GetValue(c.NameID)
	ints := []uint{
		g.GetValue(storage.Attribute{int(0x5B0 + (id * 6)), false}),
		g.GetValue(storage.Attribute{int(0x5B1 + (id * 6)), false}),
		g.GetValue(storage.Attribute{int(0x5B2 + (id * 6)), false}),
		g.GetValue(storage.Attribute{int(0x5B3 + (id * 6)), false}),
		g.GetValue(storage.Attribute{int(0x5B4 + (id * 6)), false}),
		g.GetValue(storage.Attribute{int(0x5B5 + (id * 6)), false}),
	}
	var s string
	for _, i := range ints {
		s += getCharFromInt(i)
	}
	return s
}

func getCharFromInt(i uint) string {
	m := map[uint]string{
		0xA0: "A",
		0xA1: "B",
		0xA2: "C",
		0xA3: "D",
		0xA4: "E",
		0xA5: "F",
		0xA6: "G",
		0xA7: "H",
		0xA8: "I",
		0xA9: "J",
		0xAA: "K",
		0xAB: "L",
		0xAC: "M",
		0xAD: "N",
		0xAE: "O",
		0xAF: "P",
		0xB0: "Q",
		0xB1: "R",
		0xB2: "S",
		0xB3: "T",
		0xB4: "U",
		0xB5: "V",
		0xB6: "W",
		0xB7: "X",
		0xB8: "Y",
		0xB9: "Z",
		0xBA: "a",
		0xBB: "b",
		0xBC: "c",
		0xBD: "d",
		0xBE: "e",
		0xBF: "f",
		0xC0: "g",
		0xC1: "h",
		0xC2: "i",
		0xC3: "j",
		0xC4: "k",
		0xC5: "l",
		0xC6: "m",
		0xC7: "n",
		0xC8: "o",
		0xC9: "p",
		0xCA: "q",
		0xCB: "r",
		0xCC: "s",
		0xCD: "t",
		0xCE: "u",
		0xCF: "v",
		0xD0: "w",
		0xD1: "x",
		0xD2: "y",
		0xD3: "z",
		0xD4: "0",
		0xD5: "1",
		0xD6: "2",
		0xD7: "3",
		0xD8: "4",
		0xD9: "5",
		0xDA: "6",
		0xDB: "7",
		0xDC: "8",
		0xDD: "9",
	}
	return m[i]
}