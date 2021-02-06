package api

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

func NewGame(todos []Todo, rounds int, extraRounds int, timePerRound int) Game {

	return &game{
		todos:        todos,
		rounds:       rounds,
		extraRounds:  extraRounds,
		timePerRound: timePerRound,
	}
}

type Todo struct {
	Name string
}

func (t Todo) String() string {
	return t.Name
}

type Game interface {
	Start(listener chan Info)
}

type Score struct {
	Todo  Todo
	Score int
}

func (s Score) String() string {
	return fmt.Sprintf("[%v]: %v", s.Todo, s.Score)
}

type Info struct {
	IsEnded bool
	Round   int
	Total   int
	Scores  []Score
}

func (i Info) String() string {
	sb := strings.Builder{}

	sb.WriteString(fmt.Sprintf(" Round: %2d of %d\n\n", i.Round, i.Total))

	for i, score := range i.Scores {
		sb.WriteString(fmt.Sprintf("%2dº %2v\n", i+1, score))
	}

	if i.IsEnded {
		if len(i.Scores) > 0 {
			sb.WriteString("\n")
			sb.WriteString(fmt.Sprintf("[%s] is the Winner!", i.Scores[0].Todo))
		}
	}
	return sb.String()
}

type game struct {
	rounds       int
	timePerRound int
	extraRounds  int
	todos        []Todo
	randSource   rand.Source
}

func (g *game) Start(listener chan Info) {
	rounds := g.rounds
	scores := make([]Score, 0)

	g.randSource = rand.NewSource(time.Now().UTC().UnixNano())

	for round := 1; round <= rounds; round++ {
		roundWinner := draw(g.todos, g.randSource)
		scores = updateAndNotify(listener, scores, roundWinner, round, rounds)
	}
}

func updateAndNotify(listener chan Info, scores []Score, roundWinner Todo, round int, rounds int) []Score {
	var info Info
	scores, info = updateScore(scores, roundWinner, round, rounds)
	listener <- info
	return scores
}

func updateScore(scores []Score, winner Todo, round int, rounds int) ([]Score, Info) {
	index, found := search(scores, winner)
	if found {
		scores[index].Score++
	} else {
		scores = append(scores, Score{
			Todo:  winner,
			Score: 1,
		})
	}

	sort.Slice(scores, func(i, j int) bool {
		return scores[i].Score > scores[j].Score
	})
	return scores, Info{
		IsEnded: round == rounds,
		Round:   round,
		Total:   rounds,
		Scores:  scores,
	}
}

func search(scores []Score, winner Todo) (int, bool) {
	for i, s := range scores {
		if s.Todo == winner {
			return i, true
		}
	}
	return -1, false
}

func draw(todos []Todo, source rand.Source) Todo {
	n := len(todos)
	drawn := rand.New(source).Intn(n)
	return todos[drawn]
}
