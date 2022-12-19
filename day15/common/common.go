package common

import (
	"aoc/shared"
	"aoc/shared/types"
	"strconv"
	"strings"
)

type Sensor struct {
	Position types.Vector
	Beacon   *Beacon
}

type Beacon struct {
	Position types.Vector
}

func ParseInput(inputFile string) ([]*Sensor, error) {
	lines, err := shared.ReadFile(inputFile)
	if err != nil {
		return nil, err
	}

	sensors := []*Sensor{}
	for _, line := range lines {
		stripped := strings.ReplaceAll(line, "Sensor at x=", "")
		split := strings.Split(stripped, ": closest beacon is at x=")

		sensorPos := strings.Split(split[0], ", y=")
		beaconPos := strings.Split(split[1], ", y=")

		beaconX, err := strconv.ParseInt(beaconPos[0], 10, 0)
		if err != nil {
			return nil, err
		}

		beaconY, err := strconv.ParseInt(beaconPos[1], 10, 0)
		if err != nil {
			return nil, err
		}

		beacon := new(Beacon)
		beacon.Position = types.Vector{X: int(beaconX), Y: int(beaconY)}

		sensorX, err := strconv.ParseInt(sensorPos[0], 10, 0)
		if err != nil {
			return nil, err
		}

		sensorY, err := strconv.ParseInt(sensorPos[1], 10, 0)
		if err != nil {
			return nil, err
		}

		sensor := new(Sensor)
		sensor.Beacon = beacon
		sensor.Position = types.Vector{X: int(sensorX), Y: int(sensorY)}
		sensors = append(sensors, sensor)
	}

	return sensors, nil
}
