// Copyright 2014 The Go-NetCDF Authors. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// These files are autogenerated from nc_double.go using generate.sh

package netcdf

import (
	"unsafe"
)

// #include <stdlib.h>
// #include <netcdf.h>
import "C"

// WriteInt64 writes data as the entire data for variable v.
func (v Var) WriteInt64(data []int64) error {
	if err := v.okData(NC_INT64, len(data)); err != nil {
		return err
	}
	return newError(C.nc_put_var_longlong(C.int(v.f), C.int(v.id), (*C.longlong)(unsafe.Pointer(&data[0]))))
}

// ReadInt64 reads the entire variable v into data, which must have enough
// space for all the values (i.e. len(data) must be at least v.Len()).
func (v Var) ReadInt64(data []int64) error {
	if err := v.okData(NC_INT64, len(data)); err != nil {
		return err
	}
	return newError(C.nc_get_var_longlong(C.int(v.f), C.int(v.id), (*C.longlong)(unsafe.Pointer(&data[0]))))
}

// WriteInt64 sets the value of attribute a to val.
func (a Attr) WriteInt64(val []int64) error {
	// TODO: check Type is NC_DOUBLE and len(val) is corrent
	cname := C.CString(a.name)
	defer C.free(unsafe.Pointer(cname))
	return newError(C.nc_put_att_longlong(C.int(a.v.f), C.int(a.v.id), cname,
		C.nc_type(NC_INT64), C.size_t(len(val)), (*C.longlong)(unsafe.Pointer(&val[0]))))
}

// ReadInt64 returns the attribute value.
func (a Attr) ReadInt64() (val []int64, err error) {
	// TODO: check Type is NC_DOUBLE
	n, err := a.Len()
	if err != nil {
		return nil, err
	}
	cname := C.CString(a.name)
	defer C.free(unsafe.Pointer(cname))
	val = make([]int64, n)
	err = newError(C.nc_get_att_longlong(C.int(a.v.f), C.int(a.v.id), cname,
		(*C.longlong)(unsafe.Pointer(&val[0]))))
	return
}
