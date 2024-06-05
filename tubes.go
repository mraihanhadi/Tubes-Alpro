package main

import (
	"bufio"
	"fmt"
	"os"
)

const nmax = 100

type Study struct {
	isiMateri  string
	babMateri  string
	namaMatkul string
}

type quiz struct {
	soal                  string
	opsiBenar             string
	nomorSoal, totalBenar int
	opsi                  [4]string
	materiQuiz            string
}

type Form struct {
	Judul string
	komen string
}

type tugas struct {
	judulTugas  string
	soal        string
	materiTugas string
}

type tabForm [nmax]Form

var tabelForm tabForm

type tabstudy [nmax]Study

var tabelStudy tabstudy

type tabQuiz [nmax]quiz

var tabelQuiz tabQuiz

type tabTugas [nmax]tugas

var tabelTugas tabTugas

type Student struct {
	nama         string
	nim          string
	jawabanTugas [nmax]string
	nilaiTugas   [nmax]int
	quizJawaban  [nmax]string
	nilaiQuiz    int
}

type tabStudent [nmax]Student

var tabelStudent tabStudent

var currentStudent Student

func main() {
	var countMateri, countQuiz, countTugas, countForm, countStudent int
	var input int
	var scanner = bufio.NewScanner(os.Stdin)
	intro()
	for {
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
				menuGuru(scanner, &countMateri, &countQuiz, &countTugas, &countForm, &countStudent)
			case 2:
				menuSiswa(scanner, &countMateri, &countQuiz, &countTugas, &countForm, &countStudent)
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

func inputStudentData(scanner *bufio.Scanner, countStudent *int) {
	fmt.Println("Masukkan Nama dan NIM Anda:")
	fmt.Print("Nama: ")
	if scanner.Scan() {
		currentStudent.nama = scanner.Text()
	}
	fmt.Print("NIM: ")
	if scanner.Scan() {
		currentStudent.nim = scanner.Text()
	}
	tabelStudent[*countStudent].nama = currentStudent.nama
	tabelStudent[*countStudent].nim = currentStudent.nim
	*countStudent++
	fmt.Printf("Selamat datang, %s (NIM: %s)\n", currentStudent.nama, currentStudent.nim)
}

func menuSiswa(scanner *bufio.Scanner, countMateri, countQuiz, countTugas, countForm, countStudent *int) {
	inputStudentData(scanner, &*countStudent)
	var input int
	for {
		fmt.Println("=========\t Pilihan Menu Siswa \t=========")
		fmt.Println("0. Keluar")
		fmt.Println("1. Materi Pembelajaran")
		fmt.Println("2. Tugas")
		fmt.Println("3. Quiz")
		fmt.Println("4. Forum Diskusi")
		fmt.Println("==================================================")
		fmt.Print("Pilih Menu: ")
		fmt.Scanln(&input)
		if input == 0 || input == 1 || input == 2 || input == 3 || input == 4 {
			switch input {
			case 0:
				return
			case 1:
				materiMurid(scanner, *countMateri)
			case 2:
				tugasMurid(scanner, *countTugas, *countStudent)
			case 3:
				quizMurid(scanner, *countQuiz, *countStudent)
			case 4:
				forumMurid(scanner, &*countForm)
			}
		}
	}
}

func menuGuru(scanner *bufio.Scanner, countMateri, countQuiz, countTugas, countForm, countStudent *int) {
	var input int
	for {
		fmt.Println("=========\t Pilihan Menu Guru \t=========")
		fmt.Println("0. Keluar")
		fmt.Println("1. Materi Pembelajaran")
		fmt.Println("2. Tugas")
		fmt.Println("3. Quiz")
		fmt.Println("4. Forum Diskusi")
		fmt.Println("5. Data Siswa dan Nilai")
		fmt.Println("6. Jumlah Siswa per Materi/Tugas")
		fmt.Println("======================================")
		fmt.Print("Pilih Menu: ")
		fmt.Scanln(&input)
		if input == 0 || input == 1 || input == 2 || input == 3 || input == 4 || input == 5 || input == 6 {
			switch input {
			case 0:
				return
			case 1:
				materiGuru(scanner, &*countMateri)
			case 2:
				tugasGuru(scanner, &*countTugas, *countStudent)
			case 3:
				quizGuru(scanner, countQuiz)
			case 4:
				forumGuru(scanner, &*countForm)
			case 5:
				lihatDataSiswa(*countStudent)
			case 6:
				tampilJumlahSiswaPerMateriTugas(*countMateri, *countTugas, *countStudent)
			}
		}
	}
}

func inMateri(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0

	for scanner.Scan() {
		if count >= nmax {
			fmt.Println("Jumlah maksimum materi telah tercapai.")
			return nmax
		}
		materi := Study{}
		materi.namaMatkul = scanner.Text()

		if !scanner.Scan() {
			return count
		}
		materi.babMateri = scanner.Text()

		if !scanner.Scan() {
			return count
		}
		materi.isiMateri = scanner.Text()

		tabelStudy[count] = materi
		count++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Gagal membaca file:", err)
	}

	return count
}

func searchStudy(namaMatkul, babMateri string, n int) int {
	for i := 0; i < n; i++ {
		if tabelStudy[i].namaMatkul == namaMatkul && tabelStudy[i].babMateri == babMateri {
			return i
		}
	}
	return -1 // Tidak ditemukan
}

func editMateri(index int, materi Study) {
	if index >= 0 && index < nmax {
		tabelStudy[index].babMateri = materi.babMateri
		tabelStudy[index].isiMateri = materi.isiMateri
		fmt.Println("Materi berhasil diubah.")
	} else {
		fmt.Println("Indeks tidak valid.")
	}
}

func hapusMateri(index *int) {
	if *index >= 0 && *index < nmax {
		for i := *index; i < nmax-1; i++ {
			tabelStudy[i] = tabelStudy[i+1]
		}
		tabelStudy[nmax-1] = Study{}
		fmt.Println("Materi berhasil dihapus.")
		(*index)--
	} else {
		fmt.Println("Indeks tidak valid.")
	}
}

func materiGuru(scanner *bufio.Scanner, countMateri *int) {
	var input int
	for {
		fmt.Println("=========\t  Pilihan Menu  \t=========")
		fmt.Println("0. Keluar")
		fmt.Println("1. Masukkan Materi")
		fmt.Println("2. Edit Materi")
		fmt.Println("3. Hapus Materi")
		fmt.Println("======================================")
		fmt.Print("Pilih Menu: ")
		fmt.Scanln(&input)
		if input == 0 || input == 1 || input == 2 || input == 3 {
			switch input {
			case 0:
				return
			case 1:
				fmt.Println("Masukkan nama file materi: ")
				var filename string
				fmt.Scanln(&filename)
				*countMateri = inMateri(filename)
				fmt.Printf("Total Materi yang berhasil dibaca: %d\n", *countMateri)
			case 2:
				var namaMatkul, babMateri string
				var materi Study
				fmt.Print("Masukkan nama matkul yang ingin diedit: ")
				if scanner.Scan() {
					namaMatkul = scanner.Text()
				}
				fmt.Print("Masukkan bab materi yang ingin diedit: ")
				if scanner.Scan() {
					babMateri = scanner.Text()
				}
				index := searchStudy(namaMatkul, babMateri, *countMateri)
				if index != -1 {
					fmt.Print("Masukkan bab materi baru: ")
					if scanner.Scan() {
						materi.babMateri = scanner.Text()
					}
					fmt.Print("Masukkan isi materi baru: ")
					if scanner.Scan() {
						materi.isiMateri = scanner.Text()
					}
					editMateri(index, materi)
				} else {
					fmt.Println("Materi tidak ditemukan.")
				}
			case 3:
				var namaMatkul, babMateri string
				fmt.Print("Masukkan nama matkul yang ingin dihapus: ")
				if scanner.Scan() {
					namaMatkul = scanner.Text()
				}
				fmt.Print("Masukkan bab materi yang ingin dihapus: ")
				if scanner.Scan() {
					babMateri = scanner.Text()
				}
				index := searchStudy(namaMatkul, babMateri, *countMateri)
				if index != -1 {
					hapusMateri(&index)
				} else {
					fmt.Println("Materi tidak ditemukan.")
				}
			}
		}
	}
}

func inTugas(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Gagal membuka file:", err)
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0

	for scanner.Scan() {
		if count >= nmax {
			fmt.Println("Jumlah maksimum tugas telah tercapai.")
			break
		}
		t := tugas{}
		t.materiTugas = scanner.Text()

		if !scanner.Scan() {
			break
		}
		t.judulTugas = scanner.Text()

		if !scanner.Scan() {
			break
		}
		t.soal = scanner.Text()

		tabelTugas[count] = t
		count++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Gagal membaca file:", err)
	}

	return count
}

func searchTugas(materiTugas, judulTugas string, countTugas int) int {
	for i := 0; i < countTugas; i++ {
		if tabelTugas[i].materiTugas == materiTugas && tabelTugas[i].judulTugas == judulTugas {
			return i
		}
	}
	return -1
}

func editTugas(index int, t tugas) {
	if index >= 0 && index < nmax {
		tabelTugas[index] = t
		fmt.Println("Tugas berhasil diedit.")
	} else {
		fmt.Println("Indeks tidak valid.")
	}
}

func hapusTugas(index *int) {
	if *index >= 0 && *index < nmax {
		for i := *index; i < nmax-1; i++ {
			tabelTugas[i] = tabelTugas[i+1]
		}
		// Mengosongkan elemen terakhir
		tabelTugas[nmax-1] = tugas{}
		fmt.Println("Tugas berhasil dihapus.")
		(*index)--
	} else {
		fmt.Println("Indeks tidak valid.")
	}
}

func tugasGuru(scanner *bufio.Scanner, countTugas *int, countStudent int) {
	var input int
	fmt.Println("=========\t Pilihan Menu Tugas  \t=========")
	fmt.Println("0. Keluar")
	fmt.Println("1. Masukkan Tugas")
	fmt.Println("2. Edit Tugas")
	fmt.Println("3. Hapus Tugas")
	fmt.Println("4. Nilai tugas siswa")
	fmt.Println("======================================")
	fmt.Print("Pilih Menu: ")
	fmt.Scanln(&input)
	if input == 0 || input == 1 || input == 2 || input == 3 || input == 4 {
		switch input {
		case 0:
			return
		case 1:
			fmt.Println("Masukkan nama file tugas: ")
			var filename string
			fmt.Scanln(&filename)
			*countTugas = inTugas(filename)
			fmt.Printf("Total Tugas yang berhasil dibaca: %d\n", *countTugas)
		case 2:
			var materiTugas, judulTugas string
			var t tugas
			fmt.Print("Masukkan materi tugas yang ingin diedit: ")
			if scanner.Scan() {
				materiTugas = scanner.Text()
			}
			fmt.Print("Masukkan judul tugas yang ingin diedit: ")
			if scanner.Scan() {
				judulTugas = scanner.Text()
			}
			index := searchTugas(materiTugas, judulTugas, *countTugas)
			if index != -1 {
				fmt.Print("Masukkan judul tugas yang baru: ")
				if scanner.Scan() {
					t.judulTugas = scanner.Text()
				}
				fmt.Print("Masukkan soal tugas yang baru: ")
				if scanner.Scan() {
					t.judulTugas = scanner.Text()
				}
				editTugas(index, t)
			} else {
				fmt.Println("Tugas tidak ditemukan.")
			}
		case 3:
			var materiTugas, judulTugas string
			fmt.Print("Masukkan materi tugas yang ingin dihapus: ")
			if scanner.Scan() {
				materiTugas = scanner.Text()
			}
			fmt.Print("Masukkan judul tugas yang ingin dihapus: ")
			if scanner.Scan() {
				judulTugas = scanner.Text()
			}
			index := searchStudy(materiTugas, judulTugas, *countTugas)
			if index != -1 {
				hapusMateri(&index)
			} else {
				fmt.Println("Tugas tidak ditemukan.")
			}
		case 4:
			lihatJawabanTugas(scanner, countStudent)
		}

	}
}

func inQuiz(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0

	for {
		if count >= nmax {
			fmt.Println("Jumlah maksimum quiz telah tercapai.")
			return nmax
		}
		q := quiz{}

		if !scanner.Scan() {
			break
		}
		q.materiQuiz = scanner.Text()

		if !scanner.Scan() {
			break
		}
		q.nomorSoal = count + 1

		if !scanner.Scan() {
			break
		}
		q.soal = scanner.Text()

		for i := 0; i < 4; i++ {
			if !scanner.Scan() {
				return count
			}
			q.opsi[i] = scanner.Text()
		}

		if !scanner.Scan() {
			return count
		}
		q.opsiBenar = scanner.Text()

		tabelQuiz[count] = q
		count++
	}
	return count
}

func searchQuiz(materiQuiz string, nomorSoal int, n int) int {
	for i := 0; i < n; i++ {
		if tabelQuiz[i].materiQuiz == materiQuiz && tabelQuiz[i].nomorSoal == nomorSoal {
			return i
		}
	}
	return -1
}

func editQuiz(index int, q quiz) {
	if index >= 0 && index < nmax {
		tabelQuiz[index].soal = q.soal
		for i := 0; i < 4; i++ {
			tabelQuiz[index].opsi[i] = q.opsi[i]
		}
		tabelQuiz[index].opsiBenar = q.opsiBenar
		fmt.Println("Quiz berhasil diedit.")
	} else {
		fmt.Println("Indeks tidak valid.")
	}
}

func hapusQuiz(index *int, countQuiz *int) {
	for i := *index; i < *countQuiz-1; i++ {
		tabelQuiz[i] = tabelQuiz[i+1]
	}

	tabelQuiz[*countQuiz-1] = quiz{}
	(*countQuiz)--

	for i := 0; i < *countQuiz; i++ {
		tabelQuiz[i].nomorSoal = i + 1
	}

	fmt.Println("Quiz berhasil dihapus.")
}

func quizGuru(scanner *bufio.Scanner, countQuiz *int) {
	var input int
	for {
		fmt.Println("=========\t Pilih Menu \t=========")
		fmt.Println("0. Keluar")
		fmt.Println("1. Masukkan Quiz")
		fmt.Println("2. Edit Quiz")
		fmt.Println("3. Hapus Quiz")
		fmt.Println("4. Tampilkan Nilai Quiz")
		fmt.Println("======================================")
		fmt.Print("Pilih Menu: ")
		fmt.Scanln(&input)
		if input == 0 || input == 1 || input == 2 || input == 3 || input == 4 {
			switch input {
			case 0:
				return
			case 1:
				fmt.Println("Masukkan nama file quiz: ")
				var filename string
				fmt.Scanln(&filename)
				*countQuiz = inQuiz(filename)
				fmt.Printf("Total Quiz yang berhasil dibaca: %d\n", *countQuiz)
			case 2:
				var materiQuiz string
				var nomorSoal int
				var q quiz
				fmt.Print("Masukkan materi quiz yang ingin diedit: ")
				if scanner.Scan() {
					materiQuiz = scanner.Text()
				}
				fmt.Print("Masukkan nomor soal yang ingin diedit: ")
				fmt.Scanln(&nomorSoal)
				index := searchQuiz(materiQuiz, nomorSoal, *countQuiz)
				if index != -1 {
					fmt.Print("Masukkan soal quiz baru: ")
					if scanner.Scan() {
						q.soal = scanner.Text()
					}
					fmt.Print("Masukkan opsi baru: ")
					for i := 0; i < 4; i++ {
						if !scanner.Scan() {
							return
						}
						q.opsi[i] = scanner.Text()
					}
					fmt.Print("Masukkan opsi yang benar: ")
					if scanner.Scan() {
						q.opsiBenar = scanner.Text()
					}
					editQuiz(index, q)
				} else {
					fmt.Println("Quiz tidak ditemukan.")
				}
			case 3:
				var materiQuiz string
				var nomorSoal int
				fmt.Print("Masukkan materi quiz yang ingin dihapus: ")
				if scanner.Scan() {
					materiQuiz = scanner.Text()
				}
				fmt.Print("Masukkan nomor soal yang ingin dihapus: ")
				fmt.Scanln(&nomorSoal)
				index := searchQuiz(materiQuiz, nomorSoal, *countQuiz)
				if index != -1 {
					hapusQuiz(&index, &*countQuiz)
				} else {
					fmt.Println("Quiz tidak ditemukan.")
				}
			}
		}
	}
}

func inForumGuru(scanner *bufio.Scanner, countForm *int) {
	if *countForm >= nmax {
		fmt.Println("Jumlah maksimum forum telah tercapai.")
		return
	}

	fmt.Println("Masukkan judul forum: ")
	var judul string
	if scanner.Scan() {
		judul = scanner.Text()
	}

	fmt.Println("Masukkan komentar pembuka: ")
	var komen string
	if scanner.Scan() {
		komen = scanner.Text()
	}

	tabelForm[*countForm] = Form{Judul: judul, komen: komen}
	*countForm++

	fmt.Println("Forum berhasil dibuat.")
}

func searchForm(judul string) int {
	for i := 0; i < nmax; i++ {
		if tabelForm[i].Judul == judul {
			return i
		}
	}
	return -1
}

func editForum(index int, scanner *bufio.Scanner) {
	fmt.Println("Masukkan judul baru (tekan Enter untuk melewati): ")
	if scanner.Scan() {
		newTitle := scanner.Text()
		if newTitle != "" {
			tabelForm[index].Judul = newTitle
		}
	}

	fmt.Println("Masukkan komentar baru (tekan Enter untuk melewati): ")
	if scanner.Scan() {
		newComment := scanner.Text()
		if newComment != "" {
			tabelForm[index].komen = newComment
		}
	}

	fmt.Println("Forum berhasil diupdate.")
}

func hapusForum(index int, countForm *int) {
	for i := index; i < *countForm-1; i++ {
		tabelForm[i] = tabelForm[i+1]
	}

	*countForm--
	fmt.Println("Forum berhasil dihapus.")
}

func forumGuru(scanner *bufio.Scanner, countForm *int) {
	var input int
	for {
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
				inForumGuru(scanner, &*countForm)
			case 2:
				var judul string
				fmt.Print("Masukkan judul forum yang ingin diedit:")
				if scanner.Scan() {
					judul = scanner.Text()
				}
				index := searchForm(judul)
				if index != -1 {
					editForum(index, scanner)
				} else {
					fmt.Println("Forum tidak ditemukan.")
				}
			case 3:
				var judul string
				fmt.Scanln(&judul)
				index := searchForm(judul)
				if index != -1 {
					hapusForum(index, &*countForm)
				}
			}
		}
	}
}

