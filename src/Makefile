SHELL=/bin/bash

all = compile

compile: clean
	@go build main.go

creds:


clean:
ifneq (,$(wildcard ./main.exe))
	@rm ./main.exe 
endif
ifneq (,$(wildcard ./main))
	@rm ./main 
endif