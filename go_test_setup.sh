# !bin/sh
# This script is used to setup the environment for running the go tests. with colors

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
BLUE='\033[0;34m'
PURPLE='\033[0;35m'
CYAN='\033[0;36m'
NC='\033[0m' # No Color

# install grc if not installed depending on the OS
if [ -x "$(command -v grc)" ]; then
    echo -e "${GREEN}grc is installed${NC}"
else
    echo -e "${RED}grc is not installed${NC}"
    if [ -x "$(command -v apt-get)" ]; then
        echo -e "${GREEN}Installing grc${NC}"
        sudo apt-get install grc
    elif [ -x "$(command -v yum)" ]; then
        echo -e "${GREEN}Installing grc${NC}"
        sudo yum install grc
    elif [ -x "$(command -v brew)" ]; then
        echo -e "${GREEN}Installing grc${NC}"
        brew install grc
    else
        echo -e "${RED}grc is not installed and cannot be installed${NC}"
        exit 1
    fi
fi

mkdir ~/.grc

# create the config file for grc
echo -e "${GREEN}Creating the config file for grc${NC}"
echo -e """
# Go
\bgo.* test\b
conf.gotest
""" > ~/.grc/grc.conf

# create the goconfig file for grc
echo -e "${GREEN}Creating the goconfig file for grc${NC}"
echo -e """
regexp==== RUN .*
colour=blue
-
regexp=--- PASS: .*
colour=green
-
regexp=^PASS$
colour=green
-
regexp=^(ok|\?) .*
colour=magenta
-
regexp=--- FAIL: .*
colour=red
-
regexp=[^\s]+\.go(:\d+)?
colour=cyan
""" > ~/.grc/conf.gotest

# in ~/.bashrc add alias for go test -> grc go test
echo -e "${GREEN}Adding the following line to your bashrc${NC}"
# if line ``alias go='grc go`` exist then do nothing
if grep -q "alias go='grc go'" ~/.bash_profile; then
    echo -e "${YELLOW}alias go='grc go' already exist${NC}"
else
    echo -e "${GREEN}alias go='grc go' added${NC}"
    echo "alias go='grc go'" >> ~/.bash_profile
fi
source ~/.bash_profile