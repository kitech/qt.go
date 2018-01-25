// 只是实例代码，用来生成moc_myClass  
#include"myClass.h"  
#include<iostream>  
using std::cout;  
using std::endl;  
myClass::myClass()  
{  
    connect(this, SIGNAL(signalFunc),  
            this, SLOT(slotFunc));  
    connect(this, SIGNAL(signalFunc2),  
            this, SLOT(slotFunc2));  
}  
myClass::~myClass()  
{  
}  
void myClass::slotFunc(double d)  
{  
    cout << "slotFunc" << endl;  
}  
int myClass::slotFunc2(char c)  
{  
    cout << "slotFunc2" << endl;  
    return c;  
}  
int myClass::slotFunc3(int c)  
{  
    cout << "slotFunc3" << endl;  
    return c;  
}  
void myClass::trigger()  
{  
}  
void myClass::trigger2()  
{
}

