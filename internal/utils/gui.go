package utils

import (
	"fmt"
	"image/color"
	"regexp"
	"strconv"
	"tiles/internal/models"
)

// checkGridSize checks if a newly created tile's coordinates are bugger than the sizes of the checkGridSize
//
// @param ui The graphic interface
// @param tile The newly created tile
func checkGridSize(ui *models.UI, x, y int) {
	// if necessary fix the size of the grid
	if x > ui.Columns {
		ui.Columns = x
	}
	if y > ui.Rows {
		ui.Rows = y
	}
}

// ParseHexColor parses the hexadecimal color of a tile in a color.Color object
//
// @param s The hex string
// @return color.Color The object
// @error An error that can occur
func ParseHexColor(hex string) (color.Color, error) {
	if len(hex) != 7 || hex[0] != '#' {
		return nil, fmt.Errorf("Formato colore non valido: %s", hex)
	}

	r, err := strconv.ParseUint(hex[1:3], 16, 8)
	if err != nil {
		return nil, err
	}
	g, err := strconv.ParseUint(hex[3:5], 16, 8)
	if err != nil {
		return nil, err
	}
	b, err := strconv.ParseUint(hex[5:7], 16, 8)
	if err != nil {
		return nil, err
	}

	return color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: 0xff}, nil
}

// IsValidHexColor checks if the hexadecimal color is correct
//
// @param s The hexadecimal string
// @return bool Returns True if the color is corretc, False otherwise
func IsValidHexColor(s string) bool {
	re := regexp.MustCompile(`^#[0-9a-fA-F]{6}$`)
	return re.MatchString(s)
}
