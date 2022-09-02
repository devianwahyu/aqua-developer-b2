# **Aqua Developer Batch 2 Day 08**

Tanggal: 31 Agustus 2022

Materi: Cloud Native

## **Definisi**
Cloud native architecture adalah pendekatan yang digunakan untuk design, constructing, dan operating workload yang dibangung di cloud dan memiliki banyak kelebihan.

## **Pillars**
Cloud infrastructure (Microservices, Modern Design, Containers, Backing Services, Automation)

## **Twelve Factor Application**
1. Codebase
2. Dependencies
3. Configuration
4. Backing Service
5. Build, Release, Run
6. Process
7. Port Binding
8. Concurrency
9. Dev/Prod Parity
10. Logging
11. Admin Process
12. Disposability

## **Containers**

![docker_logo](./assets/logo-docker.png)

Container seperti diawal, merupakan salah satu pillar dari cloud native app. Salah satu tool untuk melakukan containerization adalah Docker.

## **Containerization**
Containerization merupakan abstraksi dari app layer yang dibungkus dalam tempat, terdiri atas code dan dependencies.

Contoh analoginya adalah, misalkan kita mengembangkan aplikasi bersama orang lain, kemudian kita menggunakan sebuah framework. Ternyata framework yang kita pakai tidak sama versinnya dengan punya teman kita. Perbedaan ini dapat menyebabkan error. Karenanya supaya tidak perlu gonta-ganti versi di local komputer, dapat menggunakan docker.

## **Docker Architecture**
Docker menggunakan arsitektur client-server. 

![docker_architecture](./assets/architecture.svg)

## **Docker VS VM**
Perbedaan antara cara kerja docker dengan vm

![docker_vs_vm](./assets/containers-vs-virtual-machines.jpg)

## **Term in Docker**
### **Docker registry**
- Sebuah registry adalah servis yang mengandung respositories of images
- Local registry adalah tempat images disimpan di local
- Remote registry adalah tempat images disimpan di online (Azure, GCP, Docker)

### **Dockerfile**
Dockerfile merupakan text document yang mengandung semua command yang user dapat lakukan untuk assemble image

### **Docker Image**
Docker image merupakan read-only template dengan instruksi untuk membangun docker container

### **Docker Container**
Merupakan runnable instance dari image, sederhananya adalah image (kumpulan code & dependencies) yang sudah dapat dijalankan di berbagai environment docker