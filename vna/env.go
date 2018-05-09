package vna

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
// 初始化检测环境
//

var (
	gProvinceNames = make(map[string]string)
	gCitiesNames   = make(map[string]string)
)

var (
	logger = log.New(os.Stderr, "GoVNA", log.Lshortfile)
)

// 初始化检测环境
func InitDetectorEnv(base string) {
	// 检查data目录是否存在
	if !fileExists(base) {
		logger.Println("Data directory [data] not exists, create...")
		os.MkdirAll(base, os.ModePerm)
	}

	loadProvinces(base,
		"prov-army.csv",
		"prov-civil.csv",
		"prov-spec.csv")

	loadCities(base,
		"city-civil.csv",
		"city-army.csv",
		"city-embassy.csv",
		"city-spec.csv",
		"city-wj.csv")
}

func loadProvinces(base string, names ...string) {
	for _, name := range names {
		file := filepath.Join(base, name)
		logger.Println("Loading provinces file: ", file)
		loadFileToMemory(file, name, gProvinceNames)
	}
}

func loadCities(base string, names ...string) {
	for _, name := range names {
		file := filepath.Join(base, name)
		logger.Println("Loading cities file: ", file)
		loadFileToMemory(file, name, gCitiesNames)
	}
}

func loadFileToMemory(file string, fileName string, target map[string]string) {
	// 检查文件是否存在
	if !fileExists(file) {
		logger.Println("Data file not exist, download from github server. datafile:", fileName)
		// 如果不存在，从GigHub中下载
		resp, he := http.Get(fmt.Sprintf("https://raw.githubusercontent.com/parkingwang/go-vna/master/data/%s", fileName))
		if nil != he {
			logger.Println("Cannot download data from github server:", fileName)
			panic(he)
		}

		f, fe := os.Create(file)
		if nil != fe {
			logger.Println("Cannot create data file:", file)
			panic(fe)
		}

		io.Copy(f, resp.Body)
	}
	// load provinces
	fields, err := ReadFields(file)
	if nil != err {
		panic(err)
	}

	for _, field := range fields {
		target[field.Short] = field.Name
	}
}

func fileExists(path string) bool {
	_, e := os.Stat(path)
	if nil != e && os.IsNotExist(e) {
		return false
	} else {
		return true
	}
}
