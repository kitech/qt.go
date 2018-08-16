package qtqml

// /usr/include/qt/QtCore/qabstractitemmodel.h
// #include <qabstractitemmodel.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
*/
// import "C"
import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"unsafe"

	"github.com/kitech/qt.go/qtcore"
	"github.com/kitech/qt.go/qtrt"
)

//  ext block end

/*
full example:

////////
func init() {
	qtqml.RegisterModel("AListModel", NewAListModel)
}

type AListModel struct {
    mdlop *qtcore.QAbstractListModel
}

func NewAListModel(mdlop *qtcore.QAbstractListModel) qtqml.QGoListModel_ITF_RO2 {
	this := &AListModel{mdlop: mdlop}
	return this
}

func (this *AListModel) RowCount() int {
	return rand.Intn(50) + 1
}

func (this *AListModel) Data(index *qtcore.QModelIndex, role int) *qtcore.QVariant {
	log.Println("hehehhe", index.Row(), role)
	retv := qtcore.NewQVariant_15(fmt.Sprintf("retfromgo,row%d-role%d", index.Row(), role))
	return retv
}

func (this *AListModel) RoleNames() map[int]string {
	roles := map[int]string{
		qtcore.Qt__UserRole + 1: "name",
		qtcore.Qt__UserRole + 2: "hue",
		qtcore.Qt__UserRole + 3: "status",
		qtcore.Qt__UserRole + 4: "age",
	}
	return roles
}
*/

/*
qml usage:

GoListModel{typeName:"AListModel"}
*/

//  body block begin

func init() { init_cb_GoListModel() }
func init_cb_GoListModel() {
	qmlRegisterType_QGoListModel()
	dmobj := qtcore.NewQObjectFromPointer(unsafe.Pointer(nil))
	qtrt.SetAllInheritCallback(dmobj, "NewGoListModel", _cb_NewGoListModel)
	qtrt.SetAllInheritCallback(dmobj, "DeleteGoListModel", _cb_DeleteGoListModel)
}

func _cb_NewGoListModel(cthat uint64, name string) uint64 {
	// log.Println("NewGoListModel", cthat, unsafe.Pointer(uintptr(cthat)), name)
	newfn, ok := regedQmlModelTypes[name]
	if !ok {
		log.Println("qml model type not found:", name)
		return 0
	}

	mdlop := qtcore.NewQAbstractListModelFromPointer(unsafe.Pointer(uintptr(cthat)))
	fwdmdl := newfn(mdlop)
	mdlin := _NewQGoListModelInternal(mdlop)
	mdlin.initfwd(fwdmdl)
	newedQmlModels[cthat] = mdlin

	return cthat
}
func _cb_DeleteGoListModel(cthat uint64, name string, gothis uint64) uint64 {
	// log.Println("DeleteGoListModel", cthat, unsafe.Pointer(uintptr(cthat)), name, gothis, len(newedQmlModels))
	if mdlin, ok := newedQmlModels[cthat]; ok {
		delete(newedQmlModels, cthat)
		mdlin.dtor()
	} else {
		log.Println("GoListModel internal object not found:", unsafe.Pointer(uintptr(cthat)), name)
	}
	// log.Println("DeleteGoListModel", cthat, unsafe.Pointer(uintptr(cthat)), name, gothis, len(newedQmlModels))
	return 0
}

func qmlRegisterType_QGoListModel() {
	qtrt.InvokeQtFunc6("C_qmlRegisterType_QGoListModel", qtrt.FFI_TYPE_POINTER)
}

//////////
// mdlop argument used to implementation modifiable model
var regedQmlModelTypes = map[string]func(*qtcore.QAbstractListModel) QGoListModel_ITF_RO2{}
var newedQmlModels = map[uint64]*_QGoListModelInternal{}

func RegisterModel(name string, newfn func(*qtcore.QAbstractListModel) QGoListModel_ITF_RO2) {
	if f, ok := regedQmlModelTypes[name]; ok {
		log.Println("already registered:", name, f)
	} else {
		regedQmlModelTypes[name] = newfn
	}
}

type QGoListModel_ITF_RO2 interface {
	// qtcore.QObject_ITF

	RowCount() int
	Data(*qtcore.QModelIndex, int) *qtcore.QVariant
	RoleNames() map[int]string
}

type QGoListModel_ITF_RW2 interface {
	QGoListModel_ITF_RO2
	SetData()
}

type _QGoListModelInternal struct {
	*qtcore.QObject

	fwd QGoListModel_ITF_RO2
}

func _NewQGoListModelInternal(mdlop *qtcore.QAbstractListModel) *_QGoListModelInternal {
	this := &_QGoListModelInternal{}
	this.QObject = mdlop.QObject_PTR()

	return this
}

