package template

type Theme struct {
	White     ColorTuple
	Orange    ColorTuple
	Yellow    ColorTuple
	Green     ColorTuple
	Purple    ColorTuple
	Blue      ColorTuple
	Red       ColorTuple
	Lightgrey ColorTuple
	Grey      ColorTuple
	Black     ColorTuple
	Invisible ColorTuple
	Error     ColorTuple
}

func GenerateIconPalette() []ColorTuple {
	return []ColorTuple{
		{
			CATHEX: "#cdd6f4",
			Hex:    "#d4d4d5",
		},
		{
			CATHEX: "#fab387",
			Hex:    "#f0a988",
		},
		{
			CATHEX: "#f9e2af",
			Hex:    "#e5d487",
		},
		{
			CATHEX: "#a6e3a1",
			Hex:    "#37d99e",
		},
		{
			CATHEX: "#cba6f7",
			Hex:    "#c397d8",
		},
		{

			CATHEX: "#89dceb",
			Hex:    "#5fb0fc",
		},
		{
			CATHEX: "#f38ba8",
			Hex:    "#e87979",
		},
		{
			CATHEX: "#585b70",
			Hex:    "#45484c",
		},
		{
			CATHEX: "#1e1e2e",
			Hex:    "#191d22",
		},
		{
			CATHEX: "#11111b",
			Hex:    "#101317",
		},
		{
			CATHEX: "#f2cdcd",
			Hex:    "#d4d4d5",
		},
		{
			CATHEX: "#f5c2e7",
			Hex:    "#c397d8",
		},
		{
			CATHEX: "#eba0ac",
			Hex:    "#e87979",
		},
		{
			CATHEX: "#94e2d5",
			Hex:    "#37d99e",
		},
		{
			CATHEX: "#74c7ec",
			Hex:    "#5fb0fc",
		},
		{
			CATHEX: "#89b4fa",
			Hex:    "#5fb0fc",
		},
		{
			CATHEX: "#b4befe",
			Hex:    "#c397d8",
		},
		{
			CATHEX: "#bac2de",
			Hex:    "#d4d4d5",
		},
		{
			CATHEX: "#a6adc8",
			Hex:    "#45484c",
		},
		{
			CATHEX: "#9399b2",
			Hex:    "#45484c",
		},
		{
			CATHEX: "#7f849c",
			Hex:    "#45484c",
		},
		{
			CATHEX: "#6c7086",
			Hex:    "#d4d4d5",
		},
		{
			CATHEX: "#45475a",
			Hex:    "#45484c",
		},
		{
			CATHEX: "#313244",
			Hex:    "#45484c",
		},
		{
			CATHEX: "#181825",
			Hex:    "#191d22",
		},
	}
}

// ColorTuple is a color in Hex, RGB, and HSL.
type ColorTuple struct {
	CATHEX string
	Hex    string
}
