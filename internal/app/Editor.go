package app

import (
	"ChronoTrigger/internal/app/characters"
	"ChronoTrigger/internal/app/inventory"
	"ChronoTrigger/internal/app/party"
	"ChronoTrigger/internal/app/storage"
	"ChronoTrigger/internal/app/utils"
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
	fmt.Println(utils.DecodeName(characters.Crono(), GameData))
	fmt.Println(utils.DecodeName(characters.Marle(), GameData))
	fmt.Println(utils.DecodeName(characters.Lucca(), GameData))
	fmt.Println(utils.DecodeName(characters.Frog(), GameData))
	fmt.Println(utils.DecodeName(characters.Robo(), GameData))
	fmt.Println(utils.DecodeName(characters.Ayla(), GameData))
	fmt.Println(utils.DecodeName(characters.Magus(), GameData))
}