func (this *_QGoListModelInternal) initfwd(fwd QGoListModel_ITF_RO2) {
	this.fwd = fwd

	amdl := this
	qtrt.SetAllInheritCallback(amdl, "rowCount", this.RowCount)
	qtrt.SetAllInheritCallback(amdl, "data", this.Data)
	qtrt.SetAllInheritCallback(amdl, "roleNames", this.RoleNames)

	qtrt.SetAllInheritCallback(amdl, "index", this.Index)
	qtrt.SetAllInheritCallback(amdl, "sibling", this.Sibling)
	qtrt.SetAllInheritCallback(amdl, "dropMimeData", this.DropMimeData)
	qtrt.SetAllInheritCallback(amdl, "flags", this.Flags)

}
func (this *_QGoListModelInternal) dtor() {
	amdl := this
	qtrt.UnsetAllInheritCallback(amdl, "rowCount")
	qtrt.UnsetAllInheritCallback(amdl, "data")
	qtrt.UnsetAllInheritCallback(amdl, "roleNames")

	qtrt.UnsetAllInheritCallback(amdl, "index")
	qtrt.UnsetAllInheritCallback(amdl, "sibling")
	qtrt.UnsetAllInheritCallback(amdl, "dropMimeData")
	qtrt.UnsetAllInheritCallback(amdl, "flags")
	qtrt.UnsetAllInheritCallbackAll(amdl)

	this.fwd = nil
	this.SetCthis(unsafe.Pointer(nil))
	this.QObject = nil
	// runtime.SetFinalizer(this, func(obj interface{}) { log.Println("golist model internal obj dtored") })
}

func (this *_QGoListModelInternal) RowCount() int {
	return this.fwd.RowCount()
}
func (this *_QGoListModelInternal) Data(index *qtcore.QModelIndex, role int) uint64 {
	rv := this.fwd.Data(index, role)
	qtrt.ReleaseOwnerToQt(rv)
	return uint64(uintptr(rv.GetCthis()))
}
func (this *_QGoListModelInternal) RoleNames() uint64 {
	dat0 := this.fwd.RoleNames()
	roles := map[string]int{}
	for key, val := range dat0 {
		roles[val] = key
	}
	dat, _ := json.Marshal(roles)
	return uint64(uintptr(qtrt.CString(string(dat))))

}

func (this *_QGoListModelInternal) Index(row, column int) uint64 {
	// log.Println("hehehhe", row, column)
	midx := qtcore.NewQModelIndex()
	return uint64(uintptr(midx.GetCthis()))
}

func (this *_QGoListModelInternal) Sibling() uint64 {
	log.Println("hehehhe")
	midx := qtcore.NewQModelIndex()
	return uint64(uintptr(midx.GetCthis()))
}
func (this *_QGoListModelInternal) DropMimeData() uint64 {
	log.Println("hehehhe")
	return uint64(qtrt.GoBool2C(true))
}
func (this *_QGoListModelInternal) Flags() uint64 {
	log.Println("hehehhe")
	return 0
}

/////////////

/*

 */
type QGoListModel struct {
	*qtcore.QAbstractItemModel
}
type QGoListModel_ITF interface {
	qtcore.QAbstractItemModel_ITF
	QGoListModel_PTR() *QGoListModel
}

func (ptr *QGoListModel) QGoListModel_PTR() *QGoListModel { return ptr }

func (this *QGoListModel) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAbstractItemModel.GetCthis()
	}
}
func (this *QGoListModel) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractItemModel = qtcore.NewQAbstractItemModelFromPointer(cthis)
}
func NewQGoListModelFromPointer(cthis unsafe.Pointer) *QGoListModel {
	bcthis0 := qtcore.NewQAbstractItemModelFromPointer(cthis)
	return &QGoListModel{bcthis0}
}
func (*QGoListModel) NewFromPointer(cthis unsafe.Pointer) *QGoListModel {
	return NewQGoListModelFromPointer(cthis)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:393
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QGoListModel) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QGoListModel10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:396
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGoListModel(QObject *)

/*

 */
