package main

import (
	"fmt"
	"jvmgo/ch02/classpath"
	"strings"
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
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Could not find or load main class %s\n", cmd.class)
	}
	fmt.Printf("class data:%v\n", classData)
}
