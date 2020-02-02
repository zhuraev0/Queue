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

func Test_queue_with_multiple_element(t *testing.T) {
	list := people{}
	list.eQueue("1")
	list.eQueue("Masha")
	list.eQueue("Sasha")
	if list.len() != 3 {
		t.Error("after adding 3 elements size must be 3, got: ", list.len())
	}
	if list.Prev() != "1" {
		t.Error("The first in queue != \"1\", got: ", list.Prev())
	}
	if list.Next() != "Sasha" {
		t.Error("The last in queue != \"Sasha\", got: ", list.Next())
	}
	err := list.Next()
	if err != nil {
		t.Error("After add last in queue \"55\", got: ", list.Next())
	}
	list.deQueue()
	list.eQueue("Dasha")
	if list.Prev() == "Dasha" {
		t.Error("After remove the first element in queue \"Masha\", the first in queue must be \"Dasha\", got: ", list.Prev())
	}
}

func Test_queue_for_deleting_item(t *testing.T) {
	list := people{}
	list.eQueue(5)
	list.deQueue()
	if list.Prev() != nil {
		t.Error("deleting one prev queue , got: ", list)
	}
}
