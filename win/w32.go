package win

import (
	"syscall"
	"unsafe"
)

const CW_USEDEFAULT = ^0x7fffffff

// Bitmap compression constants
const (
	BI_RGB       = 0
	BI_RLE8      = 1
	BI_RLE4      = 2
	BI_BITFIELDS = 3
	BI_JPEG      = 4
	BI_PNG       = 5
)

const (
	DIB_RGB_COLORS = 0
	DIB_PAL_COLORS = 1
)

type (
	BOOL    int32
	HRESULT int32
)

type (
	ATOM      uint16
	HANDLE    uintptr
	HGLOBAL   HANDLE
	HINSTANCE HANDLE
	LCID      uint32
	LCTYPE    uint32
)

type (
	COLORREF     uint32
	HBITMAP      HGDIOBJ
	HBRUSH       HGDIOBJ
	HDC          HANDLE
	HFONT        HGDIOBJ
	HGDIOBJ      HANDLE
	HENHMETAFILE HANDLE
	HPALETTE     HGDIOBJ
	HPEN         HGDIOBJ
	HREGION      HGDIOBJ
)

type (
	HACCEL    HANDLE
	HCURSOR   HANDLE
	HDWP      HANDLE
	HICON     HANDLE
	HMENU     HANDLE
	HMONITOR  HANDLE
	HRAWINPUT HANDLE
	HWND      HANDLE
)

type WNDCLASSEX struct {
	CbSize        uint32
	Style         uint32
	LpfnWndProc   uintptr
	CbClsExtra    int32
	CbWndExtra    int32
	HInstance     HINSTANCE
	HIcon         HICON
	HCursor       HCURSOR
	HbrBackground HBRUSH
	LpszMenuName  *uint16
	LpszClassName *uint16
	HIconSm       HICON
}

