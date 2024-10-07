package device

type ciscoDevice struct {
	PlaybookDIRSwitch string
	PlaybookDIRRouter string
}

func (dev *ciscoDevice) InitRouter() IRouter {
	return &ciscoRouter{
		PlaybookDIR: dev.PlaybookDIRRouter,
	}
}

func (dev *ciscoDevice) InitSwitch() ISwitch {
	return &ciscoSwitch{}
}
