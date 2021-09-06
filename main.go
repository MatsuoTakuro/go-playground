package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 0; i < 3; i++ {
		/*
		   go func() {
		       fmt.Println(i)
		   }()
		*/
		go func(i int) {
			fmt.Println(i)
		}(i)
		time.Sleep(100 * time.Millisecond)
	}
	/*
	   (実行結果)
	   0
	   2
	   1
	   (0,1,2が順不同で出力)
	*/
}
