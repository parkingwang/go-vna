package vna

import (
	"fmt"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

const (
	VNumTypeCivil         = iota // 普通民用车牌
	VNumTypeNewEnergy     = iota // 新能源车牌
	VNumTypePolice        = iota // 警察车牌
	VNumTypeWJ2012        = iota // 武警2012式车牌
	VNumTypeHKMacao       = iota // 港澳车辆
	VNumTypeAviation      = iota // 民航车牌
	VNumTypeConsulate2007 = iota // 领事馆政府车牌
	VNumTypeOldEmbassy    = iota // 使馆车牌
	VNumTypeEmbassy       = iota // 新式使馆车牌
	VNumTypePLA2012       = iota // 军队车辆
	VNumTypeUnknown       = iota // 未知车牌
	VNumTypeConsulate2018 = iota // 2018式领事馆车牌
)

const (
	VNumTypeNameCivil      = "CIVIL"
	VNumTypeNameNewEnergy  = "NEW_ENERGY"
	VNumTypeNamePolice     = "POLICE"
	VNumTypeNameWJ2012     = "WJ2012"
	VNumTypeNameHKMacao    = "HK_MACAO"
	VNumTypeNameAviation   = "AVIATION"
	VNumTypeNameCON2007    = "CONSULATE"
	VNumTypeNameCON2018    = "CONSULATE2018"
	VNumTypeNameOldEmbassy = "OLD_EMBASSY"
	VNumTypeNameEmbassy    = "EMBASSY"
	VNumTypeNamePLA2012    = "PLA2012"
	VNumTypeNameUnknown    = "UNKNOWN"
)

// 检测结果
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
