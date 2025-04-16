package uuid

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type UUID uuid.UUID

func (u *UUID) ScanUUID(v pgtype.UUID) error {
	if !v.Valid {
		return fmt.Errorf("cannot scan NULL into *uuid.UUID")
	}

	*u = v.Bytes
	return nil
}

func (u UUID) UUIDValue() (pgtype.UUID, error) {
	return pgtype.UUID{Bytes: [16]byte(u), Valid: true}, nil
}

type NullUUID uuid.NullUUID

func (u *NullUUID) ScanUUID(v pgtype.UUID) error {
	*u = NullUUID{UUID: uuid.UUID(v.Bytes), Valid: v.Valid}
	return nil
}

func (u NullUUID) UUIDValue() (pgtype.UUID, error) {
	return pgtype.UUID{Bytes: [16]byte(u.UUID), Valid: u.Valid}, nil
}

func TryWrapUUIDEncodePlan(value interface{}) (plan pgtype.WrappedEncodePlanNextSetter, nextValue interface{}, ok bool) {
	switch value := value.(type) {
	case uuid.UUID:
		return &wrapUUIDEncodePlan{}, UUID(value), true
	case uuid.NullUUID:
		return &wrapNullUUIDEncodePlan{}, NullUUID(value), true
	}

	return nil, nil, false
}

type wrapUUIDEncodePlan struct {
	next pgtype.EncodePlan
}

func (plan *wrapUUIDEncodePlan) SetNext(next pgtype.EncodePlan) { plan.next = next }

func (plan *wrapUUIDEncodePlan) Encode(value interface{}, buf []byte) (newBuf []byte, err error) {
	return plan.next.Encode(UUID(value.(uuid.UUID)), buf)
}

type wrapNullUUIDEncodePlan struct {
	next pgtype.EncodePlan
}

func (plan *wrapNullUUIDEncodePlan) SetNext(next pgtype.EncodePlan) { plan.next = next }

func (plan *wrapNullUUIDEncodePlan) Encode(value interface{}, buf []byte) (newBuf []byte, err error) {
	return plan.next.Encode(NullUUID(value.(uuid.NullUUID)), buf)
}

func TryWrapUUIDScanPlan(target interface{}) (plan pgtype.WrappedScanPlanNextSetter, nextDst interface{}, ok bool) {
	switch target := target.(type) {
	case *uuid.UUID:
		return &wrapUUIDScanPlan{}, (*UUID)(target), true
	case *uuid.NullUUID:
		return &wrapNullUUIDScanPlan{}, (*NullUUID)(target), true
	}

	return nil, nil, false
}

type wrapUUIDScanPlan struct {
	next pgtype.ScanPlan
}

func (plan *wrapUUIDScanPlan) SetNext(next pgtype.ScanPlan) { plan.next = next }

func (plan *wrapUUIDScanPlan) Scan(src []byte, dst interface{}) error {
	return plan.next.Scan(src, (*UUID)(dst.(*uuid.UUID)))
}

type wrapNullUUIDScanPlan struct {
	next pgtype.ScanPlan
}

func (plan *wrapNullUUIDScanPlan) SetNext(next pgtype.ScanPlan) { plan.next = next }

func (plan *wrapNullUUIDScanPlan) Scan(src []byte, dst interface{}) error {
	return plan.next.Scan(src, (*NullUUID)(dst.(*uuid.NullUUID)))
}

type UUIDCodec struct {
	pgtype.UUIDCodec
}

func (UUIDCodec) DecodeValue(tm *pgtype.Map, oid uint32, format int16, src []byte) (interface{}, error) {
	if src == nil {
		return nil, nil
	}

	var target uuid.UUID
	scanPlan := tm.PlanScan(oid, format, &target)
	if scanPlan == nil {
		return nil, fmt.Errorf("PlanScan did not find a plan")
	}

	err := scanPlan.Scan(src, &target)
	if err != nil {
		return nil, err
	}

	return target, nil
}

// Register registers the github.com/google/uuid integration with a pgtype.Map.
func Register(tm *pgtype.Map) {
	tm.TryWrapEncodePlanFuncs = append([]pgtype.TryWrapEncodePlanFunc{TryWrapUUIDEncodePlan}, tm.TryWrapEncodePlanFuncs...)
	tm.TryWrapScanPlanFuncs = append([]pgtype.TryWrapScanPlanFunc{TryWrapUUIDScanPlan}, tm.TryWrapScanPlanFuncs...)

	name := "uuid"
	arrayName := "_" + name

	tm.RegisterType(&pgtype.Type{
		Name:  "uuid",
		OID:   pgtype.UUIDOID,
		Codec: UUIDCodec{},
	})

	// https://github.com/jackc/pgx/blob/ff9c26d8516879a5d48ae9d729c994d58e64777d/pgtype/register_default_pg_types.go#L5
	var value uuid.UUID
	tm.RegisterDefaultPgType(value, name)  // T
	tm.RegisterDefaultPgType(&value, name) // *T

	var sliceT []uuid.UUID
	tm.RegisterDefaultPgType(sliceT, arrayName)  // []T
	tm.RegisterDefaultPgType(&sliceT, arrayName) // *[]T

	var slicePtrT []*uuid.UUID
	tm.RegisterDefaultPgType(slicePtrT, arrayName)  // []*T
	tm.RegisterDefaultPgType(&slicePtrT, arrayName) // *[]*T

	var arrayOfT pgtype.Array[uuid.UUID]
	tm.RegisterDefaultPgType(arrayOfT, arrayName)  // Array[T]
	tm.RegisterDefaultPgType(&arrayOfT, arrayName) // *Array[T]

	var arrayOfPtrT pgtype.Array[*uuid.UUID]
	tm.RegisterDefaultPgType(arrayOfPtrT, arrayName)  // Array[*T]
	tm.RegisterDefaultPgType(&arrayOfPtrT, arrayName) // *Array[*T]

	var flatArrayOfT pgtype.FlatArray[uuid.UUID]
	tm.RegisterDefaultPgType(flatArrayOfT, arrayName)  // FlatArray[T]
	tm.RegisterDefaultPgType(&flatArrayOfT, arrayName) // *FlatArray[T]

	var flatArrayOfPtrT pgtype.FlatArray[*uuid.UUID]
	tm.RegisterDefaultPgType(flatArrayOfPtrT, arrayName)  // FlatArray[*T]
	tm.RegisterDefaultPgType(&flatArrayOfPtrT, arrayName) // *FlatArray[*T]
}
