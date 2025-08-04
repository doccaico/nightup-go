package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"gopkg.in/ini.v1"

	"github.com/doccaico/nightup-go"
)

func usage() {
	fmt.Println("Usage:\n  nightup LANG")
	fmt.Println("\nRequirements:")
	fmt.Println("\tjq, 7za, curl")
	fmt.Println("\nSupported languages:")
	fmt.Println("\tzig, odin, v, go")
	os.Exit(0)
}

func getInstallPath(ini_path string, lang string) string {
	cfg, err := ini.Load(ini_path)
	if err != nil {
		panic(err)
	}
	var path = cfg.Section(nightup.SectionName).Key(lang).String()
	return path
}

func checkRequirements() {
	for _, exe := range [...]string{"jq", "7za", "curl"} {
		_, err := exec.LookPath(exe)
		if err != nil {
			panic(err)
		}
	}
}

func main() {

	if len(os.Args) != 2 {
		usage()
	}

	checkRequirements()

	home_path, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	download_path := filepath.Join(home_path, "Downloads")
	ini_path := filepath.Join(home_path, nightup.IniFileName)
	lang := os.Args[1]

	// cd ~/Downloadsする
	err = os.Chdir(download_path)
	if err != nil {
		panic(err)
	}

	switch lang {
	case "zig":
		var install_path = getInstallPath(ini_path, "zig")
		nightup.RemoveDirIfExist(install_path)
		nightup.ZigInstall(install_path)
	case "odin":
		var install_path = getInstallPath(ini_path, "odin")
		nightup.RemoveDirIfExist(install_path)
		nightup.OdinInstall(install_path)
	case "v":
		var install_path = getInstallPath(ini_path, "v")
		nightup.RemoveDirIfExist(install_path)
		nightup.VInstall(install_path)
	case "-h", "--help":
		usage()
	default:
		fmt.Printf("error: unsupported lang: %#q\n\n", lang)
		usage()
	}
}
