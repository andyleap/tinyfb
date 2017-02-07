package tinyfb

import (
	"image"
	"image/draw"
	"reflect"
	"runtime"
	"sync"
	"syscall"
	"unsafe"

	"github.com/andyleap/tinyfb/win"
)

type tinyFB struct {
	title  string
	width  int32
	height int32

	wnd win.HWND

	window_hdc         win.HDC
	surface_width      int32
	surface_height     int32
	bitmap_header      *win.BITMAPINFO
	bitmap_header_data []byte

	buffer     *image.RGBA
	bufferlock sync.Mutex

	wc win.WNDCLASSEX

	char func(char string, mods int)
}

func (t *tinyFB) wndProc(hwnd win.HWND, msg uint32, wParam, lParam uintptr) (result uintptr) {
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
		if t.char != nil {
			char := ""
			switch int(wParam) {
			case win.VK_UP:
				char = "Up"
			case win.VK_DOWN:
				char = "Down"
			case win.VK_LEFT:
				char = "Left"
			case win.VK_RIGHT:
				char = "Right"
			case win.VK_INSERT:
				char = "Insert"
			case win.VK_DELETE:
				char = "Delete"
			}
			if char != "" {
				t.char(char, 0)
			}
		}
	case uint32(win.WM_CHAR):
		if t.char != nil {
			t.char(string(rune(int(wParam))), 0)
		}
	default:
		result = win.DefWindowProc(hwnd, msg, wParam, lParam)
	}
	return
}

func New(title string, width, height int32) TinyFB {
	t := &tinyFB{title: title, width: width, height: height}
	t.buffer = image.NewRGBA(image.Rect(0, 0, int(width), int(height)))
	return t
}

func (t *tinyFB) Run() {
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
	t.bitmap_header_data = make([]byte, uint32(unsafe.Sizeof(win.BITMAPINFOHEADER{})+4*3))

	t.bitmap_header = (*win.BITMAPINFO)(unsafe.Pointer(&t.bitmap_header_data[0]))
	t.bitmap_header.BmiHeader.BiSize = uint32(unsafe.Sizeof(t.bitmap_header.BmiHeader) + 4*3)
	t.bitmap_header.BmiHeader.BiSizeImage = uint32(len(t.buffer.Pix) * 4)
	t.bitmap_header.BmiHeader.BiPlanes = 1
	t.bitmap_header.BmiHeader.BiBitCount = 32
	t.bitmap_header.BmiHeader.BiCompression = win.BI_BITFIELDS
	t.bitmap_header.BmiHeader.BiWidth = int32(t.surface_width)
	t.bitmap_header.BmiHeader.BiHeight = -int32(t.surface_height)
	t.bitmap_header.BmiHeader.BiClrUsed = 0

	masks := []uint32{}
	masksHeader := (*reflect.SliceHeader)(unsafe.Pointer(&masks))
	masksHeader.Data = uintptr(unsafe.Pointer(&t.bitmap_header.BmiColors))
	masksHeader.Len = 3
	masksHeader.Cap = 3

	masks[0] = 0x000000FF
	masks[1] = 0x0000FF00
	masks[2] = 0x00F00000

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

func (t *tinyFB) Update(buffer *image.RGBA) {
	t.bufferlock.Lock()
	draw.Draw(t.buffer, t.buffer.Rect, buffer, image.Point{0, 0}, draw.Src)
	t.bufferlock.Unlock()
	win.InvalidateRect(t.wnd, nil, true)
	win.SendMessage(t.wnd, win.WM_PAINT, 0, 0)
}

func (t *tinyFB) Close() {
	win.DestroyWindow(t.wnd)
}

func (t *tinyFB) Char(char func(char string, mods int)) {
	t.char = char
}
