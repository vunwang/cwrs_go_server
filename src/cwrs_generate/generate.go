package main

import (
	"cwrs_go_server/src/cwrs_utils"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// =========必填参数 只修改这几个参数其他不要改===========
	pkg := "sys_user_ceshi"       //包名
	comment := "用户身份"             //模块
	table := "sys_user_dept_role" //表名
	tableSelName := "sudr"        //表别名
	// =========必填参数 只修改这几个参数其他不要改===========

	//mysql数据库配置
	//用户名
	user := "api"
	//密码
	password := "Djr9WpVtPJBzg"
	//域名/ip
	host := "dc.jstydzkj.com"
	//端口
	port := 3306
	//数据库名
	dbName := "plcsh"

	// 连接数据库
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, dbName))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 获取表字段
	columns, err := GetTableColumns(db, table)
	if err != nil {
		panic(err)
	}

	// 获取当前目录
	proDir, _ := os.Getwd()
	baseDir := proDir + "/src/server/" + pkg

	// 构建模板数据
	var fields []FieldInfo
	var primaryKey FieldInfo
	for _, col := range columns {
		field := FieldInfo{
			ColumnName:   col.ColumnName,
			GoName:       ToGoName(col.ColumnName),
			GoNameLower:  cwrs_utils.LowerFirst(ToGoName(col.ColumnName)),
			GoType:       ToGoType(col.DataType),
			JsonTag:      cwrs_utils.LowerFirst(ToGoName(col.ColumnName)),
			Comment:      col.ColumnComment,
			IsPrimaryKey: col.ColumnKey == "PRI",
			IsRequired:   col.IsNullable == "NO",
			TableSelName: tableSelName,
		}
		fields = append(fields, field)
		if field.IsPrimaryKey {
			primaryKey = field
		}
	}

	data := TemplateData{
		TableName:            table,
		TableSelName:         tableSelName,
		Pkg:                  pkg,
		StructName:           ToGoName(table),
		StructNameLower:      cwrs_utils.LowerFirst(ToGoName(table)),
		Comment:              comment,
		Columns:              fields,
		PrimaryKey:           primaryKey,
		ServiceVarName:       cwrs_utils.LowerFirst(ToGoName(table)) + "ServiceImpl",
		ServiceStructName:    ToGoName(table) + "Service",
		ControllerStructName: cwrs_utils.LowerFirst(ToGoName(table)) + "Controller",
		DaoVarName:           cwrs_utils.LowerFirst(ToGoName(table)) + "DaoImpl",
		DaoStructName:        ToGoName(table) + "Dao",
	}

	// 生成目录
	os.MkdirAll(filepath.Join(baseDir, "pojo"), 0755)
	os.MkdirAll(filepath.Join(baseDir, "service"), 0755)
	os.MkdirAll(filepath.Join(baseDir, "dao"), 0755)
	os.MkdirAll(filepath.Join(baseDir, "controller"), 0755)

	// 生成文件
	tempDir := proDir + "/src/cwrs_generate/templates/"
	GenerateFile(tempDir+"entity.tmpl", filepath.Join(baseDir, "pojo", data.StructName+"Entity.go"), data)
	GenerateFile(tempDir+"dao.tmpl", filepath.Join(baseDir, "dao", data.StructName+"Dao.go"), data)
	GenerateFile(tempDir+"service.tmpl", filepath.Join(baseDir, "service", data.StructName+"Service.go"), data)
	GenerateFile(tempDir+"controller.tmpl", filepath.Join(baseDir, "controller", data.StructName+"Controller.go"), data)

	// 追加路由
	AppendRouter(proDir+"/src/cwrs_routes/CwrsRouter.go", data)

	fmt.Println("✅ 代码生成完成！")
}
