Utilité
Le dossier pkg contient le code qui peut être réutilisé par d'autres projets. Il est divisé en sous-dossiers selon les fonctionnalités, comme la configuration et les utilitaires.

Explication détaillée
Le dossier pkg est utile car il contient du code réutilisable qui peut être partagé entre différents projets. Il est divisé en sous-dossiers selon les fonctionnalités :

config : Contient le code pour charger et gérer la configuration de votre application.
utils : Contient des fonctions utilitaires qui peuvent être utilisées dans différents endroits de votre application.

Exemple
Supposons que vous avez une application web simple. Le dossier pkg pourrait ressembler à ceci :

Copier
pkg/
├── config/
│   └── config.go
└── utils/
    └── utils.go