func NewQGoListModel(parent qtcore.QObject_ITF /*777 QObject **/) *QGoListModel {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QGoListModelC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGoListModelFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QGoListModel")
	return gothis
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:396
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGoListModel(QObject *)

/*

 */
func NewQGoListModel__() *QGoListModel {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN18QGoListModelC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGoListModelFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QGoListModel")
	return gothis
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:397
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGoListModel()

/*

 */
func DeleteQGoListModel(this *QGoListModel) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QGoListModelD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:399
// index:0
// Public virtual Visibility=Default Availability=Available
// [24] QModelIndex index(int, int, const QModelIndex &) const

/*
Returns the index of the item in the model specified by the given row, column and parent index.

When reimplementing this function in a subclass, call createIndex() to generate model indexes that other components can use to refer to items in your model.

See also createIndex().
*/
func (this *QGoListModel) Index(row int, column int, parent qtcore.QModelIndex_ITF) *qtcore.QModelIndex /*123*/ {
	var convArg2 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg2 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QGoListModel5indexEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, qtcore.DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:399
// index:0
// Public virtual Visibility=Default Availability=Available
// [24] QModelIndex index(int, int, const QModelIndex &) const

/*
Returns the index of the item in the model specified by the given row, column and parent index.

When reimplementing this function in a subclass, call createIndex() to generate model indexes that other components can use to refer to items in your model.

See also createIndex().
*/
func (this *QGoListModel) Index__(row int) *qtcore.QModelIndex /*123*/ {
	// arg: 1, int=Int, =Invalid, , Invalid
	column := int(0)
	// arg: 2, const QModelIndex &=LValueReference, QModelIndex=Record, , Invalid
	var convArg2 = qtcore.NewQModelIndex()
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QGoListModel5indexEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, qtcore.DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:399
// index:0
// Public virtual Visibility=Default Availability=Available
// [24] QModelIndex index(int, int, const QModelIndex &) const

/*
Returns the index of the item in the model specified by the given row, column and parent index.

When reimplementing this function in a subclass, call createIndex() to generate model indexes that other components can use to refer to items in your model.

See also createIndex().
*/
func (this *QGoListModel) Index__1(row int, column int) *qtcore.QModelIndex /*123*/ {
	// arg: 2, const QModelIndex &=LValueReference, QModelIndex=Record, , Invalid
	var convArg2 = qtcore.NewQModelIndex()
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QGoListModel5indexEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, qtcore.DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:400
// index:0
// Public virtual Visibility=Default Availability=Available
// [24] QModelIndex sibling(int, int, const QModelIndex &) const

/*
Returns the sibling at row and column for the item at index, or an invalid QModelIndex if there is no sibling at that location.

sibling() is just a convenience function that finds the item's parent, and uses it to retrieve the index of the child item in the specified row and column.

This method can optionally be overridden for implementation-specific optimization.

See also index(), QModelIndex::row(), and QModelIndex::column().
*/
func (this *QGoListModel) Sibling(row int, column int, idx qtcore.QModelIndex_ITF) *qtcore.QModelIndex /*123*/ {
	var convArg2 unsafe.Pointer
	if idx != nil && idx.QModelIndex_PTR() != nil {
		convArg2 = idx.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QGoListModel7siblingEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, qtcore.DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:401
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool dropMimeData(const QMimeData *, Qt::DropAction, int, int, const QModelIndex &)

/*
Handles the data supplied by a drag and drop operation that ended with the given action.

Returns true if the data and action were handled by the model; otherwise returns false.

The specified row, column and parent indicate the location of an item in the model where the operation ended. It is the responsibility of the model to complete the action at the correct location.

For instance, a drop action on an item in a QTreeView can result in new items either being inserted as children of the item specified by row, column, and parent, or as siblings of the item.

When row and column are -1 it means that the dropped data should be considered as dropped directly on parent. Usually this will mean appending the data as child items of parent. If row and column are greater than or equal zero, it means that the drop occurred just before the specified row and column in the specified parent.

The mimeTypes() member is called to get the list of acceptable MIME types. This default implementation assumes the default implementation of mimeTypes(), which returns a single default MIME type. If you reimplement mimeTypes() in your custom model to return multiple MIME types, you must reimplement this function to make use of them.

See also supportedDropActions(), canDropMimeData(), and Using drag and drop with item views.
*/
func (this *QGoListModel) DropMimeData(data qtcore.QMimeData_ITF /*777 const QMimeData **/, action int, row int, column int, parent qtcore.QModelIndex_ITF) bool {
	var convArg0 unsafe.Pointer
	if data != nil && data.QMimeData_PTR() != nil {
		convArg0 = data.QMimeData_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg4 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QGoListModel12dropMimeDataEPK9QMimeDataN2Qt10DropActionEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, action, row, column, convArg4)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:404
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] Qt::ItemFlags flags(const QModelIndex &) const

/*
Returns the item flags for the given index.

The base class implementation returns a combination of flags that enables the item (ItemIsEnabled) and allows it to be selected (ItemIsSelectable).

See also Qt::ItemFlags.
*/
func (this *QGoListModel) Flags(index qtcore.QModelIndex_ITF) int {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QGoListModel5flagsERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return int(rv)
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
}

//  keep block end
