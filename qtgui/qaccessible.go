package qtgui

// /usr/include/qt/QtGui/qaccessible.h
// #include <qaccessible.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

/*

 */
type QAccessible struct {
	*qtrt.CObject
}
type QAccessible_ITF interface {
	QAccessible_PTR() *QAccessible
}

func (ptr *QAccessible) QAccessible_PTR() *QAccessible { return ptr }

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

/*
If a QAccessibleInterface implementation exists for the given object, this function returns a pointer to the implementation; otherwise it returns 0.

The function calls all installed factory functions (from most recently installed to least recently installed) until one is found that provides an interface for the class of object. If no factory can provide an accessibility implementation for the class the function loads installed accessibility plugins, and tests if any of the plugins can provide the implementation.

If no implementation for the object's class is available, the function tries to find an implementation for the object's parent class, using the above strategy.

All interfaces are managed by an internal cache and should not be deleted.
*/
func (this *QAccessible) QueryAccessibleInterface(arg0 qtcore.QObject_ITF /*777 QObject **/) *QAccessibleInterface /*777 QAccessibleInterface **/ {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QObject_PTR() != nil {
		convArg0 = arg0.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QAccessible24queryAccessibleInterfaceEP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAccessibleInterfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QAccessible_QueryAccessibleInterface(arg0 qtcore.QObject_ITF /*777 QObject **/) *QAccessibleInterface /*777 QAccessibleInterface **/ {
	var nilthis *QAccessible
	rv := nilthis.QueryAccessibleInterface(arg0)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:415
// index:0
// Public static Visibility=Default Availability=Available
// [4] QAccessible::Id uniqueId(QAccessibleInterface *)

/*
Returns the unique ID for the QAccessibleInterface iface.
*/
func (this *QAccessible) UniqueId(iface QAccessibleInterface_ITF /*777 QAccessibleInterface **/) uint {
	var convArg0 unsafe.Pointer
	if iface != nil && iface.QAccessibleInterface_PTR() != nil {
		convArg0 = iface.QAccessibleInterface_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QAccessible8uniqueIdEP20QAccessibleInterface", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}
func QAccessible_UniqueId(iface QAccessibleInterface_ITF /*777 QAccessibleInterface **/) uint {
	var nilthis *QAccessible
	rv := nilthis.UniqueId(iface)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:416
// index:0
// Public static Visibility=Default Availability=Available
// [8] QAccessibleInterface * accessibleInterface(QAccessible::Id)

/*
Returns the QAccessibleInterface belonging to the id.

Returns 0 if the id is invalid.
*/
func (this *QAccessible) AccessibleInterface(uniqueId uint) *QAccessibleInterface /*777 QAccessibleInterface **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QAccessible19accessibleInterfaceEj", qtrt.FFI_TYPE_POINTER, uniqueId)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAccessibleInterfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
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

/*
Call this function to ensure that manually created interfaces are properly memory managed.

Must only be called exactly once per interface iface. This is implicitly called when calling queryAccessibleInterface, calling this function is only required when QAccessibleInterfaces are instantiated with the "new" operator. This is not recommended, whenever possible use the default functions and let queryAccessibleInterface() take care of this.

When it is necessary to reimplement the QAccessibleInterface::child() function and returning the child after constructing it, this function needs to be called.
*/
func (this *QAccessible) RegisterAccessibleInterface(iface QAccessibleInterface_ITF /*777 QAccessibleInterface **/) uint {
	var convArg0 unsafe.Pointer
	if iface != nil && iface.QAccessibleInterface_PTR() != nil {
		convArg0 = iface.QAccessibleInterface_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QAccessible27registerAccessibleInterfaceEP20QAccessibleInterface", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}
func QAccessible_RegisterAccessibleInterface(iface QAccessibleInterface_ITF /*777 QAccessibleInterface **/) uint {
	var nilthis *QAccessible
	rv := nilthis.RegisterAccessibleInterface(iface)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:418
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void deleteAccessibleInterface(QAccessible::Id)

/*
Removes the interface belonging to this id from the cache and deletes it. The id becomes invalid an may be re-used by the cache.
*/
func (this *QAccessible) DeleteAccessibleInterface(uniqueId uint) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QAccessible25deleteAccessibleInterfaceEj", qtrt.FFI_TYPE_POINTER, uniqueId)
	qtrt.ErrPrint(err, rv)
}
func QAccessible_DeleteAccessibleInterface(uniqueId uint) {
	var nilthis *QAccessible
	nilthis.DeleteAccessibleInterface(uniqueId)
}

// /usr/include/qt/QtGui/qaccessible.h:424
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void updateAccessibility(QAccessibleEvent *)

/*
Notifies about a change that might be relevant for accessibility clients.

event provides details about the change. These include the source of the change and the nature of the change. The event should contain enough information give meaningful notifications.

For example, the type ValueChange indicates that the position of a slider has been changed.

Call this function whenever the state of your accessible object or one of its sub-elements has been changed either programmatically (e.g. by calling QLabel::setText()) or by user interaction.

If there are no accessibility tools listening to this event, the performance penalty for calling this function is small, but if determining the parameters of the call is expensive you can test QAccessible::isActive() to avoid unnecessary computation.
*/
func (this *QAccessible) UpdateAccessibility(event QAccessibleEvent_ITF /*777 QAccessibleEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QAccessibleEvent_PTR() != nil {
		convArg0 = event.QAccessibleEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QAccessible19updateAccessibilityEP16QAccessibleEvent", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
func QAccessible_UpdateAccessibility(event QAccessibleEvent_ITF /*777 QAccessibleEvent **/) {
	var nilthis *QAccessible
	nilthis.UpdateAccessibility(event)
}

// /usr/include/qt/QtGui/qaccessible.h:426
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool isActive()

/*
Returns true if the platform requested accessibility information.

This function will return false until a tool such as a screen reader accessed the accessibility framework. It is still possible to use QAccessible::queryAccessibleInterface() even if accessibility is not active. But there will be no notifications sent to the platform.

It is recommended to use this function to prevent expensive notifications via updateAccessibility() when they are not needed.
*/
func (this *QAccessible) IsActive() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QAccessible8isActiveEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
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
// [-2] void setActive(bool)

/*

 */
func (this *QAccessible) SetActive(active bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QAccessible9setActiveEb", qtrt.FFI_TYPE_POINTER, active)
	qtrt.ErrPrint(err, rv)
}
func QAccessible_SetActive(active bool) {
	var nilthis *QAccessible
	nilthis.SetActive(active)
}

// /usr/include/qt/QtGui/qaccessible.h:428
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setRootObject(QObject *)

/*
Sets the root object of the accessible objects of this application to object. All other accessible objects are reachable using object navigation from the root object.

Normally, it isn't necessary to call this function, because Qt sets the QApplication object as the root object immediately before the event loop is entered in QApplication::exec().

Use QAccessible::installRootObjectHandler() to redirect the function call to a customized handler function.

See also queryAccessibleInterface().
*/
func (this *QAccessible) SetRootObject(object qtcore.QObject_ITF /*777 QObject **/) {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QAccessible13setRootObjectEP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
func QAccessible_SetRootObject(object qtcore.QObject_ITF /*777 QObject **/) {
	var nilthis *QAccessible
	nilthis.SetRootObject(object)
}

// /usr/include/qt/QtGui/qaccessible.h:430
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void cleanup()

/*

 */
func (this *QAccessible) Cleanup() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QAccessible7cleanupEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
}
func QAccessible_Cleanup() {
	var nilthis *QAccessible
	nilthis.Cleanup()
}

func DeleteQAccessible(this *QAccessible) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QAccessibleD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*
This enum type defines accessible event types.



The values for this enum are defined to be the same as those defined in the IAccessible2 and MSAA specifications.

*/
type QAccessible__Event = int

//
const QAccessible__SoundPlayed QAccessible__Event = 1

//
const QAccessible__Alert QAccessible__Event = 2

//
const QAccessible__ForegroundChanged QAccessible__Event = 3

//
const QAccessible__MenuStart QAccessible__Event = 4

//
const QAccessible__MenuEnd QAccessible__Event = 5

//
const QAccessible__PopupMenuStart QAccessible__Event = 6

//
const QAccessible__PopupMenuEnd QAccessible__Event = 7

//
const QAccessible__ContextHelpStart QAccessible__Event = 12

//
const QAccessible__ContextHelpEnd QAccessible__Event = 13

//
const QAccessible__DragDropStart QAccessible__Event = 14

//
const QAccessible__DragDropEnd QAccessible__Event = 15

//
const QAccessible__DialogStart QAccessible__Event = 16

//
const QAccessible__DialogEnd QAccessible__Event = 17

//
const QAccessible__ScrollingStart QAccessible__Event = 18

//
const QAccessible__ScrollingEnd QAccessible__Event = 19

//
const QAccessible__MenuCommand QAccessible__Event = 24

//
const QAccessible__ActionChanged QAccessible__Event = 257

//
const QAccessible__ActiveDescendantChanged QAccessible__Event = 258

//
const QAccessible__AttributeChanged QAccessible__Event = 259

//
const QAccessible__DocumentContentChanged QAccessible__Event = 260

//
const QAccessible__DocumentLoadComplete QAccessible__Event = 261

//
const QAccessible__DocumentLoadStopped QAccessible__Event = 262

//
const QAccessible__DocumentReload QAccessible__Event = 263

//
const QAccessible__HyperlinkEndIndexChanged QAccessible__Event = 264

//
const QAccessible__HyperlinkNumberOfAnchorsChanged QAccessible__Event = 265

//
const QAccessible__HyperlinkSelectedLinkChanged QAccessible__Event = 266

//
const QAccessible__HypertextLinkActivated QAccessible__Event = 267

//
const QAccessible__HypertextLinkSelected QAccessible__Event = 268

//
const QAccessible__HyperlinkStartIndexChanged QAccessible__Event = 269

//
const QAccessible__HypertextChanged QAccessible__Event = 270

//
const QAccessible__HypertextNLinksChanged QAccessible__Event = 271

//
const QAccessible__ObjectAttributeChanged QAccessible__Event = 272

//
const QAccessible__PageChanged QAccessible__Event = 273

//
const QAccessible__SectionChanged QAccessible__Event = 274

//
const QAccessible__TableCaptionChanged QAccessible__Event = 275

//
const QAccessible__TableColumnDescriptionChanged QAccessible__Event = 276

//
const QAccessible__TableColumnHeaderChanged QAccessible__Event = 277

//
const QAccessible__TableModelChanged QAccessible__Event = 278

//
const QAccessible__TableRowDescriptionChanged QAccessible__Event = 279

//
const QAccessible__TableRowHeaderChanged QAccessible__Event = 280

//
const QAccessible__TableSummaryChanged QAccessible__Event = 281

//
const QAccessible__TextAttributeChanged QAccessible__Event = 282

//
const QAccessible__TextCaretMoved QAccessible__Event = 283

//
const QAccessible__TextColumnChanged QAccessible__Event = 285

//
const QAccessible__TextInserted QAccessible__Event = 286

//
const QAccessible__TextRemoved QAccessible__Event = 287

//
const QAccessible__TextUpdated QAccessible__Event = 288

//
const QAccessible__TextSelectionChanged QAccessible__Event = 289

//
const QAccessible__VisibleDataChanged QAccessible__Event = 290

//
const QAccessible__ObjectCreated QAccessible__Event = 32768

//
const QAccessible__ObjectDestroyed QAccessible__Event = 32769

//
const QAccessible__ObjectShow QAccessible__Event = 32770

//
const QAccessible__ObjectHide QAccessible__Event = 32771

//
const QAccessible__ObjectReorder QAccessible__Event = 32772

//
const QAccessible__Focus QAccessible__Event = 32773

//
const QAccessible__Selection QAccessible__Event = 32774

//
const QAccessible__SelectionAdd QAccessible__Event = 32775

//
const QAccessible__SelectionRemove QAccessible__Event = 32776

//
const QAccessible__SelectionWithin QAccessible__Event = 32777

//
const QAccessible__StateChanged QAccessible__Event = 32778

//
const QAccessible__LocationChanged QAccessible__Event = 32779

//
const QAccessible__NameChanged QAccessible__Event = 32780

//
const QAccessible__DescriptionChanged QAccessible__Event = 32781

//
const QAccessible__ValueChanged QAccessible__Event = 32782

//
const QAccessible__ParentChanged QAccessible__Event = 32783

//
const QAccessible__HelpChanged QAccessible__Event = 32928

//
const QAccessible__DefaultActionChanged QAccessible__Event = 32944

//
const QAccessible__AcceleratorChanged QAccessible__Event = 32960

//
const QAccessible__InvalidEvent QAccessible__Event = 32961

func (this *QAccessible) EventItemName(val int) string {
	switch val {
	case QAccessible__SoundPlayed: // 1
		return "SoundPlayed"
	case QAccessible__Alert: // 2
		return "Alert"
	case QAccessible__ForegroundChanged: // 3
		return "ForegroundChanged"
	case QAccessible__MenuStart: // 4
		return "MenuStart"
	case QAccessible__MenuEnd: // 5
		return "MenuEnd"
	case QAccessible__PopupMenuStart: // 6
		return "PopupMenuStart"
	case QAccessible__PopupMenuEnd: // 7
		return "PopupMenuEnd"
	case QAccessible__ContextHelpStart: // 12
		return "ContextHelpStart"
	case QAccessible__ContextHelpEnd: // 13
		return "ContextHelpEnd"
	case QAccessible__DragDropStart: // 14
		return "DragDropStart"
	case QAccessible__DragDropEnd: // 15
		return "DragDropEnd"
	case QAccessible__DialogStart: // 16
		return "DialogStart"
	case QAccessible__DialogEnd: // 17
		return "DialogEnd"
	case QAccessible__ScrollingStart: // 18
		return "ScrollingStart"
	case QAccessible__ScrollingEnd: // 19
		return "ScrollingEnd"
	case QAccessible__MenuCommand: // 24
		return "MenuCommand"
	case QAccessible__ActionChanged: // 257
		return "ActionChanged"
	case QAccessible__ActiveDescendantChanged: // 258
		return "ActiveDescendantChanged"
	case QAccessible__AttributeChanged: // 259
		return "AttributeChanged"
	case QAccessible__DocumentContentChanged: // 260
		return "DocumentContentChanged"
	case QAccessible__DocumentLoadComplete: // 261
		return "DocumentLoadComplete"
	case QAccessible__DocumentLoadStopped: // 262
		return "DocumentLoadStopped"
	case QAccessible__DocumentReload: // 263
		return "DocumentReload"
	case QAccessible__HyperlinkEndIndexChanged: // 264
		return "HyperlinkEndIndexChanged"
	case QAccessible__HyperlinkNumberOfAnchorsChanged: // 265
		return "HyperlinkNumberOfAnchorsChanged"
	case QAccessible__HyperlinkSelectedLinkChanged: // 266
		return "HyperlinkSelectedLinkChanged"
	case QAccessible__HypertextLinkActivated: // 267
		return "HypertextLinkActivated"
	case QAccessible__HypertextLinkSelected: // 268
		return "HypertextLinkSelected"
	case QAccessible__HyperlinkStartIndexChanged: // 269
		return "HyperlinkStartIndexChanged"
	case QAccessible__HypertextChanged: // 270
		return "HypertextChanged"
	case QAccessible__HypertextNLinksChanged: // 271
		return "HypertextNLinksChanged"
	case QAccessible__ObjectAttributeChanged: // 272
		return "ObjectAttributeChanged"
	case QAccessible__PageChanged: // 273
		return "PageChanged"
	case QAccessible__SectionChanged: // 274
		return "SectionChanged"
	case QAccessible__TableCaptionChanged: // 275
		return "TableCaptionChanged"
	case QAccessible__TableColumnDescriptionChanged: // 276
		return "TableColumnDescriptionChanged"
	case QAccessible__TableColumnHeaderChanged: // 277
		return "TableColumnHeaderChanged"
	case QAccessible__TableModelChanged: // 278
		return "TableModelChanged"
	case QAccessible__TableRowDescriptionChanged: // 279
		return "TableRowDescriptionChanged"
	case QAccessible__TableRowHeaderChanged: // 280
		return "TableRowHeaderChanged"
	case QAccessible__TableSummaryChanged: // 281
		return "TableSummaryChanged"
	case QAccessible__TextAttributeChanged: // 282
		return "TextAttributeChanged"
	case QAccessible__TextCaretMoved: // 283
		return "TextCaretMoved"
	case QAccessible__TextColumnChanged: // 285
		return "TextColumnChanged"
	case QAccessible__TextInserted: // 286
		return "TextInserted"
	case QAccessible__TextRemoved: // 287
		return "TextRemoved"
	case QAccessible__TextUpdated: // 288
		return "TextUpdated"
	case QAccessible__TextSelectionChanged: // 289
		return "TextSelectionChanged"
	case QAccessible__VisibleDataChanged: // 290
		return "VisibleDataChanged"
	case QAccessible__ObjectCreated: // 32768
		return "ObjectCreated"
	case QAccessible__ObjectDestroyed: // 32769
		return "ObjectDestroyed"
	case QAccessible__ObjectShow: // 32770
		return "ObjectShow"
	case QAccessible__ObjectHide: // 32771
		return "ObjectHide"
	case QAccessible__ObjectReorder: // 32772
		return "ObjectReorder"
	case QAccessible__Focus: // 32773
		return "Focus"
	case QAccessible__Selection: // 32774
		return "Selection"
	case QAccessible__SelectionAdd: // 32775
		return "SelectionAdd"
	case QAccessible__SelectionRemove: // 32776
		return "SelectionRemove"
	case QAccessible__SelectionWithin: // 32777
		return "SelectionWithin"
	case QAccessible__StateChanged: // 32778
		return "StateChanged"
	case QAccessible__LocationChanged: // 32779
		return "LocationChanged"
	case QAccessible__NameChanged: // 32780
		return "NameChanged"
	case QAccessible__DescriptionChanged: // 32781
		return "DescriptionChanged"
	case QAccessible__ValueChanged: // 32782
		return "ValueChanged"
	case QAccessible__ParentChanged: // 32783
		return "ParentChanged"
	case QAccessible__HelpChanged: // 32928
		return "HelpChanged"
	case QAccessible__DefaultActionChanged: // 32944
		return "DefaultActionChanged"
	case QAccessible__AcceleratorChanged: // 32960
		return "AcceleratorChanged"
	case QAccessible__InvalidEvent: // 32961
		return "InvalidEvent"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QAccessible_EventItemName(val int) string {
	var nilthis *QAccessible
	return nilthis.EventItemName(val)
}

/*
This enum defines the role of an accessible object. The roles are:


*/
type QAccessible__Role = int

//
const QAccessible__NoRole QAccessible__Role = 0

//
const QAccessible__TitleBar QAccessible__Role = 1

//
const QAccessible__MenuBar QAccessible__Role = 2

//
const QAccessible__ScrollBar QAccessible__Role = 3

//
const QAccessible__Grip QAccessible__Role = 4

//
const QAccessible__Sound QAccessible__Role = 5

//
const QAccessible__Cursor QAccessible__Role = 6

//
const QAccessible__Caret QAccessible__Role = 7

//
const QAccessible__AlertMessage QAccessible__Role = 8

//
const QAccessible__Window QAccessible__Role = 9

//
const QAccessible__Client QAccessible__Role = 10

//
const QAccessible__PopupMenu QAccessible__Role = 11

//
const QAccessible__MenuItem QAccessible__Role = 12

//
const QAccessible__ToolTip QAccessible__Role = 13

//
const QAccessible__Application QAccessible__Role = 14

//
const QAccessible__Document QAccessible__Role = 15

//
const QAccessible__Pane QAccessible__Role = 16

//
const QAccessible__Chart QAccessible__Role = 17

//
const QAccessible__Dialog QAccessible__Role = 18

//
const QAccessible__Border QAccessible__Role = 19

//
const QAccessible__Grouping QAccessible__Role = 20

//
const QAccessible__Separator QAccessible__Role = 21

//
const QAccessible__ToolBar QAccessible__Role = 22

//
const QAccessible__StatusBar QAccessible__Role = 23

//
const QAccessible__Table QAccessible__Role = 24

//
const QAccessible__ColumnHeader QAccessible__Role = 25

//
const QAccessible__RowHeader QAccessible__Role = 26

//
const QAccessible__Column QAccessible__Role = 27

//
const QAccessible__Row QAccessible__Role = 28

//
const QAccessible__Cell QAccessible__Role = 29

//
const QAccessible__Link QAccessible__Role = 30

//
const QAccessible__HelpBalloon QAccessible__Role = 31

//
const QAccessible__Assistant QAccessible__Role = 32

//
const QAccessible__List QAccessible__Role = 33

//
const QAccessible__ListItem QAccessible__Role = 34

//
const QAccessible__Tree QAccessible__Role = 35

//
const QAccessible__TreeItem QAccessible__Role = 36

//
const QAccessible__PageTab QAccessible__Role = 37

//
const QAccessible__PropertyPage QAccessible__Role = 38

//
const QAccessible__Indicator QAccessible__Role = 39

//
const QAccessible__Graphic QAccessible__Role = 40

//
const QAccessible__StaticText QAccessible__Role = 41

//
const QAccessible__EditableText QAccessible__Role = 42

//
const QAccessible__Button QAccessible__Role = 43

//
const QAccessible__PushButton QAccessible__Role = 43

//
const QAccessible__CheckBox QAccessible__Role = 44

//
const QAccessible__RadioButton QAccessible__Role = 45

//
const QAccessible__ComboBox QAccessible__Role = 46

//
const QAccessible__ProgressBar QAccessible__Role = 48

//
const QAccessible__Dial QAccessible__Role = 49

//
const QAccessible__HotkeyField QAccessible__Role = 50

//
const QAccessible__Slider QAccessible__Role = 51

//
const QAccessible__SpinBox QAccessible__Role = 52

//
const QAccessible__Canvas QAccessible__Role = 53

//
const QAccessible__Animation QAccessible__Role = 54

//
const QAccessible__Equation QAccessible__Role = 55

//
const QAccessible__ButtonDropDown QAccessible__Role = 56

//
const QAccessible__ButtonMenu QAccessible__Role = 57

//
const QAccessible__ButtonDropGrid QAccessible__Role = 58

//
const QAccessible__Whitespace QAccessible__Role = 59

//
const QAccessible__PageTabList QAccessible__Role = 60

//
const QAccessible__Clock QAccessible__Role = 61

//
const QAccessible__Splitter QAccessible__Role = 62

//
const QAccessible__LayeredPane QAccessible__Role = 128

//
const QAccessible__Terminal QAccessible__Role = 129

//
const QAccessible__Desktop QAccessible__Role = 130

//
const QAccessible__Paragraph QAccessible__Role = 131

//
const QAccessible__WebDocument QAccessible__Role = 132

//
const QAccessible__Section QAccessible__Role = 133

//
const QAccessible__ColorChooser QAccessible__Role = 1028

//
const QAccessible__Footer QAccessible__Role = 1038

//
const QAccessible__Form QAccessible__Role = 1040

//
const QAccessible__Heading QAccessible__Role = 1044

//
const QAccessible__Note QAccessible__Role = 1051

//
const QAccessible__ComplementaryContent QAccessible__Role = 1068

//
const QAccessible__UserRole QAccessible__Role = 65535

func (this *QAccessible) RoleItemName(val int) string {
	switch val {
	case QAccessible__NoRole: // 0
		return "NoRole"
	case QAccessible__TitleBar: // 1
		return "TitleBar"
	case QAccessible__MenuBar: // 2
		return "MenuBar"
	case QAccessible__ScrollBar: // 3
		return "ScrollBar"
	case QAccessible__Grip: // 4
		return "Grip"
	case QAccessible__Sound: // 5
		return "Sound"
	case QAccessible__Cursor: // 6
		return "Cursor"
	case QAccessible__Caret: // 7
		return "Caret"
	case QAccessible__AlertMessage: // 8
		return "AlertMessage"
	case QAccessible__Window: // 9
		return "Window"
	case QAccessible__Client: // 10
		return "Client"
	case QAccessible__PopupMenu: // 11
		return "PopupMenu"
	case QAccessible__MenuItem: // 12
		return "MenuItem"
	case QAccessible__ToolTip: // 13
		return "ToolTip"
	case QAccessible__Application: // 14
		return "Application"
	case QAccessible__Document: // 15
		return "Document"
	case QAccessible__Pane: // 16
		return "Pane"
	case QAccessible__Chart: // 17
		return "Chart"
	case QAccessible__Dialog: // 18
		return "Dialog"
	case QAccessible__Border: // 19
		return "Border"
	case QAccessible__Grouping: // 20
		return "Grouping"
	case QAccessible__Separator: // 21
		return "Separator"
	case QAccessible__ToolBar: // 22
		return "ToolBar"
	case QAccessible__StatusBar: // 23
		return "StatusBar"
	case QAccessible__Table: // 24
		return "Table"
	case QAccessible__ColumnHeader: // 25
		return "ColumnHeader"
	case QAccessible__RowHeader: // 26
		return "RowHeader"
	case QAccessible__Column: // 27
		return "Column"
	case QAccessible__Row: // 28
		return "Row"
	case QAccessible__Cell: // 29
		return "Cell"
	case QAccessible__Link: // 30
		return "Link"
	case QAccessible__HelpBalloon: // 31
		return "HelpBalloon"
	case QAccessible__Assistant: // 32
		return "Assistant"
	case QAccessible__List: // 33
		return "List"
	case QAccessible__ListItem: // 34
		return "ListItem"
	case QAccessible__Tree: // 35
		return "Tree"
	case QAccessible__TreeItem: // 36
		return "TreeItem"
	case QAccessible__PageTab: // 37
		return "PageTab"
	case QAccessible__PropertyPage: // 38
		return "PropertyPage"
	case QAccessible__Indicator: // 39
		return "Indicator"
	case QAccessible__Graphic: // 40
		return "Graphic"
	case QAccessible__StaticText: // 41
		return "StaticText"
	case QAccessible__EditableText: // 42
		return "EditableText"
	case QAccessible__Button: // 43
		return "Button,PushButton"
		// case QAccessible__PushButton: // 43
		// return ""
	case QAccessible__CheckBox: // 44
		return "CheckBox"
	case QAccessible__RadioButton: // 45
		return "RadioButton"
	case QAccessible__ComboBox: // 46
		return "ComboBox"
	case QAccessible__ProgressBar: // 48
		return "ProgressBar"
	case QAccessible__Dial: // 49
		return "Dial"
	case QAccessible__HotkeyField: // 50
		return "HotkeyField"
	case QAccessible__Slider: // 51
		return "Slider"
	case QAccessible__SpinBox: // 52
		return "SpinBox"
	case QAccessible__Canvas: // 53
		return "Canvas"
	case QAccessible__Animation: // 54
		return "Animation"
	case QAccessible__Equation: // 55
		return "Equation"
	case QAccessible__ButtonDropDown: // 56
		return "ButtonDropDown"
	case QAccessible__ButtonMenu: // 57
		return "ButtonMenu"
	case QAccessible__ButtonDropGrid: // 58
		return "ButtonDropGrid"
	case QAccessible__Whitespace: // 59
		return "Whitespace"
	case QAccessible__PageTabList: // 60
		return "PageTabList"
	case QAccessible__Clock: // 61
		return "Clock"
	case QAccessible__Splitter: // 62
		return "Splitter"
	case QAccessible__LayeredPane: // 128
		return "LayeredPane"
	case QAccessible__Terminal: // 129
		return "Terminal"
	case QAccessible__Desktop: // 130
		return "Desktop"
	case QAccessible__Paragraph: // 131
		return "Paragraph"
	case QAccessible__WebDocument: // 132
		return "WebDocument"
	case QAccessible__Section: // 133
		return "Section"
	case QAccessible__ColorChooser: // 1028
		return "ColorChooser"
	case QAccessible__Footer: // 1038
		return "Footer"
	case QAccessible__Form: // 1040
		return "Form"
	case QAccessible__Heading: // 1044
		return "Heading"
	case QAccessible__Note: // 1051
		return "Note"
	case QAccessible__ComplementaryContent: // 1068
		return "ComplementaryContent"
	case QAccessible__UserRole: // 65535
		return "UserRole"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QAccessible_RoleItemName(val int) string {
	var nilthis *QAccessible
	return nilthis.RoleItemName(val)
}

/*
This enum specifies string information that an accessible object returns.


*/
type QAccessible__Text = int

// The name of the object. This can be used both as an identifier or a short description by accessible clients.
const QAccessible__Name QAccessible__Text = 0

// A short text describing the object.
const QAccessible__Description QAccessible__Text = 1

// The value of the object.
const QAccessible__Value QAccessible__Text = 2

// A longer text giving information about how to use the object.
const QAccessible__Help QAccessible__Text = 3

// The keyboard shortcut that executes the object's default action.
const QAccessible__Accelerator QAccessible__Text = 4

//
const QAccessible__DebugDescription QAccessible__Text = 5

//
const QAccessible__UserText QAccessible__Text = 65535

func (this *QAccessible) TextItemName(val int) string {
	switch val {
	case QAccessible__Name: // 0
		return "Name"
	case QAccessible__Description: // 1
		return "Description"
	case QAccessible__Value: // 2
		return "Value"
	case QAccessible__Help: // 3
		return "Help"
	case QAccessible__Accelerator: // 4
		return "Accelerator"
	case QAccessible__DebugDescription: // 5
		return "DebugDescription"
	case QAccessible__UserText: // 65535
		return "UserText"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QAccessible_TextItemName(val int) string {
	var nilthis *QAccessible
	return nilthis.TextItemName(val)
}

/*


 */
type QAccessible__RelationFlag = int

//
const QAccessible__Label QAccessible__RelationFlag = 1

//
const QAccessible__Labelled QAccessible__RelationFlag = 2

//
const QAccessible__Controller QAccessible__RelationFlag = 4

//
const QAccessible__Controlled QAccessible__RelationFlag = 8

//
const QAccessible__AllRelations QAccessible__RelationFlag = -1

func (this *QAccessible) RelationFlagItemName(val int) string {
	switch val {
	case QAccessible__Label: // 1
		return "Label"
	case QAccessible__Labelled: // 2
		return "Labelled"
	case QAccessible__Controller: // 4
		return "Controller"
	case QAccessible__Controlled: // 8
		return "Controlled"
	case QAccessible__AllRelations: // -1
		return "AllRelations"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QAccessible_RelationFlagItemName(val int) string {
	var nilthis *QAccessible
	return nilthis.RelationFlagItemName(val)
}

/*
QAccessibleInterface supports several sub interfaces. In order to provide more information about some objects, their accessible representation should implement one or more of these interfaces.

Note: When subclassing one of these interfaces, QAccessibleInterface::interface_cast() needs to be implemented.



See also QAccessibleInterface::interface_cast(), QAccessibleTextInterface, QAccessibleValueInterface, QAccessibleActionInterface, QAccessibleTableInterface, and QAccessibleTableCellInterface.

*/
type QAccessible__InterfaceType = int

// For text that supports selections or is more than one line. Simple labels do not need to implement this interface.
const QAccessible__TextInterface QAccessible__InterfaceType = 0

//
const QAccessible__EditableTextInterface QAccessible__InterfaceType = 1

// For objects that are used to manipulate a value, for example slider or scroll bar.
const QAccessible__ValueInterface QAccessible__InterfaceType = 2

// For interactive objects that allow the user to trigger an action. Basically everything that allows for example mouse interaction.
const QAccessible__ActionInterface QAccessible__InterfaceType = 3

//
const QAccessible__ImageInterface QAccessible__InterfaceType = 4

// For lists, tables and trees.
const QAccessible__TableInterface QAccessible__InterfaceType = 5

// For cells in a TableInterface object.
const QAccessible__TableCellInterface QAccessible__InterfaceType = 6

func (this *QAccessible) InterfaceTypeItemName(val int) string {
	switch val {
	case QAccessible__TextInterface: // 0
		return "TextInterface"
	case QAccessible__EditableTextInterface: // 1
		return "EditableTextInterface"
	case QAccessible__ValueInterface: // 2
		return "ValueInterface"
	case QAccessible__ActionInterface: // 3
		return "ActionInterface"
	case QAccessible__ImageInterface: // 4
		return "ImageInterface"
	case QAccessible__TableInterface: // 5
		return "TableInterface"
	case QAccessible__TableCellInterface: // 6
		return "TableCellInterface"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QAccessible_InterfaceTypeItemName(val int) string {
	var nilthis *QAccessible
	return nilthis.InterfaceTypeItemName(val)
}

/*
This enum describes different types of text boundaries. It follows the IAccessible2 API and is used in the QAccessibleTextInterface.



See also QAccessibleTextInterface.

*/
type QAccessible__TextBoundaryType = int

// Use individual characters as boundary.
const QAccessible__CharBoundary QAccessible__TextBoundaryType = 0

// Use words as boundaries.
const QAccessible__WordBoundary QAccessible__TextBoundaryType = 1

// Use sentences as boundary.
const QAccessible__SentenceBoundary QAccessible__TextBoundaryType = 2

// Use paragraphs as boundary.
const QAccessible__ParagraphBoundary QAccessible__TextBoundaryType = 3

// Use newlines as boundary.
const QAccessible__LineBoundary QAccessible__TextBoundaryType = 4

// No boundary (use the whole text).
const QAccessible__NoBoundary QAccessible__TextBoundaryType = 5

func (this *QAccessible) TextBoundaryTypeItemName(val int) string {
	switch val {
	case QAccessible__CharBoundary: // 0
		return "CharBoundary"
	case QAccessible__WordBoundary: // 1
		return "WordBoundary"
	case QAccessible__SentenceBoundary: // 2
		return "SentenceBoundary"
	case QAccessible__ParagraphBoundary: // 3
		return "ParagraphBoundary"
	case QAccessible__LineBoundary: // 4
		return "LineBoundary"
	case QAccessible__NoBoundary: // 5
		return "NoBoundary"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QAccessible_TextBoundaryTypeItemName(val int) string {
	var nilthis *QAccessible
	return nilthis.TextBoundaryTypeItemName(val)
}

//  body block end

//  keep block begin

func init_unused_10801() {
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
}

//  keep block end
