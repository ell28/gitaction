package gitaction

import "fmt"

func main() {
	fmt.Println("Hello, world!")
	fmt.Println("Hello, world2!")
	fmt.Println("Test Commit 3.")
	fmt.Println("Test Commit 4.")
	countNumbers()
}

func countNumbers() {
	count := 0
	for i := 0; i <= 10; i++ {
		count++
	}
	fmt.Println("Total numbers:", count)
}
