package sighting

// RareCutOff is the largest number of sightings a species may have in
// the record and still be labelled rarely recorded.
const RareCutOff = 2

// RareSpecies returns the GBIF keys of the species seen at most
// RareCutOff times, so a page can badge their sightings as rarely
// recorded. The whole record is counted, not just what a page shows.
func RareSpecies(sightings []Sighting) map[int]bool {
	counts := make(map[int]int)
	for _, s := range sightings {
		counts[s.Species.GBIFKey]++
	}

	rare := make(map[int]bool)
	for key, n := range counts {
		if n <= RareCutOff {
			rare[key] = true
		}
	}
	return rare
}
