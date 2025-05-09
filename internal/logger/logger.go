package logger

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
)

var (
	systemPrefix  = lipgloss.NewStyle().SetString("[ SYSTEM ]").Foreground(lipgloss.Color("#0ea5e9"))
	warningPrefix = lipgloss.NewStyle().SetString("[ WARNING ]").Foreground(lipgloss.Color("#f59e0b"))
	errorPrefix   = lipgloss.NewStyle().SetString("[ ERROR ]").Foreground(lipgloss.Color("#f43f5e"))
	userPrefix    = lipgloss.NewStyle().SetString("[ USER ]").Foreground(lipgloss.Color("#14b8a6"))
	agentPrefix   = lipgloss.NewStyle().SetString("[ AGENT ]").Foreground(lipgloss.Color("#d946ef"))
	timeFormatter = lipgloss.NewStyle().Foreground(lipgloss.Color("#737373")).Render
	keyFormatter  = lipgloss.NewStyle().Foreground(lipgloss.Color("#737373")).Render
)

type Logger struct {
	writer io.Writer
}

func New() *Logger {
	return &Logger{writer: os.Stdout}
}

func (l *Logger) System(message string, args ...any) {
	fmt.Fprintf(l.writer, "%s %s %s %s\n", getTimestamp(), systemPrefix, message, formatKeyValueArgs(args))
}

func (l *Logger) Warning(message string, args ...any) {
	fmt.Fprintf(l.writer, "%s %s %s %s\n", getTimestamp(), warningPrefix, message, formatKeyValueArgs(args))
}

func (l *Logger) Error(message string, args ...any) {
	fmt.Fprintf(l.writer, "%s %s %s %s\n", getTimestamp(), errorPrefix, message, formatKeyValueArgs(args))
}

func (l *Logger) User(message string, args ...any) {
	fmt.Fprintf(l.writer, "%s %s %s %s\n", getTimestamp(), userPrefix, message, formatKeyValueArgs(args))
}

func (l *Logger) Agent(message string, args ...any) {
	fmt.Fprintf(l.writer, "%s %s %s %s\n", getTimestamp(), agentPrefix, message, formatKeyValueArgs(args))
}

func formatKeyValueArgs(args []any) string {
	if len(args) == 0 {
		return ""
	}

	var sb strings.Builder
	for i := 0; i < len(args); i += 2 {
		if i+1 >= len(args) {
			break
		}

		key, ok := args[i].(string)
		if !ok {
			continue
		}

		value := args[i+1]
		switch v := value.(type) {
		case string:
			sb.WriteString(keyFormatter(key) + "=")
			sb.WriteString(fmt.Sprintf("%q", v))
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, bool:
			sb.WriteString(keyFormatter(key) + "=")
			sb.WriteString(fmt.Sprintf("%+v", v))
		default:
			jsonBytes, err := json.MarshalIndent(value, "", " ")
			if err == nil && len(jsonBytes) > 2 && (jsonBytes[0] == '{' || jsonBytes[0] == '[') {
				sb.WriteString("\n" + keyFormatter(key) + "=")
				lines := strings.Split(string(jsonBytes), "\n")
				for i, line := range lines {
					if i > 0 && i < len(lines)-1 {
						sb.WriteString(" ")
					}
					sb.WriteString(line)
					if i < len(lines)-1 {
						sb.WriteString("\n")
					}
				}
			} else {
				sb.WriteString(keyFormatter(key) + "=")
				sb.WriteString(fmt.Sprintf("%+v", v))
			}
		}

		if i != len(args)-2 {
			sb.WriteString(" ")
		}
	}
	return sb.String()
}

func getTimestamp() string {
	return timeFormatter(time.Now().Format("15:04:05"))
}
