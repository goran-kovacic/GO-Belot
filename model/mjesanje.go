package model

type Mjesanje struct{
	Entitet

	Rezultat func() Rezultat

	Stiglja  bool
	Belot bool
}