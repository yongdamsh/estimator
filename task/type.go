package task

type Feature struct {
	Name string
}

type Task struct {
	Feature Feature
	Name    string
	OrigEst string
	CurEst  string
	Elapsed string
}
