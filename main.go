package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		// Pesan error ke stderr supaya tidak tercampur dengan output normal (stdout)
		fmt.Fprintln(os.Stderr, "usage: disk-usage-analyzer <path>")
		// Exit code non-zero = tanda gagal, dibaca oleh shell/CI
		os.Exit(1)
	}

	// Sampai sini dijamin os.Args[1] ada, jadi aman diakses
	path := os.Args[1]

	// ReadDir mengembalikan 2 nilai: daftar entry dan error.
	// Kalau gagal (path tidak ada, permission denied), entries kosong dan err terisi.
	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	// range menghasilkan (index, nilai); index tidak kita pakai jadi pakai _
	for _, entry := range entries {
		// IsDir() true kalau entry ini folder
		kind := "file"
		if entry.IsDir() {
			kind = "dir"
		}
		fmt.Println(kind, entry.Name())

	}

}
