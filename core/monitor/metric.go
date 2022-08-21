package monitor

type Pair struct {
	Name  *string `json:"name"`
	Value *string `json:"value"`
}

func (p *Pair) GetName() string {
	if p != nil && p.Name != nil {
		return *p.Name
	}
	return ""
}

func (p *Pair) GetValue() string {
	if p != nil && p.Value != nil {
		return *p.Value
	}
	return ""
}

type Metric struct {
	// 命名空间
	fqName string
	// 名称
	pairs []*Pair
	// 值
	variable []string
	// error
	err error
}
