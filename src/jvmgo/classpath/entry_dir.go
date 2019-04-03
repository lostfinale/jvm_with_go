package classpath

import (
	"path/filepath"
	"io/ioutil"
)

//文件夹入口，不支子文件查询
type DirEntry struct {
	absDir string
}

func newDirEntry(path string) *DirEntry{
	//获取绝对路径
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir}
}

func (de *DirEntry) readClass(className string) ([]byte, Entry, error) {

	fileName := filepath.Join(de.absDir, className)
	data, err := ioutil.ReadFile(fileName)
	return data, de, err

}


func (de *DirEntry) String() string {
	return de.absDir
}



