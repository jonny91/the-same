package main

import (
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"sync"
)

var (
	wg             sync.WaitGroup
	calculatorList map[string]ICalculator
)

func main() {
	rootPath, _ := os.Getwd()
	ReadConfig()
	prepareCalculator()
	if !path.IsAbs(AppConfig.Folder) {
		AppConfig.Folder = path.Join(rootPath, AppConfig.Folder)
	}

	err := filepath.Walk(AppConfig.Folder, func(path string, info fs.FileInfo, err error) error {
		go func() {
			wg.Add(1)
			defer wg.Done()

			if !info.IsDir() {
				bytes, err := os.ReadFile(path)
				if err != nil {
					fmt.Println(err)
					return
				}

				for _, calc := range calculatorList {
					calc.Calc(path, bytes)
				}
			}
		}()
		return nil
	})
	if err != nil {
		panic(err)
	}
	wg.Wait()

	result := GetResult(calculatorList)
	if result == "{}" {
		fmt.Println("没有重复内容")
	} else {
		fmt.Println(result)
	}
}

func prepareCalculator() {
	calculatorList = map[string]ICalculator{}
	if AppConfig.Mode == 1 {
		calculatorList["sha1"] = NewSha1Calc()
	} else {
		calculatorList["md5"] = NewMd5Calc()
	}
}
