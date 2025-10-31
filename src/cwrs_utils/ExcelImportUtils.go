package cwrs_utils

import (
	"fmt"
	"mime/multipart"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/xuri/excelize/v2"
)

// ExcelImportOption 导入选项
type ExcelImportOption struct {
	SheetName   string   // 默认 "Sheet1"
	HeaderRow   int      // 表头所在行（从1开始），默认 1
	StartRow    int      // 数据起始行（从1开始），默认 2
	MaxRows     int      // 最大导入行数，0 表示不限制
	TimeLayouts []string // 时间解析格式，如 []string{"2006-01-02", "2006/01/02"}
}

// 默认选项
var defaultImportOption = ExcelImportOption{
	SheetName:   "Sheet1",
	HeaderRow:   1,
	StartRow:    2,
	MaxRows:     0,
	TimeLayouts: []string{"2006-01-02", "2006-01-02 15:04:05", "2006/01/02", "2006/01/02 15:04:05"},
}

// ImportExcelToStructSlice 从 multipart.File 导入 Excel 到结构体 slice
// data 必须是指向 struct slice 的指针，如 &[]User{}
func ImportExcelToStructSlice(file multipart.File, data interface{}, opts ...ExcelImportOption) error {
	option := defaultImportOption
	if len(opts) > 0 {
		option = opts[0]
	}

	// 检查 v 是否为 slice 指针
	rv := reflect.ValueOf(data)
	if rv.Kind() != reflect.Ptr || rv.Elem().Kind() != reflect.Slice {
		return fmt.Errorf("v 必须是指向 slice 的指针")
	}

	sliceValue := rv.Elem()
	elemType := sliceValue.Type().Elem()

	// 支持 *struct
	for elemType.Kind() == reflect.Ptr {
		elemType = elemType.Elem()
	}
	if elemType.Kind() != reflect.Struct {
		return fmt.Errorf("slice 元素必须是 struct 或 *struct")
	}

	// 解析结构体字段的 excel tag
	fieldMap := make(map[string]struct {
		Index  int
		Type   reflect.Type
		Format string
	})
	var fieldOrder []string

	for i := 0; i < elemType.NumField(); i++ {
		field := elemType.Field(i)
		tag := field.Tag.Get("excel")
		if tag == "" || tag == "-" {
			continue
		}

		parts := strings.Split(tag, ",")
		title := strings.TrimSpace(parts[0])
		if title == "" {
			continue
		}

		format := ""
		for _, opt := range parts[1:] {
			opt = strings.TrimSpace(opt)
			if strings.HasPrefix(opt, "format=") {
				format = strings.TrimPrefix(opt, "format=")
			}
		}

		fieldMap[title] = struct {
			Index  int
			Type   reflect.Type
			Format string
		}{Index: i, Type: field.Type, Format: format}
		fieldOrder = append(fieldOrder, title)
	}

	if len(fieldMap) == 0 {
		return fmt.Errorf("未找到任何带 excel tag 的字段")
	}

	// 读取 Excel
	f, err := excelize.OpenReader(file)
	if err != nil {
		return fmt.Errorf("打开 Excel 失败: %w", err)
	}
	defer f.Close()

	rows, err := f.GetRows(option.SheetName)
	if err != nil {
		return fmt.Errorf("读取 Sheet 失败: %w", err)
	}

	if len(rows) < option.HeaderRow {
		return fmt.Errorf("Excel 行数不足，无法读取表头")
	}

	// 获取表头（列标题）
	headerRow := rows[option.HeaderRow-1] // 转为 0-based

	// 建立列标题 → 列索引 映射
	colIndexMap := make(map[string]int)
	for j, title := range headerRow {
		colIndexMap[strings.TrimSpace(title)] = j
	}

	// 验证必要列是否存在
	for _, title := range fieldOrder {
		if _, exists := colIndexMap[title]; !exists {
			return fmt.Errorf("缺少必要列: %s", title)
		}
	}

	// 开始解析数据行
	dataStart := option.StartRow - 1 // 转为 0-based
	dataEnd := len(rows)
	if option.MaxRows > 0 && dataStart+option.MaxRows < dataEnd {
		dataEnd = dataStart + option.MaxRows
	}

	for i := dataStart; i < dataEnd; i++ {
		if i >= len(rows) {
			break
		}
		row := rows[i]
		if len(row) == 0 {
			continue // 跳过空行
		}

		// 检查是否整行为空
		allEmpty := true
		for _, cell := range row {
			if strings.TrimSpace(cell) != "" {
				allEmpty = false
				break
			}
		}
		if allEmpty {
			continue
		}

		// 创建新结构体实例
		var elem reflect.Value
		if sliceValue.Type().Elem().Kind() == reflect.Ptr {
			elem = reflect.New(elemType)
		} else {
			elem = reflect.New(elemType).Elem()
		}

		// 填充字段
		for _, title := range fieldOrder {
			colIdx, exists := colIndexMap[title]
			if !exists {
				continue
			}

			var cellValue string
			if colIdx < len(row) {
				cellValue = strings.TrimSpace(row[colIdx])
			}

			fieldInfo := fieldMap[title]
			field := elem
			if sliceValue.Type().Elem().Kind() == reflect.Ptr {
				field = field.Elem()
			}
			targetField := field.Field(fieldInfo.Index)

			if err := setFieldValue(targetField, cellValue, fieldInfo.Type, fieldInfo.Format, option.TimeLayouts); err != nil {
				rowNum := i + 1 // Excel 行号从 1 开始
				return fmt.Errorf("第 %d 行，字段 %s 赋值失败: %w", rowNum, title, err)
			}
		}

		// 添加到 slice
		if sliceValue.Type().Elem().Kind() == reflect.Ptr {
			sliceValue.Set(reflect.Append(sliceValue, elem))
		} else {
			sliceValue.Set(reflect.Append(sliceValue, elem))
		}
	}

	return nil
}

