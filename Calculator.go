package main

import (
	"encoding/json"
	"path/filepath"
	"sync"
)

type ICalculator interface {
	Calc(path string, bytes []byte) string
	GetSameList() map[string][]string
}

type Calculator struct {
	// md5 : 文件path[]
	results map[string][]string
	// md5: 文件path[]
	SameList map[string][]string
	wMutex   sync.Mutex
}

func GetResult(calculators map[string]ICalculator) string {
	for _, calculator := range calculators {
		sameList := calculator.GetSameList()
		jStr, err := json.Marshal(sameList)
		if err == nil {
			return string(jStr)
		}
	}
	return "{}"
}

func (this *Calculator) hasSame(path, c string) {
	this.wMutex.Lock()
	r := this.results[c]
	p, _ := filepath.Rel(AppConfig.Folder, path)
	if r == nil {
		r = []string{p}
	} else {
		r = append(r, p)
		// 有重复的
		this.SameList[c] = r
	}
	this.results[c] = r
	this.wMutex.Unlock()
}
