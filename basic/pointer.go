package basic

import "fmt"

func Basic_Pointer() {
	// ddefault: pass by value (duplicate the data)
	// human1 := Human{"Usep Bulukumba", "Male", "Futsal"}
	// human2 := human1

	var human1 Human = Human{"Usep Bulukumba", "Male", "Futsal"}
	var human2 *Human = &human1 // &pointer: pass by reference (pointing to the same data)

	human2.Name = "Ajat Summarecon"

	fmt.Println(human1)
	fmt.Println(human1.Name)
	fmt.Println("======")
	fmt.Println(human2)
	fmt.Println(human2.Name)
}
