package notifyicon

import (
	"fmt"
	"keepGuestOn/pkg/win"
	"log"
	"unsafe"

	"golang.org/x/sys/windows"
)

func wndProc(hWnd uintptr, msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	case notifyIconMsg:
		switch nmsg := win.LOWORD(uint32(lParam)); nmsg {
		case win.NIN_BALLOONUSERCLICK:
			log.Print("User has clicked the balloon message")
		case win.WM_LBUTTONDOWN:
			clickHandler()
		}
	case win.WM_DESTROY:
		win.PostQuitMessage(0)
	default:
		return win.DefWindowProc(hWnd, msg, wParam, lParam)
	}
	return 0
}

func createMainWindow() (uintptr, error) {
	hInstance := win.GetModuleHandle(nil)

	wndClass := windows.StringToUTF16Ptr("MyWindow")

	var wcex win.WNDCLASSEX
	wcex.CbSize = uint32(unsafe.Sizeof(wcex))
	wcex.Style = win.CS_HREDRAW | win.CS_VREDRAW
	wcex.LpfnWndProc = windows.NewCallback(wndProc)
	wcex.HInstance = hInstance
	wcex.HCursor = win.LoadCursor(0, win.MAKEINTRESOURCE(win.IDC_ARROW))
	wcex.HbrBackground = win.COLOR_WINDOW + 1
	wcex.LpszClassName = wndClass
	if win.RegisterClassEx(&wcex) == 0 {
		return 0, win.GetLastError()
	}

	hwnd := win.CreateWindowEx(0, wndClass, windows.StringToUTF16Ptr("NotifyIcon Example"), win.WS_OVERLAPPEDWINDOW, win.CW_USEDEFAULT, win.CW_USEDEFAULT, 400, 300, 0, 0, hInstance, nil)
	if hwnd == win.NULL {
		return 0, win.GetLastError()
	}
//	win.ShowWindow(hwnd, win.SW_SHOW)

	return hwnd, nil
}

func loadIconFromResource(id uintptr) (uintptr, error) {
	hIcon := win.LoadImage(
		win.GetModuleHandle(nil),
		win.MAKEINTRESOURCE(id),
		win.IMAGE_ICON,
		0, 0,
		win.LR_DEFAULTSIZE)
	if hIcon == win.NULL {
		return 0, win.GetLastError()
	}

	return hIcon, nil
}
func loadIconFromFile(name string) (uintptr, error) {
	hIcon := win.LoadImage(
		win.NULL,
		windows.StringToUTF16Ptr(name),
		win.IMAGE_ICON,
		0, 0,
		win.LR_DEFAULTSIZE|win.LR_LOADFROMFILE)
	if hIcon == win.NULL {
		return 0, win.GetLastError()
	}

	return hIcon, nil
}

func clickHandler() {
	log.Print("User has clicked the notify icon")
}

func StartNotification(header string, body string, info string) (bool){
	hIcon, err := loadIconFromResource(10) // rsrc uses 10 for icon resource id
	if err != nil {
		hIcon, err = loadIconFromFile("icon.ico") // fallback to use file
		if err != nil {
			fmt.Print(err.Error())
			//panic(err)
		}
	}
	defer win.DestroyIcon(hIcon)
	ni, err := newNotifyIcon()
	if err != nil {
		panic(err)
	}
	defer ni.Dispose()
	//defer os.Exit(0)

	ni.SetIcon(hIcon)
	ni.SetTooltip(header)
	ni.ShowNotificationWithIcon(body, info, hIcon)
	return true; // created msg

}
