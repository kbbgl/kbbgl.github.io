package main

func main() {

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	// fmt.Println(colors)

	for color, hex := range colors {
		println(color, hex)
	}

}
