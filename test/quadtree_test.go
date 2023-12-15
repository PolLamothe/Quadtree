package test

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/floor"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/quadtree"
	"testing"
)

func TestGetQuadtree(t *testing.T) { // Test servant a vérifier que le quadtree stock bien la map
	var floorContent [][]int = floor.ReadFloorFromFile("../floor-files/validation") // on stock la map dans un tableau
	var Quad quadtree.Quadtree = quadtree.MakeFromArray(floorContent)               //on créer un quadtree qui contient toute la map
	var f *floor.Floor = &floor.Floor{QuadtreeContent: Quad, FullContent: floorContent}
	f.Content = make([][]int, len(floorContent))
	for y := 0; y < len(f.Content); y++ {
		f.Content[y] = make([]int, len(floorContent[y]))
	}
	f.QuadtreeContent.GetContent(0, 0, f.Content)
	for i := 0; i < len(floorContent); i++ {
		for x := 0; x < len(floorContent[i]); x++ {
			if floorContent[i][x] != f.Content[i][x] { // on vérifie que les valeurs du tableau et du quadtree sont les mêmes
				t.Errorf("Les valeurs ne sont pas les mêmes") //sinon on affiche une erreur
			}
		}
	}
}

func TestNodeUnitaire(t *testing.T) {
	var floorContent [][]int = [][]int{{0, 1, 0, 1}, {1, 0, 1, 0}, {0, 1, 0, 1}, {1, 0, 1, 0}}
	var Quad quadtree.Quadtree = quadtree.MakeFromArray(floorContent)
	var LastNode quadtree.Node = GetLastNode(1, 1, len(floorContent), len(floorContent[0]), Quad.Root)
	if LastNode.Content != floorContent[1][1] {
		t.Errorf("Les valeurs ne sont pas les mêmes")
	}
	if LastNode.Width != 1 {
		t.Errorf("La largeur n'est pas la bonne")
	}
	if LastNode.Height != 1 {
		t.Errorf("La hauteur n'est pas la bonne")
	}
	if LastNode.TopLeftX != 1 {
		t.Errorf("Le TopLeftX n'est pas la bonne")
	}
	if LastNode.TopLeftY != 1 {
		t.Errorf("Le TopLeftY n'est pas la bonne")
	}
}

func TestNodeOptimise(t *testing.T) { // Test servant a vérifier qu'un node définie son contenu si tout les blocs qui le constitue sont les mêmes
	var floorContent [][]int = [][]int{{0, 0, 1, 1}, {0, 0, 1, 1}, {1, 1, 0, 0}, {1, 1, 0, 0}}
	var Quad quadtree.Quadtree = quadtree.MakeFromArray(floorContent)
	var LastNode quadtree.Node = GetLastNode(1, 1, len(floorContent), len(floorContent[0]), Quad.Root)
	if LastNode.Content != 0 {
		t.Errorf("Les valeurs ne sont pas les mêmes")
	}
	if LastNode.Width != 2 {
		t.Errorf("La largeur n'est pas la bonne")
	}
	if LastNode.Height != 2 {
		t.Errorf("La hauteur n'est pas la bonne")
	}
	if LastNode.TopLeftX != 0 {
		t.Errorf("Le TopLeftX n'est pas la bonne")
	}
	if LastNode.TopLeftY != 0 {
		t.Errorf("Le TopLeftY n'est pas la bonne")
	}
}

func TestHorsMapsUnitaire(t *testing.T) {
	var floorContent [][]int = [][]int{{0, 0, 1, 1}, {0, 0, 1, 1}, {1, 1, 0, 0}, {1, 1, 0, 0}}
	var Quad quadtree.Quadtree = quadtree.MakeFromArray(floorContent)
	if Quad.GetNumberFromQuad(-1, 0, len(floorContent), len(floorContent[0]), Quad.Root) != -1 {
		t.Errorf("La valeur hors du terrain n'est pas la bonne")
	}
}

func TestGetQuadtreeHorsMap(t *testing.T) { // Test servant a vérifier que GetContent renvoie bien -1 si les coordonnées sont en dehors de la map
	var floorContent [][]int = [][]int{{0, 0, 1, 1}, {0, 0, 1, 1}, {1, 1, 0, 0}, {1, 1, 0, 0}}
	var correction [][]int = [][]int{{-1, -1, -1, -1}, {-1, 0, 0, 1}, {-1, 0, 0, 1}, {-1, 1, 1, 0}} // le tableau qui doit etre affiché au coordonnées -1,-1
	var Quad quadtree.Quadtree = quadtree.MakeFromArray(floorContent)                               //on créer un quadtree qui contient toute la map
	var f *floor.Floor = &floor.Floor{QuadtreeContent: Quad, FullContent: floorContent}
	f.Content = make([][]int, len(floorContent))
	for y := 0; y < len(f.Content); y++ {
		f.Content[y] = make([]int, len(floorContent[y]))
	}
	f.QuadtreeContent.GetContent(-1, -1, f.Content)
	for i := 0; i < len(floorContent); i++ {
		for x := 0; x < len(floorContent[i]); x++ {
			if correction[i][x] != f.Content[i][x] { // on vérifie que les valeurs du tableau et du quadtree sont les mêmes
				t.Errorf("Les valeurs ne sont pas les mêmes") //sinon on affiche une erreur
			}
		}
	}
}

func GetLastNode(indexX, indexY, height, width int, current *quadtree.Node) quadtree.Node { //fonction qui sert a ramener la Node final d'un index donné
	var currentWidth int = (*current).Width
	var currentHeight int = (*current).Height
	if indexX < 0 || indexY < 0 || indexX >= width || indexY >= height { // si les position sont en dehors du terrain
		return quadtree.Node{}
	}
	if (*current).Content != -1 { //si le node actuelle possede une valeur
		return (*current)
	}
	if (currentWidth)%2 != 0 { //si la largeur du node actuelle est impair la largeur des nodes enfants a Gauche sera augmenté
		currentWidth += 1
	}
	if (currentHeight)%2 != 0 { //si la hauteur du node actuelle est impair la hauteur des nodes enfants en Haut sera augmenté
		currentHeight += 1
	}
	var xChoice int = 0
	if (*current).TopLeftX+currentWidth/2 > indexX {
		xChoice = 1
	} else if (*current).TopLeftX+current.Width > indexX {
		xChoice = 2
	}
	var yChoice int = 0
	if (*current).TopLeftY+currentHeight/2 > indexY {
		yChoice = 1
	} else if (*current).TopLeftY+(*current).Height > indexY {
		yChoice = 2
	}
	var suivant *quadtree.Node
	if xChoice == 1 && yChoice == 1 {
		suivant = (*current).TopLeftNode
	} else if xChoice == 2 && yChoice == 1 {
		suivant = (*current).TopRightNode
	} else if xChoice == 1 && yChoice == 2 {
		suivant = (*current).BottomLeftNode
	} else if xChoice == 2 && yChoice == 2 {
		suivant = (*current).BottomRightNode
	} else {
		return quadtree.Node{}
	}
	return GetLastNode(indexX, indexY, height, width, suivant)
}
