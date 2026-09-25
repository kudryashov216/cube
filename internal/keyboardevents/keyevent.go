package main

import (
	"bufio"
	"cube/internal/output"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	ch                    rune              = 0
	kernel32              *syscall.LazyDLL  = syscall.NewLazyDLL("kernel32.dll")
	procGetStdHandle      *syscall.LazyProc = kernel32.NewProc("GetStdHandle")
	procReadConsoleInputW *syscall.LazyProc = kernel32.NewProc("ReadConsoleInputW")
)

type inputRecord struct {
	EventType uint16
	_         uint16
	Event     keyEvent
}

type keyEvent struct {
	KeyDown         uint32
	RepeatCount     uint16
	VirtualKeyCode  uint16
	VirtualScanCode uint16
	UnicodeChar     uint16
	ControlKeyState uint32
}

const (
	stdInputHandle = -10
	keyEvent_      = 0x0001
)

func GetKeyEvent(ch chan<- uint16, writer *bufio.Writer) {

	h, err := getStdHandle(stdInputHandle)

	if err != nil {
		output.Output("Failed to obtain the stdin handle: " + err.Error() + "\n")
		return
	}

	buf := make([]inputRecord, 1)

	for {
		n, err := readConsoleInput(h, buf)
		if err != nil || n == 0 {
			continue
		}

		rec := buf[0]

		if rec.EventType == keyEvent_ && rec.Event.KeyDown == 1 {

			vk := rec.Event.VirtualKeyCode

			ch <- vk

		}
	}
}

func getStdHandle(stdHandle int) (windows.Handle, error) {
	r1, _, e1 := procGetStdHandle.Call(uintptr(stdHandle))
	if r1 == ^uintptr(0) {
		return 0, e1
	}
	return windows.Handle(r1), nil
}

func readConsoleInput(h windows.Handle, records []inputRecord) (int, error) {
	numRecords := uint32(len(records))
	var numRead uint32
	r1, _, e1 := procReadConsoleInputW.Call(
		uintptr(h),
		uintptr(unsafe.Pointer(&records[0])),
		uintptr(numRecords),
		uintptr(unsafe.Pointer(&numRead)),
	)
	if r1 == 0 {
		return 0, e1
	}
	return int(numRead), nil
}
