package printpn

import (
	"github.com/01-edu/z01"
	"os"
)

func PrintProgramName() {
	dir := os.Args[0]
	var filename string
	for i := 0; i < len(dir); i++ {
		if dir[i] == '\\' {
			filename = ""
			continue
		} else if dir[i] == '.' {
			for _, e := range filename {
				z01.PrintRune(e)
			}
			break
		}
		filename += string(dir[i])
	}
}
