package pool

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"math/big"

	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/msgpencode"
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Pool) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 1 {
		err = msgp.ArrayError{Wanted: 1, Got: zb0001}
		return
	}
	err = z.Info.DecodeMsg(dc)
	if err != nil {
		err = msgp.WrapError(err, "Info")
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Pool) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 1
	err = en.Append(0x91)
	if err != nil {
		return
	}
	err = z.Info.EncodeMsg(en)
	if err != nil {
		err = msgp.WrapError(err, "Info")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Pool) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 1
	o = append(o, 0x91)
	o, err = z.Info.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "Info")
		return
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Pool) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 1 {
		err = msgp.ArrayError{Wanted: 1, Got: zb0001}
		return
	}
	bts, err = z.Info.UnmarshalMsg(bts)
	if err != nil {
		err = msgp.WrapError(err, "Info")
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Pool) Msgsize() (s int) {
	s = 1 + z.Info.Msgsize()
	return
}

// DecodeMsg implements msgp.Decodable
func (z *PoolInfo) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 9 {
		err = msgp.ArrayError{Wanted: 9, Got: zb0001}
		return
	}
	z.Address, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "Address")
		return
	}
	z.ReserveUsd, err = dc.ReadFloat64()
	if err != nil {
		err = msgp.WrapError(err, "ReserveUsd")
		return
	}
	if dc.IsNil() {
		err = dc.ReadNil()
		if err != nil {
			err = msgp.WrapError(err, "SwapFee")
			return
		}
		z.SwapFee = nil
	} else {
		{
			var zb0002 []byte
			zb0002, err = dc.ReadBytes(msgpencode.EncodeInt(z.SwapFee))
			if err != nil {
				err = msgp.WrapError(err, "SwapFee")
				return
			}
			z.SwapFee = msgpencode.DecodeInt(zb0002)
		}
	}
	z.Exchange, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "Exchange")
		return
	}
	z.Type, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "Type")
		return
	}
	var zb0003 uint32
	zb0003, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err, "Tokens")
		return
	}
	if cap(z.Tokens) >= int(zb0003) {
		z.Tokens = (z.Tokens)[:zb0003]
	} else {
		z.Tokens = make([]string, zb0003)
	}
	for za0001 := range z.Tokens {
		z.Tokens[za0001], err = dc.ReadString()
		if err != nil {
			err = msgp.WrapError(err, "Tokens", za0001)
			return
		}
	}
	var zb0004 uint32
	zb0004, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err, "Reserves")
		return
	}
	if cap(z.Reserves) >= int(zb0004) {
		z.Reserves = (z.Reserves)[:zb0004]
	} else {
		z.Reserves = make([]*big.Int, zb0004)
	}
	for za0002 := range z.Reserves {
		if dc.IsNil() {
			err = dc.ReadNil()
			if err != nil {
				err = msgp.WrapError(err, "Reserves", za0002)
				return
			}
			z.Reserves[za0002] = nil
		} else {
			{
				var zb0005 []byte
				zb0005, err = dc.ReadBytes(msgpencode.EncodeInt(z.Reserves[za0002]))
				if err != nil {
					err = msgp.WrapError(err, "Reserves", za0002)
					return
				}
				z.Reserves[za0002] = msgpencode.DecodeInt(zb0005)
			}
		}
	}
	z.Checked, err = dc.ReadBool()
	if err != nil {
		err = msgp.WrapError(err, "Checked")
		return
	}
	z.BlockNumber, err = dc.ReadUint64()
	if err != nil {
		err = msgp.WrapError(err, "BlockNumber")
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *PoolInfo) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 9
	err = en.Append(0x99)
	if err != nil {
		return
	}
	err = en.WriteString(z.Address)
	if err != nil {
		err = msgp.WrapError(err, "Address")
		return
	}
	err = en.WriteFloat64(z.ReserveUsd)
	if err != nil {
		err = msgp.WrapError(err, "ReserveUsd")
		return
	}
	if z.SwapFee == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteBytes(msgpencode.EncodeInt(z.SwapFee))
		if err != nil {
			err = msgp.WrapError(err, "SwapFee")
			return
		}
	}
	err = en.WriteString(z.Exchange)
	if err != nil {
		err = msgp.WrapError(err, "Exchange")
		return
	}
	err = en.WriteString(z.Type)
	if err != nil {
		err = msgp.WrapError(err, "Type")
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Tokens)))
	if err != nil {
		err = msgp.WrapError(err, "Tokens")
		return
	}
	for za0001 := range z.Tokens {
		err = en.WriteString(z.Tokens[za0001])
		if err != nil {
			err = msgp.WrapError(err, "Tokens", za0001)
			return
		}
	}
	err = en.WriteArrayHeader(uint32(len(z.Reserves)))
	if err != nil {
		err = msgp.WrapError(err, "Reserves")
		return
	}
	for za0002 := range z.Reserves {
		if z.Reserves[za0002] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = en.WriteBytes(msgpencode.EncodeInt(z.Reserves[za0002]))
			if err != nil {
				err = msgp.WrapError(err, "Reserves", za0002)
				return
			}
		}
	}
	err = en.WriteBool(z.Checked)
	if err != nil {
		err = msgp.WrapError(err, "Checked")
		return
	}
	err = en.WriteUint64(z.BlockNumber)
	if err != nil {
		err = msgp.WrapError(err, "BlockNumber")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *PoolInfo) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 9
	o = append(o, 0x99)
	o = msgp.AppendString(o, z.Address)
	o = msgp.AppendFloat64(o, z.ReserveUsd)
	if z.SwapFee == nil {
		o = msgp.AppendNil(o)
	} else {
		o = msgp.AppendBytes(o, msgpencode.EncodeInt(z.SwapFee))
	}
	o = msgp.AppendString(o, z.Exchange)
	o = msgp.AppendString(o, z.Type)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Tokens)))
	for za0001 := range z.Tokens {
		o = msgp.AppendString(o, z.Tokens[za0001])
	}
	o = msgp.AppendArrayHeader(o, uint32(len(z.Reserves)))
	for za0002 := range z.Reserves {
		if z.Reserves[za0002] == nil {
			o = msgp.AppendNil(o)
		} else {
			o = msgp.AppendBytes(o, msgpencode.EncodeInt(z.Reserves[za0002]))
		}
	}
	o = msgp.AppendBool(o, z.Checked)
	o = msgp.AppendUint64(o, z.BlockNumber)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *PoolInfo) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 9 {
		err = msgp.ArrayError{Wanted: 9, Got: zb0001}
		return
	}
	z.Address, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Address")
		return
	}
	z.ReserveUsd, bts, err = msgp.ReadFloat64Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "ReserveUsd")
		return
	}
	if msgp.IsNil(bts) {
		bts, err = msgp.ReadNilBytes(bts)
		if err != nil {
			return
		}
		z.SwapFee = nil
	} else {
		{
			var zb0002 []byte
			zb0002, bts, err = msgp.ReadBytesBytes(bts, msgpencode.EncodeInt(z.SwapFee))
			if err != nil {
				err = msgp.WrapError(err, "SwapFee")
				return
			}
			z.SwapFee = msgpencode.DecodeInt(zb0002)
		}
	}
	z.Exchange, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Exchange")
		return
	}
	z.Type, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Type")
		return
	}
	var zb0003 uint32
	zb0003, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Tokens")
		return
	}
	if cap(z.Tokens) >= int(zb0003) {
		z.Tokens = (z.Tokens)[:zb0003]
	} else {
		z.Tokens = make([]string, zb0003)
	}
	for za0001 := range z.Tokens {
		z.Tokens[za0001], bts, err = msgp.ReadStringBytes(bts)
		if err != nil {
			err = msgp.WrapError(err, "Tokens", za0001)
			return
		}
	}
	var zb0004 uint32
	zb0004, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Reserves")
		return
	}
	if cap(z.Reserves) >= int(zb0004) {
		z.Reserves = (z.Reserves)[:zb0004]
	} else {
		z.Reserves = make([]*big.Int, zb0004)
	}
	for za0002 := range z.Reserves {
		if msgp.IsNil(bts) {
			bts, err = msgp.ReadNilBytes(bts)
			if err != nil {
				return
			}
			z.Reserves[za0002] = nil
		} else {
			{
				var zb0005 []byte
				zb0005, bts, err = msgp.ReadBytesBytes(bts, msgpencode.EncodeInt(z.Reserves[za0002]))
				if err != nil {
					err = msgp.WrapError(err, "Reserves", za0002)
					return
				}
				z.Reserves[za0002] = msgpencode.DecodeInt(zb0005)
			}
		}
	}
	z.Checked, bts, err = msgp.ReadBoolBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Checked")
		return
	}
	z.BlockNumber, bts, err = msgp.ReadUint64Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "BlockNumber")
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *PoolInfo) Msgsize() (s int) {
	s = 1 + msgp.StringPrefixSize + len(z.Address) + msgp.Float64Size
	if z.SwapFee == nil {
		s += msgp.NilSize
	} else {
		s += msgp.BytesPrefixSize + len(msgpencode.EncodeInt(z.SwapFee))
	}
	s += msgp.StringPrefixSize + len(z.Exchange) + msgp.StringPrefixSize + len(z.Type) + msgp.ArrayHeaderSize
	for za0001 := range z.Tokens {
		s += msgp.StringPrefixSize + len(z.Tokens[za0001])
	}
	s += msgp.ArrayHeaderSize
	for za0002 := range z.Reserves {
		if z.Reserves[za0002] == nil {
			s += msgp.NilSize
		} else {
			s += msgp.BytesPrefixSize + len(msgpencode.EncodeInt(z.Reserves[za0002]))
		}
	}
	s += msgp.BoolSize + msgp.Uint64Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *PoolToken) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 5 {
		err = msgp.ArrayError{Wanted: 5, Got: zb0001}
		return
	}
	z.Token, err = dc.ReadString()
	if err != nil {
		err = msgp.WrapError(err, "Token")
		return
	}
	if dc.IsNil() {
		err = dc.ReadNil()
		if err != nil {
			err = msgp.WrapError(err, "Balance")
			return
		}
		z.Balance = nil
	} else {
		{
			var zb0002 []byte
			zb0002, err = dc.ReadBytes(msgpencode.EncodeInt(z.Balance))
			if err != nil {
				err = msgp.WrapError(err, "Balance")
				return
			}
			z.Balance = msgpencode.DecodeInt(zb0002)
		}
	}
	z.Weight, err = dc.ReadUint()
	if err != nil {
		err = msgp.WrapError(err, "Weight")
		return
	}
	if dc.IsNil() {
		err = dc.ReadNil()
		if err != nil {
			err = msgp.WrapError(err, "PrecisionMultiplier")
			return
		}
		z.PrecisionMultiplier = nil
	} else {
		{
			var zb0003 []byte
			zb0003, err = dc.ReadBytes(msgpencode.EncodeInt(z.PrecisionMultiplier))
			if err != nil {
				err = msgp.WrapError(err, "PrecisionMultiplier")
				return
			}
			z.PrecisionMultiplier = msgpencode.DecodeInt(zb0003)
		}
	}
	if dc.IsNil() {
		err = dc.ReadNil()
		if err != nil {
			err = msgp.WrapError(err, "VReserve")
			return
		}
		z.VReserve = nil
	} else {
		{
			var zb0004 []byte
			zb0004, err = dc.ReadBytes(msgpencode.EncodeInt(z.VReserve))
			if err != nil {
				err = msgp.WrapError(err, "VReserve")
				return
			}
			z.VReserve = msgpencode.DecodeInt(zb0004)
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *PoolToken) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 5
	err = en.Append(0x95)
	if err != nil {
		return
	}
	err = en.WriteString(z.Token)
	if err != nil {
		err = msgp.WrapError(err, "Token")
		return
	}
	if z.Balance == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteBytes(msgpencode.EncodeInt(z.Balance))
		if err != nil {
			err = msgp.WrapError(err, "Balance")
			return
		}
	}
	err = en.WriteUint(z.Weight)
	if err != nil {
		err = msgp.WrapError(err, "Weight")
		return
	}
	if z.PrecisionMultiplier == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteBytes(msgpencode.EncodeInt(z.PrecisionMultiplier))
		if err != nil {
			err = msgp.WrapError(err, "PrecisionMultiplier")
			return
		}
	}
	if z.VReserve == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = en.WriteBytes(msgpencode.EncodeInt(z.VReserve))
		if err != nil {
			err = msgp.WrapError(err, "VReserve")
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *PoolToken) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 5
	o = append(o, 0x95)
	o = msgp.AppendString(o, z.Token)
	if z.Balance == nil {
		o = msgp.AppendNil(o)
	} else {
		o = msgp.AppendBytes(o, msgpencode.EncodeInt(z.Balance))
	}
	o = msgp.AppendUint(o, z.Weight)
	if z.PrecisionMultiplier == nil {
		o = msgp.AppendNil(o)
	} else {
		o = msgp.AppendBytes(o, msgpencode.EncodeInt(z.PrecisionMultiplier))
	}
	if z.VReserve == nil {
		o = msgp.AppendNil(o)
	} else {
		o = msgp.AppendBytes(o, msgpencode.EncodeInt(z.VReserve))
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *PoolToken) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 5 {
		err = msgp.ArrayError{Wanted: 5, Got: zb0001}
		return
	}
	z.Token, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Token")
		return
	}
	if msgp.IsNil(bts) {
		bts, err = msgp.ReadNilBytes(bts)
		if err != nil {
			return
		}
		z.Balance = nil
	} else {
		{
			var zb0002 []byte
			zb0002, bts, err = msgp.ReadBytesBytes(bts, msgpencode.EncodeInt(z.Balance))
			if err != nil {
				err = msgp.WrapError(err, "Balance")
				return
			}
			z.Balance = msgpencode.DecodeInt(zb0002)
		}
	}
	z.Weight, bts, err = msgp.ReadUintBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Weight")
		return
	}
	if msgp.IsNil(bts) {
		bts, err = msgp.ReadNilBytes(bts)
		if err != nil {
			return
		}
		z.PrecisionMultiplier = nil
	} else {
		{
			var zb0003 []byte
			zb0003, bts, err = msgp.ReadBytesBytes(bts, msgpencode.EncodeInt(z.PrecisionMultiplier))
			if err != nil {
				err = msgp.WrapError(err, "PrecisionMultiplier")
				return
			}
			z.PrecisionMultiplier = msgpencode.DecodeInt(zb0003)
		}
	}
	if msgp.IsNil(bts) {
		bts, err = msgp.ReadNilBytes(bts)
		if err != nil {
			return
		}
		z.VReserve = nil
	} else {
		{
			var zb0004 []byte
			zb0004, bts, err = msgp.ReadBytesBytes(bts, msgpencode.EncodeInt(z.VReserve))
			if err != nil {
				err = msgp.WrapError(err, "VReserve")
				return
			}
			z.VReserve = msgpencode.DecodeInt(zb0004)
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *PoolToken) Msgsize() (s int) {
	s = 1 + msgp.StringPrefixSize + len(z.Token)
	if z.Balance == nil {
		s += msgp.NilSize
	} else {
		s += msgp.BytesPrefixSize + len(msgpencode.EncodeInt(z.Balance))
	}
	s += msgp.UintSize
	if z.PrecisionMultiplier == nil {
		s += msgp.NilSize
	} else {
		s += msgp.BytesPrefixSize + len(msgpencode.EncodeInt(z.PrecisionMultiplier))
	}
	if z.VReserve == nil {
		s += msgp.NilSize
	} else {
		s += msgp.BytesPrefixSize + len(msgpencode.EncodeInt(z.VReserve))
	}
	return
}
