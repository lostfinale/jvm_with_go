package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

//通配符入口
func newWildcardEntry(path string) CompositeEntry{
	baseDir := path[:len(path) -1] //remove *
	compositeEntry := CompositeEntry{}

	//walk函数
	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}
		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			jarEntry := newZipEntry(path)
			compositeEntry = append(compositeEntry, jarEntry)
		}
		return nil
	}
	//遍历路径然后执行walk函数
	filepath.Walk(baseDir, walkFn)
	return compositeEntry

}