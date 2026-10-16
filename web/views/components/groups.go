package components

import "workshop/internal/domain/species"

// Group is a broad kind of animal, keyed for CSS (--group-<key>) and labelled for readers.
type Group struct {
	Key   string
	Label string
}

var (
	groupLand    = Group{"land", "Land mammal"}
	groupSea     = Group{"sea", "Sea mammal"}
	groupBird    = Group{"bird", "Bird"}
	groupReptile = Group{"reptile", "Reptile"}
	groupShark   = Group{"shark", "Shark or ray"}
	groupOther   = Group{"other", "Other"}
)

// Groups lists every group in the order a legend shows them.
var Groups = []Group{groupLand, groupSea, groupBird, groupReptile, groupShark, groupOther}

// seaFamilies are the mammal families that live in water: whales, dolphins, seals and sea cows.
var seaFamilies = map[string]bool{
	"Balaenopteridae": true, "Balaenidae": true, "Eschrichtiidae": true, "Neobalaenidae": true,
	"Delphinidae": true, "Phocoenidae": true, "Physeteridae": true, "Kogiidae": true,
	"Monodontidae": true, "Ziphiidae": true, "Iniidae": true, "Platanistidae": true, "Pontoporiidae": true,
	"Phocidae": true, "Otariidae": true, "Odobenidae": true, "Trichechidae": true, "Dugongidae": true,
}

// AnimalGroup sorts a species into a group from its taxonomy.
func AnimalGroup(sp species.Species) Group {
	switch sp.Class {
	case "Mammalia":
		if seaFamilies[sp.Family] {
			return groupSea
		}
		return groupLand
	case "Aves":
		return groupBird
	case "Reptilia", "Crocodylia", "Testudines", "Squamata":
		return groupReptile
	case "Elasmobranchii", "Chondrichthyes":
		return groupShark
	default:
		return groupOther
	}
}
