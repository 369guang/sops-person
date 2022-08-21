package monitor

import "fmt"

type cpuCollection struct {
}

func init() {
	registerCollector("cpu", NewCPUCollector)
}

func NewCPUCollector() (Collector, error) {
	return &cpuCollection{}, nil
}

func (c *cpuCollection) Update(ch chan<- Metric) error {

	fmt.Println("update cpu collection")

	return nil
}
