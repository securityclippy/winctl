package input


var (
	//https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-mouseinput
	MOUSEEVENTF_ABSOLUTE = 0x8000
	MOUSEEVENTF_HWHEEL = 0x01000
	MOUSEEVENTF_MOVE = 0x0001
	MOUSEEVENTF_MOVE_NOCOALESCE = 0x2000
	MOUSEEVENTF_LEFTDOWN = 0x0002
	MOUSEEVENTF_LEFTUP = 0x0004
	MOUSEEVENTF_RIGHTDOWN = 0x0008
	MOUSEEVENTF_RIGHTUP = 0x0010
	MOUSEEVENTF_MIDDLEDOWN = 0x0020
	MOUSEEVENTF_MIDDLEUP = 0x0040
	MOUSEEVENTF_VIRTUALDESK = 0x4000
	MOUSEEVENTF_WHEEL = 0x0800
	MOUSEEVENTF_XDOWN = 0x0080
	MOUSEEVENTF_XUP = 0x0100
)

type Input struct {

}

type MouseInput struct {
	//The absolute position of the mouse, or
	// the amount of motion since the last mouse
	// event was generated, depending on the value of the
	// dwFlags member. Absolute data is specified as the
	// x coordinate of the mouse;
	// relative data is specified as the number of pixels moved.
	Dx int32
	Dy int32
	MouseData uint32

	//If dwFlags does not contain
	// MOUSEEVENTF_WHEEL, MOUSEEVENTF_XDOWN, or MOUSEEVENTF_XUP,
	// then mouseData should be zero.
	DwFlags uint32
	Time uint32
	DwExtraInfo uintptr
}