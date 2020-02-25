.PHONY: all clean test
TESTS = 

all:
	@mkdir -p ./build/
	@go build -v -o ./build/ ./cmd/...

clean:
	@rm -vrf ./build/

test-1: $(wildcard tests/test-1-*)
	$(foreach test,$^, ./build/lab1-1 < $(test) && sleep 1; echo ;)


test-2: $(wildcard tests/test-2-*)
	$(foreach test,$^, ./build/lab1-2 < $(test) && sleep 1; echo ;)