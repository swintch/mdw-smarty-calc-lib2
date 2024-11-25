package main

import (
	"log"
	"os"

	"github.com/swintch/mdw-smarty-calc-lib2/calc"
	"github.com/swintch/mdw-smarty-calc-lib2/handler"
)

func main() {
	calculator := &calc.Addition{}
	handlerObj := handler.NewHandler(calculator, os.Stdout)
	err := handlerObj.Handle(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

}
