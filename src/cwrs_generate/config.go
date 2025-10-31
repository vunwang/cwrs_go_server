package main

import (
	"bufio"
	"bytes"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go/format"
	"os"
	"regexp"
	"strings"
	"text/template"
)

type ColumnInfo struct {
	ColumnName    string
	DataType      string
	ColumnComment string
	IsNullable    string
	ColumnKey     string
}

type TemplateData struct {
	TableName            string
	TableSelName         string
	Pkg                  string
	StructName           string
	StructNameLower      string
	Comment              string
	Columns              []FieldInfo
	PrimaryKey           FieldInfo
	ServiceVarName       string
	ServiceStructName    string
	ControllerStructName string
	PackagePath          string
	DaoVarName           string
	DaoStructName        string
}

type FieldInfo struct {
	ColumnName   string
	GoName       string
	GoNameLower  string
	GoType       string
	JsonTag      string
	Comment      string
	IsPrimaryKey bool
	IsRequired   bool
	DefaultValue interface{}
	TableSelName string
}

func ToGoType(mysqlType string) string {
	switch {
	case strings.Contains(mysqlType, "int"):
		return "int"
	case strings.Contains(mysqlType, "varchar"), strings.Contains(mysqlType, "text"), strings.Contains(mysqlType, "char"):
		return "string"
	case strings.Contains(mysqlType, "datetime"), strings.Contains(mysqlType, "timestamp"):
		return "string"
	case strings.Contains(mysqlType, "decimal"), strings.Contains(mysqlType, "float"), strings.Contains(mysqlType, "double"):
		return "float64"
	default:
		return "string"
	}
}

func ToGoName(columnName string) string {
	parts := strings.Split(columnName, "_")
	for i := range parts {
		if len(parts[i]) > 0 {
			parts[i] = strings.ToUpper(parts[i][:1]) + parts[i][1:]
		}
	}
	return strings.Join(parts, "")
}

// 获取表字段
func GetTableColumns(db *sql.DB, tableName string) ([]ColumnInfo, error) {
	query := `
        SELECT COLUMN_NAME, DATA_TYPE, COLUMN_COMMENT, IS_NULLABLE, COLUMN_KEY
        FROM information_schema.COLUMNS
        WHERE TABLE_SCHEMA = DATABASE() AND TABLE_NAME = ?
        ORDER BY ORDINAL_POSITION
    `
	rows, err := db.Query(query, tableName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var columns []ColumnInfo
	for rows.Next() {
		var col ColumnInfo
		if err := rows.Scan(&col.ColumnName, &col.DataType, &col.ColumnComment, &col.IsNullable, &col.ColumnKey); err != nil {
			return nil, err
		}
		columns = append(columns, col)
	}
	return columns, nil
}

func GenerateFile(templatePath, outputPath string, data interface{}) {
	// 检查文件是否存在
	if _, err := os.Stat(outputPath); err == nil {
		fmt.Printf("⚠️ 文件已存在: %s\n", outputPath)
		fmt.Print("是否覆盖？(y/n): ")

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))

		if input != "y" && input != "yes" {
			fmt.Println("❌ 操作已取消")
			return
		}
		fmt.Println("✅ 继续生成...")
	} else if !os.IsNotExist(err) {
		// 其他错误（如权限问题）
		panic(fmt.Sprintf("检查文件状态失败: %v", err))
	}

	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		panic(err)
	}

	file, err := os.Create(outputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if err := tmpl.Execute(file, data); err != nil {
		panic(err)
	}

	// 格式化代码
	content, err := os.ReadFile(outputPath)
	if err != nil {
		panic(err)
	}
	formatted, err := format.Source(content)
	if err != nil {
		fmt.Printf("%s 格式化失败: %v\n", outputPath, err)
		return
	}
	os.WriteFile(outputPath, formatted, 0644)
}

// 追加路由数据到指定文件
func AppendRouter(filePath string, data TemplateData) {
	if err := injectModuleIntoFile(filePath, data); err != nil {
		panic(err)
	}
}

