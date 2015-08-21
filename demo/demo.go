package main

import (
	"time"
	"image/color"
	"math/rand"
	"image"
	"github.com/andyleap/tinyfb"
)

func main() {
	t := tinyfb.New("test", 640, 480)
	go t.Run()
	
	i := image.NewRGBA(image.Rect(0, 0, 640, 480))
	
	
	for {
		i.SetRGBA(rand.Intn(640), rand.Intn(480), color.RGBA{
			R: uint8(rand.Intn(255)),
			G: uint8(rand.Intn(255)),
			B: uint8(rand.Intn(255)),
			A: 0,
		})
		
		t.Update(i)
		time.Sleep(time.Second / 20)
	}
	
}