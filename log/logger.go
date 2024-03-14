package log

func Info(data ...any) {
	sugar.Info(data...)
}

func Debug(data ...any) {
	sugar.Debug(data...)
}

func Error(data ...any) {
	sugar.Error(data...)
}

func Warn(data ...any) {
	sugar.Warn(data...)
}

func Fatal(data ...any) {
	sugar.Fatal(data...)
}

func Infof(template string, data ...any) {
	sugar.Infof(template, data...)
}

func Debugf(template string, data ...any) {
	sugar.Debugf(template, data...)
}

func Errorf(template string, data ...any) {
	sugar.Errorf(template, data...)
}

func Warnf(template string, data ...any) {
	sugar.Warnf(template, data...)
}

func Fatalf(template string, data ...any) {
	sugar.Fatalf(template, data...)
}

func Infow(msg string, keysAndValues ...any) {
	sugar.Infow(msg, keysAndValues...)
}

func Debugw(msg string, keysAndValues ...any) {
	sugar.Debugw(msg, keysAndValues...)
}

func Errorw(msg string, keysAndValues ...any) {
	sugar.Errorw(msg, keysAndValues...)
}

func Warnw(msg string, keysAndValues ...any) {
	sugar.Warnw(msg, keysAndValues...)
}

func Fatalw(msg string, keysAndValues ...any) {
	sugar.Fatalw(msg, keysAndValues...)
}
