package main

import (
	"fmt"
	"os"
)

type Study struct {
	namaMateri string
	babMateri  string
}

type tabstudy [100]Study

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
		fmt.Scan(&input)
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
		fmt.Scan(&input)
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

func inMateri(A *tabstudy, n *int) {
	var idx int
	var materi string
	idx = *n
	fmt.Println("Masukkan NONE untuk selesai")
	fmt.Print("Masukkan Judul Materi: ")
	fmt.Scan(&materi)
	for materi != "NONE" {
		A[idx].namaMateri = materi
		idx++
	}

}

func inQuiz() {
	var quiz string
	fmt.Println("Ketik NONE untuk selesai")
	fmt.Print("Masukkan Quiz: ")
	fmt.Scan(&quiz)

}

func inForumGuru() {
	var forum string
	fmt.Println("Ketik NONE untuk selesai")
	fmt.Print("Masukkan Forum Diskusi")
	fmt.Scan(&forum)
	for forum != "NONE" {

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
		fmt.Scan(&input)
		if input == 0 || input == 1 || input == 2 || input == 3 || input == 4 {
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
