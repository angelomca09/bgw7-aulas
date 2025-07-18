package simulator

func NewSimulatorMock() (simulator *SimulatorMock) {
	simulator = &SimulatorMock{}
	return
}

type SimulatorMock struct {
	CanCatchFunc func(hunter, prey *Subject) (canCatch bool)

	Calls struct {
		CanCatch int
	}
}

func (s *SimulatorMock) CanCatch(hunter, prey *Subject) (canCatch bool) {
	s.Calls.CanCatch++
	if s.CanCatchFunc != nil {
		return s.CanCatchFunc(hunter, prey)
	}
	return
}
