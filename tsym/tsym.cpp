#include <stdio.h>
#include <tsym.h>

Abc::Abc(int a0, int a1, int a2, int a3) {
    printf("a0=%d, a1=%d, a2=%d\n", a0, a1, a2);
}
Abc::Abc(int& a0, char **a1, int a2) {
    printf("a0=%d, a1=%p, a2=%d\n", a0, a1, a2);
    printf("%s\n", a1[0]);
}

