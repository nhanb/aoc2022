watch:
	find -name '*.go' -or -name '*.input' | entr -rc -s 'date && go build'
