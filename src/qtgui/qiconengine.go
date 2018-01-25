package qtgui

// /usr/include/qt/QtGui/qiconengine.h
// #include <qiconengine.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 64
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
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QIconEngine) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQIconEngineFromPointer(cthis unsafe.Pointer) *QIconEngine {
	return &QIconEngine{&qtrt.CObject{cthis}}
}
func (*QIconEngine) NewFromPointer(cthis unsafe.Pointer) *QIconEngine {
	return NewQIconEngineFromPointer(cthis)
}

// /usr/include/qt/QtGui/qiconengine.h:53
// index:0
// Public
// void QIconEngine()
func NewQIconEngine() *QIconEngine {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QIconEngineC1Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQIconEngineFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qiconengine.h:55
// index:0
// Public virtual
// void ~QIconEngine()
func DeleteQIconEngine(*QIconEngine) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QIconEngineD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qiconengine.h:56
// index:0
// Public pure virtual
// void paint(class QPainter *, const class QRect &, class QIcon::Mode, class QIcon::State)
func (this *QIconEngine) Paint(painter *QPainter /*444 QPainter **/, rect *qtcore.QRect, mode int, state int) {
	var convArg0 = painter.GetCthis()
	var convArg1 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QIconEngine5paintEP8QPainterRK5QRectN5QIcon4ModeENS5_5StateE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, mode, state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qiconengine.h:57
// index:0
// Public virtual
// QSize actualSize(const class QSize &, class QIcon::Mode, class QIcon::State)
func (this *QIconEngine) ActualSize(size *qtcore.QSize, mode int, state int) *qtcore.QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = size.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QIconEngine10actualSizeERK5QSizeN5QIcon4ModeENS3_5StateE", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, mode, state)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qiconengine.h:58
// index:0
// Public virtual
// QPixmap pixmap(const class QSize &, class QIcon::Mode, class QIcon::State)
func (this *QIconEngine) Pixmap(size *qtcore.QSize, mode int, state int) *QPixmap /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = size.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QIconEngine6pixmapERK5QSizeN5QIcon4ModeENS3_5StateE", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, mode, state)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qiconengine.h:60
// index:0
// Public virtual
// void addPixmap(const class QPixmap &, class QIcon::Mode, class QIcon::State)
func (this *QIconEngine) AddPixmap(pixmap *QPixmap, mode int, state int) {
	var convArg0 = pixmap.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QIconEngine9addPixmapERK7QPixmapN5QIcon4ModeENS3_5StateE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode, state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qiconengine.h:61
// index:0
// Public virtual
// void addFile(const class QString &, const class QSize &, class QIcon::Mode, class QIcon::State)
func (this *QIconEngine) AddFile(fileName *qtcore.QString, size *qtcore.QSize, mode int, state int) {
	var convArg0 = fileName.GetCthis()
	var convArg1 = size.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QIconEngine7addFileERK7QStringRK5QSizeN5QIcon4ModeENS6_5StateE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, mode, state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qiconengine.h:63
// index:0
// Public virtual
// QString key()
func (this *QIconEngine) Key() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QIconEngine3keyEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qiconengine.h:64
// index:0
// Public pure virtual
// QIconEngine * clone()
func (this *QIconEngine) Clone() *QIconEngine /*444 QIconEngine **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QIconEngine5cloneEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQIconEngineFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qiconengine.h:65
// index:0
// Public virtual
// bool read(class QDataStream &)
func (this *QIconEngine) Read(in *qtcore.QDataStream) bool {
	var convArg0 = in.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QIconEngine4readER11QDataStream", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qiconengine.h:66
// index:0
// Public virtual
// bool write(class QDataStream &)
func (this *QIconEngine) Write(out *qtcore.QDataStream) bool {
	var convArg0 = out.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QIconEngine5writeER11QDataStream", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qiconengine.h:80
// index:0
// Public virtual
// QString iconName()
func (this *QIconEngine) IconName() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QIconEngine8iconNameEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qiconengine.h:81
// index:0
// Public
// bool isNull()
func (this *QIconEngine) IsNull() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QIconEngine6isNullEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qiconengine.h:82
// index:0
// Public
// QPixmap scaledPixmap(const class QSize &, class QIcon::Mode, class QIcon::State, qreal)
func (this *QIconEngine) ScaledPixmap(size *qtcore.QSize, mode int, state int, scale float64) *QPixmap /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = size.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QIconEngine12scaledPixmapERK5QSizeN5QIcon4ModeENS3_5StateEd", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, mode, state, scale)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qiconengine.h:93
// index:0
// Public virtual
// void virtual_hook(int, void *)
func (this *QIconEngine) Virtual_hook(id int, data unsafe.Pointer /*666*/) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QIconEngine12virtual_hookEiPv", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), id, data)
	gopp.ErrPrint(err, rv)
}

type QIconEngine__IconEngineHook = int

const QIconEngine__AvailableSizesHook QIconEngine__IconEngineHook = 1
const QIconEngine__IconNameHook QIconEngine__IconEngineHook = 2
const QIconEngine__IsNullHook QIconEngine__IconEngineHook = 3
const QIconEngine__ScaledPixmapHook QIconEngine__IconEngineHook = 4

//  body block end
