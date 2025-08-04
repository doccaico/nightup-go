package nightup

import (
	"fmt"
	"os/exec"
	"strings"
)

func GoInstall(install_path string) {
	// dlを取得する
	cmd1 := exec.Command("curl", "-fsSL", "https://go.dev/dl/?mode=json", "-o", "dl")
	err := cmd1.Run()
	if err != nil {
		panic(err)
	}
	fmt.Println("Download (dl) is done")

	// Tarnameを取得する (go1.24.5.linux-amd64.tar.gz)
	out, err := exec.Command("sh", "-c", "cat dl | jq -r '.[0].files[] | select(.os == \"linux\" and .arch == \"amd64\" and .kind == \"archive\").filename'").Output()
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
	fmt.Println("Download (TAR.GZ) is done")

	// 解凍する
	cmd3 := exec.Command("tar", "xzf", tarname)
	err = cmd3.Run()
	if err != nil {
		panic(err)
	}
	fmt.Println("Extraction is done")

	// ディレクトリを移動する
	// For debug
	// cmd4 := exec.Command("mv", "-f", "go", "/home/doccaico/debug/go")
	cmd4 := exec.Command("mv", "-f", "go", install_path)
	err = cmd4.Run()
	if err != nil {
		panic(err)
	}
	fmt.Println("Moved:", install_path)

	// TAR.GZとdlを削除する
	cmd5 := exec.Command("rm", tarname)
	err = cmd5.Run()
	if err != nil {
		panic(err)
	}
	cmd6 := exec.Command("rm", "dl")
	err = cmd6.Run()
	if err != nil {
		panic(err)
	}
	fmt.Println("Unnecessary files deleted")
}
