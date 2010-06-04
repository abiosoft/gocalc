include $(GOROOT)/src/Make.$(GOARCH)

TARG=calculator
GOFILES=\
	calculator.go\
	gui.go\

include $(GOROOT)/src/Make.cmd