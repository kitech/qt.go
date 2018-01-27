/****************************************************************************
** Meta object code from reading C++ file 'qdynslotobject.h'
**
** Created by: The Qt Meta Object Compiler version 67 (Qt 5.10.0)
**
** WARNING! All changes made in this file will be lost!
*****************************************************************************/

#include "qdynslotobject.h"
#include <QtCore/qbytearray.h>
#include <QtCore/qmetatype.h>
#if !defined(Q_MOC_OUTPUT_REVISION)
#error "The header file 'qdynslotobject.h' doesn't include <QObject>."
#elif Q_MOC_OUTPUT_REVISION != 67
#error "This file was generated using the moc from 5.10.0. It"
#error "cannot be used with the include files from this version of Qt."
#error "(The moc has changed too much.)"
#endif

QT_BEGIN_MOC_NAMESPACE
QT_WARNING_PUSH
QT_WARNING_DISABLE_DEPRECATED
struct qt_meta_stringdata_QDynSlotObject_t_tmpl {
    QByteArrayData data[3];
    char stringdata0[128];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_QDynSlotObject_t_tmpl, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static qt_meta_stringdata_QDynSlotObject_t_tmpl qt_meta_stringdata_QDynSlotObject = {
    {
QT_MOC_LITERAL(0, 0, 14), // "QDynSlotObject"
QT_MOC_LITERAL(1, 15, 68), // "slot96BE12EB9AF2F6851704412FA..."
QT_MOC_LITERAL(2, 84, 0) // ""

    },
    "QDynSlotObject\0"
    "slot96BE12EB9AF2F6851704412FA2981E03E32BBD18D40F6040D01F107A20CACC07\0"
    ""
};
// #undef QT_MOC_LITERAL

static uint qt_meta_data_QDynSlotObject[] = {

 // content:
       7,       // revision
       0,       // classname
       0,    0, // classinfo
       0,    0, // methods
       0,    0, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       0,       // signalCount

       0        // eod
};

#include <QtCore>

void QDynSlotObject::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    Q_UNUSED(_o);
    Q_UNUSED(_id);
    Q_UNUSED(_c);
    Q_UNUSED(_a);

    QDynSlotObject* this_ = (QDynSlotObject*)_o;
    qDebug()<<"calling:"<< _o << _c << _id << this_->fnptr_;
    if (this_->fnptr_ != 0) {
        this_->fnptr_(_o, _c, _id, _a, this_->name_,
                      this_->argc_, this_->argtys_, this_->cbptr_);
    }
    qDebug()<<"called:"<< _o << _c << _id;
}

/*
QMetaObject QDynSlotObject::staticMetaObject = {
    { &QObject::staticMetaObject, qt_meta_stringdata_QDynSlotObject.data,
      qt_meta_data_QDynSlotObject,  qt_static_metacall, nullptr, nullptr}
};
*/


const QMetaObject *QDynSlotObject::metaObject() const
{
    qDebug()<<"hehehhe"<< QObject::d_ptr->metaObject;
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *QDynSlotObject::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, this->qt_meta_stringdata.stringdata0))
        return static_cast<void*>(this);
    return QObject::qt_metacast(_clname);
}

int QDynSlotObject::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
{
    qDebug()<<"called:"<<_c << _id;
    _id = QObject::qt_metacall(_c, _id, _a);
    if (_id < 0)
        return _id;
    if (_c == QMetaObject::InvokeMetaMethod) {
        if (_id < 1)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 1;
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        if (_id < 1)
            *reinterpret_cast<int*>(_a[0]) = -1;
        _id -= 1;
    }
    return _id;
}
QT_WARNING_POP
QT_END_MOC_NAMESPACE


/////////////////////
static QAtomicInt dyn_slot_obj_seq = 0;
QDynSlotObject::QDynSlotObject(void *fnptr, char* name, int argc, int*argtys, void*cbptr) : QObject()
{
    qt_meta_stringdata_QDynSlotObject_t *sd = (qt_meta_stringdata_QDynSlotObject_t*)(void*)(&qt_meta_stringdata_QDynSlotObject);
    printf("%d\n", sd->stringdata0[9]);
    sd->stringdata0[9] = 89;
    printf("%d\n", sd->stringdata0[9]);
    // qt_meta_stringdata_QDynSlotObject.stringdata0[9]=0;

    // this->setCallbackSlot(0, "aaaaaaaaa", 0, 0);
    this->setCallbackSlot(fnptr, name, argc, argtys, cbptr);
}

QDynSlotObject::~QDynSlotObject() {}

