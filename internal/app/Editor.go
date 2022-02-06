package app

import (
	"ChronoTrigger/internal/app/characters"
	"encoding/binary"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/mcred/savetools"
	"net/url"
)

//func showCharacterAttributes(c characters.Character, g storage.Game) {
//	//fmt.Println(utils.DecodeName(c, g))
//	//fmt.Print("Name ID: ")
//	//fmt.Println(g.getValue(c.NameID))
//	//fmt.Print("Character ID: ")
//	//fmt.Println(g.getValue(c.CharID))
//	//fmt.Print("Current HP: ")
//	//fmt.Println(g.getValue(c.HP))
//	//fmt.Print("Max Hp: ")
//	//fmt.Println(g.getValue(c.MaxHP))
//	//fmt.Print("Current MP: ")
//	//fmt.Println(g.getValue(c.MP))
//	//fmt.Print("Base Max HP:")
//	//fmt.Println(g.getValue(c.MaxMP))
//	//fmt.Print("Base Power: ")
//	//fmt.Println(g.getValue(c.BasePower))
//	//fmt.Print("Base Stamina: ")
//	//fmt.Println(g.getValue(c.BaseStamina))
//	//fmt.Print("Base Speed: ")
//	//fmt.Println(g.getValue(c.BaseSpeed))
//	//fmt.Print("Base Magic: ")
//	//fmt.Println(g.getValue(c.BaseMagic))
//	//fmt.Print("Base Hit: ")
//	//fmt.Println(g.getValue(c.BaseHit))
//	//fmt.Print("Base Evade: ")
//	//fmt.Println(g.getValue(c.BaseEvade))
//	//fmt.Print("Base Magic Defence: ")
//	//fmt.Println(g.getValue(c.BaseMagicDefense))
//	//fmt.Print("Level: ")
//	//fmt.Println(g.getValue(c.Level))
//	//fmt.Print("Experience: ")
//	//fmt.Println(g.getValue(c.Experience))
//	//fmt.Print("Helmet: ")
//	//fmt.Println(g.getValue(c.Helmet))
//	//fmt.Print("Armors: ")
//	//fmt.Println(g.getValue(c.Armor))
//	//fmt.Print("Weapon: ")
//	//fmt.Println(g.getValue(c.Weapon))
//	//fmt.Print("Relic: ")
//	//fmt.Println(g.getValue(c.Relic))
//	//fmt.Print("Exp To Level: ")
//	//fmt.Println(g.getValue(c.ExpToLevel))
//	//fmt.Print("Current Power: ")
//	//fmt.Println(g.getValue(c.CurrentPower))
//	//fmt.Print("Current Stamina: ")
//	//fmt.Println(g.getValue(c.CurrentStamina))
//	//fmt.Print("Current Speed: ")
//	//fmt.Println(g.getValue(c.CurrentSpeed))
//	//fmt.Print("Current Magic: ")
//	//fmt.Println(g.getValue(c.CurrentMagic))
//	//fmt.Print("Current Hit: ")
//	//fmt.Println(g.getValue(c.CurrentHit))
//	//fmt.Print("Current Evade: ")
//	//fmt.Println(g.getValue(c.CurrentEvade))
//	//fmt.Print("Current Magic Defence: ")
//	//fmt.Println(g.getValue(c.CurrentMagicDefense))
//	//fmt.Print("Current Attack: ")
//	//fmt.Println(g.getValue(c.CurrentAttack))
//	//fmt.Print("Current Defense: ")
//	//fmt.Println(g.getValue(c.CurrentDefense))
//	//fmt.Print("Current Max HP: ")
//	//fmt.Println(g.getValue(c.CurrentMaxHP))
//}

