curl -d {"text":"First Posting..."} -H "Content-Type: application/json" -X POST 127.0.0.1:8080/posts

curl 127.0.0.1:8080/posts
[{"text":"First Posting...","createdAt":"2019-11-03T11:49:50.317Z"}]
