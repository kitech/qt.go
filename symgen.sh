#!/bin/sh

modules="Core Gui Widgets Network Inline"
genfile=qtsymbols/qtsymbols.go



echo -e "package qtsymbols\n" > $genfile
echo -e "var Qtsymbols = map[string]string{\n" >> $genfile

echo > /tmp/sym.txt
for mod in $modules; do
    libpath="/usr/lib/libQt5$mod.so"
    echo $libpath
    symbols=$(readelf -Ws $libpath|awk '{print $8}'|awk -F@ '{print $1}')
    for sym in $symbols; do
        echo "$sym" >> /tmp/sym.txt
    done
done

symtep2=$(cat /tmp/sym.txt|sort -u|uniq -u)

for sym in $symtep2; do
    sign=$(echo $sym|c++filt)
    echo "$sym, $sign"
    # if [ x"$sym" == x"qt_version_tag" ]; then
    #    continue
    # fi
    echo "\"$sym\": \"$sign\"," >> $genfile
done

echo -e "}" >> $genfile
gofmt -w $genfile

### clean
rm -fv /tmp/sym.txt

