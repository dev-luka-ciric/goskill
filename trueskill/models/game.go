package models

import (
	"fmt"
	"math"
	"slices"

	"github.com/dev-luka-ciric/goskill/trueskill"
)

type Game struct {
	Team1 *Team
	Team2 *Team
}

func NewGame(team1, team2 *Team) Game {
	return Game{
		Team1: team1,
		Team2: team2,
	}
}

func (g *Game) MatchQuality() float64 {
	team1 := g.Team1.CalculateTeam()
	team2 := g.Team2.CalculateTeam()

	return ts.MatchQuality([]trueskill.Rating{team1, team2})
}

func calculateNewSkillsForPlayer(player *Player, team, skill trueskill.Rating) {
	fmt.Println(player.Rating.Mu(), player.Rating.Sigma(), team.Mu(), team.Sigma(), skill.Mu(), skill.Sigma())

	d := math.Pow(player.Rating.Sigma(), 2) / math.Pow(team.Sigma(), 2)
	mu := d * skill.Mu()
	sigma := d * math.Pow(skill.Sigma(), 2)
	sigma = math.Sqrt(sigma)
	player.Rating = trueskill.NewRating(mu, sigma)
}

func (g *Game) UpdateRatings(firstWon bool) {
	team1 := g.Team1.CalculateTeam()
	team2 := g.Team2.CalculateTeam()

	teams := []trueskill.Rating{team1, team2}
	if !firstWon {
		slices.Reverse(teams)
	}

	skills, _ := ts.AdjustSkills(teams, false)

	if !firstWon {
		slices.Reverse(skills)
	}

	for _, p := range g.Team1.GetPlayers() {
		calculateNewSkillsForPlayer(p, team1, skills[0])
	}

	for _, p := range g.Team2.GetPlayers() {
		calculateNewSkillsForPlayer(p, team2, skills[1])
	}
}
