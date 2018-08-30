#ifndef _DYNAMIC_META_OBJECT_H_
#define _DYNAMIC_META_OBJECT_H_

// #include <QObject>

#ifdef __cplusplus
extern "C" {
#endif
    void* DMetaObject_new(void* modat, void* metacast_fnptr, void* metacall_fnptr);
    void DMetaObject_delete(void* this_);

typedef struct QMetaObjectData { // private data
    /*const QMetaObject*/void *superdata;
    /*const QByteArrayData*/void *stringdata;
    /*const uint*/void *data;
    // typedef void (*StaticMetacallFunction)(QObject *, QMetaObject::Call, int, void **);
    /*StaticMetacallFunction*/void* static_metacall;
    /*const QMetaObject*/void * const *relatedMetaObjects;
    void *extradata; //reserved for future use
} QMetaObjectData;

#ifdef __cplusplus
};
#endif

#endif

