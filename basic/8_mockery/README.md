# mockery

## mockery  설치
```
# go install github.com/vektra/mockery/v2@latest
brew install mockery
```

## mock_user_db.go 생성
```shell
mockery --all --keeptree
```

## run test
```shell
go test -v
```