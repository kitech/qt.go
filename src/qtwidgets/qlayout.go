//  header block begin
// /usr/include/qt/QtWidgets/qlayout.h
// #include <qlayout.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 6
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
type QLayout struct {
	*qtcore.QObject
	*QLayoutItem
}

func (this *QLayout) GetCthis() unsafe.Pointer {
	return this.QObject.GetCthis()
}

// /usr/include/qt/QtWidgets/qlayout.h:63
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QLayout) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLayout10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:80
// index:0
// void QLayout(class QWidget *)
func NewQLayout(parent unsafe.Pointer) *QLayout {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLayoutC1EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQLayoutFromPointer(cthis)
	return gothis
}
func NewQLayoutFromPointer(cthis unsafe.Pointer) *QLayout {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	bcthis1 := NewQLayoutItemFromPointer(cthis)
	return &QLayout{bcthis0, bcthis1}
}

// /usr/include/qt/QtWidgets/qlayout.h:81
// index:1
// void QLayout()
func NewQLayout_1() *QLayout {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLayoutC1Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQLayoutFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qlayout.h:153
// index:2
// void QLayout(class QLayoutPrivate &, class QLayout *, class QWidget *)
func NewQLayout_2(d unsafe.Pointer, arg1 unsafe.Pointer, arg2 unsafe.Pointer) *QLayout {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLayoutC1ER14QLayoutPrivatePS_P7QWidget", ffiqt.FFI_TYPE_VOID, cthis, d, arg1, arg2)
	gopp.ErrPrint(err, rv)
	gothis := NewQLayoutFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qlayout.h:82
// index:0
// virtual
// void ~QLayout()
func DeleteQLayout(*QLayout) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLayoutD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:84
// index:0
// int margin()
func (this *QLayout) Margin() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLayout6marginEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:85
// index:0
// int spacing()
func (this *QLayout) Spacing() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLayout7spacingEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:87
// index:0
// void setMargin(int)
func (this *QLayout) SetMargin(arg0 int) {
	// 0: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLayout9setMarginEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:88
// index:0
// void setSpacing(int)
func (this *QLayout) SetSpacing(arg0 int) {
	// 0: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLayout10setSpacingEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:90
// index:0
// void setContentsMargins(int, int, int, int)
func (this *QLayout) SetContentsMargins(left int, top int, right int, bottom int) {
	// 0: (, left int, top int, right int, bottom int), (&left, &top, &right, &bottom)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLayout18setContentsMarginsEiiii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &left, &top, &right, &bottom)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:91
// index:1
// void setContentsMargins(const class QMargins &)
func (this *QLayout) SetContentsMargins_1(margins unsafe.Pointer) {
	// 1: (, margins const QMargins &), (margins)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLayout18setContentsMarginsERK8QMargins", ffiqt.FFI_TYPE_VOID, this.GetCthis(), margins)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:92
// index:0
// void getContentsMargins(int *, int *, int *, int *)
func (this *QLayout) GetContentsMargins(left unsafe.Pointer, top unsafe.Pointer, right unsafe.Pointer, bottom unsafe.Pointer) {
	// 0: (, left int *, top int *, right int *, bottom int *), (left, top, right, bottom)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLayout18getContentsMarginsEPiS0_S0_S0_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), left, top, right, bottom)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:93
// index:0
// QMargins contentsMargins()
func (this *QLayout) ContentsMargins() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLayout15contentsMarginsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:94
// index:0
// QRect contentsRect()
func (this *QLayout) ContentsRect() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLayout12contentsRectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:96
// index:0
// bool setAlignment(class QWidget *, Qt::Alignment)
func (this *QLayout) SetAlignment(w unsafe.Pointer, alignment int) {
	// 0: (, w QWidget *, QFlags<Qt::AlignmentFlag> alignment), (w, &alignment)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLayout12setAlignmentEP7QWidget6QFlagsIN2Qt13AlignmentFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), w, &alignment)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:97
// index:1
// bool setAlignment(class QLayout *, Qt::Alignment)
func (this *QLayout) SetAlignment_1(l unsafe.Pointer, alignment int) {
	// 1: (, l QLayout *, QFlags<Qt::AlignmentFlag> alignment), (l, &alignment)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLayout12setAlignmentEPS_6QFlagsIN2Qt13AlignmentFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), l, &alignment)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:100
// index:0
// void setSizeConstraint(enum QLayout::SizeConstraint)
func (this *QLayout) SetSizeConstraint(arg0 int) {
	// 0: (, QLayout::SizeConstraint arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLayout17setSizeConstraintENS_14SizeConstraintE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:101
// index:0
// QLayout::SizeConstraint sizeConstraint()
func (this *QLayout) SizeConstraint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLayout14sizeConstraintEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:102
// index:0
// void setMenuBar(class QWidget *)
func (this *QLayout) SetMenuBar(w unsafe.Pointer) {
	// 0: (, w QWidget *), (w)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLayout10setMenuBarEP7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), w)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:103
// index:0
// QWidget * menuBar()
func (this *QLayout) MenuBar() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLayout7menuBarEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:105
// index:0
// QWidget * parentWidget()
func (this *QLayout) ParentWidget() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLayout12parentWidgetEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:107
// index:0
// virtual
// void invalidate()
func (this *QLayout) Invalidate() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLayout10invalidateEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:108
// index:0
// virtual
// QRect geometry()
func (this *QLayout) Geometry() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLayout8geometryEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:109
// index:0
// bool activate()
func (this *QLayout) Activate() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLayout8activateEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:110
// index:0
// void update()
func (this *QLayout) Update() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLayout6updateEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:112
// index:0
// void addWidget(class QWidget *)
func (this *QLayout) AddWidget(w unsafe.Pointer) {
	// 0: (, w QWidget *), (w)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLayout9addWidgetEP7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), w)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:113
// index:0
// pure virtual
// void addItem(class QLayoutItem *)
func (this *QLayout) AddItem(arg0 unsafe.Pointer) {
	// 0: (, QLayoutItem * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLayout7addItemEP11QLayoutItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:115
// index:0
// void removeWidget(class QWidget *)
func (this *QLayout) RemoveWidget(w unsafe.Pointer) {
	// 0: (, w QWidget *), (w)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLayout12removeWidgetEP7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), w)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:116
// index:0
// void removeItem(class QLayoutItem *)
func (this *QLayout) RemoveItem(arg0 unsafe.Pointer) {
	// 0: (, QLayoutItem * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLayout10removeItemEP11QLayoutItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:118
// index:0
// virtual
// Qt::Orientations expandingDirections()
func (this *QLayout) ExpandingDirections() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLayout19expandingDirectionsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:119
// index:0
// virtual
// QSize minimumSize()
func (this *QLayout) MinimumSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLayout11minimumSizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:120
// index:0
// virtual
// QSize maximumSize()
func (this *QLayout) MaximumSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLayout11maximumSizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:121
// index:0
// virtual
// void setGeometry(const class QRect &)
func (this *QLayout) SetGeometry(arg0 unsafe.Pointer) {
	// 0: (, const QRect & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLayout11setGeometryERK5QRect", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:122
// index:0
// pure virtual
// QLayoutItem * itemAt(int)
func (this *QLayout) ItemAt(index int) {
	// 0: (, index int), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLayout6itemAtEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:123
// index:0
// pure virtual
// QLayoutItem * takeAt(int)
func (this *QLayout) TakeAt(index int) {
	// 0: (, index int), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLayout6takeAtEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:124
// index:0
// virtual
// int indexOf(class QWidget *)
func (this *QLayout) IndexOf(arg0 unsafe.Pointer) {
	// 0: (, QWidget * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLayout7indexOfEP7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:125
// index:0
// pure virtual
// int count()
func (this *QLayout) Count() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLayout5countEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:126
// index:0
// virtual
// bool isEmpty()
func (this *QLayout) IsEmpty() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLayout7isEmptyEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:127
// index:0
// virtual
// QSizePolicy::ControlTypes controlTypes()
func (this *QLayout) ControlTypes() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLayout12controlTypesEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:130
// index:0
// QLayoutItem * replaceWidget(class QWidget *, class QWidget *, Qt::FindChildOptions)
func (this *QLayout) ReplaceWidget(from unsafe.Pointer, to unsafe.Pointer, options int) {
	// 0: (, from QWidget *, to QWidget *, QFlags<Qt::FindChildOption> options), (from, to, &options)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLayout13replaceWidgetEP7QWidgetS1_6QFlagsIN2Qt15FindChildOptionEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), from, to, &options)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:132
// index:0
// int totalHeightForWidth(int)
func (this *QLayout) TotalHeightForWidth(w int) {
	// 0: (, w int), (&w)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLayout19totalHeightForWidthEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &w)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:133
// index:0
// QSize totalMinimumSize()
func (this *QLayout) TotalMinimumSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLayout16totalMinimumSizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:134
// index:0
// QSize totalMaximumSize()
func (this *QLayout) TotalMaximumSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLayout16totalMaximumSizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:135
// index:0
// QSize totalSizeHint()
func (this *QLayout) TotalSizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLayout13totalSizeHintEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:136
// index:0
// virtual
// QLayout * layout()
func (this *QLayout) Layout() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLayout6layoutEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:138
// index:0
// void setEnabled(_Bool)
func (this *QLayout) SetEnabled(arg0 bool) {
	// 0: (, bool arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLayout10setEnabledEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:139
// index:0
// bool isEnabled()
func (this *QLayout) IsEnabled() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLayout9isEnabledEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:142
// index:0
// static
// QSize closestAcceptableSize(const class QWidget *, const class QSize &)
func (this *QLayout) ClosestAcceptableSize(w unsafe.Pointer, s unsafe.Pointer) {
	// 0: (w const QWidget *, s const QSize &), (w, s)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLayout21closestAcceptableSizeEPK7QWidgetRK5QSize", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QLayout_ClosestAcceptableSize(w unsafe.Pointer, s unsafe.Pointer) {
	// 0: (w const QWidget *, s const QSize &), (w, s)
	var nilthis *QLayout
	nilthis.ClosestAcceptableSize(w, s)
}

// /usr/include/qt/QtWidgets/qlayout.h:145
// index:0
// void widgetEvent(class QEvent *)
func (this *QLayout) WidgetEvent(arg0 unsafe.Pointer) {
	// 0: (, QEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLayout11widgetEventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:146
// index:0
// virtual
// void childEvent(class QChildEvent *)
func (this *QLayout) ChildEvent(e unsafe.Pointer) {
	// 0: (, e QChildEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLayout10childEventEP11QChildEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:147
// index:0
// void addChildLayout(class QLayout *)
func (this *QLayout) AddChildLayout(l unsafe.Pointer) {
	// 0: (, l QLayout *), (l)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLayout14addChildLayoutEPS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), l)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:148
// index:0
// void addChildWidget(class QWidget *)
func (this *QLayout) AddChildWidget(w unsafe.Pointer) {
	// 0: (, w QWidget *), (w)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLayout14addChildWidgetEP7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), w)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:149
// index:0
// bool adoptLayout(class QLayout *)
func (this *QLayout) AdoptLayout(layout unsafe.Pointer) {
	// 0: (, layout QLayout *), (layout)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QLayout11adoptLayoutEPS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), layout)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:151
// index:0
// QRect alignmentRect(const class QRect &)
func (this *QLayout) AlignmentRect(arg0 unsafe.Pointer) {
	// 0: (, const QRect & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QLayout13alignmentRectERK5QRect", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
