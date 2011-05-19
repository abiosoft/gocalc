include $(GOROOT)/src/Make.inc

TARG=calculator
GOFILES=\
	calculator.go\
	gui.go\

include $(GOROOT)/src/Make.cmd
