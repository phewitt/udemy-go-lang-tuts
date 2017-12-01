package main

func main() {

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#44bb00",
		"white": "#ffffff",
	}

	printMap(colors)
}

func printMap(m map[string]string) {
	for key, value := range m {
		println("[", key, "]", "=", value)
	}
}
