package qtwidgets

// /usr/include/qt/QtWidgets/qitemdelegate.h
// #include <qitemdelegate.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 57
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

// void drawDisplay(QPainter *, const QStyleOptionViewItem &, const QRect &, const QString &)
func (this *QItemDelegate) InheritDrawDisplay(f func(painter *qtgui.QPainter /*777 QPainter **/, option *QStyleOptionViewItem, rect *qtcore.QRect, text string) /*void*/) {
	qtrt.SetAllInheritCallback(this, "drawDisplay", f)
}

// void drawDecoration(QPainter *, const QStyleOptionViewItem &, const QRect &, const QPixmap &)
func (this *QItemDelegate) InheritDrawDecoration(f func(painter *qtgui.QPainter /*777 QPainter **/, option *QStyleOptionViewItem, rect *qtcore.QRect, pixmap *qtgui.QPixmap) /*void*/) {
	qtrt.SetAllInheritCallback(this, "drawDecoration", f)
}

// void drawFocus(QPainter *, const QStyleOptionViewItem &, const QRect &)
func (this *QItemDelegate) InheritDrawFocus(f func(painter *qtgui.QPainter /*777 QPainter **/, option *QStyleOptionViewItem, rect *qtcore.QRect) /*void*/) {
	qtrt.SetAllInheritCallback(this, "drawFocus", f)
}

// void drawCheck(QPainter *, const QStyleOptionViewItem &, const QRect &, Qt::CheckState)
func (this *QItemDelegate) InheritDrawCheck(f func(painter *qtgui.QPainter /*777 QPainter **/, option *QStyleOptionViewItem, rect *qtcore.QRect, state int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "drawCheck", f)
}

// void drawBackground(QPainter *, const QStyleOptionViewItem &, const QModelIndex &)
func (this *QItemDelegate) InheritDrawBackground(f func(painter *qtgui.QPainter /*777 QPainter **/, option *QStyleOptionViewItem, index *qtcore.QModelIndex) /*void*/) {
	qtrt.SetAllInheritCallback(this, "drawBackground", f)
}

// void doLayout(const QStyleOptionViewItem &, QRect *, QRect *, QRect *, bool)
func (this *QItemDelegate) InheritDoLayout(f func(option *QStyleOptionViewItem, checkRect *qtcore.QRect /*777 QRect **/, iconRect *qtcore.QRect /*777 QRect **/, textRect *qtcore.QRect /*777 QRect **/, hint bool) /*void*/) {
	qtrt.SetAllInheritCallback(this, "doLayout", f)
}

// QRect rect(const QStyleOptionViewItem &, const QModelIndex &, int)
func (this *QItemDelegate) InheritRect(f func(option *QStyleOptionViewItem, index *qtcore.QModelIndex, role int) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "rect", f)
}

// bool eventFilter(QObject *, QEvent *)
func (this *QItemDelegate) InheritEventFilter(f func(object *qtcore.QObject /*777 QObject **/, event *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "eventFilter", f)
}

// bool editorEvent(QEvent *, QAbstractItemModel *, const QStyleOptionViewItem &, const QModelIndex &)
func (this *QItemDelegate) InheritEditorEvent(f func(event *qtcore.QEvent /*777 QEvent **/, model *qtcore.QAbstractItemModel /*777 QAbstractItemModel **/, option *QStyleOptionViewItem, index *qtcore.QModelIndex) bool) {
	qtrt.SetAllInheritCallback(this, "editorEvent", f)
}

// QStyleOptionViewItem setOptions(const QModelIndex &, const QStyleOptionViewItem &)
func (this *QItemDelegate) InheritSetOptions(f func(index *qtcore.QModelIndex, option *QStyleOptionViewItem) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "setOptions", f)
}

// QPixmap decoration(const QStyleOptionViewItem &, const QVariant &)
func (this *QItemDelegate) InheritDecoration(f func(option *QStyleOptionViewItem, variant *qtcore.QVariant) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "decoration", f)
}

// QPixmap * selected(const QPixmap &, const QPalette &, bool)
func (this *QItemDelegate) InheritSelected(f func(pixmap *qtgui.QPixmap, palette *qtgui.QPalette, enabled bool) unsafe.Pointer /*666*/) {
	qtrt.SetAllInheritCallback(this, "selected", f)
}

