package util

import "fmt"

type Screenbuffer struct {
	Buffer []string
}

func (b Screenbuffer) Render() {
	for _, row := range b.Buffer {
		fmt.Println(row)
	}
}

func (b *Screenbuffer) Add(s string) {
	b.Buffer = append(b.Buffer, s)
}

func (b *Screenbuffer) Flush() {
	b.Buffer = nil
}

func examples() {

	var test = Screenbuffer{[]string{}}

	test.Add(fmt.Sprintf("[H[J"))
	test.Buffer = append(test.Buffer, "this")
	test.Buffer = append(test.Buffer, "and")
	test.Buffer = append(test.Buffer, "that")
	test.Buffer = append(test.Buffer, "someother"+"boogiewoogie")
	test.Buffer = append(test.Buffer, fmt.Sprintf("%s:%d", "stringcheese", 8))
	test.Add("dingo")
	test.Add(fmt.Sprintf("%s:%d", "Liquid Pickle", 666))

	fmt.Println(test)        // as struct
	fmt.Println(test.Buffer) // as slice
	test.Render()            // print each line

	test.Flush()
	test.Add("dingo")
	test.Add(fmt.Sprintf("%s:%d", "Liquid Pickle", 666))
	test.Render() // print each line

}
