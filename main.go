package main

import (
	"fmt"
	"bytes"
	"github.com/tj/kingpin"
)


var (
	gen = kingpin.Command("gen", "create new go file.")
	genPackage = gen.Flag("p", "package name").String()
	genImport = gen.Flag("i", "import list ").Strings()
	genCode = gen.Arg("code", "code line by line").Strings()
)

type Gocode struct {
	packageName string
	imports []string
	mainCode []string
}

func (gocode *Gocode) Form(packageName string, imports []string, mainCode []string) {
	gocode.packageName = packageName
	gocode.imports = imports
	gocode.mainCode = mainCode
}


func GenPackageName(packageName string) string {
	return fmt.Sprintf("package %s", packageName)
}

func GenImports(imports []string) string {
	var buff bytes.Buffer
	for _,v := range imports {
		buff.WriteString("\t")
		buff.WriteString(fmt.Sprintf("\"%s\"", v))
		buff.WriteString("\n")
	}
	return fmt.Sprintf("import (\n%s)", buff.String())
}

func GenMainCode(mainCode []string) string {
	var buff bytes.Buffer
	for _,v := range mainCode {
		buff.WriteString("\t")
		buff.WriteString(v)
		buff.WriteString("\n")
	}
	return fmt.Sprintf("func main() {\n%s}", buff.String())
}

func (gocode *Gocode) Gen() (codeStr string) {
	codeStr = GenPackageName(gocode.packageName)
	codeStr += "\n"
	codeStr += GenImports(gocode.imports)
	codeStr += "\n"
	codeStr += GenMainCode(gocode.mainCode)
	return codeStr
}

func main() {
	switch kingpin.Parse() {
	case "gen":
		if *genPackage == "" { *genPackage="main" }
		gocode := Gocode{}
		gocode.Form(*genPackage, *genImport, *genCode)
		fmt.Println(gocode.Gen())
	default:
		panic("need 'new'")
	}
}
