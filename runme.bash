#!/bin/bash

# cleaning old files
rm mem.log
rm reads.txt
rm writes.txt
rm bat.html

getWriteSpeed () {
    dd if=/dev/zero bs=$block_size of=tstfile count=1024 > mem.log 2>&1 
    result=$(sed -n '3p' mem.log)
    echo $result | grep -Eo '[0-9.]+' > mem.log
    result=$(sed -n '3p' mem.log)
}

getReadSpeed () {
    dd if=tstfile bs=$block_size of=/dev/null count=1024 > mem.log 2>&1 
    result=$(sed -n '3p' mem.log)
    echo $result | grep -Eo '[0-9.]+' > mem.log
    result=$(sed -n '3p' mem.log)
}

for i in {4,8,16,32,64,128,192,256,512}
do
    power=$(($i))
    k="k"
    block_size="$power$k"
    getWriteSpeed
    echo $power,$result >> writes.txt
    getReadSpeed
    echo $power,$result >> reads.txt
done

rm mem.log
rm tstfile

go run plot.go
