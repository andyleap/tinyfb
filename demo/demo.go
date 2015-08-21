package main

import (
	"fmt"
	"time"
	"image/color"
	"math/rand"
	"image"
	"github.com/andyleap/tinyfb"
)

func main() {
	t := tinyfb.New("test", 640, 480)
	rand.Seed(time.Now().UnixNano())
	go t.Run()
	
	i := image.NewRGBA(image.Rect(0, 0, 640, 480))
	step := 0
	
	for {
		start := time.Now()
		for x := 0; x <= 640; x++ {
			for y := 0; y <= 480; y++ {
				i.SetRGBA(x, y, color.RGBA{
					R: uint8((x+step)/3),
					G: 0,
					B: uint8(y/2),
					A: 0,
				})
			}
		}
		end := time.Now()
		
		fmt.Println("Calc Time: ", end.Sub(start))
		
		t.Update(i)
		step += 1
		time.Sleep(time.Second / 60)
	}
	
}