package jobs

// classe JobsManager

// méthode New() -> Manager
// méthode RunAsync() -> Lance une go-routine de la méthode RunSync() -> check toutes les minutes les triggers qui sont activés et
// exécuter les tâches correspondantes
// Quand tu lances le traitement d'un nouveau trigger utilise des goroutines pour accélérer le process
