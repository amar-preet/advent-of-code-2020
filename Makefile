all: env run
env:
	@echo "GO ENV"
	GOPATH=${GOPATH} ${GO} env
	@echo ${GO} version
	@echo
run:
	l='day-1 day-2 day-3 day-4 day-5 day-6 day-7 day-8 day-9'; \
	for k in $$l; do cd $$k; pwd; echo "Running $$k code.."; go run main.go; cd ..;  \
	done		  
