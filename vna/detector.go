package vna

import (
	"errors"
	"fmt"
	"strings"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

// 指定车牌号码，返回归属地分析结果
func DetectNumber(number string) (DetectedResult, error) {
	number = strings.ToUpper(number)
	numberRune := []rune(number)
	numType, numTypeName := detectNumberType(numberRune, number)

	if VNumTypeUnknown == numType {
		return DetectedResult{}, errors.New(fmt.Sprintf("unknown number[%s]", number))
	}
	provinceShort, provinceName, cityShort, cityName := detectSpecChars(numType, numberRune)

	// 统计易错字符
	var fallRate float32
	for _, c := range numberRune {
		fallRate += gFallRateNames[string(c)]
	}

	return DetectedResult{
		Number:         number,
		NumberType:     numType,
		NumberTypeName: numTypeName,
		ProvinceName:   provinceName,
		ProvinceKey:    provinceShort,
		CityName:       cityName,
		CityKey:        cityShort,
		FallRate:       fallRate,
	}, nil
}

// 返回车牌号码类型及类型名称
func detectNumberType(numberRune []rune, numberStr string) (int, string) {
	numSize := len(numberRune)
	if !(7 == numSize || 8 == numSize) {
		return VNumTypeUnknown, "UNKNOWN"
	} else if strings.ContainsAny("VZHKEBSLJNGCQ", string(numberRune[:1])) {
		return VNumTypePLA2012, "PLA2012"
	} else if starts(numberStr, "使") {
		return VNumTypeOldEmbassy, "OLD_EMBASSY"
	} else if ends(numberStr, "使") {
		return VNumTypeEmbassy, "EMBASSY"
	} else if ends(numberStr, "领") {
		return VNumTypeConsulate, "CONSULATE"
	} else if starts(numberStr, "民航") {
		return VNumTypeAviation, "AVIATION"
	} else if starts(numberStr, "WJ") {
		return VNumTypeWJ2012, "WJ2012"
	} else if ends(numberStr, "警") {
		return VNumTypePolice, "POLICE"
	} else if starts(numberStr, "粤Z") && (ends(numberStr, "港") || ends(numberStr, "澳")) {
		return VNumTypeHKMacao, "HK_MACAO"
	} else {
		// 新能源车牌长度为8位
		if 8 == numSize {
			return VNumTypeNewEnergy, "NEW_ENERGY"
		} else {
			return VNumTypeCivil, "CIVIL"
		}
	}
}

// 分析车牌号码的省份和城市关键字符
func detectSpecChars(numType int, number []rune) (provKey string, provName string, cityKey string, cityName string) {
	switch numType {
	case VNumTypeWJ2012:
		// 武警： WJ-粤-1234X
		provKey = string(number[2:3])
		cityKey = string(number[7:8])

	case VNumTypeEmbassy:
		provKey = "使"
		cityKey = string(number[:3])

	case VNumTypeOldEmbassy:
		provKey = "使"
		cityKey = string(number[1:4])

	case VNumTypeAviation:
		provKey = "航"
		cityKey = "航"

	case VNumTypePLA2012:
		fallthrough
	case VNumTypePolice:
		fallthrough
	case VNumTypeConsulate:
		fallthrough
	case VNumTypeHKMacao:
		fallthrough
	case VNumTypeNewEnergy:
		fallthrough
	case VNumTypeCivil:
		fallthrough
	default:
		// 第一位：省份/军区
		// 第二位：城市/部队
		provKey = string(number[:1])
		cityKey = string(number[:2])
	}

	provName = gProvinceNames[provKey]
	cityName = gCitiesNames[cityKey]

	return
}
