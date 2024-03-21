package model

import "fmt"

type PartrijaTriIgraca struct {
	Partija
}

func (p PartrijaTriIgraca) GetRezultat() Rezultat {
	rez := p.Partija.GetRezultat()

	for _, m := range p.Mjesanja {
		rez.Treci += m.Rezultat().Treci
	}
	return rez
}

func (p PartrijaTriIgraca) String() string {
	rez := p.GetRezultat()
	return fmt.Sprintf("Partija TRI IGRACA, igra gotova: %v, %s : %d | %s : %d | %s : %d",
		p.IsIgraGotova(), p.Igraci[0], rez.Prvi, p.Igraci[1], rez.Drugi, p.Igraci[2], rez.Treci)
}
