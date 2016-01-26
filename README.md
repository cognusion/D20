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

All of those options are configurable.

```bash
 d20 --help
Usage of d20:
  -chars string
    	Characters to use (all alphanumeric alphanumeric-nosim numeric alphabet binary hexadecimal) (default "all")
  -count int
    	Number of strings (default 20)
  -length int
    	Length of string (default 20)
  -mangle string
    	Mangle the output (Decreases cardinality) (UC LC)
```
