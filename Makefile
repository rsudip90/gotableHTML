GOTABLE := ${GOPATH}/src/gotable
SCSS_BIN := scss

reports:
	golint *.go
	go vet *.go
	go build
	${SCSS_BIN} ${GOTABLE}/scss/report.scss ./report.css --style=expanded --sourcemap=none

test:
	./gotableReports

clean:
	rm -f *.html *.txt *.pdf *.csv *.css *.css.map
	rm -rf .sass-cache
	rm -f gotableReports

all: clean reports test
	@echo "Reports generated"