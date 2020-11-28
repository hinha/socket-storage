package vo

type UploadFile struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Result []byte `json:"result"`
}
