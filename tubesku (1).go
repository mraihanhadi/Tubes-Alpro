package main

import (
	"bufio"
	"fmt"
	"os"
)

var scanner = bufio.NewScanner(os.Stdin)

const nmax = 100

type Study struct {
	namaMateri string
	babMateri  string
	banyakBab  int
}

type quiz struct {
	soal                  string
	opsiBenar             string
	nomorSoal, totalBenar int
	opsi                  [4]string
	banyakSoal            int
}

type Form struct {
	Judul string
	komen string
}

type tabForm [nmax]Form

var tabelForm tabForm

type tabstudy [nmax]Study

var tabelStudy tabstudy

type tabQuiz [nmax]quiz

var tabelQuiz tabQuiz

func main() {
	var input int
	intro()
	for true {
		fmt.Println("=========\t Menu Study App \t==========")
		fmt.Println("0. Keluar")
		fmt.Println("1. Menu Guru")
		fmt.Println("2. Menu Siswa")
		fmt.Println("==================================================")
		fmt.Print("Pilih Menu: ")
		fmt.Scanln(&input)
		if input == 0 || input == 1 || input == 2 {
			switch input {
			case 0:
				out()
				os.Exit(0)
			case 1:
				menuGuru()
			case 2:
				menuSiswa()
			}
		}
	}
}

func intro() {
	fmt.Println("             Welcome To Our Study App      ")
	fmt.Println()
}

func out() {
	fmt.Println("Thank Youu")

}

func menuSiswa() {
	var input int
	for true {
		fmt.Println("=========\t Pilihan Menu Siswa \t=========")
		fmt.Println("0. Keluar")
		fmt.Println("1. Materi Pembelajaran")
		fmt.Println("2. Tugas")
		fmt.Println("3. Quiz")
		fmt.Println("4. Forum Diskusi")
		fmt.Println("==================================================")
		fmt.Print("Pilih Menu: ")
		fmt.Scan(&input)
		if input == 0 || input == 1 || input == 2 || input == 3 || input == 4 {
			switch input {
			case 0:
				return
			case 1:
				materiMurid()
			case 2:
				tugasMurid()
			case 3:
				quizMurid()
			case 4:

			}
		}
	}
}

func menuGuru() {
	var input int
	for true {
		fmt.Println("=========\t Pilihan Menu Guru \t=========")
		fmt.Println("0. Keluar")
		fmt.Println("1. Materi Pembelajaran")
		fmt.Println("2. Tugas")
		fmt.Println("3. Quiz")
		fmt.Println("4. Forum Diskusi")
		fmt.Println("======================================")
		fmt.Print("Pilih Menu: ")
		fmt.Scanln(&input)
		if input == 0 || input == 1 || input == 2 || input == 3 || input == 4 {
			switch input {
			case 0:
				return
			case 1:
				materiGuru()
			case 2:
				tugasGuru()
			case 3:
				quizGuru()
			case 4:
				forumGuru()
			}
		}
	}
}

func inMateri() {
	var m Study
	fmt.Println("Masukkan NONE sebagai judul untuk selesai")
	fmt.Println("Masukkan Judul Materi: ")
	for i := 0; i < 100; i++ {
		if !scanner.Scan() {
			break
		}
		m.namaMateri = scanner.Text()
		if m.namaMateri == "NONE" {
			break
		}
		fmt.Println("Masukkan Bab Materi: ")
		fmt.Scan(&m.babMateri)
		tabelStudy[i] = m
		fmt.Println("Masukkan Judul Materi: ")
		fmt.Scan(&m.namaMateri)
	}
}

func inQuiz() {
	var q quiz
	var banyakSoalQuiz int

	fmt.Println("Ketik NONE sebagai soal untuk selesai")
	for i := 0; i < 100; i++ {
		q.nomorSoal = i + 1
		fmt.Println("Masukkan Soal: ")
		if !scanner.Scan() {
			break
		}
		q.soal = scanner.Text()
		if q.soal == "NONE" {
			break
		}
		for j := 0; j < 4; j++ {
			fmt.Printf("Masukkan opsi %c: ", 'A'+j)
			if !scanner.Scan() {
				break
			}
			q.opsi[j] = scanner.Text()
		}
		for {
			fmt.Print("Pilih opsi yang benar (A/B/C/D): ")
			if !scanner.Scan() {
				break
			}
			q.opsiBenar = scanner.Text()
			if q.opsiBenar == "A" || q.opsiBenar == "B" || q.opsiBenar == "C" || q.opsiBenar == "D" {
				break
			} else {
				fmt.Println("Opsi harus A/B/C/D")
			}
		}

		tabelQuiz[i] = q
		banyakSoalQuiz++
	}
}

