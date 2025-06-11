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


