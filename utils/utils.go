package utils

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/szwtdl/aqscwlxy/types"
	"io"
	"math"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

func GetUuid(platformType string) string {
	if platformType == "1" {
		// 当前时间戳（毫秒） + 随机数
		timestamp := time.Now().UnixMilli()
		randomNum := rand.Intn(10_000_000) // 随机数范围 [0, 10_000_000)
		return fmt.Sprintf("%d%d", timestamp, randomNum)
	}
	// 生成标准 UUID
	return uuid.New().String()
}

// 转换字符串

func ProcessString(input string) string {
	switch input {
	case "对":
		return "A"
	case "错":
		return "B"
	default:
		return input
	}
}

// 计算数量

func CalculateQuantity(m []types.Exam, oValue int) int {
	B := len(m)
	C := oValue
	if C == 0 || C == B {
		return B - 2
	} else if C-2 > B {
		return B
	} else {
		return C - 2
	}
}

// 获取题目

func GetSubject(items []types.Exam) []map[string]string {
	number := CalculateQuantity(items, len(items))
	var i int = 0
	var examList []map[string]string
	// 重新排序
	rand.Shuffle(len(items), func(i, j int) {
		items[i], items[j] = items[j], items[i]
	})
	for _, exam := range items {
		if i >= number {
			break
		}
		examList = append(examList, map[string]string{
			"id":    exam.Id,
			"value": ProcessString(exam.BzAnswer),
		})
		i++
	}
	return examList
}

// FlattenDict 将嵌套字典展平
func FlattenDict(d map[string]interface{}, parentKey string, sep string) map[string]interface{} {
	result := make(map[string]interface{})
	seenKeys := make(map[string]bool)
	var flatten func(map[string]interface{}, string)
	flatten = func(currentMap map[string]interface{}, parentKey string) {
		for k, v := range currentMap {
			newKey := k
			if parentKey != "" {
				newKey = parentKey + sep + k
			}
			originalKey := newKey
			i := 1
			for seenKeys[newKey] {
				newKey = fmt.Sprintf("%s%s%d", originalKey, sep, i)
				i++
			}
			seenKeys[newKey] = true

			if subMap, ok := v.(map[string]interface{}); ok {
				flatten(subMap, newKey)
			} else {
				result[newKey] = v
			}
		}
	}

	flatten(d, parentKey)
	return result
}

// GetSign 计算签名
func GetSign(postData map[string]interface{}) string {
	// 展平字典
	flatData := FlattenDict(postData, "", ".")
	// 去除键和值中的空格
	cleanData := make(map[string]string)
	for k, v := range flatData {
		key := strings.ReplaceAll(k, " ", "")
		value := strings.ReplaceAll(fmt.Sprintf("%v", v), " ", "")
		cleanData[key] = value
	}
	// 按键名排序（忽略大小写）
	keys := make([]string, 0, len(cleanData))
	for k := range cleanData {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return strings.ToLower(keys[i]) < strings.ToLower(keys[j])
	})
	// 按排序后的键连接对应的值
	var sortedValues strings.Builder
	for _, key := range keys {
		sortedValues.WriteString(cleanData[key])
	}
	// 拼接盐值
	tmpSign := sortedValues.String() + "8d387869d4c9eb0fe3b338a8c096324e"
	// 计算 MD5 哈希值
	hash := md5.Sum([]byte(tmpSign))
	return hex.EncodeToString(hash[:])
}

// 签名 old
//func GetSign(data map[string]interface{}) string {
//	// 提取键并按字母顺序排序（忽略大小写）
//	keys := make([]string, 0, len(data))
//	for key := range data {
//		keys = append(keys, key)
//	}
//	sort.Slice(keys, func(i, j int) bool {
//		return strings.ToLower(keys[i]) < strings.ToLower(keys[j])
//	})
//
//	// 拼接按排序后的值
//	var sortedValues strings.Builder
//	for _, key := range keys {
//		sortedValues.WriteString(strings.TrimSpace(ToString(data[key])))
//	}
//
//	// 拼接盐值并计算 MD5
//	tmpSign := sortedValues.String() + "8d387869d4c9eb0fe3b338a8c096324e"
//	hash := md5.Sum([]byte(tmpSign))
//	return hex.EncodeToString(hash[:])
//}

func GetReqTime() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}

// 获取当前时间的13位时间戳

func GetReqTime13() string {
	return strconv.FormatInt(time.Now().UnixMilli(), 10)
}

// 生成随机数 1-10，有小数点
func GetRandomFloat() float64 {
	return rand.Float64() * 10
}

