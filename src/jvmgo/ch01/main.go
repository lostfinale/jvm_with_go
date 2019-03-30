package main

import "fmt"


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
	fmt.Printf("classpath:%s class:%s args:%v\n", cmd.cpOption, cmd.class, cmd.args)
}