// Window message constants
const (
	WM_APP                    = 32768
	WM_ACTIVATE               = 6
	WM_ACTIVATEAPP            = 28
	WM_AFXFIRST               = 864
	WM_AFXLAST                = 895
	WM_ASKCBFORMATNAME        = 780
	WM_CANCELJOURNAL          = 75
	WM_CANCELMODE             = 31
	WM_CAPTURECHANGED         = 533
	WM_CHANGECBCHAIN          = 781
	WM_CHAR                   = 258
	WM_CHARTOITEM             = 47
	WM_CHILDACTIVATE          = 34
	WM_CLEAR                  = 771
	WM_CLOSE                  = 16
	WM_COMMAND                = 273
	WM_COMMNOTIFY             = 68 /* OBSOLETE */
	WM_COMPACTING             = 65
	WM_COMPAREITEM            = 57
	WM_CONTEXTMENU            = 123
	WM_COPY                   = 769
	WM_COPYDATA               = 74
	WM_CREATE                 = 1
	WM_CTLCOLORBTN            = 309
	WM_CTLCOLORDLG            = 310
	WM_CTLCOLOREDIT           = 307
	WM_CTLCOLORLISTBOX        = 308
	WM_CTLCOLORMSGBOX         = 306
	WM_CTLCOLORSCROLLBAR      = 311
	WM_CTLCOLORSTATIC         = 312
	WM_CUT                    = 768
	WM_DEADCHAR               = 259
	WM_DELETEITEM             = 45
	WM_DESTROY                = 2
	WM_DESTROYCLIPBOARD       = 775
	WM_DEVICECHANGE           = 537
	WM_DEVMODECHANGE          = 27
	WM_DISPLAYCHANGE          = 126
	WM_DRAWCLIPBOARD          = 776
	WM_DRAWITEM               = 43
	WM_DROPFILES              = 563
	WM_ENABLE                 = 10
	WM_ENDSESSION             = 22
	WM_ENTERIDLE              = 289
	WM_ENTERMENULOOP          = 529
	WM_ENTERSIZEMOVE          = 561
	WM_ERASEBKGND             = 20
	WM_EXITMENULOOP           = 530
	WM_EXITSIZEMOVE           = 562
	WM_FONTCHANGE             = 29
	WM_GETDLGCODE             = 135
	WM_GETFONT                = 49
	WM_GETHOTKEY              = 51
	WM_GETICON                = 127
	WM_GETMINMAXINFO          = 36
	WM_GETTEXT                = 13
	WM_GETTEXTLENGTH          = 14
	WM_HANDHELDFIRST          = 856
	WM_HANDHELDLAST           = 863
	WM_HELP                   = 83
	WM_HOTKEY                 = 786
	WM_HSCROLL                = 276
	WM_HSCROLLCLIPBOARD       = 782
	WM_ICONERASEBKGND         = 39
	WM_INITDIALOG             = 272
	WM_INITMENU               = 278
	WM_INITMENUPOPUP          = 279
	WM_INPUT                  = 0X00FF
	WM_INPUTLANGCHANGE        = 81
	WM_INPUTLANGCHANGEREQUEST = 80
	WM_KEYDOWN                = 256
	WM_KEYUP                  = 257
	WM_KILLFOCUS              = 8
	WM_MDIACTIVATE            = 546
	WM_MDICASCADE             = 551
	WM_MDICREATE              = 544
	WM_MDIDESTROY             = 545
	WM_MDIGETACTIVE           = 553
	WM_MDIICONARRANGE         = 552
	WM_MDIMAXIMIZE            = 549
	WM_MDINEXT                = 548
	WM_MDIREFRESHMENU         = 564
	WM_MDIRESTORE             = 547
	WM_MDISETMENU             = 560
	WM_MDITILE                = 550
	WM_MEASUREITEM            = 44
	WM_GETOBJECT              = 0X003D
	WM_CHANGEUISTATE          = 0X0127
	WM_UPDATEUISTATE          = 0X0128
	WM_QUERYUISTATE           = 0X0129
	WM_UNINITMENUPOPUP        = 0X0125
	WM_MENURBUTTONUP          = 290
	WM_MENUCOMMAND            = 0X0126
	WM_MENUGETOBJECT          = 0X0124
	WM_MENUDRAG               = 0X0123
	WM_APPCOMMAND             = 0X0319
	WM_MENUCHAR               = 288
	WM_MENUSELECT             = 287
	WM_MOVE                   = 3
	WM_MOVING                 = 534
	WM_NCACTIVATE             = 134
	WM_NCCALCSIZE             = 131
	WM_NCCREATE               = 129
	WM_NCDESTROY              = 130
	WM_NCHITTEST              = 132
	WM_NCLBUTTONDBLCLK        = 163
	WM_NCLBUTTONDOWN          = 161
	WM_NCLBUTTONUP            = 162
	WM_NCMBUTTONDBLCLK        = 169
	WM_NCMBUTTONDOWN          = 167
	WM_NCMBUTTONUP            = 168
	WM_NCXBUTTONDOWN          = 171
	WM_NCXBUTTONUP            = 172
	WM_NCXBUTTONDBLCLK        = 173
	WM_NCMOUSEHOVER           = 0X02A0
	WM_NCMOUSELEAVE           = 0X02A2
	WM_NCMOUSEMOVE            = 160
	WM_NCPAINT                = 133
	WM_NCRBUTTONDBLCLK        = 166
	WM_NCRBUTTONDOWN          = 164
	WM_NCRBUTTONUP            = 165
	WM_NEXTDLGCTL             = 40
	WM_NEXTMENU               = 531
	WM_NOTIFY                 = 78
	WM_NOTIFYFORMAT           = 85
	WM_NULL                   = 0
	WM_PAINT                  = 15
	WM_PAINTCLIPBOARD         = 777
	WM_PAINTICON              = 38
	WM_PALETTECHANGED         = 785
	WM_PALETTEISCHANGING      = 784
	WM_PARENTNOTIFY           = 528
	WM_PASTE                  = 770
	WM_PENWINFIRST            = 896
	WM_PENWINLAST             = 911
	WM_POWER                  = 72
	WM_POWERBROADCAST         = 536
	WM_PRINT                  = 791
	WM_PRINTCLIENT            = 792
	WM_QUERYDRAGICON          = 55
	WM_QUERYENDSESSION        = 17
	WM_QUERYNEWPALETTE        = 783
	WM_QUERYOPEN              = 19
	WM_QUEUESYNC              = 35
	WM_QUIT                   = 18
	WM_RENDERALLFORMATS       = 774
	WM_RENDERFORMAT           = 773
	WM_SETCURSOR              = 32
	WM_SETFOCUS               = 7
	WM_SETFONT                = 48
	WM_SETHOTKEY              = 50
	WM_SETICON                = 128
	WM_SETREDRAW              = 11
	WM_SETTEXT                = 12
	WM_SETTINGCHANGE          = 26
	WM_SHOWWINDOW             = 24
	WM_SIZE                   = 5
	WM_SIZECLIPBOARD          = 779
	WM_SIZING                 = 532
	WM_SPOOLERSTATUS          = 42
	WM_STYLECHANGED           = 125
	WM_STYLECHANGING          = 124
	WM_SYSCHAR                = 262
	WM_SYSCOLORCHANGE         = 21
	WM_SYSCOMMAND             = 274
	WM_SYSDEADCHAR            = 263
	WM_SYSKEYDOWN             = 260
	WM_SYSKEYUP               = 261
	WM_TCARD                  = 82
	WM_THEMECHANGED           = 794
	WM_TIMECHANGE             = 30
	WM_TIMER                  = 275
	WM_UNDO                   = 772
	WM_USER                   = 1024
	WM_USERCHANGED            = 84
	WM_VKEYTOITEM             = 46
	WM_VSCROLL                = 277
	WM_VSCROLLCLIPBOARD       = 778
	WM_WINDOWPOSCHANGED       = 71
	WM_WINDOWPOSCHANGING      = 70
	WM_WININICHANGE           = 26
	WM_KEYFIRST               = 256
	WM_KEYLAST                = 264
	WM_SYNCPAINT              = 136
	WM_MOUSEACTIVATE          = 33
	WM_MOUSEMOVE              = 512
	WM_LBUTTONDOWN            = 513
	WM_LBUTTONUP              = 514
	WM_LBUTTONDBLCLK          = 515
	WM_RBUTTONDOWN            = 516
	WM_RBUTTONUP              = 517
	WM_RBUTTONDBLCLK          = 518
	WM_MBUTTONDOWN            = 519
	WM_MBUTTONUP              = 520
	WM_MBUTTONDBLCLK          = 521
	WM_MOUSEWHEEL             = 522
	WM_MOUSEFIRST             = 512
	WM_XBUTTONDOWN            = 523
	WM_XBUTTONUP              = 524
	WM_XBUTTONDBLCLK          = 525
	WM_MOUSELAST              = 525
	WM_MOUSEHOVER             = 0X2A1
	WM_MOUSELEAVE             = 0X2A3
	WM_CLIPBOARDUPDATE        = 0x031D
)

