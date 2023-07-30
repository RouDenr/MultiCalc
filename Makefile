CXX := g++
GO := go
CXXFLAGS := -std=c++17
BUILD_DIR := build
BIN_DIR := bin
TARGET := $(BIN_DIR)/app
HTML_DIR := html

all: build_model build_view build_viewmodel build_html
	@echo "Build complete. Run './bin/app' to execute."

build_model:
	$(CXX) $(CXXFLAGS) -shared -fPIC -o $(BUILD_DIR)/model.so src/model.cpp
	@echo "C++ Model compiled."

build_view:
	# Add any build steps for the view here (if required)
	@echo "View files processed."

build_viewmodel:
	$(GO) build -o $(BIN_DIR)/app go/main.go go/modelwrapper.go
	@echo "Golang ViewModel compiled."

build_html:
	cp $(HTML_DIR)/* $(BIN_DIR)
	@echo "HTML files copied."

run_server:
	cd $(BIN_DIR) && ./app

clean:
	rm -rf $(BUILD_DIR)/*
	rm -rf $(BIN_DIR)/*
