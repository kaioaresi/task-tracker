package cli

import (
	"fmt"
	"os"
)

func main() {

	task := os.Args[0]

	fmt.Println(task)
}
