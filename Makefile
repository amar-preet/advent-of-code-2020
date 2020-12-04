all: env 1 2 3
env:
	@echo "GO ENV"
	GOPATH=${GOPATH} ${GO} env
	@echo ${GO} version
	@echo
1:
	@echo "Running Day 1 code.."
	cd Day1; go run main.go
	@echo
2:
	@echo "Running Day 2 code.."
	cd Day2; go run main.go
	@echo
3:
	@echo "Running Day 3 code.."
	cd Day3; go run main.go		
	@echo
	
	