package revertString

import "fmt"

func Revert(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	fmt.Println("teste")
	fmt.Println("testingv5")
	return string(b)
}
