package jobqueue

import "encoding/json"

type Job struct {
	ID   string            `json:"id"`
	Type string            `json:"type"`
	Data map[string]string `json:"data"`
}

func (j *Job) ToJSON() []byte {
	b, _ := json.Marshal(j)
	return b
}

func FromJSON(data string) Job {
	var j Job

	_ = json.Unmarshal([]byte(data), &j)

	return j
}
