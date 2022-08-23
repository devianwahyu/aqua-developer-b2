# **Aqua Developer Batch 2 Day 02**

Tanggal: 23 Agustus 2022

Materi: Version Control (GIT)

## **Commands**

### **Init** 

command: 
```
    1. git init --quiet (default)
    2. git init --bare
```

fungsi: Untuk menginisialisasi git pada projek yang akan di manage

penggunaan: Masuk ke directory yang diinginkan > tulis command ```git init``` di terminal

### **Add** 

commands: 
```
    1. git add . (all changes)
    2. git add <file_name>
    3. git add <directory_name>
```

fungsi: Menambahkan perubahan dari working tree ke staging area

penggunaan: Setelah melakukan perubahan > tulis command ```git add .``` di terminal

### **Status** 

commands: 
```
    1. git status
```

fungsi: Melihat state dari suatu file, apakah dia sudah masuk ke staging atau masih di working area, atau malah masih belum di track

penggunaan: Tulis command ```git status``` di terminal

### **Remote** 

commands: 
```
    1. git remote add <remote_name> <remote_url>
    2. git remote -v (check remote)
```

fungsi: Secara sederhana untuk menyambungkan repo local dengan repo online

penggunaan: Setelah melakukan init git di directory > tulis command ```git remote add origin https://github.com/devianwahyu/aqua-developer-b2.git``` di terminal

### **Commit** 

commands: 
```
    1. git commit -a
    2. git commit -m "message"
```

fungsi: Melakukan snapshot projek yang berada pada staging area yang akan di push ke repo online

penggunaan: Setelah melakukan ```git add .``` > tulis command ```git commit -m "feat: add new feature"``` di terminal

### **Push** 

commands: 
```
    1. git push -u origin master
    2. git push -u origin (misal branch-nya banyak takut salah push)
```

fungsi: Mengirimkan perubahan yang sudah di commit ke repo online

penggunaan: Setelah melakukan ```git commit -m "feat: add new feature"``` > tulis command ```git push -u origin master``` di terminal

### **Branch** 

commands: 
```
    1. git branch <branch_name>
    2. git branch -a (list branch)
```

fungsi: Membuat branch, branch berguna untuk menyimpan pekerjaan supaya mengurangi kemungkinan terjadinya konflik apabila banyak kontributor

penggunaan: Setelah melakukan ```git init``` > tulis command ```git branch new_feature``` di terminal

### **Fetch** 

commands: 
```
    1. git fetch --all
    2. git fetch origin
```

fungsi: Mengambil metadata terbaru dari repo online, tidak terjadi transaksi perubahan code disini, lebih kearah memperbaharui jika ada struktur baru misal ada branch baru di repo online

penggunaan: Tulis command ```git fetch origin``` di terminal apabila apa perubahan di repo online

### **Pull** 

commands: 
```
    1. git pull (jika ada di branch yang sesuai)
    2. git pull origin master (jika ingin pull dari master)
```

fungsi: Secara sederhana digunakan untuk melakukan update local repo dengan remote repo

penggunaan: Tulis command ```git pull origin master``` di terminal apabila apa perubahan di repo online

### **Merge** 

commands: 
```
    1. git merge <branch_name>
```

fungsi: Digunakan untuk melakukan menggabungkan 2 branch

penggunaan: Masuk ke branch tujuan misal master ```git checkout master```> Tulis command ```git merge new_feature``` di terminal

### **Stash** 

commands: 
```
    1. git stash
    2. git stash list (melihat list stash)
    3. git stash pop <index> (mengembalikan stash ke working area)
```

fungsi: Secara sederhana digunakan untuk menyimpan sementara perubahan ke suatu lokasi, ketika perubahan ini tidak ingin di commit

penggunaan: Setelah melakukan perubahan > Tulis command ```git stash``` di terminal 

### **Commit Convention** 

pendekaan: git karma

commands: 
```
    <type>(<scope> (optional)): <subject>
    <body> (optional)
    <footer> (optional)
```

fungsi: Supaya pesan commit lebih jelas dan mudah dipahami oleh kontributor lain dari sebuah repo

penggunaan: ```git commit -m "fix(authentication): menambahkan jwt pada proses authentication"```

commit convention:
- feat: penambahan fitur
- fix: memperbaiki bug
- perf: improve performance
- docs: menambah dokumentasi
- style: melakukan style code, misal semicolon
- refactor: melakukan perubahan di code misal nama variabel
- test: menambahkan test yang tertinggal
- build: melakukan update configurasi

### **Semantic Versioning** 

commands: 
```
    V<major>.<minor>.<patch>
```

fungsi: Untuk memberikan versi pada aplikasi

penggunaan: V2.1.0

keterangan:
- major: breaking change pada API
- minor: penambahan fitur
- patch: fix, perf

### **Git Management** 

Menggunakan **Trunk Based Development**, secara sederhana tipe management ini adalah dengan melakukan pull/fetch dari master branch ke branch lain yang digunakan untuk proses development. Keuntungan dari metode ini adalah proses bisnis yang terjadi lebih singkat daripada git flow, namun tetap efektif untuk mengurangi kemungkinan terjadinya konflik.