package main

func main() {
	path := PlugCLI()

	println(path)

	text := FileHandler(path[0])

	println(text)

	ParseRenPy(text)

	// DrawGraph(tree)

}
