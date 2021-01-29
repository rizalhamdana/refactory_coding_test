# Soal 5 Simple APP

## Soal

Buatlah aplikasi client-server dengan kriteria sebagai berikut:
Setiap menit client akan mengirimkan POST request ke server, dengan payload seperti contoh berikut:

```
// Header
"X-RANDOM": "93f9h3dx"
// Body
{ "counter": 1 }
// Header
"X-RANDOM": "fe9g83jm"
// Body
{ "counter": 2 }
// Header
"X-RANDOM": "igrijd9p"
// Body
{ "counter": 3 }
```
Server akan menerima request dari client tersebut. Menyimpan body maupun header ke file server.log. Lalu memberikan response HTTP code 201. Contoh isi file server.log

```
[2020-07-28T16:23:40+07:00] Success: POST http://192.168.1.30/ {"counter": 1, "X-RANDOM": "93f9h3dx"}
[2020-07-28T16:24:40+07:00] Success: POST http://192.168.1.30/ {"counter": 2, "X-RANDOM": "fe9g83jm"}
[2020-07-28T16:25:40+07:00] Success: POST http://192.168.1.30/ {"counter": 3, "X-RANDOM": "igrijd9p"}

```
## Pembahasan

1. Program untuk menjalankan aplikasi server berada pada directory `server`
2. Program untuk menjalan aplikasi client berada pada directory `client`
3. Hasil log dalam bentuk file berada pada file `server.log` di dalam directory `server`


## How to run

### Aplikasi Server

1. Clone repository ini
2. Buka terminal dan pastikan berada pada directory soal_5
3. Jalankan perintah berikut
```
go run ./server/.
```

### Aplikasi Client

1. Clone repository ini
2. Buka terminal dan pastikan berada pada directory soal_5
3. Jalankan perintah berikut
```
go run ./client/.
```
