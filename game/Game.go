package game

type Game struct {
	Player1          Player `json:"player1"`
	Player2          Player `json:"player2"`
	Ball             int    `json:"ball"`
	GameMode         int    `json:"mode"`
	Winner           int    `json:"winner"`
	totalPoints      int
	totalPointsOld   int
	finish           bool
	winningCondition int
	final            bool
}

func (g *Game) Reset() {
	g.Player1.Reset()
	g.Player2.Reset()
	g.finish = false
	g.final = false
	g.Winner = 0
	g.Ball = 1
	g.totalPoints = 0
	g.totalPoints = g.Player1.Points + g.Player2.Points
}

func (g Game) GetGame() Game {
	return g
}

func (g *Game) SetGameMode(mode int) {
	g.GameMode = mode
	g.winningCondition = mode
}

func (g *Game) SwitchMode() {
	if g.GameMode == 21 {
		g.SetGameMode(11)
	} else {
		g.SetGameMode(21)
	}
}

func (g *Game) IncPoints(player int) {
	if g.finish == true {
		return
	}

	if player == 1 {
		g.Player1.IncPoints()
	} else {
		g.Player2.IncPoints()
	}
	g.totalPoints = g.Player1.Points + g.Player2.Points
	g.CheckFinish()
	g.DetectBallChange()
}

func (g *Game) CalcWinningCondition() int {
	breakPoints := g.GameMode - 1
	if g.Player1.Points >= breakPoints && g.Player2.Points >= breakPoints && g.Player1.Diff(g.Player2.Points) <= 2 {
		g.final = true
		if g.Player1.Points > g.Player2.Points {
			return g.Player2.Points + 2
		} else {
			return g.Player1.Points + 2
		}
	} else {
		return g.GameMode
	}
}

func (g *Game) CheckFinish() {
	g.winningCondition = g.CalcWinningCondition()

	if g.Player1.Points >= g.winningCondition {
		g.Winner = 1
		g.finish = true
	} else if g.Player2.Points >= g.winningCondition {
		g.Winner = 2
		g.finish = true
	} else {
		g.Winner = 0
		g.finish = false
	}
}

func (g *Game) DecPoints(player int) {
	if g.finish == true {
		return
	}

	if player == 1 {
		g.Player1.DecPoints()
	} else {
		g.Player2.DecPoints()
	}
	g.totalPoints = g.Player1.Points + g.Player2.Points
	g.DetectBallChange()
}

func (g *Game) ChangeBall() {
	if g.Ball == 1 {
		g.Ball = 2
	} else {
		g.Ball = 1
	}
}

func (g *Game) DetectBallChange() {
	if g.totalPointsOld == g.totalPoints {
		return
	}
	g.totalPointsOld = g.totalPoints

	if g.final {
		g.ChangeBall()
	} else {
		if g.GameMode == 11 {
			if g.totalPoints%3 == 0 {
				g.ChangeBall()
			}
		} else {
			if g.totalPoints%5 == 0 {
				g.ChangeBall()
			}
		}
	}

}

func (g *Game) Init(player1 Player, player2 Player) {
	g.Player1 = player1
	g.Player2 = player2
	g.GameMode = 21
	g.Ball = 1
	g.Reset()
}

func NewGame() *Game {
	g := Game{GameMode: 21, Ball: 1}
	player1 := NewPlayer("1", g)
	player2 := NewPlayer("2", g)
	g.Init(player1, player2)
	return &g
}
