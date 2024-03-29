# Reduced-Diagnostic-Data-Language-Pack
精简大众ODIS ODIS-E语言数据包

-tr_TR_
-sv_SE_
-sl_SI_
-ru_RU_
-ro_RO_
-pt_PT_
-pt_BR_
-pl_PL_
-nl_NL_
-ko_KR_
-ja_JP_
-it_IT_
-hu_HU_
-hr_HR_
-fr_FR_
-fi_FI_
-es_ES_
-en_US_
-en_GB_
-el_GR_
-de_DE_
-da_DK_
-cs_CZ_
-zh_CN_

language.cs_CZ_

#### 交叉编译生成最后的文件
```shell
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go && mv main Reduced-Diagnostic-Data-Language-Pack_linux_amd64
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go && mv main.exe Reduced-Diagnostic-Data-Language-Pack_windows_amd64.exe
CGO_ENABLED=0 GOOS=windows GOARCH=386 go build main.go && mv main.exe Reduced-Diagnostic-Data-Language-Pack_windows_386.exe
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go && mv main Reduced-Diagnostic-Data-Language-Pack_darwin_amd64
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build main.go && mv main Reduced-Diagnostic-Data-Language-Pack_darwin_arm64
```


| 语种  | 英文全称 | 中文全称 | 缩写 |
|:---|:---|:---|:---|
|العربية | Arabic（Israel） | 阿拉伯语（以色列) | ar_IL |
|العربية | Arabic（Egypt) | 阿拉伯语（埃及) | ar_EG |
|中文     | Chinese Simplified | 中文简体  | zh_CN  |
|中文     | Chinese Tradition |	中文繁体|	zh_TW|
|中文     | Chinese	| 中文（香港) | zh_HK |
|Nederlands | Dutch （Netherlands)|	荷兰语|	nl_NL|
|Nederlands | Dutch （Netherlands)|	荷兰语（比利时)|nl_BE|
|English|	English（United States)|	英语（美国)|	en_US|
|English|	English（Australia) |	英语（澳大利亚)|	en_AU|
|English|	English（Canada) |	英语（加拿大)|	en_CA|
|English|	English（India) |	英语（印度)|	en_IN|
|English|	English（Ireland) |	英语（爱尔兰)|	en_IE|
|English|	English（New Zealand) |	英语（新西兰)|	en_NZ|
|English|	English（Singapore) |	英语（新加波)|	en_SG|
|English|	English（South Africa) |	英语（南非)| en_ZA|
|English|	English（United Kingdom) |	英语（英国)| en_GB|
|Français|	French|	法语|	fr_FR|
|Français|	French|	法语（比利时)|	fr_BE|
|Français|	French|	法语（加拿大)|	fr_CA|
|Français|	French|	法语（瑞士)|fr_CH|
|Deutsch|	German|	德语|	de_DE|
|Deutsch|	German|	德语（列支敦斯登)|	de_LI|
|Deutsch|	German|	德语（奥地利)|	de_AT|
|Deutsch|	German|	德语（瑞士)|	de_CH|
|Italiano|	Italian|	意大利语|	it_IT|
|Italiano|	Italian|	意大利语（瑞士)|	it_CH|
|Protuguês|	Portuguese|	葡萄牙语（巴西）|	pt_BR|
|Protuguês|	Portuguese|	葡萄牙语|	pt_PT|
|Español|	Spanish|	西班牙语|	es_ES|
|Español|	Spanish|	西班牙语（美国)|	es_US|
|বাংলা|	Bengali|	孟加拉语|	bn_BD|
|বাংলা|	Bengali|	孟加拉语（印度)|	bn_IN|
|hrvatski|	Croatian|	克罗地亚语|	hr_HR|
|čeština|	Czech|	捷克语|	cs_CZ|
|Dansk|	Danish|	丹麦语|	da_DK|
|ελληνικά|	Greek|	希腊语|	el_GR|
|עברית|	Hebrew|	希伯来语（以色列)|	he_IL|
|עברית|	Hebrew|	希伯来语（以色列)|	iw_IL|



