package main
import "fmt"

const NMax = 1000 // Membatasi maksimal 1000 pendaftar

type Peserta struct {
	ID            int
	Nama          string
	TanggalDaftar string
	BidangMinat   string
}
type TabPeserta struct {
	Data  [NMax]Peserta
	Count int
}
func kecilBesar(s string) string {		// Mengubah teks menjadi huruf kecil semua menggunakan ASCII untuk kemudahan komparasi	
	var i int
	var result string
	for i = 0; i < len(s); i = i + 1 {
		if s[i] >= 'A' && s[i] <= 'Z' {
			result += string(s[i] + 32)
		} else {
			result += string(s[i])
		}
	}
	return result
}
func kecilBesarID(t TabPeserta, id int) int {	// Mencari indeks posisi ID berada
	var i int
	for i = 0; i < t.Count; i = i + 1 {
		if t.Data[i].ID == id {
			return i
		}
	}
	return -1
}
func tambahPeserta(t *TabPeserta) {		// Menambahkan data Peserta
	if t.Count >= NMax {
		fmt.Println("Kapasitas pendaftar penuh.")
	} else {
		fmt.Println("\n-- Tambah Peserta Baru --")
		fmt.Print("ID Pendaftaran: ")
		fmt.Scan(&t.Data[t.Count].ID)
		fmt.Print("Nama Lengkap (tanpa spasi/gunakan_underscore): ")
		fmt.Scan(&t.Data[t.Count].Nama)
		fmt.Print("Tanggal Daftar (DD-MM-YYYY): ")
		fmt.Scan(&t.Data[t.Count].TanggalDaftar)
		fmt.Print("Bidang Minat: ")
		fmt.Scan(&t.Data[t.Count].BidangMinat)

		t.Count = t.Count + 1
		fmt.Println("Data peserta berhasil ditambahkan!")
	}
}

func ubahPeserta(t *TabPeserta) {		// Mengubah data Peserta
	var id, idx int
	fmt.Print("\nMasukkan ID Peserta yang ingin diubah: ")
	fmt.Scan(&id)

	idx = kecilBesarID(*t, id)
	if idx != -1 {
		fmt.Print("Nama Baru (tanpa spasi): ")
		fmt.Scan(&t.Data[idx].Nama)
		fmt.Print("Tanggal Daftar Baru (DD-MM-YYYY): ")
		fmt.Scan(&t.Data[idx].TanggalDaftar)
		fmt.Print("Bidang Minat Baru: ")
		fmt.Scan(&t.Data[idx].BidangMinat)
		fmt.Println("Data berhasil diubah.")
	} else {
		fmt.Println("Peserta dengan ID tersebut tidak ditemukan.")
	}
}

func hapusPeserta(t *TabPeserta) {	// Menghapus data peserta
	var id, idx, i int
	fmt.Print("\nMasukkan ID Peserta yang ingin dihapus: ")
	fmt.Scan(&id)

	idx = kecilBesarID(*t, id)
	if idx != -1 {
		for i = idx; i < t.Count-1; i = i + 1 {
			t.Data[i] = t.Data[i+1]
		}
		t.Count = t.Count - 1
		fmt.Println("Data peserta berhasil dihapus.")
	} else {
		fmt.Println("Peserta dengan ID tersebut tidak ditemukan.")
	}
}
func cariBidangSequential(t TabPeserta, minat string) {	
	var ditemukan bool = false
	target := kecilBesar(minat)

	fmt.Printf("\n[Sequential Search] Hasil pencarian untuk bidang minat '%s':\n", target)
	for i := 0; i < t.Count; i = i + 1 {
		if kecilBesar(t.Data[i].BidangMinat) == target {
			// Lebar ID diubah jadi %-15d dan Minat ditambahkan ke output
			fmt.Printf("- ID: %-15d | Nama: %-20s | Minat: %-15s | Tanggal: %-10s\n", t.Data[i].ID, t.Data[i].Nama, kecilBesar(t.Data[i].BidangMinat), t.Data[i].TanggalDaftar)
			ditemukan = true
		}
	}
	if !ditemukan {
		fmt.Println("Tidak ada peserta di bidang minat tersebut.")
	}
}
func sortBidangMinat(t *TabPeserta) {		// Mengurutkan Bidang Minat sebelum Binary Search menggunakan sorting
	var i int
	for i = 1; i < t.Count; i = i + 1 {
		temp := t.Data[i]
		j := i - 1
		for j >= 0 && kecilBesar(t.Data[j].BidangMinat) > kecilBesar(temp.BidangMinat) {
			t.Data[j+1] = t.Data[j]
			j = j - 1
		}
		t.Data[j+1] = temp
	}
}

