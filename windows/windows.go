package windows

import (
	"syscall"
	"unsafe"
)

const (
	/*
	   十六进制，高四位为背景色低4位为前景色
	   0 = 黑色       8 = 灰色
	   1 = 蓝色       9 = 淡蓝色
	   2 = 绿色       A = 淡绿色
	   3 = 湖蓝色     B = 淡浅绿色
	   4 = 红色       C = 淡红色
	   5 = 紫色       D = 淡紫色
	   6 = 黄色       E = 淡黄色
	   7 = 白色       F = 亮白色
	*/
	COLOR_BLANK     = 0x00
	COLOR_BLUE      = 0x01
	COLOR_GREEN     = 0x02
	COLOR_RED       = 0x04
	COLOR_INTENSITY = 0x08

	FOREGROUND_BLUE      = 0x01
	FOREGROUND_GREEN     = 0x02
	FOREGROUND_RED       = 0x04
	FOREGROUND_INTENSITY = 0x08

	BACKGROUND_BLUE      = 0x10
	BACKGROUND_GREEN     = 0x20
	BACKGROUND_RED       = 0x40
	BACKGROUND_INTENSITY = 0x80
)

type COORD struct {
	X int16
	Y int16
}
type SMALL_RECT struct {
	Left   int16
	Top    int16
	Right  int16
	Bottom int16
}
type CONSOLE_SCREEN_BUFFER_INFO struct {
	dwSize              COORD
	dwCursorPosition    COORD
	wAttributes         uint16
	srWindow            SMALL_RECT
	dwMaximumWindowSize COORD
}

func SetConsoleTitle(title string) (int, error) {
	kernel32, err := syscall.LoadLibrary("kernel32.dll")
	defer syscall.FreeLibrary(kernel32)

	setConsoleTitle, err := syscall.GetProcAddress(kernel32, "SetConsoleTitleW")
	ret, _, err := syscall.Syscall(setConsoleTitle, 1, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))), 0, 0)

	return int(ret), err
}
func SetConsoleTextColor(color int) (int, error) {
	handler := syscall.Stdout
	kernel32, err := syscall.LoadLibrary("kernel32.dll")
	defer syscall.FreeLibrary(kernel32)

	setConsoleTextAttribute, err := syscall.GetProcAddress(kernel32, "SetConsoleTextAttribute")
	ret, _, err := syscall.Syscall(setConsoleTextAttribute, 2, uintptr(handler), uintptr(color), 0)
	return int(ret), err
}
func Clear() (int, error) {
	handler := syscall.Stdout
	buffer := CONSOLE_SCREEN_BUFFER_INFO{}

	kernel32, err := syscall.LoadLibrary("kernel32.dll")
	defer syscall.FreeLibrary(kernel32)

	getConsoleScreenBufferInfo, err := syscall.GetProcAddress(kernel32, "GetConsoleScreenBufferInfo")
	ret, _, err := syscall.Syscall(getConsoleScreenBufferInfo, 2, uintptr(handler), uintptr(unsafe.Pointer(&buffer)), 0)

	dwConSize := buffer.dwSize.X * buffer.dwSize.Y

	fillConsoleOutputCharacter, err := syscall.GetProcAddress(kernel32, "FillConsoleOutputCharacterA")
	ret, _, err = syscall.Syscall6(fillConsoleOutputCharacter, 5, uintptr(handler), uintptr(0x20), uintptr(dwConSize), 0, uintptr(unsafe.Pointer(&[...]uint32{0})), 0)

	fillConsoleOutputAttribute, err := syscall.GetProcAddress(kernel32, "FillConsoleOutputAttribute")
	ret, _, err = syscall.Syscall6(fillConsoleOutputAttribute, 5, uintptr(handler), uintptr(buffer.wAttributes), uintptr(dwConSize), 0, uintptr(unsafe.Pointer(&[...]uint32{0})), 0)

	setConsoleCursorPosition, err := syscall.GetProcAddress(kernel32, "SetConsoleCursorPosition")
	ret, _, err = syscall.Syscall(setConsoleCursorPosition, 2, uintptr(handler), uintptr(0), 0)

	return int(ret), err
}
