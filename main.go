package main

func main() {
	path := PlugCLI()

	println(path)

	text := FileHandler(path[0])

	println(text)

	g := ParseRenPy(text)

	DrawGraph(g)

}
