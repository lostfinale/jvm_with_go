package main

import (
	"fmt"
	"jvmgo/ch06/classpath"
	"jvmgo/ch06/classfile"
	"strings"
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
	//加载类
	cf := loadClass(className, cp)

	//找到main函数
	mainMethod := getMainMethod(cf)
	if mainMethod != nil {
		//执行main函数
		interpret(mainMethod)
	} else {
		fmt.Printf("Main method not found in class %s\n", cmd.class)
	}

}

func loadClass(className string, cp *classpath.Classpath) * classfile.ClassFile {
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

func getMainMethod(cf *classfile.ClassFile) *classfile.MemberInfo {
	for _, m := range cf.Methods(){
		if m.Name() == "main" && m.Descriptor() == "([Ljava/lang/String;)V" {
			return m
		}
	}
	return nil
}

