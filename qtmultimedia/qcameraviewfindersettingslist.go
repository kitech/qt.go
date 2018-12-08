package qtmultimedia

// /usr/include/qt/QtMultimedia/qcameraviewfindersettings.h
// #include <qcameraviewfindersettings.h>
// #include <QtMultimedia>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 0
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtnetwork"

//  ext block end

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
	if false {
		qtnetwork.KeepMe()
	}
}

//  keep block end

//  body block begin
type QCameraViewfinderSettingsList struct {
	*qtrt.CObject
}

// QList<T> & operator=(const QList<T> &)
func (this *QCameraViewfinderSettingsList) Operator_equal0() *QCameraViewfinderSettingsList {
	// QCameraViewfinderSettingsList_operator_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_operator_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator=(QList<T> &&)
func (this *QCameraViewfinderSettingsList) Operator_equal1() *QCameraViewfinderSettingsList {
	// QCameraViewfinderSettingsList_operator_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_operator_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// void swap(QList<T> &)
func (this *QCameraViewfinderSettingsList) Swap0() {
	// QCameraViewfinderSettingsList_swap_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_swap_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool operator==(const QList<T> &)
func (this *QCameraViewfinderSettingsList) Operator_equal_equal0() bool {
	// QCameraViewfinderSettingsList_operator_equal_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_operator_equal_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool operator!=(const QList<T> &)
func (this *QCameraViewfinderSettingsList) Operator_not_equal0() bool {
	// QCameraViewfinderSettingsList_operator_not_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_operator_not_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int size()
func (this *QCameraViewfinderSettingsList) Size0() int {
	// QCameraViewfinderSettingsList_size_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_size_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// void detach()
func (this *QCameraViewfinderSettingsList) Detach0() {
	// QCameraViewfinderSettingsList_detach_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_detach_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detachShared()
func (this *QCameraViewfinderSettingsList) DetachShared0() {
	// QCameraViewfinderSettingsList_detachShared_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_detachShared_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isDetached()
func (this *QCameraViewfinderSettingsList) IsDetached0() bool {
	// QCameraViewfinderSettingsList_isDetached_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_isDetached_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void setSharable(bool)
func (this *QCameraViewfinderSettingsList) SetSharable0() {
	// QCameraViewfinderSettingsList_setSharable_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_setSharable_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isSharedWith(const QList<T> &)
func (this *QCameraViewfinderSettingsList) IsSharedWith0() bool {
	// QCameraViewfinderSettingsList_isSharedWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_isSharedWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool isEmpty()
func (this *QCameraViewfinderSettingsList) IsEmpty0() bool {
	// QCameraViewfinderSettingsList_isEmpty_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_isEmpty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void clear()
func (this *QCameraViewfinderSettingsList) Clear0() {
	// QCameraViewfinderSettingsList_clear_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_clear_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// const T & at(int)
func (this *QCameraViewfinderSettingsList) At0() *QCameraViewfinderSettings {
	// QCameraViewfinderSettingsList_at_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_at_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QCameraViewfinderSettings{}
}

// const T & operator[](int)
func (this *QCameraViewfinderSettingsList) Operator_get_index0() *QCameraViewfinderSettings {
	// QCameraViewfinderSettingsList_operator_get_index_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_operator_get_index_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QCameraViewfinderSettings{}
}

// T & operator[](int)
func (this *QCameraViewfinderSettingsList) Operator_get_index1() *QCameraViewfinderSettings {
	// QCameraViewfinderSettingsList_operator_get_index_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_operator_get_index_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QCameraViewfinderSettings{}
}

// void reserve(int)
func (this *QCameraViewfinderSettingsList) Reserve0() {
	// QCameraViewfinderSettingsList_reserve_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_reserve_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const T &)
func (this *QCameraViewfinderSettingsList) Append0() {
	// QCameraViewfinderSettingsList_append_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_append_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const QList<T> &)
func (this *QCameraViewfinderSettingsList) Append1() {
	// QCameraViewfinderSettingsList_append_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_append_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void prepend(const T &)
func (this *QCameraViewfinderSettingsList) Prepend0() {
	// QCameraViewfinderSettingsList_prepend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_prepend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void insert(int, const T &)
func (this *QCameraViewfinderSettingsList) Insert0() {
	// QCameraViewfinderSettingsList_insert_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_insert_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void replace(int, const T &)
func (this *QCameraViewfinderSettingsList) Replace0() {
	// QCameraViewfinderSettingsList_replace_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_replace_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeAt(int)
func (this *QCameraViewfinderSettingsList) RemoveAt0() {
	// QCameraViewfinderSettingsList_removeAt_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_removeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int removeAll(const T &)
func (this *QCameraViewfinderSettingsList) RemoveAll0() int {
	// QCameraViewfinderSettingsList_removeAll_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_removeAll_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool removeOne(const T &)
func (this *QCameraViewfinderSettingsList) RemoveOne0() bool {
	// QCameraViewfinderSettingsList_removeOne_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_removeOne_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// T takeAt(int)
func (this *QCameraViewfinderSettingsList) TakeAt0() *QCameraViewfinderSettings {
	// QCameraViewfinderSettingsList_takeAt_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_takeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QCameraViewfinderSettings{}
}

// T takeFirst()
func (this *QCameraViewfinderSettingsList) TakeFirst0() *QCameraViewfinderSettings {
	// QCameraViewfinderSettingsList_takeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_takeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QCameraViewfinderSettings{}
}

// T takeLast()
func (this *QCameraViewfinderSettingsList) TakeLast0() *QCameraViewfinderSettings {
	// QCameraViewfinderSettingsList_takeLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_takeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QCameraViewfinderSettings{}
}

// void move(int, int)
func (this *QCameraViewfinderSettingsList) Move0() {
	// QCameraViewfinderSettingsList_move_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_move_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void swap(int, int)
func (this *QCameraViewfinderSettingsList) Swap1() {
	// QCameraViewfinderSettingsList_swap_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_swap_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int indexOf(const T &, int)
func (this *QCameraViewfinderSettingsList) IndexOf0() int {
	// QCameraViewfinderSettingsList_indexOf_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_indexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int lastIndexOf(const T &, int)
func (this *QCameraViewfinderSettingsList) LastIndexOf0() int {
	// QCameraViewfinderSettingsList_lastIndexOf_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_lastIndexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool contains(const T &)
func (this *QCameraViewfinderSettingsList) Contains0() bool {
	// QCameraViewfinderSettingsList_contains_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_contains_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count(const T &)
func (this *QCameraViewfinderSettingsList) Count0() int {
	// QCameraViewfinderSettingsList_count_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_count_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// QList::iterator begin()
func (this *QCameraViewfinderSettingsList) Begin0() {
	// QCameraViewfinderSettingsList_begin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_begin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator begin()
func (this *QCameraViewfinderSettingsList) Begin1() {
	// QCameraViewfinderSettingsList_begin_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_begin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cbegin()
func (this *QCameraViewfinderSettingsList) Cbegin0() {
	// QCameraViewfinderSettingsList_cbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_cbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constBegin()
func (this *QCameraViewfinderSettingsList) ConstBegin0() {
	// QCameraViewfinderSettingsList_constBegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_constBegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator end()
func (this *QCameraViewfinderSettingsList) End0() {
	// QCameraViewfinderSettingsList_end_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_end_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator end()
func (this *QCameraViewfinderSettingsList) End1() {
	// QCameraViewfinderSettingsList_end_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_end_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cend()
func (this *QCameraViewfinderSettingsList) Cend0() {
	// QCameraViewfinderSettingsList_cend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_cend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constEnd()
func (this *QCameraViewfinderSettingsList) ConstEnd0() {
	// QCameraViewfinderSettingsList_constEnd_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_constEnd_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rbegin()
func (this *QCameraViewfinderSettingsList) Rbegin0() {
	// QCameraViewfinderSettingsList_rbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_rbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rend()
func (this *QCameraViewfinderSettingsList) Rend0() {
	// QCameraViewfinderSettingsList_rend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_rend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rbegin()
func (this *QCameraViewfinderSettingsList) Rbegin1() {
	// QCameraViewfinderSettingsList_rbegin_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_rbegin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rend()
func (this *QCameraViewfinderSettingsList) Rend1() {
	// QCameraViewfinderSettingsList_rend_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_rend_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crbegin()
func (this *QCameraViewfinderSettingsList) Crbegin0() {
	// QCameraViewfinderSettingsList_crbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_crbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crend()
func (this *QCameraViewfinderSettingsList) Crend0() {
	// QCameraViewfinderSettingsList_crend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_crend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator insert(QList::iterator, const T &)
func (this *QCameraViewfinderSettingsList) Insert1() {
	// QCameraViewfinderSettingsList_insert_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_insert_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator)
func (this *QCameraViewfinderSettingsList) Erase0() {
	// QCameraViewfinderSettingsList_erase_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_erase_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator, QList::iterator)
func (this *QCameraViewfinderSettingsList) Erase1() {
	// QCameraViewfinderSettingsList_erase_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_erase_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int count()
func (this *QCameraViewfinderSettingsList) Count1() int {
	// QCameraViewfinderSettingsList_count_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_count_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int length()
func (this *QCameraViewfinderSettingsList) Length0() int {
	// QCameraViewfinderSettingsList_length_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_length_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// T & first()
func (this *QCameraViewfinderSettingsList) First0() *QCameraViewfinderSettings {
	// QCameraViewfinderSettingsList_first_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_first_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QCameraViewfinderSettings{}
}

// const T & constFirst()
func (this *QCameraViewfinderSettingsList) ConstFirst0() *QCameraViewfinderSettings {
	// QCameraViewfinderSettingsList_constFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_constFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QCameraViewfinderSettings{}
}

// const T & first()
func (this *QCameraViewfinderSettingsList) First1() *QCameraViewfinderSettings {
	// QCameraViewfinderSettingsList_first_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_first_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QCameraViewfinderSettings{}
}

// T & last()
func (this *QCameraViewfinderSettingsList) Last0() *QCameraViewfinderSettings {
	// QCameraViewfinderSettingsList_last_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_last_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QCameraViewfinderSettings{}
}

// const T & last()
func (this *QCameraViewfinderSettingsList) Last1() *QCameraViewfinderSettings {
	// QCameraViewfinderSettingsList_last_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_last_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QCameraViewfinderSettings{}
}

// const T & constLast()
func (this *QCameraViewfinderSettingsList) ConstLast0() *QCameraViewfinderSettings {
	// QCameraViewfinderSettingsList_constLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_constLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QCameraViewfinderSettings{}
}

// void removeFirst()
func (this *QCameraViewfinderSettingsList) RemoveFirst0() {
	// QCameraViewfinderSettingsList_removeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_removeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeLast()
func (this *QCameraViewfinderSettingsList) RemoveLast0() {
	// QCameraViewfinderSettingsList_removeLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_removeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool startsWith(const T &)
func (this *QCameraViewfinderSettingsList) StartsWith0() bool {
	// QCameraViewfinderSettingsList_startsWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_startsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool endsWith(const T &)
func (this *QCameraViewfinderSettingsList) EndsWith0() bool {
	// QCameraViewfinderSettingsList_endsWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_endsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> mid(int, int)
func (this *QCameraViewfinderSettingsList) Mid0() *QCameraViewfinderSettingsList {
	// QCameraViewfinderSettingsList_mid_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_mid_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// T value(int)
func (this *QCameraViewfinderSettingsList) Value0() *QCameraViewfinderSettings {
	// QCameraViewfinderSettingsList_value_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_value_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QCameraViewfinderSettings{}
}

// T value(int, const T &)
func (this *QCameraViewfinderSettingsList) Value1() *QCameraViewfinderSettings {
	// QCameraViewfinderSettingsList_value_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_value_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QCameraViewfinderSettings{}
}

// void push_back(const T &)
func (this *QCameraViewfinderSettingsList) Push_back0() {
	// QCameraViewfinderSettingsList_push_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_push_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void push_front(const T &)
func (this *QCameraViewfinderSettingsList) Push_front0() {
	// QCameraViewfinderSettingsList_push_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_push_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// T & front()
func (this *QCameraViewfinderSettingsList) Front0() *QCameraViewfinderSettings {
	// QCameraViewfinderSettingsList_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QCameraViewfinderSettings{}
}

// const T & front()
func (this *QCameraViewfinderSettingsList) Front1() *QCameraViewfinderSettings {
	// QCameraViewfinderSettingsList_front_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_front_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QCameraViewfinderSettings{}
}

// T & back()
func (this *QCameraViewfinderSettingsList) Back0() *QCameraViewfinderSettings {
	// QCameraViewfinderSettingsList_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QCameraViewfinderSettings{}
}

// const T & back()
func (this *QCameraViewfinderSettingsList) Back1() *QCameraViewfinderSettings {
	// QCameraViewfinderSettingsList_back_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_back_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QCameraViewfinderSettings{}
}

// void pop_front()
func (this *QCameraViewfinderSettingsList) Pop_front0() {
	// QCameraViewfinderSettingsList_pop_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_pop_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void pop_back()
func (this *QCameraViewfinderSettingsList) Pop_back0() {
	// QCameraViewfinderSettingsList_pop_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_pop_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool empty()
func (this *QCameraViewfinderSettingsList) Empty0() bool {
	// QCameraViewfinderSettingsList_empty_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_empty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> & operator+=(const QList<T> &)
func (this *QCameraViewfinderSettingsList) Operator_add_equal0() *QCameraViewfinderSettingsList {
	// QCameraViewfinderSettingsList_operator_add_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_operator_add_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> operator+(const QList<T> &)
func (this *QCameraViewfinderSettingsList) Operator_add0() *QCameraViewfinderSettingsList {
	// QCameraViewfinderSettingsList_operator_add_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_operator_add_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator+=(const T &)
func (this *QCameraViewfinderSettingsList) Operator_add_equal1() *QCameraViewfinderSettingsList {
	// QCameraViewfinderSettingsList_operator_add_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_operator_add_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const T &)
func (this *QCameraViewfinderSettingsList) Operator_left_shift0() *QCameraViewfinderSettingsList {
	// QCameraViewfinderSettingsList_operator_left_shift_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_operator_left_shift_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const QList<T> &)
func (this *QCameraViewfinderSettingsList) Operator_left_shift1() *QCameraViewfinderSettingsList {
	// QCameraViewfinderSettingsList_operator_left_shift_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_operator_left_shift_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QVector<T> toVector()
func (this *QCameraViewfinderSettingsList) ToVector0() {
	// QCameraViewfinderSettingsList_toVector_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_toVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QSet<T> toSet()
func (this *QCameraViewfinderSettingsList) ToSet0() {
	// QCameraViewfinderSettingsList_toSet_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_toSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList<T> fromVector(const QVector<T> &)
func (this *QCameraViewfinderSettingsList) FromVector0() *QCameraViewfinderSettingsList {
	// QCameraViewfinderSettingsList_fromVector_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_fromVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromSet(const QSet<T> &)
func (this *QCameraViewfinderSettingsList) FromSet0() *QCameraViewfinderSettingsList {
	// QCameraViewfinderSettingsList_fromSet_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_fromSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromStdList(const std::list<T> &)
func (this *QCameraViewfinderSettingsList) FromStdList0() *QCameraViewfinderSettingsList {
	// QCameraViewfinderSettingsList_fromStdList_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_fromStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// std::list<T> toStdList()
func (this *QCameraViewfinderSettingsList) ToStdList0() {
	// QCameraViewfinderSettingsList_toStdList_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_toStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::Node * detach_helper_grow(int, int)
func (this *QCameraViewfinderSettingsList) Detach_helper_grow0() {
	// QCameraViewfinderSettingsList_detach_helper_grow_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_detach_helper_grow_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper(int)
func (this *QCameraViewfinderSettingsList) Detach_helper0() {
	// QCameraViewfinderSettingsList_detach_helper_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_detach_helper_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper()
func (this *QCameraViewfinderSettingsList) Detach_helper1() {
	// QCameraViewfinderSettingsList_detach_helper_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_detach_helper_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void dealloc(QListData::Data *)
func (this *QCameraViewfinderSettingsList) Dealloc0() {
	// QCameraViewfinderSettingsList_dealloc_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_dealloc_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_construct(QList::Node *, const T &)
func (this *QCameraViewfinderSettingsList) Node_construct0() {
	// QCameraViewfinderSettingsList_node_construct_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_node_construct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *)
func (this *QCameraViewfinderSettingsList) Node_destruct0() {
	// QCameraViewfinderSettingsList_node_destruct_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_node_destruct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_copy(QList::Node *, QList::Node *, QList::Node *)
func (this *QCameraViewfinderSettingsList) Node_copy0() {
	// QCameraViewfinderSettingsList_node_copy_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_node_copy_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *, QList::Node *)
func (this *QCameraViewfinderSettingsList) Node_destruct1() {
	// QCameraViewfinderSettingsList_node_destruct_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_node_destruct_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isValidIterator(const QList::iterator &)
func (this *QCameraViewfinderSettingsList) IsValidIterator0() bool {
	// QCameraViewfinderSettingsList_isValidIterator_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_isValidIterator_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::NotArrayCompatibleLayout)
func (this *QCameraViewfinderSettingsList) Op_eq_impl0() bool {
	// QCameraViewfinderSettingsList_op_eq_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_op_eq_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::ArrayCompatibleLayout)
func (this *QCameraViewfinderSettingsList) Op_eq_impl1() bool {
	// QCameraViewfinderSettingsList_op_eq_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_op_eq_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QCameraViewfinderSettingsList) Contains_impl0() bool {
	// QCameraViewfinderSettingsList_contains_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_contains_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QCameraViewfinderSettingsList) Contains_impl1() bool {
	// QCameraViewfinderSettingsList_contains_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_contains_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QCameraViewfinderSettingsList) Count_impl0() int {
	// QCameraViewfinderSettingsList_count_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_count_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int count_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QCameraViewfinderSettingsList) Count_impl1() int {
	// QCameraViewfinderSettingsList_count_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraViewfinderSettingsList_count_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

//  body block end
