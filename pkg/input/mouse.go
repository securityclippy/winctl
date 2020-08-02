package input

import (
	"errors"
	"fmt"
	"github.com/securityclippy/winctl/pkg/proc"
	"github.com/securityclippy/winctl/pkg/syskey"
	"syscall"
	"unsafe"
)

func SendMouseEvent(hwnd syscall.Handle, event uint32) bool {
	ret, _, _ := proc.SendMessageA.Call(uintptr(hwnd), uintptr(event), 0, 0)
	return ret != 0
}

func MoveMouseToLocation(hwnd syscall.Handle, loc *POINT) error {
	return nil
}

func SendLeftClickLocation (hwnd syscall.Handle, loc *POINT) error {

	ret, _, _ := proc.PostMessageA.Call(uintptr(hwnd), syskey.WM_LBUTTONDOWN, 0, uintptr(unsafe.Pointer(loc)))
	fmt.Println(ret)

	ret, _, _ = proc.PostMessageA.Call(uintptr(hwnd), syskey.WM_LBUTTONUP, 0, uintptr(unsafe.Pointer(loc)))

	if ret != 0 {
		return nil
	}

	return errors.New("failed to send left click to location")
}


func SendLeftClick(hwnd syscall.Handle, debug bool) error {

	loc, err := GetCursorLocation()
	if err != nil {
		return err
	}
	//ret, _, _ := proc.PostMessageA.Call(uintptr(hwnd), syskey.MOUSEEVENTF_LEFTDOWN, syskey.VK_LBUTTON, 0)
	//fmt.Println("down:", ret)
	//ret, _, _ = proc.PostMessageA.Call(uintptr(hwnd), syskey.MOUSEEVENTF_LEFTUP, 0, 0)
	//fmt.Println("up", ret)

	if debug {
		fmt.Printf("sent left click to: %+v\n", loc)
	}

	ret, _, _ := proc.PostMessageA.Call(uintptr(hwnd), syskey.WM_LBUTTONDOWN, 0, uintptr(unsafe.Pointer(loc)))
	fmt.Println(ret)

	ret, _, _ = proc.PostMessageA.Call(uintptr(hwnd), syskey.WM_LBUTTONUP, 0, uintptr(unsafe.Pointer(loc)))
	fmt.Println(ret)
	return nil

}

type POINT struct {
	X int32
	Y int32
}

func GetCursorLocation() (*POINT, error) {
	p := &POINT{}
	ret, _, _ := proc.User32GetCursorPos.Call(uintptr(unsafe.Pointer(p)))

	if ret != 0 {
		return p, nil
	}

	return nil, errors.New("failed to get cursor position")
}

func GetCursorRelativeLocation(hwnd syscall.Handle, relPos *POINT) (*POINT, error) {
	ret, _, _ := proc.User32ScreenToClient.Call(uintptr(hwnd), uintptr(unsafe.Pointer(relPos)))

	if ret != 0 {
		return relPos, nil
	}

	return nil, errors.New("failed to get relative cursor position")
}