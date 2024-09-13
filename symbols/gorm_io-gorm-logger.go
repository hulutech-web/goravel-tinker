// Code generated by 'yaegi extract gorm.io/gorm/logger'. DO NOT EDIT.

package symbols

import (
	"context"
	"go/constant"
	"go/token"
	"gorm.io/gorm/logger"
	"reflect"
	"time"
)

func init() {
	Symbols["gorm.io/gorm/logger/logger"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Blue":              reflect.ValueOf(constant.MakeFromLiteral("\"\\x1b[34m\"", token.STRING, 0)),
		"BlueBold":          reflect.ValueOf(constant.MakeFromLiteral("\"\\x1b[34;1m\"", token.STRING, 0)),
		"Cyan":              reflect.ValueOf(constant.MakeFromLiteral("\"\\x1b[36m\"", token.STRING, 0)),
		"Default":           reflect.ValueOf(&logger.Default).Elem(),
		"Discard":           reflect.ValueOf(&logger.Discard).Elem(),
		"ErrRecordNotFound": reflect.ValueOf(&logger.ErrRecordNotFound).Elem(),
		"Error":             reflect.ValueOf(logger.Error),
		"ExplainSQL":        reflect.ValueOf(logger.ExplainSQL),
		"Green":             reflect.ValueOf(constant.MakeFromLiteral("\"\\x1b[32m\"", token.STRING, 0)),
		"Info":              reflect.ValueOf(logger.Info),
		"Magenta":           reflect.ValueOf(constant.MakeFromLiteral("\"\\x1b[35m\"", token.STRING, 0)),
		"MagentaBold":       reflect.ValueOf(constant.MakeFromLiteral("\"\\x1b[35;1m\"", token.STRING, 0)),
		"New":               reflect.ValueOf(logger.New),
		"Recorder":          reflect.ValueOf(&logger.Recorder).Elem(),
		"Red":               reflect.ValueOf(constant.MakeFromLiteral("\"\\x1b[31m\"", token.STRING, 0)),
		"RedBold":           reflect.ValueOf(constant.MakeFromLiteral("\"\\x1b[31;1m\"", token.STRING, 0)),
		"Reset":             reflect.ValueOf(constant.MakeFromLiteral("\"\\x1b[0m\"", token.STRING, 0)),
		"Silent":            reflect.ValueOf(logger.Silent),
		"Warn":              reflect.ValueOf(logger.Warn),
		"White":             reflect.ValueOf(constant.MakeFromLiteral("\"\\x1b[37m\"", token.STRING, 0)),
		"Yellow":            reflect.ValueOf(constant.MakeFromLiteral("\"\\x1b[33m\"", token.STRING, 0)),
		"YellowBold":        reflect.ValueOf(constant.MakeFromLiteral("\"\\x1b[33;1m\"", token.STRING, 0)),

		// type definitions
		"Config":    reflect.ValueOf((*logger.Config)(nil)),
		"Interface": reflect.ValueOf((*logger.Interface)(nil)),
		"LogLevel":  reflect.ValueOf((*logger.LogLevel)(nil)),
		"Writer":    reflect.ValueOf((*logger.Writer)(nil)),

		// interface wrapper definitions
		"_Interface": reflect.ValueOf((*_gorm_io_gorm_logger_Interface)(nil)),
		"_Writer":    reflect.ValueOf((*_gorm_io_gorm_logger_Writer)(nil)),
	}
}

// _gorm_io_gorm_logger_Interface is an interface wrapper for Interface type
type _gorm_io_gorm_logger_Interface struct {
	IValue   interface{}
	WError   func(a0 context.Context, a1 string, a2 ...interface{})
	WInfo    func(a0 context.Context, a1 string, a2 ...interface{})
	WLogMode func(a0 logger.LogLevel) logger.Interface
	WTrace   func(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error)
	WWarn    func(a0 context.Context, a1 string, a2 ...interface{})
}

func (W _gorm_io_gorm_logger_Interface) Error(a0 context.Context, a1 string, a2 ...interface{}) {
	W.WError(a0, a1, a2...)
}
func (W _gorm_io_gorm_logger_Interface) Info(a0 context.Context, a1 string, a2 ...interface{}) {
	W.WInfo(a0, a1, a2...)
}
func (W _gorm_io_gorm_logger_Interface) LogMode(a0 logger.LogLevel) logger.Interface {
	return W.WLogMode(a0)
}
func (W _gorm_io_gorm_logger_Interface) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	W.WTrace(ctx, begin, fc, err)
}
func (W _gorm_io_gorm_logger_Interface) Warn(a0 context.Context, a1 string, a2 ...interface{}) {
	W.WWarn(a0, a1, a2...)
}

// _gorm_io_gorm_logger_Writer is an interface wrapper for Writer type
type _gorm_io_gorm_logger_Writer struct {
	IValue  interface{}
	WPrintf func(a0 string, a1 ...interface{})
}

func (W _gorm_io_gorm_logger_Writer) Printf(a0 string, a1 ...interface{}) {
	W.WPrintf(a0, a1...)
}
