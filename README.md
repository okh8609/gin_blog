# gin_blog


Article 的 CRUD 業務邏輯未撰寫 ...

~~internal/model/model.go 內的共用軟刪除 callback 未完成...~~
gorm2的callback怪怪的 還要再研究~

### 測試指令參考：
``` bash
# TAGs
curl -X POST "http://127.0.0.1:8080/api/v1/tags" -F 'name=GO' -F 'created_by=khaos'
curl -X POST "http://127.0.0.1:8080/api/v1/tags" -F 'name=C#' -F 'created_by=khaos'
curl -X GET "http://localhost:8080/api/v1/tags?state=1&page=1&page_size=2" -H "accept: application/json"
curl -X PUT "http://localhost:8080/api/v1/tags/1" -F 'state=0' -F 'modified_by=khaos2'
curl -X DELETE "http://localhost:8080/api/v1/tags/2" -H "accept: application/json"

# File
curl -X POST "http://localhost:8080/upload/file" -F "type=1" -F "file=@/home/kh/img_gin.png"

# Auth (user)
curl -X POST "http://localhost:8080/auth/new" -F 'uuid=b4f2b241-4aac-4c6f-a406-dd2cf4c303f5' -F 'password=AAA080'
curl -X POST "http://localhost:8080/auth/verify" -F 'uuid=b4f2b241-4aac-4c6f-a406-dd2cf4c303f5' -F 'password=AAA080'
curl -X PUT "http://localhost:8080/auth" -F 'uuid=b4f2b241-4aac-4c6f-a406-dd2cf4c303f5' -F 'password=AAA0800'
curl -X POST "http://localhost:8080/auth/verify" -F 'uuid=b4f2b241-4aac-4c6f-a406-dd2cf4c303f5' -F 'password=AAA080'
curl -X POST "http://localhost:8080/auth/verify" -F 'uuid=b4f2b241-4aac-4c6f-a406-dd2cf4c303f5' -F 'password=AAA0800'
curl -X DELETE "http://localhost:8080/auth" -F 'uuid=b4f2b241-4aac-4c6f-a406-dd2cf4c303f5'







```
a1137b46-1b88-48e1-a773-1aa710de848a

6a36bf75-18cf-4162-88ee-5688932ec94d