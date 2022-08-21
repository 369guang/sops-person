package monitor

import "fmt"

type diskCollection struct {
}

func init() {
	registerCollector("disk", NewDiskCollector)
}

func NewDiskCollector() (Collector, error) {
	return &diskCollection{}, nil
}

func (c *diskCollection) Update(ch chan<- Metric) error {

	fmt.Println("update disk collection")

	return nil
}
