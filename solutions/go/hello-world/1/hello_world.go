package greeting

import(
    "fmt"
)
// HelloWorld greets the world.
func HelloWorld() string {
	return "Hello, World!"
}
// main is the main entry point of the program
func main(){
    greeting := HelloWorld()
    fmt.Printf("%v",greeting)
}