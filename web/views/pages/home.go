package pages

import (
	"fmt"
	"strconv"

	"workshop/internal/domain/sighting"
)

// heroLede says what the site holds, in the figures it currently has.
func heroLede(s sighting.Summary) string {
	if s.Sightings == 0 {
		return "An open record of where the world's wildlife is seen, gathered so researchers, rangers and curious people can track species over time."
	}
	return fmt.Sprintf(
		"%s open wildlife records from %s countries, gathered in one place so researchers, rangers and curious people can see where species are being seen, and where they're not.",
		number(s.Sightings), number(s.Countries),
	)
}

// number formats n with thousands separators.
func number(n int) string {
	s := strconv.Itoa(n)
	for i := len(s) - 3; i > 0; i -= 3 {
		s = s[:i] + "," + s[i:]
	}
	return s
}

// yearSpan renders the first and last years on record, or a dash when there are none.
func yearSpan(s sighting.Summary) string {
	if s.Earliest.IsZero() {
		return "–"
	}
	if s.Earliest.Year() == s.Latest.Year() {
		return strconv.Itoa(s.Latest.Year())
	}
	return fmt.Sprintf("%d–%d", s.Earliest.Year(), s.Latest.Year())
}
