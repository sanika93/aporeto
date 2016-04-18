#!/usr/bin/python

import sys
import re

myList=[]

for i in sys.argv:
	myList.append(i)

fileToRead=myList[1]
fileToWrite=myList[2]

filein = open(fileToRead,"r")
fileOut = open(fileToWrite,"w")
data = filein.read()
lines = re.split(r'(?<!\w\.\w.)(?<![A-Z][a-z]\.)(?<=\.|\?)\s',data)
lines_seen = set()
for line in lines:
	if line not in lines_seen:
		fileOut.write(line)
		lines_seen.add(line)
fileOut.close()
