package xapp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	mrand "math/rand"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"sync"
	"syscall"
	"time"
	"xapp/xdb"
	"xapp/xdb/game/query"
	"xapp/xredis"

	"github.com/beego/beego/logs"
	"github.com/gin-gonic/gin"
	consul "github.com/hashicorp/consul/api"
	"github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

var Id string
var Project string
var Running bool = true
var WaitingGroup = new(sync.WaitGroup)

var ApiV1 *gin.RouterGroup
var ApiV2 *gin.RouterGroup
var ApiV3 *gin.RouterGroup
var ApiV4 *gin.RouterGroup
var ApiV5 *gin.RouterGroup
var ApiV6 *gin.RouterGroup
var ApiV7 *gin.RouterGroup
var ApiV8 *gin.RouterGroup
var ApiV9 *gin.RouterGroup

var db *xdb.XDb = new(xdb.XDb)
var redis *xredis.XRedis = new(xredis.XRedis)
var db_query *query.Query

func IsEnvPrd() bool {

	return strings.Contains(env, "prd")
}

func IsEnvDev() bool {
	return strings.Contains(env, "dev")
}

func IsEnvTest() bool {
	return strings.Contains(env, "test")
}

func IsEnvPrePrd() bool {
	return strings.Contains(env, "preprd")
}

var env string
var router *gin.Engine

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func access_log() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyLogWriter := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyLogWriter
		startTime := time.Now()
		c.Next()
		pattern := `(/consul|/upload|/download|/swagger)`
		re := regexp.MustCompile(pattern)
		match := re.MatchString(c.Request.URL.Path)
		if match {
			return
		}
		endTime := time.Now()
		spend := endTime.Sub(startTime)
		reqbytes, _ := io.ReadAll(c.Request.Body)
		req := string(reqbytes)
		logdata := make(map[string]interface{})
		logdata["path"] = c.Request.URL.Path
		logdata["method"] = c.Request.Method
		logdata["ip"] = c.ClientIP()
		logdata["time"] = startTime.Format("2006-01-02 15:04:05")
		logdata["spend"] = spend
		logdata["status"] = c.Writer.Status()
		logdata["req"] = req
		bytes, _ := json.Marshal(logdata)
		bytes = append(bytes, []byte("\r\n")...)
		gin.DefaultWriter.Write(bytes)
	}
}

func cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "Content-Type, x-token, Content-Length, X-Requested-With")
		context.Header("Access-Control-Allow-Methods", "GET,POST")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
		context.Next()
	}
}

func Init() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath("./")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		logs.Error("读取配置文件失败", err)
		return
	}
	Id = viper.GetString("server.id")
	Project = viper.GetString("server.project")
	env = viper.GetString("server.env")
	mrand.NewSource(time.Now().UnixNano())
	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(3)
	logs.SetLogger(logs.AdapterFile, fmt.Sprintf(`{"filename":"_log/%s.log","maxsize":10485760}`, Project))
	logs.SetLogger(logs.AdapterConsole, `{"color":true}`)
	gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)
	f, _ := os.OpenFile("_log/http.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	gin.DefaultWriter = io.MultiWriter(f)
	router = gin.New()
	router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Use(access_log())
	router.Use(cors())
	ApiV1 = router.Group("/api/v1")
	ApiV2 = router.Group("/api/v2")
	ApiV3 = router.Group("/api/v3")
	ApiV4 = router.Group("/api/v4")
	ApiV5 = router.Group("/api/v5")
	ApiV6 = router.Group("/api/v6")
	ApiV7 = router.Group("/api/v7")
	ApiV8 = router.Group("/api/v8")
	ApiV9 = router.Group("/api/v9")
	if viper.GetString("consul.host") != "" {
		selfhost := viper.GetString("server.host")
		config := consul.DefaultConfig()
		config.Address = viper.GetString("consul.host") + ":" + viper.GetString("consul.port")
		consul_client, err := consul.NewClient(config)
		if err != nil {
			logs.Error("create consul client error : ", err)
			return
		}
		registration := &consul.AgentServiceRegistration{
			ID:      fmt.Sprintf("%s_%s", Project, viper.GetString("server.id")),
			Name:    Project,
			Port:    viper.GetInt("http.port"),
			Tags:    []string{viper.GetString("rpc.port")},
			Address: selfhost,
		}
		check := new(consul.AgentServiceCheck)
		check.HTTP = fmt.Sprintf("http://%s:%s/consul", viper.GetString("server.host"), viper.GetString("http.port"))
		check.Timeout = "1s"
		check.Interval = "2s"
		check.DeregisterCriticalServiceAfter = "1s"
		registration.Check = check
		if err := consul_client.Agent().ServiceRegister(registration); err != nil {
			logs.Error("register to consul error: ", err.Error())
			return
		}
		router.GET("/consul", func(c *gin.Context) {
			c.String(200, "ok")
		})
	}
	if viper.GetString("db.user") != "" {
		db.Init("db")
		db_query = query.Use(db.Gorm())
	}
	if viper.GetString("redis.host") != "" {
		redis.Init("redis")
	}
}

func Run(callback func()) {
	{
		prcport := viper.GetInt("rpc.port")
		if prcport > 0 {
			go func() {
				listener, err := net.Listen("tcp", fmt.Sprintf(":%d", prcport))
				if err != nil {
					logs.Error("Error creating rpc listener:", err.Error())
					return
				}
				logs.Debug("start rpc server at port: ", prcport)
				for {
					conn, err := listener.Accept()
					if err != nil {
						logs.Error("Error accepting rpc connection:", err.Error())
						continue
					}
					go rpc.ServeConn(conn)
				}
			}()
		}
	}
	time.Sleep(time.Microsecond * 100)
	{
		go func() {
			port := viper.GetString("http.port")
			logs.Debug("start server at port: ", port)
			router.Run(":" + port)
		}()
	}
	time.Sleep(time.Microsecond * 100)
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT) //ctrl+c 信号
	signal.Notify(sig, syscall.SIGTERM)
	WaitingGroup.Add(1)
	go func() {
		for {
			select {
			case <-sig:
				fmt.Println("server exit")
				Running = false
				WaitingGroup.Done()
			}
		}
	}()
	callback()
	WaitingGroup.Wait()
	logs.Debug("****************server exit****************")
}

func Db() *gorm.DB {
	return db.Gorm()
}

func DbQuery() *query.Query {
	return db_query
}

func Redis() *xredis.XRedis {
	return redis
}
