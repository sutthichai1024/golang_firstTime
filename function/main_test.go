package main

import "testing"

// func TestPlus(t *testing.T) {
// 	a := plus(1, 4)

// 	if a != 3 {
// 		t.Errorf("Expected result of 3 , but it was")
// 	}
// }

// func TestPrism(t *testing.T) {
// 	a := prism(4, 2, 2)

// 	if a != 16 {
// 		t.Errorf("Expected result of 3 , but it was ")
// 	}

func benchmarkFact10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fact(10)
	}
}

func benchmarkFactUsingfor10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		factUsingfor(10)
	}
}
