package cwrs_utils

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

// ExportStructSliceToExcel 从结构体 slice 自动生成 Excel 并通过 Gin 返回下载
// 要求结构体字段使用 `excel:"标题[,width=xx][,format=yy]"` tag
// 顺序按照元素声明顺序写入 Excel
// typeName 导出类型，tmpl 为模板导出，其他为普通导出
func ExportStructSliceToExcel(c *gin.Context, data interface{}, fileName, typeName string) error {
	v := reflect.ValueOf(data)
	if v.Kind() != reflect.Slice {
		return fmt.Errorf("data 必须是 slice 类型")
	}

	if typeName != "tmpl" && v.Len() == 0 {
		return fmt.Errorf("数据为空，无法导出")
	}

	// 获取元素类型（支持 *struct）
	elemType := reflect.TypeOf(data).Elem()
	for elemType.Kind() == reflect.Ptr {
		elemType = elemType.Elem()
	}
	if elemType.Kind() != reflect.Struct {
		return fmt.Errorf("slice 元素必须是 struct 或 *struct")
	}

	// 解析带 excel tag 的字段（保持声明顺序）
	var fields []struct {
		Index  int
		Title  string
		Width  float64
		Format string
	}

	for i := 0; i < elemType.NumField(); i++ {
		field := elemType.Field(i)
		tag := field.Tag.Get("excel")
		if tag == "" || tag == "-" {
			continue // 跳过未标记或显式忽略的字段
		}

		parts := strings.Split(tag, ",")
		title := strings.TrimSpace(parts[0])
		if title == "" {
			continue
		}

		width := 12.0
		format := ""

		// 解析可选参数
		for _, part := range parts[1:] {
			part = strings.TrimSpace(part)
			if strings.HasPrefix(part, "width=") {
				if w, err := strconv.ParseFloat(strings.TrimPrefix(part, "width="), 64); err == nil {
					width = w
				}
			} else if strings.HasPrefix(part, "format=") {
				format = strings.TrimPrefix(part, "format=")
			}
		}

		fields = append(fields, struct {
			Index  int
			Title  string
			Width  float64
			Format string
		}{Index: i, Title: title, Width: width, Format: format})
	}

	if len(fields) == 0 {
		return fmt.Errorf("未找到任何带 excel tag 的字段")
	}

	// 创建 Excel 文件
	f := excelize.NewFile()
	sheetName := "Sheet1"

	// 写入表头并设置列宽
	for colIdx, field := range fields {
		colName, _ := excelize.ColumnNumberToName(colIdx + 1)
		f.SetCellValue(sheetName, colName+"1", field.Title)
		f.SetColWidth(sheetName, colName, colName, field.Width)
	}

	// 写入数据行
	if typeName != "tmpl" {

		for rowIdx := 0; rowIdx < v.Len(); rowIdx++ {
			row := v.Index(rowIdx)
			if row.Kind() == reflect.Ptr {
				if row.IsNil() {
					continue
				}
				row = row.Elem()
			}

			for colIdx, field := range fields {
				fieldValue := row.Field(field.Index)
				if !fieldValue.IsValid() {
					continue
				}

				// 处理 nil 指针字段
				if fieldValue.Kind() == reflect.Ptr && fieldValue.IsNil() {
					continue
				}

				var cellValue interface{}
				switch k := fieldValue.Kind(); k {
				case reflect.String:
					cellValue = fieldValue.String()
				case reflect.Bool:
					cellValue = fieldValue.Bool()
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					cellValue = fieldValue.Int()
				case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
					cellValue = fieldValue.Uint()
				case reflect.Float32, reflect.Float64:
					cellValue = fieldValue.Float()
				default:
					cellValue = fmt.Sprintf("%v", fieldValue.Interface())
				}

				colName, _ := excelize.ColumnNumberToName(colIdx + 1)
				cell := fmt.Sprintf("%s%d", colName, rowIdx+2)
				f.SetCellValue(sheetName, cell, cellValue)

				// 应用单元格格式
				if field.Format != "" {
					styleID, err := f.NewStyle(&excelize.Style{
						CustomNumFmt: &field.Format, // 传 *string
					})
					if err == nil {
						_ = f.SetCellStyle(sheetName, cell, cell, styleID)
					}
				}
			}
		}
	}
	// 写入响应
	buf, err := f.WriteToBuffer()
	if err != nil {
		return err
	}

	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s.xlsx", fileName))
	c.Header("Content-Length", strconv.Itoa(buf.Len()))
	c.Data(http.StatusOK, "application/octet-stream", buf.Bytes())
	return nil
}
