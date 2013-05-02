package imagick

type ProgressMonitor func(text string, offset, span int, data string) bool
