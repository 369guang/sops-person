package monitor

import "fmt"

type memoryCollection struct {
}

func init() {
	registerCollector("memories", NewMemoryCollector)
}

func NewMemoryCollector() (Collector, error) {
	return &memoryCollection{}, nil
}

func (c *memoryCollection) Update(ch chan<- Metric) error {

	fmt.Println("update memory collector")

	return nil
}
