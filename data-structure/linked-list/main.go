package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func main() {
	// Inicialización de nodos
	var head *Node	//Apuntamos a la dirección de memoria del primer nodo, Ejemplo: head apunta a la dirección de la memoria en 0x0001
	var one, two, three *Node //Apuntamos a la dirección de memoria de los nodos, Ejemplo: one, two y three apuntan a la dirección de la memoria en 0x0001, 0x0002, 0x0003 respectivamente

	// Asignación de memoria
	one = &Node{}
	two = &Node{}
	three = &Node{}

	// Asignación de valores
	one.data = 1
	two.data = 2
	three.data = 3

	// Conectar nodos
	one.next = two
	two.next = three
	three.next = nil

	// Guardar la dirección del primer nodo en head
	head = one

	// Para demostrar, imprimimos los valores de la lista
	printList(head)
}

func printList(node *Node) {
	current := node
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}
