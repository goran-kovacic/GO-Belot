package model

import "fmt"

type PartijaDvaIgraca struct {
	Partija
}

func (p PartijaDvaIgraca) String() string {
	rez := p.GetRezultat()
	return fmt.Sprintf("Partija DVA IGRACA, igra gotova: %v, %s: %d | %s: %d",
		p.IsIgraGotova(), p.Igraci[0], rez.Prvi, p.Igraci[1], rez.Drugi)
}