const (
	SW_HIDE            = 0
	SW_NORMAL          = 1
	SW_SHOWNORMAL      = 1
	SW_SHOWMINIMIZED   = 2
	SW_MAXIMIZE        = 3
	SW_SHOWMAXIMIZED   = 3
	SW_SHOWNOACTIVATE  = 4
	SW_SHOW            = 5
	SW_MINIMIZE        = 6
	SW_SHOWMINNOACTIVE = 7
	SW_SHOWNA          = 8
	SW_RESTORE         = 9
	SW_SHOWDEFAULT     = 10
	SW_FORCEMINIMIZE   = 11
)

const (
	SRCCOPY = 0x00CC0020
)

// mouse button constants
const (
	MK_CONTROL  = 0x0008
	MK_LBUTTON  = 0x0001
	MK_MBUTTON  = 0x0010
	MK_RBUTTON  = 0x0002
	MK_SHIFT    = 0x0004
	MK_XBUTTON1 = 0x0020
	MK_XBUTTON2 = 0x0040
)

const (
	CS_VREDRAW         = 0x00000001
	CS_HREDRAW         = 0x00000002
	CS_KEYCVTWINDOW    = 0x00000004
	CS_DBLCLKS         = 0x00000008
	CS_OWNDC           = 0x00000020
	CS_CLASSDC         = 0x00000040
	CS_PARENTDC        = 0x00000080
	CS_NOKEYCVT        = 0x00000100
	CS_NOCLOSE         = 0x00000200
	CS_SAVEBITS        = 0x00000800
	CS_BYTEALIGNCLIENT = 0x00001000
	CS_BYTEALIGNWINDOW = 0x00002000
	CS_GLOBALCLASS     = 0x00004000
	CS_IME             = 0x00010000
	CS_DROPSHADOW      = 0x00020000
)

