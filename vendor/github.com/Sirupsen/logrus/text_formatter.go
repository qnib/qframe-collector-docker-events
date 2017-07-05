package logrus

import (
	"bytes"
	"fmt"
<<<<<<< HEAD
	"runtime"
	"sort"
	"strings"
=======
	"sort"
	"strings"
	"sync"
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
	"time"
)

const (
	nocolor = 0
	red     = 31
	green   = 32
	yellow  = 33
	blue    = 34
	gray    = 37
)

var (
	baseTimestamp time.Time
<<<<<<< HEAD
	isTerminal    bool
=======
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
)

func init() {
	baseTimestamp = time.Now()
<<<<<<< HEAD
	isTerminal = IsTerminal()
}

func miniTS() int {
	return int(time.Since(baseTimestamp) / time.Second)
=======
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
}

type TextFormatter struct {
	// Set to true to bypass checking for a TTY before outputting colors.
	ForceColors bool

	// Force disabling colors.
	DisableColors bool

	// Disable timestamp logging. useful when output is redirected to logging
	// system that already adds timestamps.
	DisableTimestamp bool

	// Enable logging the full timestamp when a TTY is attached instead of just
	// the time passed since beginning of execution.
	FullTimestamp bool

	// TimestampFormat to use for display when a full timestamp is printed
	TimestampFormat string

	// The fields are sorted by default for a consistent output. For applications
	// that log extremely frequently and don't use the JSON formatter this may not
	// be desired.
	DisableSorting bool
<<<<<<< HEAD
=======

	// QuoteEmptyFields will wrap empty fields in quotes if true
	QuoteEmptyFields bool

	// QuoteCharacter can be set to the override the default quoting character "
	// with something else. For example: ', or `.
	QuoteCharacter string

	// Whether the logger's out is to a terminal
	isTerminal bool

	sync.Once
}

func (f *TextFormatter) init(entry *Entry) {
	if len(f.QuoteCharacter) == 0 {
		f.QuoteCharacter = "\""
	}
	if entry.Logger != nil {
		f.isTerminal = IsTerminal(entry.Logger.Out)
	}
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
}

func (f *TextFormatter) Format(entry *Entry) ([]byte, error) {
	var b *bytes.Buffer
<<<<<<< HEAD
	var keys []string = make([]string, 0, len(entry.Data))
=======
	keys := make([]string, 0, len(entry.Data))
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
	for k := range entry.Data {
		keys = append(keys, k)
	}

	if !f.DisableSorting {
		sort.Strings(keys)
	}
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	prefixFieldClashes(entry.Data)

<<<<<<< HEAD
	isColorTerminal := isTerminal && (runtime.GOOS != "windows")
	isColored := (f.ForceColors || isColorTerminal) && !f.DisableColors
=======
	f.Do(func() { f.init(entry) })

	isColored := (f.ForceColors || f.isTerminal) && !f.DisableColors
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2

	timestampFormat := f.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = DefaultTimestampFormat
	}
	if isColored {
		f.printColored(b, entry, keys, timestampFormat)
	} else {
		if !f.DisableTimestamp {
			f.appendKeyValue(b, "time", entry.Time.Format(timestampFormat))
		}
		f.appendKeyValue(b, "level", entry.Level.String())
		if entry.Message != "" {
			f.appendKeyValue(b, "msg", entry.Message)
		}
		for _, key := range keys {
			f.appendKeyValue(b, key, entry.Data[key])
		}
	}

	b.WriteByte('\n')
	return b.Bytes(), nil
}

func (f *TextFormatter) printColored(b *bytes.Buffer, entry *Entry, keys []string, timestampFormat string) {
	var levelColor int
	switch entry.Level {
	case DebugLevel:
		levelColor = gray
	case WarnLevel:
		levelColor = yellow
	case ErrorLevel, FatalLevel, PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}

	levelText := strings.ToUpper(entry.Level.String())[0:4]

<<<<<<< HEAD
	if !f.FullTimestamp {
		fmt.Fprintf(b, "\x1b[%dm%s\x1b[0m[%04d] %-44s ", levelColor, levelText, miniTS(), entry.Message)
=======
	if f.DisableTimestamp {
		fmt.Fprintf(b, "\x1b[%dm%s\x1b[0m %-44s ", levelColor, levelText, entry.Message)
	} else if !f.FullTimestamp {
		fmt.Fprintf(b, "\x1b[%dm%s\x1b[0m[%04d] %-44s ", levelColor, levelText, int(entry.Time.Sub(baseTimestamp)/time.Second), entry.Message)
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
	} else {
		fmt.Fprintf(b, "\x1b[%dm%s\x1b[0m[%s] %-44s ", levelColor, levelText, entry.Time.Format(timestampFormat), entry.Message)
	}
	for _, k := range keys {
		v := entry.Data[k]
		fmt.Fprintf(b, " \x1b[%dm%s\x1b[0m=", levelColor, k)
		f.appendValue(b, v)
	}
}

<<<<<<< HEAD
func needsQuoting(text string) bool {
=======
func (f *TextFormatter) needsQuoting(text string) bool {
	if f.QuoteEmptyFields && len(text) == 0 {
		return true
	}
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
	for _, ch := range text {
		if !((ch >= 'a' && ch <= 'z') ||
			(ch >= 'A' && ch <= 'Z') ||
			(ch >= '0' && ch <= '9') ||
			ch == '-' || ch == '.') {
			return true
		}
	}
	return false
}

func (f *TextFormatter) appendKeyValue(b *bytes.Buffer, key string, value interface{}) {

	b.WriteString(key)
	b.WriteByte('=')
	f.appendValue(b, value)
	b.WriteByte(' ')
}

func (f *TextFormatter) appendValue(b *bytes.Buffer, value interface{}) {
	switch value := value.(type) {
	case string:
<<<<<<< HEAD
		if !needsQuoting(value) {
			b.WriteString(value)
		} else {
			fmt.Fprintf(b, "%q", value)
		}
	case error:
		errmsg := value.Error()
		if !needsQuoting(errmsg) {
			b.WriteString(errmsg)
		} else {
			fmt.Fprintf(b, "%q", errmsg)
=======
		if !f.needsQuoting(value) {
			b.WriteString(value)
		} else {
			fmt.Fprintf(b, "%s%v%s", f.QuoteCharacter, value, f.QuoteCharacter)
		}
	case error:
		errmsg := value.Error()
		if !f.needsQuoting(errmsg) {
			b.WriteString(errmsg)
		} else {
			fmt.Fprintf(b, "%s%v%s", f.QuoteCharacter, errmsg, f.QuoteCharacter)
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
		}
	default:
		fmt.Fprint(b, value)
	}
}
