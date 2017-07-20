package main

import (
	"fmt"
	"testing"
	"github.com/tj/assert"
)

var (
	packageName = "main"
	imports = []string{"fmt"}
	mainCode = []string{"fmt.Println(0)"}
)

func TestGenPackageName(t *testing.T) {
	s := GenPackageName(packageName)
	fmt.Println(s)
	assert.Equal(t, "package main", s)
}


func TestGenImports(t *testing.T) {
	s := GenImports(imports)
	fmt.Println(s)
	assert.Equal(t, "import (\n\t\"fmt\"\n)", s)
}


func TestGenMainCode(t *testing.T) {
	s := GenMainCode(mainCode)
	fmt.Println(s)
	assert.Equal(t, "func main() {\n\tfmt.Println(0)\n}", s)
}


func TestGen(t *testing.T) {
	gocode := Gocode{}
	gocode.Form(packageName, imports, mainCode)
	fmt.Print(gocode)
	s := gocode.Gen()
	fmt.Println(s)
	assert.Equal(t, "package main\nimport (\n\t\"fmt\"\n)\nfunc main() {\n\tfmt.Println(0)\n}", s)
}
