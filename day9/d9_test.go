package main

import "testing"

func TestActuallyCrossing(t *testing.T) {
	e1 := Edge{Point{2, 0}, Point{2, 3}}
	e2 := Edge{Point{0, 2}, Point{3, 2}}
	if !e1.isCrossing(e2) {
		t.Error("Failed crossing e1 y aligned")
	}
	if !e2.isCrossing(e1) {
		t.Error("Failed crossing e1 x aligned")
	}
}

func TestParallelCrossing(t *testing.T) {
	e1 := Edge{Point{2, 0}, Point{2, 3}}
	e2 := Edge{Point{3, 0}, Point{3, 3}}

	if !e1.isParallelTo(e2) {
		t.Error("y aligned Parallels are not parallel")
	}

	if e1.isCrossing(e2) {
		t.Error("y aligned Failed: Parallels are crossing")
	}

	e3 := Edge{Point{0, 2}, Point{3, 2}}
	e4 := Edge{Point{3, 0}, Point{0, 0}}

	if !e3.isParallelTo(e4) {
		t.Error("x aligned Parallels are not parallel")
	}

	if e3.isCrossing(e4) {
		t.Error("x aligned Failed: Parallels are crossing")
	}
}

func TestCollinear(t *testing.T) {
	{
		e1 := Edge{Point{2, 0}, Point{2, 3}}
		e2 := Edge{Point{3, 0}, Point{3, 3}}

		if e1.isCollinearTo(e2) {
			t.Error("apart parallels are collinear")
		}
	}
	{
		e1 := Edge{Point{0, 2}, Point{3, 2}}
		e2 := Edge{Point{3, 0}, Point{0, 0}}

		if e1.isCollinearTo(e2) {
			t.Error("apart parallels are collinear")
		}
	}
	{
		e1 := Edge{Point{3, 0}, Point{3, 5}}
		e2 := Edge{Point{3, 0}, Point{3, 3}}

		if !e1.isCollinearTo(e2) {
			t.Error("failed collinear")
		}
	}
	{
		e1 := Edge{Point{0, 2}, Point{3, 2}}
		e2 := Edge{Point{3, 0}, Point{0, 0}}

		if e1.isCollinearTo(e2) {
			t.Error("apart parallels are collinear")
		}
	}
}
