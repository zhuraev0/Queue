package main

import "fmt"

type OrderQueue struct {
	next, prev *OrderQueue
	Value      interface{}
}

type people struct {
	first, last *OrderQueue
	length      int
}

func (receiver *people) len() int {
	return receiver.length
}

func (receiver *people) Prev() interface{} {
	if receiver.first != nil {
		return receiver.first.Value
	}
	return nil
}

func (receiver *people) Next() interface{} {
	if receiver.last != nil {
		return receiver.last.Value
	}
	return nil
}
/*
func (receiver *people) Len() int {
	return receiver.len()
}
*/
func (receiver *people) eQueue(value interface{}) {
	if receiver.len() == 0 {
		receiver.first = &OrderQueue{
			next:  nil,
			prev:  nil,
			Value: value,
		}
		receiver.last = receiver.first
		receiver.length++
		return
	}
	receiver.length++

	current := receiver.first
	for {
		if current.next != nil {
			receiver.last = current.next
			break
		}
		current.next = &OrderQueue{
			next:  nil,
			prev:  current,
			Value: value,
		}
		current = current.next
	}
}

func (receiver *people) deQueue() people {
	if receiver.len() == 0 {
		return people{}
	}

	result := people{
		first:  receiver.first,
		last:   nil,
		length: 0,
	}
	receiver.first = receiver.first.next
	receiver.length--
	if receiver.length == 0 {
		receiver.last = receiver.first
	}
	return result
}
func main() {

	t := people{}
	t.eQueue("Vasya")
	t.eQueue("Sanya")
	t.eQueue("Vanya")
	t.eQueue("Petya")


	for i := 0; i <= 4; i++ {
		fmt.Println(t.deQueue())
	}
}
