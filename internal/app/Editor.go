package app

import (
	"ChronoTrigger/internal/app/characters"
	"ChronoTrigger/internal/app/party"
	"ChronoTrigger/internal/app/storage"
	"ChronoTrigger/internal/app/utils"
	"fmt"
)

func showCharacterAttributes(c characters.Character, g storage.Game) {
	fmt.Println(utils.DecodeName(c, g))
	fmt.Print("Name ID: ")
	fmt.Println(g.GetValue(c.NameID))
	fmt.Print("Character ID: ")
	fmt.Println(g.GetValue(c.CharID))
	fmt.Print("Current HP: ")
	fmt.Println(g.GetValue(c.HP))
	fmt.Print("Max Hp: ")
	fmt.Println(g.GetValue(c.MaxHP))
	fmt.Print("Current MP: ")
	fmt.Println(g.GetValue(c.MP))
	fmt.Print("Base Max HP:")
	fmt.Println(g.GetValue(c.MaxMP))
	fmt.Print("Base Power: ")
	fmt.Println(g.GetValue(c.BasePower))
	fmt.Print("Base Stamina: ")
	fmt.Println(g.GetValue(c.BaseStamina))
	fmt.Print("Base Speed: ")
	fmt.Println(g.GetValue(c.BaseSpeed))
	fmt.Print("Base Magic: ")
	fmt.Println(g.GetValue(c.BaseMagic))
	fmt.Print("Base Hit: ")
	fmt.Println(g.GetValue(c.BaseHit))
	fmt.Print("Base Evade: ")
	fmt.Println(g.GetValue(c.BaseEvade))
	fmt.Print("Base Magic Defence: ")
	fmt.Println(g.GetValue(c.BaseMagicDefense))
	fmt.Print("Level: ")
	fmt.Println(g.GetValue(c.Level))
	fmt.Print("Experience: ")
	fmt.Println(g.GetValue(c.Experience))
	fmt.Print("Helmet: ")
	fmt.Println(g.GetValue(c.Helmet))
	fmt.Print("Armors: ")
	fmt.Println(g.GetValue(c.Armor))
	fmt.Print("Weapon: ")
	fmt.Println(g.GetValue(c.Weapon))
	fmt.Print("Relic: ")
	fmt.Println(g.GetValue(c.Relic))
	fmt.Print("Exp To Level: ")
	fmt.Println(g.GetValue(c.ExpToLevel))
	fmt.Print("Current Power: ")
	fmt.Println(g.GetValue(c.CurrentPower))
	fmt.Print("Current Stamina: ")
	fmt.Println(g.GetValue(c.CurrentStamina))
	fmt.Print("Current Speed: ")
	fmt.Println(g.GetValue(c.CurrentSpeed))
	fmt.Print("Current Magic: ")
	fmt.Println(g.GetValue(c.CurrentMagic))
	fmt.Print("Current Hit: ")
	fmt.Println(g.GetValue(c.CurrentHit))
	fmt.Print("Current Evade: ")
	fmt.Println(g.GetValue(c.CurrentEvade))
	fmt.Print("Current Magic Defence: ")
	fmt.Println(g.GetValue(c.CurrentMagicDefense))
	fmt.Print("Current Attack: ")
	fmt.Println(g.GetValue(c.CurrentAttack))
	fmt.Print("Current Defense: ")
	fmt.Println(g.GetValue(c.CurrentDefense))
	fmt.Print("Current Max HP: ")
	fmt.Println(g.GetValue(c.CurrentMaxHP))
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

	p.LoadInventory(g)
	fmt.Println("Inventory: ")
	for _, item := range p.Inventory {
		fmt.Print(item.Name)
		fmt.Print(": ")
		fmt.Println(item.Count)
	}
}

func Run() {
	path := "./tests/files/save2.sav"
	GameData := storage.Open(path)

	showCharacterAttributes(characters.Crono(), GameData)
	showPartyAttributes(party.GetParty(), GameData)
}
