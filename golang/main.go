package golang

import "fmt"

func Main() {
	m := make(map[interface{}]int)
	m['a'] = 10
	m[10] = 11
	m["100"] = 12
	m[97] = 10000
	fmt.Println(m[97], m['a'], m["100"])
}
