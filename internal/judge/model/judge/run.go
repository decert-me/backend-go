package judge

type RunFiles struct {
	Content string `json:"content,omitempty"`
	Name    string `json:"name,omitempty"`
	Max     int    `json:"max,omitempty"`
}

type RunMainCode struct {
	Content string `json:"content"`
}

type RunCmd struct {
	Args        []string    `json:"args"`
	Env         []string    `json:"env"`
	Files       []RunFiles  `json:"files"`
	CPULimit    int64       `json:"cpuLimit"`
	ClockLimit  int64       `json:"clockLimit"`
	MemoryLimit int         `json:"memoryLimit"`
	ProcLimit   int         `json:"procLimit"`
	CPURate     float64     `json:"cpuRate"`
	CopyIn      interface{} `json:"copyIn"`
}

type Run struct {
	Cmd []RunCmd `json:"cmd"`
}

type RunResultFiles struct {
	Stderr string `json:"stderr"`
	Stdout string `json:"stdout"`
}
type RunResult struct {
	Status     string         `json:"status"`
	ExitStatus int            `json:"exitStatus"`
	Error      string         `json:"error"`
	Time       int            `json:"time"`
	Memory     int            `json:"memory"`
	RunTime    int            `json:"runTime"`
	Files      RunResultFiles `json:"files"`
}