func materiMurid(scanner *bufio.Scanner, countMateri int) {
	for i := 0; i < countMateri; i++ {
		fmt.Printf("Materi %d: %s\n", i+1, tabelStudy[i].namaMatkul)
		fmt.Printf("Bab: %s\n", tabelStudy[i].babMateri)
		fmt.Printf("Isi: %s\n", tabelStudy[i].isiMateri)
	}
	fmt.Println("Apakah Anda ingin mengikuti mata kuliah ini? (y/t)")
	if scanner.Scan() {
		answer := scanner.Text()
		if answer == "y" {
			fmt.Println("Anda telah mengikuti mata kuliah ini.")
		}
	}
}

func tugasMurid(scanner *bufio.Scanner, countTugas int, countStudent int) {
	for i := 0; i < countTugas; i++ {
		fmt.Printf("Tugas %d: %s\n", i+1, tabelTugas[i].judulTugas)
		fmt.Printf("Soal: %s\n", tabelTugas[i].soal)
	}
	fmt.Println("Masukkan nomor tugas yang ingin Anda kerjakan:")
	var nomorTugas int
	fmt.Scanln(&nomorTugas)
	if nomorTugas > 0 && nomorTugas <= countTugas {
		tugas := tabelTugas[nomorTugas-1]
		fmt.Printf("Soal: %s\n", tugas.soal)
		fmt.Println("Masukkan jawaban Anda:")
		if scanner.Scan() {
			jawaban := scanner.Text()
			for i := 0; i < nmax; i++ {
				if currentStudent.jawabanTugas[i] == "" {
					tabelStudent[countStudent-1].jawabanTugas[i] = jawaban
					return
				}
			}
			fmt.Println("Jawaban Anda telah disimpan.")
		}
	} else {
		fmt.Println("Nomor tugas tidak valid.")
	}
}

