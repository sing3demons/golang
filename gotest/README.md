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

### Unit Test VS Code Configuration

#### cmd + shift + p : >open Workspace Settings (JSON)

```Unit Test VS Code Configuration
"go.coverOnSave": true,
"go.coverOnSingleTest": true,
"go.coverageDecorator": {
    "type": "gutter",
    "coveredHighlightColor": "rgba(64,128,128,0.5)",
    "uncoveredHighlightColor": "rgba(128,64,64,0.25)",
    "coveredGutterStyle": "blockgreen",
    "uncoveredGutterStyle": "blockred"
}
```
