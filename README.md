# asynchronous-processing-practice

## Go
 * API server
 * queue / worker pool

## Python
 * task execution
 * business logic
 * failure handling

## File Structure
```
async-task-system/
├── go-server/
│   ├── cmd/
│   │   └── main.go        # HTTP entrypoint
│   ├── api/
│   │   ├── common/                # 共用的 handlers
│   │   │   └── health.go
│   │   ├── v1/
│   │   │   ├── tasks/
│   │   │   │   ├── create.go
│   │   │   │   ├── get.go
│   │   │   │   └── list.go
│   │   │   └── routes.go
│   │   ├── v2/
│   │   │   ├── tasks/
│   │   │   │   ├── create.go      # v2 新功能
│   │   │   │   ├── get.go
│   │   │   │   ├── list.go
│   │   │   │   └── batch.go       # v2 才有的批次功能
│   │   │   └── routes.go
│   │   │── router.go
│   │   ├── queue/
│   │   │   ├── publisher.go   # 業務邏輯層 (The Abstraction)
│   │   │   └── redis.go       # 具體的實作層 (The Implementation)
│   │   ├── worker/
│   │   │   └── pool.go        # worker pool (goroutines)
│   │   └── task/
│   │       └── model.go       # Task schema / status
│   └── go.mod
│
├── python-worker/
│   ├── app/
│   │   ├── main.py            # worker entrypoint
│   │   ├── executor.py        # dispatch by task type
│   │   ├── tasks/
│   │   │   ├── send_email.py
│   │   │   └── generate_report.py
│   │   └── errors.py          # retry / failure handling
│   └── requirements.txt
│
├── protocol/
│   └── task_schema.json       # shared contract
│
├── docker-compose.yml         # optional
└── README.md
```

## Data Pipeline
```
main.go
  ↓
router.go
  ↓
v1/routes.go   v2/routes.go
  ↓
handlers (create.go, get.go, ...)
  ↓
queue / worker / task
```


