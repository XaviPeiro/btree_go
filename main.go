package main

import (
	"fmt"
)

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
func main() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	for i := 1; i <= 5; i++ {
		//TIP <p>To start your debugging session, right-click your code in the editor and select the Debug option.</p> <p>We have set one <icon src="AllIcons.Debugger.Db_set_breakpoint"/> breakpoint
		// for you, but you can always add more by pressing <shortcut actionId="ToggleLineBreakpoint"/>.</p>
		fmt.Println("i =", 100/i)
	}
	fmt.Println(slices())

}

func slices() int {
	var a1 []int
	a1 = []int{1, 2, 3, 4 ,5}

	fmt.Println(a1[2:])
	fmt.Println(a1[1:])
	copy(a1[0:1], a1[4:])	
	fmt.Println(a1)
	b2 := [10]int{}
	c1 := make([]int, 0, 10)
	fmt.Println(b2)
	fmt.Println(len(b2))
	fmt.Println(c1)

	pp := 33
	xx := &pp
	println(*xx+3)

	return 0
}
