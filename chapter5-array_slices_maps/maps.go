package main

import "fmt"

func main() {
	//map is an unordered collection of key-value pairs

	// create a map
	x := make(map[string]int)

	x["key"] = 10
	fmt.Println(x)
	fmt.Println(x["key"])

	//delete the map items using delete function
	delete(x, "key")

	fmt.Println(x)

	//a better way to define map
	elements1 := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B":  "Boron",
		"C":  "Carbon",
		"N":  "Nitrogen",
		"O":  "Oxygen",
		"F":  "Fluorine",
		"Ne": "Neon",
		"a":  "",
	}

	// for check key is present there or not
	name, ok := elements1["a"]
	fmt.Println(name, ok)
	name1, ok1 := elements1["asdas"]
	fmt.Println(name1, ok1)

	//map of map
	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
	}

	if el, ok := elements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}
}
