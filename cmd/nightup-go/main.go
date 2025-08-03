package main

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/ini.v1"

	"github.com/doccaico/nightup-go"
)

func usage() {
	fmt.Println("Usage:\n nightup LANG")
	fmt.Println("\nSupported languages:")
	fmt.Println("\tzig, odin, v, go")
	os.Exit(0)
}

func getInstallPath(ini_path string, lang string) {
	// ini_pathが存在するかどうかを確認して、存在しなければパニック
	if _, err := os.Stat(ini_path); os.IsNotExist(err) {
		panic("`" + ini_path + "` is not exist")
	}

}

func main() {

	if len(os.Args) != 2 {
		usage()
	}
	// fmt.Printf("%d\n", len(os.Args))
	// nightup.Greet()

	home_path, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	// ~/Downloadsが存在するかどうかを確認して、存在しなければパニック
	download_path := filepath.Join(home_path, "Downloads")
	if _, err := os.Stat(download_path); os.IsNotExist(err) {
		panic("`" + download_path + "` is not exist")
	}

	ini_path := filepath.Join(home_path, nightup.IniFileName)
	fmt.Println(ini_path)

	lang := os.Args[1]
	switch lang {
	case "zig":
		fmt.Printf("zig\n")
		getInstallPath(ini_path, "zig")
	case "-h", "--help":
		usage()
	default:
		fmt.Printf("error: unsupported lang: %#q\n\n", lang)
		usage()
	}

}
