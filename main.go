package main
 
import "fmt"
 
var staticList []string = []string{"item1", "item2", "item3"}
 
var staticMap map[string]struct{} = map[string]struct{}{
	"item1": {},
	"item2": {},
	"item3": {},
}
 
func main() {
	var m = make(map[string]string)
 
	for key := range staticMap {
		fmt.Printf("Key: %s\n", key)
	}
 
}
