package models

import (
	"math"

	"goskill/trueskill"
)

type Team struct {
	Players []*Player
}

func NewDefaultTeam() Team {
	return Team{
		Players: []*Player{},
	}
}

func NewTeam(players ...*Player) Team {
	return Team{
		Players: players,
	}
}

func (t *Team) AddPlayer(p *Player) {
	t.Players = append(t.Players, p)
}

func (t *Team) GetPlayers() []*Player {
	return t.Players
}

func (t *Team) CalculateTeam() trueskill.Rating {
	mu, sigma := 0.0, 0.0
	for _, p := range t.Players {
		mu += p.Rating.Mu()
		sigma += math.Pow(p.Rating.Sigma(), 2)
	}

	return trueskill.NewRating(mu, math.Sqrt(sigma))
}