void QDynSlotObject::setCallbackSlot(void *fnptr, char* name, int argc, int*argtys, void*cbptr)
{
    fnptr_ = (void(*)(QObject*, int, int, void*, char*, int, int*, void*))fnptr;
    name_ = name;
    argc_ = argc;
    argtys_ = argtys;
    cbptr_ = cbptr;

    // rewrite qt_meta_stringdata_ and qt_meta_data_ and staticMetaObject

    qt_meta_stringdata_QDynSlotObject_t sd;
    uint ud[] = {};

    int clsnamesize = 32; qptrdiff clsnameoffset = 0;
    int slotnamesize = 0; qptrdiff slotnameoffset = 33;
    int endelemsize = 0; qptrdiff endelemoffset = 98;
    // memset(this->qt_meta_stringdata.stringdata0, 'A', 128);
    // this->qt_meta_stringdata.data[0].size = clsnamesize; // change example
    // this->qt_meta_stringdata.data[0].offset = clsnameoffset;
    // this->qt_meta_stringdata.data[1].offset = slotnameoffset;
    // this->qt_meta_stringdata.data[2].size = endelemsize;
    // this->qt_meta_stringdata.data[2].offset = endelemoffset;

    uint64_t rdseq = qrand(); rdseq |= uint64_t(qrand()) << 32;
    QByteArray rdseqbin((char*)&rdseq, 8);
    char* rdseqhex = rdseqbin.toHex().toUpper().data();
    qDebug()<<rdseqbin.toHex().toUpper();
    strcpy(&this->qt_meta_stringdata.stringdata0[14], rdseqhex);

    slotnamesize = strlen(name);
    // this->qt_meta_stringdata.data[1].size = slotnamesize;
    // strcpy(&this->qt_meta_stringdata.stringdata0[slotnameoffset], name);
    // strcpy(&this->qt_meta_stringdata.stringdata0[0], "QDynSlotObject999\0");
    qDebug()<<(void*)(&this->qt_meta_stringdata.stringdata0[0]);

    /*
    this->qt_meta_data[0] = 7;
    this->qt_meta_data[4] = 0;  // method count
    this->qt_meta_data[5] = 14; // method offset
    this->qt_meta_data[14] = 1;
    this->qt_meta_data[16] = 19;
    this->qt_meta_data[17] = 2;
    this->qt_meta_data[18] = 0x0a;
    this->qt_meta_data[19] = QMetaType::Void;
    */

    this->staticMetaObject = {
         { &QObject::staticMetaObject, this->qt_meta_stringdata.data,
           this->qt_meta_data, qt_static_metacall, nullptr, nullptr},
    };
}

QDemoObject::QDemoObject(){}
QDemoObject::~QDemoObject(){}
void QDemoObject::dummyslot(){}

extern "C" void callback_slot(void*_o, int _c, int _id, void*_a,
                              char*name, int argc, int*argtys, void*cbptr) {
    qDebug()<<"callback_slot called:"<<_o<<_c<<_id<<_a
            << ((QDynSlotObject*)_o)->name_;
    void**ap = (void**)_a;
    const QString* newname = (const QString*)(ap[1]);
    qDebug()<<"new object name from signal parameter 0:" << newname->toLatin1().data() << ap[1];
}

int main(int argc, char **argv)
{
    QCoreApplication app(argc, argv);
    qsrand(time(NULL));

    qDebug() << "Basic QObject class size: " << sizeof(QDemoObject);
    qDebug() << "Deep Custom QObject class size: " << sizeof(QDynSlotObject);
    qDebug() << "QMetaObject class size: " << sizeof(QMetaObject);
    qDebug() << "meta stringdata class size: " << sizeof(struct QDynSlotObject::qt_meta_stringdata_QDynSlotObject_t);

    QDemoObject* dso1 = new QDemoObject();
    QMetaObject* mto1 = (QMetaObject*)dso1->metaObject();
    qDebug() << (void*)mto1->className()  << mto1->methodCount();
    qDebug() << mto1->className()  << mto1->methodCount();

    QDynSlotObject* dso = new QDynSlotObject((void*)&callback_slot, "clicked123", 0, 0, 0);
    QMetaObject* mto = (QMetaObject*)dso->metaObject();
    qDebug() << (void*)(mto->className())  << mto->methodCount();
    qDebug() << mto->className()  << mto->methodCount();
    for (int i = 0; i < mto->methodCount(); i ++) {
        qDebug()<<mto->method(i).name()<<mto->method(i).methodSignature();
    }

    QDynSlotObject* dso2 = new QDynSlotObject((void*)&callback_slot, "released456", 0, 0, 0);
    QObject::connect(dso, SIGNAL(objectNameChanged(const QString &)), dso2, SLOT(dumnyslot()));
    dso->setObjectName(QString("beginnotdso111"));
    dso->setObjectName(QString("beginnotdso222"));

    // app.exec();
    return 0;
}
