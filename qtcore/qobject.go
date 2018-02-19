package qtcore

// /usr/include/qt/QtCore/qobject.h
// #include <qobject.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 2
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

// QObject * sender()
func (this *QObject) InheritSender(f func() unsafe.Pointer /*666*/) {
	qtrt.SetAllInheritCallback(this, "sender", f)
}

// int senderSignalIndex()
func (this *QObject) InheritSenderSignalIndex(f func() int) {
	qtrt.SetAllInheritCallback(this, "senderSignalIndex", f)
}

// int receivers(const char *)
func (this *QObject) InheritReceivers(f func(signal string) int) {
	qtrt.SetAllInheritCallback(this, "receivers", f)
}

// bool isSignalConnected(const class QMetaMethod &)
func (this *QObject) InheritIsSignalConnected(f func(signal *QMetaMethod) bool) {
	qtrt.SetAllInheritCallback(this, "isSignalConnected", f)
}

// void timerEvent(class QTimerEvent *)
func (this *QObject) InheritTimerEvent(f func(event *QTimerEvent /*777 QTimerEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "timerEvent", f)
}

// void childEvent(class QChildEvent *)
func (this *QObject) InheritChildEvent(f func(event *QChildEvent /*777 QChildEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "childEvent", f)
}

// void customEvent(class QEvent *)
func (this *QObject) InheritCustomEvent(f func(event *QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "customEvent", f)
}

// void connectNotify(const class QMetaMethod &)
func (this *QObject) InheritConnectNotify(f func(signal *QMetaMethod) /*void*/) {
	qtrt.SetAllInheritCallback(this, "connectNotify", f)
}

// void disconnectNotify(const class QMetaMethod &)
func (this *QObject) InheritDisconnectNotify(f func(signal *QMetaMethod) /*void*/) {
	qtrt.SetAllInheritCallback(this, "disconnectNotify", f)
}

type QObject struct {
	*qtrt.CObject
}
type QObject_ITF interface {
	QObject_PTR() *QObject
}

func (ptr *QObject) QObject_PTR() *QObject { return ptr }

func (this *QObject) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QObject) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQObjectFromPointer(cthis unsafe.Pointer) *QObject {
	return &QObject{&qtrt.CObject{cthis}}
}
func (*QObject) NewFromPointer(cthis unsafe.Pointer) *QObject {
	return NewQObjectFromPointer(cthis)
}

// /usr/include/qt/QtCore/qobject.h:119
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const
func (this *QObject) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QObject10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qobject.h:124
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QObject(QObject *)
func NewQObject(parent QObject_ITF /*777 QObject **/) *QObject {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QObjectC2EPS_", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQObjectFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQObject)
	return gothis
}

// /usr/include/qt/QtCore/qobject.h:124
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QObject(QObject *)
func NewQObject__() *QObject {
	// arg: 0, QObject *=Pointer, QObject=Record,
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN7QObjectC2EPS_", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQObjectFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQObject)
	return gothis
}

