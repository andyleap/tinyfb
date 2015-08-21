package main

import (
	"time"
	"image/color"
	"math/rand"
	"math"
	"image"
	"github.com/andyleap/tinyfb"
	
	"github.com/tbogdala/noisey"
)

func main() {
	t := tinyfb.New("test", 320, 240)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	quit := false
	go func() {
		t.Run()
		quit = true
	}()
	perlin := noisey.NewPerlinGenerator(r)
	
	i := image.NewRGBA(image.Rect(0, 0, 320, 240))
	step := 0
	frame := time.Now()
	for !quit {
		for x := 0; x <= 320; x++ {
			for y := 0; y <= 240; y++ {
				noise := perlin.Get3D(float64(x)*0.03, float64(y)*0.03, float64(step)*0.05)
				val := uint8(math.Floor((noise*0.4 + 0.5) * 250))
				i.SetRGBA(x, y, color.RGBA{
					R: val,
					G: val,
					B: val,
					A: 0,
				})
			}
		}
		
		t.Update(i)
		step += 1
		end := time.Now()
		delta := end.Sub(frame).Nanoseconds() - (time.Second/60).Nanoseconds()
		if delta < 0 {
			time.Sleep(time.Duration(-delta))
		}
		
		frame = time.Now()
	}
	
}