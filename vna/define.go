package vna

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

const (
	VNumTypeCivil       = iota // 普通民用车牌
	VNumTypeNewEnergy   = iota // 新能源车牌
	VNumTypePolice      = iota // 警察车牌
	VNumTypeWJ2012      = iota // 武警2012式车牌
	VNumTypeHKMacao     = iota // 港澳车辆
	VNumTypeAviation    = iota // 民航车牌
	VNumTypeConsulate   = iota // 领馆政府车牌
	VNumTypeOldEmbassy  = iota // 使馆车牌
	VNumTypeEmbassy     = iota // 新式使馆车牌
	VNumTypePLA2012     = iota // 军队车辆
	VNumTypeUnknown     = iota // 未知车牌
)
