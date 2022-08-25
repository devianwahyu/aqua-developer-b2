# **Aqua Developer Batch 2 Day 04**

Tanggal: 25 Agustus 2022

Materi: Pengantar Golang

## **Definisi**
Bahasa pemrograman multi-purpose yang bisa dibilang baru, digunakan untuk membuat software dengan performa yang mantep. Dikembangkan oleh Google 2007 akhir dan release v1.0 di tahun 2012.

## **Company Pengguna**
Ada banyak perusahaan besar yang memakai Golang sebagai bahasa untuk handle service mereka, seperti Google, eFisher, Tokopedia, Gojek, Uber, dll.

## **Data Type**
1. Numeric Non Decimal (uint, int rune, dll)
2. Numeric Decimal (float, double)
3. Boolean
4. String

## **Variable**
### Anatomi/Contoh
```
// Biasa
var varA string
var varB string = "Variable B"

// Tanpa Nilai Data Type
var varA = "Variable A"

// Shorthand
varA := "Variable A"

// Multivariable
var varA, varB string
varA = "Tes A"
varB = "Tes B"

varA, varB := "Test A", "Test B"

// Unused
_ := "Ga Dipakai"
```

## **Constant**
### Anatomi/Contoh
```
const constA string = "Konstanta A"

// Multiple Declare
const (
    constA = "halo",
    constB = "hola"
)
```

## **Condition**
### Anatomi/Contoh
```
// IF-ELSE
i := 10

if i > 5 {
    fmt.Println("besar")
} else if i == 10 {
    fmt.Println("sama")
} else {
    fmt.Println("kecil")
}

// SWITCH-CASE
switch i{
    case 11:
        fmt.Println("besar")
    case 10:
        fmt.Println("sama")
    case 9:
        fmt.Println("kecil")
    default:
        fmt.Println("Join eFishery Mantep")
}
```

## **Looping**
### Anatomi/Contoh
```
var array = [4]string{"a", "b", "c", "d"}

// Range
for _, value := range array {
    fmt.Println(value)
}

// Loogin for biasa
for i := 0; i < len(array); i++ {
    fmt.Println(value)
}

// Break, Continue
if break then out looping
if continue then i++
```

## **Function**
### Anatomi/Contoh
```
// Biasa multiple value
func swap(x int, y int) (int, int) {
    return y, x
}

// Parameter dengan tipe sama
func swap(x, y int) (int, int) {
    return y, x
}
```

## **Struct**
### Anatomi/Contoh
```
type school struct{
    name    string
    address string
}

// Embedded
type student struct{
    name    string
    grade   int
    sekolah school
}

// Combine Slice
var allStudent = []student{
    {name: "joko", grade: 10, sekolah: {school{name: "Budi Utomo", adress: "Surabaya"}}}
}
```

## **Data Structure**
### Anatomi/Contoh
```
// Map
var chicken map[string]int
chicken = map[string]int{}

chicken["januari] = 50

fmt.Println(chicken["januari"])

// Array
var pig = [2]string{"pig1", "pig2"}
var cow = [...]string{"cow1", "cow2", "cown"}
```

## **Slice**
### Anatomi/Contoh
```
var animals = []string{"sapi", "kida", "trenggiling"}

/*
memiliki 3 atribut, yaitu
1. ptr      : pointer
2. len      : length
3. cap      : capacity
*/

var someAnimal = animals[1:3] // ["kida", "trenggiling"]
len(animals) // 3
cap(animals) // 3

animals.append("jerapah")
cap(animals) // 6 => 2 x len(animals) awal
```

## **Defer**
### Anatomi/Contoh
```
defer fmt.Println("hello")
fmt.Println("world")

// output => world kemudian hello
```