#include <string.h>
#include <stdio.h>

#include "dynamic_meta_object.h"

#include <QObject>

///// Dynamic Emulate object
class DemuObject : public QObject {
public: // Q_OBJECT extends rewrite
    QMetaObject staticMetaObject;
    virtual const QMetaObject *metaObject() const;
    virtual void *qt_metacast(const char *);
    virtual int qt_metacall(QMetaObject::Call, int, void **);
    static inline QString tr(const char *s, const char *c = nullptr, int n = -1)
    { return QObject::staticMetaObject.tr(s, c, n); }
    static inline QString trUtf8(const char *s, const char *c = nullptr, int n = -1)
    { return QObject::staticMetaObject.tr(s, c, n); }
private:
public:
    Q_DECL_HIDDEN_STATIC_METACALL static void qt_static_metacall(QObject *, QMetaObject::Call, int, void **);
    struct QPrivateSignal {};

    virtual QMetaObject *metaObject() ;
public:
    explicit DemuObject();
    virtual ~DemuObject();
    void* (*qt_metacast_fnptr)(const char*);
    int (*qt_metacall_fnptr)(int, int, void**);

};


// publics
DemuObject::DemuObject () : QObject() {}
DemuObject::~DemuObject() {
    // mannually emit destroy signal
    // auto fnptr = this->staticMetaObject.d.static_metacall;
    // (fnptr)(this, 0, 0, 0)
}

#define ALOG() printf("in %s:%d %s\n", __FILE__, __LINE__, __FUNCTION__);
const QMetaObject* DemuObject::metaObject() const {
    ALOG();
    return &this->staticMetaObject;
}
QMetaObject* DemuObject::metaObject() {
    ALOG();
    return &this->staticMetaObject;
}

void* DemuObject::qt_metacast(const char* _clname) {
    ALOG();
    return this->qt_metacast_fnptr(_clname);
}

int DemuObject::qt_metacall(QMetaObject::Call call, int id, void **arguments) {
    ALOG();
    return this->qt_metacall_fnptr(call, id, arguments);
}

void DemuObject::qt_static_metacall(QObject *, QMetaObject::Call , int, void **) {
    ALOG();
}

// internals

#include <QDebug>
// externs
void* DMetaObject_new(void* modat, void* metacast_fnptr, void* metacall_fnptr) {
    DemuObject* dmo = new DemuObject();
    dmo->staticMetaObject.d.superdata = ((QMetaObject*)modat)->d.superdata;
    dmo->staticMetaObject.d.stringdata = ((QMetaObject*)modat)->d.stringdata;
    dmo->staticMetaObject.d.data = ((QMetaObject*)modat)->d.data;
    dmo->staticMetaObject.d.static_metacall = DemuObject::qt_static_metacall;
    dmo->qt_metacast_fnptr = (void* (*)(const char*))(metacast_fnptr);
    dmo->qt_metacall_fnptr = (int (*)(int, int, void**))(metacall_fnptr);
    qDebug()<<dmo;
    qDebug()<<dmo->staticMetaObject.d.superdata;
    qDebug()<<dmo->staticMetaObject.d.stringdata;
    qDebug()<<dmo->staticMetaObject.d.data;
    qDebug()<<dmo->staticMetaObject.d.static_metacall;

    return dmo;
}
void DMetaObject_delete(void* this_) {
    delete (DemuObject*)(void*)(0);
}

// 参考：
// https://github.com/GIPdA/DynamicMetaobject
// https://doc.qt.io/archives/qq/qq16-dynamicqobject.html
//
