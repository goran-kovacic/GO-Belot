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
	kreirajPartijuDvaIgracaSTRUCT()
}

// func kreirajMjesanjaDvaPara() []model.MjesanjeDvaUnosa {
// 	// mjesanja := make([]model.Mjesanje, 0)

// 	mjesanja := []model.MjesanjeDvaUnosa{}

// 	mjesanje1 := model.MjesanjeDvaUnosa{
// 		BodovaPrviUnos:  10,
// 		BodovaDrugiUnos: 152,
// 		ZvanjePrviUnos:  0,
// 		ZvanjeDrugiUnos: 20,
// 	}

// 	mjesanja = append(mjesanja, mjesanje1)

// 	mjesanje2 := model.MjesanjeDvaUnosa{
// 		BodovaPrviUnos:  152,
// 		BodovaDrugiUnos: 10,
// 		ZvanjePrviUnos:  0,
// 		ZvanjeDrugiUnos: 20,
// 	}
// 	mjesanja = append(mjesanja, mjesanje2)

// 	fmt.Println(mjesanja)

// 	return mjesanja
// }

// func kreirajPartijuDvaIgraca() {
// 	partija := &model.PartijaDvaIgraca{}
// 	partija.DoKolikoSeIgra = 501
// 	// partija.Lokacija = *lokacija
// 	partija.Unosi = *igrac1
// 	partija.Igraci = []model.Igrac{*igrac1, *igrac2}
// 	// partija.MjesanjeDvaPara = kreirajMjesanjaDvaPara()
// 	partija.MjesanjeDvaPara= kreirajMjesanjaDvaPara()
// 	partije = append(partije, &partija.Partija)
// }

func kreirajPartijuDvaIgracaSTRUCT() {
	partija := modelStruct.PartijaDvaIgraca{}
	partija.DoKolikoSeIgra = 501
	partija.Unosi = igrac1
	partija.Igraci = []modelStruct.Igrac{igrac1, igrac2}
	// partija.MjesanjeDvaPara = kreirajMjesanjaDvaPara()
	// partija.MjesanjeDvaPara= kreirajMjesanjaDvaPara()
	partija.Mjesanja = kreirajMjesanjaDvaParaSTRUCT()
	partije = append(partije, partija)
}

func kreirajMjesanjaDvaParaSTRUCT() []modelStruct.Mjesanje {
	// mjesanja := make([]modelStruct.Mjesanje, 0)

	mjesanja := []modelStruct.Mjesanje{}

	mjesanje1 := modelStruct.MjesanjeDvaUnosa{
		BodovaPrviUnos:  10,
		BodovaDrugiUnos: 152,
		ZvanjePrviUnos:  0,
		ZvanjeDrugiUnos: 20,
		Stiglja:         false,
		Belot:           false,
		Entitet:         modelStruct.Entitet{ID: 5},
	}

	mjesanja = append(mjesanja, mjesanje1)

	mjesanje2 := modelStruct.MjesanjeDvaUnosa{
		BodovaPrviUnos:  152,
		BodovaDrugiUnos: 10,
		ZvanjePrviUnos:  0,
		ZvanjeDrugiUnos: 20,
	}
	mjesanja = append(mjesanja, mjesanje2)

	fmt.Println(mjesanja)

	return mjesanja
}

func kreirajIgraca(ime, prezime string) modelStruct.Igrac {
	igrac := modelStruct.Igrac{
		Ime:     ime,
		Prezime: prezime,
	}
	return igrac
}
