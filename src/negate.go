package main

import (
    "fmt"
)
func negate(myBoolean bool) bool {
     return !myBoolean
}
func main() {
   truth := true
   truth = negate(truth)
   fmt.Println(truth)
   lies := false
   lies = negate(lies)
   fmt.Println(lies)
}
   
