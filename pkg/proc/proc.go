package proc

import "syscall"

var (

	User32  = syscall.MustLoadDLL("user32.dll")
	EnumWindows = User32.MustFindProc("EnumWindows")
	GetWindowTextW = User32.MustFindProc("GetWindowTextW")
	SendMessageA = User32.MustFindProc("SendMessageA")
	SetWindowTextA = User32.MustFindProc("SetWindowTextA")
	PostMessageA = User32.MustFindProc("PostMessageA")
	FindWindowW = User32.MustFindProc("FindWindowW")

	User32GetCursorPos = User32.MustFindProc("GetCursorPos")
	User32ScreenToClient = User32.MustFindProc("ScreenToClient")
	GetAsyncKeyState = User32.MustFindProc("GetAsyncKeyState")
)


