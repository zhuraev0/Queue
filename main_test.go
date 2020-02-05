package main

import (
	"testing"
)

func TestList_add(t *testing.T) {
	list := people{}
	if list.len() != 0 {
		t.Error("empty queue length must be zero, got:", list.len())
	}
	if list.Prev() != nil {
		t.Error("first element of queue must have value of 0, got:", list.Prev())
	}
	if list.Next() != nil {
		t.Error("last element of queue must have value of 0, got:", list.Next())
	}
	if list.Next() != list.Prev() {
		t.Error("first and last elements of one item queue must have same priorities, got:", list.Next(), ";", list.Prev())
	}
}

func TestOneItem(t *testing.T) {
	list := people{}
	list.eQueue(1)
	if list.len() != 1 {
		t.Error("after adding one 1, got:", list.len())
	}
	if list.Prev() != 1 {
		t.Error("have first element, got:", list.Prev())
	}
	if list.Next() != 1 {
		t.Error("have last element, got:", list.Next())
	}
	if list.Next() != list.Prev() {
		t.Error("first and last elements of one item queue must have same values, got:", list.Next(), ";", list.Prev())
	}
}
func TestMultipleElement(t *testing.T) {
	list := people{}
	list.eQueue("Vasya")
	list.eQueue("Sanya")
	list.eQueue("Vanya")
	firstQueue := list.Prev()
	if firstQueue != "Jack" {
		t.Error("The first in queue must be == \"Vasya\", got: ", firstQueue)
	}
	lastQueue := list.Next()
	if lastQueue != "Vanya" {
		t.Error("The last in queue must be == \"Vanya\", got: ", lastQueue)
	}
	list.eQueue("Petya")
	if list.Next() != "Petya" {
		t.Error("After add last in queue \"Petya\", the last must be, got: ", list.Next())
	}
	list.deQueue()
	if list.Prev() == "Vasya" {
		t.Error("After remove the first element in queue \"Vasya\", first in queue must be == \"Sanya\", got: ", list.Prev())
	}
}
func TestDeletingItem(t *testing.T) {
	list := people{}
	list.eQueue(5)
	list.deQueue()
	if list.Prev() != nil {
		t.Error("deleting one prev queue , got: ", list)
	}
}
