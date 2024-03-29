# DioNa: DNA Matcher Web Application

> Program penerapan string matching dan regular expression dalam bentuk website lokal sederhana.
>
> > _Tugas Besar 3 IF2211 Strategi Algoritma_ <br> _Penerapan String Matching dan Regular Expression dalam DNA Pattern Matching_ <br> _Semester II 2021/2022_ <br>

## Table of Contents

- [General Info](#general-information)
- [Technologies Used](#technologies-used)
- [Features](#features)
- [Screenshots](#screenshots)
- [Setup](#setup)
- [Usage](#usage)
- [Contact](#contact)

## General Information

<p> Sebuah aplikasi DNA Pattern Matching dengan memanfaatkan algoritma String Matching dan Regular Expression </p> 
<p> Dapat diakses melalui pranala berikut: https://dnamatch-fend.netlify.app/ <p>

_Program ini dibuat sebagai pemenuhan Tugas Besar 3 IF2211 Strategi Algoritma Semester II 2021/2022._

## Technologies Used

### Languages

- Golang
- JavaScript
- MySQL

### Frameworks / Libraries

- Gin
- Gorm
- ReactJS

## Features

Fitur yang dihadirkan oleh website ini adalah:

- Menerima input penyakit baru berupa nama penyakit dan sequence DNA nya
- Memprediksi seseorang menderita penyakit tertentu berdasarkan sequence DNA-nya
- Memiliki halaman yang menampilkan urutan hasil prediksi dengan kolom pencarian di dalamnya
- Menghitung tingkat kemiripan DNA pengguna dengan DNA penyakit pada tes DNA

## Screenshots
![image](https://user-images.githubusercontent.com/40627156/165942576-b936b187-bd1c-4dd6-ac1e-d8d56fabe553.png)

![image](https://user-images.githubusercontent.com/40627156/165942721-255fac8b-ebf0-4856-bb49-ed2ca29f7c11.png)

![image](https://user-images.githubusercontent.com/40627156/165942804-d892a2ce-1426-4a3b-aa40-c56d1963e9e2.png)


## Setup
### Menyiapkan Basis Data
Buat basis data baru bernama dnamatching dengan menjalankan query 
```
  create database dnamatching
```
<br>

Jalankan basis data tersebut menggunakan query
```
  use dnamatching
```
<br>

Buat tabel baru bernama **penyakits** untuk menyimpan data penyakit menggunakan query berikut
```
  create table penyakits( 
  nama_penyakit varchar(100) not null, 
  dna_penyakit varchar(255) not null, 
  primary key(nama_penyakit) 
  ); 
```
<br>

Buat tabel baru bernama **riwayats** untuk menyimpan data riwayat penyakit menggunakan query berikut
```
  create table riwayats( 
  tanggal_pred date not null, 
  nama_pasien varchar(100) not null, 
  nama_penyakit varchar(100) not null, 
  similarity decimal(10,2) not null, 
  status varchar(11) not null, 
  foreign key(nama_penyakit) references penyakits(nama_penyakit) 
  ); 
```

## Usage
### Menjalankan Backend
Buka folder **api** dengan terminal dan gunakan perintah berikut
```
  go run main.go
```

### Menjalankan Frontend
Buka folder **frontend** dengan terminal setelah database sudah dibuat dan bagian backend sudah jalan dengan menggunakan perintah berikut
```
  npm install
  npm start
```
## Contact

Dibuat oleh Kelompok 23

- William Manuel Kurniawan (13520020)
- Fachry Dennis Heraldi (13520139) <a href="https://github.com/dennisheraldi">GitHub</a>
- Thirafi Najwan Kurniatama (13520157)