func inForumGuru() {
	var f Form

	fmt.Println("Ketik NONE sebagai judul/komen untuk selesai")
	fmt.Print("Masukkan Judul Form Diskusi: ")
	if scanner.Scan() {
		f.Judul = scanner.Text()
	}
	for i := 0; i < 100; i++ {
		fmt.Println("Masukkan komen:")
		if !scanner.Scan() {
			break
		}
		f.komen = scanner.Text()
		if f.komen == "NONE" {
			break
		}
		tabelForm[i] = f
	}
}

func cariMateri() {
	var cariMateri string
	fmt.Println("Ketik NONE untuk selesai")
	fmt.Print("Masukkan Judul Materi: ")
	fmt.Scan(&cariMateri)
	for cariMateri != " NONE" {

	}
}

func cariQuiz() {
	var kerjakan string
	fmt.Println("Ketik NONE untuk selesai")
	fmt.Print("Masukkan Judul Quiz yang ingin dikerjakan:")
	fmt.Scan(&kerjakan)
}

func materiGuru() {
	var input int
	for true {
		fmt.Println("=========\t Pilihan Menu  \t=========")
		fmt.Println("0. Keluar")
		fmt.Println("1. Masukkan Materi")
		fmt.Println("2. Edit Materi")
		fmt.Println("3. Hapus Materi")
		fmt.Println("======================================")
		fmt.Print("Pilih Menu: ")
		fmt.Scanln(&input)
		if input == 0 || input == 1 || input == 2 || input == 3 || input == 4 {
			switch input {
			case 0:
				return
			case 1:
				inMateri()
			case 2:

			case 3:

			case 4:

			}
		}
	}
}

func quizGuru() {
	var input int
	for true {
		fmt.Println("=========\t Pilihan Menu  \t=========")
		fmt.Println("0. Keluar")
		fmt.Println("1. Masukkan Quiz")
		fmt.Println("2. Edit Quiz")
		fmt.Println("3. Hapus Quiz")
		fmt.Println("======================================")
		fmt.Print("Pilih Menu: ")
		fmt.Scanln(&input)
		if input == 0 || input == 1 || input == 2 || input == 3 {
			switch input {
			case 0:
				return
			case 1:
				inQuiz()
			case 2:

			case 3:

			case 4:

			}
		}
	}
}

func forumGuru() {
	var input int
	for true {
		fmt.Println("=========\t Pilihan Menu Forum  \t=========")
		fmt.Println("0. Keluar")
		fmt.Println("1. Masukkan Forum")
		fmt.Println("2. Edit Forum")
		fmt.Println("3. Hapus Forum")
		fmt.Println("======================================")
		fmt.Print("Pilih Menu: ")
		fmt.Scanln(&input)
		if input == 0 || input == 1 || input == 2 || input == 3 {
			switch input {
			case 0:
				return
			case 1:
				inForumGuru()
			case 2:

			case 3:

			case 4:

			}
		}
	}
}

func tugasGuru() {
	var input int
	for true {
		fmt.Println("=========\t Pilihan Menu Tugas  \t=========")
		fmt.Println("0. Keluar")
		fmt.Println("1. Masukkan Tugas")
		fmt.Println("2. Edit Tugas")
		fmt.Println("3. Hapus Tugas")
		fmt.Println("======================================")
		fmt.Print("Pilih Menu: ")
		fmt.Scan(&input)
		if input == 0 || input == 1 || input == 2 || input == 3 {
			switch input {
			case 0:
				return
			case 1:

			case 2:

			case 3:

			case 4:

			}
		}
	}
}

func quizMurid() {
	var input int
	for true {
		fmt.Println("=========\t Pilihan Menu Quiz  \t=========")
		fmt.Println("0. Keluar")
		fmt.Println("1. Pilih Materi Quiz")
		fmt.Println("======================================")
		fmt.Print("Pilih Menu: ")
		fmt.Scan(&input)
		if input == 0 || input == 1 || input == 2 || input == 3 {
			switch input {
			case 0:
				return
			case 1:

			case 2:

			case 3:

			case 4:

			}
		}
	}
}

func tugasMurid() {
	var input int
	for true {
		fmt.Println("=========\t Pilihan Menu Tugas  \t=========")
		fmt.Println("0. Keluar")
		fmt.Println("1. Masukkan Materi Tugas")
		fmt.Println("======================================")
		fmt.Print("Pilih Menu: ")
		fmt.Scan(&input)
		if input == 0 || input == 1 {
			switch input {
			case 0:
				return
			case 1:

			case 2:

			case 3:

			case 4:

			}
		}
	}
}

func materiMurid() {
	var input int
	for true {
		fmt.Println("=========\t Pilihan Menu  \t=========")
		fmt.Println("0. Keluar")
		fmt.Println("1. ")
		fmt.Println("======================================")
		fmt.Print("Pilih Menu: ")
		fmt.Scan(&input)
		if input == 0 || input == 1 {
			switch input {
			case 0:
				return
			case 1:

			case 2:

			case 3:

			case 4:

			}
		}
	}
}