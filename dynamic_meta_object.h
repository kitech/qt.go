#ifndef _DYNAMIC_META_OBJECT_H_
#define _DYNAMIC_META_OBJECT_H_

// #include <QObject>

class QMetaObjectData { // private data
 public:
    /*const QMetaObject*/void *superdata;
    /*const QByteArrayData*/void *stringdata;
    /*const uint*/void *data;
    // typedef void (*StaticMetacallFunction)(QObject *, QMetaObject::Call, int, void **);
    /*StaticMetacallFunction*/void* static_metacall;
    /*const QMetaObject*/void * const *relatedMetaObjects;
    void *extradata; //reserved for future use
};

class DMetaObject {
 public: // Q_OBJECT extends
    // static const QMetaObject staticMetaObject;
    // virtual const QMetaObject *metaObject() const;
    // virtual void *qt_metacast(const char*);
    // virtual int qt_metacall(QMetaObject::Call, int, void **);
    // static inline QString tr(const char *s, const char *c = nullptr, int n = -1)
    //     { return staticMetaObject.tr(s, c, n); }
    // static inline QString trUtf8(const char *s, const char *c = nullptr, int n = -1)
    //     { return staticMetaObject.tr(s, c, n); }

 public:
    QMetaObjectData* staticMetaObject;

 public: // Q_OBJECT extends rewrite

    virtual const void* metaObject() const;
    virtual void* qt_metacast(const char*);
    virtual int qt_metacall(/*QMetaObject::Call*/ int call, int id, void **arguments);

    virtual void* metaObject();

 public:
    explicit DMetaObject();
    virtual ~DMetaObject();
    void* (*qt_metacast_fnptr)(const char*);
    int (*qt_metacall_fnptr)(int, int, void**);

};

#ifdef __cplusplus
extern "C" {
#endif
    void* DMetaObject_new(void* modat, void* metacast_fnptr, void* metacall_fnptr);
    void DMetaObject_delete(void* this_);
#ifdef __cplusplus
};
#endif

#endif

