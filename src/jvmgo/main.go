package main

import (
	"fmt"
	"jvmgo/classpath"
	"strings"
	"jvmgo/rtda/heap"
)

func main() {

	cm := parseCmd()
	if cm.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cm.helpFlag || cm.class == "" {
		printUsage()
	} else {
		startJVM(cm)
	}

}

func startJVM(cmd *Cmd) {
	//解析classpath
	cp := classpath.Parse(cmd.XjareOption, cmd.cpOption)
	className := strings.Replace(cmd.class, ".", "/", -1)

	classLoader := heap.NewClassLoader(cp, cmd.verboseClassFlag)
	mainClass := classLoader.LoadClass(className)
	mainMethod := mainClass.GetMainMethod()
	if mainMethod != nil {
		//执行main函数
		interpret(mainMethod, cmd.verboseInstFlag, cmd.args)
	} else {
		fmt.Printf("Main method not found in class %s\n", cmd.class)
	}

}

