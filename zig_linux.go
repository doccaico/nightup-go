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
	out, err := exec.Command("sh", "-c", "cat index.json | jq -r '.master.\"x86_64-linux\".tarball'").Output()
	if err != nil {
		panic(err)
	}
	url := strings.TrimRight(string(out), "\n")
	fmt.Println("Download URL =>", url)

	// TAR.XZをダウンロードする
	cmd2 := exec.Command("curl", "-fsSOL", url)
	err = cmd2.Run()
	if err != nil {
		panic(err)
	}
	fmt.Println("Download (TAR.XZ) is done")

	// 解凍する
	tarname := filepath.Base(url)
	cmd3 := exec.Command("tar", "xf", tarname)
	err = cmd3.Run()
	if err != nil {
		panic(err)
	}
	fmt.Println("Extraction is done")

	// インストール先のディレクトリが存在したら削除する
	if _, err := os.Stat(install_path); !os.IsNotExist(err) {
		cmd4 := exec.Command("rm", "-rf", install_path)
		err = cmd4.Run()
		if err != nil {
			panic(err)
		}
		fmt.Println("Removed:", install_path)
	}

	// ディレクトリを移動する
	src := strings.TrimSuffix(tarname, ".tar.xz")
	cmd5 := exec.Command("mv", "-f", src, install_path)
	err = cmd5.Run()
	if err != nil {
		panic(err)
	}
	fmt.Println("Moved:", install_path)

	// TAR.XZとindex.jsonを削除する
	cmd6 := exec.Command("rm", tarname)
	err = cmd6.Run()
	if err != nil {
		panic(err)
	}
	cmd7 := exec.Command("rm", "index.json")
	err = cmd7.Run()
	if err != nil {
		panic(err)
	}
	fmt.Println("Unnecessary files deleted")
}
