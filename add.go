package main

import "syscall/js"
import "strconv"

var document = js.Global().Get("document")

func main() {
	operand1 := document.Call("getElementById", "operand_1")
	operand2 := document.Call("getElementById", "operand_2")
	add := document.Call("getElementById", "add")
	result := document.Call("getElementById", "result")

	add.Set("onclick", js.NewCallback(func([]js.Value) {

		value1, _ := strconv.Atoi(operand1.Get("value").String())
		value2, _ := strconv.Atoi(operand2.Get("value").String())
		sum := value1 + value2
		result.Set("innerText", sum)

	}))
	select {}
}