func quizMurid(scanner *bufio.Scanner, countQuiz int, countStudent int) {
	for i := 0; i < countQuiz; i++ {
		fmt.Printf("Soal %d: %s\n", i+1, tabelQuiz[i].soal)
		for j := 0; j < 4; j++ {
			fmt.Printf("Opsi %d: %s\n", j+1, tabelQuiz[i].opsi[j])
		}
		fmt.Println("Masukkan nomor opsi jawaban Anda:")
		var nomorOpsi int
		fmt.Scanln(&nomorOpsi)
		if nomorOpsi > 0 && nomorOpsi <= 4 {
			currentStudent.quizJawaban[tabelQuiz[i].nomorSoal] = tabelQuiz[i].opsi[nomorOpsi-1]
			if tabelQuiz[i].opsiBenar == tabelQuiz[i].opsi[nomorOpsi-1] {
				currentStudent.nilaiQuiz += 1
			}
		} else {
			fmt.Println("Nomor opsi tidak valid.")
		}
	}
	tabelStudent[countStudent-1].nilaiQuiz = currentStudent.nilaiQuiz * 100 / countQuiz
	fmt.Printf("Nilai Quiz Anda: %d\n", currentStudent.nilaiQuiz*100/countQuiz)
}

func forumMurid(scanner *bufio.Scanner, countForum *int) {
	var input int
	for {
		fmt.Println("=========\t Pilihan Menu Forum  \t=========")
		fmt.Println("0. Keluar")
		fmt.Println("1. Reply Forum")
		fmt.Println("2. Tampilkan Forum")
		fmt.Println("======================================")
		fmt.Print("Pilih Menu: ")
		fmt.Scanln(&input)
		if input == 0 || input == 1 || input == 2 {
			switch input {
			case 0:
				return
			case 1:
				replyForum(scanner)
			case 2:
				tampilkanForum(*countForum)
			}
		}
	}
}

