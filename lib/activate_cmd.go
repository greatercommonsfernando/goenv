package lib

const activateCmdFileName string = "activate.bat"
const activateCmdContents string = `@echo off

CALL :GET_PARENT_PATH "%~dp0" PARENT_PATH
CALL :GET_PARENT_PATH "%PARENT_PATH%" PARENT_PATH

IF "%GO_ENV%"=="%PARENT_PATH%" ( GOTO :NON_OP_GOENV )

IF "%GO_ENV%"=="" ( GOTO :SET_GOENV ) ELSE ( GOTO :SWITCH_GOENV )

GOTO :EOF

:NON_OP_GOENV
echo Already activated goenv in %GO_ENV%
GOTO :EOF

:SET_GOENV
SET GO_ENV=%PARENT_PATH%
SET GOPATH=%GO_ENV%
SET BIN_PATH=%GOPATH%\bin
SET OLD_PATH=%PATH%
SET PATH=%OLD_PATH%;%BIN_PATH%
echo Activating goenv in %GO_ENV%
GOTO :EOF

:SWITCH_GOENV
SET GO_ENV=%PARENT_PATH%
SET GOPATH=%GO_ENV%
SET BIN_PATH=%GOPATH%\bin
SET PATH=%OLD_PATH%;%BIN_PATH%
echo Switching to goenv in %GO_ENV%
GOTO :EOF

:GET_PARENT_PATH
:: use temp variable to hold the path, so we can substring
SET TMP_PARENT_PATH=%~dp1
:: strip the trailing slash, so we can call it again to get its parent
SET %2=%TMP_PARENT_PATH:~0,-1%
GOTO :EOF
`
