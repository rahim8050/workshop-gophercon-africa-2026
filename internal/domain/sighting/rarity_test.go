package sighting

import (
	"testing"
	"workshop/internal/domain/species"

	"github.com/stretchr/testify/assert"
)

func TestRareSpecies(t *testing.T) {
	t.Parallel()

	zebra := Sighting{Species: species.Species{GBIFKey: 2441176}}
	lion := Sighting{Species: species.Species{GBIFKey: 5219317}}

	tt := map[string]struct {
		input []Sighting
		want  map[int]bool
	}{
		"No sightings": {want: map[int]bool{}},
		"Seen twice or fewer": {
			input: []Sighting{zebra, zebra, lion},
			want:  map[int]bool{2441176: true, 5219317: true},
		},
		"Seen more often": {
			input: []Sighting{zebra, zebra, zebra, lion},
			want:  map[int]bool{5219317: true},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tc.want, RareSpecies(tc.input))
		})
	}
}
