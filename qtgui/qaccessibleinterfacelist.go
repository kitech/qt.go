package qtgui

// /usr/include/qt/QtGui/qaccessible.h
// #include <qaccessible.h>
// #include <QtGui>

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
}

//  keep block end

//  body block begin
type QAccessibleInterfaceList struct {
	*qtrt.CObject
}

// QList<T> & operator=(const QList<T> &)
func (this *QAccessibleInterfaceList) Operator_equal0() *QAccessibleInterfaceList {
	// QAccessibleInterfaceList_operator_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_operator_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator=(QList<T> &&)
func (this *QAccessibleInterfaceList) Operator_equal1() *QAccessibleInterfaceList {
	// QAccessibleInterfaceList_operator_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_operator_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// void swap(QList<T> &)
func (this *QAccessibleInterfaceList) Swap0() {
	// QAccessibleInterfaceList_swap_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_swap_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool operator==(const QList<T> &)
func (this *QAccessibleInterfaceList) Operator_equal_equal0() bool {
	// QAccessibleInterfaceList_operator_equal_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_operator_equal_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool operator!=(const QList<T> &)
func (this *QAccessibleInterfaceList) Operator_not_equal0() bool {
	// QAccessibleInterfaceList_operator_not_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_operator_not_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int size()
func (this *QAccessibleInterfaceList) Size0() int {
	// QAccessibleInterfaceList_size_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_size_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// void detach()
func (this *QAccessibleInterfaceList) Detach0() {
	// QAccessibleInterfaceList_detach_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_detach_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detachShared()
func (this *QAccessibleInterfaceList) DetachShared0() {
	// QAccessibleInterfaceList_detachShared_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_detachShared_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isDetached()
func (this *QAccessibleInterfaceList) IsDetached0() bool {
	// QAccessibleInterfaceList_isDetached_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_isDetached_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void setSharable(bool)
func (this *QAccessibleInterfaceList) SetSharable0() {
	// QAccessibleInterfaceList_setSharable_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_setSharable_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isSharedWith(const QList<T> &)
func (this *QAccessibleInterfaceList) IsSharedWith0() bool {
	// QAccessibleInterfaceList_isSharedWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_isSharedWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool isEmpty()
func (this *QAccessibleInterfaceList) IsEmpty0() bool {
	// QAccessibleInterfaceList_isEmpty_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_isEmpty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void clear()
func (this *QAccessibleInterfaceList) Clear0() {
	// QAccessibleInterfaceList_clear_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_clear_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// const T & at(int)
func (this *QAccessibleInterfaceList) At0() *QAccessibleInterface {
	// QAccessibleInterfaceList_at_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_at_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAccessibleInterface{}
}

// const T & operator[](int)
func (this *QAccessibleInterfaceList) Operator_get_index0() *QAccessibleInterface {
	// QAccessibleInterfaceList_operator_get_index_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_operator_get_index_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAccessibleInterface{}
}

// T & operator[](int)
func (this *QAccessibleInterfaceList) Operator_get_index1() *QAccessibleInterface {
	// QAccessibleInterfaceList_operator_get_index_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_operator_get_index_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAccessibleInterface{}
}

// void reserve(int)
func (this *QAccessibleInterfaceList) Reserve0() {
	// QAccessibleInterfaceList_reserve_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_reserve_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const T &)
func (this *QAccessibleInterfaceList) Append0() {
	// QAccessibleInterfaceList_append_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_append_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const QList<T> &)
func (this *QAccessibleInterfaceList) Append1() {
	// QAccessibleInterfaceList_append_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_append_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void prepend(const T &)
func (this *QAccessibleInterfaceList) Prepend0() {
	// QAccessibleInterfaceList_prepend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_prepend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void insert(int, const T &)
func (this *QAccessibleInterfaceList) Insert0() {
	// QAccessibleInterfaceList_insert_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_insert_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void replace(int, const T &)
func (this *QAccessibleInterfaceList) Replace0() {
	// QAccessibleInterfaceList_replace_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_replace_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeAt(int)
func (this *QAccessibleInterfaceList) RemoveAt0() {
	// QAccessibleInterfaceList_removeAt_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_removeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int removeAll(const T &)
func (this *QAccessibleInterfaceList) RemoveAll0() int {
	// QAccessibleInterfaceList_removeAll_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_removeAll_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool removeOne(const T &)
func (this *QAccessibleInterfaceList) RemoveOne0() bool {
	// QAccessibleInterfaceList_removeOne_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_removeOne_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// T takeAt(int)
func (this *QAccessibleInterfaceList) TakeAt0() *QAccessibleInterface {
	// QAccessibleInterfaceList_takeAt_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_takeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAccessibleInterface{}
}

// T takeFirst()
func (this *QAccessibleInterfaceList) TakeFirst0() *QAccessibleInterface {
	// QAccessibleInterfaceList_takeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_takeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAccessibleInterface{}
}

// T takeLast()
func (this *QAccessibleInterfaceList) TakeLast0() *QAccessibleInterface {
	// QAccessibleInterfaceList_takeLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_takeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAccessibleInterface{}
}

// void move(int, int)
func (this *QAccessibleInterfaceList) Move0() {
	// QAccessibleInterfaceList_move_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_move_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void swap(int, int)
func (this *QAccessibleInterfaceList) Swap1() {
	// QAccessibleInterfaceList_swap_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_swap_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int indexOf(const T &, int)
func (this *QAccessibleInterfaceList) IndexOf0() int {
	// QAccessibleInterfaceList_indexOf_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_indexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int lastIndexOf(const T &, int)
func (this *QAccessibleInterfaceList) LastIndexOf0() int {
	// QAccessibleInterfaceList_lastIndexOf_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_lastIndexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool contains(const T &)
func (this *QAccessibleInterfaceList) Contains0() bool {
	// QAccessibleInterfaceList_contains_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_contains_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count(const T &)
func (this *QAccessibleInterfaceList) Count0() int {
	// QAccessibleInterfaceList_count_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_count_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// QList::iterator begin()
func (this *QAccessibleInterfaceList) Begin0() {
	// QAccessibleInterfaceList_begin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_begin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator begin()
func (this *QAccessibleInterfaceList) Begin1() {
	// QAccessibleInterfaceList_begin_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_begin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cbegin()
func (this *QAccessibleInterfaceList) Cbegin0() {
	// QAccessibleInterfaceList_cbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_cbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constBegin()
func (this *QAccessibleInterfaceList) ConstBegin0() {
	// QAccessibleInterfaceList_constBegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_constBegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator end()
func (this *QAccessibleInterfaceList) End0() {
	// QAccessibleInterfaceList_end_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_end_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator end()
func (this *QAccessibleInterfaceList) End1() {
	// QAccessibleInterfaceList_end_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_end_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cend()
func (this *QAccessibleInterfaceList) Cend0() {
	// QAccessibleInterfaceList_cend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_cend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constEnd()
func (this *QAccessibleInterfaceList) ConstEnd0() {
	// QAccessibleInterfaceList_constEnd_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_constEnd_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rbegin()
func (this *QAccessibleInterfaceList) Rbegin0() {
	// QAccessibleInterfaceList_rbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_rbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rend()
func (this *QAccessibleInterfaceList) Rend0() {
	// QAccessibleInterfaceList_rend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_rend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rbegin()
func (this *QAccessibleInterfaceList) Rbegin1() {
	// QAccessibleInterfaceList_rbegin_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_rbegin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rend()
func (this *QAccessibleInterfaceList) Rend1() {
	// QAccessibleInterfaceList_rend_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_rend_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crbegin()
func (this *QAccessibleInterfaceList) Crbegin0() {
	// QAccessibleInterfaceList_crbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_crbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crend()
func (this *QAccessibleInterfaceList) Crend0() {
	// QAccessibleInterfaceList_crend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_crend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator insert(QList::iterator, const T &)
func (this *QAccessibleInterfaceList) Insert1() {
	// QAccessibleInterfaceList_insert_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_insert_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator)
func (this *QAccessibleInterfaceList) Erase0() {
	// QAccessibleInterfaceList_erase_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_erase_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator, QList::iterator)
func (this *QAccessibleInterfaceList) Erase1() {
	// QAccessibleInterfaceList_erase_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_erase_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int count()
func (this *QAccessibleInterfaceList) Count1() int {
	// QAccessibleInterfaceList_count_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_count_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int length()
func (this *QAccessibleInterfaceList) Length0() int {
	// QAccessibleInterfaceList_length_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_length_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// T & first()
func (this *QAccessibleInterfaceList) First0() *QAccessibleInterface {
	// QAccessibleInterfaceList_first_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_first_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAccessibleInterface{}
}

// const T & constFirst()
func (this *QAccessibleInterfaceList) ConstFirst0() *QAccessibleInterface {
	// QAccessibleInterfaceList_constFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_constFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAccessibleInterface{}
}

// const T & first()
func (this *QAccessibleInterfaceList) First1() *QAccessibleInterface {
	// QAccessibleInterfaceList_first_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_first_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAccessibleInterface{}
}

// T & last()
func (this *QAccessibleInterfaceList) Last0() *QAccessibleInterface {
	// QAccessibleInterfaceList_last_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_last_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAccessibleInterface{}
}

// const T & last()
func (this *QAccessibleInterfaceList) Last1() *QAccessibleInterface {
	// QAccessibleInterfaceList_last_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_last_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAccessibleInterface{}
}

// const T & constLast()
func (this *QAccessibleInterfaceList) ConstLast0() *QAccessibleInterface {
	// QAccessibleInterfaceList_constLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_constLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAccessibleInterface{}
}

// void removeFirst()
func (this *QAccessibleInterfaceList) RemoveFirst0() {
	// QAccessibleInterfaceList_removeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_removeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeLast()
func (this *QAccessibleInterfaceList) RemoveLast0() {
	// QAccessibleInterfaceList_removeLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_removeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool startsWith(const T &)
func (this *QAccessibleInterfaceList) StartsWith0() bool {
	// QAccessibleInterfaceList_startsWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_startsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool endsWith(const T &)
func (this *QAccessibleInterfaceList) EndsWith0() bool {
	// QAccessibleInterfaceList_endsWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_endsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> mid(int, int)
func (this *QAccessibleInterfaceList) Mid0() *QAccessibleInterfaceList {
	// QAccessibleInterfaceList_mid_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_mid_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// T value(int)
func (this *QAccessibleInterfaceList) Value0() *QAccessibleInterface {
	// QAccessibleInterfaceList_value_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_value_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAccessibleInterface{}
}

// T value(int, const T &)
func (this *QAccessibleInterfaceList) Value1() *QAccessibleInterface {
	// QAccessibleInterfaceList_value_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_value_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAccessibleInterface{}
}

// void push_back(const T &)
func (this *QAccessibleInterfaceList) Push_back0() {
	// QAccessibleInterfaceList_push_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_push_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void push_front(const T &)
func (this *QAccessibleInterfaceList) Push_front0() {
	// QAccessibleInterfaceList_push_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_push_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// T & front()
func (this *QAccessibleInterfaceList) Front0() *QAccessibleInterface {
	// QAccessibleInterfaceList_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAccessibleInterface{}
}

// const T & front()
func (this *QAccessibleInterfaceList) Front1() *QAccessibleInterface {
	// QAccessibleInterfaceList_front_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_front_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAccessibleInterface{}
}

// T & back()
func (this *QAccessibleInterfaceList) Back0() *QAccessibleInterface {
	// QAccessibleInterfaceList_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAccessibleInterface{}
}

// const T & back()
func (this *QAccessibleInterfaceList) Back1() *QAccessibleInterface {
	// QAccessibleInterfaceList_back_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_back_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAccessibleInterface{}
}

// void pop_front()
func (this *QAccessibleInterfaceList) Pop_front0() {
	// QAccessibleInterfaceList_pop_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_pop_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void pop_back()
func (this *QAccessibleInterfaceList) Pop_back0() {
	// QAccessibleInterfaceList_pop_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_pop_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool empty()
func (this *QAccessibleInterfaceList) Empty0() bool {
	// QAccessibleInterfaceList_empty_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_empty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> & operator+=(const QList<T> &)
func (this *QAccessibleInterfaceList) Operator_add_equal0() *QAccessibleInterfaceList {
	// QAccessibleInterfaceList_operator_add_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_operator_add_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> operator+(const QList<T> &)
func (this *QAccessibleInterfaceList) Operator_add0() *QAccessibleInterfaceList {
	// QAccessibleInterfaceList_operator_add_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_operator_add_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator+=(const T &)
func (this *QAccessibleInterfaceList) Operator_add_equal1() *QAccessibleInterfaceList {
	// QAccessibleInterfaceList_operator_add_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_operator_add_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const T &)
func (this *QAccessibleInterfaceList) Operator_left_shift0() *QAccessibleInterfaceList {
	// QAccessibleInterfaceList_operator_left_shift_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_operator_left_shift_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const QList<T> &)
func (this *QAccessibleInterfaceList) Operator_left_shift1() *QAccessibleInterfaceList {
	// QAccessibleInterfaceList_operator_left_shift_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_operator_left_shift_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QVector<T> toVector()
func (this *QAccessibleInterfaceList) ToVector0() {
	// QAccessibleInterfaceList_toVector_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_toVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QSet<T> toSet()
func (this *QAccessibleInterfaceList) ToSet0() {
	// QAccessibleInterfaceList_toSet_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_toSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList<T> fromVector(const QVector<T> &)
func (this *QAccessibleInterfaceList) FromVector0() *QAccessibleInterfaceList {
	// QAccessibleInterfaceList_fromVector_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_fromVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromSet(const QSet<T> &)
func (this *QAccessibleInterfaceList) FromSet0() *QAccessibleInterfaceList {
	// QAccessibleInterfaceList_fromSet_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_fromSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromStdList(const std::list<T> &)
func (this *QAccessibleInterfaceList) FromStdList0() *QAccessibleInterfaceList {
	// QAccessibleInterfaceList_fromStdList_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_fromStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// std::list<T> toStdList()
func (this *QAccessibleInterfaceList) ToStdList0() {
	// QAccessibleInterfaceList_toStdList_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_toStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::Node * detach_helper_grow(int, int)
func (this *QAccessibleInterfaceList) Detach_helper_grow0() {
	// QAccessibleInterfaceList_detach_helper_grow_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_detach_helper_grow_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper(int)
func (this *QAccessibleInterfaceList) Detach_helper0() {
	// QAccessibleInterfaceList_detach_helper_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_detach_helper_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper()
func (this *QAccessibleInterfaceList) Detach_helper1() {
	// QAccessibleInterfaceList_detach_helper_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_detach_helper_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void dealloc(QListData::Data *)
func (this *QAccessibleInterfaceList) Dealloc0() {
	// QAccessibleInterfaceList_dealloc_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_dealloc_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_construct(QList::Node *, const T &)
func (this *QAccessibleInterfaceList) Node_construct0() {
	// QAccessibleInterfaceList_node_construct_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_node_construct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *)
func (this *QAccessibleInterfaceList) Node_destruct0() {
	// QAccessibleInterfaceList_node_destruct_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_node_destruct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_copy(QList::Node *, QList::Node *, QList::Node *)
func (this *QAccessibleInterfaceList) Node_copy0() {
	// QAccessibleInterfaceList_node_copy_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_node_copy_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *, QList::Node *)
func (this *QAccessibleInterfaceList) Node_destruct1() {
	// QAccessibleInterfaceList_node_destruct_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_node_destruct_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isValidIterator(const QList::iterator &)
func (this *QAccessibleInterfaceList) IsValidIterator0() bool {
	// QAccessibleInterfaceList_isValidIterator_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_isValidIterator_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::NotArrayCompatibleLayout)
func (this *QAccessibleInterfaceList) Op_eq_impl0() bool {
	// QAccessibleInterfaceList_op_eq_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_op_eq_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::ArrayCompatibleLayout)
func (this *QAccessibleInterfaceList) Op_eq_impl1() bool {
	// QAccessibleInterfaceList_op_eq_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_op_eq_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QAccessibleInterfaceList) Contains_impl0() bool {
	// QAccessibleInterfaceList_contains_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_contains_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QAccessibleInterfaceList) Contains_impl1() bool {
	// QAccessibleInterfaceList_contains_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_contains_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QAccessibleInterfaceList) Count_impl0() int {
	// QAccessibleInterfaceList_count_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_count_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int count_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QAccessibleInterfaceList) Count_impl1() int {
	// QAccessibleInterfaceList_count_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAccessibleInterfaceList_count_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

//  body block end
