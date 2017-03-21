package main

// NOTE: THIS FILE WAS PRODUCED BY THE
// ZEBRAPACK CODE GENERATION TOOL (github.com/glycerine/zebrapack)
// DO NOT EDIT

import "github.com/glycerine/zebrapack/msgp"

// DecodeMsg implements msgp.Decodable
// We treat empty fields as if we read a Nil from the wire.
func (z *A) DecodeMsg(dc *msgp.Reader) (err error) {
	var sawTopNil bool
	if dc.IsNil() {
		sawTopNil = true
		err = dc.ReadNil()
		if err != nil {
			return
		}
		dc.PushAlwaysNil()
	}

	var field []byte
	_ = field
	const maxFields0zdww = 6

	// -- templateDecodeMsgZid starts here--
	var totalEncodedFields0zdww uint32
	totalEncodedFields0zdww, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zdww := totalEncodedFields0zdww
	missingFieldsLeft0zdww := maxFields0zdww - totalEncodedFields0zdww

	var nextMiss0zdww int = -1
	var found0zdww [maxFields0zdww]bool
	var curField0zdww int

doneWithStruct0zdww:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zdww > 0 || missingFieldsLeft0zdww > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zdww, missingFieldsLeft0zdww, msgp.ShowFound(found0zdww[:]), decodeMsgFieldOrder0zdww)
		if encodedFieldsLeft0zdww > 0 {
			encodedFieldsLeft0zdww--
			curField0zdww, err = dc.ReadInt()
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss0zdww < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zdww = 0
			}
			for nextMiss0zdww < maxFields0zdww && (found0zdww[nextMiss0zdww] || decodeMsgFieldSkip0zdww[nextMiss0zdww]) {
				nextMiss0zdww++
			}
			if nextMiss0zdww == maxFields0zdww {
				// filled all the empty fields!
				break doneWithStruct0zdww
			}
			missingFieldsLeft0zdww--
			curField0zdww = nextMiss0zdww
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zdww)
		switch curField0zdww {
		// -- templateDecodeMsgZid ends here --

		case 0:
			// zid 0 for "name"
			found0zdww[0] = true
			z.Name, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case 1:
			// zid 1 for "Bday"
			found0zdww[1] = true
			z.Bday, err = dc.ReadTime()
			if err != nil {
				panic(err)
			}
		case 2:
			// zid 2 for "phone"
			found0zdww[2] = true
			z.Phone, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case 3:
			// zid 3 for "Sibs"
			found0zdww[3] = true
			z.Sibs, err = dc.ReadInt()
			if err != nil {
				panic(err)
			}
		case 4:
			// zid 4 for "GPA"
			found0zdww[4] = true
			z.GPA, err = dc.ReadFloat64()
			if err != nil {
				panic(err)
			}
		case 5:
			// zid 5 for "Friend"
			found0zdww[5] = true
			z.Friend, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		default:
			err = dc.Skip()
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss0zdww != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of A
var decodeMsgFieldOrder0zdww = []string{"name", "Bday", "phone", "Sibs", "GPA", "Friend"}

var decodeMsgFieldSkip0zdww = []bool{false, false, false, false, false, false}

// fieldsNotEmpty supports omitempty tags
func (z *A) fieldsNotEmpty(isempty []bool) uint32 {
	if len(isempty) == 0 {
		return 6
	}
	var fieldsInUse uint32 = 6
	isempty[0] = (len(z.Name) == 0) // string, omitempty
	if isempty[0] {
		fieldsInUse--
	}
	isempty[1] = (z.Bday.IsZero()) // time.Time, omitempty
	if isempty[1] {
		fieldsInUse--
	}
	isempty[2] = (len(z.Phone) == 0) // string, omitempty
	if isempty[2] {
		fieldsInUse--
	}
	isempty[3] = (z.Sibs == 0) // number, omitempty
	if isempty[3] {
		fieldsInUse--
	}
	isempty[4] = (z.GPA == 0) // number, omitempty
	if isempty[4] {
		fieldsInUse--
	}
	isempty[5] = (!z.Friend) // bool, omitempty
	if isempty[5] {
		fieldsInUse--
	}

	return fieldsInUse
}

// EncodeMsg implements msgp.Encodable
func (z *A) EncodeMsg(en *msgp.Writer) (err error) {

	// honor the omitempty tags
	var empty_zjbp [6]bool
	fieldsInUse_zycq := z.fieldsNotEmpty(empty_zjbp[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zycq)
	if err != nil {
		return err
	}

	if !empty_zjbp[0] {
		// zid 0 for "name"
		err = en.Append(0x0)
		if err != nil {
			return err
		}
		err = en.WriteString(z.Name)
		if err != nil {
			panic(err)
		}
	}

	if !empty_zjbp[1] {
		// zid 1 for "Bday"
		err = en.Append(0x1)
		if err != nil {
			return err
		}
		err = en.WriteTime(z.Bday)
		if err != nil {
			panic(err)
		}
	}

	if !empty_zjbp[2] {
		// zid 2 for "phone"
		err = en.Append(0x2)
		if err != nil {
			return err
		}
		err = en.WriteString(z.Phone)
		if err != nil {
			panic(err)
		}
	}

	if !empty_zjbp[3] {
		// zid 3 for "Sibs"
		err = en.Append(0x3)
		if err != nil {
			return err
		}
		err = en.WriteInt(z.Sibs)
		if err != nil {
			panic(err)
		}
	}

	if !empty_zjbp[4] {
		// zid 4 for "GPA"
		err = en.Append(0x4)
		if err != nil {
			return err
		}
		err = en.WriteFloat64(z.GPA)
		if err != nil {
			panic(err)
		}
	}

	if !empty_zjbp[5] {
		// zid 5 for "Friend"
		err = en.Append(0x5)
		if err != nil {
			return err
		}
		err = en.WriteBool(z.Friend)
		if err != nil {
			panic(err)
		}
	}

	return
}

// MarshalMsg implements msgp.Marshaler
func (z *A) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())

	// honor the omitempty tags
	var empty [6]bool
	fieldsInUse := z.fieldsNotEmpty(empty[:])
	o = msgp.AppendMapHeader(o, fieldsInUse)

	if !empty[0] {
		// zid 0 for "name"
		o = append(o, 0x0)
		o = msgp.AppendString(o, z.Name)
	}

	if !empty[1] {
		// zid 1 for "Bday"
		o = append(o, 0x1)
		o = msgp.AppendTime(o, z.Bday)
	}

	if !empty[2] {
		// zid 2 for "phone"
		o = append(o, 0x2)
		o = msgp.AppendString(o, z.Phone)
	}

	if !empty[3] {
		// zid 3 for "Sibs"
		o = append(o, 0x3)
		o = msgp.AppendInt(o, z.Sibs)
	}

	if !empty[4] {
		// zid 4 for "GPA"
		o = append(o, 0x4)
		o = msgp.AppendFloat64(o, z.GPA)
	}

	if !empty[5] {
		// zid 5 for "Friend"
		o = append(o, 0x5)
		o = msgp.AppendBool(o, z.Friend)
	}

	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *A) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.UnmarshalMsgWithCfg(bts, nil)
}
func (z *A) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields1zgin = 6

	// -- templateUnmarshalMsgZid starts here--
	var totalEncodedFields1zgin uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zgin, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft1zgin := totalEncodedFields1zgin
	missingFieldsLeft1zgin := maxFields1zgin - totalEncodedFields1zgin

	var nextMiss1zgin int = -1
	var found1zgin [maxFields1zgin]bool
	var curField1zgin int

