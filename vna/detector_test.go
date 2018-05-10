package vna

import (
	"fmt"
	"testing"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

func TestDetectNumberNewEnergy(t *testing.T) {
	InitDetectorEnv("../" + DataDirName)
	checkIs(t, "NEW_ENERGY", "粤BF49883", "广东省", "深圳市")
	checkIs(t, "NEW_ENERGY", "粤BD12345", "广东省", "深圳市")
	checkIs(t, "NEW_ENERGY", "粤B01234F", "广东省", "深圳市")
	checkIs(t, "NEW_ENERGY", "粤B01234D", "广东省", "深圳市")
	checkIs(t, "NEW_ENERGY", "京AD12345", "北京市", "中央国家机关")
	checkIs(t, "NEW_ENERGY", "京GD12345", "北京市", "北京远郊区")
}

func TestDetectNumberCivil(t *testing.T) {
	InitDetectorEnv("../" + DataDirName)
	checkIs(t, "CIVIL", "京A12345", "北京市", "中央国家机关")
	checkIs(t, "CIVIL", "新B12345", "新疆维吾尔自治区", "昌吉回族自治州、五家渠市")
	checkIs(t, "CIVIL", "粤O00001", "广东省", "省直机关")
}

func TestDetectNumberPolice(t *testing.T) {
	InitDetectorEnv("../" + DataDirName)
	checkIs(t, "POLICE", "粤A1234警", "广东省", "广州市")
}

func TestDetectNumberWJ2012(t *testing.T) {
	InitDetectorEnv("../" + DataDirName)
	checkIs(t, "WJ2012", "WJ粤1006X", "广东省", "消防部队")
	checkIs(t, "WJ2012", "WJ粤0001B", "广东省", "边防部队")
	checkIs(t, "WJ2012", "WJ粤00011", "广东省", "内卫部队")
}

func TestDetectNumberHKMacao(t *testing.T) {
	InitDetectorEnv("../" + DataDirName)
	checkIs(t, "HK_MACAO", "粤Z1234港", "广东省", "港澳车辆")
	checkIs(t, "HK_MACAO", "粤Z1234澳", "广东省", "港澳车辆")
}

func TestDetectNumberAviation(t *testing.T) {
	InitDetectorEnv("../" + DataDirName)
	checkIs(t, "AVIATION", "民航B9016", "民航", "民航")
	checkIs(t, "AVIATION", "民航F0441", "民航", "民航")
}

func TestDetectNumberConsulate(t *testing.T) {
	InitDetectorEnv("../" + DataDirName)
	checkIs(t, "CONSULATE", "粤A1006领", "广东省", "广州市")
}

func TestDetectNumberPLA2012(t *testing.T) {
	InitDetectorEnv("../" + DataDirName)
	checkIs(t, "PLA2012", "ZB00221", "中央军委", "总政治部二级部")
	checkIs(t, "PLA2012", "NB02151", "南京军区", "政治部")
	checkIs(t, "PLA2012", "BA09007", "北京军区", "司令部")
}

func TestDetectNumberEmbassy(t *testing.T) {
	InitDetectorEnv("../" + DataDirName)
	checkIs(t, "OLD_EMBASSY", "使189001", "大使馆", "巴勒斯坦")
	checkIs(t, "EMBASSY", "189001使", "大使馆", "巴勒斯坦")
	checkIs(t, "EMBASSY", "238001使", "大使馆", "土库曼斯坦")
}

func checkIs(t *testing.T, numType string, number string, province string, city string) {
	dr, err := DetectNumber(number)
	if nil != err {
		t.Error(err)
	}
	fmt.Println("Testing:", number, dr)
	if numType != dr.NumberTypeName {
		t.Errorf("Number type not match, should: %s, was %s", numType, dr.NumberTypeName)
	}
	if province != dr.ProvinceName {
		t.Errorf("PROVINCE not match, should: %s, was: %s", province, dr.ProvinceName)
	}
	if city != dr.CityName {
		t.Errorf("CITY not match, should: %s, was: %s", city, dr.CityName)
	}
}