// /usr/include/qt/QtCore/qobject.h:125
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QObject()
func DeleteQObject(this *QObject) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QObjectD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qobject.h:127
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QObject) Event(event QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QObject5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobject.h:128
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool eventFilter(QObject *, QEvent *)
func (this *QObject) EventFilter(watched QObject_ITF /*777 QObject **/, event QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if watched != nil && watched.QObject_PTR() != nil {
		convArg0 = watched.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg1 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QObject11eventFilterEPS_P6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobject.h:145
// index:0
// Public Visibility=Default Availability=Available
// [8] QString objectName() const
func (this *QObject) ObjectName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QObject10objectNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qobject.h:146
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setObjectName(const QString &)
func (this *QObject) SetObjectName(name string) {
	var tmpArg0 = NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QObject13setObjectNameERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:148
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isWidgetType() const
func (this *QObject) IsWidgetType() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QObject12isWidgetTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobject.h:149
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isWindowType() const
func (this *QObject) IsWindowType() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QObject12isWindowTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobject.h:151
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool signalsBlocked() const
func (this *QObject) SignalsBlocked() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QObject14signalsBlockedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobject.h:152
// index:0
// Public Visibility=Default Availability=Available
// [1] bool blockSignals(_Bool)
func (this *QObject) BlockSignals(b bool) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QObject12blockSignalsEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), b)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobject.h:154
// index:0
// Public Visibility=Default Availability=Available
// [8] QThread * thread() const
func (this *QObject) Thread() *QThread /*777 QThread **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QObject6threadEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQThreadFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qobject.h:155
// index:0
// Public Visibility=Default Availability=Available
// [-2] void moveToThread(QThread *)
func (this *QObject) MoveToThread(thread QThread_ITF /*777 QThread **/) {
	var convArg0 unsafe.Pointer
	if thread != nil && thread.QThread_PTR() != nil {
		convArg0 = thread.QThread_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QObject12moveToThreadEP7QThread", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:157
// index:0
// Public Visibility=Default Availability=Available
// [4] int startTimer(int, Qt::TimerType)
func (this *QObject) StartTimer(interval int, timerType int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QObject10startTimerEiN2Qt9TimerTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), interval, timerType)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qobject.h:157
// index:0
// Public Visibility=Default Availability=Available
// [4] int startTimer(int, Qt::TimerType)
func (this *QObject) StartTimer__(interval int) int {
	// arg: 1, Qt::TimerType=Elaborated, Qt::TimerType=Enum,
	timerType := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN7QObject10startTimerEiN2Qt9TimerTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), interval, timerType)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qobject.h:165
// index:0
// Public Visibility=Default Availability=Available
// [-2] void killTimer(int)
func (this *QObject) KillTimer(id int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QObject9killTimerEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), id)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:208
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QObjectList & children() const
func (this *QObject) Children() *QObjectList {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QObject8childrenEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQObjectListFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQObjectList)
	return rv2
}

