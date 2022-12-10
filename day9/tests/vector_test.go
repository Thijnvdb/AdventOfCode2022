package common

import (
	"aoc/day9/common"
	"aoc/shared/types"
	"testing"
)

func TestGetDistanceBetweenWhen2Appart(t *testing.T) {
	point1 := common.NewRopePoint()
	point2 := common.NewRopePoint()
	point2.Position = types.Vector{X: 2, Y: 1}

	if common.GetDistanceBetweenKnots(point1, point2) != common.GetDistanceBetweenKnots(point2, point1) {
		t.Error("received 2 different results for same values")
	}

	dist := common.GetDistanceBetweenKnots(point1, point2)
	if dist != 2 {
		t.Errorf("result was not 2, but %v", dist)
	}
}

func TestGetDistanceBetweenWhen1Appart(t *testing.T) {
	point1 := common.NewRopePoint()
	point2 := common.NewRopePoint()
	point2.Position = types.Vector{X: -1, Y: -1}

	if common.GetDistanceBetweenKnots(point1, point2) != common.GetDistanceBetweenKnots(point2, point1) {
		t.Error("received 2 different results for same values")
	}

	for x := -1; x < 2; x++ {
		for y := -1; y < 2; y++ {
			if x == 0 && y == 0 {
				continue
			}
			v := types.Vector{X: x, Y: y}
			point2.Position = v

			dist := common.GetDistanceBetweenKnots(point1, point2)
			if dist != 1 {
				t.Errorf("distance was not 1, but %v", dist)
			}
		}
	}
}
