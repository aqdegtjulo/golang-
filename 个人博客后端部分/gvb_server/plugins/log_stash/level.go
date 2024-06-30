package log_stash

import "encoding/json"

type Level int

const (
	DebugLevel   Level = 0
	InfoLevel    Level = 1
	WarningLevel Level = 2
	ErrorLevel   Level = 3
)

func (s Level) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.string())
}

func (s Level) string() string {
	var str string
	switch s {
	case DebugLevel:
		str = "Debug"
	case InfoLevel:
		str = "Info"
	case WarningLevel:
		str = "Warning"
	case ErrorLevel:
		str = "Error"
	default:
		str = "其他"
	}
	return str

}
