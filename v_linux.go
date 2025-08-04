package nightup

import (
	"fmt"
	"os/exec"
	"strings"
)

func VInstall(install_path string) {
	// latestを取得する
	cmd1 := exec.Command("curl", "-fsSOL", "https://api.github.com/repos/vlang/v/releases/latest")
	err := cmd1.Run()
	if err != nil {
		panic(err)
	}
	fmt.Println("Download (latest) is done")

	// ダウンロードURLを取得する
	out, err := exec.Command("sh", "-c", "cat latest | jq -r '.assets[] | select(.name == \"v_linux.zip\").browser_download_url'").Output()
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
	cmd3 := exec.Command("7za", "x", "-aoa", "v_linux.zip", "-bso0", "-bsp0")
	err = cmd3.Run()
	if err != nil {
		panic(err)
	}
	fmt.Println("Extraction is done")

	// ディレクトリを移動する
	cmd4 := exec.Command("mv", "-f", "v", install_path)
	err = cmd4.Run()
	if err != nil {
		panic(err)
	}
	fmt.Println("Moved:", install_path)

	// ZIPとlatestを削除する
	cmd5 := exec.Command("rm", "v_linux.zip")
	err = cmd5.Run()
	if err != nil {
		panic(err)
	}
	cmd6 := exec.Command("rm", "latest")
	err = cmd6.Run()
	if err != nil {
		panic(err)
	}
	fmt.Println("Unnecessary files deleted")
}
