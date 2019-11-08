## extlog

Log extension module to reformat text logs emitted by Go `log` module as structured JSON logs.

This library transparently converts the log format, hence you may not need to change existing code that uses the `log` module.




### Usage 

At the start of the application initialize the extlog module, which sets up the logger writer to a custom writer wrapper (os.Stderr by default.)


```
import "source.golabs.io/nishanth.shetty/extlog"
```

either in init() or main() call the below functions

```
extlog.Init(log.LstdFlags | log.Lshortfile)
```

`extlog.Init()` takes logger flags as input argument in order to parse the log written, extlog should be initialized with the `log.Flags()` value

----

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
