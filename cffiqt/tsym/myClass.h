#ifndef _MYCLASS_H_
#define _MYCLASS_H_

#include <QtCore>

class myClass : public QObject
{  
    Q_OBJECT

 public:  
    myClass();  
    ~myClass();  

    void trigger();
    void trigger2();  

signals:
    long signalFunc(double aa);  
    int signalFunc2(char ab, int ac);  

protected slots:
        short slotFunc(double ad);  
        int slotFunc2(char ae);
        int slotFunc3(int af);

public slots:
    void slotFunc4(double ba);  
            int slotFunc5(char bb);
            int slotFunc6(int bc);
};

#endif