// QRect doCheck(const QStyleOptionViewItem &, const QRect &, const QVariant &)
func (this *QItemDelegate) InheritDoCheck(f func(option *QStyleOptionViewItem, bounding *qtcore.QRect, variant *qtcore.QVariant) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "doCheck", f)
}

// QRect textRectangle(QPainter *, const QRect &, const QFont &, const QString &)
func (this *QItemDelegate) InheritTextRectangle(f func(painter *qtgui.QPainter /*777 QPainter **/, rect *qtcore.QRect, font *qtgui.QFont, text string) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "textRectangle", f)
}

/*

 */
type QItemDelegate struct {
	*QAbstractItemDelegate
}
type QItemDelegate_ITF interface {
	QAbstractItemDelegate_ITF
	QItemDelegate_PTR() *QItemDelegate
}

func (ptr *QItemDelegate) QItemDelegate_PTR() *QItemDelegate { return ptr }

func (this *QItemDelegate) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAbstractItemDelegate.GetCthis()
	}
}
func (this *QItemDelegate) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractItemDelegate = NewQAbstractItemDelegateFromPointer(cthis)
}
func NewQItemDelegateFromPointer(cthis unsafe.Pointer) *QItemDelegate {
	bcthis0 := NewQAbstractItemDelegateFromPointer(cthis)
	return &QItemDelegate{bcthis0}
}
func (*QItemDelegate) NewFromPointer(cthis unsafe.Pointer) *QItemDelegate {
	return NewQItemDelegateFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:59
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QItemDelegate) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QItemDelegate10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:63
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QItemDelegate(QObject *)