doneWithStruct1zgin:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zgin > 0 || missingFieldsLeft1zgin > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zgin, missingFieldsLeft1zgin, msgp.ShowFound(found1zgin[:]), unmarshalMsgFieldOrder1zgin)
		if encodedFieldsLeft1zgin > 0 {
			encodedFieldsLeft1zgin--
			curField1zgin, bts, err = nbs.ReadIntBytes(bts)
			if err != nil {
				panic(err)
				return
			}
		} else {
			//missing fields need handling
			if nextMiss1zgin < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zgin = 0
			}
			for nextMiss1zgin < maxFields1zgin && (found1zgin[nextMiss1zgin] || unmarshalMsgFieldSkip1zgin[nextMiss1zgin]) {
				nextMiss1zgin++
			}
			if nextMiss1zgin == maxFields1zgin {
				// filled all the empty fields!
				break doneWithStruct1zgin
			}
			missingFieldsLeft1zgin--
			curField1zgin = nextMiss1zgin
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zgin)
		switch curField1zgin {
		// -- templateUnmarshalMsgZid ends here --

		case 0:
			// zid 0 for "name"
			found1zgin[0] = true
			z.Name, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case 1:
			// zid 1 for "Bday"
			found1zgin[1] = true
			z.Bday, bts, err = nbs.ReadTimeBytes(bts)

			if err != nil {
				panic(err)
			}
		case 2:
			// zid 2 for "phone"
			found1zgin[2] = true
			z.Phone, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case 3:
			// zid 3 for "Sibs"
			found1zgin[3] = true
			z.Sibs, bts, err = nbs.ReadIntBytes(bts)

			if err != nil {
				panic(err)
			}
		case 4:
			// zid 4 for "GPA"
			found1zgin[4] = true
			z.GPA, bts, err = nbs.ReadFloat64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case 5:
			// zid 5 for "Friend"
			found1zgin[5] = true
			z.Friend, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss1zgin != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of A
var unmarshalMsgFieldOrder1zgin = []string{"name", "Bday", "phone", "Sibs", "GPA", "Friend"}

var unmarshalMsgFieldSkip1zgin = []bool{false, false, false, false, false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *A) Msgsize() (s int) {
	s = 1 + 1 + msgp.StringPrefixSize + len(z.Name) + 1 + msgp.TimeSize + 1 + msgp.StringPrefixSize + len(z.Phone) + 1 + msgp.IntSize + 1 + msgp.Float64Size + 1 + msgp.BoolSize
	return
}

// DecodeMsg implements msgp.Decodable
// We treat empty fields as if we read a Nil from the wire.
func (z *Big) DecodeMsg(dc *msgp.Reader) (err error) {
	var sawTopNil bool
	if dc.IsNil() {
		sawTopNil = true
		err = dc.ReadNil()
		if err != nil {
			return
		}
		dc.PushAlwaysNil()
	}

	var field []byte
	_ = field
	const maxFields2zsxo = 5

	// -- templateDecodeMsgZid starts here--
	var totalEncodedFields2zsxo uint32
	totalEncodedFields2zsxo, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2zsxo := totalEncodedFields2zsxo
	missingFieldsLeft2zsxo := maxFields2zsxo - totalEncodedFields2zsxo

	var nextMiss2zsxo int = -1
	var found2zsxo [maxFields2zsxo]bool
	var curField2zsxo int

doneWithStruct2zsxo:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zsxo > 0 || missingFieldsLeft2zsxo > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zsxo, missingFieldsLeft2zsxo, msgp.ShowFound(found2zsxo[:]), decodeMsgFieldOrder2zsxo)
		if encodedFieldsLeft2zsxo > 0 {
			encodedFieldsLeft2zsxo--
			curField2zsxo, err = dc.ReadInt()
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss2zsxo < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2zsxo = 0
			}
			for nextMiss2zsxo < maxFields2zsxo && (found2zsxo[nextMiss2zsxo] || decodeMsgFieldSkip2zsxo[nextMiss2zsxo]) {
				nextMiss2zsxo++
			}
			if nextMiss2zsxo == maxFields2zsxo {
				// filled all the empty fields!
				break doneWithStruct2zsxo
			}
			missingFieldsLeft2zsxo--
			curField2zsxo = nextMiss2zsxo
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zsxo)
		switch curField2zsxo {
		// -- templateDecodeMsgZid ends here --

		case 0:
			// zid 0 for "Slice"
			found2zsxo[0] = true
			var zlse uint32
			zlse, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Slice) >= int(zlse) {
				z.Slice = (z.Slice)[:zlse]
			} else {
				z.Slice = make([]S2, zlse)
			}
			for zuwh := range z.Slice {
				err = z.Slice[zuwh].DecodeMsg(dc)
				if err != nil {
					panic(err)
				}
			}
		case 1:
			// zid 1 for "Transform"
			found2zsxo[1] = true
			var zbzx uint32
			zbzx, err = dc.ReadMapHeader()
			if err != nil {
				panic(err)
			}
			if z.Transform == nil && zbzx > 0 {
				z.Transform = make(map[int]*S2, zbzx)
			} else if len(z.Transform) > 0 {
				for key, _ := range z.Transform {
					delete(z.Transform, key)
				}
			}
			for zbzx > 0 {
				zbzx--
				var zzrg int
				var zahz *S2
				zzrg, err = dc.ReadInt()
				if err != nil {
					panic(err)
				}
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}

					if zahz != nil {
						dc.PushAlwaysNil()
						err = zahz.DecodeMsg(dc)
						if err != nil {
							return
						}
						dc.PopAlwaysNil()
					}
				} else {
					// not Nil, we have something to read

					if zahz == nil {
						zahz = new(S2)
					}
					err = zahz.DecodeMsg(dc)
					if err != nil {
						panic(err)
					}
				}
				z.Transform[zzrg] = zahz
			}
		case 2:
			// zid 2 for "Myptr"
			found2zsxo[2] = true
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}

				if z.Myptr != nil {
					dc.PushAlwaysNil()
					err = z.Myptr.DecodeMsg(dc)
					if err != nil {
						return
					}
					dc.PopAlwaysNil()
				}
			} else {
				// not Nil, we have something to read

				if z.Myptr == nil {
					z.Myptr = new(S2)
				}
				err = z.Myptr.DecodeMsg(dc)
				if err != nil {
					panic(err)
				}
			}
		case 3:
			// zid 3 for "Myarray"
			found2zsxo[3] = true
			if dc.AlwaysNil {
				// nothing more here
			} else if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
			} else {

				var zwpi uint32
				zwpi, err = dc.ReadArrayHeader()
				if err != nil {
					panic(err)
				}
				if !dc.IsNil() && zwpi != 3 {
					err = msgp.ArrayError{Wanted: 3, Got: zwpi}
					return
				}
			}
			for zxmq := range z.Myarray {
				z.Myarray[zxmq], err = dc.ReadString()
				if err != nil {
					panic(err)
				}
			}
		case 4:
			// zid 4 for "MySlice"
			found2zsxo[4] = true
			var zkha uint32
			zkha, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.MySlice) >= int(zkha) {
				z.MySlice = (z.MySlice)[:zkha]
			} else {
				z.MySlice = make([]string, zkha)
			}
			for zmro := range z.MySlice {
				z.MySlice[zmro], err = dc.ReadString()
				if err != nil {
					panic(err)
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss2zsxo != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Big
var decodeMsgFieldOrder2zsxo = []string{"Slice", "Transform", "Myptr", "Myarray", "MySlice"}

var decodeMsgFieldSkip2zsxo = []bool{false, false, false, false, false}

// fieldsNotEmpty supports omitempty tags
func (z *Big) fieldsNotEmpty(isempty []bool) uint32 {
	if len(isempty) == 0 {
		return 5
	}
	var fieldsInUse uint32 = 5
	isempty[0] = (len(z.Slice) == 0) // string, omitempty
	if isempty[0] {
		fieldsInUse--
	}
	isempty[1] = (len(z.Transform) == 0) // string, omitempty
	if isempty[1] {
		fieldsInUse--
	}
	isempty[2] = (z.Myptr == nil) // pointer, omitempty
	if isempty[2] {
		fieldsInUse--
	}
	isempty[3] = (len(z.Myarray) == 0) // string, omitempty
	if isempty[3] {
		fieldsInUse--
	}
	isempty[4] = (len(z.MySlice) == 0) // string, omitempty
	if isempty[4] {
		fieldsInUse--
	}

	return fieldsInUse
}

// EncodeMsg implements msgp.Encodable
func (z *Big) EncodeMsg(en *msgp.Writer) (err error) {

	// honor the omitempty tags
	var empty_zzmj [5]bool
	fieldsInUse_zfgd := z.fieldsNotEmpty(empty_zzmj[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zfgd)
	if err != nil {
		return err
	}

	if !empty_zzmj[0] {
		// zid 0 for "Slice"
		err = en.Append(0x0)
		if err != nil {
			return err
		}
		err = en.WriteArrayHeader(uint32(len(z.Slice)))
		if err != nil {
			panic(err)
		}
		for zuwh := range z.Slice {
			err = z.Slice[zuwh].EncodeMsg(en)
			if err != nil {
				panic(err)
			}
		}
	}

	if !empty_zzmj[1] {
		// zid 1 for "Transform"
		err = en.Append(0x1)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.Transform)))
		if err != nil {
			panic(err)
		}
		for zzrg, zahz := range z.Transform {
			err = en.WriteInt(zzrg)
			if err != nil {
				panic(err)
			}
			if zahz == nil {
				err = en.WriteNil()
				if err != nil {
					return
				}
			} else {
				err = zahz.EncodeMsg(en)
				if err != nil {
					panic(err)
				}
			}
		}
	}

	if !empty_zzmj[2] {
		// zid 2 for "Myptr"
		err = en.Append(0x2)
		if err != nil {
			return err
		}
		if z.Myptr == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.Myptr.EncodeMsg(en)
			if err != nil {
				panic(err)
			}
		}
	}

	if !empty_zzmj[3] {
		// zid 3 for "Myarray"
		err = en.Append(0x3)
		if err != nil {
			return err
		}
		err = en.WriteArrayHeader(3)
		if err != nil {
			panic(err)
		}
		for zxmq := range z.Myarray {
			err = en.WriteString(z.Myarray[zxmq])
			if err != nil {
				panic(err)
			}
		}
	}

	if !empty_zzmj[4] {
		// zid 4 for "MySlice"
		err = en.Append(0x4)
		if err != nil {
			return err
		}
		err = en.WriteArrayHeader(uint32(len(z.MySlice)))
		if err != nil {
			panic(err)
		}
		for zmro := range z.MySlice {
			err = en.WriteString(z.MySlice[zmro])
			if err != nil {
				panic(err)
			}
		}
	}

	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Big) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())

	// honor the omitempty tags
	var empty [5]bool
	fieldsInUse := z.fieldsNotEmpty(empty[:])
	o = msgp.AppendMapHeader(o, fieldsInUse)

	if !empty[0] {
		// zid 0 for "Slice"
		o = append(o, 0x0)
		o = msgp.AppendArrayHeader(o, uint32(len(z.Slice)))
		for zuwh := range z.Slice {
			o, err = z.Slice[zuwh].MarshalMsg(o)
			if err != nil {
				panic(err)
			}
		}
	}

	if !empty[1] {
		// zid 1 for "Transform"
		o = append(o, 0x1)
		o = msgp.AppendMapHeader(o, uint32(len(z.Transform)))
		for zzrg, zahz := range z.Transform {
			o = msgp.AppendInt(o, zzrg)
			if zahz == nil {
				o = msgp.AppendNil(o)
			} else {
				o, err = zahz.MarshalMsg(o)
				if err != nil {
					panic(err)
				}
			}
		}
	}

	if !empty[2] {
		// zid 2 for "Myptr"
		o = append(o, 0x2)
		if z.Myptr == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Myptr.MarshalMsg(o)
			if err != nil {
				panic(err)
			}
		}
	}

	if !empty[3] {
		// zid 3 for "Myarray"
		o = append(o, 0x3)
		o = msgp.AppendArrayHeader(o, 3)
		for zxmq := range z.Myarray {
			o = msgp.AppendString(o, z.Myarray[zxmq])
		}
	}

	if !empty[4] {
		// zid 4 for "MySlice"
		o = append(o, 0x4)
		o = msgp.AppendArrayHeader(o, uint32(len(z.MySlice)))
		for zmro := range z.MySlice {
			o = msgp.AppendString(o, z.MySlice[zmro])
		}
	}

	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Big) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.UnmarshalMsgWithCfg(bts, nil)
}
func (z *Big) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields3zvji = 5

	// -- templateUnmarshalMsgZid starts here--
	var totalEncodedFields3zvji uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3zvji, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft3zvji := totalEncodedFields3zvji
	missingFieldsLeft3zvji := maxFields3zvji - totalEncodedFields3zvji

	var nextMiss3zvji int = -1
	var found3zvji [maxFields3zvji]bool
	var curField3zvji int

