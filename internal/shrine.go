package internal

type ShrineManager interface {
	SetState(float64)
	GetState() float64
}

type HPShrineManager struct {
	hp float64
}

func (s *HPShrineManager) SetState(state float64) {
	s.hp = state
}

func (s *HPShrineManager) GetState() float64 {
	return s.hp
}
