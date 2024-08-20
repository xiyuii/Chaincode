package chaincode

// 定义用户结构体
type User struct {
	UserID       string   `json:"userID"`
	UserType     string   `json:"userType"`
	Wallet       string   `json:"wallet"`
	Models       []*Model `json:"models"`
}

// 定义模型结构体
type Model struct {
	ModelID      string `json:"modelID"`      // 提交模型的名称
	API            string `json:"api"`            // 调用所需的API
	Version      string     `json:"version"`
	Owner        string	    `json:"owner"`
	SubmissionTime string `json:"submissionTime"` // 提交时间
	Description  string     `json:"description"`
}

// 历史查询结果结构体
type HistoryQueryResult struct {
	Record    *Model `json:"record"`
	TxId      string `json:"txId"`
	Timestamp string `json:"timestamp"`
	IsDelete  bool   `json:"isDelete"`
}


