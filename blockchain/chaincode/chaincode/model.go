package chaincode

/*
定义用户结构体
用户ID
用户类型
实名认证信息哈希,包括用户注册的姓名、身份证号、手机号、注册平台同意协议签名的哈希
模型列表
*/
type User struct {
	UserID       string   `json:"userID"`
	UserType     string   `json:"userType"`
	RealInfoHash string   `json:"realInfoHash"`
	ModelList    []*Model `json:"modelList"`
}

/*
定义大模型结构体
溯源码
研发者输入
发布者输入
共享者输入
使用者输入
*/
type Model struct {
	Traceability_code string        `json:"traceability_code"`
	Developer_input      Developer_input  `json:"developer_input"`
	Publisher_input     Publisher_input `json:"publisher_input"`
	Sharer_input      Sharer_input  `json:"sharer_input"`
	User_input        User_input    `json:"user_input"`
}

// 用于处理历史查询结果的 HistoryQueryResult 结构
type HistoryQueryResult struct {
	Record    *Model `json:"record"`
	TxId      string `json:"txId"`
	Timestamp string `json:"timestamp"`
	IsDelete  bool   `json:"isDelete"`
}

/*
研发者
研发者的溯源码（自动生成）
AI模型名称
研发批次
训练参数
研发机构或组织
*/
type Developer_input struct {
	De_TraceCode       string `json:"traceCode"`       
	De_AIModelName     string `json:"aiModelName"`     
	De_PublishTime      string `json:"publishTime"`
	De_DevelopmentBatch string `json:"developmentBatch"`   
	De_TrainingParams  string `json:"trainingParams"`  
	De_ResearchOrg     string `json:"researchOrg"` 
	De_Txid      string `json:"de_txid"`
	De_Timestamp string `json:"de_timestamp"`    
}

/*
发布者
发布组织名称
发布平台
发布信息
曾发布其他模型记录
联系方式
*/
type Publisher_input struct {
	Pu_OrganizationName   string `json:"organizationName"` 
	Pu_Platform           string `json:"platform"`           
	Pu_PublishInfo        string `json:"publishInfo"`        
	Pu_OtherModelsHistory string `json:"otherModelsHistory"` 
	Pu_ContactInfo        string `json:"contactInfo"`
	Pu_Txid      string `json:"pu_txid"`
	Pu_Timestamp string `json:"pu_timestamp"`        
}

/*
共享者
共享时间
使用时间
电话
使用单位或组织
使用单位联系方式
*/
type Sharer_input struct {
	Sh_SharingTime       string `json:"sharingTime"`       
	Sh_UsageTime         string `json:"usageTime"`         
	Sh_PhoneNumber       string `json:"phoneNumber"`      
	Sh_UsingOrganization string `json:"usingOrganization"` 
	Sh_ContactInfo       string `json:"contactInfo"`
	Sh_Txid	  string `json:"sh_txid"`
	Sh_Timestamp string `json:"sh_timestamp"`       
}

/*
反馈者
反馈时间
销售时间
反馈用户名称
反馈用户位置
反馈用户联系方式
*/
type User_input struct {
	U_FeedbackTime      string `json:"feedbackTime"`      
	U_SalesTime         string `json:"salesTime"`         
	U_UserName          string `json:"userName"`          
	U_UserLocation      string `json:"userLocation"`      
	U_UserContactInfo   string `json:"userContactInfo"`
	U_Txid      string `json:"u_txid"`
	U_Timestamp string `json:"u_timestamp"`   
}