/*
Constructs an item delegate with the given parent.
*/
func (*QItemDelegate) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QItemDelegate {
	return NewQItemDelegate(parent)
}
func NewQItemDelegate(parent qtcore.QObject_ITF /*777 QObject **/) *QItemDelegate {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QItemDelegateC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQItemDelegateFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QItemDelegate")
	return gothis
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:63
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QItemDelegate(QObject *)

/*
Constructs an item delegate with the given parent.
*/
func (*QItemDelegate) NewForInheritp() *QItemDelegate {
	return NewQItemDelegatep()
}
func NewQItemDelegatep() *QItemDelegate {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN13QItemDelegateC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQItemDelegateFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QItemDelegate")
	return gothis
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:64
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QItemDelegate()

/*

 */
func DeleteQItemDelegate(this *QItemDelegate) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QItemDelegateD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:66
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasClipping() const

/*

 */
func (this *QItemDelegate) HasClipping() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QItemDelegate11hasClippingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:67
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setClipping(bool)

/*

 */
func (this *QItemDelegate) SetClipping(clip bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QItemDelegate11setClippingEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), clip)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:70
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void paint(QPainter *, const QStyleOptionViewItem &, const QModelIndex &) const

/*
Reimplemented from QAbstractItemDelegate::paint().

Renders the delegate using the given painter and style option for the item specified by index.

When reimplementing this function in a subclass, you should update the area held by the option's rect variable, using the option's state variable to determine the state of the item to be displayed, and adjust the way it is painted accordingly.

For example, a selected item may need to be displayed differently to unselected items, as shown in the following code:


      if (option.state & QStyle::State_Selected)
          painter->fillRect(option.rect, option.palette.highlight());

      int size = qMin(option.rect.width(), option.rect.height());
      int brightness = index.model()->data(index, Qt::DisplayRole).toInt();
      double radius = (size / 2.0) - (brightness / 255.0 * size / 2.0);
      if (radius == 0.0)
          return;

      painter->save();
      painter->setRenderHint(QPainter::Antialiasing, true);
      painter->setPen(Qt::NoPen);
      if (option.state & QStyle::State_Selected)
          painter->setBrush(option.palette.highlightedText());
      else
      ...



After painting, you should ensure that the painter is returned to its the state it was supplied in when this function was called. For example, it may be useful to call QPainter::save() before painting and QPainter::restore() afterwards.

See also QStyle::State.
*/
func (this *QItemDelegate) Paint(painter qtgui.QPainter_ITF /*777 QPainter **/, option QStyleOptionViewItem_ITF, index qtcore.QModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if option != nil && option.QStyleOptionViewItem_PTR() != nil {
		convArg1 = option.QStyleOptionViewItem_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg2 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QItemDelegate5paintEP8QPainterRK20QStyleOptionViewItemRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:73
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint(const QStyleOptionViewItem &, const QModelIndex &) const

/*
Reimplemented from QAbstractItemDelegate::sizeHint().

Returns the size needed by the delegate to display the item specified by index, taking into account the style information provided by option.

When reimplementing this function, note that in case of text items, QItemDelegate adds a margin (i.e. 2 * QStyle::PM_FocusFrameHMargin) to the length of the text.
*/
func (this *QItemDelegate) SizeHint(option QStyleOptionViewItem_ITF, index qtcore.QModelIndex_ITF) *qtcore.QSize /*123*/ {
	var convArg0 unsafe.Pointer
	if option != nil && option.QStyleOptionViewItem_PTR() != nil {
		convArg0 = option.QStyleOptionViewItem_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg1 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QItemDelegate8sizeHintERK20QStyleOptionViewItemRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:77
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QWidget * createEditor(QWidget *, const QStyleOptionViewItem &, const QModelIndex &) const

/*
Reimplemented from QAbstractItemDelegate::createEditor().

Returns the widget used to edit the item specified by index for editing. The parent widget and style option are used to control how the editor widget appears.

See also QAbstractItemDelegate::createEditor().
*/
func (this *QItemDelegate) CreateEditor(parent QWidget_ITF /*777 QWidget **/, option QStyleOptionViewItem_ITF, index qtcore.QModelIndex_ITF) *QWidget /*777 QWidget **/ {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if option != nil && option.QStyleOptionViewItem_PTR() != nil {
		convArg1 = option.QStyleOptionViewItem_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg2 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QItemDelegate12createEditorEP7QWidgetRK20QStyleOptionViewItemRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:81
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setEditorData(QWidget *, const QModelIndex &) const

/*
Reimplemented from QAbstractItemDelegate::setEditorData().

Sets the data to be displayed and edited by the editor from the data model item specified by the model index.

The default implementation stores the data in the editor widget's user property.

See also QMetaProperty::isUser().
*/
func (this *QItemDelegate) SetEditorData(editor QWidget_ITF /*777 QWidget **/, index qtcore.QModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if editor != nil && editor.QWidget_PTR() != nil {
		convArg0 = editor.QWidget_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg1 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QItemDelegate13setEditorDataEP7QWidgetRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:82
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setModelData(QWidget *, QAbstractItemModel *, const QModelIndex &) const

/*
Reimplemented from QAbstractItemDelegate::setModelData().

Gets data from the editor widget and stores it in the specified model at the item index.

The default implementation gets the value to be stored in the data model from the editor widget's user property.

See also QMetaProperty::isUser().
*/
func (this *QItemDelegate) SetModelData(editor QWidget_ITF /*777 QWidget **/, model qtcore.QAbstractItemModel_ITF /*777 QAbstractItemModel **/, index qtcore.QModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if editor != nil && editor.QWidget_PTR() != nil {
		convArg0 = editor.QWidget_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if model != nil && model.QAbstractItemModel_PTR() != nil {
		convArg1 = model.QAbstractItemModel_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg2 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QItemDelegate12setModelDataEP7QWidgetP18QAbstractItemModelRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:84
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void updateEditorGeometry(QWidget *, const QStyleOptionViewItem &, const QModelIndex &) const

/*
Reimplemented from QAbstractItemDelegate::updateEditorGeometry().

Updates the editor for the item specified by index according to the style option given.
*/
func (this *QItemDelegate) UpdateEditorGeometry(editor QWidget_ITF /*777 QWidget **/, option QStyleOptionViewItem_ITF, index qtcore.QModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if editor != nil && editor.QWidget_PTR() != nil {
		convArg0 = editor.QWidget_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if option != nil && option.QStyleOptionViewItem_PTR() != nil {
		convArg1 = option.QStyleOptionViewItem_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg2 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QItemDelegate20updateEditorGeometryEP7QWidgetRK20QStyleOptionViewItemRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:89
// index:0
// Public Visibility=Default Availability=Available
// [8] QItemEditorFactory * itemEditorFactory() const

/*
Returns the editor factory used by the item delegate. If no editor factory is set, the function will return null.

See also setItemEditorFactory().
*/
func (this *QItemDelegate) ItemEditorFactory() *QItemEditorFactory /*777 QItemEditorFactory **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QItemDelegate17itemEditorFactoryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQItemEditorFactoryFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setItemEditorFactory(QItemEditorFactory *)

/*
Sets the editor factory to be used by the item delegate to be the factory specified. If no editor factory is set, the item delegate will use the default editor factory.

See also itemEditorFactory().
*/
func (this *QItemDelegate) SetItemEditorFactory(factory QItemEditorFactory_ITF /*777 QItemEditorFactory **/) {
	var convArg0 unsafe.Pointer
	if factory != nil && factory.QItemEditorFactory_PTR() != nil {
		convArg0 = factory.QItemEditorFactory_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QItemDelegate20setItemEditorFactoryEP18QItemEditorFactory", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:93
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void drawDisplay(QPainter *, const QStyleOptionViewItem &, const QRect &, const QString &) const

/*
Renders the item view text within the rectangle specified by rect using the given painter and style option.
*/
func (this *QItemDelegate) DrawDisplay(painter qtgui.QPainter_ITF /*777 QPainter **/, option QStyleOptionViewItem_ITF, rect qtcore.QRect_ITF, text string) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if option != nil && option.QStyleOptionViewItem_PTR() != nil {
		convArg1 = option.QStyleOptionViewItem_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if rect != nil && rect.QRect_PTR() != nil {
		convArg2 = rect.QRect_PTR().GetCthis()
	}
	var tmpArg3 = qtcore.NewQString5(text)
	var convArg3 = tmpArg3.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QItemDelegate11drawDisplayEP8QPainterRK20QStyleOptionViewItemRK5QRectRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:95
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void drawDecoration(QPainter *, const QStyleOptionViewItem &, const QRect &, const QPixmap &) const

/*
Renders the decoration pixmap within the rectangle specified by rect using the given painter and style option.
*/
func (this *QItemDelegate) DrawDecoration(painter qtgui.QPainter_ITF /*777 QPainter **/, option QStyleOptionViewItem_ITF, rect qtcore.QRect_ITF, pixmap qtgui.QPixmap_ITF) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if option != nil && option.QStyleOptionViewItem_PTR() != nil {
		convArg1 = option.QStyleOptionViewItem_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if rect != nil && rect.QRect_PTR() != nil {
		convArg2 = rect.QRect_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if pixmap != nil && pixmap.QPixmap_PTR() != nil {
		convArg3 = pixmap.QPixmap_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QItemDelegate14drawDecorationEP8QPainterRK20QStyleOptionViewItemRK5QRectRK7QPixmap", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:97
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void drawFocus(QPainter *, const QStyleOptionViewItem &, const QRect &) const

/*
Renders the region within the rectangle specified by rect, indicating that it has the focus, using the given painter and style option.
*/
func (this *QItemDelegate) DrawFocus(painter qtgui.QPainter_ITF /*777 QPainter **/, option QStyleOptionViewItem_ITF, rect qtcore.QRect_ITF) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if option != nil && option.QStyleOptionViewItem_PTR() != nil {
		convArg1 = option.QStyleOptionViewItem_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if rect != nil && rect.QRect_PTR() != nil {
		convArg2 = rect.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QItemDelegate9drawFocusEP8QPainterRK20QStyleOptionViewItemRK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:99
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void drawCheck(QPainter *, const QStyleOptionViewItem &, const QRect &, Qt::CheckState) const

/*
Renders a check indicator within the rectangle specified by rect, using the given painter and style option, using the given state.
*/
func (this *QItemDelegate) DrawCheck(painter qtgui.QPainter_ITF /*777 QPainter **/, option QStyleOptionViewItem_ITF, rect qtcore.QRect_ITF, state int) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if option != nil && option.QStyleOptionViewItem_PTR() != nil {
		convArg1 = option.QStyleOptionViewItem_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if rect != nil && rect.QRect_PTR() != nil {
		convArg2 = rect.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QItemDelegate9drawCheckEP8QPainterRK20QStyleOptionViewItemRK5QRectN2Qt10CheckStateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, state)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:101
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void drawBackground(QPainter *, const QStyleOptionViewItem &, const QModelIndex &) const

/*
Renders the item background for the given index, using the given painter and style option.

This function was introduced in  Qt 4.2.
*/
func (this *QItemDelegate) DrawBackground(painter qtgui.QPainter_ITF /*777 QPainter **/, option QStyleOptionViewItem_ITF, index qtcore.QModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if option != nil && option.QStyleOptionViewItem_PTR() != nil {
		convArg1 = option.QStyleOptionViewItem_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg2 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QItemDelegate14drawBackgroundEP8QPainterRK20QStyleOptionViewItemRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:104
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void doLayout(const QStyleOptionViewItem &, QRect *, QRect *, QRect *, bool) const

/*

 */
func (this *QItemDelegate) DoLayout(option QStyleOptionViewItem_ITF, checkRect qtcore.QRect_ITF /*777 QRect **/, iconRect qtcore.QRect_ITF /*777 QRect **/, textRect qtcore.QRect_ITF /*777 QRect **/, hint bool) {
	var convArg0 unsafe.Pointer
	if option != nil && option.QStyleOptionViewItem_PTR() != nil {
		convArg0 = option.QStyleOptionViewItem_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if checkRect != nil && checkRect.QRect_PTR() != nil {
		convArg1 = checkRect.QRect_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if iconRect != nil && iconRect.QRect_PTR() != nil {
		convArg2 = iconRect.QRect_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if textRect != nil && textRect.QRect_PTR() != nil {
		convArg3 = textRect.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QItemDelegate8doLayoutERK20QStyleOptionViewItemP5QRectS4_S4_b", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3, hint)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:107
// index:0
// Protected Visibility=Default Availability=Available
// [16] QRect rect(const QStyleOptionViewItem &, const QModelIndex &, int) const

/*

 */
func (this *QItemDelegate) Rect(option QStyleOptionViewItem_ITF, index qtcore.QModelIndex_ITF, role int) *qtcore.QRect /*123*/ {
	var convArg0 unsafe.Pointer
	if option != nil && option.QStyleOptionViewItem_PTR() != nil {
		convArg0 = option.QStyleOptionViewItem_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg1 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QItemDelegate4rectERK20QStyleOptionViewItemRK11QModelIndexi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, role)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:109
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool eventFilter(QObject *, QEvent *)

/*
Reimplemented from QObject::eventFilter().

Returns true if the given editor is a valid QWidget and the given event is handled; otherwise returns false. The following key press events are handled by default:


Tab
Backtab
Enter
Return
Esc


In the case of Tab, Backtab, Enter and Return key press events, the editor's data is comitted to the model and the editor is closed. If the event is a Tab key press the view will open an editor on the next item in the view. Likewise, if the event is a Backtab key press the view will open an editor on the previous item in the view.

If the event is a Esc key press event, the editor is closed without committing its data.

See also commitData() and closeEditor().
*/
func (this *QItemDelegate) EventFilter(object qtcore.QObject_ITF /*777 QObject **/, event qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg1 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QItemDelegate11eventFilterEP7QObjectP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:110
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool editorEvent(QEvent *, QAbstractItemModel *, const QStyleOptionViewItem &, const QModelIndex &)

/*
Reimplemented from QAbstractItemDelegate::editorEvent().
*/
func (this *QItemDelegate) EditorEvent(event qtcore.QEvent_ITF /*777 QEvent **/, model qtcore.QAbstractItemModel_ITF /*777 QAbstractItemModel **/, option QStyleOptionViewItem_ITF, index qtcore.QModelIndex_ITF) bool {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if model != nil && model.QAbstractItemModel_PTR() != nil {
		convArg1 = model.QAbstractItemModel_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if option != nil && option.QStyleOptionViewItem_PTR() != nil {
		convArg2 = option.QStyleOptionViewItem_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg3 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QItemDelegate11editorEventEP6QEventP18QAbstractItemModelRK20QStyleOptionViewItemRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:113
// index:0
// Protected Visibility=Default Availability=Available
// [192] QStyleOptionViewItem setOptions(const QModelIndex &, const QStyleOptionViewItem &) const

/*

 */
func (this *QItemDelegate) SetOptions(index qtcore.QModelIndex_ITF, option QStyleOptionViewItem_ITF) *QStyleOptionViewItem /*123*/ {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if option != nil && option.QStyleOptionViewItem_PTR() != nil {
		convArg1 = option.QStyleOptionViewItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QItemDelegate10setOptionsERK11QModelIndexRK20QStyleOptionViewItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStyleOptionViewItemFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStyleOptionViewItem)
	return rv2
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:116
// index:0
// Protected Visibility=Default Availability=Available
// [32] QPixmap decoration(const QStyleOptionViewItem &, const QVariant &) const

/*

 */
func (this *QItemDelegate) Decoration(option QStyleOptionViewItem_ITF, variant qtcore.QVariant_ITF) *qtgui.QPixmap /*123*/ {
	var convArg0 unsafe.Pointer
	if option != nil && option.QStyleOptionViewItem_PTR() != nil {
		convArg0 = option.QStyleOptionViewItem_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if variant != nil && variant.QVariant_PTR() != nil {
		convArg1 = variant.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QItemDelegate10decorationERK20QStyleOptionViewItemRK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:117
// index:0
// Protected Visibility=Default Availability=Available
// [8] QPixmap * selected(const QPixmap &, const QPalette &, bool) const

/*

 */
func (this *QItemDelegate) Selected(pixmap qtgui.QPixmap_ITF, palette qtgui.QPalette_ITF, enabled bool) *qtgui.QPixmap /*777 QPixmap **/ {
	var convArg0 unsafe.Pointer
	if pixmap != nil && pixmap.QPixmap_PTR() != nil {
		convArg0 = pixmap.QPixmap_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if palette != nil && palette.QPalette_PTR() != nil {
		convArg1 = palette.QPalette_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QItemDelegate8selectedERK7QPixmapRK8QPaletteb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, enabled)
	qtrt.ErrPrint(err, rv)
	return qtgui.NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:119
// index:0
// Protected Visibility=Default Availability=Available
// [16] QRect doCheck(const QStyleOptionViewItem &, const QRect &, const QVariant &) const

/*

 */
func (this *QItemDelegate) DoCheck(option QStyleOptionViewItem_ITF, bounding qtcore.QRect_ITF, variant qtcore.QVariant_ITF) *qtcore.QRect /*123*/ {
	var convArg0 unsafe.Pointer
	if option != nil && option.QStyleOptionViewItem_PTR() != nil {
		convArg0 = option.QStyleOptionViewItem_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if bounding != nil && bounding.QRect_PTR() != nil {
		convArg1 = bounding.QRect_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if variant != nil && variant.QVariant_PTR() != nil {
		convArg2 = variant.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QItemDelegate7doCheckERK20QStyleOptionViewItemRK5QRectRK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qitemdelegate.h:121
// index:0
// Protected Visibility=Default Availability=Available
// [16] QRect textRectangle(QPainter *, const QRect &, const QFont &, const QString &) const

/*

 */
func (this *QItemDelegate) TextRectangle(painter qtgui.QPainter_ITF /*777 QPainter **/, rect qtcore.QRect_ITF, font qtgui.QFont_ITF, text string) *qtcore.QRect /*123*/ {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if rect != nil && rect.QRect_PTR() != nil {
		convArg1 = rect.QRect_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if font != nil && font.QFont_PTR() != nil {
		convArg2 = font.QFont_PTR().GetCthis()
	}
	var tmpArg3 = qtcore.NewQString5(text)
	var convArg3 = tmpArg3.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QItemDelegate13textRectangleEP8QPainterRK5QRectRK5QFontRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
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
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end
