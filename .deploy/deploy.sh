#!/bin/bash

source ~/.profile

name='datecalc'

cd "$HOME"/compile/${name}-go || {
    echo "Status: $?"
    exit 4
}

echo "------------------------------------"
env | grep PATH
env | grep LOADED
echo "------------------------------------"

masterrtc=0

echo '### LINUX #########################################################'
stages="stp,testta3 lgkk,testta3 stp,prodta3 lgkk,prodta3 hema,test,01T hema,prod"
for UMG in ${stages}
do
    cd /tmp/ || exit 1
    remotecommander.rb -d "\$HOME/bin/" -s "${name}" -g ${UMG} || {
        echo "Status: $?"
        masterrtc=2
    }
done
rm -v /tmp/${name}


echo '### AIX #########################################################'
stages="stp,testta2 stp,prodta2 pfif"
cd /tmp/ || exit 1
cp ${name}.aix ${name}
for UMG in ${stages}
do
    remotecommander.rb -d "\$HOME/bin/" -s "${name}" -g ${UMG}  || {
        echo "Status: $?"
        masterrtc=2
    }
done

rm -v /tmp/${name}*
exit ${masterrtc}
