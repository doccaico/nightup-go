# nightup
My updater tool written in Go.

## Dependencies
- [jq](https://github.com/jqlang/jq)
- 7zip (7za)
- curl

## How to use
```sh
$ nightup [lang/software]
```

## Supported languages/software
- [zig](https://github.com/ziglang/zig)
- [odin](https://github.com/odin-lang/Odin)
- [v](https://github.com/vlang/v)
- [go](https://github.com/golang/go)
- [vim](https://github.com/vim/vim-win32-installer)(~/Downloadsへダウンロードしてエクスプローラーを開くだけ)

## Build
```sh
# Windows
$ build.bat

# WSL2 (or Linux Native?)
$ ./build.sh
```

## Config(INI file)
```ini
; ~/.nightup
[Windows]
zig=C:\Langs\zig
odin=C:\Langs\odin
v=C:\Langs\v
go=C:\Langs\go

[Linux]
zig=/home/doccaico/langs/zig
odin=/home/doccaico/langs/odin
v=/home/doccaico/langs/v
go=/home/doccaico/langs/go
```
