package nightup

import (
	"fmt"
	"os/exec"
	"strings"
)

func VimInstall(install_path string) {
	// latestを取得する
	cmd1 := exec.Command("curl", "-fsSOL", "https://api.github.com/repos/vim/vim-win32-installer/releases/latest")
	err := cmd1.Run()
	if err != nil {
		panic(err)
	}
	fmt.Println("Download (latest) is done")

	// ダウンロードURLを取得する
	out, err := exec.Command("cmd", "/c", "type latest | jq -r .assets[]^|select(.name^|endswith(\"_x64.exe\")).browser_download_url").Output()
	if err != nil {
		panic(err)
	}
	url := strings.TrimRight(string(out), "\r\n")
	fmt.Println("Download URL =>", url)

	// EXEをダウンロードする
	cmd2 := exec.Command("curl", "-fsSOL", url)
	err = cmd2.Run()
	if err != nil {
		panic(err)
	}
	fmt.Println("Download (EXE) is done")

	// latestを削除する
	cmd3 := exec.Command("cmd", "/c", "del latest")
	err = cmd3.Run()
	if err != nil {
		panic(err)
	}
	fmt.Println("Unnecessary files deleted")

	// エクスプローラーを開く
	cmd4 := exec.Command("cmd", "/c", "start", "explorer", ".")
	err = cmd4.Run()
	if err != nil {
		panic(err)
	}
	fmt.Println("Opened EXPLORER.EXE")
}
