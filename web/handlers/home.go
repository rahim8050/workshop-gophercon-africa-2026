package handlers

import (
	"context"
	"log/slog"
	"net/http"
	"time"
	"workshop/internal/domain/sighting"
	"workshop/web/views/components"
	"workshop/web/views/pages"
)

type SightingsLister interface {
	List(ctx context.Context, filter sighting.ListFilter) ([]sighting.Sighting, error)
}

// latestShown caps how many sightings the home page lists; the header's figures and the map cover them all.
const latestShown = 50

func Home(logger *slog.Logger, sightings SightingsLister) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		items, err := sightings.List(ctx, sighting.ListFilter{})
		if err != nil {
			logger.ErrorContext(ctx, "handlers: listing sightings: "+err.Error())
			renderError(ctx, w, logger, http.StatusInternalServerError, "Unable to load sightings.")
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		latest := items[:min(len(items), latestShown)]
		points := components.MapPoints(items, len(latest))

		home := pages.HomeProps{
			Sightings: latest,
			Summary:   sighting.Summarise(items),
			Points:    points,
			Rare:      sighting.RareSpecies(items),
			ByYear:    sighting.CountByYear(items),
			Spot:      spotlight(items),
		}
		if err = pages.Home(home).Render(ctx, w); err != nil {
			logger.ErrorContext(ctx, "handlers: rendering sightings: "+err.Error())
			renderError(ctx, w, logger, http.StatusInternalServerError, "Unable to render page.")
		}
	}
}

// spotlight picks the species of the day and counts how much of it is on record.
func spotlight(items []sighting.Sighting) pages.SpotlightProps {
	now := time.Now()
	s, photo, ok := components.PickSpotlight(items, now)
	props := pages.SpotlightProps{Sighting: s, Photo: photo, Day: now, Show: ok}
	if !ok {
		return props
	}
	for _, other := range items {
		if other.Species.GBIFKey == s.Species.GBIFKey {
			props.OnRecord++
		}
	}
	return props
}
