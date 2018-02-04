package qtgui

// /usr/include/qt/QtGui/qaccessible.h
// #include <qaccessible.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  ext block end

//  body block begin

type QAccessible struct {
	*qtrt.CObject
}

func (this *QAccessible) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QAccessible) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQAccessibleFromPointer(cthis unsafe.Pointer) *QAccessible {
	return &QAccessible{&qtrt.CObject{cthis}}
}
func (*QAccessible) NewFromPointer(cthis unsafe.Pointer) *QAccessible {
	return NewQAccessibleFromPointer(cthis)
}

// /usr/include/qt/QtGui/qaccessible.h:414
// index:0
// Public static Visibility=Default Availability=Available
// [8] QAccessibleInterface * queryAccessibleInterface(QObject *)
func (this *QAccessible) QueryAccessibleInterface(arg0 *qtcore.QObject /*777 QObject **/) *QAccessibleInterface /*777 QAccessibleInterface **/ {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QAccessible24queryAccessibleInterfaceEP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQAccessibleInterfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}
func QAccessible_QueryAccessibleInterface(arg0 *qtcore.QObject /*777 QObject **/) *QAccessibleInterface /*777 QAccessibleInterface **/ {
	var nilthis *QAccessible
	rv := nilthis.QueryAccessibleInterface(arg0)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:415
// index:0
// Public static Visibility=Default Availability=Available
// [4] QAccessible::Id uniqueId(QAccessibleInterface *)
func (this *QAccessible) UniqueId(iface *QAccessibleInterface /*777 QAccessibleInterface **/) uint {
	var convArg0 = iface.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QAccessible8uniqueIdEP20QAccessibleInterface", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}
func QAccessible_UniqueId(iface *QAccessibleInterface /*777 QAccessibleInterface **/) uint {
	var nilthis *QAccessible
	rv := nilthis.UniqueId(iface)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:416
// index:0
// Public static Visibility=Default Availability=Available
// [8] QAccessibleInterface * accessibleInterface(QAccessible::Id)
func (this *QAccessible) AccessibleInterface(uniqueId uint) *QAccessibleInterface /*777 QAccessibleInterface **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QAccessible19accessibleInterfaceEj", qtrt.FFI_TYPE_POINTER, uniqueId)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQAccessibleInterfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}
func QAccessible_AccessibleInterface(uniqueId uint) *QAccessibleInterface /*777 QAccessibleInterface **/ {
	var nilthis *QAccessible
	rv := nilthis.AccessibleInterface(uniqueId)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:417
// index:0
// Public static Visibility=Default Availability=Available
// [4] QAccessible::Id registerAccessibleInterface(QAccessibleInterface *)
func (this *QAccessible) RegisterAccessibleInterface(iface *QAccessibleInterface /*777 QAccessibleInterface **/) uint {
	var convArg0 = iface.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QAccessible27registerAccessibleInterfaceEP20QAccessibleInterface", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}
func QAccessible_RegisterAccessibleInterface(iface *QAccessibleInterface /*777 QAccessibleInterface **/) uint {
	var nilthis *QAccessible
	rv := nilthis.RegisterAccessibleInterface(iface)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:418
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void deleteAccessibleInterface(QAccessible::Id)
func (this *QAccessible) DeleteAccessibleInterface(uniqueId uint) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QAccessible25deleteAccessibleInterfaceEj", qtrt.FFI_TYPE_POINTER, uniqueId)
	gopp.ErrPrint(err, rv)
}
func QAccessible_DeleteAccessibleInterface(uniqueId uint) {
	var nilthis *QAccessible
	nilthis.DeleteAccessibleInterface(uniqueId)
}

// /usr/include/qt/QtGui/qaccessible.h:424
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void updateAccessibility(QAccessibleEvent *)
func (this *QAccessible) UpdateAccessibility(event *QAccessibleEvent /*777 QAccessibleEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QAccessible19updateAccessibilityEP16QAccessibleEvent", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
}
func QAccessible_UpdateAccessibility(event *QAccessibleEvent /*777 QAccessibleEvent **/) {
	var nilthis *QAccessible
	nilthis.UpdateAccessibility(event)
}

// /usr/include/qt/QtGui/qaccessible.h:426
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool isActive()
func (this *QAccessible) IsActive() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QAccessible8isActiveEv", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QAccessible_IsActive() bool {
	var nilthis *QAccessible
	rv := nilthis.IsActive()
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:427
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setActive(_Bool)
func (this *QAccessible) SetActive(active bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QAccessible9setActiveEb", qtrt.FFI_TYPE_POINTER, active)
	gopp.ErrPrint(err, rv)
}
func QAccessible_SetActive(active bool) {
	var nilthis *QAccessible
	nilthis.SetActive(active)
}

// /usr/include/qt/QtGui/qaccessible.h:428
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setRootObject(QObject *)
func (this *QAccessible) SetRootObject(object *qtcore.QObject /*777 QObject **/) {
	var convArg0 = object.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QAccessible13setRootObjectEP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
}
func QAccessible_SetRootObject(object *qtcore.QObject /*777 QObject **/) {
	var nilthis *QAccessible
	nilthis.SetRootObject(object)
}

// /usr/include/qt/QtGui/qaccessible.h:430
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void cleanup()
func (this *QAccessible) Cleanup() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QAccessible7cleanupEv", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
}
func QAccessible_Cleanup() {
	var nilthis *QAccessible
	nilthis.Cleanup()
}

func DeleteQAccessible(this *QAccessible) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QAccessibleD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

type QAccessible__Event = int

const QAccessible__SoundPlayed QAccessible__Event = 1
const QAccessible__Alert QAccessible__Event = 2
const QAccessible__ForegroundChanged QAccessible__Event = 3
const QAccessible__MenuStart QAccessible__Event = 4
const QAccessible__MenuEnd QAccessible__Event = 5
const QAccessible__PopupMenuStart QAccessible__Event = 6
const QAccessible__PopupMenuEnd QAccessible__Event = 7
const QAccessible__ContextHelpStart QAccessible__Event = 12
const QAccessible__ContextHelpEnd QAccessible__Event = 13
const QAccessible__DragDropStart QAccessible__Event = 14
const QAccessible__DragDropEnd QAccessible__Event = 15
const QAccessible__DialogStart QAccessible__Event = 16
const QAccessible__DialogEnd QAccessible__Event = 17
const QAccessible__ScrollingStart QAccessible__Event = 18
const QAccessible__ScrollingEnd QAccessible__Event = 19
const QAccessible__MenuCommand QAccessible__Event = 24
const QAccessible__ActionChanged QAccessible__Event = 257
const QAccessible__ActiveDescendantChanged QAccessible__Event = 258
const QAccessible__AttributeChanged QAccessible__Event = 259
const QAccessible__DocumentContentChanged QAccessible__Event = 260
const QAccessible__DocumentLoadComplete QAccessible__Event = 261
const QAccessible__DocumentLoadStopped QAccessible__Event = 262
const QAccessible__DocumentReload QAccessible__Event = 263
const QAccessible__HyperlinkEndIndexChanged QAccessible__Event = 264
const QAccessible__HyperlinkNumberOfAnchorsChanged QAccessible__Event = 265
const QAccessible__HyperlinkSelectedLinkChanged QAccessible__Event = 266
const QAccessible__HypertextLinkActivated QAccessible__Event = 267
const QAccessible__HypertextLinkSelected QAccessible__Event = 268
const QAccessible__HyperlinkStartIndexChanged QAccessible__Event = 269
const QAccessible__HypertextChanged QAccessible__Event = 270
const QAccessible__HypertextNLinksChanged QAccessible__Event = 271
const QAccessible__ObjectAttributeChanged QAccessible__Event = 272
const QAccessible__PageChanged QAccessible__Event = 273
const QAccessible__SectionChanged QAccessible__Event = 274
const QAccessible__TableCaptionChanged QAccessible__Event = 275
const QAccessible__TableColumnDescriptionChanged QAccessible__Event = 276
const QAccessible__TableColumnHeaderChanged QAccessible__Event = 277
const QAccessible__TableModelChanged QAccessible__Event = 278
const QAccessible__TableRowDescriptionChanged QAccessible__Event = 279
const QAccessible__TableRowHeaderChanged QAccessible__Event = 280
const QAccessible__TableSummaryChanged QAccessible__Event = 281
const QAccessible__TextAttributeChanged QAccessible__Event = 282
const QAccessible__TextCaretMoved QAccessible__Event = 283
const QAccessible__TextColumnChanged QAccessible__Event = 285
const QAccessible__TextInserted QAccessible__Event = 286
const QAccessible__TextRemoved QAccessible__Event = 287
const QAccessible__TextUpdated QAccessible__Event = 288
const QAccessible__TextSelectionChanged QAccessible__Event = 289
const QAccessible__VisibleDataChanged QAccessible__Event = 290
const QAccessible__ObjectCreated QAccessible__Event = 32768
const QAccessible__ObjectDestroyed QAccessible__Event = 32769
const QAccessible__ObjectShow QAccessible__Event = 32770
const QAccessible__ObjectHide QAccessible__Event = 32771
const QAccessible__ObjectReorder QAccessible__Event = 32772
const QAccessible__Focus QAccessible__Event = 32773
const QAccessible__Selection QAccessible__Event = 32774
const QAccessible__SelectionAdd QAccessible__Event = 32775
const QAccessible__SelectionRemove QAccessible__Event = 32776
const QAccessible__SelectionWithin QAccessible__Event = 32777
const QAccessible__StateChanged QAccessible__Event = 32778
const QAccessible__LocationChanged QAccessible__Event = 32779
const QAccessible__NameChanged QAccessible__Event = 32780
const QAccessible__DescriptionChanged QAccessible__Event = 32781
const QAccessible__ValueChanged QAccessible__Event = 32782
const QAccessible__ParentChanged QAccessible__Event = 32783
const QAccessible__HelpChanged QAccessible__Event = 32928
const QAccessible__DefaultActionChanged QAccessible__Event = 32944
const QAccessible__AcceleratorChanged QAccessible__Event = 32960
const QAccessible__InvalidEvent QAccessible__Event = 32961

type QAccessible__Role = int

const QAccessible__NoRole QAccessible__Role = 0
const QAccessible__TitleBar QAccessible__Role = 1
const QAccessible__MenuBar QAccessible__Role = 2
const QAccessible__ScrollBar QAccessible__Role = 3
const QAccessible__Grip QAccessible__Role = 4
const QAccessible__Sound QAccessible__Role = 5
const QAccessible__Cursor QAccessible__Role = 6
const QAccessible__Caret QAccessible__Role = 7
const QAccessible__AlertMessage QAccessible__Role = 8
const QAccessible__Window QAccessible__Role = 9
const QAccessible__Client QAccessible__Role = 10
const QAccessible__PopupMenu QAccessible__Role = 11
const QAccessible__MenuItem QAccessible__Role = 12
const QAccessible__ToolTip QAccessible__Role = 13
const QAccessible__Application QAccessible__Role = 14
const QAccessible__Document QAccessible__Role = 15
const QAccessible__Pane QAccessible__Role = 16
const QAccessible__Chart QAccessible__Role = 17
const QAccessible__Dialog QAccessible__Role = 18
const QAccessible__Border QAccessible__Role = 19
const QAccessible__Grouping QAccessible__Role = 20
const QAccessible__Separator QAccessible__Role = 21
const QAccessible__ToolBar QAccessible__Role = 22
const QAccessible__StatusBar QAccessible__Role = 23
const QAccessible__Table QAccessible__Role = 24
const QAccessible__ColumnHeader QAccessible__Role = 25
const QAccessible__RowHeader QAccessible__Role = 26
const QAccessible__Column QAccessible__Role = 27
const QAccessible__Row QAccessible__Role = 28
const QAccessible__Cell QAccessible__Role = 29
const QAccessible__Link QAccessible__Role = 30
const QAccessible__HelpBalloon QAccessible__Role = 31
const QAccessible__Assistant QAccessible__Role = 32
const QAccessible__List QAccessible__Role = 33
const QAccessible__ListItem QAccessible__Role = 34
const QAccessible__Tree QAccessible__Role = 35
const QAccessible__TreeItem QAccessible__Role = 36
const QAccessible__PageTab QAccessible__Role = 37
const QAccessible__PropertyPage QAccessible__Role = 38
const QAccessible__Indicator QAccessible__Role = 39
const QAccessible__Graphic QAccessible__Role = 40
const QAccessible__StaticText QAccessible__Role = 41
const QAccessible__EditableText QAccessible__Role = 42
const QAccessible__Button QAccessible__Role = 43
const QAccessible__PushButton QAccessible__Role = 43
const QAccessible__CheckBox QAccessible__Role = 44
const QAccessible__RadioButton QAccessible__Role = 45
const QAccessible__ComboBox QAccessible__Role = 46
const QAccessible__ProgressBar QAccessible__Role = 48
const QAccessible__Dial QAccessible__Role = 49
const QAccessible__HotkeyField QAccessible__Role = 50
const QAccessible__Slider QAccessible__Role = 51
const QAccessible__SpinBox QAccessible__Role = 52
const QAccessible__Canvas QAccessible__Role = 53
const QAccessible__Animation QAccessible__Role = 54
const QAccessible__Equation QAccessible__Role = 55
const QAccessible__ButtonDropDown QAccessible__Role = 56
const QAccessible__ButtonMenu QAccessible__Role = 57
const QAccessible__ButtonDropGrid QAccessible__Role = 58
const QAccessible__Whitespace QAccessible__Role = 59
const QAccessible__PageTabList QAccessible__Role = 60
const QAccessible__Clock QAccessible__Role = 61
const QAccessible__Splitter QAccessible__Role = 62
const QAccessible__LayeredPane QAccessible__Role = 128
const QAccessible__Terminal QAccessible__Role = 129
const QAccessible__Desktop QAccessible__Role = 130
const QAccessible__Paragraph QAccessible__Role = 131
const QAccessible__WebDocument QAccessible__Role = 132
const QAccessible__Section QAccessible__Role = 133
const QAccessible__ColorChooser QAccessible__Role = 1028
const QAccessible__Footer QAccessible__Role = 1038
const QAccessible__Form QAccessible__Role = 1040
const QAccessible__Heading QAccessible__Role = 1044
const QAccessible__Note QAccessible__Role = 1051
const QAccessible__ComplementaryContent QAccessible__Role = 1068
const QAccessible__UserRole QAccessible__Role = 65535

type QAccessible__Text = int

const QAccessible__Name QAccessible__Text = 0
const QAccessible__Description QAccessible__Text = 1
const QAccessible__Value QAccessible__Text = 2
const QAccessible__Help QAccessible__Text = 3
const QAccessible__Accelerator QAccessible__Text = 4
const QAccessible__DebugDescription QAccessible__Text = 5
const QAccessible__UserText QAccessible__Text = 65535

type QAccessible__RelationFlag = int

const QAccessible__Label QAccessible__RelationFlag = 1
const QAccessible__Labelled QAccessible__RelationFlag = 2
const QAccessible__Controller QAccessible__RelationFlag = 4
const QAccessible__Controlled QAccessible__RelationFlag = 8
const QAccessible__AllRelations QAccessible__RelationFlag = -1

type QAccessible__InterfaceType = int

const QAccessible__TextInterface QAccessible__InterfaceType = 0
const QAccessible__EditableTextInterface QAccessible__InterfaceType = 1
const QAccessible__ValueInterface QAccessible__InterfaceType = 2
const QAccessible__ActionInterface QAccessible__InterfaceType = 3
const QAccessible__ImageInterface QAccessible__InterfaceType = 4
const QAccessible__TableInterface QAccessible__InterfaceType = 5
const QAccessible__TableCellInterface QAccessible__InterfaceType = 6

type QAccessible__TextBoundaryType = int

const QAccessible__CharBoundary QAccessible__TextBoundaryType = 0
const QAccessible__WordBoundary QAccessible__TextBoundaryType = 1
const QAccessible__SentenceBoundary QAccessible__TextBoundaryType = 2
const QAccessible__ParagraphBoundary QAccessible__TextBoundaryType = 3
const QAccessible__LineBoundary QAccessible__TextBoundaryType = 4
const QAccessible__NoBoundary QAccessible__TextBoundaryType = 5

//  body block end
