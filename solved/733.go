var Image [][]int
var Filled [50][50]bool
var Color int
var NewColor int
var Row int
var Col int

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	Image = image
	Color = image[sr][sc]
	NewColor = newColor
	Row = len(image)
	Col = len(image[0])
	for i := 0; i < Row; i++ {
		for j := 0; j < Col; j++ {
			Filled[i][j] = false
		}
	}

	fill(sr, sc)

	return Image
}

func fill(cr, cc int) {
	changeColor(cr, cc)

	if 0 < cr {
		if changeColor(cr - 1, cc) {
			fill(cr - 1, cc)
		}
	}
	if cr < Row - 1 {
		if changeColor(cr + 1, cc) {
			fill(cr + 1, cc)
		}
	}
	if 0 < cc {
		if changeColor(cr, cc - 1) {
			fill(cr, cc - 1)
		}
	}
	if cc < Col - 1 {		
		if changeColor(cr, cc + 1) {
			fill(cr, cc + 1)
		}
	}
}

func changeColor(cr, cc int) bool {
	if Image[cr][cc] == Color && !Filled[cr][cc] {
		Image[cr][cc] = NewColor
		Filled[cr][cc] = true
		return true
	}
	return false
}