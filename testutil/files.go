package testutil

import (
	"io"
	"os"
)

func FileExists(fname string) bool {
	_, err := os.Stat(fname)
	return !os.IsNotExist(err)
}

func ReadAll(fname string) string {
	file, _ := os.Open(fname)       // ファイルを開く
	byteData, _ := io.ReadAll(file) // ファイル全体をメモリへ読み込む
	return string(byteData)
}
