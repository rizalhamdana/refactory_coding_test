
# Soal 2 (JSON Manipulation)

## Soal

We have JSON Data: [check here](https://gist.github.com/dhamanutd/6993984928506eea49908c2e3fcbc628)


Your tasks:  
1. Find users who doesn't have any phone numbers. 
2. Find users who have articles. 
3. Find users who have "annis" on their name. 
4. Find users who have articles on year 2020. 
5. Find users who are born on 1986. 
6. Find articles that contain "tips" on the title. 
7. Find articles published before August 2019.


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
2. Buka terminal dan pastikan berada di directory `soal_2`
3. Build program dengan menjalankan perintah berikut
```
go build .
```
4. Jalankan program dengan menjalankan perintah berikut
```
./soal_2
```





