//  header block begin
// /usr/include/qt/QtGui/qiconengine.h
// #include <qiconengine.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 70
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"

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
}

//  ext block end

//  body block begin
type QIconEngine struct {
	*qtrt.CObject
}

func (this *QIconEngine) GetCthis() unsafe.Pointer {
	return this.Cthis
}

// /usr/include/qt/QtGui/qiconengine.h:53
// index:0
// void QIconEngine()
func NewQIconEngine() *QIconEngine {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QIconEngineC1Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQIconEngineFromPointer(cthis)
	return gothis
}
func NewQIconEngineFromPointer(cthis unsafe.Pointer) *QIconEngine {
	return &QIconEngine{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qiconengine.h:55
// index:0
// virtual
// void ~QIconEngine()
func DeleteQIconEngine(*QIconEngine) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QIconEngineD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qiconengine.h:56
// index:0
// pure virtual
// void paint(class QPainter *, const class QRect &, class QIcon::Mode, class QIcon::State)
func (this *QIconEngine) Paint(painter unsafe.Pointer, rect unsafe.Pointer, mode int, state int) {
	// 0: (, painter QPainter *, rect const QRect &, mode QIcon::Mode, state QIcon::State), (painter, rect, &mode, &state)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QIconEngine5paintEP8QPainterRK5QRectN5QIcon4ModeENS5_5StateE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), painter, rect, &mode, &state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qiconengine.h:57
// index:0
// virtual
// QSize actualSize(const class QSize &, class QIcon::Mode, class QIcon::State)
func (this *QIconEngine) ActualSize(size unsafe.Pointer, mode int, state int) {
	// 0: (, size const QSize &, mode QIcon::Mode, state QIcon::State), (size, &mode, &state)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QIconEngine10actualSizeERK5QSizeN5QIcon4ModeENS3_5StateE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), size, &mode, &state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qiconengine.h:58
// index:0
// virtual
// QPixmap pixmap(const class QSize &, class QIcon::Mode, class QIcon::State)
func (this *QIconEngine) Pixmap(size unsafe.Pointer, mode int, state int) {
	// 0: (, size const QSize &, mode QIcon::Mode, state QIcon::State), (size, &mode, &state)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QIconEngine6pixmapERK5QSizeN5QIcon4ModeENS3_5StateE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), size, &mode, &state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qiconengine.h:60
// index:0
// virtual
// void addPixmap(const class QPixmap &, class QIcon::Mode, class QIcon::State)
func (this *QIconEngine) AddPixmap(pixmap unsafe.Pointer, mode int, state int) {
	// 0: (, pixmap const QPixmap &, mode QIcon::Mode, state QIcon::State), (pixmap, &mode, &state)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QIconEngine9addPixmapERK7QPixmapN5QIcon4ModeENS3_5StateE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pixmap, &mode, &state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qiconengine.h:61
// index:0
// virtual
// void addFile(const class QString &, const class QSize &, class QIcon::Mode, class QIcon::State)
func (this *QIconEngine) AddFile(fileName unsafe.Pointer, size unsafe.Pointer, mode int, state int) {
	// 0: (, fileName const QString &, size const QSize &, mode QIcon::Mode, state QIcon::State), (fileName, size, &mode, &state)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QIconEngine7addFileERK7QStringRK5QSizeN5QIcon4ModeENS6_5StateE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), fileName, size, &mode, &state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qiconengine.h:63
// index:0
// virtual
// QString key()
func (this *QIconEngine) Key() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QIconEngine3keyEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qiconengine.h:64
// index:0
// pure virtual
// QIconEngine * clone()
func (this *QIconEngine) Clone() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QIconEngine5cloneEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qiconengine.h:65
// index:0
// virtual
// bool read(class QDataStream &)
func (this *QIconEngine) Read(in unsafe.Pointer) {
	// 0: (, in QDataStream &), (in)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QIconEngine4readER11QDataStream", ffiqt.FFI_TYPE_VOID, this.GetCthis(), in)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qiconengine.h:66
// index:0
// virtual
// bool write(class QDataStream &)
func (this *QIconEngine) Write(out unsafe.Pointer) {
	// 0: (, out QDataStream &), (out)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QIconEngine5writeER11QDataStream", ffiqt.FFI_TYPE_VOID, this.GetCthis(), out)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qiconengine.h:77
// index:0
// virtual
// QList<QSize> availableSizes(class QIcon::Mode, class QIcon::State)
func (this *QIconEngine) AvailableSizes(mode int, state int) {
	// 0: (, mode QIcon::Mode, state QIcon::State), (&mode, &state)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QIconEngine14availableSizesEN5QIcon4ModeENS0_5StateE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &mode, &state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qiconengine.h:80
// index:0
// virtual
// QString iconName()
func (this *QIconEngine) IconName() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QIconEngine8iconNameEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qiconengine.h:81
// index:0
// bool isNull()
func (this *QIconEngine) IsNull() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QIconEngine6isNullEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qiconengine.h:82
// index:0
// QPixmap scaledPixmap(const class QSize &, class QIcon::Mode, class QIcon::State, qreal)
func (this *QIconEngine) ScaledPixmap(size unsafe.Pointer, mode int, state int, scale float64) {
	// 0: (, size const QSize &, mode QIcon::Mode, state QIcon::State, scale qreal), (size, &mode, &state, &scale)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QIconEngine12scaledPixmapERK5QSizeN5QIcon4ModeENS3_5StateEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), size, &mode, &state, &scale)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qiconengine.h:93
// index:0
// virtual
// void virtual_hook(int, void *)
func (this *QIconEngine) Virtual_hook(id int, data unsafe.Pointer) {
	// 0: (, id int, data void *), (&id, data)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QIconEngine12virtual_hookEiPv", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &id, data)
	gopp.ErrPrint(err, rv)
}

//  body block end
