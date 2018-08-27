#include <string.h>

#include "dynamic_meta_object.h"

// publics
DMetaObject::DMetaObject(){}
DMetaObject::~DMetaObject() {
    // mannually emit destroy signal
    void *fnptr = this->staticMetaObject->static_metacall;
    ((void (*)(void*, int, int, void**))(fnptr))(this, 0, 0, 0);
}

const void* DMetaObject::metaObject() const {
    return this->staticMetaObject;
}
void* DMetaObject::metaObject() {
    return this->staticMetaObject;
}

void* DMetaObject::qt_metacast(const char* _clname) {
    return this->qt_metacast_fnptr(_clname);
}

int DMetaObject::qt_metacall(/*QMetaObject::Call*/ int call, int id, void **arguments) {
    return this->qt_metacall_fnptr(call, id, arguments);
}

// internals

// externs
void* DMetaObject_new(void* modat, void* metacast_fnptr, void* metacall_fnptr) {
    DMetaObject* dmo = new DMetaObject();
    dmo->staticMetaObject = (QMetaObjectData*)modat;
    dmo->qt_metacast_fnptr = (void* (*)(const char*))(metacast_fnptr);
    dmo->qt_metacall_fnptr = (int (*)(int, int, void**))(metacall_fnptr);

    return dmo;
}
void DMetaObject_delete(void* this_) {
    delete (DMetaObject*)(void*)(0);
}

// 参考：
// https://github.com/GIPdA/DynamicMetaobject
// https://doc.qt.io/archives/qq/qq16-dynamicqobject.html
//
