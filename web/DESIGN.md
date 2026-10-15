# Design

Sightings Observatory: an open record of wildlife sightings. The feel is a **data observatory**:
crisp, factual and serious, with the figures up front.

## Colours (`:root` in `assets/css/app.css`)

| Property | Hex | Use |
|----------|-----|-----|
| `--ink` | `#14213D` | Header, hero, footer background; links on light |
| `--ink-raised` | `#1C2C4F` | Hover on dark |
| `--ink-line` | `rgba(255,255,255,.12)` | Dividers on dark |
| `--gold` | `#C8962E` | Buttons, eyebrow text, focus ring. Only on ink, never as text on light |
| `--gold-hover` | `#D9A843` | Button hover |
| `--paper` | `#FAFAF7` | Cards |
| `--page` | `#ECECEE` | Page background (chosen by the attendee: grey, not white) |
| `--text` | `#1C1C1E` | Body text |
| `--muted` | `#52525B` | Secondary text on light |
| `--on-ink` / `--on-ink-muted` | `#FAFAF7` / `#B9C2D6` | Text on dark |

## Lettering

Inter (Google Fonts) throughout: 700 for headings with tight letter-spacing, uppercase small
labels with wide tracking, tabular figures for numbers.

## Layout

- Content width `--wrap` 1200px, side padding `--gutter` (16–32px).
- Radii: 10px for cards, 8px for buttons, 6px for nav links.

## Components

- **Header**: sticky ink bar with the leaf-and-eye logo, the name and an "Open wildlife records"
  tagline, nav links (Coverage, Map, Records, About the data) and a gold "Browse records" button. On a
  phone the nav drops to its own row, and below 420px the tagline and button hide.
- **Hero**: ink banner with an eyebrow ("Open biodiversity data · Sourced from GBIF"), the headline
  "Every sighting is a data point for conservation.", a lede built from live figures, two buttons.
- **Stats strip**: five figures from `sighting.Summarise` (sightings, species, countries,
  observers, years covered) and the latest record date. Five columns, two on a phone.
- **Map** (`#map`): Leaflet on OpenStreetMap tiles, every sighting with a position (not only
  the 50 cards), clustered. Markers and cluster bubbles are gold `--map-marker` with an ink
  `--map-marker-stroke`; corners use `--radius`. 16:9 on desktop, 4:5 on a phone. Popups link to
  a card only when that sighting is among the latest shown.
- **Cards**: paper on the grey page, light shadow.
- **Footer** (`#about`): what the data is and where it comes from (GBIF), with a licence note.
- **Logo**: `assets/images/logo.svg`, a gold leaf-shaped eye on an ink rounded square.
