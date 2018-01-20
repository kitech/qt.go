
class Abc {
 public:
    Abc(int a0, int a1, int a2, int a3);
    Abc(int& a0, char **a1, int a2 = 9);
    Abc(char*ch);
    int length();

    char * _s;
};
