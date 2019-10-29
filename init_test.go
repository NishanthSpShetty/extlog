package extlog

import "testing"

func TestInit(t *testing.T) {
	want := true
	if got := Init(); got != want {
		t.Errorf("got %v ", got)
	}

}

//text:="2019/10/29 16:45:33 /Users/nishanth/projects/programs/golang/golog/main.go:15: Logging..."

func TestFileLenTest(t *testing.T) {
	text := "/Users/nishanth/projects/programs/golang/golog/main.go:15: Logging..."

	//calculated length of long file name string
	want := 59

	got := getFileLen(text, log.Lshortfile)
	if got != want {
		t.Errorf("want %d got %d", got, want)
	}
}
