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

func (g *Game) SetValue(a Attribute, v int) {
	if a.Is16Bit {
		b := make([]byte, 2)
		binary.LittleEndian.PutUint16(b, uint16(v))
		g.Data[a.Location] = b[0]
		g.Data[a.Location + 1] = b[1]
	} else {
		g.Data[a.Location] = byte(v)
	}
}


func Open(path string) Game {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Could not open %s", path)
		fmt.Println("Could not open #{path}")
		os.Exit(1)
	}
	return Game{Data:f}
}

func (g *Game) Save() {
	generateChecksum(g)
}

func generateChecksum(g *Game) {
	for slot := 1; slot <= 3; slot ++ {
		checksum := 0
		max := slot * 0xA00
		min := max - 0xA00
		for i := min; i < max; i++ {
			if checksum > 0xFFFF {
				checksum -= 0xFFFF
			}
			if i % 2 == 0 {
				checksum += g.GetValue(Attribute{i, true})
			}
		}
		l := 0x1FF0 + ((slot - 1) * 2)
		g.SetValue(Attribute{l, true}, checksum)
	}
}