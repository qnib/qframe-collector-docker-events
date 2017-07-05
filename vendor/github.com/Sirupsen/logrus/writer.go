package logrus

import (
	"bufio"
	"io"
	"runtime"
)

func (logger *Logger) Writer() *io.PipeWriter {
	return logger.WriterLevel(InfoLevel)
}

func (logger *Logger) WriterLevel(level Level) *io.PipeWriter {
<<<<<<< HEAD
	reader, writer := io.Pipe()

	var printFunc func(args ...interface{})
	switch level {
	case DebugLevel:
		printFunc = logger.Debug
	case InfoLevel:
		printFunc = logger.Info
	case WarnLevel:
		printFunc = logger.Warn
	case ErrorLevel:
		printFunc = logger.Error
	case FatalLevel:
		printFunc = logger.Fatal
	case PanicLevel:
		printFunc = logger.Panic
	default:
		printFunc = logger.Print
	}

	go logger.writerScanner(reader, printFunc)
=======
	return NewEntry(logger).WriterLevel(level)
}

func (entry *Entry) Writer() *io.PipeWriter {
	return entry.WriterLevel(InfoLevel)
}

func (entry *Entry) WriterLevel(level Level) *io.PipeWriter {
	reader, writer := io.Pipe()

	var printFunc func(args ...interface{})

	switch level {
	case DebugLevel:
		printFunc = entry.Debug
	case InfoLevel:
		printFunc = entry.Info
	case WarnLevel:
		printFunc = entry.Warn
	case ErrorLevel:
		printFunc = entry.Error
	case FatalLevel:
		printFunc = entry.Fatal
	case PanicLevel:
		printFunc = entry.Panic
	default:
		printFunc = entry.Print
	}

	go entry.writerScanner(reader, printFunc)
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
	runtime.SetFinalizer(writer, writerFinalizer)

	return writer
}

<<<<<<< HEAD
func (logger *Logger) writerScanner(reader *io.PipeReader, printFunc func(args ...interface{})) {
=======
func (entry *Entry) writerScanner(reader *io.PipeReader, printFunc func(args ...interface{})) {
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		printFunc(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
<<<<<<< HEAD
		logger.Errorf("Error while reading from Writer: %s", err)
=======
		entry.Errorf("Error while reading from Writer: %s", err)
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
	}
	reader.Close()
}

func writerFinalizer(writer *io.PipeWriter) {
	writer.Close()
}
