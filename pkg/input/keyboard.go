package input

import (
	"github.com/securityclippy/winctl/pkg/proc"
	"strconv"
	"syscall"
	"time"
	"github.com/securityclippy/winctl/pkg/syskey"
)

func sendMessageAKeyDown(hwnd syscall.Handle, key uint32) uintptr {

	r0, _, _ := syscall.Syscall6(proc.SendMessageA.Addr(), 4, uintptr(hwnd), uintptr(syskey.WM_KEYDOWN), uintptr(key), uintptr(0), 0, 0)

	return uintptr(r0)
}

func sendMessageAKeyChar(hwnd syscall.Handle, key uint32) uintptr {

	r0, _, _ := syscall.Syscall6(proc.SendMessageA.Addr(), 4, uintptr(hwnd), uintptr(syskey.WM_CHAR), uintptr(key), uintptr(1), 0, 0)

	return uintptr(r0)
}

func sendMessageAKeyUp(hwnd syscall.Handle, key uint32) uintptr {

	r0, _, _ := syscall.Syscall6(proc.SendMessageA.Addr(), 4, uintptr(hwnd), uintptr(syskey.WM_KEYUP), uintptr(key), uintptr(1), 0, 0)

	return uintptr(r0)

}



func SendMessageA(hwnd syscall.Handle, key uint32, sendChar bool) uintptr {
	ret := sendMessageAKeyDown(hwnd, key)
	if ret != 0 {
		return ret
	}
	if sendChar {
		ret = sendMessageAKeyChar(hwnd, key)
		if ret != 0 {
			return ret

		}
	}

	ret = sendMessageAKeyUp(hwnd, key)
	if ret != 0 {
		return ret
	}

	return 0
}


func SendMessageALongPress(hwnd syscall.Handle, key uint32, duration time.Duration) uintptr {

	timeout := time.After(duration)
	tick := time.Tick(5 * time.Millisecond)

	sendMessageAKeyDown(hwnd, key)
	foo:
	for {
		select {
		case <- timeout:
			break foo
		case <- tick:
			sendMessageAKeyChar(hwnd, key)
		}
	}

	return sendMessageAKeyUp(hwnd, key)

}



func StringToUint32(s string) uint32 {
	result, _ := strconv.ParseUint(s, 10, 64)
	return uint32(result)

}


func postMessageAKeyDown(hwnd syscall.Handle, key uint32) uintptr {

	r0, _, _ := syscall.Syscall6(proc.PostMessageA.Addr(), 4, uintptr(hwnd), uintptr(syskey.WM_KEYDOWN), uintptr(key), uintptr(0), 0, 0)

	return uintptr(r0)
}

func postMessageAKeyChar(hwnd syscall.Handle, key uint32) uintptr {

	r0, _, _ := syscall.Syscall6(proc.PostMessageA.Addr(), 4, uintptr(hwnd), uintptr(syskey.WM_CHAR), uintptr(key), uintptr(1), 0, 0)

	return uintptr(r0)
}

func postMessageAKeyUp(hwnd syscall.Handle, key uint32) uintptr {

	r0, _, _ := syscall.Syscall6(proc.PostMessageA.Addr(), 4, uintptr(hwnd), uintptr(syskey.WM_KEYUP), uintptr(key), uintptr(1), 0, 0)

	return uintptr(r0)

}



func PostMessageA(hwnd syscall.Handle, key uint32, sendChar bool) uintptr {
	ret := postMessageAKeyDown(hwnd, key)
	if ret != 0 {
		return ret
	}
	if sendChar {
		ret = postMessageAKeyChar(hwnd, key)
		if ret != 0 {
			return ret

		}
	}

	ret = postMessageAKeyUp(hwnd, key)
	if ret != 0 {
		return ret
	}

	return 0
}
