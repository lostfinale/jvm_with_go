package main

import (
	"fmt"
	"jvmgo/ch03/classpath"
	"strings"
	"jvmgo/ch03/classfile"
)


func main() {

	cm := parseCmd()
	if cm.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cm.helpFlag|| cm.class == "" {
		printUsage()
	} else {
		startJVM(cm)
	}

}

func startJVM(cmd *Cmd) {

	cp := classpath.Parse(cmd.XjareOption, cmd.cpOption)
	fmt.Printf("classpath:%s class:%s args:%v\n", cmd.cpOption, cmd.class, cmd.args)
	className := strings.Replace(cmd.class, ".", "/", -1)

	cf := loadClass(className, cp)

	fmt.Println(cmd.class)
	printClassInfo(cf)
}

func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {

	classData, _, err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}

	cf, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}

	return cf
}

func printClassInfo(cf *classfile.ClassFile) {
	fmt.Printf("version: %v.%v\n", cf.MajorVersion(), cf.MinorVersion())
	fmt.Printf("constants count: %v\n", len(cf.ConstantPool()))

	fmt.Printf("acces flags: 0x%x\n", cf.AccessFlags())
	fmt.Printf("this class: %v\n", cf.ClassName())
	fmt.Printf("super class: %v\n", cf.SuperClassName())
	fmt.Printf("interfaces: %v\n", cf.InterfaceName())
	fmt.Printf("fields count: %v\n", len(cf.Fields()))
	for _, f := range cf.Fields() {
		fmt.Printf("    %s\n", f.Name())
	}

	fmt.Printf("method coutn: %v\n", len(cf.Methods()))

	for _, m := range cf.Methods() {
		fmt.Printf("    %s\n", m.Name())
	}
}