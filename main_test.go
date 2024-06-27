package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestBasicLinkedListOperations(t *testing.T) {

	ll := NewLInkedList()

	if ll.size != 0 {
		t.Fatalf("expected size : %v, actual size : %v", 0, ll.size)
	}

	if !ll.IsEmpty() {
		t.Fatal("expected empty")
	}

	for _, v := range []int{5, 4, 99} {
		ll.Add(v)
	}

	if ll.size == 0 {
		t.Fatalf("expected size : %v, actual size : %v", 3, ll.size)
	}

	if ll.IsEmpty() {
		t.Fatal("expected not empty")
	}

	expected := "5 -> 4 -> 99"
	if ll.Print() != expected {
		t.Fatalf("expected : %v, actual %v", expected, ll.Print())
	}

}

func TestFindOperations(t *testing.T) {

	ll := NewLInkedList()

	data := []int{5, 4, 99}

	for _, v := range data {
		ll.Add(v)
	}

	for i, v := range data {
		actualIndex := ll.FindIndex(v)
		if i != actualIndex {
			t.Fatalf("Expected index : %v, actual %v", i, actualIndex)
		}

		if _, ok := ll.Find(v); !ok {
			t.Fatalf("Expected %v to be found", v)
		}
	}

	if _, ok := ll.Find(999); ok {
		t.Fatalf("Expected %v to not be found", 999)
	}

}

func TestCheckOperation(t *testing.T) {

	ll := NewLInkedList()

	data := []string{"med", "cat", "banana"}
	for _, value := range data {
		if !ll.Check(value) {
			t.Fatalf("Check Expected %v, actual %v", true, false)
		}
	}

	expected := strings.Join(data, " -> ")
	if ll.Print() != expected {
		t.Fatalf("expected : %v, actual %v", expected, ll.Print())
	}

	if ll.Check("cat") {
		t.Fatalf("Check V Expected %v, actual %v", false, false)
	}

	if ll.Print() != "cat -> med -> banana" {
		t.Fatalf("Print aftec Check expected %v, actual %v", "cat -> med -> banana", ll.Print())
	}

	ll.Check("banana")
	if ll.Print() != "banana -> cat -> med" {
		t.Fatalf("Print aftec Check expected %v, actual %v", "banana -> cat -> med", ll.Print())
	}
}

func TestCacheOperations(t *testing.T) {

	c := NewCache()

	if c.Print() != "" {
		t.Fatalf("Cache print expected %v, actual %v", "", c.Print())
	}

	data := []string{"med", "cat", "banana"}
	for _, value := range data {
		c.Check(value)
	}

	expected := strings.Join(data, " -> ")
	actual := c.Print()
	fmt.Println(actual)
	if expected != actual {
		t.Fatalf("Cache print expected %v, actual %v", c.Print(), strings.Join(data, " -> "))
	}

}
