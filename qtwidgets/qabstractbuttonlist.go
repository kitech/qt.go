package qtwidgets

// /usr/include/qt/QtWidgets/qabstractbutton.h
// #include <qabstractbutton.h>
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
type QAbstractButtonList struct {
	*qtrt.CObject
}

// QList<T> & operator=(const QList<T> &)
func (this *QAbstractButtonList) Operator_equal_0() *QAbstractButtonList {
	// QAbstractButtonList_operator_equal_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_operator_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator=(QList<T> &&)
func (this *QAbstractButtonList) Operator_equal_1() *QAbstractButtonList {
	// QAbstractButtonList_operator_equal_1()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_operator_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// void swap(QList<T> &)
func (this *QAbstractButtonList) Swap_0() {
	// QAbstractButtonList_swap_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_swap_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool operator==(const QList<T> &)
func (this *QAbstractButtonList) Operator_equal_equal_0() bool {
	// QAbstractButtonList_operator_equal_equal_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_operator_equal_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool operator!=(const QList<T> &)
func (this *QAbstractButtonList) Operator_not_equal_0() bool {
	// QAbstractButtonList_operator_not_equal_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_operator_not_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int size()
func (this *QAbstractButtonList) Size_0() int {
	// QAbstractButtonList_size_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_size_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// void detach()
func (this *QAbstractButtonList) Detach_0() {
	// QAbstractButtonList_detach_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_detach_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detachShared()
func (this *QAbstractButtonList) DetachShared_0() {
	// QAbstractButtonList_detachShared_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_detachShared_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isDetached()
func (this *QAbstractButtonList) IsDetached_0() bool {
	// QAbstractButtonList_isDetached_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_isDetached_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void setSharable(bool)
func (this *QAbstractButtonList) SetSharable_0() {
	// QAbstractButtonList_setSharable_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_setSharable_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isSharedWith(const QList<T> &)
func (this *QAbstractButtonList) IsSharedWith_0() bool {
	// QAbstractButtonList_isSharedWith_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_isSharedWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool isEmpty()
func (this *QAbstractButtonList) IsEmpty_0() bool {
	// QAbstractButtonList_isEmpty_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_isEmpty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void clear()
func (this *QAbstractButtonList) Clear_0() {
	// QAbstractButtonList_clear_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_clear_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// const T & at(int)
func (this *QAbstractButtonList) At_0() *QAbstractButton {
	// QAbstractButtonList_at_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_at_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAbstractButton{}
}

// const T & operator[](int)
func (this *QAbstractButtonList) Operator_get_index_0() *QAbstractButton {
	// QAbstractButtonList_operator_get_index_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_operator_get_index_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAbstractButton{}
}

// T & operator[](int)
func (this *QAbstractButtonList) Operator_get_index_1() *QAbstractButton {
	// QAbstractButtonList_operator_get_index_1()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_operator_get_index_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAbstractButton{}
}

// void reserve(int)
func (this *QAbstractButtonList) Reserve_0() {
	// QAbstractButtonList_reserve_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_reserve_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const T &)
func (this *QAbstractButtonList) Append_0() {
	// QAbstractButtonList_append_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_append_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const QList<T> &)
func (this *QAbstractButtonList) Append_1() {
	// QAbstractButtonList_append_1()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_append_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void prepend(const T &)
func (this *QAbstractButtonList) Prepend_0() {
	// QAbstractButtonList_prepend_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_prepend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void insert(int, const T &)
func (this *QAbstractButtonList) Insert_0() {
	// QAbstractButtonList_insert_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_insert_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void replace(int, const T &)
func (this *QAbstractButtonList) Replace_0() {
	// QAbstractButtonList_replace_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_replace_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeAt(int)
func (this *QAbstractButtonList) RemoveAt_0() {
	// QAbstractButtonList_removeAt_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_removeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int removeAll(const T &)
func (this *QAbstractButtonList) RemoveAll_0() int {
	// QAbstractButtonList_removeAll_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_removeAll_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool removeOne(const T &)
func (this *QAbstractButtonList) RemoveOne_0() bool {
	// QAbstractButtonList_removeOne_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_removeOne_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// T takeAt(int)
func (this *QAbstractButtonList) TakeAt_0() *QAbstractButton {
	// QAbstractButtonList_takeAt_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_takeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAbstractButton{}
}

// T takeFirst()
func (this *QAbstractButtonList) TakeFirst_0() *QAbstractButton {
	// QAbstractButtonList_takeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_takeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAbstractButton{}
}

// T takeLast()
func (this *QAbstractButtonList) TakeLast_0() *QAbstractButton {
	// QAbstractButtonList_takeLast_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_takeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAbstractButton{}
}

// void move(int, int)
func (this *QAbstractButtonList) Move_0() {
	// QAbstractButtonList_move_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_move_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void swap(int, int)
func (this *QAbstractButtonList) Swap_1() {
	// QAbstractButtonList_swap_1()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_swap_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int indexOf(const T &, int)
func (this *QAbstractButtonList) IndexOf_0() int {
	// QAbstractButtonList_indexOf_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_indexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int lastIndexOf(const T &, int)
func (this *QAbstractButtonList) LastIndexOf_0() int {
	// QAbstractButtonList_lastIndexOf_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_lastIndexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool contains(const T &)
func (this *QAbstractButtonList) Contains_0() bool {
	// QAbstractButtonList_contains_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_contains_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count(const T &)
func (this *QAbstractButtonList) Count_0() int {
	// QAbstractButtonList_count_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_count_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// QList::iterator begin()
func (this *QAbstractButtonList) Begin_0() {
	// QAbstractButtonList_begin_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_begin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator begin()
func (this *QAbstractButtonList) Begin_1() {
	// QAbstractButtonList_begin_1()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_begin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cbegin()
func (this *QAbstractButtonList) Cbegin_0() {
	// QAbstractButtonList_cbegin_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_cbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constBegin()
func (this *QAbstractButtonList) ConstBegin_0() {
	// QAbstractButtonList_constBegin_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_constBegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator end()
func (this *QAbstractButtonList) End_0() {
	// QAbstractButtonList_end_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_end_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator end()
func (this *QAbstractButtonList) End_1() {
	// QAbstractButtonList_end_1()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_end_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cend()
func (this *QAbstractButtonList) Cend_0() {
	// QAbstractButtonList_cend_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_cend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constEnd()
func (this *QAbstractButtonList) ConstEnd_0() {
	// QAbstractButtonList_constEnd_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_constEnd_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rbegin()
func (this *QAbstractButtonList) Rbegin_0() {
	// QAbstractButtonList_rbegin_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_rbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rend()
func (this *QAbstractButtonList) Rend_0() {
	// QAbstractButtonList_rend_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_rend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rbegin()
func (this *QAbstractButtonList) Rbegin_1() {
	// QAbstractButtonList_rbegin_1()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_rbegin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rend()
func (this *QAbstractButtonList) Rend_1() {
	// QAbstractButtonList_rend_1()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_rend_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crbegin()
func (this *QAbstractButtonList) Crbegin_0() {
	// QAbstractButtonList_crbegin_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_crbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crend()
func (this *QAbstractButtonList) Crend_0() {
	// QAbstractButtonList_crend_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_crend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator insert(QList::iterator, const T &)
func (this *QAbstractButtonList) Insert_1() {
	// QAbstractButtonList_insert_1()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_insert_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator)
func (this *QAbstractButtonList) Erase_0() {
	// QAbstractButtonList_erase_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_erase_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator, QList::iterator)
func (this *QAbstractButtonList) Erase_1() {
	// QAbstractButtonList_erase_1()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_erase_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int count()
func (this *QAbstractButtonList) Count_1() int {
	// QAbstractButtonList_count_1()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_count_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int length()
func (this *QAbstractButtonList) Length_0() int {
	// QAbstractButtonList_length_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_length_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// T & first()
func (this *QAbstractButtonList) First_0() *QAbstractButton {
	// QAbstractButtonList_first_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_first_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAbstractButton{}
}

// const T & constFirst()
func (this *QAbstractButtonList) ConstFirst_0() *QAbstractButton {
	// QAbstractButtonList_constFirst_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_constFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAbstractButton{}
}

// const T & first()
func (this *QAbstractButtonList) First_1() *QAbstractButton {
	// QAbstractButtonList_first_1()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_first_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAbstractButton{}
}

// T & last()
func (this *QAbstractButtonList) Last_0() *QAbstractButton {
	// QAbstractButtonList_last_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_last_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAbstractButton{}
}

// const T & last()
func (this *QAbstractButtonList) Last_1() *QAbstractButton {
	// QAbstractButtonList_last_1()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_last_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAbstractButton{}
}

// const T & constLast()
func (this *QAbstractButtonList) ConstLast_0() *QAbstractButton {
	// QAbstractButtonList_constLast_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_constLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAbstractButton{}
}

// void removeFirst()
func (this *QAbstractButtonList) RemoveFirst_0() {
	// QAbstractButtonList_removeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_removeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeLast()
func (this *QAbstractButtonList) RemoveLast_0() {
	// QAbstractButtonList_removeLast_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_removeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool startsWith(const T &)
func (this *QAbstractButtonList) StartsWith_0() bool {
	// QAbstractButtonList_startsWith_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_startsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool endsWith(const T &)
func (this *QAbstractButtonList) EndsWith_0() bool {
	// QAbstractButtonList_endsWith_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_endsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> mid(int, int)
func (this *QAbstractButtonList) Mid_0() *QAbstractButtonList {
	// QAbstractButtonList_mid_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_mid_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// T value(int)
func (this *QAbstractButtonList) Value_0() *QAbstractButton {
	// QAbstractButtonList_value_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_value_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAbstractButton{}
}

// T value(int, const T &)
func (this *QAbstractButtonList) Value_1() *QAbstractButton {
	// QAbstractButtonList_value_1()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_value_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAbstractButton{}
}

// void push_back(const T &)
func (this *QAbstractButtonList) Push_back_0() {
	// QAbstractButtonList_push_back_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_push_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void push_front(const T &)
func (this *QAbstractButtonList) Push_front_0() {
	// QAbstractButtonList_push_front_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_push_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// T & front()
func (this *QAbstractButtonList) Front_0() *QAbstractButton {
	// QAbstractButtonList_front_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAbstractButton{}
}

// const T & front()
func (this *QAbstractButtonList) Front_1() *QAbstractButton {
	// QAbstractButtonList_front_1()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_front_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAbstractButton{}
}

// T & back()
func (this *QAbstractButtonList) Back_0() *QAbstractButton {
	// QAbstractButtonList_back_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAbstractButton{}
}

// const T & back()
func (this *QAbstractButtonList) Back_1() *QAbstractButton {
	// QAbstractButtonList_back_1()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_back_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAbstractButton{}
}

// void pop_front()
func (this *QAbstractButtonList) Pop_front_0() {
	// QAbstractButtonList_pop_front_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_pop_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void pop_back()
func (this *QAbstractButtonList) Pop_back_0() {
	// QAbstractButtonList_pop_back_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_pop_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool empty()
func (this *QAbstractButtonList) Empty_0() bool {
	// QAbstractButtonList_empty_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_empty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> & operator+=(const QList<T> &)
func (this *QAbstractButtonList) Operator_add_equal_0() *QAbstractButtonList {
	// QAbstractButtonList_operator_add_equal_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_operator_add_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> operator+(const QList<T> &)
func (this *QAbstractButtonList) Operator_add_0() *QAbstractButtonList {
	// QAbstractButtonList_operator_add_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_operator_add_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator+=(const T &)
func (this *QAbstractButtonList) Operator_add_equal_1() *QAbstractButtonList {
	// QAbstractButtonList_operator_add_equal_1()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_operator_add_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const T &)
func (this *QAbstractButtonList) Operator_left_shift_0() *QAbstractButtonList {
	// QAbstractButtonList_operator_left_shift_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_operator_left_shift_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const QList<T> &)
func (this *QAbstractButtonList) Operator_left_shift_1() *QAbstractButtonList {
	// QAbstractButtonList_operator_left_shift_1()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_operator_left_shift_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QVector<T> toVector()
func (this *QAbstractButtonList) ToVector_0() {
	// QAbstractButtonList_toVector_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_toVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QSet<T> toSet()
func (this *QAbstractButtonList) ToSet_0() {
	// QAbstractButtonList_toSet_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_toSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList<T> fromVector(const QVector<T> &)
func (this *QAbstractButtonList) FromVector_0() *QAbstractButtonList {
	// QAbstractButtonList_fromVector_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_fromVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromSet(const QSet<T> &)
func (this *QAbstractButtonList) FromSet_0() *QAbstractButtonList {
	// QAbstractButtonList_fromSet_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_fromSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromStdList(const std::list<T> &)
func (this *QAbstractButtonList) FromStdList_0() *QAbstractButtonList {
	// QAbstractButtonList_fromStdList_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_fromStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// std::list<T> toStdList()
func (this *QAbstractButtonList) ToStdList_0() {
	// QAbstractButtonList_toStdList_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_toStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::Node * detach_helper_grow(int, int)
func (this *QAbstractButtonList) Detach_helper_grow_0() {
	// QAbstractButtonList_detach_helper_grow_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_detach_helper_grow_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper(int)
func (this *QAbstractButtonList) Detach_helper_0() {
	// QAbstractButtonList_detach_helper_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_detach_helper_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper()
func (this *QAbstractButtonList) Detach_helper_1() {
	// QAbstractButtonList_detach_helper_1()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_detach_helper_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void dealloc(QListData::Data *)
func (this *QAbstractButtonList) Dealloc_0() {
	// QAbstractButtonList_dealloc_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_dealloc_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_construct(QList::Node *, const T &)
func (this *QAbstractButtonList) Node_construct_0() {
	// QAbstractButtonList_node_construct_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_node_construct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *)
func (this *QAbstractButtonList) Node_destruct_0() {
	// QAbstractButtonList_node_destruct_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_node_destruct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_copy(QList::Node *, QList::Node *, QList::Node *)
func (this *QAbstractButtonList) Node_copy_0() {
	// QAbstractButtonList_node_copy_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_node_copy_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *, QList::Node *)
func (this *QAbstractButtonList) Node_destruct_1() {
	// QAbstractButtonList_node_destruct_1()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_node_destruct_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isValidIterator(const QList::iterator &)
func (this *QAbstractButtonList) IsValidIterator_0() bool {
	// QAbstractButtonList_isValidIterator_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_isValidIterator_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::NotArrayCompatibleLayout)
func (this *QAbstractButtonList) Op_eq_impl_0() bool {
	// QAbstractButtonList_op_eq_impl_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_op_eq_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::ArrayCompatibleLayout)
func (this *QAbstractButtonList) Op_eq_impl_1() bool {
	// QAbstractButtonList_op_eq_impl_1()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_op_eq_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QAbstractButtonList) Contains_impl_0() bool {
	// QAbstractButtonList_contains_impl_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_contains_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QAbstractButtonList) Contains_impl_1() bool {
	// QAbstractButtonList_contains_impl_1()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_contains_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QAbstractButtonList) Count_impl_0() int {
	// QAbstractButtonList_count_impl_0()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_count_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int count_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QAbstractButtonList) Count_impl_1() int {
	// QAbstractButtonList_count_impl_1()
	rv, err := qtrt.InvokeQtFunc6("QAbstractButtonList_count_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

//  body block end
