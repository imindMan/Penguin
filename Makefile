all: make

CC = go
OUTPUT_FOLDER = ./bin 
SOURCES = ./cmd/penguin

make:
	mkdir $(OUTPUT_FOLDER)
	$(CC) build $(SOURCES)
	mv ./penguin $(OUTPUT_FOLDER)
