# **Aqua Developer Batch 2 Day 05**

Tanggal: 26 Agustus 2022

Materi: Golang Lanjutan

## **Pointer**
Pointer merupakan salah satu jenis variabel istimewa di Golang. Fungsinya adalah untuk menyimpan data berupa alamat memori dari variabel lain atau sering dibilang data reference. Karena dia menyimpan data berupa reference, makanya ketika terjadi perubahan data di alamat, maka nilai dari data pointer juga berubah.

```
// Contoh
package main

import "fmt"

func main() {
    // variabel biasa
    var a int = 10

    // pointer
    var b *int = &a

    fmt.Println(a) // 10
    fmt.Println(&a) // misalkan 0x281029
    fmt.Println(b) // sama dengan &a = 0x281029

    // Nilai a diubah
    a = 20

    fmt.Println(b) // tetap = 0x281029
    fmt.Println(*b) // berubah menjadi 20
}
```

## **Method**
Method pada golang mirip dengan function, bedanya untuk method memiliki argumen receiver. Dengan argumen ini method dapat mengakses properti yang dimiliki oleh argumen. Biasa argumen ini bertipe struct.

```
// Contoh
package main

import "fmt"

type Student struct {
    Name    string
    age     int
}

func (s Student) GetStudentName() string {
    return s.Name
}

funct (s *Student) SetStudentName(name string) string {
    s.Name = name
    return s.Name
}

func main() {
    var s1 = Student{"Joko", 20}
    fmt.Println(s1.GetStudentName()) // Joko
    fmt.Println(s1.SetStudentName("Budi")) // Budi
}
```

## **Public and Private**
Public dan private merupakan salah satu hal mendasar dalam bahasa pemrograman. Fungsinya adalah untuk mengurangi kemungkinan adanya kebocoran data. Pada golang implementasi dari private dan public ini sangat mudah, cukup dengan menggunakan huruf kapital pada variabel/fungsi yang ingin dijadikan public.

```
// Contoh
// File a
package lib

func PublicFunction(name string) string {
    return "Hello "+name
}

func privateFunction(name string) string {
    return "Holla "+name
}

// File main
package main

import (
    "fmt"
    "lib"
)

func main() {
    fmt.Println(lib.PublicFunction("Bowo")) // Hello Bowo
    fmt.Println(lib.privateFunction("Bowo")) // Error
}
```

## **Interface**
Interface merupakan custom type yang digunakan untuk membuat signature method(blueprint dari sebuah method). Interface bersifat abstract,

```
// Contoh
package main

import (
    "fmt"
    "strconv"
)

type Student struct {
    Name    string
    Age     int
}

type Greeting interface {
    NormalGreeting()
    AdvanceGreeting()
}

func (s Student) NormalGreeting() string {
    return "Hello "+s.Name
}

func (s Student) AdvanceGreeting() string {
    return "Hello "+s.Name+", your age is"+strconv.Itoa(s.Age)
}

func main() {
    var s1 = Student{"Mega", 20}
    fmt.Println(s1.NormalGreeting()) // Hello Mega
    fmt.Println(s1.AdvanceGreeting()) // Hello Mega, your age is 20
}
```

## **GoRoutine**
GoRoutine merupakan bagian penting dalam concurrent programming pada GO.GoRoutine beralan secara multicore. Konsep concurrent sendiri adalah ketika ada sebuah task, dia akan dilemparkan ke thread lain diluar main thread, kemudian ketika proses sudah selesai akan dilemparkan kembali hasilnya ke main thread. Konsep ini mungkin mirip dengan async-await jika teman-teman lebih familiar.

```
// Contoh sederhana goroutine
package main

import (
    "fmt"
    "sync"
)

func main() {
    wg := sync.WaitGroup{}
    
    wg.Add(2)

    go func() {
        defer wg.Done()
        fmt.Println("Hello")
    }()

    go func() {
        defer wg.Done()
        fmt.Println("World")
    }()

    wg.Wait()
}

// Secara kegunaan mungkin dalam kasus diatas tidak terlalu nampak kasiatnya, namun ketika proses yang berjalan banyak goroutine sangat mantep
```

Kemudian dalam concurrent ada sebuah anomali kalau tidak salah namanya rest condition. Jadi kondisi ini terjadi apabila terdapat banyak request yang berjalan, dan terjadi crash. Misalkan ada 2 request berjalan, servis 1 sudah sampai titik C sedangkan servis 2 masih sampai titik B, maka titik akan ter-reset ke titik B. Terdapat 3 solusi yang dapat dipakai, yaitu locking, Mutex, dan atomic variabel.