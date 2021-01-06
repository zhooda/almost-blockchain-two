ifeq ($(OS), Windows_NT)
	EXECUTABLE=abc2.exe
	CLEAN_COMMAND=rmdir
	CLEAN_FLAGS=/s /q
	SRC_DIR=.\src
else
	EXECUTABLE=abc2
	CLEAN_COMMAND=rm -rf
	SRC_DIR=./src
endif

build:
	cd $(SRC_DIR); \
	go build -v -o ../bin/$(EXECUTABLE) main.go

run:
	cd $(SRC_DIR); \
	go run main.go

clean:
	$(CLEAN_COMMAND) "bin" $(CLEAN_FLAGS)