हिंदी	Hindi	印度语	hi_IN
|Magyar	Hungarian	匈牙利语	hu_HU
|Indonesian	印度尼西亚语	in_ID
|日本語の言語	Japanese	日语	ja_JP
|한국의	Korean	韩语（朝鲜语）	ko_KR
|Bahasa Melayu	Malay	马来语	ms_MY
فارسی	Perisan	波斯语	fa_IR
Polski	Polish	波兰语	pl_PL
româna	Romanian	罗马尼亚语	ro_RO
Русский	Russian	俄罗斯语	ru_RU
српски	Serbian	塞尔维亚语	sr_RS
Svenska	Swedish	瑞典语	sv_SE
ไทย	Thai	泰语	th_TH
Türkçe	Turkey	土耳其语	tr_TR
اردو	Urdu	乌尔都语	ur_PK
tiếng việt	Vietnamese	越南语	vi_VN
catalá	Catalan	加泰隆语（西班牙)	ca_ES
latviešu	Latviesu	拉脱维亚语	lv_LV
Lietuvių	Lithuanian	立陶宛语	lt_LT
Norsk bokmal	Norwegian	挪威语	nb_NO
Slovenčina	slovencina	斯洛伐克语	sk_SK
Slovenščina	Slovenian	斯洛文尼亚语	sl_SI
български	bulgarian	保加利亚语	bg_BG
українська	Ukrainian	乌克兰语	uk_UA
Tagalog	Filipino	菲律宾语	tl_PH
Suomi	Finnish	芬兰语	fi_FI
Afrikaans	Afrikaans	南非语	af_ZA
Rumantsch	Romansh	罗曼什语（瑞士)	rm_CH
ဗမာ	Burmese（Zawgyi)	缅甸语（民间)	my_ZG
ဗမာ	Burmese（Paduak)	缅甸语（官方)	my_MM
ខ្មែរ	Khmer	柬埔寨语	km_KH
አማርኛ	Amharic	阿姆哈拉语（埃塞俄比亚)	am_ET
беларуская	Belarusian	白俄罗斯语	be_BY
eesti	Estonian	爱沙尼亚语	et_EE
Kiswahili	Swahili	斯瓦希里语（坦桑尼亚)	sw_TZ
isiZulu	Zulu	祖鲁语（南非)	zu_ZA
azərbaycanca	Azerbaijani	阿塞拜疆语	az_AZ
Հայերէն	Armenian	亚美尼亚语（亚美尼亚)	hy_AM
ქართული	Georgian	格鲁吉亚语（格鲁吉亚)	ka_GE
ລາວ	Laotian	老挝语（老挝)	lo_LA
Монгол	Mongolian	蒙古语	mn_MN
नेपाली	Nepali	尼泊尔语	ne_NP
қазақ тілі	Kazakh	哈萨克语	kk_KZ
Galego	Galician	加利西亚语	 gl-rES
íslenska	 Icelandic	冰岛语	is-rIS
ಕನ್ನಡ	Kannada	坎纳达语	kn-rIN
кыргыз тили; قىرعىز تىلى	Kyrgyz	吉尔吉斯语	ky-rKG
മലയാളം	Malayalam	马拉亚拉姆语	ml-rIN
मराठी	Marathi	马拉提语/马拉地语	 mr-rIN
தமிழ்	Tamil	泰米尔语	ta-rIN
македонски јазик	Macedonian	马其顿语	mk-rMK
తెలుగు	Telugu	泰卢固语	te-rIN
Ўзбек тили	Uzbek	乌兹别克语	uz-rUZ
Euskara	Basque	巴斯克语	eu-rES
සිංහල	Sinhala	僧加罗语（斯里兰卡)	si_LK