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
		logger.Println("Data directory is not exists, create now.")
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
		path := filepath.Join(base, name)
		logger.Println("Loading provinces file: ", path)
		downloadIfNotExists(path, name)
		loadFileToMemory(path, gProvinceNames)
	}
}

func loadCities(base string, names ...string) {
	for _, name := range names {
		path := filepath.Join(base, name)
		logger.Println("Loading cities file: ", path)
		downloadIfNotExists(path, name)
		loadFileToMemory(path, gCitiesNames)
	}
}

func loadFileToMemory(file string, destMap map[string]string) {
	fields, err := ReadRecords(file)
	if nil != err {
		panic(err)
	}

	for _, field := range fields {
		destMap[field.Key] = field.Value
	}
}

func downloadIfNotExists(path string, name string) {
	// 检查文件是否存在
	if !fileExists(path) {
		logger.Println("Data file is not exist, download from github server. file:", name)
		// 如果不存在，从GigHub中下载
		resp, he := http.Get(fmt.Sprintf("https://raw.githubusercontent.com/parkingwang/go-vna/master/data/%s", name))
		if nil != he {
			logger.Println("Cannot download data file from github server:", name)
			panic(he)
		}

		f, fe := os.Create(path)
		if nil != fe {
			logger.Println("Cannot create data file:", path)
			panic(fe)
		}

		io.Copy(f, resp.Body)
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