doneWithStruct3zvji:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3zvji > 0 || missingFieldsLeft3zvji > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zvji, missingFieldsLeft3zvji, msgp.ShowFound(found3zvji[:]), unmarshalMsgFieldOrder3zvji)
		if encodedFieldsLeft3zvji > 0 {
			encodedFieldsLeft3zvji--
			curField3zvji, bts, err = nbs.ReadIntBytes(bts)
			if err != nil {
				panic(err)
				return
			}
		} else {
			//missing fields need handling
			if nextMiss3zvji < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3zvji = 0
			}
			for nextMiss3zvji < maxFields3zvji && (found3zvji[nextMiss3zvji] || unmarshalMsgFieldSkip3zvji[nextMiss3zvji]) {
				nextMiss3zvji++
			}
			if nextMiss3zvji == maxFields3zvji {
				// filled all the empty fields!
				break doneWithStruct3zvji
			}
			missingFieldsLeft3zvji--
			curField3zvji = nextMiss3zvji
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3zvji)
		switch curField3zvji {
		// -- templateUnmarshalMsgZid ends here --

		case 0:
			// zid 0 for "Slice"
			found3zvji[0] = true
			if nbs.AlwaysNil {
				(z.Slice) = (z.Slice)[:0]
			} else {

				var zgye uint32
				zgye, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Slice) >= int(zgye) {
					z.Slice = (z.Slice)[:zgye]
				} else {
					z.Slice = make([]S2, zgye)
				}
				for zuwh := range z.Slice {
					bts, err = z.Slice[zuwh].UnmarshalMsg(bts)
					if err != nil {
						panic(err)
					}
					if err != nil {
						panic(err)
					}
				}
			}
		case 1:
			// zid 1 for "Transform"
			found3zvji[1] = true
			if nbs.AlwaysNil {
				if len(z.Transform) > 0 {
					for key, _ := range z.Transform {
						delete(z.Transform, key)
					}
				}

			} else {

				var zauk uint32
				zauk, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.Transform == nil && zauk > 0 {
					z.Transform = make(map[int]*S2, zauk)
				} else if len(z.Transform) > 0 {
					for key, _ := range z.Transform {
						delete(z.Transform, key)
					}
				}
				for zauk > 0 {
					var zzrg int
					var zahz *S2
					zauk--
					zzrg, bts, err = nbs.ReadIntBytes(bts)
					if err != nil {
						panic(err)
					}
					if nbs.AlwaysNil {
						if zahz != nil {
							zahz.UnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != zahz {
								zahz.UnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if zahz == nil {
								zahz = new(S2)
							}
							bts, err = zahz.UnmarshalMsg(bts)
							if err != nil {
								panic(err)
							}
							if err != nil {
								panic(err)
							}
						}
					}
					z.Transform[zzrg] = zahz
				}
			}
		case 2:
			// zid 2 for "Myptr"
			found3zvji[2] = true
			if nbs.AlwaysNil {
				if z.Myptr != nil {
					z.Myptr.UnmarshalMsg(msgp.OnlyNilSlice)
				}
			} else {
				// not nbs.AlwaysNil
				if msgp.IsNil(bts) {
					bts = bts[1:]
					if nil != z.Myptr {
						z.Myptr.UnmarshalMsg(msgp.OnlyNilSlice)
					}
				} else {
					// not nbs.AlwaysNil and not IsNil(bts): have something to read

					if z.Myptr == nil {
						z.Myptr = new(S2)
					}
					bts, err = z.Myptr.UnmarshalMsg(bts)
					if err != nil {
						panic(err)
					}
					if err != nil {
						panic(err)
					}
				}
			}
		case 3:
			// zid 3 for "Myarray"
			found3zvji[3] = true
			var zbas uint32
			zbas, bts, err = nbs.ReadArrayHeaderBytes(bts)
			if err != nil {
				panic(err)
			}
			if !nbs.IsNil(bts) && zbas != 3 {
				err = msgp.ArrayError{Wanted: 3, Got: zbas}
				return
			}
			for zxmq := range z.Myarray {
				z.Myarray[zxmq], bts, err = nbs.ReadStringBytes(bts)

				if err != nil {
					panic(err)
				}
			}
		case 4:
			// zid 4 for "MySlice"
			found3zvji[4] = true
			if nbs.AlwaysNil {
				(z.MySlice) = (z.MySlice)[:0]
			} else {

				var zqgg uint32
				zqgg, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.MySlice) >= int(zqgg) {
					z.MySlice = (z.MySlice)[:zqgg]
				} else {
					z.MySlice = make([]string, zqgg)
				}
				for zmro := range z.MySlice {
					z.MySlice[zmro], bts, err = nbs.ReadStringBytes(bts)

					if err != nil {
						panic(err)
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss3zvji != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Big
var unmarshalMsgFieldOrder3zvji = []string{"Slice", "Transform", "Myptr", "Myarray", "MySlice"}

var unmarshalMsgFieldSkip3zvji = []bool{false, false, false, false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Big) Msgsize() (s int) {
	s = 1 + 1 + msgp.ArrayHeaderSize
	for zuwh := range z.Slice {
		s += z.Slice[zuwh].Msgsize()
	}
	s += 1 + msgp.MapHeaderSize
	if z.Transform != nil {
		for zzrg, zahz := range z.Transform {
			_ = zahz
			_ = zzrg
			s += msgp.IntSize
			if zahz == nil {
				s += msgp.NilSize
			} else {
				s += zahz.Msgsize()
			}
		}
	}
	s += 1
	if z.Myptr == nil {
		s += msgp.NilSize
	} else {
		s += z.Myptr.Msgsize()
	}
	s += 1 + msgp.ArrayHeaderSize
	for zxmq := range z.Myarray {
		s += msgp.StringPrefixSize + len(z.Myarray[zxmq])
	}
	s += 1 + msgp.ArrayHeaderSize
	for zmro := range z.MySlice {
		s += msgp.StringPrefixSize + len(z.MySlice[zmro])
	}
	return
}

// DecodeMsg implements msgp.Decodable
// We treat empty fields as if we read a Nil from the wire.
func (z *S2) DecodeMsg(dc *msgp.Reader) (err error) {
	var sawTopNil bool
	if dc.IsNil() {
		sawTopNil = true
		err = dc.ReadNil()
		if err != nil {
			return
		}
		dc.PushAlwaysNil()
	}

	var field []byte
	_ = field
	const maxFields4zpjt = 6

	// -- templateDecodeMsgZid starts here--
	var totalEncodedFields4zpjt uint32
	totalEncodedFields4zpjt, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft4zpjt := totalEncodedFields4zpjt
	missingFieldsLeft4zpjt := maxFields4zpjt - totalEncodedFields4zpjt

	var nextMiss4zpjt int = -1
	var found4zpjt [maxFields4zpjt]bool
	var curField4zpjt int

doneWithStruct4zpjt:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zpjt > 0 || missingFieldsLeft4zpjt > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zpjt, missingFieldsLeft4zpjt, msgp.ShowFound(found4zpjt[:]), decodeMsgFieldOrder4zpjt)
		if encodedFieldsLeft4zpjt > 0 {
			encodedFieldsLeft4zpjt--
			curField4zpjt, err = dc.ReadInt()
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss4zpjt < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss4zpjt = 0
			}
			for nextMiss4zpjt < maxFields4zpjt && (found4zpjt[nextMiss4zpjt] || decodeMsgFieldSkip4zpjt[nextMiss4zpjt]) {
				nextMiss4zpjt++
			}
			if nextMiss4zpjt == maxFields4zpjt {
				// filled all the empty fields!
				break doneWithStruct4zpjt
			}
			missingFieldsLeft4zpjt--
			curField4zpjt = nextMiss4zpjt
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zpjt)
		switch curField4zpjt {
		// -- templateDecodeMsgZid ends here --

		case 1:
			// zid 1 for "beta"
			found4zpjt[1] = true
			z.B, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case 2:
			// zid 2 for "ralph"
			found4zpjt[2] = true
			var zphf uint32
			zphf, err = dc.ReadMapHeader()
			if err != nil {
				panic(err)
			}
			if z.R == nil && zphf > 0 {
				z.R = make(map[string]uint8, zphf)
			} else if len(z.R) > 0 {
				for key, _ := range z.R {
					delete(z.R, key)
				}
			}
			for zphf > 0 {
				zphf--
				var ztcd string
				var zchc uint8
				ztcd, err = dc.ReadString()
				if err != nil {
					panic(err)
				}
				zchc, err = dc.ReadUint8()
				if err != nil {
					panic(err)
				}
				z.R[ztcd] = zchc
			}
		case 3:
			// zid 3 for "P"
			found4zpjt[3] = true
			z.P, err = dc.ReadUint16()
			if err != nil {
				panic(err)
			}
		case 4:
			// zid 4 for "Q"
			found4zpjt[4] = true
			z.Q, err = dc.ReadUint32()
			if err != nil {
				panic(err)
			}
		case 5:
			// zid 5 for "T"
			found4zpjt[5] = true
			z.T, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		default:
			err = dc.Skip()
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss4zpjt != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of S2
var decodeMsgFieldOrder4zpjt = []string{"alpha", "beta", "ralph", "P", "Q", "T"}

var decodeMsgFieldSkip4zpjt = []bool{true, false, false, false, false, false}

// fieldsNotEmpty supports omitempty tags
func (z *S2) fieldsNotEmpty(isempty []bool) uint32 {
	if len(isempty) == 0 {
		return 5
	}
	var fieldsInUse uint32 = 5
	isempty[1] = (len(z.B) == 0) // string, omitempty
	if isempty[1] {
		fieldsInUse--
	}
	isempty[2] = (len(z.R) == 0) // string, omitempty
	if isempty[2] {
		fieldsInUse--
	}
	isempty[3] = (z.P == 0) // number, omitempty
	if isempty[3] {
		fieldsInUse--
	}
	isempty[4] = (z.Q == 0) // number, omitempty
	if isempty[4] {
		fieldsInUse--
	}
	isempty[5] = (z.T == 0) // number, omitempty
	if isempty[5] {
		fieldsInUse--
	}

	return fieldsInUse
}

// EncodeMsg implements msgp.Encodable
func (z *S2) EncodeMsg(en *msgp.Writer) (err error) {

	// honor the omitempty tags
	var empty_zjab [6]bool
	fieldsInUse_zzft := z.fieldsNotEmpty(empty_zjab[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zzft)
	if err != nil {
		return err
	}

	if !empty_zjab[1] {
		// zid 1 for "beta"
		err = en.Append(0x1)
		if err != nil {
			return err
		}
		err = en.WriteString(z.B)
		if err != nil {
			panic(err)
		}
	}

	if !empty_zjab[2] {
		// zid 2 for "ralph"
		err = en.Append(0x2)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.R)))
		if err != nil {
			panic(err)
		}
		for ztcd, zchc := range z.R {
			err = en.WriteString(ztcd)
			if err != nil {
				panic(err)
			}
			err = en.WriteUint8(zchc)
			if err != nil {
				panic(err)
			}
		}
	}

	if !empty_zjab[3] {
		// zid 3 for "P"
		err = en.Append(0x3)
		if err != nil {
			return err
		}
		err = en.WriteUint16(z.P)
		if err != nil {
			panic(err)
		}
	}

	if !empty_zjab[4] {
		// zid 4 for "Q"
		err = en.Append(0x4)
		if err != nil {
			return err
		}
		err = en.WriteUint32(z.Q)
		if err != nil {
			panic(err)
		}
	}

	if !empty_zjab[5] {
		// zid 5 for "T"
		err = en.Append(0x5)
		if err != nil {
			return err
		}
		err = en.WriteInt64(z.T)
		if err != nil {
			panic(err)
		}
	}

	return
}

// MarshalMsg implements msgp.Marshaler
func (z *S2) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())

	// honor the omitempty tags
	var empty [6]bool
	fieldsInUse := z.fieldsNotEmpty(empty[:])
	o = msgp.AppendMapHeader(o, fieldsInUse)

	if !empty[1] {
		// zid 1 for "beta"
		o = append(o, 0x1)
		o = msgp.AppendString(o, z.B)
	}

	if !empty[2] {
		// zid 2 for "ralph"
		o = append(o, 0x2)
		o = msgp.AppendMapHeader(o, uint32(len(z.R)))
		for ztcd, zchc := range z.R {
			o = msgp.AppendString(o, ztcd)
			o = msgp.AppendUint8(o, zchc)
		}
	}

	if !empty[3] {
		// zid 3 for "P"
		o = append(o, 0x3)
		o = msgp.AppendUint16(o, z.P)
	}

	if !empty[4] {
		// zid 4 for "Q"
		o = append(o, 0x4)
		o = msgp.AppendUint32(o, z.Q)
	}

	if !empty[5] {
		// zid 5 for "T"
		o = append(o, 0x5)
		o = msgp.AppendInt64(o, z.T)
	}

	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *S2) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.UnmarshalMsgWithCfg(bts, nil)
}
func (z *S2) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields5zswr = 6

	// -- templateUnmarshalMsgZid starts here--
	var totalEncodedFields5zswr uint32
	if !nbs.AlwaysNil {
		totalEncodedFields5zswr, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft5zswr := totalEncodedFields5zswr
	missingFieldsLeft5zswr := maxFields5zswr - totalEncodedFields5zswr

	var nextMiss5zswr int = -1
	var found5zswr [maxFields5zswr]bool
	var curField5zswr int

doneWithStruct5zswr:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft5zswr > 0 || missingFieldsLeft5zswr > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zswr, missingFieldsLeft5zswr, msgp.ShowFound(found5zswr[:]), unmarshalMsgFieldOrder5zswr)
		if encodedFieldsLeft5zswr > 0 {
			encodedFieldsLeft5zswr--
			curField5zswr, bts, err = nbs.ReadIntBytes(bts)
			if err != nil {
				panic(err)
				return
			}
		} else {
			//missing fields need handling
			if nextMiss5zswr < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss5zswr = 0
			}
			for nextMiss5zswr < maxFields5zswr && (found5zswr[nextMiss5zswr] || unmarshalMsgFieldSkip5zswr[nextMiss5zswr]) {
				nextMiss5zswr++
			}
			if nextMiss5zswr == maxFields5zswr {
				// filled all the empty fields!
				break doneWithStruct5zswr
			}
			missingFieldsLeft5zswr--
			curField5zswr = nextMiss5zswr
		}
		//fmt.Printf("switching on curField: '%v'\n", curField5zswr)
		switch curField5zswr {
		// -- templateUnmarshalMsgZid ends here --

		case 1:
			// zid 1 for "beta"
			found5zswr[1] = true
			z.B, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case 2:
			// zid 2 for "ralph"
			found5zswr[2] = true
			if nbs.AlwaysNil {
				if len(z.R) > 0 {
					for key, _ := range z.R {
						delete(z.R, key)
					}
				}

			} else {

				var zugo uint32
				zugo, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.R == nil && zugo > 0 {
					z.R = make(map[string]uint8, zugo)
				} else if len(z.R) > 0 {
					for key, _ := range z.R {
						delete(z.R, key)
					}
				}
				for zugo > 0 {
					var ztcd string
					var zchc uint8
					zugo--
					ztcd, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						panic(err)
					}
					zchc, bts, err = nbs.ReadUint8Bytes(bts)

					if err != nil {
						panic(err)
					}
					z.R[ztcd] = zchc
				}
			}
		case 3:
			// zid 3 for "P"
			found5zswr[3] = true
			z.P, bts, err = nbs.ReadUint16Bytes(bts)

			if err != nil {
				panic(err)
			}
		case 4:
			// zid 4 for "Q"
			found5zswr[4] = true
			z.Q, bts, err = nbs.ReadUint32Bytes(bts)

			if err != nil {
				panic(err)
			}
		case 5:
			// zid 5 for "T"
			found5zswr[5] = true
			z.T, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss5zswr != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of S2
