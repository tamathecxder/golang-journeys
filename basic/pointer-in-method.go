package basic

import "fmt"

type Shawty struct {
	Name string
}

func (shawty *Shawty) getShawty() {
	var baddieNames []string = []string{
		"Jessica", "Natalia", "Amanda",
	}

	for _, name := range baddieNames {
		if shawty.Name == name {
			shawty.Name = shawty.Name + " is a baddie!"
			return
		}
	}

	shawty.Name = shawty.Name + " is a regular girl"
}

func Basic_PointerInMethod() {
	urShawty := Shawty{"Nancy"}   // default
	urShawty2 := Shawty{"Amanda"} // pointer

	fmt.Println(urShawty.Name)
	fmt.Println(urShawty2.Name)

	urShawty.getShawty()
	urShawty2.getShawty()

	fmt.Println(urShawty.Name)
	fmt.Println(urShawty2.Name)
}
