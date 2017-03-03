reports:
	go build
	./gotableReports

clean:
	rm -f *.html *.txt *.pdf *.csv
	rm -f gotableReports

all: reports
	@echo "Reports generated"