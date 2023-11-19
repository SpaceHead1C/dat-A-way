package log

import "go.uber.org/zap"

type Logger struct {
	*zap.Logger
}

func New() (*Logger, error) {
	l, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	return &Logger{l}, nil
}

func GlobalLogger() *Logger {
	return &Logger{zap.L()}
}

func (l *Logger) Panicw(msg string, keysAndValues ...any) {
	l.Sugar().Panicw(msg, keysAndValues...)
}

func (l *Logger) Panicf(template string, args ...any) {
	l.Sugar().Panicf(template, args...)
}

func (l *Logger) Panicln(args ...any) {
	l.Sugar().Panicln(args...)
}

func (l *Logger) DPanicw(msg string, keysAndValues ...any) {
	l.Sugar().DPanicw(msg, keysAndValues...)
}

func (l *Logger) DPanicf(template string, args ...any) {
	l.Sugar().DPanicf(template, args...)
}

func (l *Logger) DPanicln(args ...any) {
	l.Sugar().DPanicln(args...)
}

func (l *Logger) Fatalw(msg string, keysAndValues ...any) {
	l.Sugar().Fatalw(msg, keysAndValues...)
}

func (l *Logger) Fatalf(template string, args ...any) {
	l.Sugar().Fatalf(template, args...)
}

func (l *Logger) Fatalln(args ...any) {
	l.Sugar().Fatalln(args...)
}

func (l *Logger) Errorw(msg string, keysAndValues ...any) {
	l.Sugar().Errorw(msg, keysAndValues...)
}

func (l *Logger) Errorf(template string, args ...any) {
	l.Sugar().Errorf(template, args...)
}

func (l *Logger) Errorln(args ...any) {
	l.Sugar().Errorln(args)
}

func (l *Logger) Warnw(msg string, keysAndValues ...any) {
	l.Sugar().Warnw(msg, keysAndValues...)
}

func (l *Logger) Warnf(template string, args ...any) {
	l.Sugar().Warnf(template, args...)
}

func (l *Logger) Warnln(args ...any) {
	l.Sugar().Warnln(args...)
}

func (l *Logger) Infow(msg string, keysAndValues ...any) {
	l.Sugar().Infow(msg, keysAndValues...)
}

func (l *Logger) Infof(template string, args ...any) {
	l.Sugar().Infof(template, args...)
}

func (l *Logger) Infoln(args ...any) {
	l.Sugar().Infoln(args...)
}

func (l *Logger) Debugw(msg string, keysAndValues ...any) {
	l.Sugar().Debugw(msg, keysAndValues...)
}

func (l *Logger) Debugf(template string, args ...any) {
	l.Sugar().Debugf(template, args...)
}

func (l *Logger) Debugln(args ...any) {
	l.Sugar().Debugln(args...)
}
