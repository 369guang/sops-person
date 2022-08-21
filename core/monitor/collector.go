package monitor

type Collector interface {
	Update(ch chan<- Metric) error
}

var (
	// 工厂注册信息
	factories = make(map[string]func() (Collector, error))
	// 记录状态
	collectorState = make(map[string]bool)
)

// 注册
func registerCollector(collector string, factory func() (Collector, error)) {
	collectorState[collector] = true
	factories[collector] = factory
}

type NodeCollector struct {
	Collectors map[string]Collector
}

func NewNodeCollector() (*NodeCollector, error) {

	collectors := make(map[string]Collector)
	for key, _ := range collectorState {
		collector, err := factories[key]()
		if err != nil {
			return nil, err
		}
		collectors[key] = collector
	}

	return &NodeCollector{Collectors: collectors}, nil
}

//nc, err := collector.NewNodeCollector()
//if err != nil {
//panic(err)
//}
//var collectors []string
//for n := range nc.Collectors {
//collectors = append(collectors, n)
//}
//sort.Strings(collectors)
//for _, c := range collectors {
//fmt.Println("collector", c)
//}
//
//fmt.Println("更新阶段")
