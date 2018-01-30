package qtwidgets

// /usr/include/qt/QtWidgets/qshortcut.h
// #include <qshortcut.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 25
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"

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
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin
type QShortcut struct {
	*qtcore.QObject
}

func (this *QShortcut) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QShortcut) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQShortcutFromPointer(cthis unsafe.Pointer) *QShortcut {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QShortcut{bcthis0}
}
func (*QShortcut) NewFromPointer(cthis unsafe.Pointer) *QShortcut {
	return NewQShortcutFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qshortcut.h:55
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QShortcut) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QShortcut10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qshortcut.h:63
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QShortcut(QWidget *)
func NewQShortcut(parent *QWidget /*777 QWidget **/) *QShortcut {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QShortcutC2EP7QWidget", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQShortcutFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qshortcut.h:64
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QShortcut(const QKeySequence &, QWidget *, const char *, const char *, Qt::ShortcutContext)
func NewQShortcut_1(key *qtgui.QKeySequence, parent *QWidget /*777 QWidget **/, member string, ambiguousMember string, context int) *QShortcut {
	var convArg0 = key.GetCthis()
	var convArg1 = parent.GetCthis()
	var convArg2 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg2)
	var convArg3 = qtrt.CString(ambiguousMember)
	defer qtrt.FreeMem(convArg3)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QShortcutC2ERK12QKeySequenceP7QWidgetPKcS6_N2Qt15ShortcutContextE", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, context)
	gopp.ErrPrint(err, rv)
	gothis := NewQShortcutFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qshortcut.h:67
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QShortcut()
func DeleteQShortcut(*QShortcut) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QShortcutD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qshortcut.h:69
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setKey(const QKeySequence &)
func (this *QShortcut) SetKey(key *qtgui.QKeySequence) {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QShortcut6setKeyERK12QKeySequence", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qshortcut.h:70
// index:0
// Public Visibility=Default Availability=Available
// [8] QKeySequence key()
func (this *QShortcut) Key() *qtgui.QKeySequence /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QShortcut3keyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQKeySequenceFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qshortcut.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setEnabled(_Bool)
func (this *QShortcut) SetEnabled(enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QShortcut10setEnabledEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qshortcut.h:73
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEnabled()
func (this *QShortcut) IsEnabled() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QShortcut9isEnabledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qshortcut.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setContext(Qt::ShortcutContext)
func (this *QShortcut) SetContext(context int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QShortcut10setContextEN2Qt15ShortcutContextE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), context)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qshortcut.h:76
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::ShortcutContext context()
func (this *QShortcut) Context() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QShortcut7contextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qshortcut.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWhatsThis(const QString &)
func (this *QShortcut) SetWhatsThis(text *qtcore.QString) {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QShortcut12setWhatsThisERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qshortcut.h:79
// index:0
// Public Visibility=Default Availability=Available
// [8] QString whatsThis()
func (this *QShortcut) WhatsThis() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QShortcut9whatsThisEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qshortcut.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAutoRepeat(_Bool)
func (this *QShortcut) SetAutoRepeat(on bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QShortcut13setAutoRepeatEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qshortcut.h:82
// index:0
// Public Visibility=Default Availability=Available
// [1] bool autoRepeat()
func (this *QShortcut) AutoRepeat() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QShortcut10autoRepeatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qshortcut.h:84
// index:0
// Public Visibility=Default Availability=Available
// [4] int id()
func (this *QShortcut) Id() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QShortcut2idEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qshortcut.h:86
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QWidget * parentWidget()
func (this *QShortcut) ParentWidget() *QWidget /*777 QWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QShortcut12parentWidgetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qshortcut.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void activated()
func (this *QShortcut) Activated() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QShortcut9activatedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qshortcut.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void activatedAmbiguously()
func (this *QShortcut) ActivatedAmbiguously() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QShortcut20activatedAmbiguouslyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qshortcut.h:94
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QShortcut) Event(e *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QShortcut5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

//  body block end
