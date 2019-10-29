package extlog

import (
	"os"

	"fmt"
	"io"
	"log"
)

type LogWriter struct {

	//internal writer to be used by the writer which should be set by the application
	internal_writer io.Writer

	//we need to read the log flags set, but if the extlog initialized before it was set
	//in application we will end up reading default flags, so we will read from the log itself.
	//flags int
}

type Meta struct {
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
	if flags&(log.Lshortfile|log.Llongfile) != 0 {
		//file name present
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
	}

	if !found {
		//no file information were sent ?
		return file_len
	}

	//+1 to remove the tailing space
	return file_len + index + 1
}

func getTimestampLen(flags int) int {

	/**
	order of flag evaluation in log
	 Ldate|Ltime|Lmicroseconds
		-- Ldate (8)+SPACE +2'\'
		-- Ltime (6)+SPACE +2'.'
		-- Lmicorseconds 1'.'+(6)
	*/

	timestamp_len := 0
	if flags&(log.Ldate) != 0 {
		timestamp_len = timestamp_len + 11
	}

	if flags&(log.Ltime) != 0 {
		timestamp_len = timestamp_len + 9
	}

	if flags&(log.Lmicorseconds) != 0 {
		timestamp_len = timestamp_len + 7
	}

	return timestamp_len
}

//extractLogMeta extract the log meta data details from the text
func extractLogMeta(data []byte, flags int) Meta {

	//timestamp field
	timestamp_len := getTimestampLen(flags)

	//skip the tailing space ?

	text := string(data)

	timestamp := text[:timestamp_len]

	filename_start := (timestamp_len)

	filename_len := getFileLen(text[filename_start:], flags)

	fine_name_end = filename_index + filename_len

	filename = text[filename_start:file_name_end]

	return Meta{Timestamp: timestamp, Filename: filename, MessageIndex: file_name_end}
}

func (logWriter LogWriter) Write(data []byte) (n int, err error) {
	//write the given chunk of data to the writer sink
	//now we will write to std

	fmt.Println("Recived data on LogWriter..")

	//convert the text into json dataset

	//extract log meta content from the data
	extractLogMeta(data, log.Flags())

	n, err = logWriter.internal_writer.Write(data)

	fmt.Println("EndOfLogWriter")

	return
}

type LogMeta struct {
}

var logMeta = LogMeta{}

//SetupLogger setup standard logger with our custom logging properties
func SetupLogger(writer io.Writer) {

	log.SetOutput(LogWriter{internal_writer: writer})
	logMeta.flags = log.Flags()

}

//Init should  change this func signature.
func Init() bool {
	SetupLogger(os.Stderr)
	return true
}

func GetLogMeta() LogMeta {
	return logMeta
}
