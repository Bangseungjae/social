What tech are we using, pretty simple:
- Go 1.22
- Docker
- Postgres running on Docker
- Swagger for docs
- Golang migrate for migrations

//실행
```bash
go run cmd/api/*.go
# or
air

```

```text
go get -u github.com/go-chi/chi/v5

go get github.com/go-chi/cors
```

핫 리로드

https://github.com/air-verse/air
```text
go install github.com/air-verse/air@latest

// start
air init
// then created .air.toml
```
그리고 tmp폴더로 세팅되 있는것들을 bin으로 수정
```text
exclude_dir = ["assets", "bin", "vendor", "testdata", "web", "docs", "scripts"]

cmd = "go build -o ./bin/main ./cmd/api/*.go"
```
- exclude_dir 에서 "web", "docs", "scripts" 추가
- cmd 수정

![img.png](img/img.png)

다음처럼 코드상에서 
```go
addr: ":8080", -> addr: ":3000",
```

으로 바꾸면 핫리로딩되며 서버 주소가 3000번 포트로 다시시작됨

direnv
```text
brew install direnv

// 실행
## BASH -> ~/.bashrc에 아래 내용 추가
eval "$(direnv hook bash)"

## ZSH -> ~/.zshrc에 아래 내용 추가
eval "$(direnv hook zsh)"

direnv allow
```

### DB migrate
```shell
migrate create -seq -ext sql -dir ./cmd/migrate/migrations create_users

go install -tags "postgres,mysql" github.com/golang-migrate/migrate/v4/cmd/migrate@latest

migrate -path=./cmd/migrate/migrations -database="postgres://admin:adminpassword@localhost/socialnetwork?sslmode=disable" up

make migration posts_create

make migrate-up 
```


```shell
psql -U admin -d socialnetwork
\dt
```

### Validator
https://github.com/go-playground/validator

```shell
go get github.com/go-playground/validator/v10

# Then import the validator package into your own code.
import "github.com/go-playground/validator/v10"
```

### Logger
```go
go get -u go.uber.org/zap

// 설탕문법
	logger := zap.Must(zap.NewProduction()).Sugar() // .Sugar로 생성
	app.logger.Infow( // Info -> Infow
		"server has started",
		zap.String("addr", app.config.addr),
		"env", app.config.env,
	)
```


### web
```shell
cd web
npm create vite@latest . -- --template react-ts
```