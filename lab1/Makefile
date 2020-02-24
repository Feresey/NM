.PHONY: all clean test
TESTS = $(wildcard tests/test-*)

all:
	@mkdir -p ./build/
	@go build -v -o ./build/

clean:
	@rm -vrf ./build/

test: $(TESTS)
	$(foreach test,$^, ./build/lab1 < $(test) && sleep 1;)