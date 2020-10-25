package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	var n int
	var s []string
	var cadena string
	fmt.Println("Numero de cadenas a ingresar: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {

		fmt.Scan(&cadena)

		s = append(s, cadena)
	}
	sort.Strings(s)

	file, err := os.Create("ascendente.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	for _, v := range s {
		file.WriteString(v + "\n")
	}

	sort.Sort(sort.Reverse(sort.StringSlice(s)))

	file2, err2 := os.Create("descendente.txt")
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	defer file2.Close()
	for _, v := range s {
		file2.WriteString(v + "\n")
	}

}
