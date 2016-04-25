:: Copyright 2016 fluter

@echo off

set exe=gopaste.exe

set action=all

if not -%1-==-- set action=%1

2>NUL CALL :make_%action%
IF ERRORLEVEL 1 CALL :default
exit /b

:make_all
	go build -o %exe% cmd/paste.go
	goto end
:make_test
	go test -v -cover -covermode atomic -coverprofile coverage.out
	goto end
:make_clean
	del %exe%
	goto end
:default
	echo make.bat: *** No rule to make target '%action%'.  Stop.

:end
	VER > NULL
	GOTO :EOF
