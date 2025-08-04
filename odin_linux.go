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
	out, err := exec.Command("sh", "-c", "cat nightly.json | jq -r '.last_updated'").Output()
	if err != nil {
		panic(err)
	}
	url := strings.TrimRight(string(out), "\r\n")
	// "2025-08-03T20:14:55.584916+00:00"から"2025-08-03"を取得する
	date := url[:len("yyyy-mm-dd")]
	// ファイル名を設定する
	tarname := fmt.Sprintf("odin-linux-amd64-nightly%%2B%s.tar.gz", date)
	// 最終的なダウンロード先のURLを作成する
	dl_url := fmt.Sprintf("https://f001.backblazeb2.com/file/odin-binaries/nightly/%s", tarname)
	fmt.Println("Download URL =>", dl_url)

	// TAR.GZをダウンロードする
	cmd2 := exec.Command("curl", "-fsSOL", dl_url)
	err = cmd2.Run()
	if err != nil {
		panic(err)
	}
	fmt.Println("Download (TAR.GZ) is done")

	// 空ディレクトリ(odin)を作成する
	cmd3 := exec.Command("mkdir", "-p", "odin")
	err = cmd3.Run()
	if err != nil {
		panic(err)
	}
	fmt.Println("Create (odin) is done")

	// 解凍する
	cmd4 := exec.Command("tar", "xzf", tarname, "-C", "odin", "--strip-components=1")
	err = cmd4.Run()
	if err != nil {
		panic(err)
	}
	fmt.Println("Extraction is done")

	// ディレクトリを移動する
	cmd5 := exec.Command("mv", "-f", "odin", install_path)
	err = cmd5.Run()
	if err != nil {
		panic(err)
	}
	fmt.Println("Moved:", install_path)

	// TAR.GZとnightly.jsonを削除する
	cmd6 := exec.Command("rm", tarname)
	err = cmd6.Run()
	if err != nil {
		panic(err)
	}
	cmd7 := exec.Command("rm", "nightly.json")
	err = cmd7.Run()
	if err != nil {
		panic(err)
	}
	fmt.Println("Unnecessary files deleted")
}
