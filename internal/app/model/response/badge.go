package response

import "backend-go/internal/app/model"

type PermitClaimBadgeRes struct {
	Func      string           `json:"func"`
	QuestData *model.Extradata `json:"quest_data,omitempty"`
	Uri       string           `json:"uri"`
	Sign      string           `json:"sign"`
}
