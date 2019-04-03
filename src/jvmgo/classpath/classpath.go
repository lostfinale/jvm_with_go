package classpath

import (
	"os"
	"path/filepath"
)

type Classpath struct {
	bootClasspath Entry
	extClassPath Entry
	userClassPath Entry
}


func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}


	//启动classpath和扩展classpath
	cp.parseBootAndExtClaspath(jreOption)

	//用户classpath
	cp.pareUserClasspath(cpOption)
	return cp
}


func (cp *Classpath) parseBootAndExtClaspath(jreOption string) {
	jreDir := getJreDir(jreOption)

	// jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	cp.bootClasspath = newWildcardEntry(jreLibPath)
	// jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	cp.extClassPath = newWildcardEntry(jreExtPath)
}

func (cp *Classpath) pareUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	cp.userClassPath = newEntry(cpOption)
}


func (cp *Classpath) ReadClass(className string)([]byte, Entry, error) {
	className = className + ".class"
	if data, entry, err := cp.bootClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	if data, entry, err := cp.extClassPath.readClass(className); err == nil {
		return data, entry, err
	}
	return cp.userClassPath.readClass(className)
}

func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}

	if exists("./jre") {
		return "./jre"
	}

	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}
	panic("Can not find jre folder!")
}

func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}