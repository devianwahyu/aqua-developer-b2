package main

import "fmt"

// Declare Struct
type person struct {
	name string
	age  int
	kos  kos
}

type kos struct {
	name    string
	address string
}

func main() {
	// Init Struct
	var manusia person
	manusia.name = "Devian"
	manusia.age = 22
	manusia.kos = kos{
		name:    "Kos Mantep",
		address: "Jl. Joyoagung No. 666",
	}

	// Print Struct
	fmt.Printf("Nama: %s\nAge: %d\nKos: %s\nAlamat: %s\n\n",
		manusia.name,
		manusia.age,
		manusia.kos.name,
		manusia.kos.address,
	)

	// Edit Struct
	manusia.name = "Jokowi"

	fmt.Printf("Nama: %s\nAge: %d\nKos: %s\nAlamat: %s\n\n",
		manusia.name,
		manusia.age,
		manusia.kos.name,
		manusia.kos.address,
	)

	// Array
	var people = [2]person{{
		name: "Luhut",
		age:  68,
		kos: kos{
			name:    "Istana Presiden",
			address: "Jakarta",
		},
	}}

	// Assign Manual
	people[1] = person{
		name: "Amien",
		age:  80,
		kos: kos{
			name:    "Istana Presiden",
			address: "Jakarta",
		},
	}

	// Print Array
	for _, value := range people {
		fmt.Printf("Nama: %s\nAge: %d\nKos: %s\nAlamat: %s\n\n",
			value.name,
			value.age,
			value.kos.name,
			value.kos.address,
		)
	}

	// Edit Array
	people[0] = person{
		name: "Budiono",
		age:  77,
		kos: kos{
			name:    "Istana Presiden",
			address: "Jakarta",
		},
	}

	for _, value := range people {
		fmt.Printf("Nama: %s\nAge: %d\nKos: %s\nAlamat: %s\n\n",
			value.name,
			value.age,
			value.kos.name,
			value.kos.address,
		)
	}

	// Slice
	var peopleSlice = []person{{
		name: "Luhut",
		age:  68,
		kos: kos{
			name:    "Istana Presiden",
			address: "Jakarta",
		},
	}}

	// Assign Manual
	newSlice := append(peopleSlice, person{
		name: "Amien",
		age:  80,
		kos: kos{
			name:    "Istana Presiden",
			address: "Jakarta",
		},
	})

	// Print SLice
	for _, value := range newSlice {
		fmt.Printf("Nama: %s\nAge: %d\nKos: %s\nAlamat: %s\n\n",
			value.name,
			value.age,
			value.kos.name,
			value.kos.address,
		)
	}

	// Edit Slice
	newSlice[0] = person{
		name: "Megawati",
		age:  77,
		kos: kos{
			name:    "Istana Presiden",
			address: "Jakarta",
		},
	}

	// Print Slice
	for _, value := range newSlice {
		fmt.Printf("Nama: %s\nAge: %d\nKos: %s\nAlamat: %s\n\n",
			value.name,
			value.age,
			value.kos.name,
			value.kos.address,
		)
	}

	// Map
	var peopleMap map[string]person
	peopleMap = map[string]person{}

	// Assign Map
	peopleMap["person1"] = person{
		name: "Supri",
		age:  77,
		kos: kos{
			name:    "Istana Presiden",
			address: "Jakarta",
		},
	}

	// Print Map
	fmt.Printf("Nama: %s\nAge: %d\nKos: %s\nAlamat: %s\n\n",
		peopleMap["person1"].name,
		peopleMap["person1"].age,
		peopleMap["person1"].kos.name,
		peopleMap["person1"].kos.address,
	)

	// Edit Map
	peopleMap["person1"] = person{
		name: "Sukroini",
		age:  77,
		kos: kos{
			name:    "Istana Presiden",
			address: "Jakarta",
		},
	}

	// Print Map
	fmt.Printf("Nama: %s\nAge: %d\nKos: %s\nAlamat: %s\n\n",
		peopleMap["person1"].name,
		peopleMap["person1"].age,
		peopleMap["person1"].kos.name,
		peopleMap["person1"].kos.address)
}
