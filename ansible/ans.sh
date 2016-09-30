#!/bin/bash

key=$(vagrant ssh-config | grep -w IdentityFile | grep -o "/home.*" | head -n1)
ANSHOST=ans.hosts
echo "[hosts]" > $ANSHOST
echo "192.168.33.10" >> $ANSHOST
echo "192.168.33.11" >> $ANSHOST
echo "key=$key"
ansible hosts -u vagrant --private-key=$key -i $ANSHOST -m ping
ansible hosts -u vagrant --private-key=$key -i $ANSHOST -m shell -a "ls -l /"
