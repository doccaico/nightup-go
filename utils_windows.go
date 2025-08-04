package nightup

import (
	"fmt"
	"os"
	"os/exec"
)

func RemoveDirIfExist(install_path string) {
	if _, err := os.Stat(install_path); !os.IsNotExist(err) {
		cmd := exec.Command("cmd", "/c", fmt.Sprintf("rmdir /s /q %s", install_path))
		err = cmd.Run()
		if err != nil {
			panic(err)
		}
		fmt.Println("Removed:", install_path)
	}
}
