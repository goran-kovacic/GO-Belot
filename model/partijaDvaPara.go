package model

import "fmt"

type PartijaDvaPara struct {
	Partija
}

func (p PartijaDvaPara) String() string {
	rez := p.GetRezultat()
	return fmt.Sprintf("Partija DVA PARA, igra gotova: %v, %s i %s : %d | %s i %s : %d", 
	p.IsIgraGotova(), p.Igraci[0], p.Igraci[1], rez.Prvi, p.Igraci[2], p.Igraci[3], rez.Drugi)

}