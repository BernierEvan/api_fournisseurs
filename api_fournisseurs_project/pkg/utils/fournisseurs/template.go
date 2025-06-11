package fournisseurs

import (
    
    "time"
)

// Produit : structure contenant les infos récupérées
type Produit struct {
    Nom            string  `json:"nom"`
    Reference      string  `json:"reference"`
    PrixUnitaire   float64 `json:"prix_unitaire"`
    Stock          int     `json:"stock"`
    Distributeur   string  `json:"distributeur"`
    RefFournisseur string  `json:"ref_fournisseur"`
}

func SimulateResultsSearch(ref string)(*Produit, error){
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
}package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// Request the HTML page.
	res, err := http.Get("http://metalsucks.net")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".sidebar-reviews article .content-block").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		band := s.Find("a").Text()
		title := s.Find("i").Text()
		fmt.Printf("Review %d: %s - %s\n", i, band, title)
	})
