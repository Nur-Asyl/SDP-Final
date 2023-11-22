package internal

type Shrine interface {
	Register(Hero)
	Deregister(Hero)
	NotifyAll(float64)
}

type HPShrine struct {
	BarManager ShrineManager
	heroes     []Hero
}

func (m *HPShrine) Register(h Hero) {
	m.heroes = append(m.heroes, h)
}

func (m *HPShrine) Deregister(h Hero) {
	m.heroes = removeFromslice(m.heroes, h)
}

func (m HPShrine) NotifyAll() {
	for _, h := range m.heroes {
		h.update(m.BarManager.GetState())
	}
}

func removeFromslice(observerList []Hero, h Hero) []Hero {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if h.GetName() == observer.GetName() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}
