package model

type Rezultat struct{
	Prvi int
	Drugi int
	Treci int

}

func (r Rezultat) IsPocetak() bool{ // mozda pointer (r *Rezultat)
	return r.Prvi == 0 && r.Drugi == 0 && r.Treci == 0
}