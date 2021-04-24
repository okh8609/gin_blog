# gin_blog


Article 的 CRUD 業務邏輯未撰寫 ...

~~internal/model/model.go 內的共用軟刪除 callback 未完成...~~
gorm2的callback怪怪的 還要再研究~

### 測試指令參考：
``` bash
curl -X POST "http://127.0.0.1:8080/api/v1/tags" -F 'name=GO' -F 'created_by=khaos'
curl -X POST "http://127.0.0.1:8080/api/v1/tags" -F 'name=C#' -F 'created_by=khaos'

curl -X GET "http://localhost:8080/api/v1/tags?state=1&page=1&page_size=2" -H "accept: application/json"

curl -X PUT "http://localhost:8080/api/v1/tags/1" -F 'state=0' -F 'modified_by=khaos2'

curl -X DELETE "http://localhost:8080/api/v1/tags/2" -H "accept: application/json"
```