func cariBidangBinary(t *TabPeserta, minat string) {
	sortBidangMinat(t)
	target := kecilBesar(minat)
	left := 0
	right := t.Count - 1 
	idxMatch := -1
	ketemu := false 

	for left <= right && !ketemu {
		mid := (left + right) / 2
		midVal := kecilBesar(t.Data[mid].BidangMinat)
		
		if midVal == target {
			idxMatch = mid
			ketemu = true 
		} else if midVal < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if idxMatch != -1 {
		fmt.Printf("\n[Binary Search] Hasil pencarian untuk bidang minat '%s':\n", target)
		start := idxMatch 
		for start > 0 && kecilBesar(t.Data[start-1].BidangMinat) == target {
			start = start - 1
		}
		end := idxMatch 
		for end < t.Count-1 && kecilBesar(t.Data[end+1].BidangMinat) == target {
			end = end + 1
		}
		for i := start; i <= end; i = i + 1 {
			// Lebar ID diubah jadi %-15d dan Minat ditambahkan ke output
			fmt.Printf("- ID: %-15d | Nama: %-20s | Minat: %-15s | Tanggal: %-10s\n", t.Data[i].ID, t.Data[i].Nama, kecilBesar(t.Data[i].BidangMinat), t.Data[i].TanggalDaftar)
		}
	} else {
		fmt.Println("Tidak ada peserta di bidang minat tersebut.")
	}
}
func selectionSortID(t *TabPeserta, isAscending bool) {		// Mengurutkan ID peserta (Selection Sort)
	var i, j int
	for i = 0; i < t.Count-1; i = i + 1 {
		idx := i
		for j = i + 1; j < t.Count; j = j + 1 {
			if isAscending {
				if t.Data[j].ID < t.Data[idx].ID {
					idx = j
				}
			} else {
				if t.Data[j].ID > t.Data[idx].ID {
					idx = j
				}
			}
		}
		temp := t.Data[i]
		t.Data[i] = t.Data[idx]
		t.Data[idx] = temp
	}
}
func insertionSortID(t *TabPeserta, isAscending bool) {		// Mengurutkan ID peserta (Insertion Sort)
	var i int
	for i = 1; i < t.Count; i = i + 1 {
		temp := t.Data[i]
		j := i - 1
		if isAscending {
			for j >= 0 && t.Data[j].ID > temp.ID {
				t.Data[j+1] = t.Data[j]
				j = j - 1
			}
		} else {
			for j >= 0 && t.Data[j].ID < temp.ID {
				t.Data[j+1] = t.Data[j]
				j = j - 1
			}
		}
		t.Data[j+1] = temp
	}
}

func tampilkanStatistik(t TabPeserta) {
	fmt.Println("\n--- Statistik KursusIn ---")
	fmt.Printf("Total Peserta Aktif: %d orang\n", t.Count)

	var listMinat [NMax]string
	var listJumlah [NMax]int
	var i int
	var ada bool
	rekapCount := 0

	for i = 0; i < t.Count; i = i + 1 {
		minatSaatIni := kecilBesar(t.Data[i].BidangMinat)
		ada = false

		for j := 0; j < rekapCount && !ada; j = j + 1 {
			if listMinat[j] == minatSaatIni {
				listJumlah[j] = listJumlah[j] + 1
				ada = true
			}
		}
		if !ada {
			listMinat[rekapCount] = minatSaatIni
			listJumlah[rekapCount] = 1
			rekapCount = rekapCount + 1
		}
	}
	fmt.Println("Jumlah Pendaftar per Bidang Minat:")
	for i := 0; i < rekapCount; i = i + 1 {
		fmt.Printf("- %-15s: %d orang\n", listMinat[i], listJumlah[i])
	}
}

func tampilkanSemua(t TabPeserta) {
	if t.Count == 0 {
		fmt.Println("Belum ada data peserta yang terdaftar.")
	} else {
		// Garis diperpanjang dan ID diberi jatah 15 spasi
		fmt.Printf("%-15s | %-20s | %-18s | %-15s\n", "ID", "Nama", "Bidang Minat", "Tanggal Daftar")
		fmt.Println("-------------------------------------------------------------------------------")
		for i := 0; i < t.Count; i = i + 1 {
			// Memanggil kecilBesar() pada BidangMinat agar tampilannya seragam saat dicetak
			fmt.Printf("%-15d | %-20s | %-18s | %-15s\n", t.Data[i].ID, t.Data[i].Nama, kecilBesar(t.Data[i].BidangMinat), t.Data[i].TanggalDaftar)
		}
		fmt.Println("-------------------------------------------------------------------------------")
	}
}

func main() {
	var tabel TabPeserta
	var pilihan int
	var running bool = true

	for running {
		fmt.Println("\n=== SISTEM PENDAFTARAN KURSUS ONLINE (KURSUSIN) ===")
		fmt.Println("1. Tambah Peserta Baru")
		fmt.Println("2. Ubah Data Peserta")
		fmt.Println("3. Hapus Data Peserta")
		fmt.Println("4. Cari Peserta Berdasarkan Bidang Minat")
		fmt.Println("5. Urutkan Peserta")
		fmt.Println("6. Tampilkan Statistik")
		fmt.Println("7. Tampilkan Semua Peserta")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tambahPeserta(&tabel)
		case 2:
			ubahPeserta(&tabel)
		case 3:
			hapusPeserta(&tabel)
		case 4:
			if tabel.Count == 0 {
				fmt.Println("Data masih kosong!")
			} else {
				var subMenu, minat string
				fmt.Println("\n-- Metode Pencarian Bidang Minat --")
				fmt.Println("1. Sequential Search")
				fmt.Println("2. Binary Search")
				fmt.Print("Pilih metode (1/2): ")
				fmt.Scan(&subMenu)
				
				fmt.Print("Masukkan Bidang Minat yang dicari: ")
				fmt.Scan(&minat)

				if subMenu == "1" {
					cariBidangSequential(tabel, minat)
				} else if subMenu == "2" {
					cariBidangBinary(&tabel, minat)
				} else {
					fmt.Println("Metode tidak valid.")
				}
			}
		case 5:
			if tabel.Count == 0 {
				fmt.Println("Data masih kosong!")
			} else {
				var metodeSort, arahSort string
				fmt.Println("\n-- Metode Sorting ID --")
				fmt.Println("1. Selection Sort")
				fmt.Println("2. Insertion Sort")
				fmt.Print("Pilih metode (1/2): ")
				fmt.Scan(&metodeSort)

				if metodeSort == "1" || metodeSort == "2" {
					fmt.Println("\n-- Arah Pengurutan --")
					fmt.Println("1. Ascending (Kecil ke Besar)")
					fmt.Println("2. Descending (Besar ke Kecil)")
					fmt.Print("Pilih arah (1/2): ")
					fmt.Scan(&arahSort)

					isAscending := (arahSort == "1")

					if metodeSort == "1" {
						selectionSortID(&tabel, isAscending)
						fmt.Println("Berhasil diurutkan dengan Selection Sort.")
					} else {
						insertionSortID(&tabel, isAscending)
						fmt.Println("Berhasil diurutkan dengan Insertion Sort.")
					}
					tampilkanSemua(tabel)
				} else {
					fmt.Println("Pilihan metode sorting tidak valid.")
				}
			}
		case 6:
			tampilkanStatistik(tabel)
		case 7:
			tampilkanSemua(tabel)
		case 0:
			fmt.Println("Keluar dari program. Terima kasih!")
			running = false
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}