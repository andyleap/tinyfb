package tinyfb

import (
	"image"
	"image/draw"

	//	"golang.org/x/mobile/event/key"
	//	"golang.org/x/mobile/event/mouse"

	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
)

type tinyFB struct {
	title         string
	width, height int

	w screen.Window

	char func(char string, mods int)
	key  func(key string, mods int, press bool)
}

func New(title string, width, height int32) TinyFB {
	return &tinyFB{
		title:  title,
		width:  int(width),
		height: int(height),
	}
}

func (t *tinyFB) Run() {
	driver.Main(func(s screen.Screen) {
		w, err := s.NewWindow(&screen.NewWindowOptions{Title: t.title, Width: t.width, Height: t.height})
		t.w = w
		if err != nil {
			panic(err)
			return
		}
		defer w.Release()

		buf, _ := s.NewBuffer(image.Point{t.width, t.height})
		tex, _ := s.NewTexture(image.Point{t.width, t.height})

		for {
			e := w.NextEvent()

			switch et := e.(type) {
			case *image.RGBA:
				draw.Draw(buf.RGBA(), buf.Bounds(), et, image.Point{}, draw.Src)

				tex.Upload(image.Point{}, buf, buf.Bounds())
				w.Copy(image.Point{}, tex, tex.Bounds(), screen.Src, nil)
				w.Publish()
			case string:
				if et == "quit" {
					return
				}
			}
		}
	})
}

func (t *tinyFB) Update(buffer *image.RGBA) {
	if t.w != nil {
		t.w.Send(buffer)
	}
}

func (t *tinyFB) Close() {
	t.w.Send("quit")
}

func (t *tinyFB) Char(char func(char string, mods int)) {
	t.char = char
}

func (t *tinyFB) Key(key func(key string, mods int, press bool)) {
	t.key = key
}
