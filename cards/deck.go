package main

import "fmt"

//the deck type is essentially a slice of type string, but because we are making this extra type, it gives us the ability to create custom functions THAT ONLY WORK WITH THAT DECK TYPE.

type deck []string

//aka acest type (deck) imprumuta comportamentul lui []string
// aka deck =====  []string
//apoi merg in celalalt fisier si inlocuiesc unde am slice de stringuri []string cu deck.

//creez o functie care apartine acestui deck si cand calluim functia, we should print out each individual card inside the deck.

//numele functiei de mai jos este print
//(d deck este RECEIVER - sets up methods on variables that we create)
// ce inseamna asta? ca orice variabila care are TYPE deck are acum acces la functia/metoda print()

func (d deck) numeFunctie() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// d-ul de mai sus e un fel de variabila temporara. imagineaza-ti ca ai luat variabila initiala pe care o ai in main.go "cards" si ai folosit-o aici in loc de "d"
//aka cum ai pus in for loop-ul cu care ai iterat range numeleFunctiei - aici e aceeasi treaba doar ca folosesti variabila d, CARE NU E CREATA NICAIERI ALTUNDEVA (un fel de this)
// PRIN CONVENTIE (nu-i obligatoriu) daca numele type-ului e deck, pui d, sau de, sau dec la numele variabilei respective
