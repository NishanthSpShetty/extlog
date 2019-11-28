package extlog

import (
	"io"
	"log"
	"os"
	"strings"
)

//LogWriter log writer wrapper over io.Writer
type LogWriter struct {

	//internal writer to be used by the writer which should be set by the application
	internal_writer io.Writer
	ServiceName     string
	Flags           int

	//we need to read the log flags set, but if the extlog initialized before it was set
	//in application we will end up reading default flags, so we will read from the log itself.
	//flags int
}

// Field represent the log fields
type Field struct {
	//meta fields
	Timestamp string
	Filename  string
	Message   string
}

func getFileLen(text string, flags int) int {
	/*
		 Lshortfile|Llongfile
			-- we have to read till we find ': '
	*/

	file_len := 0

	if flags&(log.Lshortfile|log.Llongfile) == 0 {
		//file name is not present
		return file_len

	}

	//iterate till we find the ': ' in the texts,
	prev_char := rune(0)
	index := 0
	found := false

	for idx, curr_char := range text {
		if curr_char == ' ' && prev_char == ':' {
			index = idx
			found = true
			break
		}
		prev_char = curr_char
	}

	if !found {
		//no file information were sent ?
		return 0
	}

	//-1 to remove the tailing :SPACE"
	return file_len + index - 1
}

func getTimestampLen(flags int) int {

	/**
	order of flag evaluation in log
	 Ldate|Ltime|Lmicroseconds
		-- Ldate (8)+SPACE +2'\'
		-- Ltime (6)+SPACE +2'.'
		-- Lmicroseconds 1'.'+(6)
	*/

	timestamp_len := 0
	if flags&(log.Ldate) != 0 {
		timestamp_len = timestamp_len + 11
	}

	if flags&(log.Ltime) != 0 {
		timestamp_len = timestamp_len + 9
	}

	if flags&(log.Lmicroseconds) != 0 {
		timestamp_len = timestamp_len + 7
	}

	return timestamp_len
}

//extractLogField extract the log meta data details from the text
func extractLogField(data []byte, flags int) Field {

	//timestamp field
	timestamp_len := getTimestampLen(flags)

	text := string(data)

	//skip the tailing space ?
	timestamp := text[:timestamp_len-1]

	file_name_end := timestamp_len

	filename := ""
	//extract the filename and line number
	if flags&(log.Lshortfile|log.Llongfile) != 0 {
		filename_start := (timestamp_len)

		filename_len := getFileLen(text[filename_start:], flags)

		file_name_end = filename_start + filename_len

		filename = text[filename_start:file_name_end]

		//increase the index to remove :SPACE
		file_name_end = file_name_end + 2
	}

	//get the log message

	//trim the last new line char added by logger
	lastCharIndex := len(text)

	newline := '\n'
	if rune(text[lastCharIndex-1]) == newline {
		lastCharIndex = lastCharIndex - 1
	}
	msg := text[file_name_end:lastCharIndex]
	return Field{Timestamp: timestamp, Filename: filename, Message: msg}
}

func escapeSpecialChar(text string) string {
	//escape the escape('\') character before escaping any other character
	text = strings.Replace(text, "\\", "\\\\", -1)
	//escape double quote and return the string for now
	return strings.Replace(text, "\"", "\\\"", -1)
}

func toJSON(serviceName string, field *Field) string {
	return `{"service_name": "` + serviceName + `", "timestamp": "` + field.Timestamp + `", "file": "` + field.Filename + `", "message": "` + escapeSpecialChar(field.Message) + `"}`
}

func (logWriter LogWriter) Write(data []byte) (n int, err error) {

	//extract log meta content from the data
	field := extractLogField(data, logWriter.Flags)

	text := toJSON(logWriter.ServiceName, &field)

	//dump the bytes fom text
	n, err = logWriter.internal_writer.Write([]byte(text + "\n"))

	return
}

//SetupLogger setup standard logger with our custom logging properties
func SetupLogger(writer io.Writer, serviceName string, flags int) {

	log.SetOutput(LogWriter{internal_writer: writer, ServiceName: serviceName, Flags: flags})

}

//Init initialize the logger using stderr as a output file
func Init(serviceName string, flags int) bool {
	SetupLogger(os.Stderr, serviceName, flags)
	return true
}
