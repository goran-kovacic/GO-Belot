package model

type Igrac struct{
	Entitet

	Ime string
	Prezime string
	URLSlike string
	Spol	int
}

func (i Igrac) String() string{
	return i.Ime + " " + i.Prezime
}