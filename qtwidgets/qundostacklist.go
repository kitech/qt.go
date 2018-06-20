package qtwidgets

// /usr/include/qt/QtWidgets/qundostack.h
// #include <qundostack.h>
// #include <QtWidgets>

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
}

//  keep block end

//  body block begin
type QUndoStackList struct {
	*qtrt.CObject
}

// QList<T> & operator=(const QList<T> &)
func (this *QUndoStackList) Operator_equal_0() *QUndoStackList {
	// QUndoStackList_operator_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_operator_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator=(QList<T> &&)
func (this *QUndoStackList) Operator_equal_1() *QUndoStackList {
	// QUndoStackList_operator_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_operator_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// void swap(QList<T> &)
func (this *QUndoStackList) Swap_0() {
	// QUndoStackList_swap_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_swap_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool operator==(const QList<T> &)
func (this *QUndoStackList) Operator_equal_equal_0() bool {
	// QUndoStackList_operator_equal_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_operator_equal_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool operator!=(const QList<T> &)
func (this *QUndoStackList) Operator_not_equal_0() bool {
	// QUndoStackList_operator_not_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_operator_not_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int size()
func (this *QUndoStackList) Size_0() int {
	// QUndoStackList_size_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_size_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// void detach()
func (this *QUndoStackList) Detach_0() {
	// QUndoStackList_detach_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_detach_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detachShared()
func (this *QUndoStackList) DetachShared_0() {
	// QUndoStackList_detachShared_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_detachShared_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isDetached()
func (this *QUndoStackList) IsDetached_0() bool {
	// QUndoStackList_isDetached_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_isDetached_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void setSharable(bool)
func (this *QUndoStackList) SetSharable_0() {
	// QUndoStackList_setSharable_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_setSharable_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isSharedWith(const QList<T> &)
func (this *QUndoStackList) IsSharedWith_0() bool {
	// QUndoStackList_isSharedWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_isSharedWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool isEmpty()
func (this *QUndoStackList) IsEmpty_0() bool {
	// QUndoStackList_isEmpty_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_isEmpty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void clear()
func (this *QUndoStackList) Clear_0() {
	// QUndoStackList_clear_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_clear_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// const T & at(int)
func (this *QUndoStackList) At_0() *QUndoStack {
	// QUndoStackList_at_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_at_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QUndoStack{}
}

// const T & operator[](int)
func (this *QUndoStackList) Operator_get_index_0() *QUndoStack {
	// QUndoStackList_operator_get_index_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_operator_get_index_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QUndoStack{}
}

// T & operator[](int)
func (this *QUndoStackList) Operator_get_index_1() *QUndoStack {
	// QUndoStackList_operator_get_index_1()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_operator_get_index_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QUndoStack{}
}

// void reserve(int)
func (this *QUndoStackList) Reserve_0() {
	// QUndoStackList_reserve_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_reserve_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const T &)
func (this *QUndoStackList) Append_0() {
	// QUndoStackList_append_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_append_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const QList<T> &)
func (this *QUndoStackList) Append_1() {
	// QUndoStackList_append_1()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_append_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void prepend(const T &)
func (this *QUndoStackList) Prepend_0() {
	// QUndoStackList_prepend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_prepend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void insert(int, const T &)
func (this *QUndoStackList) Insert_0() {
	// QUndoStackList_insert_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_insert_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void replace(int, const T &)
func (this *QUndoStackList) Replace_0() {
	// QUndoStackList_replace_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_replace_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeAt(int)
func (this *QUndoStackList) RemoveAt_0() {
	// QUndoStackList_removeAt_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_removeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int removeAll(const T &)
func (this *QUndoStackList) RemoveAll_0() int {
	// QUndoStackList_removeAll_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_removeAll_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool removeOne(const T &)
func (this *QUndoStackList) RemoveOne_0() bool {
	// QUndoStackList_removeOne_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_removeOne_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// T takeAt(int)
func (this *QUndoStackList) TakeAt_0() *QUndoStack {
	// QUndoStackList_takeAt_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_takeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QUndoStack{}
}

// T takeFirst()
func (this *QUndoStackList) TakeFirst_0() *QUndoStack {
	// QUndoStackList_takeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_takeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QUndoStack{}
}

// T takeLast()
func (this *QUndoStackList) TakeLast_0() *QUndoStack {
	// QUndoStackList_takeLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_takeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QUndoStack{}
}

// void move(int, int)
func (this *QUndoStackList) Move_0() {
	// QUndoStackList_move_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_move_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void swap(int, int)
func (this *QUndoStackList) Swap_1() {
	// QUndoStackList_swap_1()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_swap_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int indexOf(const T &, int)
func (this *QUndoStackList) IndexOf_0() int {
	// QUndoStackList_indexOf_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_indexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int lastIndexOf(const T &, int)
func (this *QUndoStackList) LastIndexOf_0() int {
	// QUndoStackList_lastIndexOf_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_lastIndexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool contains(const T &)
func (this *QUndoStackList) Contains_0() bool {
	// QUndoStackList_contains_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_contains_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count(const T &)
func (this *QUndoStackList) Count_0() int {
	// QUndoStackList_count_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_count_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// QList::iterator begin()
func (this *QUndoStackList) Begin_0() {
	// QUndoStackList_begin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_begin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator begin()
func (this *QUndoStackList) Begin_1() {
	// QUndoStackList_begin_1()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_begin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cbegin()
func (this *QUndoStackList) Cbegin_0() {
	// QUndoStackList_cbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_cbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constBegin()
func (this *QUndoStackList) ConstBegin_0() {
	// QUndoStackList_constBegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_constBegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator end()
func (this *QUndoStackList) End_0() {
	// QUndoStackList_end_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_end_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator end()
func (this *QUndoStackList) End_1() {
	// QUndoStackList_end_1()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_end_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cend()
func (this *QUndoStackList) Cend_0() {
	// QUndoStackList_cend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_cend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constEnd()
func (this *QUndoStackList) ConstEnd_0() {
	// QUndoStackList_constEnd_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_constEnd_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rbegin()
func (this *QUndoStackList) Rbegin_0() {
	// QUndoStackList_rbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_rbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rend()
func (this *QUndoStackList) Rend_0() {
	// QUndoStackList_rend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_rend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rbegin()
func (this *QUndoStackList) Rbegin_1() {
	// QUndoStackList_rbegin_1()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_rbegin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rend()
func (this *QUndoStackList) Rend_1() {
	// QUndoStackList_rend_1()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_rend_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crbegin()
func (this *QUndoStackList) Crbegin_0() {
	// QUndoStackList_crbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_crbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crend()
func (this *QUndoStackList) Crend_0() {
	// QUndoStackList_crend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_crend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator insert(QList::iterator, const T &)
func (this *QUndoStackList) Insert_1() {
	// QUndoStackList_insert_1()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_insert_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator)
func (this *QUndoStackList) Erase_0() {
	// QUndoStackList_erase_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_erase_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator, QList::iterator)
func (this *QUndoStackList) Erase_1() {
	// QUndoStackList_erase_1()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_erase_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int count()
func (this *QUndoStackList) Count_1() int {
	// QUndoStackList_count_1()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_count_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int length()
func (this *QUndoStackList) Length_0() int {
	// QUndoStackList_length_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_length_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// T & first()
func (this *QUndoStackList) First_0() *QUndoStack {
	// QUndoStackList_first_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_first_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QUndoStack{}
}

// const T & constFirst()
func (this *QUndoStackList) ConstFirst_0() *QUndoStack {
	// QUndoStackList_constFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_constFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QUndoStack{}
}

// const T & first()
func (this *QUndoStackList) First_1() *QUndoStack {
	// QUndoStackList_first_1()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_first_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QUndoStack{}
}

// T & last()
func (this *QUndoStackList) Last_0() *QUndoStack {
	// QUndoStackList_last_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_last_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QUndoStack{}
}

// const T & last()
func (this *QUndoStackList) Last_1() *QUndoStack {
	// QUndoStackList_last_1()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_last_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QUndoStack{}
}

// const T & constLast()
func (this *QUndoStackList) ConstLast_0() *QUndoStack {
	// QUndoStackList_constLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_constLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QUndoStack{}
}

// void removeFirst()
func (this *QUndoStackList) RemoveFirst_0() {
	// QUndoStackList_removeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_removeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeLast()
func (this *QUndoStackList) RemoveLast_0() {
	// QUndoStackList_removeLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_removeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool startsWith(const T &)
func (this *QUndoStackList) StartsWith_0() bool {
	// QUndoStackList_startsWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_startsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool endsWith(const T &)
func (this *QUndoStackList) EndsWith_0() bool {
	// QUndoStackList_endsWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_endsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> mid(int, int)
func (this *QUndoStackList) Mid_0() *QUndoStackList {
	// QUndoStackList_mid_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_mid_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// T value(int)
func (this *QUndoStackList) Value_0() *QUndoStack {
	// QUndoStackList_value_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_value_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QUndoStack{}
}

// T value(int, const T &)
func (this *QUndoStackList) Value_1() *QUndoStack {
	// QUndoStackList_value_1()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_value_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QUndoStack{}
}

// void push_back(const T &)
func (this *QUndoStackList) Push_back_0() {
	// QUndoStackList_push_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_push_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void push_front(const T &)
func (this *QUndoStackList) Push_front_0() {
	// QUndoStackList_push_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_push_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// T & front()
func (this *QUndoStackList) Front_0() *QUndoStack {
	// QUndoStackList_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QUndoStack{}
}

// const T & front()
func (this *QUndoStackList) Front_1() *QUndoStack {
	// QUndoStackList_front_1()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_front_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QUndoStack{}
}

// T & back()
func (this *QUndoStackList) Back_0() *QUndoStack {
	// QUndoStackList_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QUndoStack{}
}

// const T & back()
func (this *QUndoStackList) Back_1() *QUndoStack {
	// QUndoStackList_back_1()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_back_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QUndoStack{}
}

// void pop_front()
func (this *QUndoStackList) Pop_front_0() {
	// QUndoStackList_pop_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_pop_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void pop_back()
func (this *QUndoStackList) Pop_back_0() {
	// QUndoStackList_pop_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_pop_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool empty()
func (this *QUndoStackList) Empty_0() bool {
	// QUndoStackList_empty_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_empty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> & operator+=(const QList<T> &)
func (this *QUndoStackList) Operator_add_equal_0() *QUndoStackList {
	// QUndoStackList_operator_add_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_operator_add_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> operator+(const QList<T> &)
func (this *QUndoStackList) Operator_add_0() *QUndoStackList {
	// QUndoStackList_operator_add_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_operator_add_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator+=(const T &)
func (this *QUndoStackList) Operator_add_equal_1() *QUndoStackList {
	// QUndoStackList_operator_add_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_operator_add_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const T &)
func (this *QUndoStackList) Operator_left_shift_0() *QUndoStackList {
	// QUndoStackList_operator_left_shift_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_operator_left_shift_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const QList<T> &)
func (this *QUndoStackList) Operator_left_shift_1() *QUndoStackList {
	// QUndoStackList_operator_left_shift_1()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_operator_left_shift_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QVector<T> toVector()
func (this *QUndoStackList) ToVector_0() {
	// QUndoStackList_toVector_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_toVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QSet<T> toSet()
func (this *QUndoStackList) ToSet_0() {
	// QUndoStackList_toSet_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_toSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList<T> fromVector(const QVector<T> &)
func (this *QUndoStackList) FromVector_0() *QUndoStackList {
	// QUndoStackList_fromVector_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_fromVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromSet(const QSet<T> &)
func (this *QUndoStackList) FromSet_0() *QUndoStackList {
	// QUndoStackList_fromSet_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_fromSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromStdList(const std::list<T> &)
func (this *QUndoStackList) FromStdList_0() *QUndoStackList {
	// QUndoStackList_fromStdList_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_fromStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// std::list<T> toStdList()
func (this *QUndoStackList) ToStdList_0() {
	// QUndoStackList_toStdList_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_toStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::Node * detach_helper_grow(int, int)
func (this *QUndoStackList) Detach_helper_grow_0() {
	// QUndoStackList_detach_helper_grow_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_detach_helper_grow_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper(int)
func (this *QUndoStackList) Detach_helper_0() {
	// QUndoStackList_detach_helper_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_detach_helper_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper()
func (this *QUndoStackList) Detach_helper_1() {
	// QUndoStackList_detach_helper_1()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_detach_helper_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void dealloc(QListData::Data *)
func (this *QUndoStackList) Dealloc_0() {
	// QUndoStackList_dealloc_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_dealloc_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_construct(QList::Node *, const T &)
func (this *QUndoStackList) Node_construct_0() {
	// QUndoStackList_node_construct_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_node_construct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *)
func (this *QUndoStackList) Node_destruct_0() {
	// QUndoStackList_node_destruct_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_node_destruct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_copy(QList::Node *, QList::Node *, QList::Node *)
func (this *QUndoStackList) Node_copy_0() {
	// QUndoStackList_node_copy_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_node_copy_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *, QList::Node *)
func (this *QUndoStackList) Node_destruct_1() {
	// QUndoStackList_node_destruct_1()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_node_destruct_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isValidIterator(const QList::iterator &)
func (this *QUndoStackList) IsValidIterator_0() bool {
	// QUndoStackList_isValidIterator_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_isValidIterator_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::NotArrayCompatibleLayout)
func (this *QUndoStackList) Op_eq_impl_0() bool {
	// QUndoStackList_op_eq_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_op_eq_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::ArrayCompatibleLayout)
func (this *QUndoStackList) Op_eq_impl_1() bool {
	// QUndoStackList_op_eq_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_op_eq_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QUndoStackList) Contains_impl_0() bool {
	// QUndoStackList_contains_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_contains_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QUndoStackList) Contains_impl_1() bool {
	// QUndoStackList_contains_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_contains_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QUndoStackList) Count_impl_0() int {
	// QUndoStackList_count_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_count_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int count_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QUndoStackList) Count_impl_1() int {
	// QUndoStackList_count_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QUndoStackList_count_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

//  body block end
