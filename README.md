# buzzwords : Buzzword generator package

## Overview [![GoDoc](https://godoc.org/github.com/hajhatten/buzzwords?status.svg)](https://godoc.org/github.com/hajhatten/buzzwords)

Go package to easily generate buzzwords, with the appropriate suffix. Can also be used to generate fake loading messages meaning almost nothing. Feel free to add more words if there's anything you're missing.

## Install

```
go get github.com/hajhatten/buzzwords
```

## Example

```
package main

import (
	"fmt"
	"github.com/hajhatten/buzzwords"
)

func main() {
	// Generates a sentence of buzzwords "petabyte databases"
	fmt.Println(buzzwords.BuzzWords())
	
	// Generates a sentence with a suffix, 
	// for example "petabyte databases as a service"
	fmt.Println(buzzwords.WithSuffix()) 

	// Generates a sentence with a verb in front, 
	// for example "reticulating petabyte databases"
	fmt.Println(buzzwords.WithVerb())

	// Generates a sentence with a verb in front and suffix, 
	// for example "reticulating petabyte databases as a service"
	fmt.Println(buzzwords.WithVerbAndSuffix())
}
```

## Author

Rikard Gynnerstedt

## License

MIT.
