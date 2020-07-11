package app

import (
	"ChronoTrigger/internal/app/characters"
	"ChronoTrigger/internal/app/storage"
	"fmt"
)

func Run()  {
	path := "/Users/mcred/Library/Application Support/OpenEmu/SNES9x/Battery Saves/Chrono Trigger (USA).sav"
	GameData := storage.Open(path)


	fmt.Println(GameData.GetValue(characters.Chrono().HP))
	fmt.Println(GameData.GetValue(characters.Marle().HP))
	fmt.Println(GameData.GetValue(characters.Luca().HP))
	fmt.Println(GameData.GetValue(characters.Frog().HP))
	fmt.Println(GameData.GetValue(characters.Robo().HP))
}