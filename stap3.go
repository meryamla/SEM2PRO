package main

import (
	"fmt"
	"time"
)

// BekendeKentekens bevat een hard-coded reeks van bekende kentekens
var BekendeKentekens = []string{"ABC123", "XYZ789", "123ABC"}

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

// WelkomstBericht genereert het welkomstbericht op basis van het huidige dagdeel
func WelkomstBericht() string {
	groet := BepaalGroet()
	return fmt.Sprintf("%s! Welkom bij Fonteyn Vakantieparken", groet)
}

// ControleerKenteken controleert of het opgegeven kenteken bekend is
func ControleerKenteken(opgegevenKenteken string) string {
	for _, bekendKenteken := range BekendeKentekens {
		if bekendKenteken == opgegevenKenteken {
			return WelkomstBericht()
		}
	}

	return "U heeft helaas geen toegang tot het parkeerterrein"
}

func main() {
	// Vraag de gebruiker om het kenteken in te voeren
	var opgegevenKenteken string
	fmt.Print("Voer uw kenteken in: ")
	fmt.Scanln(&opgegevenKenteken)

	// Controleer of het kenteken bekend is
	resultaat := ControleerKenteken(opgegevenKenteken)

	// Toon het resultaat
	fmt.Println(resultaat)
}