//func showPartyAttributes(p party.Party, g storage.Game) {
//	//fmt.Print("Member1: ")
//	//fmt.Println(g.getValue(p.Member1))
//	//fmt.Print("Member2: ")
//	//fmt.Println(g.getValue(p.Member2))
//	//fmt.Print("Member3: ")
//	//fmt.Println(g.getValue(p.Member3))
//	//fmt.Print("SaveCount: ")
//	//fmt.Println(g.getValue(p.SaveCount))
//	//fmt.Print("Gold: ")
//	//fmt.Println(g.getValue(p.Gold))
//	//fmt.Print("MilliSeconds: ")
//	//fmt.Println(g.getValue(p.MilliSeconds))
//	//fmt.Print("Seconds: ")
//	//fmt.Println(g.getValue(p.Seconds))
//	//fmt.Print("Minutes: ")
//	//fmt.Println(g.getValue(p.Minutes))
//	//fmt.Print("Hours: ")
//	//fmt.Println(g.getValue(p.Hours))
//	//fmt.Print("World: ")
//	//fmt.Println(g.getValue(p.World))
//	//fmt.Print("PosX: ")
//	//fmt.Println(g.getValue(p.PosX))
//	//fmt.Print("PosY: ")
//	//fmt.Println(g.getValue(p.PosY))
//	//
//	//p.LoadInventory(g)
//	//fmt.Println("Inventory: ")
//	//for _, item := range p.Inventory {
//	//	fmt.Print(item.Name)
//	//	fmt.Print(": ")
//	//	fmt.Println(item.Count)
//	//}
//}

func (a *App) loadCard() {
	d := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
		if err != nil {
			dialog.ShowError(err, a.mainWin)
			return
		}
		if err == nil && reader == nil {
			return
		}

		a.card, err = savetools.Load(reader.URI().Path(), 3, 0xA00)

		a.crono.Init(&a.card)
		a.marle.Init(&a.card)
		a.lucca.Init(&a.card)

		if err != nil {
			dialog.ShowError(err, a.mainWin)
			return
		}
		defer reader.Close()
	}, a.mainWin)
	d.Resize(fyne.NewSize(700, 500))
	d.SetFilter(storage.NewExtensionFileFilter([]string{".sav"}))
	d.Show()
}

/*
Thanks to https://github.com/mikearnos/snessum for help with this
*/
func generateChecksum(c savetools.Card) savetools.Card {
	for slot := 1; slot <= 3; slot++ {
		var checksum uint = 0
		max := (slot * 0xA00) - 2
		min := (slot - 1) * 0xA00
		for i := max; i >= min; i -= 2 {
			if checksum > 0xFFFF {
				checksum -= 0xFFFF
			}
			checksum += uint(c.GetValue(savetools.Attribute{i, 16, binary.LittleEndian}))
		}
		checksum &= 0xFFFF
		l := 0x1FF0 + ((slot - 1) * 2)
		c.SetValue(savetools.Attribute{l, 16, binary.LittleEndian}, int(checksum))
	}
	return c
}

func (a *App) saveCard() {
	card := generateChecksum(a.card)
	err := card.Save()
	if err != nil {
		dialog.ShowError(err, a.mainWin)
		return
	}
}

func (a *App) loadUI() {
	link, _ := url.Parse("https://github.com/mcred/chrono-trigger-save-editor")
	mainMenu := fyne.NewMainMenu(
		fyne.NewMenu("File",
			fyne.NewMenuItem("Open", a.loadCard),
			fyne.NewMenuItem("Save", a.saveCard),
		),
		fyne.NewMenu("Help",
			fyne.NewMenuItem("About", func() {
				dialog.ShowCustom("About", "Ok", container.NewVBox(
					widget.NewLabel("Chrono Trigger Save File Editor."),
					widget.NewHyperlink("Help and more information on Github", link),
					widget.NewLabel("License: MIT"),
				), a.mainWin)
			}),
		),
	)
	a.mainWin.SetMainMenu(mainMenu)
}

func (a *App) init() {
	a.mainWin.Resize(fyne.NewSize(800, 600))
	a.loadUI()
	if a.card.Path == "" {
		a.loadCard()
	}
}

type App struct {
	app     fyne.App
	mainWin fyne.Window
	card    savetools.Card
	crono   characters.Character
	marle   characters.Character
	lucca   characters.Character
}

func Run() {
	a := app.New()
	w := a.NewWindow("Chrono Trigger Save File Editor")
	ui := &App{app: a, mainWin: w, card: savetools.Card{}, crono: characters.Crono(), marle: characters.Marle(), lucca: characters.Lucca()}
	ui.init()
	w.SetContent(container.NewVBox(
		ui.crono.GenerateForm(&ui.card),
		ui.marle.GenerateForm(&ui.card),
		ui.lucca.GenerateForm(&ui.card),
	))
	w.ShowAndRun()
}
