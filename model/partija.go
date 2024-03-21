package model

import (
	"fmt"
)

type Partija struct {
	Entitet

	DoKolikoSeIgra int
	Lokacija       Lokacija
	Unosi          Igrac
	Mjesanja       []Mjesanje
	Igraci         []Igrac

	MjesanjeDvaPara   []MjesanjeDvaUnosa
	MjesanjeTriIgraca []MjesanjeTriUnosa
	MjesanjeDvaIgraca []MjesanjeDvaUnosa
	
}



func (p Partija) String() string {
	rez := p.GetRezultat()
	number := len(p.Igraci)

	if number == 2 {
		
		return fmt.Sprintf("Partija DVA IGRACA, igra gotova: %v, %s: %d | %s: %d",
			p.IsIgraGotova(), p.Igraci[0], rez.Prvi, p.Igraci[1], rez.Drugi)
	} else if number == 3 {
		return fmt.Sprintf("Partija TRI IGRACA, igra gotova: %v, %s : %d | %s : %d | %s : %d",
			p.IsIgraGotova(), p.Igraci[0], rez.Prvi, p.Igraci[1], rez.Drugi, p.Igraci[2], rez.Treci)
	} else {
		return fmt.Sprintf("Partija DVA PARA, igra gotova: %v, %s i %s : %d | %s i %s : %d",
			p.IsIgraGotova(), p.Igraci[0], p.Igraci[1], rez.Prvi, p.Igraci[2], p.Igraci[3], rez.Drugi)
	}

}

func (p Partija) GetRezultat() Rezultat {
	// rezultat := Rezultat{Prvi: 0, Drugi: 0, Treci: 0}
	rezultat := Rezultat{}

	for _, m := range p.Mjesanja {
		rezultat.Prvi += m.Rezultat().Prvi
		rezultat.Drugi += m.Rezultat().Drugi
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

func (p Partija) IsIgraGotovaV2() bool {
	rez := p.GetRezultat()

	if rez.Treci == 0 {
		return rez.Prvi == rez.Drugi || rez.Prvi > p.DoKolikoSeIgra || rez.Drugi > p.DoKolikoSeIgra
	}

	if rez.Prvi == rez.Drugi || rez.Prvi == rez.Treci || rez.Drugi == rez.Treci {
		return false
	}

	if rez.Prvi > p.DoKolikoSeIgra || rez.Drugi > p.DoKolikoSeIgra || rez.Treci > p.DoKolikoSeIgra {
		return true
	}

	return false
}

func (p *Partija) IsIgraGotovaV3() bool {
	r := p.GetRezultat()
	if r.Treci == 0 {
		return r.Prvi == r.Drugi && (r.Prvi > p.DoKolikoSeIgra || r.Drugi > p.DoKolikoSeIgra)
	} else {
		if r.Prvi == r.Drugi || r.Prvi == r.Treci || r.Drugi == r.Treci {
			return false
		}
		return r.Prvi > p.DoKolikoSeIgra || r.Drugi > p.DoKolikoSeIgra || r.Treci > p.DoKolikoSeIgra
	}
}
