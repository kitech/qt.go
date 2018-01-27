#ifndef _QDynSlotObject_H_
#define _QDynSlotObject_H_

#include <QObject>

class QDemoObject : public QObject {
    Q_OBJECT;
 public:
    QDemoObject();
    virtual ~QDemoObject();

 public slots:
     void dummyslot();
};


class QDynSlotObject : public QObject
{  
    // Q_OBJECT
 public:
    QMetaObject staticMetaObject;
    virtual const QMetaObject *metaObject() const;
    virtual void *qt_metacast(const char *);
    virtual int qt_metacall(QMetaObject::Call, int, void **);
    static inline QString tr(const char *s, const char *c = nullptr, int n = -1)
    { return QObject::staticMetaObject.tr(s, c, n); }
    static inline QString trUtf8(const char *s, const char *c = nullptr, int n = -1)
    { return QObject::staticMetaObject.tr(s, c, n); }
 private:
    __attribute__((visibility("hidden"))) static void qt_static_metacall(QObject *, QMetaObject::Call, int, void **);
    struct QPrivateSignal {};

 public:
    QDynSlotObject(void *fnptr, char* name, int argc, int*argtys, void* cbptr);
    ~QDynSlotObject();
    void setCallbackSlot(void *fnptr, char* name, int argc, int*argtys, void* cbptr);

    void (*fnptr_)(QObject*, int, int, void*, char*name, int argc, int*argtys, void*cbptr) = 0;
    char* name_;
    int argc_;
    int* argtys_;
    void*cbptr_;

 public:
    ///
    struct qt_meta_stringdata_QDynSlotObject_t {
        QByteArrayData data[3];
        char stringdata0[128];  // classname123\0firstslotname\0\0
    };

#define _QT_MOC_LITERAL_TMP(idx, ofs, len)                                   \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len,        \
                                                            qptrdiff(offsetof(qt_meta_stringdata_QDynSlotObject_t, stringdata0) + ofs \
                                                                     - idx * sizeof(QByteArrayData)) \
                                                            )

    struct qt_meta_stringdata_QDynSlotObject_t qt_meta_stringdata = {
        {
            _QT_MOC_LITERAL_TMP(0, 0, 30), // "QDemoObject"
            _QT_MOC_LITERAL_TMP(1, 31, 9), // "dummyslot"
            _QT_MOC_LITERAL_TMP(2, 41, 0) // ""
            // { Q_REFCOUNT_INITIALIZE_STATIC, 0, 11, 0, 0 }, // 还是这个初始化方式不对
            // { Q_REFCOUNT_INITIALIZE_STATIC, 0, 9, 0, 12 },
            // { Q_REFCOUNT_INITIALIZE_STATIC, 0, 0, 0, 22 }
        },
        "QDynSlotObject0123456789ABCDEF\0dumnyslot\0"
        // uint64的object编号
        /*
            {
                { Q_REFCOUNT_INITIALIZE_STATIC, 0, 32, 0, 0 },
                { Q_REFCOUNT_INITIALIZE_STATIC, 0, 64, 0, 33 },
                { Q_REFCOUNT_INITIALIZE_STATIC, 0, 0, 0, 98 },
            },
            {},
        */
        };
#undef _QT_MOC_LITERAL_TMP
    uint qt_meta_data[128] = {
        // content:
        7,       // revision
        0,       // classname
        0,    0, // classinfo
        1,   14, // methods
        0,    0, // properties
        0,    0, // enums/sets
        0,    0, // constructors
        0,       // flags
        0,       // signalCount

        // slots: name, argc, parameters, tag, flags
        1,    0,   19,    2, 0x0a /* Public */,

        // slots: parameters
        QMetaType::Void,

        0        // eod
    };
};

// TODO 放在qt.inline项目中吧，
// 注意，尽量避免链接进进程，尽量使用ffi_call
#ifdef __cplusplus
extern "C"{
#endif
    QDynSlotObject* QDynSlotObject_new();
    void QDynSlotObject_delete(QDynSlotObject*o);
#ifdef __cplusplus
};
#endif


/*
class QDynSlotObject : public QObject
{
    public:
# 8 "qdynslotobject.h"
#pragma GCC diagnostic push
# 8 "qdynslotobject.h"
   
# 8 "qdynslotobject.h"
#pragma GCC diagnostic ignored "-Wsuggest-override"
# 8 "qdynslotobject.h"
    static const QMetaObject staticMetaObject; virtual const QMetaObject *metaObject() const; virtual void *qt_metacast(const char *); virtual int qt_metacall(QMetaObject::Call, int, void **); static inline QString tr(const char *s, const char *c = nullptr, int n = -1) { return staticMetaObject.tr(s, c, n); } static inline QString trUtf8(const char *s, const char *c = nullptr, int n = -1) { return staticMetaObject.tr(s, c, n); } private:
# 8 "qdynslotobject.h"
#pragma GCC diagnostic ignored "-Wattributes"
# 8 "qdynslotobject.h"
    __attribute__((visibility("hidden"))) static void qt_static_metacall(QObject *, QMetaObject::Call, int, void **);
# 8 "qdynslotobject.h"
#pragma GCC diagnostic pop
# 8 "qdynslotobject.h"
    struct QPrivateSignal {};
 public:
    QDynSlotObject();
    ~QDynSlotObject();
};

*/
#endif