const (
	IDC_ARROW       = 32512
	IDC_IBEAM       = 32513
	IDC_WAIT        = 32514
	IDC_CROSS       = 32515
	IDC_UPARROW     = 32516
	IDC_SIZENWSE    = 32642
	IDC_SIZENESW    = 32643
	IDC_SIZEWE      = 32644
	IDC_SIZENS      = 32645
	IDC_SIZEALL     = 32646
	IDC_NO          = 32648
	IDC_HAND        = 32649
	IDC_APPSTARTING = 32650
	IDC_HELP        = 32651
	IDC_ICON        = 32641
	IDC_SIZE        = 32640
)

const (
	WS_OVERLAPPED       = 0X00000000
	WS_POPUP            = 0X80000000
	WS_CHILD            = 0X40000000
	WS_MINIMIZE         = 0X20000000
	WS_VISIBLE          = 0X10000000
	WS_DISABLED         = 0X08000000
	WS_CLIPSIBLINGS     = 0X04000000
	WS_CLIPCHILDREN     = 0X02000000
	WS_MAXIMIZE         = 0X01000000
	WS_CAPTION          = 0X00C00000
	WS_BORDER           = 0X00800000
	WS_DLGFRAME         = 0X00400000
	WS_VSCROLL          = 0X00200000
	WS_HSCROLL          = 0X00100000
	WS_SYSMENU          = 0X00080000
	WS_THICKFRAME       = 0X00040000
	WS_GROUP            = 0X00020000
	WS_TABSTOP          = 0X00010000
	WS_MINIMIZEBOX      = 0X00020000
	WS_MAXIMIZEBOX      = 0X00010000
	WS_TILED            = 0X00000000
	WS_ICONIC           = 0X20000000
	WS_SIZEBOX          = 0X00040000
	WS_OVERLAPPEDWINDOW = 0X00000000 | 0X00C00000 | 0X00080000 | 0X00040000 | 0X00020000 | 0X00010000
	WS_POPUPWINDOW      = 0X80000000 | 0X00800000 | 0X00080000
	WS_CHILDWINDOW      = 0X40000000
)

