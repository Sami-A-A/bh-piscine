package printpn

import (
	"os"

	"github.com/01-edu/z01"
)

func PrintProgramName() {
	dir := os.Args[0]
	var filename string
	for i := 0; i < len(dir); i++ {
		if dir[i] == '/' {
			filename = ""
			continue
		} else if dir[i] == '.' && dir[i+1] != '/' || dir[i+1] != '.' {
			for _, e := range filename {
				z01.PrintRune(e)
			}
			for _, e := range dir {
				z01.PrintRune(e)
			}
			break
		}
		filename += string(dir[i])
	}
}
