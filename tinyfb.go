package tinyfb

import (
	"image"
)

type TinyFB interface {
	Run()
	Update(buffer *image.RGBA)
	Close()
	Char(func(char string, mods int))
	Key(func(key string, mods int, press bool))
}
