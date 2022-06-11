# gotest

### coverage: % of statements

```
go test github.com/sing3demons/gotest/services -cover -v
```

### test services/grade_test.go (เทส เฉพาะฟังก์ชั่น) -run=<func>

```go test services/grade_test.go
go test github.com/sing3demons/gotest/services -v -run=TestCheckGrade
```

### test ทุกที่

```go test services/grade_test.go
go test ./... -v
```
