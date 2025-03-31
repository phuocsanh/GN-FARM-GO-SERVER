package main

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// // 1
	// sugar:= zap.NewExample().Sugar()
	// sugar.Infof("Hello name:%s, age:%d", "Haha", 40)

	// // logger
	// logger := zap.NewExample()
	// logger.Info("Hello", zap.String("name", "Sanh"), zap.Int("age", 40))

	// 2
	// logger := zap.NewExample()
	// logger.Info("Hello NewExample")

	// loggerD,_ := zap.NewDevelopment()
	// loggerD.Info("Hello NewDevelopment")

	// loggerP,_ := zap.NewProduction()
	// loggerP.Info("Hello NewProduction")

    // 3
	endCoder := getEndCodeLog()
	sync := getWriteSync()
	core :=  zapcore.NewCore(endCoder,sync, zapcore.InfoLevel)
	logger := zap.New(core, zap.AddCaller())
	logger.Info("Info log", zap.Int("line ",1))
	logger.Error("Info log", zap.Int("line ",2))
}
func getEndCodeLog() zapcore.Encoder {

	endCodeConfig := zap.NewProductionEncoderConfig()
	endCodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder // timetamps 1721898691.973886 ->  2024-07-25T16:11:31.956+0700  
	endCodeConfig.TimeKey = "time" // ts -> time
	endCodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder // info -> INFO
	endCodeConfig.EncodeCaller = zapcore.ShortCallerEncoder // cli/main.log.go:19
	return zapcore.NewJSONEncoder(endCodeConfig)

}

func getWriteSync() zapcore.WriteSyncer {
	// Ensure the log directory exists
	if err := os.MkdirAll("./log", os.ModePerm); err != nil {
		fmt.Printf("Failed to create log directory: %v\n", err)
		return nil
	}

	file, err := os.OpenFile("./log/log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("Failed to open log file: %v\n", err)
		return nil
	}

	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)

	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}