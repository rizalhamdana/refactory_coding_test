
# Soal 3 (JSON Manipulation)

## Soal

We have JSON Data: [check here](https://gist.github.com/dhamanutd/97aa0d2131903ea8c071721032c7b2a3)


Your tasks: 

1. Find items in Meeting Room.
2. Find all electronic devices.
3. Find all furnitures.
4. Find all items was purchased at 16 Januari 2020.
5. Find all items with brown color.




## Pembahasan

Source code untuk soal ini ditulis dengan struktur directory sebagai berikut:
1. Directory `handler`: Digunakan untuk memanggil semua fungsi logic yang terkait dengan perintah soal
2. Directory `usecase`: Berisi tentang logika untuk mencari output sesuai dengan perintah soal
3. Directory `repo`: Digunakan untuk mengambil semua data yang ada dari file JSON dan memappingkannya sesuai dengan struktur model
4. Directory `model`: Berisi beberapa `struct` yang merepresentasi kan struktur data dari JSON
5. Directory `resources`: Berisi file JSON yang akan dimanipulasi
6. File `main.go`: File untuk eksekusi program

## How To Run

1. Clone Repository ini
2. Buka terminal dan pastikan berada di directory `soal_3`
3. Build program dengan menjalankan perintah berikut
```
go build .
```
4. Jalankan program dengan menjalankan perintah berikut
```
./soal_3
```





