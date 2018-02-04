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
func (this *QTableWidgetItem) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQTableWidgetItemFromPointer(cthis unsafe.Pointer) *QTableWidgetItem {
	return &QTableWidgetItem{&qtrt.CObject{cthis}}
}
func (*QTableWidgetItem) NewFromPointer(cthis unsafe.Pointer) *QTableWidgetItem {
	return NewQTableWidgetItemFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTableWidgetItem(int)
func NewQTableWidgetItem(type_ int) *QTableWidgetItem {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTableWidgetItemC2Ei", ffiqt.FFI_TYPE_POINTER, type_)
	gopp.ErrPrint(err, rv)
	gothis := NewQTableWidgetItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTableWidgetItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qtablewidget.h:83
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QTableWidgetItem(const QString &, int)
func NewQTableWidgetItem_1(text *qtcore.QString, type_ int) *QTableWidgetItem {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTableWidgetItemC2ERK7QStringi", ffiqt.FFI_TYPE_POINTER, convArg0, type_)
	gopp.ErrPrint(err, rv)
	gothis := NewQTableWidgetItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTableWidgetItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qtablewidget.h:84
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QTableWidgetItem(const QIcon &, const QString &, int)
func NewQTableWidgetItem_2(icon *qtgui.QIcon, text *qtcore.QString, type_ int) *QTableWidgetItem {
	var convArg0 = icon.GetCthis()
	var convArg1 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTableWidgetItemC2ERK5QIconRK7QStringi", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1, type_)
	gopp.ErrPrint(err, rv)
	gothis := NewQTableWidgetItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTableWidgetItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qtablewidget.h:86
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QTableWidgetItem()
func DeleteQTableWidgetItem(this *QTableWidgetItem) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTableWidgetItemD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:88
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QTableWidgetItem * clone()
func (this *QTableWidgetItem) Clone() *QTableWidgetItem /*777 QTableWidgetItem **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTableWidgetItem5cloneEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTableWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:90
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QTableWidget * tableWidget()
func (this *QTableWidgetItem) TableWidget() *QTableWidget /*777 QTableWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTableWidgetItem11tableWidgetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTableWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:92
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int row()
func (this *QTableWidgetItem) Row() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTableWidgetItem3rowEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtablewidget.h:93
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int column()
func (this *QTableWidgetItem) Column() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTableWidgetItem6columnEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtablewidget.h:95
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setSelected(_Bool)
func (this *QTableWidgetItem) SetSelected(select_ bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTableWidgetItem11setSelectedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), select_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:96
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isSelected()
func (this *QTableWidgetItem) IsSelected() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTableWidgetItem10isSelectedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtablewidget.h:98
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::ItemFlags flags()
func (this *QTableWidgetItem) Flags() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTableWidgetItem5flagsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:99
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFlags(Qt::ItemFlags)
func (this *QTableWidgetItem) SetFlags(flags int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTableWidgetItem8setFlagsE6QFlagsIN2Qt8ItemFlagEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:101
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString text()
func (this *QTableWidgetItem) Text() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTableWidgetItem4textEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:103
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setText(const QString &)
func (this *QTableWidgetItem) SetText(text *qtcore.QString) {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTableWidgetItem7setTextERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:105
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QIcon icon()
func (this *QTableWidgetItem) Icon() *qtgui.QIcon /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTableWidgetItem4iconEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQIcon)
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:107
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setIcon(const QIcon &)
func (this *QTableWidgetItem) SetIcon(icon *qtgui.QIcon) {
	var convArg0 = icon.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTableWidgetItem7setIconERK5QIcon", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:109
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString statusTip()
func (this *QTableWidgetItem) StatusTip() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTableWidgetItem9statusTipEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:111
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setStatusTip(const QString &)
func (this *QTableWidgetItem) SetStatusTip(statusTip *qtcore.QString) {
	var convArg0 = statusTip.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTableWidgetItem12setStatusTipERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:114
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString toolTip()
func (this *QTableWidgetItem) ToolTip() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTableWidgetItem7toolTipEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:116
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setToolTip(const QString &)
func (this *QTableWidgetItem) SetToolTip(toolTip *qtcore.QString) {
	var convArg0 = toolTip.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTableWidgetItem10setToolTipERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:120
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString whatsThis()
func (this *QTableWidgetItem) WhatsThis() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTableWidgetItem9whatsThisEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:122
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setWhatsThis(const QString &)
func (this *QTableWidgetItem) SetWhatsThis(whatsThis *qtcore.QString) {
	var convArg0 = whatsThis.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTableWidgetItem12setWhatsThisERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:125
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QFont font()
func (this *QTableWidgetItem) Font() *qtgui.QFont /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTableWidgetItem4fontEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQFont)
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:127
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setFont(const QFont &)
func (this *QTableWidgetItem) SetFont(font *qtgui.QFont) {
	var convArg0 = font.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTableWidgetItem7setFontERK5QFont", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:129
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int textAlignment()
func (this *QTableWidgetItem) TextAlignment() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTableWidgetItem13textAlignmentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtablewidget.h:131
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setTextAlignment(int)
func (this *QTableWidgetItem) SetTextAlignment(alignment int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTableWidgetItem16setTextAlignmentEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), alignment)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:134
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QColor backgroundColor()
func (this *QTableWidgetItem) BackgroundColor() *qtgui.QColor /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTableWidgetItem15backgroundColorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQColor)
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:136
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setBackgroundColor(const QColor &)
func (this *QTableWidgetItem) SetBackgroundColor(color *qtgui.QColor) {
	var convArg0 = color.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTableWidgetItem18setBackgroundColorERK6QColor", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:139
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QBrush background()
func (this *QTableWidgetItem) Background() *qtgui.QBrush /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTableWidgetItem10backgroundEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:141
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setBackground(const QBrush &)
func (this *QTableWidgetItem) SetBackground(brush *qtgui.QBrush) {
	var convArg0 = brush.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTableWidgetItem13setBackgroundERK6QBrush", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:144
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QColor textColor()
func (this *QTableWidgetItem) TextColor() *qtgui.QColor /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTableWidgetItem9textColorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQColor)
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:146
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setTextColor(const QColor &)
func (this *QTableWidgetItem) SetTextColor(color *qtgui.QColor) {
	var convArg0 = color.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTableWidgetItem12setTextColorERK6QColor", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:149
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QBrush foreground()
func (this *QTableWidgetItem) Foreground() *qtgui.QBrush /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTableWidgetItem10foregroundEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:151
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setForeground(const QBrush &)
func (this *QTableWidgetItem) SetForeground(brush *qtgui.QBrush) {
	var convArg0 = brush.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTableWidgetItem13setForegroundERK6QBrush", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:154
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::CheckState checkState()
func (this *QTableWidgetItem) CheckState() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTableWidgetItem10checkStateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:156
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setCheckState(Qt::CheckState)
func (this *QTableWidgetItem) SetCheckState(state int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTableWidgetItem13setCheckStateEN2Qt10CheckStateE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:159
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QSize sizeHint()
func (this *QTableWidgetItem) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTableWidgetItem8sizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:161
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setSizeHint(const QSize &)
func (this *QTableWidgetItem) SetSizeHint(size *qtcore.QSize) {
	var convArg0 = size.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTableWidgetItem11setSizeHintERK5QSize", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:164
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant data(int)
func (this *QTableWidgetItem) Data(role int) *qtcore.QVariant /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTableWidgetItem4dataEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), role)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtWidgets/qtablewidget.h:165
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setData(int, const QVariant &)
func (this *QTableWidgetItem) SetData(role int, value *qtcore.QVariant) {
	var convArg1 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTableWidgetItem7setDataEiRK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), role, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:170
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void read(QDataStream &)
func (this *QTableWidgetItem) Read(in *qtcore.QDataStream) {
	var convArg0 = in.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTableWidgetItem4readER11QDataStream", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:171
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void write(QDataStream &)
func (this *QTableWidgetItem) Write(out *qtcore.QDataStream) {
	var convArg0 = out.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTableWidgetItem5writeER11QDataStream", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:175
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int type()
func (this *QTableWidgetItem) Type() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTableWidgetItem4typeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

type QTableWidgetItem__ItemType = int

const QTableWidgetItem__Type QTableWidgetItem__ItemType = 0
const QTableWidgetItem__UserType QTableWidgetItem__ItemType = 1000

//  body block end