// setFieldValue 将字符串值转换为目标类型并设置到字段
func setFieldValue(field reflect.Value, value string, fieldType reflect.Type, format string, timeLayouts []string) error {
	if value == "" {
		// 空值：设置零值（或保持原样）
		return nil
	}

	// 解引用指针类型
	targetType := fieldType
	if targetType.Kind() == reflect.Ptr {
		targetType = targetType.Elem()
	}

	var finalValue reflect.Value

	switch targetType.Kind() {
	case reflect.String:
		finalValue = reflect.ValueOf(value)

	case reflect.Bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return fmt.Errorf("无法解析为布尔值: %v", err)
		}
		finalValue = reflect.ValueOf(b)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fmt.Errorf("无法解析为整数: %v", err)
		}
		finalValue = reflect.ValueOf(i).Convert(targetType)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		u, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return fmt.Errorf("无法解析为无符号整数: %v", err)
		}
		finalValue = reflect.ValueOf(u).Convert(targetType)

	case reflect.Float32, reflect.Float64:
		f, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return fmt.Errorf("无法解析为浮点数: %v", err)
		}
		finalValue = reflect.ValueOf(f).Convert(targetType)

	case reflect.Struct:
		if targetType == reflect.TypeOf(time.Time{}) {
			var t time.Time
			var err error
			layouts := timeLayouts
			if format != "" {
				layouts = []string{format}
			}
			for _, layout := range layouts {
				t, err = time.Parse(layout, value)
				if err == nil {
					break
				}
			}
			if err != nil {
				return fmt.Errorf("无法解析时间为 %v: %w", layouts, err)
			}
			finalValue = reflect.ValueOf(t)
		} else {
			return fmt.Errorf("不支持的 struct 类型: %s", targetType)
		}

	default:
		return fmt.Errorf("不支持的字段类型: %s", targetType.Kind())
	}

	// 处理指针类型
	if fieldType.Kind() == reflect.Ptr {
		ptr := reflect.New(targetType)
		ptr.Elem().Set(finalValue)
		finalValue = ptr
	}

	field.Set(finalValue)
	return nil
}
