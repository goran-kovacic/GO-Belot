package modelStruct

import "fmt"

type PartijaSucelje interface {
	String() string
}

type Partija struct {
	Entitet
	DoKolikoSeIgra int
	Unosi          Igrac
	Mjesanja       []Mjesanje
	Igraci         []Igrac
}

// func (p Partija) String() string {
// 	return "IZ NADKLASE"
// }

func (p Partija) GetRezultat() Rezultat {
	rezultat := Rezultat{Prvi: 0, Drugi: 0, Treci: 0}
	// rezultat := Rezultat{}

	for _, m := range p.Mjesanja {
		rezultat.Prvi += m.GetRezultat().Prvi
		rezultat.Drugi += m.GetRezultat().Drugi
	}
	return rezultat
}

func (p Partija) IsIgraGotova() bool {
	rezultat := p.GetRezultat()

	if rezultat.Treci == 0 {
		if rezultat.Prvi == rezultat.Drugi {
			return false
		} else {
			return rezultat.Prvi > p.DoKolikoSeIgra || rezultat.Drugi > p.DoKolikoSeIgra
		}
	} else {
		if rezultat.Prvi == rezultat.Drugi || rezultat.Prvi == rezultat.Treci || rezultat.Drugi == rezultat.Treci {
			return false
		}
		if rezultat.Prvi > p.DoKolikoSeIgra || rezultat.Drugi > p.DoKolikoSeIgra || rezultat.Treci > p.DoKolikoSeIgra {
			return true
		}
	}
	return false
}

type PartijaDvaIgraca struct {
	Partija
}

func (p PartijaDvaIgraca) String() string {
	rez := p.Partija.GetRezultat()
	return fmt.Sprintf("Partija DVA IGRACA, igra gotova: %v, %s: %d | %s: %d",
		p.IsIgraGotova(), p.Igraci[0], rez.Prvi, p.Igraci[1], rez.Drugi)
}
