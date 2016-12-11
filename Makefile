.DEFAULT_GOAL := build
GOVENDOR_TAG := v1.0.8
LINTER_TAG := v1.0.3

build:
	# gox -osarch="windows/386" -output="signtool"
	GOOS=windows GOARCH=386 go build -o signtool.exe main.go

dist: build
	tar -zcvf "mono-signtool.tar.gz" --directory="." "./signtool.exe"

# Setups linter configuration for tests
setup-linter:
	@if [ "$$(which gometalinter)" = "" ]; then \
		go get -u -v github.com/alecthomas/gometalinter; \
		cd $$GOPATH/src/github.com/alecthomas/gometalinter;\
		git checkout tags/$(LINTER_TAG);\
		go install;\
		gometalinter --install;\
	fi

# Runs tests
test: setup-linter
	gometalinter --vendor --fast --errors --dupl-threshold=100 --cyclo-over=25 --min-occurrences=5 --disable=gas --disable=gotype ./...
