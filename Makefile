GOTABLE := ${GOPATH}/src/gotable

reports:
	golint *.go
	go vet *.go
	go build
	cp ${GOTABLE}/report.css .
	cp ${GOTABLE}/table.tmpl .

test:
	./gotableReports

clean:
	rm -f *.html *.txt *.pdf *.csv *.css *.css.map *.tmpl
	rm -rf .sass-cache
	rm -f gotableReports

all: clean reports test
	@echo "Reports generated"