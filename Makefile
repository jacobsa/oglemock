include $(GOROOT)/src/Make.inc

TARG = github.com/jacobsa/oglemock
GOFILES = \
	action.go \
	call_expectation.go \

include $(GOROOT)/src/Make.pkg
