all: env 1 2 3 4
env:
	@echo "GO ENV"
	GOPATH=${GOPATH} ${GO} env
	@echo ${GO} version
	@echo
1:
	@echo "Running Day 1 code.."
	cd day1; go run main.go
	@echo
2:
	@echo "Running Day 2 code.."
	cd day2; go run main.go
	@echo
3:
	@echo "Running Day 3 code.."
	cd day3; go run main.go		
	@echo
4:
	@echo "Running Day 4 code.."
	cd day4; go run main.go		
	@echo	
	