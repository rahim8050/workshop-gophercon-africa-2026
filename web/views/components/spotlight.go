package components

import (
	"fmt"
	"time"

	"workshop/internal/domain/sighting"
)

// PickSpotlight chooses one sighting with a photo for the day, the same one all day for everyone.
func PickSpotlight(sightings []sighting.Sighting, day time.Time) (sighting.Sighting, sighting.Media, bool) {
	var withPhoto []int
	for i, s := range sightings {
		if len(s.Media) > 0 && s.Media[0].Type == "StillImage" {
			withPhoto = append(withPhoto, i)
		}
	}
	if len(withPhoto) == 0 {
		return sighting.Sighting{}, sighting.Media{}, false
	}
	s := sightings[withPhoto[(day.Year()*366+day.YearDay())%len(withPhoto)]]
	return s, s.Media[0], true
}

// onRecordText says how many sightings of the species are on record.
func onRecordText(n int) string {
	if n <= 1 {
		return "It is the only sighting of this species on record."
	}
	return fmt.Sprintf("There are %d sightings of this species on record.", n)
}
