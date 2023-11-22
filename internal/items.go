package internal

import "module/internal/domain"

var (
	HookAttack              = domain.GetAttack("hook")
	RootAttack              = domain.GetAttack("root")
	UltimateDismemberAttack = domain.GetAttack("ultimateDismember")
)

var (
	MeteorCast    = domain.GetSpell("meteor")
	BeemLaserCast = domain.GetSpell("beemlaser")
)

var (
	Buff = &domain.Buff{}

	SpeedBuff = &domain.SpeedBuff{
		Buff: Buff,
	}

	StrengthAndSpeedBuff = &domain.StrengthBuff{
		Buff: SpeedBuff,
	}
)
