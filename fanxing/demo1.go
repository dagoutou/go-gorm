package main

import "fmt"

func test[T any](t T) T {
	return t
}

type Tmap[k string | int, v string | int]map[k]v

func main() {
	a := test("aaaa") + test("bbbb")
	fmt.Println(a)
	m := make(Tmap[int, string])
	m[1] = "adfsf"
	fmt.Println(m)

}
