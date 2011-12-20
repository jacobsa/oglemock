include $(GOROOT)/src/Make.inc

TARG = github.com/jacobsa/oglemock
GOFILES = \
	action.go \
	call_expectation.go \
	controller.go \
	return.go \

include $(GOROOT)/src/Make.pkg
