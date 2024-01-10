# Documentation mode multijoueur local quadtree

## Introduction
L'extension du mode multijoueur quadtree est comme son nom l'indique, une extension permettant à 2 joueurs ayant leur 2 
PC connectés
au même réseau local de jouer ensemble sur une même partie.

## Fonctionnement
Cette extension utilise le protcole tcp pour communiquer entre les 2 ordinateurs. Il y'aura un serveur et un client.
Le serveur sera l'ordianteur sur lequelle la partie est lancée en premier. Ce sera donc lui qui décidera de la map et 
des extensions utilisées.
Le client lui devra spécifier l'adresse ip du serveur et le port sur lequelle il écoute pour pouvoir se connecter.

## Déroulement des étapes
1. Le serveur lance une partie en choisissant la map et les extensions
2. Le client se connecte au serveur
3. Le serveur envoie la map, les extensions, la position du joueur serveur, *ainsi que les bloc qui ont étés générés si 
l'extension de la génération infinie est activée* au client
> il faut que le joueur serveur puisse jouer même si le client n'est pas encore connecté, c'est pour cela que le serveur 
> envoie la position du 
> joueur serveur au client
4. Le client et le serveur vont maintenant s'échanger les intéractions qui ont eu lieu sur la map (déplacement, génération de bloc, pose de portail etc...)

## Explication des intéractions
**Le protocole TCP permet d'envoyer des données sous forme de bytes. Il faut donc définir un format pour les données qui 
seront envoyées.
Notamment pour que le receveur ai connaisance de l'usage qu'il doit faire des données reçues.**

Nous utiliserons donc des map qui seront utilisées ainsi :

Un attribut "API" sera utilisé qui permettera de communiquer l'usage a faire des données envoyées. *Ex: si nous envoyons 
la touche qui a été pressée par le joueur nous donnerons comme valeur a cet attribut SendKey*

Ainsi le receveur pourra savoir qu'il doit utiliser la donnée reçue pour simuler une touche pressée.

Un attribut "DATA" qui permettera de communiquer les données a utiliser. *Ex: si nous envoyons la touche qui a été pressée 
par le joueur nous donnerons comme valeur a cet attribut la touche pressée*

### Format d'envoie

Il faut donc que nous envoyons une map sous forme byte, ce qui n'est pas directement possible en go. Nous allons donc
utiliser la fonction *json.Marshall* qui permet de convertir une map en string. Nous pourrons ensuite convertir cette string
en byte et l'envoyer.

### Problème de "Double écriture"
**Le protocole tcp ouvre un "canal" entre les deux parties et qui leur permet de communiquer librement, cela ne fonctionne 
donc pas sous forme de requete mais j'ai codé un systeme pour utiliser ce canal similairement a des requêtes**

Cela peut donc poser un problème si l'un des joueurs envoie plusieurs données trop rapidement et que l'autre joueur n'a
pas eu le temp de traiter la première, cela pourrait donc causer une désynchronisation des deux joueurs.

Solution :

Nous allons donc utiliser un systeme de "requête" qui permettera de traiter les données une par une.

Pour cela le partie qui envoie les données devra attendre que le receveur lui renvoie un message de confirmation avant
d'envoyer la prochaine donnée.

#### Problème de "Double écriture" du à la réception des requêtes

Puisque la réception des requêtes ce déroule parallement au jeux, quand l'un des parties reçoit une requête, elle doit 
la traiter et renvoyer un message de confirmation. 
Mais si elle envoie ce message de confirmation en même temp que une action est effectuée, 
cela peut poser problème. car elle va donc écrire deux message en même temp sur le canal.

Solution :

Nous allons donc utiliser un systeme de "blocage" lorsque une requête est reçue, la partie qui la reçoit va bloquer le 
traitement du jeux grace à une boucle for jusqu'a ce qu'elle ai traité la requête et renvoyé un message de confirmation.

## Envoie des blocs générés par l'extension "génération infinie"

L'extension de génération infinie permet de générer des blocs de manière infinie, il faut donc que chaque partie communique 
a l'autre les blocs qui ont étés générés et leur valeur.

### Problème d'envoie trop conséquent

Dans le cas ou le joueur génère beaucoup de blocs avant que le client se connecte, cela peut poser problème car le 
serveur va devoir envoyer beaucoup de données d'un coup au client. et le buffer du canal tcp ne pourra pas toutes les
contenir.

Solution :

Nous allons donc utiliser un systeme de "paquet" qui permettera de séparer les données en plusieurs paquets de 10 blocs
et de les envoyer un par un.

### Problème de stockage des blocs générés

Les blocs générés reçue doivent donc être mis en mémoire en attendant que la coordonée du bloc soit utilisée par le client
pour pouvoir le générer. Mais cela peut poser problème si beaucoup de blocs sont recus et que le client ne les utilise pas
rapidement. Cela peut faire crasher le programme du a un overflow.

Solution :

Nous allons donc socker ces blocs dans un fichier et les supprimer une fois utilisés.