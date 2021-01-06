ifeq ($(OS), Windows_NT)
	EXECUTABLE=abc2.exe
	CLEAN_COMMAND=rmdir
	CLEAN_FLAGS=/s /q
else
	EXECUTABLE=abc2
	CLEAN_COMMAND=rm -rf
endif

build:
	go build -v -o bin/$(EXECUTABLE) src/main.go

run:
	go run src/main.go

clean:
	$(CLEAN_COMMAND) "bin" $(CLEAN_FLAGS)