.PHONY: all clean run

GO := go
TARGET := hello/main

${TARGET}:
	cd hello && ${GO} build main.go

run: ${TARGET}
	cd hello && ${GO} run main.go

clean:
	$(RM) ${TARGET}

all: ${TARGET}
