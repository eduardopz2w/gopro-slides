// +build OMIT
package main

import (
	"container/list"
	"fmt"
)

func main() {
	// Crea una lista y pone número en ella
	l := list.New()
	// Agregamos elementos a la lista
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	// Insertamos elementos en la lista
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)

	// Recorremos la lista y se imprime el contenido.
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

}
