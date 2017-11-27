package lib

const activateBashFileName string = "activate.sh"
const activateBashContents string = `#!/bin/bash

PARENT_PATH="$( cd "$( dirname $( dirname "${BASH_SOURCE[0]}" ) )" && pwd )"

if [ "${GO_ENV}" == "" ]; then
    GO_ENV=${PARENT_PATH}
    echo Activating goenv in ${GO_ENV}
    export GO_ENV

    GOPATH=${GO_ENV}
    export GOPATH

    BIN_PATH=${GOPATH}/bin
    OLD_PATH=${PATH}
    PATH=${PATH}:${BIN_PATH}
    export OLD_PATH
    export PATH
elif [ "${GO_ENV}" != ${PARENT_PATH} ]; then
    GO_ENV=${PARENT_PATH}
    echo Switching to goenv in ${GO_ENV}
    GOPATH=${GO_ENV}
    export GOPATH

    BIN_PATH=${GOPATH}/bin
    PATH=${OLD_PATH}:${BIN_PATH}
    export PATH
else
    echo Already activated goenv in ${GO_ENV}
fi
`
