package main

import "testing"

/*
func Test_sum(t *testing.T) {
	x := 5
	y := 5

	want := 10
	got := sum(x,y)

	t.Logf("Testing a sum of: %d %d", x, y)
	if want != got {
		t.Errorf("Sum was incorrect. Want: %d, but got: %d", want,got)
	}
}


 */

func Test_sum(t *testing.T) {
	tables := []struct{
		x int
		y int
		want int
	}{
		{1,1,2},
		{5,5,10},
		{-1, 5, 4},
	}

	for _, table := range tables {
		got := sum(table.x, table.y)
		if got != table.want {
			t.Errorf()
		}
	}
}