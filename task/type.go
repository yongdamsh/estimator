package task

type Feature struct {
	Name string `json:"name"`
}

type Task struct {
	Feature Feature `json:"feature"`
	Name    string  `json:"name"`
	OrigEst string  `json:"originalEstimatedTime"`
	CurEst  string  `json:"currentEstimatedTime"`
	Elapsed string  `json:"elapsedTime"`
}
