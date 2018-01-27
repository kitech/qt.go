#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <tsym.h>

Abc::Abc(int a0, int a1, int a2, int a3) {
    printf("a0=%d, a1=%d, a2=%d\n", a0, a1, a2);
}
Abc::Abc(int& a0, char **a1, int a2) {
    printf("a0=%d, a1=%p, a2=%d\n", a0, a1, a2);
    printf("%s\n", a1[0]);
}

Abc::Abc(char*ch){
    printf("a0=%p, a0=%s\n", ch, ch);
    _s = (char*)ch;
}

int Abc::length() {
    printf("a0=%p, a0=%s\n", _s, _s);
    return strlen(_s);
}
