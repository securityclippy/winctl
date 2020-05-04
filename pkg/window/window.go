package window

import (
	"fmt"
	"github.com/securityclippy/winctl/pkg/proc"
	"syscall"
	"unsafe"
)

func EnumWindows(enumFunc uintptr, lparam uintptr) (err error) {
	r1, _, e1 := syscall.Syscall(proc.EnumWindows.Addr(), 2, uintptr(enumFunc), uintptr(lparam), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func GetText(hwnd syscall.Handle, str *uint16, maxCount int32) (len int32, err error) {
	r0, _, e1 := syscall.Syscall(proc.GetWindowTextW.Addr(), 3, uintptr(hwnd), uintptr(unsafe.Pointer(str)), uintptr(maxCount))
	len = int32(r0)
	if len == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func FindByTitle(title string) (syscall.Handle, error) {
	var hwnd syscall.Handle
	cb := syscall.NewCallback(func(h syscall.Handle, p uintptr) uintptr {
		b := make([]uint16, 200)
		_, err := GetText(h, &b[0], int32(len(b)))
		if err != nil {
			// ignore the error
			return 1 // continue enumeration
		}
		if syscall.UTF16ToString(b) == title {
			// note the window
			hwnd = h
			return 0 // stop enumeration
		}
		return 1 // continue enumeration
	})
	EnumWindows(cb, 0)
	if hwnd == 0 {
		return 0, fmt.Errorf("No window with title '%s' found", title)
	}
	return hwnd, nil
}


func SetTitle(hwnd syscall.Handle, title string) uintptr {
	s, _ := syscall.UTF16PtrFromString(title)
	ret, _, _ := syscall.Syscall(proc.SetWindowTextA.Addr(), 1, uintptr(hwnd), uintptr(unsafe.Pointer(s)), 0)
	return uintptr(ret)
}

const PROCESS_ALL_ACCESS = 0x1F0FFF

func ptr(val interface{}) uintptr {
	switch val.(type) {
	case string:
		return uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(val.(string))))
	case int:
		return uintptr(val.(int))
	default:
		return uintptr(0)
	}
}

func OpenProcessHandle(processId int) syscall.Handle {
	kernel32 := syscall.MustLoadDLL("kernel32.dll")
	p := kernel32.MustFindProc("OpenProcess")
	handle, _, _ := p.Call(ptr(PROCESS_ALL_ACCESS), ptr(true), ptr(processId))
	return syscall.Handle(handle)
}


func GetHandle(cls string, win string) (ret syscall.Handle, err error) {

	cls = "GxWindowClass"
	// class will always be GxWindowClass for now
	lpszClass, err := syscall.UTF16PtrFromString(cls)
	if err != nil {
		return
	}
	lpszWindow, err := syscall.UTF16PtrFromString(win)
	if err != nil {
		return
	}


	r0, _, e1 := syscall.Syscall(proc.FindWindowW.Addr(), 2, uintptr(unsafe.Pointer(lpszClass)), uintptr(unsafe.Pointer(lpszWindow)), 0)
	//ret = HWND(r0)
	if ret == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return syscall.Handle(r0), nil
}