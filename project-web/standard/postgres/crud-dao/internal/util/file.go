// Go Api server
// @jeffotoni
// 2019-05-14

package util

import (
	"io/ioutil"
	"os"
	"path"
	"strings"
)

//  Exist only file
func FileExist(name string) bool {
	if stat, err := os.Stat(name); err == nil && !stat.IsDir() {
		return true
	}
	return false
}

func DirExist(path string) bool {
	if stat, err := os.Stat(path); err == nil && stat.IsDir() {
		return true
	}
	return false
}

func CreateDirIfNotExist(dir string) bool {
	if !DirExist(dir) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return false
		}
	}
	return true
}

func CreateDirIfNotExistNotFile(dir string) {
	vet := strings.Split(dir, "/")
	tamanho_i := len(vet)
	tamanho_f := tamanho_i - 1
	var dirNew string
	if tamanho_i > 1 {
		dirNew = strings.Join(vet[0:tamanho_f], "/")
		//criar diretorio ou nao
		if !DirExist(dirNew) {
			os.MkdirAll(dirNew, 0755)
		}
	}
}

func RemoveAllDir(pathlocal string) {
	dir, _ := ioutil.ReadDir(pathlocal)
	for _, d := range dir {
		os.RemoveAll(path.Join([]string{pathlocal, d.Name()}...))
	}
}

func DeleteFile(path string) {
	os.Remove(path)
}
