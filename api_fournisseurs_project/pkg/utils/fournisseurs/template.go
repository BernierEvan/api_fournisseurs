// template.go est le squelette du code  de recherche avec les différents API des différents sites -> Il fautdra rajouter le fait de faire des imports automatiques de package et le fait que les packages doivent communiquer entre eux pour que lorsque l'on rajoute l'API d'un nouveau fournisseur, on ait pas besoin de changer le code.


package fournisseurs // package

import (
    
    "time" //importe le package temps deja disponible dans go
)

type Produit struct { // creation de la structure produit qui récuperera
    Nom            string  `json:"nom"`
    Reference      string  `json:"reference"`
    PrixUnitaire   float64 `json:"prix_unitaire"`
    Stock          int     `json:"stock"`
    Distributeur   string  `json:"distributeur"`
    RefFournisseur string  `json:"ref_fournisseur"`
}

func SimulateResultsSearch(ref string)(*Produit, error){ //création de la fonction simulation de résultat pour voir si l'affichage du résultat marcherait bien
	time.Sleep(500 * time.Millisecond)

	produit := &Produit{
		Nom: "Carte ESP32",
		Reference: ref,
		PrixUnitaire: 9.99,
		Stock: 42,
		Distributeur: "DFRobot",
		RefFournisseur: "DFR0478",
	}
	return produit, nil
}


	// To see the output of the Example while running the test suite (go test), simply
	// remove the leading "x" before Output on the next line. This will cause the
	// example to fail (all the "real" tests should pass).

	// xOutput: voluntarily fail the Example output.
} package main

import ( //importation des différents packages nécessaires
	"fmt"
	"log"
	"net/http"
	"github.com/PuerkitoBio/goquery"
)

func main() { //function principale
	// Request the HTML page.
	res, err := http.Get("http://metalsucks.net") //on accède à un site web
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 { // si le code est a 200 c'est a dire réussie
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body) //alors on load le body dans un document
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".sidebar-reviews article .content-block").Each(func(i int, s *goquery.Selection) { // et enfin on recherche dedans si il y a les éléments que l'on veut. C'est a dire ici les éléments de la structure produit crée précédemment dans le template.go
		// For each item found, get the band and title
		band := s.Find("a").Text() // permet de recupérer le texte ou se trouve x element et de le mettre dans la variable band.
		title := s.Find("i").Text() // permet de recupérer le texte ou ser trouve x element et de le mettre dans la variable title
		fmt.Printf("Review %d: %s - %s\n", i, band, title) //permet de print et d'associer chaque valeur % avec celles qui suivent dans le fmt.Printf
	})
