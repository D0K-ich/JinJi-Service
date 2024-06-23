package logs

import (
	"errors"
	"io"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	Level 		string 		`yaml:"level"`
	//Formatter 	*logrus.Formatter 	`yaml:"formatter"`
	Output 		io.Writer			`yaml:"output"`
}

func(c *Config) Validate() (err error) {
	if c == nil				{return errors.New("nil log config")}

	if c.Level == "" 		{return errors.New("nil level config")}
	//if c.Formatter == nil 	{return errors.New("nil formater config")}
	return
}

var cfg Config

func SetConf(config *Config) (err error) {
	if err = config.Validate(); err != nil {return}
	cfg = *config
	return
}

func NewLog() (logger *zap.Logger) {
	var err error

	var log_lvl zap.AtomicLevel
	if log_lvl, err = zap.ParseAtomicLevel(cfg.Level); err != nil {
		panic("Failed to create new conf" + err.Error())
	}

	var encode_config = zapcore.EncoderConfig{
		TimeKey			: "ts",
		LevelKey		: "level",
		NameKey			: "logger",
		CallerKey		: "caller",
		FunctionKey		: zapcore.OmitKey,
		MessageKey		: "msg",
		StacktraceKey	: "stacktrace",
		LineEnding		: zapcore.DefaultLineEnding,
		EncodeLevel		: zapcore.CapitalColorLevelEncoder,
		EncodeTime		: zapcore.RFC3339TimeEncoder,
		EncodeDuration	: zapcore.SecondsDurationEncoder,
		EncodeCaller	: zapcore.ShortCallerEncoder,
	}

	var default_cfg = zap.Config{
		Level				: zap.NewAtomicLevelAt(log_lvl.Level()),
		Development			: true,
		DisableCaller		: false,
		DisableStacktrace	: false,
		Sampling			: nil,
		Encoding			: "console",
		EncoderConfig		: encode_config,
		OutputPaths			: []string{"../logs/logs.txt", "stderr"},
		ErrorOutputPaths	: []string{"stderr"},
		InitialFields		: nil,
	}

	if logger, err = default_cfg.Build(); err != nil {panic(err)}
	return
}
