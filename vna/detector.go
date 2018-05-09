package vna

import (
	"errors"
	"fmt"
	"strings"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

type DetectedResult struct {
	Number         string // 车牌号码
	NumberType     int    // 车牌号码类型
	NumberTypeName string // 车牌号码类型名称
	ProvinceName   string // 所属省份全称
	ProvinceKey    string // 所属省份查询Key
	CityName       string // 所属城市全称
	CityKey        string // 所属城市查询Key
}

func (dr DetectedResult) String() string {
	return fmt.Sprintf(`number: %s, type: %d, type_name: %s, province: %s, province_name: %s, city: %s, city_name:%s`,
		dr.Number, dr.NumberType, dr.NumberTypeName,
		dr.ProvinceKey, dr.ProvinceName, dr.CityKey, dr.CityName)
}

///

// 指定车牌号码，返回归属地分析结果
func DetectNumber(number string) (DetectedResult, error) {
	numType, numTypeName := DetectNumberType(number)

	if VNumTypeUnknown == numType {
		return DetectedResult{}, errors.New(fmt.Sprintf("unknown number[%s]", number))
	}

	provinceShort, provinceName, cityShort, cityName := DetectSpecChars(numType, number)

	return DetectedResult{
		Number:         number,
		NumberType:     numType,
		NumberTypeName: numTypeName,
		ProvinceName:   provinceName,
		ProvinceKey:    provinceShort,
		CityName:       cityName,
		CityKey:        cityShort,
	}, nil
}

// 返回车牌号码类型及类型名称
func DetectNumberType(numberStr string) (int, string) {
	numberStr = strings.ToUpper(numberStr)
	numberRune := []rune(numberStr)
	numSize := sizeOf(numberStr)
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
func DetectSpecChars(numType int, numberS string) (provKey string, provName string, cityKey string, cityName string) {
	number := []rune(numberS)
	switch numType {
	case VNumTypeWJ2012:
		// 武警： WJ-粤-1234X
		provKey = string(number[2:3])
		cityKey = string(number[7:8])

	case VNumTypePLA2012:
		// KA·00001
		provKey = string(number[:1])
		// V字头的二级单位Key需要3位
		if "V" == provKey {
			cityKey = string(number[:3])
		} else {
			cityKey = string(number[:2])
		}

	case VNumTypeEmbassy:
		provKey = "使"
		cityKey = string(number[:3])

	case VNumTypeOldEmbassy:
		provKey = "使"
		cityKey = string(number[1:4])

	case VNumTypeAviation:
		provKey = "航"
		cityKey = "航"

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
		// 第一位：省份
		// 第二位：城市
		provKey = string(number[:1])
		cityKey = string(number[:2])
	}

	provName = gProvinceNames[provKey]
	cityName = gCitiesNames[cityKey]

	return
}
