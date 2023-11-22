package domain

type IBuff interface {
	GetDescription() string
}

type Buff struct{}

func (b *Buff) GetDescription() string {
	return "BUFFS:\n"
}

type SpeedBuff struct {
	Buff IBuff
}

func (s *SpeedBuff) GetDescription() string {
	buffDescription := s.Buff.GetDescription()
	return buffDescription + "25% move speed\n"
}

type StrengthBuff struct {
	Buff IBuff
}

func (s *StrengthBuff) GetDescription() string {
	buffDescription := s.Buff.GetDescription()
	return buffDescription + "10% hp, 10% damage\n"
}
