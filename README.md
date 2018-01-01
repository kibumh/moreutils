# moreutils
[moreutils](https://joeyh.name/code/moreutils/) implemented in go.


# Getting Started

```
go get github.com/kibumh/moreutils/...
```


# sponge
sponge sponges os.Stdin and writes back to os.Stdout

## Example
```
sed -e 'SOME_SCRIPT' data.txt > data.txt
```
will end up with an empty file.
```
sed -e 'SOME_SCRIPT' data.txt | sponge data.txt
```
will solve this problem.


# ts
Timestamp prefixer.

## Example

```
$ ping 8.8.8.8
PING 8.8.8.8 (8.8.8.8) 56(84) bytes of data.
64 bytes from 8.8.8.8: icmp_seq=1 ttl=56 time=34.7 ms
64 bytes from 8.8.8.8: icmp_seq=2 ttl=56 time=33.6 ms
^C
$ ping 8.8.8.8 | ~/work/go/bin/ts
1-2 12:04:10 64 bytes from 8.8.8.8: icmp_seq=1 ttl=56 time=36.5 ms
1-2 12:04:11 64 bytes from 8.8.8.8: icmp_seq=2 ttl=56 time=35.1 ms
^C
$ 
```

## TODO
- [ ] add a format argument.
