package main

import "github.com/01-edu/z01"

type Door struct {
	state bool
}

const(
	OPEN bool = true
	CLOSE bool = false
)

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func OpenDoor(ptrDoor *Door) bool {
	PrintStr("Door Opening...")
	return ptrDoor.state
}

func CloseDoor(ptrDoor *Door) bool {
	PrintStr("Door Closing...")
	return !ptrDoor.state
}

func IsDoorOpen(ptrDoor *Door) bool {
	PrintStr("is the Door opened ?")
	return ptrDoor.state
}

func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("is the Door closed ?")
	return !ptrDoor.state
}

func main() {
	door := &Door{}
	if IsDoorClose(door) {
		OpenDoor(door)
	} else if IsDoorOpen(door) {
		CloseDoor(door)
	}
}