const (
	VK_LBUTTON             = 1
	VK_RBUTTON             = 2
	VK_CANCEL              = 3
	VK_MBUTTON             = 4
	VK_XBUTTON1            = 5
	VK_XBUTTON2            = 6
	VK_BACK                = 8
	VK_TAB                 = 9
	VK_CLEAR               = 12
	VK_RETURN              = 13
	VK_SHIFT               = 16
	VK_CONTROL             = 17
	VK_MENU                = 18
	VK_PAUSE               = 19
	VK_CAPITAL             = 20
	VK_KANA                = 0x15
	VK_HANGEUL             = 0x15
	VK_HANGUL              = 0x15
	VK_JUNJA               = 0x17
	VK_FINAL               = 0x18
	VK_HANJA               = 0x19
	VK_KANJI               = 0x19
	VK_ESCAPE              = 0x1B
	VK_CONVERT             = 0x1C
	VK_NONCONVERT          = 0x1D
	VK_ACCEPT              = 0x1E
	VK_MODECHANGE          = 0x1F
	VK_SPACE               = 32
	VK_PRIOR               = 33
	VK_NEXT                = 34
	VK_END                 = 35
	VK_HOME                = 36
	VK_LEFT                = 37
	VK_UP                  = 38
	VK_RIGHT               = 39
	VK_DOWN                = 40
	VK_SELECT              = 41
	VK_PRINT               = 42
	VK_EXECUTE             = 43
	VK_SNAPSHOT            = 44
	VK_INSERT              = 45
	VK_DELETE              = 46
	VK_HELP                = 47
	VK_LWIN                = 0x5B
	VK_RWIN                = 0x5C
	VK_APPS                = 0x5D
	VK_SLEEP               = 0x5F
	VK_NUMPAD0             = 0x60
	VK_NUMPAD1             = 0x61
	VK_NUMPAD2             = 0x62
	VK_NUMPAD3             = 0x63
	VK_NUMPAD4             = 0x64
	VK_NUMPAD5             = 0x65
	VK_NUMPAD6             = 0x66
	VK_NUMPAD7             = 0x67
	VK_NUMPAD8             = 0x68
	VK_NUMPAD9             = 0x69
	VK_MULTIPLY            = 0x6A
	VK_ADD                 = 0x6B
	VK_SEPARATOR           = 0x6C
	VK_SUBTRACT            = 0x6D
	VK_DECIMAL             = 0x6E
	VK_DIVIDE              = 0x6F
	VK_F1                  = 0x70
	VK_F2                  = 0x71
	VK_F3                  = 0x72
	VK_F4                  = 0x73
	VK_F5                  = 0x74
	VK_F6                  = 0x75
	VK_F7                  = 0x76
	VK_F8                  = 0x77
	VK_F9                  = 0x78
	VK_F10                 = 0x79
	VK_F11                 = 0x7A
	VK_F12                 = 0x7B
	VK_F13                 = 0x7C
	VK_F14                 = 0x7D
	VK_F15                 = 0x7E
	VK_F16                 = 0x7F
	VK_F17                 = 0x80
	VK_F18                 = 0x81
	VK_F19                 = 0x82
	VK_F20                 = 0x83
	VK_F21                 = 0x84
	VK_F22                 = 0x85
	VK_F23                 = 0x86
	VK_F24                 = 0x87
	VK_NUMLOCK             = 0x90
	VK_SCROLL              = 0x91
	VK_LSHIFT              = 0xA0
	VK_RSHIFT              = 0xA1
	VK_LCONTROL            = 0xA2
	VK_RCONTROL            = 0xA3
	VK_LMENU               = 0xA4
	VK_RMENU               = 0xA5
	VK_BROWSER_BACK        = 0xA6
	VK_BROWSER_FORWARD     = 0xA7
	VK_BROWSER_REFRESH     = 0xA8
	VK_BROWSER_STOP        = 0xA9
	VK_BROWSER_SEARCH      = 0xAA
	VK_BROWSER_FAVORITES   = 0xAB
	VK_BROWSER_HOME        = 0xAC
	VK_VOLUME_MUTE         = 0xAD
	VK_VOLUME_DOWN         = 0xAE
	VK_VOLUME_UP           = 0xAF
	VK_MEDIA_NEXT_TRACK    = 0xB0
	VK_MEDIA_PREV_TRACK    = 0xB1
	VK_MEDIA_STOP          = 0xB2
	VK_MEDIA_PLAY_PAUSE    = 0xB3
	VK_LAUNCH_MAIL         = 0xB4
	VK_LAUNCH_MEDIA_SELECT = 0xB5
	VK_LAUNCH_APP1         = 0xB6
	VK_LAUNCH_APP2         = 0xB7
	VK_OEM_1               = 0xBA
	VK_OEM_PLUS            = 0xBB
	VK_OEM_COMMA           = 0xBC
	VK_OEM_MINUS           = 0xBD
	VK_OEM_PERIOD          = 0xBE
	VK_OEM_2               = 0xBF
	VK_OEM_3               = 0xC0
	VK_OEM_4               = 0xDB
	VK_OEM_5               = 0xDC
	VK_OEM_6               = 0xDD
	VK_OEM_7               = 0xDE
	VK_OEM_8               = 0xDF
	VK_OEM_102             = 0xE2
	VK_PROCESSKEY          = 0xE5
	VK_PACKET              = 0xE7
	VK_ATTN                = 0xF6
	VK_CRSEL               = 0xF7
	VK_EXSEL               = 0xF8
	VK_EREOF               = 0xF9
	VK_PLAY                = 0xFA
	VK_ZOOM                = 0xFB
	VK_NONAME              = 0xFC
	VK_PA1                 = 0xFD
	VK_OEM_CLEAR           = 0xFE
)

