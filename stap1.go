package main

import (
	"fmt"
	"time"
)

// BepaalGroet bepaalt het juiste groetbericht op basis van het huidige dagdeel
func BepaalGroet() string {
	huidigeTijd := time.Now()
	uren := huidigeTijd.Hour()

	var groet string

	switch {
	case uren >= 7 && uren < 12:
		groet = "Goedemorgen"
	case uren >= 12 && uren < 18:
		groet = "Goedemiddag"
	case uren >= 18 && uren < 23:
		groet = "Goedenavond"
	default:
		groet = "Sorry, de parkeerplaats is 's nachts gesloten"
	}

	return groet
}

// haalt bericht op basis van tijd
func WelkomstBericht() string {
	groet := BepaalGroet()
	return fmt.Sprintf("%s! Welkom bij Fonteyn Vakantieparken", groet)
}

// haalt welcoms bericht op
func main() {
	welkomstBericht := WelkomstBericht()
	fmt.Println(welkomstBericht)
}
