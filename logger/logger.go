package logger

import (
    "log"
    "net/http"
    "time"
	"runtime"
)

func Logger(inner http.Handler, name string) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        //time of the incoming request so we can calculate the duration until response
        start := time.Now()

        inner.ServeHTTP(w, r)

        //output to the log console
        log.Printf(
            "REQUEST %s\t%s\t%s\t%s",
            r.Method,
            r.RequestURI,
            name,
            time.Since(start),
        )
    })
}

func WriteLog(msg string, lType string) {
	pc := make([]uintptr, 10)
    runtime.Callers(2, pc)

	log.Printf(
		"%s\t%s\t%s",
		lType,
        runtime.FuncForPC(pc[1]).Name(),
        msg,
    )
}

func Info(msg string) {
	WriteLog(msg, "INFO")
}

func Debug(msg string) {
	WriteLog(msg, "DEBUG")
}

func Error(msg string) {
	WriteLog(msg, "ERROR")
}