const (
	MAPVK_VK_TO_VSC    = 0
	MAPVK_VSC_TO_VK    = 1
	MAPVK_VK_TO_CHAR   = 2
	MAPVK_VSC_TO_VK_EX = 3
)

type MSG struct {
	HWnd    HWND
	Message uint32
	WParam  uintptr
	LParam  uintptr
	Time    uint32
	Pt      POINT
}

type POINT struct {
	X, Y int32
}

type RECT struct {
	Left, Top, Right, Bottom int32
}

type BITMAPINFOHEADER struct {
	BiSize          uint32
	BiWidth         int32
	BiHeight        int32
	BiPlanes        uint16
	BiBitCount      uint16
	BiCompression   uint32
	BiSizeImage     uint32
	BiXPelsPerMeter int32
	BiYPelsPerMeter int32
	BiClrUsed       uint32
	BiClrImportant  uint32
}

type RGBQUAD struct {
	RgbBlue     byte
	RgbGreen    byte
	RgbRed      byte
	RgbReserved byte
}

type BITMAPINFO struct {
	BmiHeader BITMAPINFOHEADER
	BmiColors *RGBQUAD
}

var (
	// Library
	libgdi32    syscall.Handle
	libuser32   syscall.Handle
	liboleaut32 syscall.Handle
	libkernel32 syscall.Handle

	stretchDIBits uintptr

	getClientRect    uintptr
	defWindowProc    uintptr
	loadCursor       uintptr
	registerClassEx  uintptr
	adjustWindowRect uintptr
	createWindowEx   uintptr
	showWindow       uintptr
	getDC            uintptr
	getMessage       uintptr
	translateMessage uintptr
	dispatchMessage  uintptr
	invalidateRect   uintptr
	validateRect     uintptr
	sendMessage      uintptr
	postMessage      uintptr
	destroyWindow    uintptr
	mapVirtualKey    uintptr

	sysAllocString uintptr

	getLastError uintptr
)

