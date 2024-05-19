// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// type Quiz struct {
// 	nomorSoal int
// 	soal      string
// 	opsi      [4]string
// 	opsiBenar string
// }

// func main() {
// 	var quizzes [100]Quiz
// 	var banyakSoalQuiz int

// 	scanner := bufio.NewScanner(os.Stdin)

// 	for i := 0; i < 100; i++ {
// 		var q Quiz
// 		q.nomorSoal = i + 1

// 		fmt.Println("Masukkan Soal (ketik 'SELESAI' untuk selesai): ")
// 		if !scanner.Scan() {
// 			break
// 		}
// 		q.soal = scanner.Text()
// 		if q.soal == "SELESAI" {
// 			break
// 		}

// 		for j := 0; j < 4; j++ {
// 			fmt.Printf("Masukkan opsi %c: ", 'A'+j)
// 			if !scanner.Scan() {
// 				break
// 			}
// 			q.opsi[j] = scanner.Text()
// 		}

// 		for {
// 			fmt.Print("Pilih opsi yang benar (A/B/C/D): ")
// 			if !scanner.Scan() {
// 				break
// 			}
// 			q.opsiBenar = scanner.Text()
// 			q.opsiBenar = q.opsiBenar[:1] // Take only the first character
// 			if q.opsiBenar == "A" || q.opsiBenar == "B" || q.opsiBenar == "C" || q.opsiBenar == "D" {
// 				break
// 			}
// 			fmt.Println("Opsi harus A/B/C/D")
// 		}

// 		quizzes[i] = q
// 		banyakSoalQuiz++
// 	}

// 	if err := scanner.Err(); err != nil {
// 		fmt.Println("Error reading input:", err)
// 		return
// 	}

// 	// Print the quizzes
// 	for i := 0; i < banyakSoalQuiz; i++ {
// 		fmt.Printf("Nomor Soal: %d\n", quizzes[i].nomorSoal)
// 		fmt.Println("Soal:", quizzes[i].soal)
// 		fmt.Println("Opsi:")
// 		for j, opsi := range quizzes[i].opsi {
// 			fmt.Printf("%c. %s\n", 'A'+j, opsi)
// 		}
// 		fmt.Println("Jawaban yang benar:", quizzes[i].opsiBenar)
// 	}
// }
