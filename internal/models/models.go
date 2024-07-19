package models

import (
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

// Piastrella represents a tile of the system
//
// X Y The coordinates of the tile
type Piastrella struct {
	X int
	Y int
}

// Properties represents the properties of a tile
//
// Color The color of the tile
// Intensity The intensity of the color
// Parent The parent the tile is referring to
// Size The number of tiles that refers to the tile
// BlockIntensity The sum of the intensity of the tiles that refers to the tile
type Properties struct {
	Color          string
	Intensity      int
	Parent         Piastrella
	Size           int
	BlockIntensity int
}

// Ruleset represents the required tiles that need to be adjacents to the tile
//
// Color The color of the adjacent tile
// Count The number of the adjacents tile of color @color
type Ruleset struct {
	Color string
	Count int
}

// Rule represents a rule of the system
//
// Ruleset The list of required adjacent tiles
// Color The color that the rule gives to teh tile
// Usage The number of times the rule has been used
type Rule struct {
	Ruleset []Ruleset
	Color   string
	Usage   int
}

// Piano The system tiles-rules
//
// Tiles The set of all the tiles of the system
// Rules The list of all the rules of the systems
type Piano struct {
	Tiles map[Piastrella]*Properties
	Rules *[]Rule
}

// Find is a method that finds the root of the block the tile x belongs to using path compression
//
// @param x The tile
// @return piastrella The root of the block
func (p Piano) Find(x Piastrella) Piastrella {
	// find the root of the block
	if p.Tiles[x].Parent != x {
		p.Tiles[x].Parent = p.Find(p.Tiles[x].Parent) // path compression
	}
	return p.Tiles[x].Parent
}

// Union is a method that unifies to sets of tiles using union by size
//
// @param x y The coordinates of the tile
func (p Piano) Union(x, y Piastrella) {
	// unify the two blocks depending on their sizes
	rootX := p.Find(x)
	rootY := p.Find(y)

	if rootX != rootY {
		if p.Tiles[rootX].Size > p.Tiles[rootY].Size {
			p.Tiles[rootY].Parent = rootX
			p.Tiles[rootX].Size += p.Tiles[rootY].Size + 1
			p.Tiles[rootY].Size = 0
			p.Tiles[rootX].BlockIntensity += p.Tiles[rootY].BlockIntensity
			p.Tiles[rootY].BlockIntensity = p.Tiles[rootY].Intensity
		} else if p.Tiles[rootY].Size > p.Tiles[rootX].Size {
			p.Tiles[rootX].Parent = rootY
			p.Tiles[rootY].Size += p.Tiles[rootX].Size + 1
			p.Tiles[rootX].Size = 0
			p.Tiles[rootY].BlockIntensity += p.Tiles[rootX].BlockIntensity
			p.Tiles[rootX].BlockIntensity = p.Tiles[rootX].Intensity
		} else {
			p.Tiles[rootX].Parent = rootY
			p.Tiles[rootY].Size++
			p.Tiles[rootY].BlockIntensity += p.Tiles[rootX].BlockIntensity
		}
	}
}

// Add is a method taht creates the tile and colors it
//
// @param x The tile
// @param c The color
// @param i The intensity of the color
func (p Piano) Add(x Piastrella, c string, i int) {
	if _, exists := p.Tiles[x]; !exists {
		p.Tiles[x] = &Properties{Color: c, Intensity: i, Parent: x, Size: 0, BlockIntensity: i}
	} else {
		root := p.Find(x)
		p.Tiles[x].Color = c
		oldInt := p.Tiles[x].Intensity
		p.Tiles[x].Intensity = i
		p.Tiles[root].BlockIntensity += i - oldInt
	}
}

// UI The graphic interface
//
// App The app itself
// Window The window of the app
// Grid The grid representing the piano
// Buttons The buttons of the several functions
// Columns The number of columns of the grid
// Rows The number of row of the grid
type UI struct {
	App     fyne.App
	Window  fyne.Window
	Grid    *fyne.Container
	Rects   [][]*canvas.Rectangle
	Buttons *fyne.Container
	List    *fyne.Container
	Rules   *[]*widget.Label
	Columns int
	Rows    int
}

// UpdateCell updates the grid when a new tile is coloured
//
// @param x y The coordinates of the tile (y must be shifted)
// @param hex The RGBA color
func (ui UI) UpdateCell(x, y int, hex color.Color) {
	ui.Rects[y][x].FillColor = hex
	ui.Rects[y][x].Refresh()
}

// CreateItem create a new element of the rules's list
//
// @param r The rule to add to the list
// @return fyne.CanvasObject The newly created element
func (ui UI) CreateRule(rule Rule) string {
	res, rs := "", ""
	for _, r := range rule.Ruleset {
		rs = rs + strconv.Itoa(r.Count) + " " + r.Color + " "
	}
	res = res + "Usage: " + strconv.Itoa(rule.Usage) + " | Color: " + rule.Color + " " + rs
	return res
}

// AddRule adds a rule to the graphic list
//
// @param p The system tiles-rules
func (ui UI) AddRule(p Piano) {
	row := widget.NewLabel((ui.CreateRule((*p.Rules)[len(*p.Rules)-1])))
	*ui.Rules = append(*ui.Rules, row)
	ui.List.Add(row)
}

// UpdateRule updates the list of rule when a rule is used
//
// @param index The index of the used rule
// @param reg The rule used
func (ui UI) UpdateRule(index int, reg Rule) {
	(*ui.Rules)[index].SetText(ui.CreateRule(reg)) 
	ui.List.Refresh()
}
