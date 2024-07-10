
# Nom du binaire à produire
BINARY_NAME=hangman.exe

# Liste explicite des fichiers source
SOURCES := src/*.go \
 
# Commande de build
build:
	@echo "Construction du projet..."
	go build -o $(BINARY_NAME) $(SOURCES)

# Commande pour nettoyer le projet (supprimer le binaire)
clean:
	@echo "Nettoyage..."
	rm $(BINARY_NAME)
 
# Commande pour exécuter le programme
run: 
	@echo "Execution du programme..."
	./$(BINARY_NAME)

restart: clean build run

# Option 'phony' pour indiquer que 'clean', 'run', et 'build' ne sont pas des fichiers
.PHONY: build clean run restart