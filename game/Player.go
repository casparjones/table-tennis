package game

type Player struct {
	Name        string `json:"name"`
	defaultName string
	playerNo    string
	Points      int `json:"points"`
}

func (p *Player) SetName(name string) {
	if name != "" {
		p.Name = name
	} else {
		p.SetDefaultName()
	}
}

func (p *Player) IncPoints() {
	p.Points++
}

func (p *Player) DecPoints() {
	if p.Points > 0 {
		p.Points--
	}
}

func (p *Player) SetPoints(point int) {
	if point > -1 {
		p.Points = point
	}
}

func (p *Player) SetDefaultName() {
	p.SetName(p.defaultName + " " + p.playerNo)
}

func (p *Player) Diff(points int) int {
	if points > p.Points {
		return points - p.Points
	} else {
		return p.Points - points
	}
}

func (p *Player) Reset() *Player {
	p.SetPoints(0)
	// p.SetDefaultName()

	return p
}

func NewPlayer(playerNo string, g Game) Player {
	p := Player{
		defaultName: "Player",
		playerNo:    playerNo,
	}
	p.Reset().SetDefaultName()
	return p
}
