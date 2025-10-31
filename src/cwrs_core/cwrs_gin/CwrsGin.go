package cwrs_gin

import (
	_ "cwrs_go_server/docs" // 导入生成的 docs 包
	"cwrs_go_server/src/cwrs_core/cwrs_middleware"
	"cwrs_go_server/src/cwrs_core/cwrs_viper"
	"cwrs_go_server/src/cwrs_core/cwrs_zap_logger"
	"cwrs_go_server/src/cwrs_routes"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"strings"
)

func InitGin() *gin.Engine {
	gin.SetMode(cwrs_viper.GlobalViper.GetString("gin.mode"))

	r := gin.New()
	r.Use(Cors())
	r.Use(gin.Logger(), gin.Recovery())
	r.Use(cwrs_zap_logger.OperationLog()) // 关键：OperationLog 必须在 ResponseInterceptor 之前注册
	r.Use(cwrs_zap_logger.BodyReader())
	r.Use(cwrs_zap_logger.ResponseInterceptor())
	//引入swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	group := r.Group("/api")
	group.Static("/images", "./static/images") //将该分组下的/images路径映射到本地文件系统的/static/images目录
	group.Use(cwrs_middleware.SysNotAuthLogger())
	//不做鉴权的
	{
		//内部访问接口
		cwrs_routes.NotAuthRoutes(group)
		//外部访问接口
		cwrs_routes.ExternalNotAuthApi(group)
	}
	//做鉴权的
	group.Use(cwrs_middleware.JWTAuthMiddleware(), cwrs_middleware.SysAuthLogger())
	{
		//内部访问接口
		cwrs_routes.AuthRoutes(group)
		//外部访问接口
		cwrs_routes.ExternalAuthApi(group)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})

	fmt.Println("Gin Initialize OK !")
	return r
}

// Cors 处理跨域请求,支持options访问
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method               //请求方法
		origin := c.Request.Header.Get("Origin") //请求头部
		var headerKeys []string                  // 声明请求头keys
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			//c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", origin)                             // 这是允许的跨域请求的源
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE") //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			//header的类型
			c.Header("Access-Control-Allow-Headers", "Content-Length, Token, Menu-Id, Origin, Host, Connection, Content-Type, Referer")
			//允许跨域设置  可以返回其他子段
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Referer,Content-Type") // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")                                                                                                // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false")                                                                                       //  跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json")                                                                                                   // 设置返回格式是json
		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		// 处理请求
		c.Next() //  处理请求
	}
}
