package nightup

import (
	"fmt"
	"os/exec"
	"strings"
)

func OdinInstall(install_path string) {
	// nightly.jsonを取得する
	cmd1 := exec.Command("curl", "-fsSOL", "https://f001.backblazeb2.com/file/odin-binaries/nightly.json")
	err := cmd1.Run()
	if err != nil {
		panic(err)
	}
	fmt.Println("Download (nightly.json) is done")

	// ダウンロードURLを取得する
	out, err := exec.Command("cmd", "/c", "type nightly.json | jq -r .last_updated").Output()
	if err != nil {
		panic(err)
	}
	url := strings.TrimRight(string(out), "\r\n")
	// "2025-08-03T20:14:55.584916+00:00"から"2025-08-03"を取得する
	date := url[:len("yyyy-mm-dd")]
	// ファイル名を設定する
	tarname := fmt.Sprintf("odin-windows-amd64-nightly%%2B%s.zip", date)
	// 最終的なダウンロード先のURLを作成する
	dl_url := fmt.Sprintf("https://f001.backblazeb2.com/file/odin-binaries/nightly/%s", tarname)
	fmt.Println("Download URL =>", dl_url)

	// ZIPをダウンロードする
	cmd2 := exec.Command("curl", "-fsSOL", dl_url)
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

	// ディレクトリ名を変更する(dist -> odin)
	cmd4 := exec.Command("cmd", "/c", "move dist odin > nul")
	err = cmd4.Run()
	if err != nil {
		panic(err)
	}
	fmt.Println("Rename is done (dist -> odin)")

	// ディレクトリを移動する
	cmd5 := exec.Command("cmd", "/c", fmt.Sprintf("move odin %s > nul", install_path))
	err = cmd5.Run()
	if err != nil {
		panic(err)
	}
	fmt.Println("Moved:", install_path)

	// ZIPとnightly.jsonを削除する
	cmd6 := exec.Command("cmd", "/c", fmt.Sprintf("del %s", tarname))
	err = cmd6.Run()
	if err != nil {
		panic(err)
	}
	cmd7 := exec.Command("cmd", "/c", "del nightly.json")
	err = cmd7.Run()
	if err != nil {
		panic(err)
	}
	fmt.Println("Unnecessary files deleted")
}
