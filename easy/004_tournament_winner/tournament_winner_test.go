package tournament_winner

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSortedSquaredArray(t *testing.T) {
	tests := []struct {
		name         string
		competitions [][]string
		results      []int
		want         string
	}{
		{
			name:         "1",
			competitions: [][]string{{"HTML", "C#"}, {"C#", "Python"}, {"Python", "HTML"}},
			results:      []int{0, 0, 1},
			want:         "Python",
		},
		{
			name: "2",
			competitions: [][]string{
				{"HTML", "Java"},
				{"Java", "Python"},
				{"Python", "HTML"},
				{"C#", "Python"},
				{"Java", "C#"},
				{"C#", "HTML"},
				{"SQL", "C#"},
				{"HTML", "SQL"},
				{"SQL", "Python"},
				{"SQL", "Java"}},
			results: []int{0, 1, 1, 1, 0, 1, 0, 1, 1, 0},
			want:    "C#",
		},
		{
			name: "3",
			competitions: [][]string{
				{"HTML", "Java"},
				{"Java", "Python"},
				{"Python", "HTML"}},
			results: []int{0, 1, 1},
			want:    "Java",
		},
	}

	var got string

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got = TournamentWinner(tt.competitions, tt.results)
			require.Equal(t, got, tt.want)
		})
	}
}
