ifeq ($(OS), Windows_NT)
	EXECUTABLE=abc2.exe
	CLEAN_COMMAND=rmdir
	CLEAN_FLAGS=/s /q
	SRC_DIR=.\src
	RUN_CMD=.\bin\abc2.exe
else
	EXECUTABLE=abc2
	CLEAN_COMMAND=rm -rf
	SRC_DIR=./src
	RUN_CMD=./bin/abc2
endif

build:
	cd $(SRC_DIR); \
	go build -v -o ../bin/$(EXECUTABLE) main.go

run:
	make build && \
	$(RUN_CMD) $(ARGS)

clean:
	$(CLEAN_COMMAND) "bin" $(CLEAN_FLAGS)