func init() {
	// Library
	libgdi32, _ = syscall.LoadLibrary("gdi32.dll")
	libuser32, _ = syscall.LoadLibrary("user32.dll")
	liboleaut32, _ = syscall.LoadLibrary("oleaut32.dll")
	libkernel32, _ = syscall.LoadLibrary("kernel32.dll")

	stretchDIBits, _ = syscall.GetProcAddress(libgdi32, "StretchDIBits")

	getClientRect, _ = syscall.GetProcAddress(libuser32, "GetClientRect")
	defWindowProc, _ = syscall.GetProcAddress(libuser32, "DefWindowProcW")
	loadCursor, _ = syscall.GetProcAddress(libuser32, "LoadCursorW")
	registerClassEx, _ = syscall.GetProcAddress(libuser32, "RegisterClassExW")
	adjustWindowRect, _ = syscall.GetProcAddress(libuser32, "AdjustWindowRect")
	createWindowEx, _ = syscall.GetProcAddress(libuser32, "CreateWindowExW")
	showWindow, _ = syscall.GetProcAddress(libuser32, "ShowWindow")
	getDC, _ = syscall.GetProcAddress(libuser32, "GetDC")
	getMessage, _ = syscall.GetProcAddress(libuser32, "GetMessageW")
	translateMessage, _ = syscall.GetProcAddress(libuser32, "TranslateMessage")
	dispatchMessage, _ = syscall.GetProcAddress(libuser32, "DispatchMessageW")
	invalidateRect, _ = syscall.GetProcAddress(libuser32, "InvalidateRect")
	validateRect, _ = syscall.GetProcAddress(libuser32, "ValidateRect")
	sendMessage, _ = syscall.GetProcAddress(libuser32, "SendMessageW")
	postMessage, _ = syscall.GetProcAddress(libuser32, "PostMessageW")
	destroyWindow, _ = syscall.GetProcAddress(libuser32, "DestroyWindow")
	mapVirtualKey, _ = syscall.GetProcAddress(libuser32, "MapVirtualKeyW")

	sysAllocString, _ = syscall.GetProcAddress(liboleaut32, "SysAllocString")

	getLastError, _ = syscall.GetProcAddress(libkernel32, "GetLastError")
}

func StretchDIBits(hdcDest HDC, nXDest, nYDest, nWidthDest, nHeightDest int32, nXSrc, nYSrc, nWidthSrc, nHeightSrc int32, lpBits uintptr, lpBitsInfo *BITMAPINFO, iUsage uint32, dwRop int64) bool {
	ret, _, _ := syscall.Syscall15(stretchDIBits, 13,
		uintptr(hdcDest),
		uintptr(nXDest),
		uintptr(nYDest),
		uintptr(nWidthDest),
		uintptr(nHeightDest),
		uintptr(nXSrc),
		uintptr(nYSrc),
		uintptr(nWidthSrc),
		uintptr(nHeightSrc),
		uintptr(lpBits),
		uintptr(unsafe.Pointer(lpBitsInfo)),
		uintptr(iUsage),
		uintptr(dwRop),
		0,
		0)

	return ret != 0
}

func GetClientRect(hWnd HWND, rect *RECT) bool {
	ret, _, _ := syscall.Syscall(getClientRect, 2,
		uintptr(hWnd),
		uintptr(unsafe.Pointer(rect)),
		0)

	return ret != 0
}

func DefWindowProc(hWnd HWND, Msg uint32, wParam, lParam uintptr) uintptr {
	ret, _, _ := syscall.Syscall6(defWindowProc, 4,
		uintptr(hWnd),
		uintptr(Msg),
		wParam,
		lParam,
		0,
		0)

	return ret
}

func LoadCursor(hInstance HINSTANCE, lpCursorName *uint16) HCURSOR {
	ret, _, _ := syscall.Syscall(loadCursor, 2,
		uintptr(hInstance),
		uintptr(unsafe.Pointer(lpCursorName)),
		0)

	return HCURSOR(ret)
}

func RegisterClassEx(windowClass *WNDCLASSEX) ATOM {
	ret, _, _ := syscall.Syscall(registerClassEx, 1,
		uintptr(unsafe.Pointer(windowClass)),
		0,
		0)

	return ATOM(ret)
}

func AdjustWindowRect(lpRect *RECT, dwStyle uint32, bMenu bool) bool {
	ret, _, _ := syscall.Syscall(adjustWindowRect, 3,
		uintptr(unsafe.Pointer(lpRect)),
		uintptr(dwStyle),
		uintptr(BoolToBOOL(bMenu)))

	return ret != 0
}

