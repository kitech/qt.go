//  header block begin
// /usr/include/qt/QtWidgets/qshortcut.h
// #include <qshortcut.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 26
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"
import "qtgui"

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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qshortcut.h:55
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QShortcut) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QShortcut10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qshortcut.h:63
// index:0
// void QShortcut(class QWidget *)
func NewQShortcut(parent unsafe.Pointer) *QShortcut {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QShortcutC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QShortcut{cthis}
}

// /usr/include/qt/QtWidgets/qshortcut.h:64
// index:1
// void QShortcut(const class QKeySequence &, class QWidget *, const char *, const char *, Qt::ShortcutContext)
func NewQShortcut_1(key unsafe.Pointer, parent unsafe.Pointer, member unsafe.Pointer, ambiguousMember unsafe.Pointer, context int) *QShortcut {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QShortcutC2ERK12QKeySequenceP7QWidgetPKcS6_N2Qt15ShortcutContextE", ffiqt.FFI_TYPE_VOID, cthis, key, parent, member, ambiguousMember, &context)
	gopp.ErrPrint(err, rv)
	return &QShortcut{cthis}
}

// /usr/include/qt/QtWidgets/qshortcut.h:67
// index:0
// virtual
// void ~QShortcut()
func DeleteQShortcut(*QShortcut) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QShortcutD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qshortcut.h:69
// index:0
// void setKey(const class QKeySequence &)
func (this *QShortcut) SetKey(key unsafe.Pointer) {
	// 0: (, const QKeySequence & key), (key)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QShortcut6setKeyERK12QKeySequence", ffiqt.FFI_TYPE_VOID, this.cthis, key)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qshortcut.h:70
// index:0
// QKeySequence key()
func (this *QShortcut) Key() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QShortcut3keyEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qshortcut.h:72
// index:0
// void setEnabled(_Bool)
func (this *QShortcut) SetEnabled(enable bool) {
	// 0: (, bool enable), (&enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QShortcut10setEnabledEb", ffiqt.FFI_TYPE_VOID, this.cthis, &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qshortcut.h:73
// index:0
// bool isEnabled()
func (this *QShortcut) IsEnabled() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QShortcut9isEnabledEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qshortcut.h:75
// index:0
// void setContext(Qt::ShortcutContext)
func (this *QShortcut) SetContext(context int) {
	// 0: (, Qt::ShortcutContext context), (&context)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QShortcut10setContextEN2Qt15ShortcutContextE", ffiqt.FFI_TYPE_VOID, this.cthis, &context)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qshortcut.h:76
// index:0
// Qt::ShortcutContext context()
func (this *QShortcut) Context() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QShortcut7contextEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qshortcut.h:78
// index:0
// void setWhatsThis(const class QString &)
func (this *QShortcut) SetWhatsThis(text unsafe.Pointer) {
	// 0: (, const QString & text), (text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QShortcut12setWhatsThisERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qshortcut.h:79
// index:0
// QString whatsThis()
func (this *QShortcut) WhatsThis() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QShortcut9whatsThisEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qshortcut.h:81
// index:0
// void setAutoRepeat(_Bool)
func (this *QShortcut) SetAutoRepeat(on bool) {
	// 0: (, bool on), (&on)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QShortcut13setAutoRepeatEb", ffiqt.FFI_TYPE_VOID, this.cthis, &on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qshortcut.h:82
// index:0
// bool autoRepeat()
func (this *QShortcut) AutoRepeat() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QShortcut10autoRepeatEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qshortcut.h:84
// index:0
// int id()
func (this *QShortcut) Id() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QShortcut2idEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qshortcut.h:86
// index:0
// inline
// QWidget * parentWidget()
func (this *QShortcut) ParentWidget() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QShortcut12parentWidgetEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qshortcut.h:90
// index:0
// void activated()
func (this *QShortcut) Activated() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QShortcut9activatedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qshortcut.h:91
// index:0
// void activatedAmbiguously()
func (this *QShortcut) ActivatedAmbiguously() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QShortcut20activatedAmbiguouslyEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
