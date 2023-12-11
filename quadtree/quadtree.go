package quadtree

// Quadtree est la structure de données pour les arbres
// quaternaires. Les champs non exportés sont :
//   - Width, Height : la taille en cases de la zone représentée
//     par l'arbre.
//   - Root : le nœud qui est la racine de l'arbre.
type Quadtree struct {
	Width, Height int
	Root          *node
}

// node représente un nœud d'arbre quaternaire. Les champs sont :
//   - TopLeftX, TopLeftY : les coordonnées (en cases) de la case
//     située en haut à gauche de la zone du terrain représentée
//     par ce nœud.
//   - width, height :  la taille en cases de la zone représentée
//     par ce nœud.
//   - content : le type de terrain de la zone représentée par ce
//     nœud (seulement s'il s'agit d'une feuille).
//   - xxxNode : Une représentation de la partie xxx de la zone
//     représentée par ce nœud, différent de nil si et seulement
//     si le nœud actuel n'est pas une feuille.
type node struct {
	TopLeftX, TopLeftY int
	width, height      int
	content            int
	topLeftNode        *node
	topRightNode       *node
	bottomLeftNode     *node
	bottomRightNode    *node
}
