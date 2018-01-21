package qtwidgets

// /usr/include/qt/QtWidgets/qtablewidget.h
// #include <qtablewidget.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 9
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
type QTableWidgetItem struct {
	*qtrt.CObject
}

func (this *QTableWidgetItem) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func NewQTableWidgetItemFromPointer(cthis unsafe.Pointer) *QTableWidgetItem {
	return &QTableWidgetItem{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtWidgets/qtablewidget.h:82
// index:0
// Public
// void QTableWidgetItem(int)
func NewQTableWidgetItem(type_ int) *QTableWidgetItem {
	cthis := qtrt.Calloc(1, 256) // 48
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTableWidgetItemC2Ei", ffiqt.FFI_TYPE_VOID, cthis, &type_)
	gopp.ErrPrint(err, rv)
	gothis := NewQTableWidgetItemFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qtablewidget.h:83
// index:1
// Public
// void QTableWidgetItem(const class QString &, int)
func NewQTableWidgetItem_1(text *qtcore.QString, type_ int) *QTableWidgetItem {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTableWidgetItemC2ERK7QStringi", ffiqt.FFI_TYPE_VOID, cthis, convArg0, &type_)
	gopp.ErrPrint(err, rv)
	gothis := NewQTableWidgetItemFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qtablewidget.h:84
// index:2
// Public
// void QTableWidgetItem(const class QIcon &, const class QString &, int)
func NewQTableWidgetItem_2(icon *qtgui.QIcon, text *qtcore.QString, type_ int) *QTableWidgetItem {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = icon.GetCthis()
	var convArg1 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTableWidgetItemC2ERK5QIconRK7QStringi", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1, &type_)
	gopp.ErrPrint(err, rv)
	gothis := NewQTableWidgetItemFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qtablewidget.h:86
// index:0
// Public virtual
// void ~QTableWidgetItem()
func DeleteQTableWidgetItem(*QTableWidgetItem) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTableWidgetItemD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:88
// index:0
// Public virtual
// QTableWidgetItem * clone()
func (this *QTableWidgetItem) Clone() *QTableWidgetItem /*444 QTableWidgetItem **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTableWidgetItem5cloneEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTableWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:90
// index:0
// Public inline
// QTableWidget * tableWidget()
func (this *QTableWidgetItem) TableWidget() *QTableWidget /*444 QTableWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTableWidgetItem11tableWidgetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTableWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:92
// index:0
// Public inline
// int row()
func (this *QTableWidgetItem) Row() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTableWidgetItem3rowEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qtablewidget.h:93
// index:0
// Public inline
// int column()
func (this *QTableWidgetItem) Column() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTableWidgetItem6columnEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qtablewidget.h:95
// index:0
// Public inline
// void setSelected(_Bool)
func (this *QTableWidgetItem) SetSelected(select_ bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTableWidgetItem11setSelectedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &select_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:96
// index:0
// Public inline
// bool isSelected()
func (this *QTableWidgetItem) IsSelected() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTableWidgetItem10isSelectedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtablewidget.h:98
// index:0
// Public inline
// Qt::ItemFlags flags()
func (this *QTableWidgetItem) Flags() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTableWidgetItem5flagsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:101
// index:0
// Public inline
// QString text()
func (this *QTableWidgetItem) Text() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTableWidgetItem4textEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:103
// index:0
// Public inline
// void setText(const class QString &)
func (this *QTableWidgetItem) SetText(text *qtcore.QString) {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTableWidgetItem7setTextERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:105
// index:0
// Public inline
// QIcon icon()
func (this *QTableWidgetItem) Icon() *qtgui.QIcon /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTableWidgetItem4iconEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:107
// index:0
// Public inline
// void setIcon(const class QIcon &)
func (this *QTableWidgetItem) SetIcon(icon *qtgui.QIcon) {
	var convArg0 = icon.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTableWidgetItem7setIconERK5QIcon", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:109
// index:0
// Public inline
// QString statusTip()
func (this *QTableWidgetItem) StatusTip() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTableWidgetItem9statusTipEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:111
// index:0
// Public inline
// void setStatusTip(const class QString &)
func (this *QTableWidgetItem) SetStatusTip(statusTip *qtcore.QString) {
	var convArg0 = statusTip.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTableWidgetItem12setStatusTipERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:114
// index:0
// Public inline
// QString toolTip()
func (this *QTableWidgetItem) ToolTip() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTableWidgetItem7toolTipEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:116
// index:0
// Public inline
// void setToolTip(const class QString &)
func (this *QTableWidgetItem) SetToolTip(toolTip *qtcore.QString) {
	var convArg0 = toolTip.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTableWidgetItem10setToolTipERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:120
// index:0
// Public inline
// QString whatsThis()
func (this *QTableWidgetItem) WhatsThis() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTableWidgetItem9whatsThisEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:122
// index:0
// Public inline
// void setWhatsThis(const class QString &)
func (this *QTableWidgetItem) SetWhatsThis(whatsThis *qtcore.QString) {
	var convArg0 = whatsThis.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTableWidgetItem12setWhatsThisERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:125
// index:0
// Public inline
// QFont font()
func (this *QTableWidgetItem) Font() *qtgui.QFont /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTableWidgetItem4fontEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:127
// index:0
// Public inline
// void setFont(const class QFont &)
func (this *QTableWidgetItem) SetFont(font *qtgui.QFont) {
	var convArg0 = font.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTableWidgetItem7setFontERK5QFont", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:129
// index:0
// Public inline
// int textAlignment()
func (this *QTableWidgetItem) TextAlignment() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTableWidgetItem13textAlignmentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qtablewidget.h:131
// index:0
// Public inline
// void setTextAlignment(int)
func (this *QTableWidgetItem) SetTextAlignment(alignment int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTableWidgetItem16setTextAlignmentEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &alignment)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:134
// index:0
// Public inline
// QColor backgroundColor()
func (this *QTableWidgetItem) BackgroundColor() *qtgui.QColor /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTableWidgetItem15backgroundColorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:136
// index:0
// Public inline
// void setBackgroundColor(const class QColor &)
func (this *QTableWidgetItem) SetBackgroundColor(color *qtgui.QColor) {
	var convArg0 = color.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTableWidgetItem18setBackgroundColorERK6QColor", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:139
// index:0
// Public inline
// QBrush background()
func (this *QTableWidgetItem) Background() *qtgui.QBrush /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTableWidgetItem10backgroundEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:141
// index:0
// Public inline
// void setBackground(const class QBrush &)
func (this *QTableWidgetItem) SetBackground(brush *qtgui.QBrush) {
	var convArg0 = brush.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTableWidgetItem13setBackgroundERK6QBrush", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:144
// index:0
// Public inline
// QColor textColor()
func (this *QTableWidgetItem) TextColor() *qtgui.QColor /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTableWidgetItem9textColorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:146
// index:0
// Public inline
// void setTextColor(const class QColor &)
func (this *QTableWidgetItem) SetTextColor(color *qtgui.QColor) {
	var convArg0 = color.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTableWidgetItem12setTextColorERK6QColor", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:149
// index:0
// Public inline
// QBrush foreground()
func (this *QTableWidgetItem) Foreground() *qtgui.QBrush /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTableWidgetItem10foregroundEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:151
// index:0
// Public inline
// void setForeground(const class QBrush &)
func (this *QTableWidgetItem) SetForeground(brush *qtgui.QBrush) {
	var convArg0 = brush.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTableWidgetItem13setForegroundERK6QBrush", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:154
// index:0
// Public inline
// Qt::CheckState checkState()
func (this *QTableWidgetItem) CheckState() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTableWidgetItem10checkStateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:156
// index:0
// Public inline
// void setCheckState(Qt::CheckState)
func (this *QTableWidgetItem) SetCheckState(state int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTableWidgetItem13setCheckStateEN2Qt10CheckStateE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:159
// index:0
// Public inline
// QSize sizeHint()
func (this *QTableWidgetItem) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTableWidgetItem8sizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:161
// index:0
// Public inline
// void setSizeHint(const class QSize &)
func (this *QTableWidgetItem) SetSizeHint(size *qtcore.QSize) {
	var convArg0 = size.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTableWidgetItem11setSizeHintERK5QSize", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:164
// index:0
// Public virtual
// QVariant data(int)
func (this *QTableWidgetItem) Data(role int) *qtcore.QVariant /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTableWidgetItem4dataEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &role)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:165
// index:0
// Public virtual
// void setData(int, const class QVariant &)
func (this *QTableWidgetItem) SetData(role int, value *qtcore.QVariant) {
	var convArg1 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTableWidgetItem7setDataEiRK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &role, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:170
// index:0
// Public virtual
// void read(class QDataStream &)
func (this *QTableWidgetItem) Read(in *qtcore.QDataStream) {
	var convArg0 = in.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTableWidgetItem4readER11QDataStream", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:171
// index:0
// Public virtual
// void write(class QDataStream &)
func (this *QTableWidgetItem) Write(out *qtcore.QDataStream) {
	var convArg0 = out.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTableWidgetItem5writeER11QDataStream", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:175
// index:0
// Public inline
// int type()
func (this *QTableWidgetItem) Type() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTableWidgetItem4typeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

//  body block end
