.PHONY: run
run:
	go build ./cmd/forger
	./forger
	rm ./forger
