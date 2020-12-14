all: env run
env:
	@echo "GO ENV"
	GOPATH=${GOPATH} ${GO} env
	@echo ${GO} version
	@echo
run:
	l='day-01 day-02 day-03 day-04 day-05 day-06 day-07 day-08 day-09 day-10 day-11 day-12 day-13'; \
	for k in $$l; do cd $$k; pwd; echo "Running $$k code.."; go run main.go; cd ..;  \
	done		  
