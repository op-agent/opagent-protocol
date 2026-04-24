package op

// type ThreadOperationParams struct {
// 	OpType OpType `json:"opType"`
// 	Data   any    `json:"data"`
// }

// type ThreadOperationResult struct {
// 	Meta    `json:"_meta,omitempty"`
// 	Content string `json:"content"`
// }

// func (*ThreadOperationResult) isResult() {}

// func (r *ThreadOperationResult) UnmarshalJSON(data []byte) error {
// 	var result ThreadOperationResult // avoid recursion
// 	if err := json.Unmarshal(data, &result); err != nil {
// 		return err
// 	}
// 	*r = ThreadOperationResult(result)
// 	return nil
// }

// type BindThreadData struct {
// 	TaskID   string `json:"taskID"`
// 	ThreadID string `json:"threadID"`
// }

// type ThreadStorage struct {
// 	ThreadID  string `json:"threadID" bson:"threadID"`
// 	Meta      Meta   `json:"meta,omitempty" bson:"meta,omitempty"`
// 	UID       string `json:"uid,omitempty" bson:"uid,omitempty"`
// 	CreatedAt int64  `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
// 	UpdatedAt int64  `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
// }

// type ThreadStorageQuery struct {
// 	ThreadID      string         `json:"threadID,omitempty" bson:"threadID,omitempty"`
// 	UID           string         `json:"uid,omitempty" bson:"uid,omitempty"`
// 	MetaFilter    map[string]any `json:"metaFilter,omitempty" bson:"metaFilter,omitempty"`
// 	CreatedAtFrom int64          `json:"createdAtFrom,omitempty" bson:"createdAtFrom,omitempty"`
// 	CreatedAtTo   int64          `json:"createdAtTo,omitempty" bson:"createdAtTo,omitempty"`
// 	Limit         int64          `json:"limit,omitempty" bson:"limit,omitempty"`
// 	Offset        int64          `json:"offset,omitempty" bson:"offset,omitempty"`
// 	SortBy        string         `json:"sortBy,omitempty" bson:"sortBy,omitempty"`
// 	Desc          bool           `json:"desc,omitempty" bson:"desc,omitempty"`
// }

// type ThreadStorageQueryResult struct {
// 	Threads []*ThreadStorage `json:"threads"`
// 	Total   int64            `json:"total"`
// 	Limit   int64            `json:"limit,omitempty"`
// 	Offset  int64            `json:"offset,omitempty"`
// }
