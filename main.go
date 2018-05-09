package main

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

import (
	"fmt"
	"github.com/go-vna/vna"
)

func main() {

	vna.InitDetectorEnv("./data/")

	fmt.Println(vna.DetectNumber("粤BF49883"))
	fmt.Println(vna.DetectNumber("粤A12345"))
	fmt.Println(vna.DetectNumber("赣AD9999"))
	fmt.Println(vna.DetectNumber("赣AD999警"))
	fmt.Println(vna.DetectNumber("贵O11111"))
	fmt.Println(vna.DetectNumber("KA20003"))
	fmt.Println(vna.DetectNumber("VA20003"))
	fmt.Println(vna.DetectNumber("WJ粤7710B"))
	fmt.Println(vna.DetectNumber("WJ云01026"))
	fmt.Println(vna.DetectNumber("WJ云0102X"))
}
