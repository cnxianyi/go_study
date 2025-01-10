package zap

import (
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

// go get -u go.uber.org/zap
// go get -u github.com/natefinch/lumberjack

// 默认配置
func InitDefault() {
	// NewProduction预设配置 初始化 logger
	// 默认配置只会输出到控制台
	logger, _ := zap.NewProduction()
	defer logger.Sync() // 安全关闭

	// 手动使用 SugaredLogger模式. 默认Logger模式
	sugar := logger.Sugar()
	// Info 不支持结构化. 无法识别  %s
	sugar.Info("user %s logged in", "18888888888") // {"level":"info","ts":1736478442.6244428,"caller":"pzap/pzap.go:21","msg":"user %s
	// Infof 支持 %s
	sugar.Infof("user %s logged in", "18888888888") // {"level":"info","ts":1736478442.624498,"caller":"pzap/pzap.go:23","msg":"user 18888888888 logged in"}
	// Infow 支持 键值对的方式.
	sugar.Infow("failed to fetch URL",
		"usr", "http",
		"attempt", 3,
		"backoff", time.Second,
	) // {"level":"info","ts":1736478442.624502,"caller":"pzap/pzap.go:25","msg":"failed to fetch URL","usr":"http","attempt":3,"backoff":1}
	// Infoln 相较于 Info 多了换行. 类似 Println
	sugar.Infoln("Failed to fetch URL: %s", "http") // {"level":"info","ts":1736478442.6245189,"caller":"pzap/pzap.go:31","msg":"Failed to fetch URL: %s http"}
}

// 将日志保存到文件
func Init() {
	// 1. 创建目录
	logDir := "./log"
	os.Mkdir(logDir, 0755)

	// 2. 创建文件
	logFile := filepath.Join(logDir, "app.log")
	file, _ := os.OpenFile(logFile,
		os.O_CREATE| // 不存在则创建
			os.O_APPEND| // 追加模式写入文件
			os.O_RDWR, // 读写文件
		0644, // 权限
	)

	// 3. 配置编码器
	encodingConfig := zap.NewProductionEncoderConfig()
	encodingConfig.EncodeTime = zapcore.ISO8601TimeEncoder // 修改时间编码

	// 4. 创建 zapcore 配置
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encodingConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(file)),
		zap.InfoLevel,
	)

	Logger = zap.New(core, zap.AddCaller())
}
