# 中国车牌号码归属地分析库- Golang

> 根据给定完整的车牌号码，分析并返回此号码所属的归属地/部门等信息。

## 支持分析号码类型

- 全国所有省份、城市级别的民用车牌号码；
- 全世界的驻华大使馆；
- 2012式武警车牌；
- 2012式军队车牌；

## Install

```bash
go get github.com/parkingwang/go-vna
```

## 使用

```go
vna.InitDetectorEnv("./data/")

result, err := vna.DetectNumber("粤BF49883")
if nil != err {
    panic(err)
}

fmt.Println(result)
```

返回的输出结果会有如下数据（示例）：

```
number: 粤BF49883, type: 1, type_name: NEW_ENERGY, province: 粤, province_name: 广东省, city: 粤B, city_name:深圳市
number: 粤A12345, type: 0, type_name: CIVIL, province: 粤, province_name: 广东省, city: 粤A, city_name:广州市
number: 赣AD9999, type: 0, type_name: CIVIL, province: 赣, province_name: 江西省, city: 赣A, city_name:南昌市
number: 赣AD999警, type: 2, type_name: POLICE, province: 赣, province_name: 江西省, city: 赣A, city_name:南昌市
number: 贵O11111, type: 0, type_name: CIVIL, province: 贵, province_name: 贵州省, city: 贵O, city_name:
number: KA20003, type: 9, type_name: PLA2012, province: K, province_name: 空军, city: KA, city_name:司令部
number: VA20003, type: 9, type_name: PLA2012, province: V, province_name: 北京卫戍区, city: VA2, city_name:总后勤部总部
number: WJ粤7710B, type: 3, type_name: WJ2012, province: 粤, province_name: 广东省, city: B, city_name:边防部队
number: WJ云01026, type: 3, type_name: WJ2012, province: 云, province_name: 云南省, city: 6, city_name:内卫部队
number: WJ云0102X, type: 3, type_name: WJ2012, province: 云, province_name: 云南省, city: X, city_name:消防部队
```

## 分析结果字段

```go
type DetectedResult struct {
	Number         string // 车牌号码
	NumberType     int    // 车牌号码类型
	NumberTypeName string // 车牌号码类型名称
	ProvinceName   string // 所属省份全称
	ProvinceShort  string // 所属省份简称
	CityName       string // 所属城市全称
	CityShort      string // 所属城市简称
}
```

其中：

- `Number`： 当前检测的车牌号码；
- `NumberType`： 当前检测的车牌号码类型Int值；
- `NumberTypeName`： 当前检测的车牌号码名称；
- `ProvinceName`： 表示所属省份名称，如果是军队车牌则表示所属军区；
- `ProvinceKey`： 所属省份查询Key。使用此Key来检索省份名称；
- `CityName`： 表示所属城市名称名称，如果是军队车牌则表示所属部队；
- `CityKey`： 所属城市查询Key。使用此Key来检索城市名称；

## 支持检测的号码类型

- `VNumTypeCivil`[**CIVIL**]       : 普通民用车牌
- `VNumTypeNewEnergy`[**NEW_ENERGY**]   : 新能源车牌
- `VNumTypePolice`[**POLICE**]      : 警察车牌
- `VNumTypeWJ2012`[**WJ2012**]      : 武警2012式车牌
- `VNumTypeHKMacao`[**HK_MACAO**]     : 港澳车辆
- `VNumTypeAviation`[**AVIATION**]    : 民航车牌
- `VNumTypeConsulate`[**CONSULATE**]   : 领馆政府车牌
- `VNumTypeOldEmbassy`[**OLD_EMBASSY**]  : 使馆车牌
- `VNumTypeEmbassy`[**EMBASSY**]     : 新式使馆车牌
- `VNumTypePLA2012`[**PLA2012**]     : 军队车辆
- `VNumTypeUnknown`[**UNKNOWN**]     : 未知车牌