// /usr/include/qt/QtCore/qobject.h:210
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setParent(QObject *)
func (this *QObject) SetParent(parent QObject_ITF /*777 QObject **/) {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QObject9setParentEPS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:211
// index:0
// Public Visibility=Default Availability=Available
// [-2] void installEventFilter(QObject *)
func (this *QObject) InstallEventFilter(filterObj QObject_ITF /*777 QObject **/) {
	var convArg0 unsafe.Pointer
	if filterObj != nil && filterObj.QObject_PTR() != nil {
		convArg0 = filterObj.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QObject18installEventFilterEPS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:212
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeEventFilter(QObject *)
func (this *QObject) RemoveEventFilter(obj QObject_ITF /*777 QObject **/) {
	var convArg0 unsafe.Pointer
	if obj != nil && obj.QObject_PTR() != nil {
		convArg0 = obj.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QObject17removeEventFilterEPS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:214
// index:0
// Public static Visibility=Default Availability=Available
// [8] QMetaObject::Connection connect(const QObject *, const char *, const QObject *, const char *, Qt::ConnectionType)
func (this *QObject) Connect(sender QObject_ITF /*777 const QObject **/, signal string, receiver QObject_ITF /*777 const QObject **/, member string, arg4 int) int {
	var convArg0 unsafe.Pointer
	if sender != nil && sender.QObject_PTR() != nil {
		convArg0 = sender.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(signal)
	defer qtrt.FreeMem(convArg1)
	var convArg2 unsafe.Pointer
	if receiver != nil && receiver.QObject_PTR() != nil {
		convArg2 = receiver.QObject_PTR().GetCthis()
	}
	var convArg3 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg3)
	rv, err := qtrt.InvokeQtFunc6("_ZN7QObject7connectEPKS_PKcS1_S3_N2Qt14ConnectionTypeE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, arg4)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}
func QObject_Connect(sender QObject_ITF /*777 const QObject **/, signal string, receiver QObject_ITF /*777 const QObject **/, member string, arg4 int) int {
	var nilthis *QObject
	rv := nilthis.Connect(sender, signal, receiver, member, arg4)
	return rv
}

// /usr/include/qt/QtCore/qobject.h:214
// index:0
// Public static Visibility=Default Availability=Available
// [8] QMetaObject::Connection connect(const QObject *, const char *, const QObject *, const char *, Qt::ConnectionType)
func (this *QObject) Connect__(sender QObject_ITF /*777 const QObject **/, signal string, receiver QObject_ITF /*777 const QObject **/, member string) int {
	var convArg0 unsafe.Pointer
	if sender != nil && sender.QObject_PTR() != nil {
		convArg0 = sender.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(signal)
	defer qtrt.FreeMem(convArg1)
	var convArg2 unsafe.Pointer
	if receiver != nil && receiver.QObject_PTR() != nil {
		convArg2 = receiver.QObject_PTR().GetCthis()
	}
	var convArg3 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg3)
	// arg: 4, Qt::ConnectionType=Elaborated, Qt::ConnectionType=Enum,
	arg4 := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN7QObject7connectEPKS_PKcS1_S3_N2Qt14ConnectionTypeE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, arg4)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qobject.h:217
// index:1
// Public static Visibility=Default Availability=Available
// [8] QMetaObject::Connection connect(const QObject *, const QMetaMethod &, const QObject *, const QMetaMethod &, Qt::ConnectionType)
func (this *QObject) Connect_1(sender QObject_ITF /*777 const QObject **/, signal QMetaMethod_ITF, receiver QObject_ITF /*777 const QObject **/, method QMetaMethod_ITF, type_ int) int {
	var convArg0 unsafe.Pointer
	if sender != nil && sender.QObject_PTR() != nil {
		convArg0 = sender.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if signal != nil && signal.QMetaMethod_PTR() != nil {
		convArg1 = signal.QMetaMethod_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if receiver != nil && receiver.QObject_PTR() != nil {
		convArg2 = receiver.QObject_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if method != nil && method.QMetaMethod_PTR() != nil {
		convArg3 = method.QMetaMethod_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QObject7connectEPKS_RK11QMetaMethodS1_S4_N2Qt14ConnectionTypeE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, type_)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}
func QObject_Connect_1(sender QObject_ITF /*777 const QObject **/, signal QMetaMethod_ITF, receiver QObject_ITF /*777 const QObject **/, method QMetaMethod_ITF, type_ int) int {
	var nilthis *QObject
	rv := nilthis.Connect_1(sender, signal, receiver, method, type_)
	return rv
}

// /usr/include/qt/QtCore/qobject.h:217
// index:1
// Public static Visibility=Default Availability=Available
// [8] QMetaObject::Connection connect(const QObject *, const QMetaMethod &, const QObject *, const QMetaMethod &, Qt::ConnectionType)
func (this *QObject) Connect_1_(sender QObject_ITF /*777 const QObject **/, signal QMetaMethod_ITF, receiver QObject_ITF /*777 const QObject **/, method QMetaMethod_ITF) int {
	var convArg0 unsafe.Pointer
	if sender != nil && sender.QObject_PTR() != nil {
		convArg0 = sender.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if signal != nil && signal.QMetaMethod_PTR() != nil {
		convArg1 = signal.QMetaMethod_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if receiver != nil && receiver.QObject_PTR() != nil {
		convArg2 = receiver.QObject_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if method != nil && method.QMetaMethod_PTR() != nil {
		convArg3 = method.QMetaMethod_PTR().GetCthis()
	}
	// arg: 4, Qt::ConnectionType=Elaborated, Qt::ConnectionType=Enum,
	type_ := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN7QObject7connectEPKS_RK11QMetaMethodS1_S4_N2Qt14ConnectionTypeE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, type_)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qobject.h:221
// index:2
// Public inline Visibility=Default Availability=Available
// [8] QMetaObject::Connection connect(const QObject *, const char *, const char *, Qt::ConnectionType) const
func (this *QObject) Connect_2(sender QObject_ITF /*777 const QObject **/, signal string, member string, type_ int) int {
	var convArg0 unsafe.Pointer
	if sender != nil && sender.QObject_PTR() != nil {
		convArg0 = sender.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(signal)
	defer qtrt.FreeMem(convArg1)
	var convArg2 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg2)
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QObject7connectEPKS_PKcS3_N2Qt14ConnectionTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, type_)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qobject.h:221
// index:2
// Public inline Visibility=Default Availability=Available
// [8] QMetaObject::Connection connect(const QObject *, const char *, const char *, Qt::ConnectionType) const
func (this *QObject) Connect_2_(sender QObject_ITF /*777 const QObject **/, signal string, member string) int {
	var convArg0 unsafe.Pointer
	if sender != nil && sender.QObject_PTR() != nil {
		convArg0 = sender.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(signal)
	defer qtrt.FreeMem(convArg1)
	var convArg2 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg2)
	// arg: 3, Qt::ConnectionType=Elaborated, Qt::ConnectionType=Enum,
	type_ := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QObject7connectEPKS_PKcS3_N2Qt14ConnectionTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, type_)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qobject.h:343
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool disconnect(const QObject *, const char *, const QObject *, const char *)
func (this *QObject) Disconnect(sender QObject_ITF /*777 const QObject **/, signal string, receiver QObject_ITF /*777 const QObject **/, member string) bool {
	var convArg0 unsafe.Pointer
	if sender != nil && sender.QObject_PTR() != nil {
		convArg0 = sender.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(signal)
	defer qtrt.FreeMem(convArg1)
	var convArg2 unsafe.Pointer
	if receiver != nil && receiver.QObject_PTR() != nil {
		convArg2 = receiver.QObject_PTR().GetCthis()
	}
	var convArg3 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg3)
	rv, err := qtrt.InvokeQtFunc6("_ZN7QObject10disconnectEPKS_PKcS1_S3_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QObject_Disconnect(sender QObject_ITF /*777 const QObject **/, signal string, receiver QObject_ITF /*777 const QObject **/, member string) bool {
	var nilthis *QObject
	rv := nilthis.Disconnect(sender, signal, receiver, member)
	return rv
}

// /usr/include/qt/QtCore/qobject.h:345
// index:1
// Public static Visibility=Default Availability=Available
// [1] bool disconnect(const QObject *, const QMetaMethod &, const QObject *, const QMetaMethod &)
func (this *QObject) Disconnect_1(sender QObject_ITF /*777 const QObject **/, signal QMetaMethod_ITF, receiver QObject_ITF /*777 const QObject **/, member QMetaMethod_ITF) bool {
	var convArg0 unsafe.Pointer
	if sender != nil && sender.QObject_PTR() != nil {
		convArg0 = sender.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if signal != nil && signal.QMetaMethod_PTR() != nil {
		convArg1 = signal.QMetaMethod_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if receiver != nil && receiver.QObject_PTR() != nil {
		convArg2 = receiver.QObject_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if member != nil && member.QMetaMethod_PTR() != nil {
		convArg3 = member.QMetaMethod_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QObject10disconnectEPKS_RK11QMetaMethodS1_S4_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QObject_Disconnect_1(sender QObject_ITF /*777 const QObject **/, signal QMetaMethod_ITF, receiver QObject_ITF /*777 const QObject **/, member QMetaMethod_ITF) bool {
	var nilthis *QObject
	rv := nilthis.Disconnect_1(sender, signal, receiver, member)
	return rv
}

// /usr/include/qt/QtCore/qobject.h:347
// index:2
// Public inline Visibility=Default Availability=Available
// [1] bool disconnect(const char *, const QObject *, const char *) const
func (this *QObject) Disconnect_2(signal string, receiver QObject_ITF /*777 const QObject **/, member string) bool {
	var convArg0 = qtrt.CString(signal)
	defer qtrt.FreeMem(convArg0)
	var convArg1 unsafe.Pointer
	if receiver != nil && receiver.QObject_PTR() != nil {
		convArg1 = receiver.QObject_PTR().GetCthis()
	}
	var convArg2 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg2)
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QObject10disconnectEPKcPKS_S1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobject.h:347
// index:2
// Public inline Visibility=Default Availability=Available
// [1] bool disconnect(const char *, const QObject *, const char *) const
func (this *QObject) Disconnect_2_() bool {
	// arg: 0, const char *=Pointer, =Invalid,
	var convArg0 unsafe.Pointer
	// arg: 1, const QObject *=Pointer, QObject=Record,
	var convArg1 unsafe.Pointer
	// arg: 2, const char *=Pointer, =Invalid,
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QObject10disconnectEPKcPKS_S1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobject.h:347
// index:2
// Public inline Visibility=Default Availability=Available
// [1] bool disconnect(const char *, const QObject *, const char *) const
func (this *QObject) Disconnect_2_1(signal string) bool {
	var convArg0 = qtrt.CString(signal)
	defer qtrt.FreeMem(convArg0)
	// arg: 1, const QObject *=Pointer, QObject=Record,
	var convArg1 unsafe.Pointer
	// arg: 2, const char *=Pointer, =Invalid,
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QObject10disconnectEPKcPKS_S1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobject.h:347
// index:2
// Public inline Visibility=Default Availability=Available
// [1] bool disconnect(const char *, const QObject *, const char *) const
func (this *QObject) Disconnect_2_2(signal string, receiver QObject_ITF /*777 const QObject **/) bool {
	var convArg0 = qtrt.CString(signal)
	defer qtrt.FreeMem(convArg0)
	var convArg1 unsafe.Pointer
	if receiver != nil && receiver.QObject_PTR() != nil {
		convArg1 = receiver.QObject_PTR().GetCthis()
	}
	// arg: 2, const char *=Pointer, =Invalid,
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QObject10disconnectEPKcPKS_S1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobject.h:350
// index:3
// Public inline Visibility=Default Availability=Available
// [1] bool disconnect(const QObject *, const char *) const
func (this *QObject) Disconnect_3(receiver QObject_ITF /*777 const QObject **/, member string) bool {
	var convArg0 unsafe.Pointer
	if receiver != nil && receiver.QObject_PTR() != nil {
		convArg0 = receiver.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QObject10disconnectEPKS_PKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobject.h:350
// index:3
// Public inline Visibility=Default Availability=Available
// [1] bool disconnect(const QObject *, const char *) const
func (this *QObject) Disconnect_3_(receiver QObject_ITF /*777 const QObject **/) bool {
	var convArg0 unsafe.Pointer
	if receiver != nil && receiver.QObject_PTR() != nil {
		convArg0 = receiver.QObject_PTR().GetCthis()
	}
	// arg: 1, const char *=Pointer, =Invalid,
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QObject10disconnectEPKS_PKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobject.h:391
// index:0
// Public Visibility=Default Availability=Available
// [-2] void dumpObjectTree()
func (this *QObject) DumpObjectTree() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QObject14dumpObjectTreeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:394
// index:1
// Public Visibility=Default Availability=Available
// [-2] void dumpObjectTree() const
func (this *QObject) DumpObjectTree_1() {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QObject14dumpObjectTreeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:392
// index:0
// Public Visibility=Default Availability=Available
// [-2] void dumpObjectInfo()
func (this *QObject) DumpObjectInfo() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QObject14dumpObjectInfoEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:395
// index:1
// Public Visibility=Default Availability=Available
// [-2] void dumpObjectInfo() const
func (this *QObject) DumpObjectInfo_1() {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QObject14dumpObjectInfoEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:398
// index:0
// Public Visibility=Default Availability=Available
// [1] bool setProperty(const char *, const QVariant &)
func (this *QObject) SetProperty(name string, value QVariant_ITF) bool {
	var convArg0 = qtrt.CString(name)
	defer qtrt.FreeMem(convArg0)
	var convArg1 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg1 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QObject11setPropertyEPKcRK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobject.h:399
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant property(const char *) const
func (this *QObject) Property(name string) *QVariant /*123*/ {
	var convArg0 = qtrt.CString(name)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QObject8propertyEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtCore/qobject.h:404
// index:0
// Public static Visibility=Default Availability=Available
// [4] uint registerUserData()
func (this *QObject) RegisterUserData() uint {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QObject16registerUserDataEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}
func QObject_RegisterUserData() uint {
	var nilthis *QObject
	rv := nilthis.RegisterUserData()
	return rv
}

// /usr/include/qt/QtCore/qobject.h:405
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUserData(uint, QObjectUserData *)
func (this *QObject) SetUserData(id uint, data QObjectUserData_ITF /*777 QObjectUserData **/) {
	var convArg1 unsafe.Pointer
	if data != nil && data.QObjectUserData_PTR() != nil {
		convArg1 = data.QObjectUserData_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QObject11setUserDataEjP15QObjectUserData", qtrt.FFI_TYPE_POINTER, this.GetCthis(), id, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:406
// index:0
// Public Visibility=Default Availability=Available
// [8] QObjectUserData * userData(uint) const
func (this *QObject) UserData(id uint) *QObjectUserData /*777 QObjectUserData **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QObject8userDataEj", qtrt.FFI_TYPE_POINTER, this.GetCthis(), id)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQObjectUserDataFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qobject.h:410
// index:0
// Public Visibility=Default Availability=Available
// [-2] void destroyed(QObject *)
func (this *QObject) Destroyed(arg0 QObject_ITF /*777 QObject **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QObject_PTR() != nil {
		convArg0 = arg0.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QObject9destroyedEPS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:410
// index:0
// Public Visibility=Default Availability=Available
// [-2] void destroyed(QObject *)
func (this *QObject) Destroyed__() {
	// arg: 0, QObject *=Pointer, QObject=Record,
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN7QObject9destroyedEPS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:414
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QObject * parent() const
func (this *QObject) Parent() *QObject /*777 QObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QObject6parentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qobject.h:416
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool inherits(const char *) const
func (this *QObject) Inherits(classname string) bool {
	var convArg0 = qtrt.CString(classname)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QObject8inheritsEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobject.h:420
// index:0
// Public Visibility=Default Availability=Available
// [-2] void deleteLater()
func (this *QObject) DeleteLater() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QObject11deleteLaterEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:423
// index:0
// Protected Visibility=Default Availability=Available
// [8] QObject * sender() const
func (this *QObject) Sender() *QObject /*777 QObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QObject6senderEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qobject.h:424
// index:0
// Protected Visibility=Default Availability=Available
// [4] int senderSignalIndex() const
func (this *QObject) SenderSignalIndex() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QObject17senderSignalIndexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qobject.h:425
// index:0
// Protected Visibility=Default Availability=Available
// [4] int receivers(const char *) const
func (this *QObject) Receivers(signal string) int {
	var convArg0 = qtrt.CString(signal)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QObject9receiversEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qobject.h:426
// index:0
// Protected Visibility=Default Availability=Available
// [1] bool isSignalConnected(const QMetaMethod &) const
func (this *QObject) IsSignalConnected(signal QMetaMethod_ITF) bool {
	var convArg0 unsafe.Pointer
	if signal != nil && signal.QMetaMethod_PTR() != nil {
		convArg0 = signal.QMetaMethod_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QObject17isSignalConnectedERK11QMetaMethod", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobject.h:428
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void timerEvent(QTimerEvent *)
func (this *QObject) TimerEvent(event QTimerEvent_ITF /*777 QTimerEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QTimerEvent_PTR() != nil {
		convArg0 = event.QTimerEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QObject10timerEventEP11QTimerEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:429
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void childEvent(QChildEvent *)
func (this *QObject) ChildEvent(event QChildEvent_ITF /*777 QChildEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QChildEvent_PTR() != nil {
		convArg0 = event.QChildEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QObject10childEventEP11QChildEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:430
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void customEvent(QEvent *)
func (this *QObject) CustomEvent(event QEvent_ITF /*777 QEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QObject11customEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:432
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void connectNotify(const QMetaMethod &)
func (this *QObject) ConnectNotify(signal QMetaMethod_ITF) {
	var convArg0 unsafe.Pointer
	if signal != nil && signal.QMetaMethod_PTR() != nil {
		convArg0 = signal.QMetaMethod_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QObject13connectNotifyERK11QMetaMethod", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobject.h:433
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void disconnectNotify(const QMetaMethod &)
func (this *QObject) DisconnectNotify(signal QMetaMethod_ITF) {
	var convArg0 unsafe.Pointer
	if signal != nil && signal.QMetaMethod_PTR() != nil {
		convArg0 = signal.QMetaMethod_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QObject16disconnectNotifyERK11QMetaMethod", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

//  body block end

//  keep block begin

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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
}

//  keep block end