func CreateWindowEx(dwExStyle uint32, lpClassName, lpWindowName *uint16, dwStyle uint32, x, y, nWidth, nHeight int32, hWndParent HWND, hMenu HMENU, hInstance HINSTANCE, lpParam unsafe.Pointer) HWND {
	ret, _, _ := syscall.Syscall12(createWindowEx, 12,
		uintptr(dwExStyle),
		uintptr(unsafe.Pointer(lpClassName)),
		uintptr(unsafe.Pointer(lpWindowName)),
		uintptr(dwStyle),
		uintptr(x),
		uintptr(y),
		uintptr(nWidth),
		uintptr(nHeight),
		uintptr(hWndParent),
		uintptr(hMenu),
		uintptr(hInstance),
		uintptr(lpParam))

	return HWND(ret)
}

func ShowWindow(hWnd HWND, nCmdShow int32) bool {
	ret, _, _ := syscall.Syscall(showWindow, 2,
		uintptr(hWnd),
		uintptr(nCmdShow),
		0)

	return ret != 0
}

func GetDC(hWnd HWND) HDC {
	ret, _, _ := syscall.Syscall(getDC, 1,
		uintptr(hWnd),
		0,
		0)

	return HDC(ret)
}

func GetMessage(msg *MSG, hWnd HWND, msgFilterMin, msgFilterMax uint32) BOOL {
	ret, _, _ := syscall.Syscall6(getMessage, 4,
		uintptr(unsafe.Pointer(msg)),
		uintptr(hWnd),
		uintptr(msgFilterMin),
		uintptr(msgFilterMax),
		0,
		0)

	return BOOL(ret)
}

func TranslateMessage(msg *MSG) bool {
	ret, _, _ := syscall.Syscall(translateMessage, 1,
		uintptr(unsafe.Pointer(msg)),
		0,
		0)

	return ret != 0
}

func DispatchMessage(msg *MSG) uintptr {
	ret, _, _ := syscall.Syscall(dispatchMessage, 1,
		uintptr(unsafe.Pointer(msg)),
		0,
		0)

	return ret
}

func InvalidateRect(hWnd HWND, lpRect *RECT, bErase bool) bool {
	ret, _, _ := syscall.Syscall(invalidateRect, 3,
		uintptr(hWnd),
		uintptr(unsafe.Pointer(lpRect)),
		uintptr(BoolToBOOL(bErase)))

	return ret != 0
}

func ValidateRect(hWnd HWND, lpRect *RECT) bool {
	ret, _, _ := syscall.Syscall(validateRect, 2,
		uintptr(hWnd),
		uintptr(unsafe.Pointer(lpRect)),
		0)

	return ret != 0
}

func SendMessage(hWnd HWND, msg uint32, wParam, lParam uintptr) uintptr {
	ret, _, _ := syscall.Syscall6(sendMessage, 4,
		uintptr(hWnd),
		uintptr(msg),
		wParam,
		lParam,
		0,
		0)

	return ret
}

func PostMessage(hWnd HWND, msg uint32, wParam, lParam uintptr) uintptr {
	ret, _, _ := syscall.Syscall6(postMessage, 4,
		uintptr(hWnd),
		uintptr(msg),
		wParam,
		lParam,
		0,
		0)

	return ret
}

func DestroyWindow(hWnd HWND) uintptr {
	ret, _, _ := syscall.Syscall(destroyWindow, 1,
		uintptr(hWnd),
		0,
		0)

	return ret
}

func MapVirtualKey(code uint32, mapType uint32) uint32 {
	ret, _, _ := syscall.Syscall(mapVirtualKey, 2,
		uintptr(code),
		uintptr(mapType),
		0)

	return uint32(ret)
}

func BoolToBOOL(value bool) BOOL {
	if value {
		return 1
	}

	return 0
}

func SysAllocString(s string) *uint16 /*BSTR*/ {
	ret, _, _ := syscall.Syscall(sysAllocString, 1,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(s))),
		0,
		0)

	return (*uint16) /*BSTR*/ (unsafe.Pointer(ret))
}

func StringToBSTR(value string) *uint16 {
	return SysAllocString(value)
}

func GetLastError() uint32 {
	ret, _, _ := syscall.Syscall(getLastError, 0,
		0,
		0,
		0)

	return uint32(ret)
}
