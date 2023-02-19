package mathfunc

import (
	"fmt"
)

func sqroot(num int64) {
	sqrt := num * num
	fmt.Println(sqrt)
}

func cuberoot(num int64) {
	cubert := (num * num) * num
	fmt.Println(cubert)
}
