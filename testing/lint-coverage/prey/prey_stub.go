package prey

import "testdoubles/positioner"

func NewPreyStub() (prey *PreyStub) {
	prey = &PreyStub{}
	return
}

type PreyStub struct {
	GetSpeedFunc    func() (speed float64)
	GetPositionFunc func() (position *positioner.Position)
}

func (p *PreyStub) GetSpeed() (speed float64) {
	if p.GetSpeedFunc != nil {
		return p.GetSpeedFunc()
	}
	return
}

func (p *PreyStub) GetPosition() (position *positioner.Position) {
	if p.GetPositionFunc != nil {
		return p.GetPositionFunc()
	}
	return
}
