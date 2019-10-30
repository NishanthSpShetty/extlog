package extlog

import (
	"log"
	"testing"
)

func TestInit(t *testing.T) {
	want := true
	if got := Init(log.Ldate | log.Llongfile); got != want {
		t.Errorf("got %v ", got)
	}

}

//text:="2019/10/29 16:45:33 /Users/nishanth/projects/programs/golang/golog/main.go:15: Logging..."

func TestTimestampLenForLdate(t *testing.T) {
	want := 11
	got := getTimestampLen(log.Ldate)

	if got != want {
		t.Errorf("want %d got %d", got, want)
	}

}

func TestTimestampLenForLtime(t *testing.T) {
	want := 9
	got := getTimestampLen(log.Ltime)

	if got != want {
		t.Errorf("want %d got %d", got, want)
	}

}

func TestTimestampLenForLmicroseconds(t *testing.T) {
	want := 7
	got := getTimestampLen(log.Lmicroseconds)

	if got != want {
		t.Errorf("want %d got %d", got, want)
	}

}

func TestTimestampLen(t *testing.T) {
	want := 27
	got := getTimestampLen(log.Ldate | log.Ltime | log.Lmicroseconds)

	if got != want {
		t.Errorf("want %d got %d", got, want)
	}

}

func TestFileLenLongLength(t *testing.T) {
	text := "/Users/nishanth/projects/programs/golang/golog/main.go:15: Logging..."

	//calculated length of long file name string
	want := 57
	got := getFileLen(text, log.Lshortfile)
	if got != want {
		t.Errorf("want %d got %d", got, want)
	}
}

func TestFileLenWhenEmpty(t *testing.T) {
	text := ""

	want := 0

	got := getFileLen(text, log.Lshortfile)
	if got != want {
		t.Errorf("want %d got %d", got, want)
	}

}

func TestExtractLogField(t *testing.T) {

	text := "2019/10/29 16:45:33 /Users/nishanth/projects/programs/golang/golog/main.go:15: Connected to database"

	field := extractLogField([]byte(text), log.LstdFlags|log.Lshortfile)

	expectedTimestamp := "2019/10/29 16:45:33"
	expectedFilename := "/Users/nishanth/projects/programs/golang/golog/main.go:15"
	expectedMessage := "Connected to database"

	if field.Timestamp != expectedTimestamp {
		t.Errorf("Timestamp =>  want: %s. got: %s. ", expectedTimestamp, field.Timestamp)
	}

	if field.Filename != expectedFilename {
		t.Errorf("Filename  =>  want: %s. got: %s. ", expectedFilename, field.Filename)
	}

	if field.Message != expectedMessage {
		t.Errorf("Message =>  want: %s. got: %s. ", expectedMessage, field.Message)
	}

}

func TestExtractLogFieldWithQuotes(t *testing.T) {

	text := "2019/10/29 16:45:33 main.go:15: user \\\"alex\\\" is logged out from the system"

	field := extractLogField([]byte(text), log.LstdFlags|log.Lshortfile)

	expectedTimestamp := "2019/10/29 16:45:33"
	expectedFilename := "main.go:15"
	expectedMessage := "user \\\"alex\\\" is logged out from the system"

	if field.Timestamp != expectedTimestamp {
		t.Errorf("Timestamp =>  want: %s. got: %s. ", expectedTimestamp, field.Timestamp)
	}

	if field.Filename != expectedFilename {
		t.Errorf("Filename  =>  want: %s. got: %s. ", expectedFilename, field.Filename)
	}

	if field.Message != expectedMessage {
		t.Errorf("Message =>  want: %s. got: %s. ", expectedMessage, field.Message)
	}

}

func TesttoJson(t *testing.T) {
	text := `2019/10/29 16:45:33 /Users/nishanth/projects/programs/golang/golog/main.go:15: Connected to database {user :\"Alex\"} \ and a slash`

	field := extractLogField([]byte(text), log.LstdFlags|log.Lshortfile)

	json := toJSON(&field)

	expected := `{"timestamp": "2019/10/29 16:45:33", "file": "/Users/nishanth/projects/programs/golang/golog/main.go:15", "message": "Connected to database {user :\\"Alex\\"}" \\ and a slash}`

	if expected != json {

		t.Errorf("  want: %s.\n got: %s. ", expected, json)
	}
}
