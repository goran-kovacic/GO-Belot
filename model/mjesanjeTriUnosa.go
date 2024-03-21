package model

type MjesanjeTriUnosa struct {
	MjesanjeDvaUnosa
	BodovaTreciUnos int
	ZvanjeTreciUnos int
}

func (m MjesanjeTriUnosa) GetRezultat() Rezultat {
	rezultat := m.MjesanjeDvaUnosa.GetRezultat()
	rezultat.Treci = m.BodovaTreciUnos + m.ZvanjeTreciUnos
	return rezultat
}
