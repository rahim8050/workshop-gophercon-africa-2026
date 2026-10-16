package components

import (
	"fmt"
	"strconv"

	"workshop/internal/domain/sighting"
)

// years lists every year from the first sighting to the last, so a quiet year stays a gap.
func years(sum sighting.Summary) []int {
	ys := make([]int, 0, sum.Latest.Year()-sum.Earliest.Year()+1)
	for y := sum.Earliest.Year(); y <= sum.Latest.Year(); y++ {
		ys = append(ys, y)
	}
	return ys
}

// peak finds the busiest year and its count.
func peak(byYear map[int]int) (year, count int) {
	for y, n := range byYear {
		if n > count || (n == count && y < year) {
			year, count = y, n
		}
	}
	return year, count
}

// busiest is the highest count in any one year.
func busiest(byYear map[int]int) int {
	_, n := peak(byYear)
	return n
}

// bar sets --bar to a year's share of the busiest, for the CSS to draw.
func bar(count, longest int) map[string]string {
	if longest <= 0 {
		return map[string]string{"--bar": "0%"}
	}
	return map[string]string{"--bar": strconv.Itoa(count*100/longest) + "%"}
}

// yearTitle is the tooltip on one year's bar.
func yearTitle(year, n int) string {
	if n == 1 {
		return fmt.Sprintf("%d: 1 sighting", year)
	}
	return fmt.Sprintf("%d: %d sightings", year, n)
}

// timelineCaption names the busiest year under the bars.
func timelineCaption(byYear map[int]int) string {
	y, n := peak(byYear)
	return fmt.Sprintf("Sightings per year · busiest %d, with %d", y, n)
}

// timelineLabel says in words what the bars show.
func timelineLabel(byYear map[int]int, sum sighting.Summary) string {
	y, n := peak(byYear)
	return fmt.Sprintf("Sightings per year from %d to %d, most in %d with %d",
		sum.Earliest.Year(), sum.Latest.Year(), y, n)
}
