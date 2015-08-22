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
	t := tinyfb.New("test", 400, 400)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	quit := false
	go func() {
		t.Run()
		quit = true
	}()
	perlin := noisey.NewOpenSimplexGenerator(r)
	
	i := image.NewRGBA(image.Rect(0, 0, 400, 400))
	step := 0
	xstep := 0
	ystep := 0
	frame := time.Now()
	for !quit {
		for x := 0; x <= 400; x++ {
			for y := 0; y <= 400; y++ {
				noise := perlin.Get3D(float64(x+xstep)*0.03, float64(y+ystep)*0.03, float64(step)*0.01)
				val := uint8(math.Floor((noise*0.4 + 0.5) * 250))
				i.SetRGBA(x, y, color.RGBA{
					R: val,
					G: val,
					B: val,
					A: 0,
				})
			}
		}
		if t.Keys[tinyfb.VK_RIGHT] {
			xstep+=10
		}
		if t.Keys[tinyfb.VK_LEFT] {
			xstep-=10
		}
		if t.Keys[tinyfb.VK_UP] {
			ystep-=10
		}
		if t.Keys[tinyfb.VK_DOWN] {
			ystep+=10
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