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
	ModelID       string `json:"modelID"`
	API           string `json:"api"`
	Version       string `json:"version"`
	Owner         string `json:"owner"`
	SubmissionTime string `json:"submissionTime"`
	CallID        string `json:"callID"`
	Description   string `json:"description"`
	ModelName     string `json:"modelName"`  
}


// 历史查询结果结构体
type HistoryQueryResult struct {
	Record    *Model `json:"record"`
	TxId      string `json:"txId"`
	Timestamp string `json:"timestamp"`
	IsDelete  bool   `json:"isDelete"`
}


