package errcode

var (
	Success                   = NewError(0, "成功")
	ServerError               = NewError(10000000, "服務內部錯誤")
	InvalidParams             = NewError(10000001, "導入參數錯誤")
	NotFound                  = NewError(10000002, "找不到")
	UnauthorizedAuthNotExist  = NewError(10000003, "驗證失敗，ID或密碼錯誤")
	UnauthorizedTokenError    = NewError(10000004, "驗證失敗，Token錯誤")
	UnauthorizedTokenTimeout  = NewError(10000005, "驗證失敗，Token超時")
	UnauthorizedTokenGenerate = NewError(10000006, "驗證失敗，Token生成失敗")
	TooManyRequests           = NewError(10000007, "請求過多")
)


var (
	ErrorGetTagListFail = NewError(20010001, "獲取標籤列表失敗")
	ErrorCreateTagFail  = NewError(20010002, "創建標籤失敗")
	ErrorUpdateTagFail  = NewError(20010003, "更新標籤失敗")
	ErrorDeleteTagFail  = NewError(20010004, "刪除標籤失敗")
	ErrorCountTagFail   = NewError(20010005, "統計標籤失敗")

	ErrorGetArticleFail    = NewError(20020001, "獲取單個文章失敗")
	ErrorGetArticlesFail   = NewError(20020002, "獲取多個文章失敗")
	ErrorCreateArticleFail = NewError(20020003, "創建文章失敗")
	ErrorUpdateArticleFail = NewError(20020004, "更新文章失敗")
	ErrorDeleteArticleFail = NewError(20020005, "刪除文章失敗")

	ErrorUploadFileFail = NewError(20030001, "上傳文件失敗")

	ErrorGetAuthFail    = NewError(20040001, "取得帳戶失敗")
	ErrorCreateAuthFail = NewError(20040002, "創建帳戶失敗")
	ErrorUpdateAuthFail = NewError(20040003, "更新帳戶失敗")
	ErrorDeleteAuthFail = NewError(20040004, "刪除帳戶失敗")
)
