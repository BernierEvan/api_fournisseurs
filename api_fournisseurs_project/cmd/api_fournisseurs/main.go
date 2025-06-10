package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/PuerkitoBio/goquery"
)

// PageInfo représente les informations que nous voulons extraire
type PageInfo struct {
	Title string `json:"title"`
	// Ajoutez d'autres champs selon les informations que vous souhaitez extraire
}

func fetchPageInfo(url string) (*PageInfo, error) {
	// Faire une requête HTTP GET pour obtenir le contenu de la page
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Vérifier que la requête a réussi
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("la requête a échoué avec le code de statut: %d", res.StatusCode)
	}

	// Utiliser goquery pour parser le document HTML
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	// Extraire le titre de la page
	title := doc.Find("title").Text()

	// Créer une structure PageInfo avec les informations extraites
	pageInfo := &PageInfo{
		Title: title,
		// Ajoutez d'autres champs ici
	}

	return pageInfo, nil
}

func pageInfoHandler(w http.ResponseWriter, r *http.Request) {
	// Récupérer l'URL depuis les paramètres de la requête
	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(w, "Le paramètre 'url' est requis", http.StatusBadRequest)
		return
	}

	// Récupérer les informations de la page
	info, err := fetchPageInfo(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convertir les informations en JSON et les envoyer en réponse
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
}

func main() {
	// Configurer les routes
	http.HandleFunc("/api/pageinfo", pageInfoHandler)

	// Démarrer le serveur
	fmt.Println("Le serveur est démarré sur le port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
