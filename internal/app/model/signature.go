package model

type Signature struct {
	MessageId string `json:"messageId"`
	Address   string `json:"address"`
	Uri       string `json:"uri"`
}

type Extradata struct {
	StartTs uint32 `json:"startTs"`
	EndTs   uint32 `json:"endTs"`
	Supply  uint64 `json:"supply"`
}

type QuestData struct {
	Title              string    `json:"title"`
	Description        string    `json:"description"`
	Uri                string    `json:"uri"`
	Metadata           string    `json:"metadata"`
	Extradata          Extradata `json:"extradata"`
	Creator            string    `json:"creator"`
	IsDraft            bool      `json:"isDraft"`
	SubmittedTimestamp int64     `json:"submittedTimestamp"`
}

type QuestDataSign struct {
	StartTs uint64 `json:"startTs"`
	EndTs   uint64 `json:"endTs"`
	Supply  uint64 `json:"supply"`
	Title   string `json:"title"`
	Uri     string `json:"uri"`
	PermitSign
}

type PermitSign struct {
	Contract string `json:"contract"`
	Sender   string `json:"sender"`
}
