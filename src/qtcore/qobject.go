//  header block begin
// /usr/include/qt/QtCore/qobject.h
// #include <qobject.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 10
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"

func init() {
	if false {
		reflect.TypeOf(123)
	}
	if false {
		reflect.TypeOf(unsafe.Sizeof(0))
	}
	if false {
		fmt.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
}

//  ext block end

//  body block begin
type QObject struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qobject.h:118
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QObject) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QObject10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:123
// index:0
// void QObject(class QObject *)
func NewQObject(parent unsafe.Pointer) *QObject {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QObjectC2EPS_", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QObject{cthis}
}

// /usr/include/qt/QtCore/qobject.h:124
// index:0
// virtual
// void ~QObject()
func DeleteQObject(*QObject) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QObjectD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:126
// index:0
// virtual
// bool event(class QEvent *)
func (this *QObject) Event(event unsafe.Pointer) {
	// 0: (, QEvent * event), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QObject5eventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.cthis, event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:127
// index:0
// virtual
// bool eventFilter(class QObject *, class QEvent *)
func (this *QObject) EventFilter(watched unsafe.Pointer, event unsafe.Pointer) {
	// 0: (, QObject * watched, QEvent * event), (watched, event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QObject11eventFilterEPS_P6QEvent", ffiqt.FFI_TYPE_VOID, this.cthis, watched, event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:144
// index:0
// QString objectName()
func (this *QObject) ObjectName() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QObject10objectNameEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:145
// index:0
// void setObjectName(const class QString &)
func (this *QObject) SetObjectName(name unsafe.Pointer) {
	// 0: (, const QString & name), (name)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QObject13setObjectNameERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, name)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:147
// index:0
// inline
// bool isWidgetType()
func (this *QObject) IsWidgetType() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QObject12isWidgetTypeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:148
// index:0
// inline
// bool isWindowType()
func (this *QObject) IsWindowType() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QObject12isWindowTypeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:150
// index:0
// inline
// bool signalsBlocked()
func (this *QObject) SignalsBlocked() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QObject14signalsBlockedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:151
// index:0
// bool blockSignals(_Bool)
func (this *QObject) BlockSignals(b bool) {
	// 0: (, bool b), (&b)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QObject12blockSignalsEb", ffiqt.FFI_TYPE_VOID, this.cthis, &b)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:153
// index:0
// QThread * thread()
func (this *QObject) Thread() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QObject6threadEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:154
// index:0
// void moveToThread(class QThread *)
func (this *QObject) MoveToThread(thread unsafe.Pointer) {
	// 0: (, QThread * thread), (thread)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QObject12moveToThreadEP7QThread", ffiqt.FFI_TYPE_VOID, this.cthis, thread)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:156
// index:0
// int startTimer(int, Qt::TimerType)
func (this *QObject) StartTimer(interval int, timerType int) {
	// 0: (, int interval, Qt::TimerType timerType), (&interval, &timerType)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QObject10startTimerEiN2Qt9TimerTypeE", ffiqt.FFI_TYPE_VOID, this.cthis, &interval, &timerType)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:164
// index:0
// void killTimer(int)
func (this *QObject) KillTimer(id int) {
	// 0: (, int id), (&id)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QObject9killTimerEi", ffiqt.FFI_TYPE_VOID, this.cthis, &id)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:207
// index:0
// inline
// const QObjectList & children()
func (this *QObject) Children() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QObject8childrenEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:209
// index:0
// void setParent(class QObject *)
func (this *QObject) SetParent(parent unsafe.Pointer) {
	// 0: (, QObject * parent), (parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QObject9setParentEPS_", ffiqt.FFI_TYPE_VOID, this.cthis, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:210
// index:0
// void installEventFilter(class QObject *)
func (this *QObject) InstallEventFilter(filterObj unsafe.Pointer) {
	// 0: (, QObject * filterObj), (filterObj)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QObject18installEventFilterEPS_", ffiqt.FFI_TYPE_VOID, this.cthis, filterObj)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:211
// index:0
// void removeEventFilter(class QObject *)
func (this *QObject) RemoveEventFilter(obj unsafe.Pointer) {
	// 0: (, QObject * obj), (obj)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QObject17removeEventFilterEPS_", ffiqt.FFI_TYPE_VOID, this.cthis, obj)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:213
// index:0
// static
// QMetaObject::Connection connect(const class QObject *, const char *, const class QObject *, const char *, Qt::ConnectionType)
func (this *QObject) Connect(sender unsafe.Pointer, signal unsafe.Pointer, receiver unsafe.Pointer, member unsafe.Pointer, arg4 int) {
	// 0: (const QObject * sender, const char * signal, const QObject * receiver, const char * member, Qt::ConnectionType arg4), (sender, signal, receiver, member, arg4)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QObject7connectEPKS_PKcS1_S3_N2Qt14ConnectionTypeE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QObject_Connect(sender unsafe.Pointer, signal unsafe.Pointer, receiver unsafe.Pointer, member unsafe.Pointer, arg4 int) {
	// 0: (const QObject * sender, const char * signal, const QObject * receiver, const char * member, Qt::ConnectionType arg4), (sender, signal, receiver, member, arg4)
	var nilthis *QObject
	nilthis.Connect(sender, signal, receiver, member, arg4)
}

// /usr/include/qt/QtCore/qobject.h:216
// index:1
// static
// QMetaObject::Connection connect(const class QObject *, const class QMetaMethod &, const class QObject *, const class QMetaMethod &, Qt::ConnectionType)
func (this *QObject) Connect_1(sender unsafe.Pointer, signal unsafe.Pointer, receiver unsafe.Pointer, method unsafe.Pointer, type_ int) {
	// 1: (const QObject * sender, const QMetaMethod & signal, const QObject * receiver, const QMetaMethod & method, Qt::ConnectionType type), (sender, signal, receiver, method, type_)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QObject7connectEPKS_RK11QMetaMethodS1_S4_N2Qt14ConnectionTypeE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QObject_Connect_1(sender unsafe.Pointer, signal unsafe.Pointer, receiver unsafe.Pointer, method unsafe.Pointer, type_ int) {
	// 1: (const QObject * sender, const QMetaMethod & signal, const QObject * receiver, const QMetaMethod & method, Qt::ConnectionType type), (sender, signal, receiver, method, type_)
	var nilthis *QObject
	nilthis.Connect_1(sender, signal, receiver, method, type_)
}

// /usr/include/qt/QtCore/qobject.h:220
// index:2
// inline
// QMetaObject::Connection connect(const class QObject *, const char *, const char *, Qt::ConnectionType)
func (this *QObject) Connect_2(sender unsafe.Pointer, signal unsafe.Pointer, member unsafe.Pointer, type_ int) {
	// 2: (, const QObject * sender, const char * signal, const char * member, Qt::ConnectionType type), (sender, signal, member, &type_)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QObject7connectEPKS_PKcS3_N2Qt14ConnectionTypeE", ffiqt.FFI_TYPE_VOID, this.cthis, sender, signal, member, &type_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:342
// index:0
// static
// bool disconnect(const class QObject *, const char *, const class QObject *, const char *)
func (this *QObject) Disconnect(sender unsafe.Pointer, signal unsafe.Pointer, receiver unsafe.Pointer, member unsafe.Pointer) {
	// 0: (const QObject * sender, const char * signal, const QObject * receiver, const char * member), (sender, signal, receiver, member)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QObject10disconnectEPKS_PKcS1_S3_", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QObject_Disconnect(sender unsafe.Pointer, signal unsafe.Pointer, receiver unsafe.Pointer, member unsafe.Pointer) {
	// 0: (const QObject * sender, const char * signal, const QObject * receiver, const char * member), (sender, signal, receiver, member)
	var nilthis *QObject
	nilthis.Disconnect(sender, signal, receiver, member)
}

// /usr/include/qt/QtCore/qobject.h:344
// index:1
// static
// bool disconnect(const class QObject *, const class QMetaMethod &, const class QObject *, const class QMetaMethod &)
func (this *QObject) Disconnect_1(sender unsafe.Pointer, signal unsafe.Pointer, receiver unsafe.Pointer, member unsafe.Pointer) {
	// 1: (const QObject * sender, const QMetaMethod & signal, const QObject * receiver, const QMetaMethod & member), (sender, signal, receiver, member)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QObject10disconnectEPKS_RK11QMetaMethodS1_S4_", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QObject_Disconnect_1(sender unsafe.Pointer, signal unsafe.Pointer, receiver unsafe.Pointer, member unsafe.Pointer) {
	// 1: (const QObject * sender, const QMetaMethod & signal, const QObject * receiver, const QMetaMethod & member), (sender, signal, receiver, member)
	var nilthis *QObject
	nilthis.Disconnect_1(sender, signal, receiver, member)
}

// /usr/include/qt/QtCore/qobject.h:346
// index:2
// inline
// bool disconnect(const char *, const class QObject *, const char *)
func (this *QObject) Disconnect_2(signal unsafe.Pointer, receiver unsafe.Pointer, member unsafe.Pointer) {
	// 2: (, const char * signal, const QObject * receiver, const char * member), (signal, receiver, member)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QObject10disconnectEPKcPKS_S1_", ffiqt.FFI_TYPE_VOID, this.cthis, signal, receiver, member)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:349
// index:3
// inline
// bool disconnect(const class QObject *, const char *)
func (this *QObject) Disconnect_3(receiver unsafe.Pointer, member unsafe.Pointer) {
	// 3: (, const QObject * receiver, const char * member), (receiver, member)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QObject10disconnectEPKS_PKc", ffiqt.FFI_TYPE_VOID, this.cthis, receiver, member)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:351
// index:4
// static
// bool disconnect(const struct QMetaObject::Connection &)
func (this *QObject) Disconnect_4(arg0 int) {
	// 4: (const QMetaObject::Connection & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QObject10disconnectERKN11QMetaObject10ConnectionE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QObject_Disconnect_4(arg0 int) {
	// 4: (const QMetaObject::Connection & arg0), (arg0)
	var nilthis *QObject
	nilthis.Disconnect_4(arg0)
}

// /usr/include/qt/QtCore/qobject.h:390
// index:0
// void dumpObjectTree()
func (this *QObject) DumpObjectTree() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QObject14dumpObjectTreeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:393
// index:1
// void dumpObjectTree()
func (this *QObject) DumpObjectTree_1() {
	// 1: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QObject14dumpObjectTreeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:391
// index:0
// void dumpObjectInfo()
func (this *QObject) DumpObjectInfo() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QObject14dumpObjectInfoEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:394
// index:1
// void dumpObjectInfo()
func (this *QObject) DumpObjectInfo_1() {
	// 1: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QObject14dumpObjectInfoEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:397
// index:0
// bool setProperty(const char *, const class QVariant &)
func (this *QObject) SetProperty(name unsafe.Pointer, value unsafe.Pointer) {
	// 0: (, const char * name, const QVariant & value), (name, value)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QObject11setPropertyEPKcRK8QVariant", ffiqt.FFI_TYPE_VOID, this.cthis, name, value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:398
// index:0
// QVariant property(const char *)
func (this *QObject) Property(name unsafe.Pointer) {
	// 0: (, const char * name), (name)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QObject8propertyEPKc", ffiqt.FFI_TYPE_VOID, this.cthis, name)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:399
// index:0
// QList<QByteArray> dynamicPropertyNames()
func (this *QObject) DynamicPropertyNames() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QObject20dynamicPropertyNamesEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:403
// index:0
// static
// uint registerUserData()
func (this *QObject) RegisterUserData() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QObject16registerUserDataEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QObject_RegisterUserData() {
	// 0: (), ()
	var nilthis *QObject
	nilthis.RegisterUserData()
}

// /usr/include/qt/QtCore/qobject.h:404
// index:0
// void setUserData(uint, class QObjectUserData *)
func (this *QObject) SetUserData(id uint, data unsafe.Pointer) {
	// 0: (, uint id, QObjectUserData * data), (&id, data)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QObject11setUserDataEjP15QObjectUserData", ffiqt.FFI_TYPE_VOID, this.cthis, &id, data)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:405
// index:0
// QObjectUserData * userData(uint)
func (this *QObject) UserData(id uint) {
	// 0: (, uint id), (&id)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QObject8userDataEj", ffiqt.FFI_TYPE_VOID, this.cthis, &id)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:409
// index:0
// void destroyed(class QObject *)
func (this *QObject) Destroyed(arg0 unsafe.Pointer) {
	// 0: (, QObject * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QObject9destroyedEPS_", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:413
// index:0
// inline
// QObject * parent()
func (this *QObject) Parent() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QObject6parentEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:415
// index:0
// inline
// bool inherits(const char *)
func (this *QObject) Inherits(classname unsafe.Pointer) {
	// 0: (, const char * classname), (classname)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QObject8inheritsEPKc", ffiqt.FFI_TYPE_VOID, this.cthis, classname)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:419
// index:0
// void deleteLater()
func (this *QObject) DeleteLater() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QObject11deleteLaterEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
