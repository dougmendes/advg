run:
    go run cmd/api/main.go

enviroment:
    docker-compose up -d

tdd path="pkg":
    go test watch --race -r {{path}}

unit path="pkg":
    ginkgo run -r --randomize-all --randomize-suites --fail-on-pending --keep-going --cover --coverprofile=coverprofile.out --race --trace --timeout=1m {{path}}

mockgen:
    go generate ./pkg/...