// 格式化时间，将秒数转换为 HH:mm:ss 格式

func FormatTime(duration int64) string {
	hour := duration / 3600
	minute := (duration % 3600) / 60
	second := duration % 60
	// 格式化为两位数字，不足补零
	pad := func(num int64) string {
		return fmt.Sprintf("%02d", num)
	}
	return fmt.Sprintf("%s:%s:%s", pad(hour), pad(minute), pad(second))
}

func JsonUnmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func JsonMarshal(v interface{}) []byte {
	data, _ := json.Marshal(v)
	return data
}

// 将接口类型转换为字符串

//func ToString(value interface{}) string {
//	switch v := value.(type) {
//	case string:
//		return v
//	case int, int8, int16, int32, int64:
//		return strings.TrimSpace(fmt.Sprintf("%d", v))
//	case float32, float64:
//		return strings.TrimSpace(fmt.Sprintf("%f", v))
//	default:
//		return ""
//	}
//}

func ToString(value interface{}) string {
	return fmt.Sprintf("%v", value)
}

func ToInt(value interface{}) (int, error) {
	switch v := value.(type) {
	case string:
		if floatVal, err := strconv.ParseFloat(v, 64); err == nil {
			// 将字符串解析为 float64，再转为 int
			return int(floatVal), nil
		}
		return strconv.Atoi(v)
	case int, int8, int16, int32, int64:
		return v.(int), nil
	case float32, float64:
		return int(v.(float64)), nil
	default:
		return 0, fmt.Errorf("invalid type")
	}
}

func Sleep(seconds int) {
	time.Sleep(time.Duration(seconds) * time.Second)
}

func ParseAnswerString(data string, delimiter string) []string {
	// 按照分隔符分割字符串并去除每个项的空白字符
	var result []string
	items := strings.Split(data, delimiter)

	for _, item := range items {
		// 去除两端空白字符并检查是否为空
		trimmedItem := strings.TrimSpace(item)
		if trimmedItem != "" {
			result = append(result, trimmedItem)
		}
	}

	return result
}

func RemovePercentAndConvert(s string) (int, error) {
	s = strings.TrimSuffix(s, "%")
	// 将字符串转换为浮动数字
	floatVal, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, err
	}
	// 将浮动数字取整
	return int(floatVal), nil
}

func ConvertToInterfaceMap(input map[string]string) map[string]interface{} {
	output := make(map[string]interface{})
	for key, value := range input {
		output[key] = value
	}
	return output
}

// 计算数量

func CalculateNum(m []int, oValue int) int {
	B := len(m) // m 的长度
	C := oValue // o.value 的值
	if C == 0 || C == B {
		return B - 2
	}
	if C-2 > B {
		return B
	}
	return C - 2
}

// 本地图片转base64

func ImageToBase64(imagePath string) (string, error) {
	image, err := os.Open(imagePath)
	if err != nil {
		return "", err
	}
	defer func(image *os.File) {
		_ = image.Close()
	}(image)
	imageData, err := io.ReadAll(image)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(imageData), nil
}

func GenerateRules(duration int, spacing int, srDuration int, validPhotoTime bool, ruleType, action string) types.RulesData {
	if spacing == 0 {
		spacing = 900
	}

	rulesData := types.RulesData{
		Rules:     []types.Rule{},
		NextRule:  nil,
		BeginTime: 0,
		Index:     0,
		Spacing:   spacing,
	}

	Ue := spacing
	ot := duration
	at := Ue
	if ot < Ue {
		at = ot
	}

	xt := srDuration
	if xt == 0 {
		xt = rand.Intn(at + 1)
	}

	Jt := int(math.Ceil(float64(ot) / float64(at)))
	for Vt := 0; Vt <= Jt; Vt++ {
		qn := xt + at*Vt
		if qn > ot {
			continue
		}

		rule := types.Rule{
			Time:   qn,
			Type:   ruleType,
			Action: action,
		}
		rulesData.Rules = append(rulesData.Rules, rule)

		if rulesData.NextRule == nil && qn > srDuration {
			rulesData.Index = Vt
			rulesData.NextRule = &rule
		}
	}

	if xt == 0 {
		rulesData.Index = 0
		if len(rulesData.Rules) > 0 {
			rulesData.NextRule = &rulesData.Rules[0]
		}
	}

	if !validPhotoTime && rulesData.NextRule == nil && len(rulesData.Rules) > 0 {
		rulesData.NextRule = &rulesData.Rules[0]
	}

	return rulesData
}
