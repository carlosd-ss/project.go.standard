// Go Api server
// @jeffotoni
// 2019-01-21

package util

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func SetCount(store string, sm *sync.Map) {

	valor := GetCount(store, sm)
	valor = valor + 1
	sm.Store(store, valor)
}

func GetCount(store string, sm *sync.Map) int {

	result, _ := sm.Load(store)
	x := fmt.Sprintf("%d", result)
	y, _ := strconv.Atoi(x)
	//fmt.Println("result:", y)
	return y
}

func Getwd(level int) string {

	pathNow, _ := os.Getwd()

	pathV := strings.Split(pathNow, "/")

	pathNew := pathV[:len(pathV)-level]

	pathNewOrg := strings.Join(pathNew, "/")

	return pathNewOrg
}
