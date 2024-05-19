// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// type Form struct {
// 	Judul string
// 	Komen string
// }

// type tabForm [100]Form

// var tabelForm tabForm

// func inForumGuru() {
// 	var f Form
// 	scanner := bufio.NewScanner(os.Stdin)

// 	fmt.Println("Ketik NONE untuk selesai")
// 	fmt.Print("Masukkan Judul Form Diskusi: ")
// 	if scanner.Scan() {
// 		f.Judul = scanner.Text()
// 	}
// 	for i := 0; i < 100; i++ {
// 		fmt.Println("Masukkan komen:")
// 		if !scanner.Scan() {
// 			break
// 		}
// 		f.Komen = scanner.Text()
// 		if strings.ToLower(f.Komen) == "none" {
// 			break
// 		}
// 		tabelForm[i] = f
// 	}

// 	fmt.Println("Diskusi berhasil disimpan.")
// 	// Print the stored comments for verification
// 	for i := 0; i < 100; i++ {
// 		if tabelForm[i].Komen != "" {
// 			fmt.Printf("Judul: %s, Komen: %s\n", tabelForm[i].Judul, tabelForm[i].Komen)
// 		} else {
// 			break
// 		}
// 	}
// }

// func main() {
// 	inForumGuru()
// }