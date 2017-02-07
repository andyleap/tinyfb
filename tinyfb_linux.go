package tinyfb

import (
	"image"
	"image/draw"

	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/keybind"
	"github.com/BurntSushi/xgbutil/xevent"
	"github.com/BurntSushi/xgbutil/xgraphics"
	"github.com/BurntSushi/xgbutil/xwindow"
)

type tinyFB struct {
	x    *xgbutil.XUtil
	win  *xwindow.Window
	img  *xgraphics.Image
	char func(char string, mods int)
}

func New(title string, width, height int32) TinyFB {
	X, _ := xgbutil.NewConn()
	t := &tinyFB{
		x: X,
	}

	keybind.Initialize(X)

	win, _ := xwindow.Generate(X)
	t.win = win
	win.Create(X.RootWin(), 0, 0, int(width), int(height), xproto.CwBackingPixel, 0x606060ff)

	win.Listen(xproto.EventMaskKeyPress, xproto.EventMaskKeyRelease)

	win.WMGracefulClose(
		func(w *xwindow.Window) {
			xevent.Detach(w.X, w.Id)
			keybind.Detach(w.X, w.Id)
			w.Destroy()
			xevent.Quit(X)
		})

	win.Map()

	xevent.KeyPressFun(
		func(X *xgbutil.XUtil, e xevent.KeyPressEvent) {
			if t.char != nil {
				keyStr := keybind.LookupString(X, e.State, e.Detail)
				t.char(keyStr, int(e.State))
			}
		}).Connect(X, win.Id)

	img := xgraphics.New(X, image.Rect(0, 0, int(width), int(height)))
	t.img = img
	img.XSurfaceSet(win.Id)
	img.XDraw()
	img.XPaint(win.Id)

	return t
}

func (t *tinyFB) Run() {
	xevent.Main(t.x)
}

func (t *tinyFB) Update(buffer *image.RGBA) {
	draw.Draw(t.img, t.img.Rect, buffer, image.Point{0, 0}, draw.Src)
	t.img.XDraw()
	t.img.XPaint(t.win.Id)
}

func (t *tinyFB) Close() {
	t.win.Unmap()
	xevent.Quit(t.x)
}

func (t *tinyFB) Char(char func(char string, mods int)) {
	t.char = char
}
