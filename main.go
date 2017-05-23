package okgo

// OKGO holds our execution blocks.
type OKGO struct {
	err    *error
	blocks []func() error
}

// NewOKGO constructs an OKGO for us.
func NewOKGO(err *error) OKGO {
	return OKGO{
		err:    err,
		blocks: make([]func() error, 0),
	}
}

// On chains execution blocks.
func (o OKGO) On(block func() error) OKGO {
	o.blocks = append(o.blocks, block)
	return o
}

// Run starts the execution chain.
func (o OKGO) Run() {
	for _, block := range o.blocks {
		if err := block(); err != nil {
			*o.err = err
			break
		}
	}
}
