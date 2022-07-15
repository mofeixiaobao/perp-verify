
all: install

install: go.sum
		go install -mod=readonly  ./verify

go.sum: go.mod
		@echo "--> Ensure dependencies have not been modified"
		GO111MODULE=on go mod verify
