package main

import (
    "fmt"
	"time"
    "io/ioutil"
    "os"
	"strings"
)

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
			if (strings.Contains(info.Name(), "-tr_TR_") || strings.Contains(info.Name(), "-sv_SE_") || strings.Contains(info.Name(), "-sl_SI_") || strings.Contains(info.Name(), "-ru_RU_") || strings.Contains(info.Name(), "-ro_RO_") || strings.Contains(info.Name(), "-pt_PT_") || strings.Contains(info.Name(), "-pt_BR_") || strings.Contains(info.Name(), "-pl_PL_") || strings.Contains(info.Name(), "-nl_NL_") || strings.Contains(info.Name(), "-ja_JP_") || strings.Contains(info.Name(), "-ko_KR_") || strings.Contains(info.Name(), "-it_IT_") || strings.Contains(info.Name(), "-hu_HU_") || strings.Contains(info.Name(), "-hr_HR_") || strings.Contains(info.Name(), "-fr_FR_") || strings.Contains(info.Name(), "-fi_FI_") || strings.Contains(info.Name(), "-es_ES_") || strings.Contains(info.Name(), "-en_GB_") || strings.Contains(info.Name(), "-el_GR_") || strings.Contains(info.Name(), "-da_DK_") || strings.Contains(info.Name(), "-cs_CZ_")) {
				fmt.Printf("正在删除需要精简的语言包文件, 相对路径是: %s/%s\n", path, info.Name())
				err := os.Remove(path+"/"+info.Name())
				if err!= nil {
                    fmt.Printf("文件删除失败，当前文件的相对路径为: %s/%s\n", path, info.Name())
                }
			}
		}
	}
}
func main() {
	fmt.Printf("欢迎使用ODIS/ODIS-E Postsetup语言包精简工具,程序将会在10秒后自动开始\n")
	time.Sleep(10 * time.Second)
    getCurrentDir(".")
	fmt.Printf("程序已经精简完成，可以检查一下Postsetup数据包的体积啦\n")
	time.Sleep(20 * time.Second)
}