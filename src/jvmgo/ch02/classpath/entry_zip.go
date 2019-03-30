package classpath

import (
	"archive/zip"
	"io/ioutil"
	"errors"
	"path/filepath"
)

//Zip或者Jar文件入口
type ZipEntry struct {

	absPath string

}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath}
}

func (ze *ZipEntry) readClass(className string) ([]byte, Entry, error) {

	rc, err := zip.OpenReader(ze.absPath)
	if err != nil {
		return nil , nil, err
	}
	defer rc.Close()


	//找到对应的文件
	classFile := ze.findClass(className, rc)

	if classFile == nil {
		return nil , nil, errors.New("class not found:" + className)
	}

	//读取数据
	data, err := readClass(classFile)

	return data, ze, err

}

func (ze *ZipEntry) String() string {
	return ze.absPath
}


func (ze *ZipEntry) findClass(className string, zipRc *zip.ReadCloser) *zip.File {
	for _, f := range zipRc.File {
		if f.Name == className {
			return f
		}
	}
	return nil
}

func readClass(classFile *zip.File) ([]byte, error) {

	rc, err := classFile.Open()
	if err != nil {
		return nil, err
	}

	defer  rc.Close()

	data, err := ioutil.ReadAll(rc)


	if err != nil {
		return nil , err
	}

	return data, nil
}
