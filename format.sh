set -e
set -x

gofmt -w .

go mod tidy
