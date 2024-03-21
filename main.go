package main

import (
	"bela/model"
	"fmt"
)

var (
	partije                        []*model.Partija
	igrac1, igrac2, igrac3, igrac4 *model.Igrac
	lokacija                       *model.Lokacija
)

func main() {
	kreirajRucno()

	for _, p := range partije {
		fmt.Println(p)
	}

}

func kreirajRucno() {
	partije = []*model.Partija{}

	igrac1 = kreirajIgraca("Tomislav", "Jakopec")
	igrac2 = kreirajIgraca("Marijan", "Zidar")

	kreirajPartijuDvaIgraca()
}

func kreirajMjesanjaDvaPara() []model.MjesanjeDvaUnosa {
	// mjesanja := make([]model.Mjesanje, 0)

	mjesanja := []model.MjesanjeDvaUnosa{}

	mjesanje1 := model.MjesanjeDvaUnosa{
		BodovaPrviUnos:  10,
		BodovaDrugiUnos: 152,
		ZvanjePrviUnos:  0,
		ZvanjeDrugiUnos: 20,
	}
	mjesanja = append(mjesanja, mjesanje1)

	mjesanje2 := model.MjesanjeDvaUnosa{
		BodovaPrviUnos:  152,
		BodovaDrugiUnos: 10,
		ZvanjePrviUnos:  0,
		ZvanjeDrugiUnos: 20,
	}
	mjesanja = append(mjesanja, mjesanje2)

	fmt.Println(mjesanja)

	return mjesanja
}

func kreirajPartijuDvaIgraca() {
	partija := &model.PartijaDvaIgraca{}
	partija.DoKolikoSeIgra = 501
	// partija.Lokacija = *lokacija
	partija.Unosi = *igrac1
	partija.Igraci = []model.Igrac{*igrac1, *igrac2}
	// partija.MjesanjeDvaPara = kreirajMjesanjaDvaPara()
	partija.MjesanjeDvaPara= kreirajMjesanjaDvaPara()
	partije = append(partije, &partija.Partija)

}

func kreirajIgraca(ime, prezime string) *model.Igrac {
	igrac := model.Igrac{
		Ime:     ime,
		Prezime: prezime,
	}
	return &igrac
}
