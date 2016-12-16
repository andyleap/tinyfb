package tinyfb

import (
	"image"
	"image/draw"
	"runtime"
	"sync"
	"syscall"
	"unsafe"

	"github.com/andyleap/tinyfb/win"
)

type TinyFB struct {
	title  string
	width  int32
	height int32

	wnd win.HWND

	window_hdc     win.HDC
	surface_width  int32
	surface_height int32
	bitmap_header  *win.BITMAPINFO

	buffer     *image.RGBA
	bufferlock sync.Mutex

	wc win.WNDCLASSEX

	Keys map[int]bool

	KeyDown   func(int)
	KeyUp     func(int)
	Char      func(rune)
	MouseDown func(x, y int, button int)
	MouseUp   func(x, y int, button int)
	MouseMove func(x, y int)
}

func (t *TinyFB) wndProc(hwnd win.HWND, msg uint32, wParam, lParam uintptr) (result uintptr) {
	result = 0
	switch msg {
	case uint32(win.WM_PAINT):
		var size win.RECT
		win.GetClientRect(t.wnd, &size)
		t.bufferlock.Lock()
		win.StretchDIBits(t.window_hdc, 0, 0, size.Right, size.Bottom, 0, 0, t.surface_width, t.surface_height, uintptr(unsafe.Pointer(&t.buffer.Pix[0])), t.bitmap_header, 0, win.SRCCOPY)
		t.bufferlock.Unlock()
		win.ValidateRect(t.wnd, nil)
	case uint32(win.WM_KEYDOWN):
		t.Keys[int(wParam)] = true
		if t.KeyDown != nil {
			t.KeyDown(int(wParam))
		}
	case uint32(win.WM_KEYUP):
		t.Keys[int(wParam)] = false
		if t.KeyUp != nil {
			t.KeyUp(int(wParam))
		}
	case uint32(win.WM_CHAR):
		if t.Char != nil {
			t.Char(rune(int(wParam)))
		}
	case uint32(win.WM_LBUTTONDOWN):
		if t.MouseDown != nil {
			point := uint64(lParam)
			x, y := int(point&0xFFFF), int(point>>16)
			t.MouseDown(x, y, 1)
		}
	case uint32(win.WM_LBUTTONUP):
		if t.MouseUp != nil {
			point := uint64(lParam)
			x, y := int(point&0xFFFF), int(point>>16)
			t.MouseUp(x, y, 1)
		}
	case uint32(win.WM_MOUSEMOVE):
		if t.MouseMove != nil {
			point := uint64(lParam)
			x, y := int(point&0xFFFF), int(point>>16)
			t.MouseMove(x, y)
		}
	default:
		result = win.DefWindowProc(hwnd, msg, wParam, lParam)
	}
	return
}

func New(title string, width, height int32) *TinyFB {
	t := &TinyFB{title: title, width: width, height: height, Keys: make(map[int]bool)}
	t.buffer = image.NewRGBA(image.Rect(0, 0, int(width), int(height)))
	return t
}

func (t *TinyFB) Run() {
	runtime.LockOSThread()
	var rect win.RECT

	t.wc.CbSize = uint32(unsafe.Sizeof(t.wc))
	t.wc.Style = win.CS_OWNDC | win.CS_VREDRAW | win.CS_HREDRAW
	t.wc.LpfnWndProc = syscall.NewCallback(t.wndProc)
	t.wc.HCursor = win.LoadCursor(0, (*uint16)(unsafe.Pointer(uintptr(win.IDC_ARROW))))
	t.wc.LpszClassName = win.StringToBSTR(t.title)
	win.RegisterClassEx(&t.wc)

	rect.Right = t.width
	rect.Bottom = t.height
	win.AdjustWindowRect(&rect, win.WS_POPUP|win.WS_SYSMENU|win.WS_CAPTION, false)
	rect.Right -= rect.Left
	rect.Bottom -= rect.Top

	t.surface_height = t.height
	t.surface_width = t.width

	t.wnd = win.CreateWindowEx(0, t.wc.LpszClassName, t.wc.LpszClassName, win.WS_OVERLAPPEDWINDOW & ^win.WS_MAXIMIZEBOX & ^win.WS_THICKFRAME, win.CW_USEDEFAULT, win.CW_USEDEFAULT, rect.Right, rect.Bottom, 0, 0, 0, nil)

	win.ShowWindow(t.wnd, win.SW_NORMAL)
	t.bitmap_header = &win.BITMAPINFO{}
	t.bitmap_header.BmiHeader.BiSize = uint32(unsafe.Sizeof(t.bitmap_header.BmiHeader))
	t.bitmap_header.BmiHeader.BiPlanes = 1
	t.bitmap_header.BmiHeader.BiBitCount = 32
	t.bitmap_header.BmiHeader.BiCompression = win.BI_RGB
	t.bitmap_header.BmiHeader.BiWidth = int32(t.surface_width)
	t.bitmap_header.BmiHeader.BiHeight = -int32(t.surface_height)

	t.window_hdc = win.GetDC(t.wnd)
	var msg win.MSG

	for {
		switch win.GetMessage(&msg, t.wnd, 0, 0) {
		case 0:
			return
		case -1:
			return
		}
		win.TranslateMessage(&msg)
		win.DispatchMessage(&msg)
	}
}

func (t *TinyFB) Update(buffer *image.RGBA) {
	t.bufferlock.Lock()
	draw.Draw(t.buffer, t.buffer.Rect, buffer, image.Point{0, 0}, draw.Src)
	t.bufferlock.Unlock()
	win.InvalidateRect(t.wnd, nil, true)
	win.SendMessage(t.wnd, win.WM_PAINT, 0, 0)
}

func (t *TinyFB) Close() {
	win.DestroyWindow(t.wnd)
}
