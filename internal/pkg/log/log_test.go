// Copyright 2024 孔令飞 <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/miniblog. The professional
// version of this repository is https://github.com/onexstack/onex.

package log

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"github.com/onexstack/miniblog/internal/pkg/contextx"
)

// MockLogger 用于测试的自定义 Logger
var mockLogger *zapLogger

// TestMain 初始化测试环境
func TestMain(m *testing.M) {
	opts := &Options{
		Level:             "debug",
		DisableCaller:     false,
		DisableStacktrace: false,
		Format:            "json",
		OutputPaths:       []string{"stdout"},
	}
	Init(opts)
	mockLogger = std
	m.Run()
}

// TestLoggerMethods 测试日志记录方法
func TestLoggerMethods(t *testing.T) {
	// 测试不同级别的日志记录
	assert.NotPanics(t, func() {
		Debugw("debug message", "key1", "value1")
		Infow("info message", "key2", "value2")
		Warnw("warn message", "key3", "value3")
		Errorw("error message", "key4", "value4")
	}, "Log methods should not cause a panic in this test")

	assert.Panics(t, func() {
		Panicw("fatal message", "key6", "value6") // 这会导致程序退出
	}, "Panicw should cause a panic and exit the program")
}

// TestLoggerInitialization 测试 Logger 初始化
func TestLoggerInitialization(t *testing.T) {
	opts := NewOptions()
	logger := New(opts)

	assert.NotNil(t, logger, "Logger should not be nil after initialization")
	assert.IsType(t, &zapLogger{}, logger, "Logger should be of type *zapLogger")
}

// TestSync 测试日志同步
func TestSync(t *testing.T) {
	assert.NotPanics(t, func() {
		Sync() // 确保 Sync 不会引发恐慌
	}, "Sync should not panic")
}

// 性能测试用例
func BenchmarkZapLoggerW(b *testing.B) {
	// 创建一个 zapLogger 实例（使用 zap.NewNop() 模拟 logger）
	logger := &zapLogger{z: zap.NewNop()}

	// 创建一个包含上下文值的 context
	ctx := contextx.WithRequestID(context.Background(), "request-id-12345")
	ctx = contextx.WithUserID(ctx, "user-id-67890")

	// 重复调用 W 函数，测量性能
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = logger.W(ctx)
	}
}
