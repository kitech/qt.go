//  header block begin
// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h
// #include <qgraphicslinearlayout.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 31
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
type QGraphicsLinearLayout struct {
	*QGraphicsLayout
}

func (this *QGraphicsLinearLayout) GetCthis() unsafe.Pointer {
	return this.QGraphicsLayout.GetCthis()
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:56
// index:0
// void QGraphicsLinearLayout(class QGraphicsLayoutItem *)
func NewQGraphicsLinearLayout(parent unsafe.Pointer) *QGraphicsLinearLayout {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGraphicsLinearLayoutC2EP19QGraphicsLayoutItem", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsLinearLayoutFromPointer(cthis)
	return gothis
}
func NewQGraphicsLinearLayoutFromPointer(cthis unsafe.Pointer) *QGraphicsLinearLayout {
	bcthis0 := NewQGraphicsLayoutFromPointer(cthis)
	return &QGraphicsLinearLayout{bcthis0}
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:57
// index:1
// void QGraphicsLinearLayout(Qt::Orientation, class QGraphicsLayoutItem *)
func NewQGraphicsLinearLayout_1(orientation int, parent unsafe.Pointer) *QGraphicsLinearLayout {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGraphicsLinearLayoutC2EN2Qt11OrientationEP19QGraphicsLayoutItem", ffiqt.FFI_TYPE_VOID, cthis, &orientation, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsLinearLayoutFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:58
// index:0
// virtual
// void ~QGraphicsLinearLayout()
func DeleteQGraphicsLinearLayout(*QGraphicsLinearLayout) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGraphicsLinearLayoutD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:60
// index:0
// void setOrientation(Qt::Orientation)
func (this *QGraphicsLinearLayout) SetOrientation(orientation int) {
	// 0: (, orientation Qt::Orientation), (&orientation)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGraphicsLinearLayout14setOrientationEN2Qt11OrientationE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &orientation)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:61
// index:0
// Qt::Orientation orientation()
func (this *QGraphicsLinearLayout) Orientation() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QGraphicsLinearLayout11orientationEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:63
// index:0
// inline
// void addItem(class QGraphicsLayoutItem *)
func (this *QGraphicsLinearLayout) AddItem(item unsafe.Pointer) {
	// 0: (, item QGraphicsLayoutItem *), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGraphicsLinearLayout7addItemEP19QGraphicsLayoutItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:64
// index:0
// inline
// void addStretch(int)
func (this *QGraphicsLinearLayout) AddStretch(stretch int) {
	// 0: (, stretch int), (&stretch)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGraphicsLinearLayout10addStretchEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &stretch)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:66
// index:0
// void insertItem(int, class QGraphicsLayoutItem *)
func (this *QGraphicsLinearLayout) InsertItem(index int, item unsafe.Pointer) {
	// 0: (, index int, item QGraphicsLayoutItem *), (&index, item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGraphicsLinearLayout10insertItemEiP19QGraphicsLayoutItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:67
// index:0
// void insertStretch(int, int)
func (this *QGraphicsLinearLayout) InsertStretch(index int, stretch int) {
	// 0: (, index int, stretch int), (&index, &stretch)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGraphicsLinearLayout13insertStretchEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index, &stretch)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:69
// index:0
// void removeItem(class QGraphicsLayoutItem *)
func (this *QGraphicsLinearLayout) RemoveItem(item unsafe.Pointer) {
	// 0: (, item QGraphicsLayoutItem *), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGraphicsLinearLayout10removeItemEP19QGraphicsLayoutItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:70
// index:0
// virtual
// void removeAt(int)
func (this *QGraphicsLinearLayout) RemoveAt(index int) {
	// 0: (, index int), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGraphicsLinearLayout8removeAtEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:72
// index:0
// void setSpacing(qreal)
func (this *QGraphicsLinearLayout) SetSpacing(spacing float64) {
	// 0: (, spacing qreal), (&spacing)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGraphicsLinearLayout10setSpacingEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &spacing)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:73
// index:0
// qreal spacing()
func (this *QGraphicsLinearLayout) Spacing() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QGraphicsLinearLayout7spacingEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:74
// index:0
// void setItemSpacing(int, qreal)
func (this *QGraphicsLinearLayout) SetItemSpacing(index int, spacing float64) {
	// 0: (, index int, spacing qreal), (&index, &spacing)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGraphicsLinearLayout14setItemSpacingEid", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index, &spacing)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:75
// index:0
// qreal itemSpacing(int)
func (this *QGraphicsLinearLayout) ItemSpacing(index int) {
	// 0: (, index int), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QGraphicsLinearLayout11itemSpacingEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:77
// index:0
// void setStretchFactor(class QGraphicsLayoutItem *, int)
func (this *QGraphicsLinearLayout) SetStretchFactor(item unsafe.Pointer, stretch int) {
	// 0: (, item QGraphicsLayoutItem *, stretch int), (item, &stretch)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGraphicsLinearLayout16setStretchFactorEP19QGraphicsLayoutItemi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item, &stretch)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:78
// index:0
// int stretchFactor(class QGraphicsLayoutItem *)
func (this *QGraphicsLinearLayout) StretchFactor(item unsafe.Pointer) {
	// 0: (, item QGraphicsLayoutItem *), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QGraphicsLinearLayout13stretchFactorEP19QGraphicsLayoutItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:80
// index:0
// void setAlignment(class QGraphicsLayoutItem *, Qt::Alignment)
func (this *QGraphicsLinearLayout) SetAlignment(item unsafe.Pointer, alignment int) {
	// 0: (, item QGraphicsLayoutItem *, QFlags<Qt::AlignmentFlag> alignment), (item, &alignment)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGraphicsLinearLayout12setAlignmentEP19QGraphicsLayoutItem6QFlagsIN2Qt13AlignmentFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item, &alignment)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:81
// index:0
// Qt::Alignment alignment(class QGraphicsLayoutItem *)
func (this *QGraphicsLinearLayout) Alignment(item unsafe.Pointer) {
	// 0: (, item QGraphicsLayoutItem *), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QGraphicsLinearLayout9alignmentEP19QGraphicsLayoutItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:83
// index:0
// virtual
// void setGeometry(const class QRectF &)
func (this *QGraphicsLinearLayout) SetGeometry(rect unsafe.Pointer) {
	// 0: (, rect const QRectF &), (rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGraphicsLinearLayout11setGeometryERK6QRectF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:85
// index:0
// virtual
// int count()
func (this *QGraphicsLinearLayout) Count() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QGraphicsLinearLayout5countEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:86
// index:0
// virtual
// QGraphicsLayoutItem * itemAt(int)
func (this *QGraphicsLinearLayout) ItemAt(index int) {
	// 0: (, index int), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QGraphicsLinearLayout6itemAtEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:88
// index:0
// virtual
// void invalidate()
func (this *QGraphicsLinearLayout) Invalidate() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QGraphicsLinearLayout10invalidateEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:89
// index:0
// virtual
// QSizeF sizeHint(Qt::SizeHint, const class QSizeF &)
func (this *QGraphicsLinearLayout) SizeHint(which int, constraint unsafe.Pointer) {
	// 0: (, which Qt::SizeHint, constraint const QSizeF &), (&which, constraint)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QGraphicsLinearLayout8sizeHintEN2Qt8SizeHintERK6QSizeF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &which, constraint)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:95
// index:0
// void dump(int)
func (this *QGraphicsLinearLayout) Dump(indent int) {
	// 0: (, indent int), (&indent)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QGraphicsLinearLayout4dumpEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &indent)
	gopp.ErrPrint(err, rv)
}

//  body block end
