package main

//Vérifie si c'est bien une lettre qui est rentré
func isLettre(s string) bool {
	return len(s) == 1 && ('a' <= s[0] && s[0] <= 'z' || 'A' <= s[0] && s[0] <= 'Z')
}
