package v1

import (
	"time"

	"k8s.io/klog"
)

// loggingFlags appear to be coming from klog. These should be universal for
// all controllers.
type loggingFlags struct {
	VerbosityLevel    klog.Level    `json:"v"`
	VModule           string        `json:"vmodule"` // accepting string because original type ModuleSpec is unexported from klog
	STDERRThreshold   severity      `json:"stderrthreshold"`
	AlsoLogToSTDERR   bool          `json:"alsologtostderr"`
	LogFlushFrequency time.Duration `json:"log-flush-frequency"`
	LogBacktraceAt    traceLocation `json:"log_backtrace_at"`
	LogDir            string        `json:"log_dir"`
	LogFile           string        `json:"log_file"`
	LogFileMaxSize    uint          `json:"log_file_max_size"`
	SkipHeaders       bool          `json:"skip_headers"`
	SkipLogHeaders    bool          `json:"skip_log_headers"`
}

// traceLocation is an emulation of the unexported struct traceLocation in klog https://github.com/kubernetes/klog/blob/master/klog.go#L325
type traceLocation struct {
	file string
	line int
}

// severity is an emulation of the unexported struct severity in klog https://github.com/kubernetes/klog/blob/master/klog.go#L98
type severity int32
