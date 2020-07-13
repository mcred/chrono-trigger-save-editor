package app

import (
	"ChronoTrigger/internal/app/characters"
	"ChronoTrigger/internal/app/inventory"
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

func Run() {
	path := "./tests/files/save1.sav"
	GameData := storage.Open(path)

	fmt.Println("Crono")
	fmt.Println(GameData.GetValue(characters.Crono().Helmet))
	fmt.Println(GameData.GetValue(characters.Crono().Armor))
	fmt.Println(GameData.GetValue(characters.Crono().Weapon))
	fmt.Println(GameData.GetValue(characters.Crono().Relic))

	fmt.Println(inventory.Helmets().GetValByID(GameData.GetValue(characters.Crono().Helmet)))
	fmt.Println(inventory.Armors().GetValByID(GameData.GetValue(characters.Crono().Armor)))
	fmt.Println(inventory.Swords().GetValByID(GameData.GetValue(characters.Crono().Weapon)))
	fmt.Println(inventory.Relics().GetValByID(GameData.GetValue(characters.Crono().Relic)))

	//showCharacterAttributes(characters.Crono(), GameData)

}
