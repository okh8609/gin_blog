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
curl -X POST "http://localhost:8080/auth/new" -F 'uuid=7ee4e82c-e4e9-44fb-97d8-aeeeaa9d485f' -F 'password=AAA0809'
curl -X POST "http://localhost:8080/auth/verify" -F 'uuid=7ee4e82c-e4e9-44fb-97d8-aeeeaa9d485f' -F 'password=AAA0808'
curl -X POST "http://localhost:8080/auth/verify" -F 'uuid=7ee4e82c-e4e9-44fb-97d8-aeeeaa9d485f' -F 'password=AAA0809'
curl -X PUT "http://localhost:8080/auth" -F 'uuid=7ee4e82c-e4e9-44fb-97d8-aeeeaa9d485f' -F 'password=AAA0800' -F 'new_password=AAA0808'
curl -X PUT "http://localhost:8080/auth" -F 'uuid=7ee4e82c-e4e9-44fb-97d8-aeeeaa9d485f' -F 'password=AAA0809' -F 'new_password=AAA0808'
curl -X POST "http://localhost:8080/auth/verify" -F 'uuid=7ee4e82c-e4e9-44fb-97d8-aeeeaa9d485f' -F 'password=AAA0809'
curl -X POST "http://localhost:8080/auth/verify" -F 'uuid=7ee4e82c-e4e9-44fb-97d8-aeeeaa9d485f' -F 'password=AAA0808'
curl -X DELETE "http://localhost:8080/auth" -F 'uuid=7ee4e82c-e4e9-44fb-97d8-aeeeaa9d485f'
curl -X DELETE "http://localhost:8080/auth" -F 'uuid=7ee4e82c-e4e9-44fb-97d8-aeeeaa9d485f' -F 'password=AAA0808'


```

b4e063ff-b7c9-4264-87d6-f7750b61194b
8626af3e-baa4-45cd-9a25-9fac9903ad14
0a161089-8865-4de6-afaf-f84e6ac70a2d
fd96ace3-c160-4b53-9aac-9d8e1e00b749
9f0fd6ce-08f7-4264-8a72-e0d999450fb7
89b17ed1-93d8-4a39-8ca4-f54632e370b7
b93e8622-ea22-4294-88eb-e91e81b5f3cb
7ea877e6-38f6-4ad1-a1f1-3a77f83297cd
7781dc82-eb38-4dec-8d97-e18ce224b154
39a318d8-16e5-401d-b2aa-5661a41571d8
67c0cbb5-1daa-489f-a754-b997b0d3d1be
66af6059-48c9-419f-a670-e4a6f7ba0d9b
71f3fb13-aff4-4ad0-9617-c2eb6b56506a
5aecfa35-e0b6-4ca5-b81d-c8e3fb83cb36
7ceed46e-0aad-43ff-9921-f6b1a9c71699
63e90b21-d759-472e-980b-5fa2e45aa066
c6ab774f-bf42-45bc-8592-10d6f74f4272
eca15bba-52a2-4a30-8138-d7348730b1eb
61b1baa8-37c9-4470-aa79-95fd93331269
3fdab082-127f-470d-a740-1212da496e83