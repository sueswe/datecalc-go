#!/bin/bash

source ~/.profile

cd "$HOME"/compile/datecalc-go || {
    echo "Status: $?"
    exit 4
}

echo "------------------------------------"
env | grep PATH
env | grep LOADED
echo "------------------------------------"



echo '
### LINUX #########################################################
'

stages="
stp,testta3
lgkk,testta3
"

for UMG in ${stages}
do
    cd /tmp/ || exit 1
    "$HOME"/bin/vicecersa.sh "${UMG}" datecalc \$HOME/bin/ || {
        echo "Status: $?"
        exit 2
    }
done


echo '
### AIX #########################################################
'

stages="
stp,testta2
lgkk,testta2
"

for UMG in ${stages}
do
    cd /tmp/ || exit 1
    "$HOME"/bin/vicecersa.sh "${UMG}" datecalc \$HOME/bin/ datecalc || {
        echo "Status: $?"
        exit 2
    }
done
