package vo

type UploadByteFile struct {
	Result []byte `json:"result"`
}

type UploadFile struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	UserID int    `json:"user_id"`
	Result []byte `json:"result"`
}

//{"action":"upload file","domain":"s3-socket","level":"info","msg":".csv","repository":"PutObject","time":"2020-11-30T15:43:44+07:00","type":"Success Upload Object"}
type Logger struct {
	Action  string `json:"action"`
	Level   string `json:"level"`
	Message string `json:"msg"`
	Time    string `json:"time"`
	Type    string `json:"type"`
}
