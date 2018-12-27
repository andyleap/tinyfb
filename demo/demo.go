package main

import (
	"image"
	"image/color"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/andyleap/tinyfb"

	"github.com/tbogdala/noisey"
)

func main() {
	t := tinyfb.New("test", 400, 400)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	quit := false
	go func() {
		time.Sleep(1 * time.Second)

		perlin := noisey.NewOpenSimplexGenerator(r)

		i := image.NewRGBA(image.Rect(0, 0, 400, 400))
		step := 0
		xstep := 0
		ystep := 0
		frame := time.Now()

		t.Char(func(char string, mods int) {
			switch char {
			case "Up":
				ystep -= 10
			case "Down":
				ystep += 10
			case "Left":
				xstep -= 10
			case "Right":
				xstep += 10
			default:
				log.Println(char, mods)

			}
		})

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
			t.Update(i)
			step += 1
			end := time.Now()
			delta := end.Sub(frame).Nanoseconds() - (time.Second / 60).Nanoseconds()
			if delta < 0 {
				time.Sleep(time.Duration(-delta))
			}
			frame = time.Now()
		}
	}()
	t.Run()
	quit = true
}
