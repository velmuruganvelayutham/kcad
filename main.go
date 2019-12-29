package main
import (
 "log"
 "os" 
"time"
)

func main() {
 f, err := os.OpenFile("date.out", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
 defer f.Close()
 if err != nil {
 log.Fatal("error opening the file %s", f) 
}
 for { 
now := time.Now()
f.WriteString(now.String() + "\n ")
f.Sync()
time.Sleep(2000 * time.Millisecond) 
log.Printf("sleeping \n")
}
}
