package model

type Soul struct {
	ID    int
	Title string
	Hits  int
}

type Souls struct {
	SoulsItem []*Soul
}
