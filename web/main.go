package main

import (
	"strings"
	"syscall/js"

	"pkg.amethysts.studio/renpy-graphviz/parser"
)

func main() {

	js.Global().Set("printMessage", js.FuncOf(printMessage))

	println("exiting")
	<-make(chan bool)
}

func printMessage(this js.Value, inputs []js.Value) interface{} {
	callback := inputs[len(inputs)-1:][0]

	// fmt.Println("input", inputs[0].String())

	// api.github.com/search/code?accept=application/vnd.github.v3+json&q=repo:amethysts-studio/coalescence+extension:rpy
	renpyRepoCodeLines := strings.Split(inputs[0].String(), "\n") //[]string{"label hello:", "world", "jump label2"} //getRenpyFromRepo(inputs[0].String())

	// fmt.Println("string inside Go - renpy lines", renpyRepoCodeLines)

	dotGraph := GraphWASM(renpyRepoCodeLines, inputs[1].Bool(), inputs[2].Bool())

	// fmt.Println("string inside Go - graph", dotGraph.String())

	// document := js.Global().Get("document")
	// p := document.Call("createElement", "p")
	// p.Set("innerHTML", dotGraph.String())
	// document.Get("body").Call("appendChild", p)
	callback.Invoke(js.Null(), dotGraph.String())
	return nil
}

func GraphWASM(text []string, options ...bool) parser.RenpyGraph {
    return parser.Graph(text, parser.RenpyGraphOptions{ShowScreens: true, ShowNestedScreens:false, ShowEdgesLabels: options[0], ShowAtoms: options[1]})
}
