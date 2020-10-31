package count

import (
	"fmt"
	"strings"
)

func count() {
	s := "how do you do do do    d"
	ss := strings.Split(s, " ")
	m := make(map[string]int, len(ss))

	for _, v := range ss {
		_, ok := m[v]
		if !ok {
			m[v] = 1
		} else {
			m[v]++
		}
	}
	fmt.Println(m)
}
