# d20
d20 is a trivial CLI application that outputs 20 (by default) strings of 20 (by default) 
characters for use as passwords/tokens/etc. 

## Syntax

By default, twenty strings of twenty characters, containing any alphanumeric or most "safe special" characters are output, ala:

```
oWx!s5^YJ7H;o8xr>Gg6
yvmaZA2dGzfYoil&.U(l
X-%/>HOEX1:U1Yz1qBhD
SHf0lguOqGi0Fb(bIIM%
m1WqtH3ICjp9uW;sRbID
w25HhL%a>81_$o42H3T/
nkejc2-1;9T<fDQd60U2
^U:tGK5JqS)I8>EJ5>g3
/l3jXBm9f5q:jffXH$4F
22GMAt2u-6%DW/qN/l32
onD7(1Kc_fk%e$eCFQ<9
2sF9Dx;I(3Kr&?WFee,M
.FyeIK&<?^Ngixjsn<<S
hJS^54W?PVor3A)KATt1
cqKB<^;m546e3EQmdj)P
1.TPnUE$Td2II>?uFXG6
C3T/O,V:Nd,&okmV)Zej
jtc^Nc49EPw;%9(Hwv)B
ggET5DKn84S8Ht4V2S9r
-x&!sfii0$wWAH!Z5B5K
```

Can also be used to generate keyblocks:

```bash
d20 --count 1 --length 741 --keyblock
PdbcbO8kBT1reE0OF3Ku4xrM1X05BNLqUpZ4F5vJ1Db/JyrYBHKsbxL96irdjtGPP
DzFtrsN2DCw53i9NXCmCghg0w2Z42zn6gBZxlyNmO86n3V1eIqAV8srR63O8lfK6P
VkfKAuCXkWepLzbf9a/8v9hknNBET2yNK2h4OvW2i5P1JL7cYy8nA1MRKCOJUI8g0
nNaAUFx3dUgmyxP+yJTphIUh4STiFbcbaBCY7xlDjty/lzZF3VUVHyD7Ej5djV7ne
Cb1AHk2OVO7N6zw2o2aMR9CCq1Re6cMongodK/pmMsjur5eUdpQbKyYHkUnFUxUu7
4LKN7E4vhV6YAn9kCt6SIPmIS78KjNmzUiwqrMvublJFAiCcdxwEuEaG9+Id5Lil/
M/Na4pNUgnogC2AZXsQSVSG5s2h4cu4x5+B7CxoeplF1pHNcyb3TBWw4mHfiHpv7V
RijLh8VRkcjgqZ3cv090VpoLcXnhG1rzE5XYEi0oedj2y7Xfr6ymRJ3c00OW67VHR
33bKf0kq0gYc0P1Ys3Q3zyswyTE3YwYeGZcxeblnXVis8SZnTnacEPV1XEL1AR/NK
LP7Pwx1McR+1gdg50yB6ey+qRWCg+rXNztsIdwxvPdzrTYrR8dtq1l1206FIcIgPD
IwmEAWufoGXTletnDuN15jxX9/vVNBfxxWwWQzio2mDUBf0dJ/qukWxz1a9IU78IP
vSHdm1snzJ7H6DXCa48zj+C+BR4MbIcokvi4KbKnrYA1NxrYde3FbysQrA4100T9q
4yHlriE99gnf4ts0WfwORQiFGlM3SItPjVncI9iZ6BW7S6CzGai5AjeJ/vBpovCHT
uTQjfuE+bmJUu7LEsT1A1651cSEpIAIloveChi8LRvkqIbLD30Ia5dI6PByG/1gM+
6VKxC/ITZ3LvnvKQp8FKYSWm7FBGPGAXLsqOVMs9gX
```

All of those options are configurable.

```bash
Usage of ./d20:
  -base64
    	Base64 encode the output
  -block
    	Block the output to 65 character lines
  -blocksize int
    	Slight misnomer: if --block is used, sets the line length to int (default 65)
  -chars string
    	Characters to use (all bytes alphanumeric alphanumeric-nosim numeric alphabet binary hexadecimal) (default "all")
  -count int
    	Number of strings (default 20)
  -custom string
    	A list of characters you want to use in lieu of '--chars' (repeat for prevalence)
  -keyblock
    	Shortcut to '--char bytes --base64 --block --blocksize 65' (HINT: --length 741, perhaps?)
  -length int
    	Length of string (default 20)
  -mangle string
    	Mangle the output (WARN: Decreases cardinality, should not be used with --base64) (UC LC)
  -pin int
    	Shortcut to '--char numeric --length int'
  -separator string
    	What character or string should each value be separated with? (default "\n")
  -unique
    	Ensure generated strings are unique
```
