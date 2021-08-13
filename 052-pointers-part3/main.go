package main

import (
	"fmt"
	"strings"
)

type school struct {
	name string
	capacity int
}


func main() {

	fmt.Printf("%-13s%-19s%-22s%-28s%-40s%-20s\n", "Type", "Name", "MapVar/SlHdrAddr", "MapHdr/BackArr/StructAddr", "Description", "Elements")
	fmt.Printf("%s\n", strings.Repeat("=", 170))

	// Similar to a slice, map also does not require a pointer to update values
	// inside function because the map variable itself is a pointer to
	// the map header.

	medals := map[string]int{
		"Alpha":   3,
		"Bravo":   10,
		"Charlie": 1,
	}

	fmt.Printf("%-13s%-19s%-22p%-28p%-40s%#[4]v\n", "map", "medals", &medals, medals, "Inside main(), original")
	updateMedals(medals)
	fmt.Printf("%-13s%-19s%-22p%-28p%-40s%#[4]v\n", "map", "medals", &medals, medals, "Inside main(), after func call")


	// Updating values of struct list inside a function
	// will also update the source as the data
	// has been properly arranged in the backing array
	privateSchool := []school{{
		name: "Amazing School",
		capacity: 10,
		},
	}
	fmt.Printf("%-13s%-19s%-22p%-28p%-40s%#[4]v\n", "struct slice", "privateSchool", &privateSchool, privateSchool, "Inside main(), original")
	updateCapacity(privateSchool)
	fmt.Printf("%-13s%-19s%-22p%-28p%-40s%#[4]v\n", "struct slice", "privateSchool", &privateSchool, privateSchool, "Inside main(), after func call")

	// However, updating struct value directly from function
	// will not work because the structs inside and outside
	// function are located in different memory address.
	// So, the original one will never "see" any changes
	// made inside function.
	anotherSchool := school{
		name: "Excellent School",
		capacity: 10,
		}

	fmt.Printf("%-13s%-19s%-22s%-28p%-40s%#[4]v\n", "struct", "anotherSchool", "-", &anotherSchool, "Inside main(), original")
	updateCapacityStruct(anotherSchool)
	fmt.Printf("%-13s%-19s%-22s%-28p%-40s%#[4]v\n", "struct", "anotherSchool", "-", &anotherSchool, "Inside main(), after func call")


	// The fix to the previous struct problem is to use pointer.
	// Now, structs inside and outside function will point to the
	// same memory location
	anotherSchoolFixed := school{
		name: "Superb School",
		capacity: 10,
		}

	fmt.Printf("%-13s%-19s%-22s%-28p%-40s%#[4]v\n", "struct", "anotherSchoolFixed", "-", &anotherSchoolFixed, "Inside main(), original")
	updateCapacityStructFixed(&anotherSchoolFixed)
	fmt.Printf("%-13s%-19s%-22s%-28p%-40s%#[4]v\n", "struct", "anotherSchoolFixed", "-", &anotherSchoolFixed, "Inside main(), after func call")
	}	

func updateMedals(m map[string]int) {
	fmt.Printf("%-13s%-19s%-22p%-28p%-40s%#[4]v\n", "map", "medals", &m, m, "Inside updateMedals(), before update")
	m["Alpha"] += 7
	m["Bravo"] += 3
	m["Charlie"] += 2
	fmt.Printf("%-13s%-19s%-22p%-28p%-40s%#[4]v\n", "map", "medals", &m, m, "Inside updateMedals(), after update")
}

func updateCapacity(s []school) {
	fmt.Printf("%-13s%-19s%-22p%-28p%-40s%#[4]v\n", "struct slice", "privateSchool", &s, s, "Inside updateCapacity(), before update")
	for i := range s {
		s[i].capacity++
	}
	fmt.Printf("%-13s%-19s%-22p%-28p%-40s%#[4]v\n", "struct slice", "privateSchool", &s, s, "Inside updateCapacity(), after update")
}

func updateCapacityStruct(s school) {
	fmt.Printf("%-13s%-19s%-22s%-28p%-40s%#[4]v\n", "struct", "anotherSchool", "-", &s,  "Inside upda...ruct(), before update")
	s.capacity++
	fmt.Printf("%-13s%-19s%-22s%-28p%-40s%#[4]v\n", "struct", "anotherSchool", "-", &s, "Inside upda...ruct(), after update")
}

func updateCapacityStructFixed(s *school) {
	fmt.Printf("%-13s%-19s%-22s%-28p%-40s%#[4]v\n", "struct", "anotherSchoolFixed", "-", s,  "Inside upda...ixed(), before update")
	s.capacity++
	// s.capacity++ above is equivalent to (*s).capacity++
	// because Go adds the indirection operator behind the scene
		
	fmt.Printf("%-13s%-19s%-22s%-28p%-40s%#[4]v\n", "struct", "anotherSchoolFixed", "-", s, "Inside upda...ixed(), after update")
}
