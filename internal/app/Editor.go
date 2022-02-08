package app

import (
	"ChronoTrigger/internal/app/characters"
	"encoding/binary"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/mcred/savetools"
	"net/url"
)

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

func (a *editorApp) loadCard() {
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

func (a *editorApp) saveCard() {
	card := generateChecksum(a.card)
	err := card.Save()
	if err != nil {
		dialog.ShowError(err, a.mainWin)
		return
	}
}

func (a *editorApp) loadUI() {
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

func (a *editorApp) init() {
	a.mainWin.Resize(fyne.NewSize(800, 600))
	a.loadUI()
	if a.card.Path == "" {
		a.loadCard()
	}
}

type editorApp struct {
	app     fyne.App
	mainWin fyne.Window
	card    savetools.Card
	crono   characters.Character
	marle   characters.Character
	lucca   characters.Character
}

// Run main entry point for chrono trigger editor
func Run() {
	a := app.New()
	w := a.NewWindow("Chrono Trigger Save File Editor")
	ui := &editorApp{app: a, mainWin: w, card: savetools.Card{}, crono: characters.Crono(), marle: characters.Marle(), lucca: characters.Lucca()}
	ui.init()
	w.SetContent(
		container.NewVScroll(container.New(layout.NewGridLayout(3),
			ui.crono.GenerateForm(),
			ui.marle.GenerateForm(),
			ui.lucca.GenerateForm(),
		)),
	)
	w.ShowAndRun()
}
