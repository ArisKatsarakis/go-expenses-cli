package model

type CommandExec struct {
	ExecuteName  string
	Name         string
	ExecutedFunc func()
}
