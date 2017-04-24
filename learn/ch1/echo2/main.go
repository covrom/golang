// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 6.
//!+

// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	stt := time.Now()

	s, sep := "", ""
	for ind, arg := range os.Args[1:] {
		s += sep + fmt.Sprintf("%d %s",ind,arg)
		sep = "\n"
	}
	fmt.Println(s)

	fmt.Printf("%dns\n",time.Since(stt).Nanoseconds())
}

//!-
