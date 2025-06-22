package main

import "fmt"

func main() {
	var cities = [4]string{
		"Cianjur", "Bandung", "Bogor", "Sukabumi",
	}

	// [start:end]
	// start = will be the init pointer of the slice to the array
	// end = will be the maximum value that slice will gonna take but it will take after the max index
	citySlices := cities[0:2]
	fmt.Println(cities)
	fmt.Println(citySlices)

	var days = [7]string{
		"Sunday",
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
	}

	fmt.Println(days)

	twoLastDays := days[len(days)-2:] // if the end is not filled out, it will takes until the rest of the array
	fmt.Println(twoLastDays)

	fmt.Println("============")

	daysBeforeWednesday := days[:3] // if the start is not filled out, it will takes all over the starting index of the array untill the end/max slice point
	fmt.Println("days before wednesday = ", daysBeforeWednesday)

	allDays := days[:] // if both sides is not filled out, it will takes all array values
	fmt.Println("all days = ", allDays)

	allDays[len(allDays)-1] = "Saturday Night"

	fmt.Println("New Days = ", days)

	fmt.Println("============")

	// is the original array impacted all will created a new array? lets find out
	fmt.Println(allDays)
	fmt.Println(len(allDays))
	fmt.Println(cap(allDays)) // 7
	newDays := append(allDays, "Mingse")
	fmt.Println(newDays)
	fmt.Println(len(newDays))
	fmt.Println(cap(newDays))        // 14
	newDays[0] = "FirstDayOnEarch"   // it will update the new array, because the capacity is overload it will create a new array + prev array values into this
	allDays[0] = "SUNDAY_CAPITALIZE" // it will update the og array and the slice also
	fmt.Println(newDays)
	fmt.Println(allDays)
	fmt.Println(days)

	fmt.Println("============")

	// make a new slice without initialize an array
	ogSlice := make([]string, 2, 3)
	ogSlice[0] = "Yudistira"
	ogSlice[1] = "Eka"
	// ogSlice[3] = "Whatever" // out of range, if u want to insert new value to the slice, you may have use the append method instead of injecting the value using index assignemnt

	fmt.Println(ogSlice)
	fmt.Println(len(ogSlice))
	fmt.Println(cap(ogSlice))

	appendedSlice := append(ogSlice, ">>END_NAME")

	fmt.Println(appendedSlice)
	fmt.Println(len(appendedSlice))
	fmt.Println(cap(appendedSlice))

	appendedSlice[0] = "ASEP"

	fmt.Println("is updated to the append?", appendedSlice)
	fmt.Println("is updated to the og?", ogSlice)

	fmt.Println("============")

	// copying the slice
	initDays := days[:]
	copyDays := make([]string, len(days), cap(days))
	copy(copyDays, initDays)

	fmt.Println(initDays)
	fmt.Println(copyDays)
}
