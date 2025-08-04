package nightup

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func ZigInstall(install_path string) {
	// index.jsonを取得する
	cmd1 := exec.Command("curl", "-fsSOL", "https://ziglang.org/download/index.json")
	err := cmd1.Run()
	if err != nil {
		panic(err)
	}
	fmt.Println("Download (index.json) is done")

	// ダウンロードURLを取得する
	out, err := exec.Command("cmd", "/c", "type index.json | jq -r .master.\"x86_64-windows\".tarball").Output()
	if err != nil {
		panic(err)
	}
	url := strings.TrimRight(string(out), "\r\n")
	fmt.Println("Download URL =>", url)

	// ZIPをダウンロードする
	cmd2 := exec.Command("curl", "-fsSOL", url)
	err = cmd2.Run()
	if err != nil {
		panic(err)
	}
	fmt.Println("Download (ZIP) is done")

	// 解凍する
	tarname := filepath.Base(url)
	cmd3 := exec.Command("7za", "x", "-aoa", tarname, "-bso0", "-bsp0")
	err = cmd3.Run()
	if err != nil {
		panic(err)
	}
	fmt.Println("Extraction is done")

	// インストール先のディレクトリが存在したら削除する
	if _, err := os.Stat(install_path); !os.IsNotExist(err) {
		cmd4 := exec.Command("cmd", "/c", fmt.Sprintf("rmdir /s /q %s", install_path))
		err = cmd4.Run()
		if err != nil {
			panic(err)
		}
		fmt.Println("Removed:", install_path)
	}

	// ディレクトリを移動する
	src := strings.TrimSuffix(tarname, ".zip")
	cmd5 := exec.Command("cmd", "/c", fmt.Sprintf("move %s %s > nul", src, install_path))
	err = cmd5.Run()
	if err != nil {
		panic(err)
	}
	fmt.Println("Moved:", install_path)

	// ZIPとindex.jsonを削除する
	cmd6 := exec.Command("cmd", "/c", fmt.Sprintf("del %s", tarname))
	err = cmd6.Run()
	if err != nil {
		panic(err)
	}
	cmd7 := exec.Command("cmd", "/c", "del index.json")
	err = cmd7.Run()
	if err != nil {
		panic(err)
	}
	fmt.Println("Unnecessary files deleted")
}
