package i18n

import (
	"io/ioutil"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
)

// Lang 语言类型
type Lang string

// TextID 代表了多语言的文本标识
type TextID string

// 语言枚举
const (
	LangCN      Lang = "cn"   // 中文简体
	LangHK      Lang = "hk"   // 香港繁体
	LangTW      Lang = "tw"   // 台湾繁体
	LangEN      Lang = "en"   // 英文
	LangDefault      = LangTW // 默认语言
)

// IsCn 是否是中文
func (l Lang) IsCn() bool {
	return l == "cn"
}

// IsHk 是否是繁体中文
func (l Lang) IsHk() bool {
	return l == "hk"
}

// IsTw 是否是台湾繁体中文
func (l Lang) IsTw() bool {
	return l == "tw"
}

// IsEn 是否是英文
func (l Lang) IsEn() bool {
	return l == "en"
}

// translations 包含翻译各种语言的文本
var translations map[TextID]map[Lang]string

// Load 从文件载入翻译
func Load(filepath string) error {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return errors.WithStack(err)
	}

	return LoadData(data)
}

// LoadData 加载内容
func LoadData(data []byte) error {
	err := toml.Unmarshal(data, &translations)
	return errors.WithStack(err)
}

// Translate 翻译为指定语言的文本内容
// 使用:
//  1. 错误文本中不需要替换内容
//     Translate(tid, lang)
//  2. 错误文本中要替换指定内容
//     Translate(tid, lang, map[string]string{
//     "xxx" => "yyy",			// 它会将文本中的"{xxx}" 替换为 "yyy"
//     ...
//     })
func Translate(id TextID, to Lang, replacement ...map[string]string) string {
	text := ""
	if lang, ok := translations[id]; ok {
		if transl, ok := lang[to]; ok {
			text = transl
		} else if transl, ok := lang[LangDefault]; ok {
			text = transl
		} else {
			return string(id)
		}
	} else {
		return string(id)
	}

	if len(replacement) > 0 && len(replacement[0]) > 0 {
		oldnew := make([]string, 0, len(replacement[0])*2)
		for key, val := range replacement[0] {
			oldnew = append(oldnew, "{"+key+"}", val)
		}
		text = strings.NewReplacer(oldnew...).Replace(text)
	}

	return text
}

// NormalizeLanguageCode 正规化语言代码
func NormalizeLanguageCode(code string) Lang {
	l := strings.ToLower(code)
	if strings.Contains(code, "_") {
		l = strings.Split(code, "_")[1]
	} else if strings.Contains(code, "-") {
		l = strings.Split(code, "-")[1]
	}

	// 统一大小写格式 全部转换小写
	l = strings.ToLower(l)

	validLang := map[Lang]bool{
		LangCN: true,
		LangHK: true,
		LangTW: true,
		LangEN: true,
	}
	if v, _ := validLang[Lang(l)]; !v {
		if strings.HasPrefix(code, "en") {
			// 兼容cli-app-lang 上传英文地区的情况，例如en-US
			return LangEN
		}
	}

	return Lang(l)
}

// DateFormat 得到一个多语言的日期格式
func DateFormat(l Lang) string {
	return Translate(TextID("date_format"), l)
}
