package h_tool

type options struct {
	a int64
	b string
	c map[string]string
}

type ServerOption func(*options)

func NewOption(opt ...ServerOption) *options  {
	r := new(options)
	for _, o := range opt {
		o(r)
	}
	return r
}

func WriteA(s int64) ServerOption {
	return func(o *options) {
		o.a = s
	}
}

func WriteB(s string) ServerOption {
	return func(o *options) {
		o.b = s
	}
}

func WriteC(s map[string]string) ServerOption {
	return func(o *options) {
		o.c = s
	}
}
