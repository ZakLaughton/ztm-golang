//--Summary:
//  Create a system monitoring dashboard using the existing dashboard
//  component structures. Each array element in the components represent
//  a 1-second sampling.
//
//--Requirements:
//* Create functions to calculate averages for each dashboard component
//* Using struct embedding, create a Dashboard structure that contains
//  each dashboard component
//* Print out a 5-second average from each component using promoted
//  methods on the Dashboard

package main

import "fmt"

type Bytes int
type Celcius float32

type BandwidthUsage struct {
	amount []Bytes
}

type CpuTemp struct {
	temp []Celcius
}

type MemoryUsage struct {
	amount []Bytes
}

type Dashboard struct {
	BandwidthUsage
	CpuTemp
	MemoryUsage
}

func (b BandwidthUsage) average() Bytes {
	var total Bytes
	for _, amount := range b.amount {
		total += amount
	}
	return total / Bytes(len(b.amount))
}

func (c CpuTemp) average() Celcius {
	var total Celcius
	for _, temp := range c.temp {
		total += temp
	}
	return total / Celcius(len(c.temp))
}

func (m MemoryUsage) average() Bytes {
	var total Bytes
	for _, amount := range m.amount {
		total += amount
	}
	return total / Bytes(len(m.amount))
}

func (d Dashboard) BandwidthAverage() Bytes {
	return d.BandwidthUsage.average()
}

func (d Dashboard) CpuTempAverage() Bytes {
	return d.BandwidthUsage.average()
}

func (d Dashboard) MemoryAverage() Bytes {
	return d.BandwidthUsage.average()
}

func main() {
	bandwidth := BandwidthUsage{[]Bytes{50000, 100000, 130000, 80000, 90000}}
	temp := CpuTemp{[]Celcius{50, 51, 53, 51, 52}}
	memory := MemoryUsage{[]Bytes{800000, 800000, 810000, 820000, 800000}}

	dash := Dashboard{bandwidth, temp, memory}
	fmt.Printf("5-second Bandwidth Average: %d\n", dash.BandwidthAverage())
	fmt.Printf("5-second CPU Temp Average: %f\n", dash.CpuTempAverage())
	fmt.Printf("5-second Memory Average: %d\n", dash.MemoryAverage())
}
