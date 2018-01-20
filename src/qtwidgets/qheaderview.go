//  header block begin
// /usr/include/qt/QtWidgets/qheaderview.h
// #include <qheaderview.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 27
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
type QHeaderView struct {
	*QAbstractItemView
}

func (this *QHeaderView) GetCthis() unsafe.Pointer {
	return this.QAbstractItemView.GetCthis()
}
func NewQHeaderViewFromPointer(cthis unsafe.Pointer) *QHeaderView {
	bcthis0 := NewQAbstractItemViewFromPointer(cthis)
	return &QHeaderView{bcthis0}
}

// /usr/include/qt/QtWidgets/qheaderview.h:55
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QHeaderView) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qheaderview.h:77
// index:0
// Public
// void QHeaderView(Qt::Orientation, class QWidget *)
func NewQHeaderView(orientation int, parent unsafe.Pointer) *QHeaderView {
	cthis := qtrt.Calloc(1, 256) // 48
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderViewC2EN2Qt11OrientationEP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, &orientation, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQHeaderViewFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qheaderview.h:78
// index:0
// Public virtual
// void ~QHeaderView()
func DeleteQHeaderView(*QHeaderView) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderViewD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:80
// index:0
// Public virtual
// void setModel(class QAbstractItemModel *)
func (this *QHeaderView) SetModel(model unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView8setModelEP18QAbstractItemModel", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), model)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:82
// index:0
// Public
// Qt::Orientation orientation()
func (this *QHeaderView) Orientation() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView11orientationEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qheaderview.h:83
// index:0
// Public
// int offset()
func (this *QHeaderView) Offset() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView6offsetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qheaderview.h:84
// index:0
// Public
// int length()
func (this *QHeaderView) Length() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView6lengthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qheaderview.h:85
// index:0
// Public virtual
// QSize sizeHint()
func (this *QHeaderView) SizeHint() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView8sizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qheaderview.h:86
// index:0
// Public virtual
// void setVisible(_Bool)
func (this *QHeaderView) SetVisible(v bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView10setVisibleEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &v)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:87
// index:0
// Public
// int sectionSizeHint(int)
func (this *QHeaderView) SectionSizeHint(logicalIndex int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView15sectionSizeHintEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &logicalIndex)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qheaderview.h:89
// index:0
// Public
// int visualIndexAt(int)
func (this *QHeaderView) VisualIndexAt(position int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView13visualIndexAtEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &position)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qheaderview.h:90
// index:0
// Public
// int logicalIndexAt(int)
func (this *QHeaderView) LogicalIndexAt(position int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView14logicalIndexAtEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &position)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qheaderview.h:92
// index:1
// Public inline
// int logicalIndexAt(int, int)
func (this *QHeaderView) LogicalIndexAt_1(x int, y int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView14logicalIndexAtEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &x, &y)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qheaderview.h:93
// index:2
// Public inline
// int logicalIndexAt(const class QPoint &)
func (this *QHeaderView) LogicalIndexAt_2(pos *qtcore.QPoint) interface{} {
	var convArg0 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView14logicalIndexAtERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qheaderview.h:95
// index:0
// Public
// int sectionSize(int)
func (this *QHeaderView) SectionSize(logicalIndex int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView11sectionSizeEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &logicalIndex)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qheaderview.h:96
// index:0
// Public
// int sectionPosition(int)
func (this *QHeaderView) SectionPosition(logicalIndex int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView15sectionPositionEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &logicalIndex)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qheaderview.h:97
// index:0
// Public
// int sectionViewportPosition(int)
func (this *QHeaderView) SectionViewportPosition(logicalIndex int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView23sectionViewportPositionEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &logicalIndex)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qheaderview.h:99
// index:0
// Public
// void moveSection(int, int)
func (this *QHeaderView) MoveSection(from int, to int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView11moveSectionEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &from, &to)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:100
// index:0
// Public
// void swapSections(int, int)
func (this *QHeaderView) SwapSections(first int, second int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView12swapSectionsEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &first, &second)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:101
// index:0
// Public
// void resizeSection(int, int)
func (this *QHeaderView) ResizeSection(logicalIndex int, size int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView13resizeSectionEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &logicalIndex, &size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:102
// index:0
// Public
// void resizeSections(class QHeaderView::ResizeMode)
func (this *QHeaderView) ResizeSections(mode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView14resizeSectionsENS_10ResizeModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:206
// index:1
// Protected
// void resizeSections()
func (this *QHeaderView) ResizeSections_1() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView14resizeSectionsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:104
// index:0
// Public
// bool isSectionHidden(int)
func (this *QHeaderView) IsSectionHidden(logicalIndex int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView15isSectionHiddenEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &logicalIndex)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qheaderview.h:105
// index:0
// Public
// void setSectionHidden(int, _Bool)
func (this *QHeaderView) SetSectionHidden(logicalIndex int, hide bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView16setSectionHiddenEib", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &logicalIndex, &hide)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:106
// index:0
// Public
// int hiddenSectionCount()
func (this *QHeaderView) HiddenSectionCount() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView18hiddenSectionCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qheaderview.h:108
// index:0
// Public inline
// void hideSection(int)
func (this *QHeaderView) HideSection(logicalIndex int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView11hideSectionEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &logicalIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:109
// index:0
// Public inline
// void showSection(int)
func (this *QHeaderView) ShowSection(logicalIndex int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView11showSectionEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &logicalIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:111
// index:0
// Public
// int count()
func (this *QHeaderView) Count() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView5countEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qheaderview.h:112
// index:0
// Public
// int visualIndex(int)
func (this *QHeaderView) VisualIndex(logicalIndex int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView11visualIndexEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &logicalIndex)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qheaderview.h:113
// index:0
// Public
// int logicalIndex(int)
func (this *QHeaderView) LogicalIndex(visualIndex int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView12logicalIndexEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &visualIndex)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qheaderview.h:115
// index:0
// Public
// void setSectionsMovable(_Bool)
func (this *QHeaderView) SetSectionsMovable(movable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView18setSectionsMovableEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &movable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:116
// index:0
// Public
// bool sectionsMovable()
func (this *QHeaderView) SectionsMovable() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView15sectionsMovableEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qheaderview.h:122
// index:0
// Public
// void setSectionsClickable(_Bool)
func (this *QHeaderView) SetSectionsClickable(clickable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView20setSectionsClickableEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &clickable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:123
// index:0
// Public
// bool sectionsClickable()
func (this *QHeaderView) SectionsClickable() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView17sectionsClickableEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qheaderview.h:129
// index:0
// Public
// void setHighlightSections(_Bool)
func (this *QHeaderView) SetHighlightSections(highlight bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView20setHighlightSectionsEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &highlight)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:130
// index:0
// Public
// bool highlightSections()
func (this *QHeaderView) HighlightSections() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView17highlightSectionsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qheaderview.h:132
// index:0
// Public
// QHeaderView::ResizeMode sectionResizeMode(int)
func (this *QHeaderView) SectionResizeMode(logicalIndex int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView17sectionResizeModeEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &logicalIndex)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qheaderview.h:133
// index:0
// Public
// void setSectionResizeMode(enum QHeaderView::ResizeMode)
func (this *QHeaderView) SetSectionResizeMode(mode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView20setSectionResizeModeENS_10ResizeModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:134
// index:1
// Public
// void setSectionResizeMode(int, enum QHeaderView::ResizeMode)
func (this *QHeaderView) SetSectionResizeMode_1(logicalIndex int, mode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView20setSectionResizeModeEiNS_10ResizeModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &logicalIndex, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:136
// index:0
// Public
// void setResizeContentsPrecision(int)
func (this *QHeaderView) SetResizeContentsPrecision(precision int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView26setResizeContentsPrecisionEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &precision)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:137
// index:0
// Public
// int resizeContentsPrecision()
func (this *QHeaderView) ResizeContentsPrecision() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView23resizeContentsPrecisionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qheaderview.h:148
// index:0
// Public
// int stretchSectionCount()
func (this *QHeaderView) StretchSectionCount() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView19stretchSectionCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qheaderview.h:150
// index:0
// Public
// void setSortIndicatorShown(_Bool)
func (this *QHeaderView) SetSortIndicatorShown(show bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView21setSortIndicatorShownEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &show)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:151
// index:0
// Public
// bool isSortIndicatorShown()
func (this *QHeaderView) IsSortIndicatorShown() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView20isSortIndicatorShownEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qheaderview.h:153
// index:0
// Public
// void setSortIndicator(int, Qt::SortOrder)
func (this *QHeaderView) SetSortIndicator(logicalIndex int, order int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView16setSortIndicatorEiN2Qt9SortOrderE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &logicalIndex, &order)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:154
// index:0
// Public
// int sortIndicatorSection()
func (this *QHeaderView) SortIndicatorSection() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView20sortIndicatorSectionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qheaderview.h:155
// index:0
// Public
// Qt::SortOrder sortIndicatorOrder()
func (this *QHeaderView) SortIndicatorOrder() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView18sortIndicatorOrderEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qheaderview.h:157
// index:0
// Public
// bool stretchLastSection()
func (this *QHeaderView) StretchLastSection() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView18stretchLastSectionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qheaderview.h:158
// index:0
// Public
// void setStretchLastSection(_Bool)
func (this *QHeaderView) SetStretchLastSection(stretch bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView21setStretchLastSectionEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &stretch)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:160
// index:0
// Public
// bool cascadingSectionResizes()
func (this *QHeaderView) CascadingSectionResizes() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView23cascadingSectionResizesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qheaderview.h:161
// index:0
// Public
// void setCascadingSectionResizes(_Bool)
func (this *QHeaderView) SetCascadingSectionResizes(enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView26setCascadingSectionResizesEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:163
// index:0
// Public
// int defaultSectionSize()
func (this *QHeaderView) DefaultSectionSize() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView18defaultSectionSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qheaderview.h:164
// index:0
// Public
// void setDefaultSectionSize(int)
func (this *QHeaderView) SetDefaultSectionSize(size int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView21setDefaultSectionSizeEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:165
// index:0
// Public
// void resetDefaultSectionSize()
func (this *QHeaderView) ResetDefaultSectionSize() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView23resetDefaultSectionSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:167
// index:0
// Public
// int minimumSectionSize()
func (this *QHeaderView) MinimumSectionSize() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView18minimumSectionSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qheaderview.h:168
// index:0
// Public
// void setMinimumSectionSize(int)
func (this *QHeaderView) SetMinimumSectionSize(size int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView21setMinimumSectionSizeEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:169
// index:0
// Public
// int maximumSectionSize()
func (this *QHeaderView) MaximumSectionSize() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView18maximumSectionSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qheaderview.h:170
// index:0
// Public
// void setMaximumSectionSize(int)
func (this *QHeaderView) SetMaximumSectionSize(size int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView21setMaximumSectionSizeEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:172
// index:0
// Public
// Qt::Alignment defaultAlignment()
func (this *QHeaderView) DefaultAlignment() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView16defaultAlignmentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qheaderview.h:175
// index:0
// Public virtual
// void doItemsLayout()
func (this *QHeaderView) DoItemsLayout() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView13doItemsLayoutEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:176
// index:0
// Public
// bool sectionsMoved()
func (this *QHeaderView) SectionsMoved() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView13sectionsMovedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qheaderview.h:177
// index:0
// Public
// bool sectionsHidden()
func (this *QHeaderView) SectionsHidden() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView14sectionsHiddenEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qheaderview.h:180
// index:0
// Public
// QByteArray saveState()
func (this *QHeaderView) SaveState() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView9saveStateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qheaderview.h:181
// index:0
// Public
// bool restoreState(const class QByteArray &)
func (this *QHeaderView) RestoreState(state *qtcore.QByteArray) interface{} {
	var convArg0 = state.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView12restoreStateERK10QByteArray", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qheaderview.h:184
// index:0
// Public virtual
// void reset()
func (this *QHeaderView) Reset() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView5resetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:187
// index:0
// Public
// void setOffset(int)
func (this *QHeaderView) SetOffset(offset int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView9setOffsetEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &offset)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:188
// index:0
// Public
// void setOffsetToSectionPosition(int)
func (this *QHeaderView) SetOffsetToSectionPosition(visualIndex int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView26setOffsetToSectionPositionEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &visualIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:189
// index:0
// Public
// void setOffsetToLastSection()
func (this *QHeaderView) SetOffsetToLastSection() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView22setOffsetToLastSectionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:190
// index:0
// Public
// void headerDataChanged(Qt::Orientation, int, int)
func (this *QHeaderView) HeaderDataChanged(orientation int, logicalFirst int, logicalLast int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView17headerDataChangedEN2Qt11OrientationEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &orientation, &logicalFirst, &logicalLast)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:193
// index:0
// Public
// void sectionMoved(int, int, int)
func (this *QHeaderView) SectionMoved(logicalIndex int, oldVisualIndex int, newVisualIndex int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView12sectionMovedEiii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &logicalIndex, &oldVisualIndex, &newVisualIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:194
// index:0
// Public
// void sectionResized(int, int, int)
func (this *QHeaderView) SectionResized(logicalIndex int, oldSize int, newSize int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView14sectionResizedEiii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &logicalIndex, &oldSize, &newSize)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:195
// index:0
// Public
// void sectionPressed(int)
func (this *QHeaderView) SectionPressed(logicalIndex int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView14sectionPressedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &logicalIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:196
// index:0
// Public
// void sectionClicked(int)
func (this *QHeaderView) SectionClicked(logicalIndex int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView14sectionClickedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &logicalIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:197
// index:0
// Public
// void sectionEntered(int)
func (this *QHeaderView) SectionEntered(logicalIndex int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView14sectionEnteredEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &logicalIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:198
// index:0
// Public
// void sectionDoubleClicked(int)
func (this *QHeaderView) SectionDoubleClicked(logicalIndex int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView20sectionDoubleClickedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &logicalIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:199
// index:0
// Public
// void sectionCountChanged(int, int)
func (this *QHeaderView) SectionCountChanged(oldCount int, newCount int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView19sectionCountChangedEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &oldCount, &newCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:200
// index:0
// Public
// void sectionHandleDoubleClicked(int)
func (this *QHeaderView) SectionHandleDoubleClicked(logicalIndex int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView26sectionHandleDoubleClickedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &logicalIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:201
// index:0
// Public
// void geometriesChanged()
func (this *QHeaderView) GeometriesChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView17geometriesChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:202
// index:0
// Public
// void sortIndicatorChanged(int, Qt::SortOrder)
func (this *QHeaderView) SortIndicatorChanged(logicalIndex int, order int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView20sortIndicatorChangedEiN2Qt9SortOrderE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &logicalIndex, &order)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:205
// index:0
// Protected
// void updateSection(int)
func (this *QHeaderView) UpdateSection(logicalIndex int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView13updateSectionEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &logicalIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:207
// index:0
// Protected
// void sectionsInserted(const class QModelIndex &, int, int)
func (this *QHeaderView) SectionsInserted(parent *qtcore.QModelIndex, logicalFirst int, logicalLast int) {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView16sectionsInsertedERK11QModelIndexii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &logicalFirst, &logicalLast)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:208
// index:0
// Protected
// void sectionsAboutToBeRemoved(const class QModelIndex &, int, int)
func (this *QHeaderView) SectionsAboutToBeRemoved(parent *qtcore.QModelIndex, logicalFirst int, logicalLast int) {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView24sectionsAboutToBeRemovedERK11QModelIndexii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &logicalFirst, &logicalLast)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:212
// index:0
// Protected
// void initialize()
func (this *QHeaderView) Initialize() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView10initializeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:214
// index:0
// Protected
// void initializeSections()
func (this *QHeaderView) InitializeSections() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView18initializeSectionsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:215
// index:1
// Protected
// void initializeSections(int, int)
func (this *QHeaderView) InitializeSections_1(start int, end int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView18initializeSectionsEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &start, &end)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:216
// index:0
// Protected virtual
// void currentChanged(const class QModelIndex &, const class QModelIndex &)
func (this *QHeaderView) CurrentChanged(current *qtcore.QModelIndex, old *qtcore.QModelIndex) {
	var convArg0 = current.GetCthis()
	var convArg1 = old.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView14currentChangedERK11QModelIndexS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:218
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QHeaderView) Event(e unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qheaderview.h:219
// index:0
// Protected virtual
// void paintEvent(class QPaintEvent *)
func (this *QHeaderView) PaintEvent(e unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:220
// index:0
// Protected virtual
// void mousePressEvent(class QMouseEvent *)
func (this *QHeaderView) MousePressEvent(e unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView15mousePressEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:221
// index:0
// Protected virtual
// void mouseMoveEvent(class QMouseEvent *)
func (this *QHeaderView) MouseMoveEvent(e unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView14mouseMoveEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:222
// index:0
// Protected virtual
// void mouseReleaseEvent(class QMouseEvent *)
func (this *QHeaderView) MouseReleaseEvent(e unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView17mouseReleaseEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:223
// index:0
// Protected virtual
// void mouseDoubleClickEvent(class QMouseEvent *)
func (this *QHeaderView) MouseDoubleClickEvent(e unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView21mouseDoubleClickEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:224
// index:0
// Protected virtual
// bool viewportEvent(class QEvent *)
func (this *QHeaderView) ViewportEvent(e unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView13viewportEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qheaderview.h:226
// index:0
// Protected virtual
// void paintSection(class QPainter *, const class QRect &, int)
func (this *QHeaderView) PaintSection(painter unsafe.Pointer, rect *qtcore.QRect, logicalIndex int) {
	var convArg1 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView12paintSectionEP8QPainterRK5QRecti", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), painter, convArg1, &logicalIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:227
// index:0
// Protected virtual
// QSize sectionSizeFromContents(int)
func (this *QHeaderView) SectionSizeFromContents(logicalIndex int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView23sectionSizeFromContentsEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &logicalIndex)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qheaderview.h:229
// index:0
// Protected virtual
// int horizontalOffset()
func (this *QHeaderView) HorizontalOffset() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView16horizontalOffsetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qheaderview.h:230
// index:0
// Protected virtual
// int verticalOffset()
func (this *QHeaderView) VerticalOffset() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView14verticalOffsetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qheaderview.h:231
// index:0
// Protected virtual
// void updateGeometries()
func (this *QHeaderView) UpdateGeometries() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView16updateGeometriesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:232
// index:0
// Protected virtual
// void scrollContentsBy(int, int)
func (this *QHeaderView) ScrollContentsBy(dx int, dy int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView16scrollContentsByEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &dx, &dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:235
// index:0
// Protected virtual
// void rowsInserted(const class QModelIndex &, int, int)
func (this *QHeaderView) RowsInserted(parent *qtcore.QModelIndex, start int, end int) {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView12rowsInsertedERK11QModelIndexii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &start, &end)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:237
// index:0
// Protected virtual
// QRect visualRect(const class QModelIndex &)
func (this *QHeaderView) VisualRect(index *qtcore.QModelIndex) interface{} {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView10visualRectERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qheaderview.h:238
// index:0
// Protected virtual
// void scrollTo(const class QModelIndex &, enum QAbstractItemView::ScrollHint)
func (this *QHeaderView) ScrollTo(index *qtcore.QModelIndex, hint int) {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHeaderView8scrollToERK11QModelIndexN17QAbstractItemView10ScrollHintE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &hint)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qheaderview.h:240
// index:0
// Protected virtual
// QModelIndex indexAt(const class QPoint &)
func (this *QHeaderView) IndexAt(p *qtcore.QPoint) interface{} {
	var convArg0 = p.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView7indexAtERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qheaderview.h:241
// index:0
// Protected virtual
// bool isIndexHidden(const class QModelIndex &)
func (this *QHeaderView) IsIndexHidden(index *qtcore.QModelIndex) interface{} {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView13isIndexHiddenERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qheaderview.h:245
// index:0
// Protected virtual
// QRegion visualRegionForSelection(const class QItemSelection &)
func (this *QHeaderView) VisualRegionForSelection(selection *qtcore.QItemSelection) interface{} {
	var convArg0 = selection.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView24visualRegionForSelectionERK14QItemSelection", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qheaderview.h:246
// index:0
// Protected
// void initStyleOption(class QStyleOptionHeader *)
func (this *QHeaderView) InitStyleOption(option unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHeaderView15initStyleOptionEP18QStyleOptionHeader", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), option)
	gopp.ErrPrint(err, rv)
}

//  body block end