func replyForum(scanner *bufio.Scanner) {
	fmt.Println("Masukkan judul forum yang ingin direply: ")
	var judul string
	if scanner.Scan() {
		judul = scanner.Text()
	}

	index := searchForm(judul)
	if index == -1 {
		fmt.Println("Forum tidak ditemukan.")
		return
	}

	fmt.Println("Masukkan balasan Anda: ")
	var reply string
	if scanner.Scan() {
		reply = scanner.Text()
	}

	tabelForm[index].komen += "\nBalasan: " + reply

	fmt.Println("Balasan berhasil ditambahkan.")
}

func tampilkanForum(countForm int) {
	if countForm == 0 {
		fmt.Println("Tidak ada forum yang tersedia.")
		return
	}

	for i := 0; i < countForm; i++ {
		fmt.Printf("Forum %d:\n", i+1)
		fmt.Printf("Judul: %s\n", tabelForm[i].Judul)
		fmt.Printf("Komentar: %s\n", tabelForm[i].komen)
		fmt.Println("------------------------------------")
	}
}

func tampilDataSiswa(countStudent int) {
	fmt.Println("Data Siswa:")
	for i := 0; i < countStudent; i++ {
		fmt.Printf("Nama: %s, NIM: %s, Nilai Quiz: %d\n", tabelStudent[i].nama, tabelStudent[i].nim, tabelStudent[i].nilaiQuiz)
		for j := 0; j < nmax; j++ {
			if tabelStudent[i].nilaiTugas[j] != 0 {
				fmt.Printf("Nilai Tugas %d: %d\n", j+1, tabelStudent[i].nilaiTugas[j])
			}
		}
	}
}