// 注入模块路由
func injectModuleIntoFile(filePath string, m TemplateData) error {
	// 1. 读取文件内容
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	lines := strings.Split(string(content), "\n")

	// 2. 找 import 块结束位置（第一个 ")" 后）
	importEnd := findImportEnd(lines)
	if importEnd == -1 {
		return fmt.Errorf("未找到 import 块")
	}

	ImportTmpl := `{{.ControllerStructName}} "cwrs_go_server/src/server/{{.Pkg}}/controller"`

	RouteBlockTmpl := `
	// {{.Comment}}相关接口
	router.POST("/{{.StructNameLower}}/add", {{.ControllerStructName}}.Add{{.StructName}})
	router.PUT("/{{.StructNameLower}}/edit", {{.ControllerStructName}}.Edit{{.StructName}})
	router.DELETE("/{{.StructNameLower}}/del", {{.ControllerStructName}}.Del{{.StructName}})
	router.GET("/{{.StructNameLower}}/detail", {{.ControllerStructName}}.Get{{.StructName}}Detail)
	router.GET("/{{.StructNameLower}}/list", {{.ControllerStructName}}.Get{{.StructName}}List)`

	importTmpl := template.Must(template.New("import").Parse(ImportTmpl))
	routeBlockTmpl := template.Must(template.New("route").Parse(RouteBlockTmpl))
	// 3. 生成 import 行
	var importBuf bytes.Buffer
	if err := importTmpl.Execute(&importBuf, m); err != nil {
		return err
	}
	importLine := importBuf.String()

	// 4. 插入 import 行（去重检查）
	if !containsImport(lines, m.ControllerStructName, m.Pkg) {
		lines = append(lines[:importEnd], append([]string{importLine}, lines[importEnd:]...)...)
	} else {
		fmt.Println("import 已存在，跳过：", importLine)
	}

	// 5. 找 AuthRoutes 函数最后一个 } 前的位置
	authRoutesEnd := findAuthRoutesEnd(lines)
	if authRoutesEnd == -1 {
		return fmt.Errorf("未找到 AuthRoutes 函数结束位置")
	}

	// 6. 生成路由块
	var routeBuf bytes.Buffer
	if err := routeBlockTmpl.Execute(&routeBuf, m); err != nil {
		return err
	}
	routeBlock := routeBuf.String()
	routeLines := strings.Split(routeBlock, "\n")

	// 7. 插入路由块（在 AuthRoutes 结束大括号前）（去重检查）
	if !containsRouteBlock(lines, routeLines) {
		lines = append(lines[:authRoutesEnd], append(routeLines, lines[authRoutesEnd:]...)...)
	} else {
		fmt.Println("路由块已存在，跳过：", routeBlock)
	}

	// 8. 写回文件
	result := strings.Join(lines, "\n")
	return os.WriteFile(filePath, []byte(result), 0644)
}

// 查找 import 块结束位置
func findImportEnd(lines []string) int {
	inImport := false
	for i, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "import (" {
			inImport = true
		} else if inImport && trimmed == ")" {
			return i // 在这个位置之后插入
		}
	}
	return -1
}

// 检查是否已导入
func containsImport(lines []string, controller, packagePath string) bool {
	pattern := fmt.Sprintf(`%s "cwrs_go_server/src/server/%s/controller"`, controller, packagePath)
	for _, line := range lines {
		if matched, _ := regexp.MatchString(pattern, line); matched {
			return true
		}
	}
	return false
}

// 检查是否已包含路由块
func containsRouteBlock(lines []string, routeLines []string) bool {
	if len(routeLines) == 0 {
		return true
	}
	if len(lines) < len(routeLines) {
		return false
	}

	for i := 0; i <= len(lines)-len(routeLines); i++ {
		match := true
		for j := 0; j < len(routeLines); j++ {
			if strings.TrimSpace(lines[i+j]) != strings.TrimSpace(routeLines[j]) {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}
	return false
}

// 查找 AuthRoutes 函数结束位置（最后一个 }）
func findAuthRoutesEnd(lines []string) int {
	inAuthRoutes := false
	braceCount := 0

	for i, line := range lines {
		if strings.Contains(line, "func AuthRoutes(") {
			inAuthRoutes = true
		}

		if inAuthRoutes {
			for _, ch := range line {
				if ch == '{' {
					braceCount++
				} else if ch == '}' {
					braceCount--
					if braceCount == 0 {
						return i // 在这个位置之前插入
					}
				}
			}
		}
	}
	return -1
}
