## extlog

	log extension module to support structured logging over a golang standard library for existing services.



Simple module to set the custom io writer for `log` package which will write the log message and its meta data in the JSON format without changing any existing application codebase.

### Usage 

at the start of the application you need to initialize extlog module, which will currently set the writer to our custom writer wrapper written over io.Writer,
currently internal writer is set to os.Stderr as default value. 

All write to log writer will be written to our custom writer sink which will convert the given log message into json data set and write it to internal writer.

```
import "source.golabs.io/nishanth.shetty/extlog"

//either in init() or main() call the below functions
extlog.Init(log.LstdFlags | log.Lshortfile)
```

`extlog.Init()` takes logger flags as input argument in order to parse the log written, extlog should be initialized with the `log.Flags()` value


### example
```
  1 package main
  2 
  3 import (
  4         "log"
  5 
  6         "source.golabs.io/nishanth.shetty/extlog"
  7 )
  8 
  9 func init() {
 10         extlog.Init(log.LstdFlags | log.Lshortfile)
 11 }
 12 
 13 func main() {
 14 
 15         log.SetFlags(log.LstdFlags | log.Lshortfile)
 16         log.Print("Logging...")
 17         log.Println("application loaded successfully")
 18 
 19 }
```

### output 
```
{"timestamp": "2019/10/29 22:47:23", "file": "main.go:16", "message": "Logging..."}
{"timestamp": "2019/10/29 22:47:23", "file": "main.go:17", "message": "application loaded successfully"}

```
