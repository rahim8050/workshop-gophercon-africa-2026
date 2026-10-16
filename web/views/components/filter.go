package components

import (
	"slices"
	"strconv"
	"strings"

	"github.com/a-h/templ"

	"workshop/internal/domain/sighting"
)

// Animal names the animal seen, as the card does.
func Animal(s sighting.Sighting) string { return displayName(s.Species) }

// Country is where it was seen.
func Country(s sighting.Sighting) string { return s.Location.Country }

// Year is when, to the year.
func Year(s sighting.Sighting) string { return strconv.Itoa(s.HappenedAt.Year()) }

// Kind is the animal group's label, as the card's badge shows it.
func Kind(s sighting.Sighting) string { return AnimalGroup(s.Species).Label }

// FilterOptions counts how often each key appears, keeping the n most common; 0 keeps all.
func FilterOptions(sightings []sighting.Sighting, n int, key func(sighting.Sighting) string) []sighting.Tally {
	return sighting.Top(sighting.CountBy(sightings, key), n)
}

// Years lists the years the sightings fall in, oldest first, which is
// the order a reader counts along, not the order they were seen in.
func Years(sightings []sighting.Sighting) []sighting.Tally {
	counts := make(map[int]int)
	seen := make(map[int]bool)
	for _, s := range sightings {
		y := s.HappenedAt.Year()
		counts[y]++
		seen[y] = true
	}

	years := make([]int, 0, len(seen))
	for y := range seen {
		years = append(years, y)
	}
	slices.Sort(years)

	rows := make([]sighting.Tally, 0, len(years))
	for _, y := range years {
		rows = append(rows, sighting.Tally{Name: strconv.Itoa(y), Count: counts[y]})
	}
	return rows
}

// FilterAttrs marks a list item with what the filter matches on.
func FilterAttrs(s sighting.Sighting) templ.Attributes {
	return templ.Attributes{
		"data-kind":    Kind(s),
		"data-animal":  Animal(s),
		"data-country": Country(s),
		"data-year":    Year(s),
		"data-search": strings.ToLower(strings.Join([]string{
			Animal(s), s.Species.ScientificName, place(s.Location), s.RecordedBy,
		}, " ")),
	}
}
