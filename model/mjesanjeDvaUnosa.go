package model

type MjesanjeDvaUnosa struct {
	Mjesanje
	BodovaPrviUnos  int
	BodovaDrugiUnos int
	ZvanjePrviUnos  int
	ZvanjeDrugiUnos int
}

func (m MjesanjeDvaUnosa) GetRezultat() Rezultat {
	return Rezultat{
		Prvi:  m.BodovaPrviUnos + m.ZvanjePrviUnos,
		Drugi: m.BodovaDrugiUnos + m.ZvanjeDrugiUnos,
	}
}

