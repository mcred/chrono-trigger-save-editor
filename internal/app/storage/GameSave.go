package storage

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"os"
)

type Attribute struct {
	Location int
	Is16Bit bool
}

type Game struct {
	Data []byte
}

func (g *Game) GetValue(a Attribute) int {
	var r int
	if a.Is16Bit {
		r = int(binary.LittleEndian.Uint16(g.Data[a.Location:a.Location+2]))
	} else {
		r = int(g.Data[a.Location])
	}
	return r
}

func Open(path string) Game {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Could not open %s", path)
		os.Exit(1)
	}
	return Game{Data:f}
}

func Save() {
	genrateChecksum()
}

func genrateChecksum() {

}