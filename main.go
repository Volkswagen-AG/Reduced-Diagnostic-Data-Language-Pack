package main

import (
    "fmt"
	"time"
    "io/ioutil"
    "os"
	"strings"
)
// Define need to delete language array
// -tr_TR_ = 土耳其语
// -sv_SE_ = 瑞典语
// -sl_SI_ = 斯洛文尼亚语
// -ru_RU_ = 俄罗斯语
// -ro_RO_ = 罗马尼亚语
// -pt_PT_ = 葡萄牙语
// -pt_BR_ = 葡萄牙语（巴西）
// -pl_PL_ = 波兰语
// -nl_NL_ = 荷兰语
// -ja_JP_ = 日语
// -ko_KR_ = 韩语（朝鲜语）
// -it_IT_ = 意大利语
// -hu_HU_ = 匈牙利语
// -hr_HR_ = 克罗地亚语
// -fr_FR  = 法语
// -fi_FI_ = 芬兰语
// -es_ES_ = 西班牙语
// -en_GB_ = 英语（英国)
// -el_GR_ = 希腊语
// -da_DK_ = 丹麦语
// -cs_CZ_ = 捷克语
var needDeleteLanguages = [...]string{"-tr_TR_","-sv_SE_","-sl_SI_","-ru_RU_","-ro_RO_","-pt_PT_","-pt_BR_","-pl_PL_","-nl_NL_","-ja_JP_","-ko_KR_","-it_IT_","-hu_HU_","-hr_HR_","-fr_FR_","-fi_FI_","-es_ES_","-en_GB_","-el_GR_","-da_DK_","-cs_CZ_"}
var needDeleteLanguagesUpper = [...]string{"-TR_TR_","-SV_SE_","-SL_SI_","-RU_RU_","-RO_RO_","-PT_PT_","-PT_BR_","-PL_PL_","-NL_NL_","-JA_JP_","-KO_KR_","-IT_IT_","-HU_HU_","-HR_HR_","-FR_FR_","-FI_FI_","-ES_ES_","-EN_GB_","-EL_GR_","-DA_DK_","-CS_CZ_"}
var needDeleleLanguageFiles = [...]string{"language.tr_TR_","language.sv_SE_","language.sl_SI_","language.ru_RU_","language.ro_RO_","language.pt_PT_","language.pt_BR_","language.pl_PL_","language.nl_NL_","language.ja_JP_","language.ko_KR_","language.it_IT_","language.hu_HU_","language.hr_HR_","language.fr_FR_","language.fi_FI_","language.es_ES_","language.en_GB_","language.el_GR_","language.da_DK_","language.cs_CZ_"}

func getCurrentDir(path string) {
    infos, err := ioutil.ReadDir(path)
	if err!= nil {
        fmt.Println(err)
    }
    for _, info := range infos {
		if (info.IsDir()) {
			fmt.Printf("当前是目录, 相对路径是: %s/%s\n", path,info.Name())
			getCurrentDir(path + "/" + info.Name())
		} else {
			for _, s := range needDeleteLanguages {
				if (strings.Contains(info.Name(), s)) {
					fmt.Printf("正在删除需要精简的语言包文件, 相对路径是: %s/%s\n", path, info.Name())
					err := os.Remove(path+"/"+info.Name())
					if err!= nil {
						fmt.Printf("文件删除失败，当前文件的相对路径为: %s/%s\n", path, info.Name())
					}
				}
			}
			for _, ss := range needDeleteLanguagesUpper {
				if (strings.Contains(info.Name(), ss)) {
					fmt.Printf("正在删除需要精简的语言包文件, 相对路径是: %s/%s\n", path, info.Name())
					err := os.Remove(path+"/"+info.Name())
					if err!= nil {
						fmt.Printf("文件删除失败，当前文件的相对路径为: %s/%s\n", path, info.Name())
					}
				}
			}
			for _, f := range needDeleleLanguageFiles {
				if (strings.Contains(info.Name(), f)) {
					fmt.Printf("正在删除需要精简的语言包文件, 相对路径是: %s/%s\n", path, info.Name())
					err := os.Remove(path+"/"+info.Name())
					if err!= nil {
						fmt.Printf("文件删除失败，当前文件的相对路径为: %s/%s\n", path, info.Name())
					}
				}
			}
		}
	}
}
func main() {
	fmt.Printf("欢迎使用ODIS/ODIS-E Postsetup语言包精简工具,程序将会在20秒后自动开始\n")
	fmt.Printf("本程序将会自动精简掉Postsetup数据中的以下语言:\n")
	fmt.Printf("1: 土耳其语 \n")
	fmt.Printf("2: 瑞典语 \n")
	fmt.Printf("3: 斯洛文尼亚语 \n")
	fmt.Printf("4: 俄罗斯语 \n")
	fmt.Printf("5: 罗马尼亚语 \n")
	fmt.Printf("6: 葡萄牙语 \n")
	fmt.Printf("7: 葡萄牙语（巴西）\n")
	fmt.Printf("8: 波兰语 \n")
	fmt.Printf("9: 荷兰语 \n")
	fmt.Printf("10: 日语 \n")
	fmt.Printf("11: 韩语（朝鲜语）\n")
	fmt.Printf("12: 意大利语 \n")
	fmt.Printf("13: 匈牙利语 \n")
	fmt.Printf("14: 克罗地亚语 \n")
	fmt.Printf("15: 法语 \n")
	fmt.Printf("16: 芬兰语 \n")
	fmt.Printf("17: 西班牙语 \n")
	fmt.Printf("18: 英语（英国) \n")
	fmt.Printf("19: 希腊语 \n")
	fmt.Printf("20: 丹麦语 \n")
	fmt.Printf("21: 捷克语 \n")
	fmt.Printf("如果你还需要精简其他语言,你可以自行修改,本程序是开源的项目,项目代码托管在Github上,代码是由 西安小马哥（VAGXMG）开发 \n")
	fmt.Printf("源代码: https://github.com/Volkswagen-AG/Reduced-Diagnostic-Data-Language-Pack \n")
	time.Sleep(20 * time.Second)
    getCurrentDir(".")
	fmt.Printf("程序已经精简完成，可以检查一下Postsetup数据包的体积啦,30秒后程序会自动退出\n")
	time.Sleep(30 * time.Second)
}