include $(GOROOT)/src/Make.inc

TARG = github.com/jacobsa/oglemock
GOFILES = \
	call_expectation.go \

include $(GOROOT)/src/Make.pkg
