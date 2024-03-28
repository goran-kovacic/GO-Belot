package main

import (
	"bela/modelStruct"
	"fmt"
)

var (
	partije                        []modelStruct.PartijaSucelje
	igrac1, igrac2, igrac3, igrac4 modelStruct.Igrac
)

func main() {
	kreirajRucno()

	for _, p := range partije {
		fmt.Println(p)
	}
}

func kreirajRucno() {
	partije = []modelStruct.PartijaSucelje{}

	igrac1 = kreirajIgraca("Tomislav", "Jakopec")
	igrac2 = kreirajIgraca("Marijan", "Zidar")
	igrac3 = kreirajIgraca("Marija", "Zimska")
	igrac4 = modelStruct.Igrac{
		Ime: "Anita",
		Prezime: "Račman",
		// Entitet: modelStruct.Entitet{ID: 1},
	}
	partijaDvaIgraca()
	partijaTriIgraca()
	partijaDvaPara()
}

func partijaDvaPara(){
	partija := modelStruct.PartijaDvaPara{}
	partija.DoKolikoSeIgra = 501
	partija.Unosi = igrac1
	partija.Igraci = []modelStruct.Igrac{igrac1, igrac2, igrac3, igrac4}
	partija.Mjesanja = mjesanjeDvaPara()
	partije = append(partije, partija)
}

func partijaTriIgraca(){
	partija := modelStruct.PartijaTriIgraca{}
	partija.DoKolikoSeIgra = 501
	partija.Unosi = igrac1
	partija.Igraci = []modelStruct.Igrac{igrac1, igrac2, igrac3}
	// partija.Igraci = append(partija.Igraci, igrac1)
	partija.Mjesanja = mjesanjeTriIgraca()
	partije = append(partije, partija)
}

func mjesanjeTriIgraca() []modelStruct.Mjesanje{
	mjesanja := []modelStruct.Mjesanje{}

	mjesanje1 := modelStruct.MjesanjeTriUnosa{
		MjesanjeDvaUnosa: modelStruct.MjesanjeDvaUnosa{
			BodovaPrviUnos: 10,
			BodovaDrugiUnos: 76,
			ZvanjePrviUnos: 0,
			ZvanjeDrugiUnos: 20,
		},
		BodovaTreciUnos: 76,
	}
	mjesanja = append(mjesanja, mjesanje1)

	for i:= 0; i<5; i++{
		mjesanja = append(mjesanja, mjesanje1)
	}

	return mjesanja
}

func partijaDvaIgraca() {
	partija := modelStruct.PartijaDvaIgraca{}
	partija.DoKolikoSeIgra = 501
	partija.Unosi = igrac1
	partija.Igraci = []modelStruct.Igrac{igrac1, igrac2}
	partija.Mjesanja = mjesanjeDvaPara()
	partije = append(partije, partija)
}

func mjesanjeDvaPara() []modelStruct.Mjesanje {
	// mjesanja := make([]modelStruct.Mjesanje, 0)

	mjesanja := []modelStruct.Mjesanje{}

	mjesanje1 := modelStruct.MjesanjeDvaUnosa{
		BodovaPrviUnos:  10,
		BodovaDrugiUnos: 152,
		ZvanjePrviUnos:  0,
		ZvanjeDrugiUnos: 20,
		Stiglja:         false,
		Belot:           false,
		Entitet:         modelStruct.Entitet{Id: 5},
	}
	mjesanja = append(mjesanja, mjesanje1)

	mjesanje2 := modelStruct.MjesanjeDvaUnosa{
		BodovaPrviUnos:  152,
		BodovaDrugiUnos: 10,
		ZvanjePrviUnos:  0,
		ZvanjeDrugiUnos: 20,
	}
	mjesanja = append(mjesanja, mjesanje2)

	// fmt.Println(mjesanja)

	return mjesanja
}

func kreirajIgraca(ime, prezime string) modelStruct.Igrac {
	igrac := modelStruct.Igrac{
		Ime:     ime,
		Prezime: prezime,
	}
	return igrac
}
