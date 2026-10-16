go-reloaded
Outil en ligne de commande écrit en Go qui corrige un texte.
Il lit un fichier d'entrée, applique des règles d'édition, puis écrit le texte corrigé dans un fichier de sortie.

Utilisation
go run . <fichier_entree> <fichier_sortie>


Exemple :

go run . sample.txt result.txt


Entrée (sample.txt) :

it (cap) was the worst of times (up) , it was
Simply add 42 (hex) and 10 (bin)
bearing a untold story ,what do you think ?


Sortie (result.txt) :

It was the worst of TIMES, it was
Simply add 66 and 2
bearing an untold story, what do you think?


Règles gérées
| Règle | Exemple | Résultat |
|---|---|---|
| (hex) : hexadécimal vers décimal | 1E (hex) | 30 |
| (bin) : binaire vers décimal | 10 (bin) | 2 |
| (up) : majuscules | go (up) | GO |
| (low) : minuscules | SHOUTING (low) | shouting |
| (cap) : première lettre en majuscule | bridge (cap) | Bridge |
| (up, n), (low, n), (cap, n) : sur les n mots précédents | so exciting (up, 2) | SO EXCITING |
| Ponctuation . , ! ? : ; collée au mot précédent | there ,and | there, and |
| Groupes ... et !? gardés ensemble | thinking ... You | thinking... You |
| Guillemets ' ' collés aux mots entourés | ' awesome ' | 'awesome' |
| a devient an devant une voyelle ou un h | a amazing | an amazing |

Organisation du code
main.go : vérifie les arguments, lit le fichier, le traite ligne par ligne, écrit le résultat
process.go : la fonction Process, qui corrige une ligne dans cet ordre :
marqueurs
ponctuation
guillemets
articles
process_test.go : les tests

Tests
go test -v
Gestion des erreurs
Mauvais nombre d'arguments : message d'utilisation et code de sortie 1
Fichier d'entrée introuvable : message clair et code de sortie 1
Nombre invalide pour (hex) ou (bin) : le mot est laissé tel quel

Limite connue
Un marqueur collé au mot (word(up)) n'est pas reconnu : il faut un espace avant le marqueur (word (up)).