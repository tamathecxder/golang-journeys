package basic

import "fmt"

type Human struct {
	Name, Gender, Hobby string
}

func Basic_AsteriskOperator() {
	var human1 Human = Human{"Usep Bulukumba", "Male", "Futsal"}
	var human2 *Human = &human1 // &pointer: pass by reference (pointing to the same data)

	human2.Name = "Ajat Summarecon"

	fmt.Println(human1)
	fmt.Println(human1.Name)
	fmt.Println("======")
	fmt.Println(human2)
	fmt.Println(human2.Name)

	fmt.Println("/////-----====;;;;;;;;;;///////")

	// human2 = &Human{"Siti Zaenudyen", "Female", "Padel"} // relocate: to the new data, separating it self from being pointing to the prev memory address
	*human2 = Human{"Siti Zaenudyen", "Female", "Padel"} // renew: all of the reference (anywhere) will be renew/repoint to this memory address

	fmt.Println("original", human1)
	fmt.Println("pointer", human2)
}
