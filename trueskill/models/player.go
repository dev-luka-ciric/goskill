package models

import (
	"github.com/dev-luka-ciric/goskill/trueskill"
)

type Player struct {
	SteamId string
	trueskill.Rating
}

func NewDefaultPlayer(steamId string) Player {
	return Player{
		SteamId: steamId,
		Rating:  ts.NewRating(),
	}
}

func NewPlayer(steamId string, mu, sigma float64) Player {
	return Player{
		SteamId: steamId,
		Rating:  trueskill.NewRating(mu, sigma),
	}
}
