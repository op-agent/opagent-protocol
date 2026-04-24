package op

type SessionReviewEntry struct {
	SessionEntryBase
	TurnID string               `json:"turnID"`
	Status ChatReviewTurnStatus `json:"status"`
}
