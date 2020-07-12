package app

import (
	"ChronoTrigger/internal/app/characters"
	"ChronoTrigger/internal/app/storage"
	"fmt"
)

func Run()  {
	path := "./tests/files/save1.sav"
	GameData := storage.Open(path)
	fmt.Println(GameData.GetValue(characters.Chrono().HP))
}