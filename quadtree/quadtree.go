package quadtree

// Quadtree est la structure de données pour les arbres
// quaternaires. Les champs non exportés sont :
//   - Width, Height : la taille en cases de la zone représentée
//     par l'arbre.
//   - Root : le nœud qui est la racine de l'arbre.
type Quadtree struct {
	width, height int
	Root          *Node
}

// Node représente un nœud d'arbre quaternaire. Les champs sont :
//   - topLeftX, topLeftY : les coordonnées (en cases) de la case
//     située en haut à gauche de la zone du terrain représentée
//     par ce nœud.
//   - Width, Height :  la taille en cases de la zone représentée
//     par ce nœud.
//   - Content : le type de terrain de la zone représentée par ce
//     nœud (seulement s'il s'agit d'une feuille).
//   - xxxNode : Une représentation de la partie xxx de la zone
//     représentée par ce nœud, différent de nil si et seulement
//     si le nœud actuel n'est pas une feuille.
type Node struct {
	TopLeftX, TopLeftY int
	Width, Height      int
	Content            int
	TopLeftNode        *Node
	TopRightNode       *Node
	BottomLeftNode     *Node
	BottomRightNode    *Node
}
