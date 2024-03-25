package modelStruct

type Igrac struct {
	Entitet
	Ime     string
	Prezime string
}

func (i Igrac) String() string {
	return i.Ime + " " + i.Prezime
}
