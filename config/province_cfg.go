// 省份代码

package config

type ProvinceInfo struct {
	FullName string `json:"fullname"`
	Logogram string `json:"logogram"`
}

var ProvinceCfg []ProvinceInfo

// 省份
func parseProvinceConfig() {

}

func GetProvinceName(name string) string {
	for _, info := range ProvinceCfg {
		if info.FullName == name {
			return info.Logogram
		}
	}
	return "other"
}