var unmarshalMsgFieldOrder5zswr = []string{"alpha", "beta", "ralph", "P", "Q", "T"}

var unmarshalMsgFieldSkip5zswr = []bool{true, false, false, false, false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *S2) Msgsize() (s int) {
	s = 1 + 1 + msgp.StringPrefixSize + len(z.B) + 1 + msgp.MapHeaderSize
	if z.R != nil {
		for ztcd, zchc := range z.R {
			_ = zchc
			_ = ztcd
			s += msgp.StringPrefixSize + len(ztcd) + msgp.Uint8Size
		}
	}
	s += 1 + msgp.Uint16Size + 1 + msgp.Uint32Size + 1 + msgp.Int64Size
	return
}

// DecodeMsg implements msgp.Decodable
// We treat empty fields as if we read a Nil from the wire.
func (z *Sys) DecodeMsg(dc *msgp.Reader) (err error) {
	var sawTopNil bool
	if dc.IsNil() {
		sawTopNil = true
		err = dc.ReadNil()
		if err != nil {
			return
		}
		dc.PushAlwaysNil()
	}

	var field []byte
	_ = field
	const maxFields6zojp = 1

	// -- templateDecodeMsgZid starts here--
	var totalEncodedFields6zojp uint32
	totalEncodedFields6zojp, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft6zojp := totalEncodedFields6zojp
	missingFieldsLeft6zojp := maxFields6zojp - totalEncodedFields6zojp

	var nextMiss6zojp int = -1
	var found6zojp [maxFields6zojp]bool
	var curField6zojp int

doneWithStruct6zojp:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft6zojp > 0 || missingFieldsLeft6zojp > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft6zojp, missingFieldsLeft6zojp, msgp.ShowFound(found6zojp[:]), decodeMsgFieldOrder6zojp)
		if encodedFieldsLeft6zojp > 0 {
			encodedFieldsLeft6zojp--
			curField6zojp, err = dc.ReadInt()
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss6zojp < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss6zojp = 0
			}
			for nextMiss6zojp < maxFields6zojp && (found6zojp[nextMiss6zojp] || decodeMsgFieldSkip6zojp[nextMiss6zojp]) {
				nextMiss6zojp++
			}
			if nextMiss6zojp == maxFields6zojp {
				// filled all the empty fields!
				break doneWithStruct6zojp
			}
			missingFieldsLeft6zojp--
			curField6zojp = nextMiss6zojp
		}
		//fmt.Printf("switching on curField: '%v'\n", curField6zojp)
		switch curField6zojp {
		// -- templateDecodeMsgZid ends here --

		case 0:
			// zid 0 for "F"
			found6zojp[0] = true
			z.F, err = dc.ReadIntf()
			if err != nil {
				panic(err)
			}
		default:
			err = dc.Skip()
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss6zojp != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Sys
var decodeMsgFieldOrder6zojp = []string{"F"}

var decodeMsgFieldSkip6zojp = []bool{false}

// fieldsNotEmpty supports omitempty tags
func (z Sys) fieldsNotEmpty(isempty []bool) uint32 {
	if len(isempty) == 0 {
		return 1
	}
	var fieldsInUse uint32 = 1
	isempty[0] = false
	if isempty[0] {
		fieldsInUse--
	}

	return fieldsInUse
}

// EncodeMsg implements msgp.Encodable
func (z Sys) EncodeMsg(en *msgp.Writer) (err error) {

	// honor the omitempty tags
	var empty_zzro [1]bool
	fieldsInUse_zlhh := z.fieldsNotEmpty(empty_zzro[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zlhh)
	if err != nil {
		return err
	}

	if !empty_zzro[0] {
		// zid 0 for "F"
		err = en.Append(0x0)
		if err != nil {
			return err
		}
		err = en.WriteIntf(z.F)
		if err != nil {
			panic(err)
		}
	}

	return
}

// MarshalMsg implements msgp.Marshaler
func (z Sys) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())

	// honor the omitempty tags
	var empty [1]bool
	fieldsInUse := z.fieldsNotEmpty(empty[:])
	o = msgp.AppendMapHeader(o, fieldsInUse)

	if !empty[0] {
		// zid 0 for "F"
		o = append(o, 0x0)
		o, err = msgp.AppendIntf(o, z.F)
		if err != nil {
			panic(err)
		}
	}

	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Sys) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.UnmarshalMsgWithCfg(bts, nil)
}
func (z *Sys) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields7zqng = 1

	// -- templateUnmarshalMsgZid starts here--
	var totalEncodedFields7zqng uint32
	if !nbs.AlwaysNil {
		totalEncodedFields7zqng, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft7zqng := totalEncodedFields7zqng
	missingFieldsLeft7zqng := maxFields7zqng - totalEncodedFields7zqng

	var nextMiss7zqng int = -1
	var found7zqng [maxFields7zqng]bool
	var curField7zqng int

doneWithStruct7zqng:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft7zqng > 0 || missingFieldsLeft7zqng > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft7zqng, missingFieldsLeft7zqng, msgp.ShowFound(found7zqng[:]), unmarshalMsgFieldOrder7zqng)
		if encodedFieldsLeft7zqng > 0 {
			encodedFieldsLeft7zqng--
			curField7zqng, bts, err = nbs.ReadIntBytes(bts)
			if err != nil {
				panic(err)
				return
			}
		} else {
			//missing fields need handling
			if nextMiss7zqng < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss7zqng = 0
			}
			for nextMiss7zqng < maxFields7zqng && (found7zqng[nextMiss7zqng] || unmarshalMsgFieldSkip7zqng[nextMiss7zqng]) {
				nextMiss7zqng++
			}
			if nextMiss7zqng == maxFields7zqng {
				// filled all the empty fields!
				break doneWithStruct7zqng
			}
			missingFieldsLeft7zqng--
			curField7zqng = nextMiss7zqng
		}
		//fmt.Printf("switching on curField: '%v'\n", curField7zqng)
		switch curField7zqng {
		// -- templateUnmarshalMsgZid ends here --

		case 0:
			// zid 0 for "F"
			found7zqng[0] = true
			z.F, bts, err = nbs.ReadIntfBytes(bts)

			if err != nil {
				panic(err)
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss7zqng != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Sys
var unmarshalMsgFieldOrder7zqng = []string{"F"}

var unmarshalMsgFieldSkip7zqng = []bool{false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Sys) Msgsize() (s int) {
	s = 1 + 1 + msgp.GuessSize(z.F)
	return
}
