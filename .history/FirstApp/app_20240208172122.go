// This is a simple Go program that prints "Hello, World!" to the console.
package main // This is the package declaration. It tells Go that this is the main package and that it should be executed as an application.
import (
	"bufio"
	"fmt"
	"os"
)// This is the import section. It tells Go that we want to use the fmt package to print to the console.

func main() {
	fmt.Print("Hello, World!")
	fmt.Print("Press any key to exit...")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
	
}