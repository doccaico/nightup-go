package nightup

import (
	"fmt"
	"os/exec"
	"strings"
)

func GoInstall(install_path string) {
	// dlを取得する
	cmd1 := exec.Command("curl", "-fsSOL", "https://go.dev/dl/?mode=json")
	err := cmd1.Run()
	if err != nil {
		panic(err)
	}
	fmt.Println("Download (dl) is done")

	// Tarnameを取得する (go1.24.5.windows-amd64.zip)
	out, err := exec.Command("cmd", "/c", "type dl | jq -r .[0].files[]^|select(.os==\"windows\"and.arch==\"amd64\"and.kind==\"archive\").filename").Output()
	if err != nil {
		panic(err)
	}
	tarname := strings.TrimRight(string(out), "\r\n")
	fmt.Println("Get tarname =>", tarname)

	// ダウンロードURLを設定する
	url := fmt.Sprintf("https://go.dev/dl/%s", tarname)
	fmt.Println("Download URL =>", url)

	// ZIPをダウンロードする
	cmd2 := exec.Command("curl", "-fsSOL", url)
	err = cmd2.Run()
	if err != nil {
		panic(err)
	}
	fmt.Println("Download (ZIP) is done")

	// 解凍する
	cmd3 := exec.Command("7za", "x", "-aoa", tarname, "-bso0", "-bsp0")
	err = cmd3.Run()
	if err != nil {
		panic(err)
	}
	fmt.Println("Extraction is done")

	// ディレクトリを移動する
	// For debug
	// cmd4 := exec.Command("cmd", "/c", fmt.Sprintf("move go %s > nul", "C:\\go"))
	cmd4 := exec.Command("cmd", "/c", fmt.Sprintf("move go %s > nul", install_path))
	err = cmd4.Run()
	if err != nil {
		panic(err)
	}
	fmt.Println("Moved:", install_path)

	// ZIPとgo.jsonを削除する
	cmd5 := exec.Command("cmd", "/c", fmt.Sprintf("del %s", tarname))
	err = cmd5.Run()
	if err != nil {
		panic(err)
	}
	cmd6 := exec.Command("cmd", "/c", "del dl")
	err = cmd6.Run()
	if err != nil {
		panic(err)
	}
	fmt.Println("Unnecessary files deleted")
}
