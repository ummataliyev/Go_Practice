package main

import (
	"fmt"
	"math"
)

func main() {
	// "Nexia", "Tico", 'Damas' ko'rganlar qilar havas
	fmt.Println("\"Nexia\", \"Tico\", 'Damas' ko'rganlar qilar havas")

	// 5 ning 4-darajasi
	fmt.Println("5 ning 4-darajasi", int(math.Pow(5, 4)))

	// 22 ni 4 ga bo'lganda qancha qoldiq qoladi?
	fmt.Println("22 ni 4 ga bo'lganda qancha qoldiq", 22%4)

	// Tomonlari 125 ga teng kvadratning yuzi va perimetrini toping
	side := 125
	fmt.Println("Tomonlari 125 ga teng kvadratning yuzi", side*side, "ga, perimetri", 4*side, "teng")

	//  Diametri 12 ga teng bo'lgan doiraning yuzini toping
	radius := 12.0 / 2
	fmt.Println("Diametri 12 ga teng bo'lgan doiraning yuzi", math.Pi*radius*radius, "ga teng")

	//  Katetlari 6 va 7 bo'lgan to'g'ri burchakli uchburchakning gipotenuzasini toping
	katet_a, katet2_b := 6.0, 7.0
	gipotenuza := math.Sqrt(math.Pow(katet_a, 2) + math.Pow(katet2_b, 2))
	fmt.Println("Katetlari 6 va 7 bo'lgan to'g'ri burchakli uchburchakning gipotenuzasi", gipotenuza)
}