func lihatJawabanTugas(scanner *bufio.Scanner, countStudent int) {
	for i := 0; i < countStudent; i++ {
		for j := 0; j < nmax; j++ {
			if tabelStudent[i].jawabanTugas[j] != "" {
				fmt.Printf("Nama Siswa: %s, NIM: %s\n", tabelStudent[i].nama, tabelStudent[i].nim)
				fmt.Printf("Jawaban: %s\n", tabelStudent[i].jawabanTugas[j])
				fmt.Println("Masukkan nilai untuk tugas ini:")
				var nilai int
				fmt.Scanln(&nilai)
				tabelStudent[i].nilaiTugas[j] = nilai
				fmt.Println("Nilai telah disimpan.")
			}
		}
	}
}

func lihatDataSiswa(countStudent int) {
	for i := 0; i < countStudent; i++ {
		fmt.Printf("Nama: %s, NIM: %s, Nilai Quiz: %d\n", tabelStudent[i].nama, tabelStudent[i].nim, tabelStudent[i].nilaiQuiz)
		for j := 0; j < nmax; j++ {
			if tabelStudent[i].nilaiTugas[j] != 0 {
				fmt.Printf("Nilai Tugas %d: %d\n", j+1, tabelStudent[i].nilaiTugas[j])
			}
		}
	}
}

func tampilJumlahSiswaPerMateriTugas(countMateri int, countTugas int, countStudent int) {
	fmt.Println("Jumlah Siswa per Materi dan Tugas:")
	for i := 0; i < countMateri; i++ {
		fmt.Printf("Materi: %s\n", tabelStudy[i].namaMatkul)
		countSiswa := 0
		for j := 0; j < countStudent; j++ {
			if tabelStudent[j].nama != "" {
				countSiswa++
			}
		}
		fmt.Printf("Jumlah Siswa: %d\n", countSiswa)
	}
	for i := 0; i < countTugas; i++ {
		fmt.Printf("Tugas: %s\n", tabelTugas[i].judulTugas)
		countSiswa := 0
		for j := 0; j < countStudent; j++ {
			if tabelStudent[j].jawabanTugas[i] != "" {
				countSiswa++
			}
		}
		fmt.Printf("Jumlah Siswa: %d\n", countSiswa)
	}
}
