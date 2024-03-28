package modelStruct

type Rezultat struct{
	Prvi int
	Drugi int
	Treci int
}

func (r Rezultat) IsPocetak() bool{ 
	return r.Prvi == 0 && r.Drugi == 0 && r.Treci == 0
}

type Entitet struct {
    Id int
}

type Mjesanje interface {
	GetRezultat() Rezultat
}

type MjesanjeDvaUnosa struct{
	BodovaPrviUnos  int
	BodovaDrugiUnos int
	ZvanjePrviUnos  int
	ZvanjeDrugiUnos int

	Entitet
	Stiglja  bool
	Belot bool
}

func (m MjesanjeDvaUnosa) GetRezultat() Rezultat{
	return Rezultat{
		Prvi: m.BodovaPrviUnos + m.ZvanjePrviUnos,
		Drugi: m.BodovaDrugiUnos + m.ZvanjeDrugiUnos,
	}
}

type MjesanjeTriUnosa struct{
	MjesanjeDvaUnosa
	
	BodovaTreciUnos int
	ZvanjeTreciUnos int

	Entitet
	Stiglja  bool
	Belot bool
}

func (m MjesanjeTriUnosa) GetRezultat() Rezultat{
	rez := m.MjesanjeDvaUnosa.GetRezultat()
	rez.Treci = m.BodovaTreciUnos + m.ZvanjeTreciUnos
	return rez
}