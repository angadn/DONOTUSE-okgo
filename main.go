package okgo

// OKGO holds our execution blocks.
type OKGO struct {
	blocks []func() bool
}

// NewOKGO constructs an OKGO for us.
func NewOKGO() OKGO {
	return OKGO{
		blocks: make([]func() bool, 0),
	}
}

// On chains execution blocks.
func (o OKGO) On(block func() bool) OKGO {
	o.blocks = append(o.blocks, block)
	return o
}

// Run starts the execution chain.
func (o OKGO) Run() {
	for _, block := range o.blocks {
		if !block() {
			break
		}
	}
}
