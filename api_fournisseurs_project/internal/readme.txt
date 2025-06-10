Le dossier internal est essentiel car il contient la logique principale de votre application. Il est divisé en plusieurs sous-dossiers pour organiser les différentes parties de votre application :

handlers : Contient les gestionnaires HTTP qui traitent les requêtes entrantes et renvoient les réponses appropriées.
models : Contient les définitions des structures de données utilisées dans votre application.
services : Contient la logique métier de votre application.
repositories : Contient le code pour interagir avec les bases de données ou d'autres sources de données.
Exemple
Supposons que vous avez une application web simple. Le dossier internal pourrait ressembler à ceci :

Copier
internal/
├── handlers/
│   └── example_handler.go
├── models/
│   └── example_model.go
├── services/
│   └── example_service.go
└── repositories/
    └── example_repository.go