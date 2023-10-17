package template

import (
	"image/color"
)

type Theme struct {
	White     Color
	Orange    Color
	Yellow    Color
	Green     Color
	Purple    Color
	Blue      Color
	Red       Color
	Lightgrey Color
	Grey      Color
	Black     Color
	Invisible Color
	Error     Color
}

func GeneratePalette() []Color {
	return []Color{
		Color{
			TemplateName: "{{white}}",
			Hex:          "#d4d4d5",
			RGB:          [3]uint32{212, 212, 213},
			HSL:          [3]float32{240, 0.01, 0.83},
		},
		Color{
			TemplateName: "{{orange}}",
			Hex:          "#f0a988",
			RGB:          [3]uint32{240, 169, 136},
			HSL:          [3]float32{19, 0.78, 0.74},
		},
		Color{
			TemplateName: "{{yellow}}",
			Hex:          "#e5d487",
			RGB:          [3]uint32{229, 212, 135},
			HSL:          [3]float32{49, 0.64, 0.71},
		},
		Color{
			TemplateName: "{{green}}",
			Hex:          "#37d99e",
			RGB:          [3]uint32{55, 217, 158},
			HSL:          [3]float32{158, 0.68, 0.53},
		},
		Color{
			TemplateName: "{{purple}}",
			Hex:          "#c397d8",
			RGB:          [3]uint32{195, 151, 216},
			HSL:          [3]float32{281, 0.45, 0.72},
		},
		Color{
			TemplateName: "{{blue}}",
			Hex:          "#5fb0fc",
			RGB:          [3]uint32{95, 176, 252},
			HSL:          [3]float32{209, 0.96, 0.68},
		},
		Color{
			TemplateName: "{{red}}",
			Hex:          "#e87979",
			RGB:          [3]uint32{232, 121, 121},
			HSL:          [3]float32{0, 0.71, 0.69},
		},
		Color{
			TemplateName: "{{lightgrey}}",
			Hex:          "#45484c",
			RGB:          [3]uint32{69, 72, 76},
			HSL:          [3]float32{214, 0.05, 0.28},
		},
		Color{
			TemplateName: "{{grey}}",
			Hex:          "#191d22",
			RGB:          [3]uint32{25, 29, 34},
			HSL:          [3]float32{213, 0.15, 0.12},
		},
		Color{
			TemplateName: "{{black}}",
			Hex:          "#101317",
			RGB:          [3]uint32{16, 19, 23},
			HSL:          [3]float32{214, 0.18, 0.08},
		},
		Color{
			TemplateName: "{{invisible}}",
			Hex:          "#ffffff00",
			RGB:          [3]uint32{0, 0, 0},
			HSL:          [3]float32{0, 0, 0},
		},
		Color{
			TemplateName: "{{error}}",
			Hex:          "#f44747",
			RGB:          [3]uint32{244, 71, 71},
			HSL:          [3]float32{0, 0.89, 0.62},
		},
	}
}

// Color is a color in Hex, RGB, and HSL.
type Color struct {
	TemplateName string
	Hex          string
	RGB          [3]uint32
	HSL          [3]float32
}

// RGBA implements color.Color
func (c Color) RGBA() (r uint32, g uint32, b uint32, a uint32) {
	return c.RGB[0], c.RGB[1], c.RGB[2], 1
}

func (c Color) getName() string {
	return c.TemplateName
}

var _ color.Color = Color{}
