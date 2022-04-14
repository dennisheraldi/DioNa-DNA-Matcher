package main

import (
	"fmt"
	"stringmatching/method"
)

func main(){
	fmt.Printf("%d",method.KMP("abacaabaccabacabaabb", "abacab"))
	fmt.Printf("%d",method.BM("abacaabaccabacabaabb", "abacab"))
}