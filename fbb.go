package main

import (
	// dot preceding the import means, use this namespace for the import
	// this means functions in the imported package don't need to have the
	// directory prefixed, soo you can use Agent() rather than agents.Agent()

	. "github.com/eyetoe/foobarbaz/agents"
)

// Global SaveFile available to all code
var SaveFile string

func main() {
	// uncomment Sandbox() to run testing function before starting game
	// see sandbox.go
	//Sandbox()

	EnvSetup()
	PickChar()
	Char := Agent{File: SaveFile}
	Char.Load()

	//fmt.Println(Char.CharacterSheet())
	//Continue()

	// for testing, Resurrect if character is dead
	Resurrect(&Char)

	// This is the main loop
	Prompt()
	return
}
