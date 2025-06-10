Contient des fonctions utilitaires qui peuvent être utilisées dans différents endroits de votre application.

package utils

import (
	"log"
)

func LogError(err error) {
	if err != nil {
		log.Printf("Erreur: %v", err)
	}
}


Cas d'utilisation

Utilitaires : Les fonctions utilitaires peuvent être utilisées dans différents endroits de votre application pour fournir des fonctionnalités communes.

