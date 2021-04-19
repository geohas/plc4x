//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

package model

import (
	"encoding/xml"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"io"
	"reflect"
	"strings"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type S7PayloadUserDataItem struct {
	ReturnCode    DataTransportErrorCode
	TransportSize DataTransportSize
	SzlId         *SzlId
	SzlIndex      uint16
	Child         IS7PayloadUserDataItemChild
}

// The corresponding interface
type IS7PayloadUserDataItem interface {
	CpuFunctionType() uint8
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

type IS7PayloadUserDataItemParent interface {
	SerializeParent(io utils.WriteBuffer, child IS7PayloadUserDataItem, serializeChildFunction func() error) error
	GetTypeName() string
}

type IS7PayloadUserDataItemChild interface {
	Serialize(io utils.WriteBuffer) error
	InitializeParent(parent *S7PayloadUserDataItem, returnCode DataTransportErrorCode, transportSize DataTransportSize, szlId *SzlId, szlIndex uint16)
	GetTypeName() string
	IS7PayloadUserDataItem
	utils.AsciiBoxer
}

func NewS7PayloadUserDataItem(returnCode DataTransportErrorCode, transportSize DataTransportSize, szlId *SzlId, szlIndex uint16) *S7PayloadUserDataItem {
	return &S7PayloadUserDataItem{ReturnCode: returnCode, TransportSize: transportSize, SzlId: szlId, SzlIndex: szlIndex}
}

func CastS7PayloadUserDataItem(structType interface{}) *S7PayloadUserDataItem {
	castFunc := func(typ interface{}) *S7PayloadUserDataItem {
		if casted, ok := typ.(S7PayloadUserDataItem); ok {
			return &casted
		}
		if casted, ok := typ.(*S7PayloadUserDataItem); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *S7PayloadUserDataItem) GetTypeName() string {
	return "S7PayloadUserDataItem"
}

func (m *S7PayloadUserDataItem) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *S7PayloadUserDataItem) LengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.LengthInBits()
}

func (m *S7PayloadUserDataItem) ParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Enum Field (returnCode)
	lengthInBits += 8

	// Enum Field (transportSize)
	lengthInBits += 8

	// Implicit Field (dataLength)
	lengthInBits += 16

	// Simple field (szlId)
	lengthInBits += m.SzlId.LengthInBits()

	// Simple field (szlIndex)
	lengthInBits += 16

	return lengthInBits
}

func (m *S7PayloadUserDataItem) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func S7PayloadUserDataItemParse(io utils.ReadBuffer, cpuFunctionType uint8) (*S7PayloadUserDataItem, error) {

	// Enum field (returnCode)
	returnCode, _returnCodeErr := DataTransportErrorCodeParse(io)
	if _returnCodeErr != nil {
		return nil, errors.Wrap(_returnCodeErr, "Error parsing 'returnCode' field")
	}

	// Enum field (transportSize)
	transportSize, _transportSizeErr := DataTransportSizeParse(io)
	if _transportSizeErr != nil {
		return nil, errors.Wrap(_transportSizeErr, "Error parsing 'transportSize' field")
	}

	// Implicit Field (dataLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	dataLength, _dataLengthErr := io.ReadUint16(16)
	_ = dataLength
	if _dataLengthErr != nil {
		return nil, errors.Wrap(_dataLengthErr, "Error parsing 'dataLength' field")
	}

	// Simple Field (szlId)
	szlId, _szlIdErr := SzlIdParse(io)
	if _szlIdErr != nil {
		return nil, errors.Wrap(_szlIdErr, "Error parsing 'szlId' field")
	}

	// Simple Field (szlIndex)
	szlIndex, _szlIndexErr := io.ReadUint16(16)
	if _szlIndexErr != nil {
		return nil, errors.Wrap(_szlIndexErr, "Error parsing 'szlIndex' field")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *S7PayloadUserDataItem
	var typeSwitchError error
	switch {
	case cpuFunctionType == 0x04: // S7PayloadUserDataItemCpuFunctionReadSzlRequest
		_parent, typeSwitchError = S7PayloadUserDataItemCpuFunctionReadSzlRequestParse(io)
	case cpuFunctionType == 0x08: // S7PayloadUserDataItemCpuFunctionReadSzlResponse
		_parent, typeSwitchError = S7PayloadUserDataItemCpuFunctionReadSzlResponseParse(io)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	// Finish initializing
	_parent.Child.InitializeParent(_parent, returnCode, transportSize, szlId, szlIndex)
	return _parent, nil
}

func (m *S7PayloadUserDataItem) Serialize(io utils.WriteBuffer) error {
	return m.Child.Serialize(io)
}

func (m *S7PayloadUserDataItem) SerializeParent(io utils.WriteBuffer, child IS7PayloadUserDataItem, serializeChildFunction func() error) error {
	io.PushContext("S7PayloadUserDataItem")

	// Enum field (returnCode)
	returnCode := CastDataTransportErrorCode(m.ReturnCode)
	_returnCodeErr := returnCode.Serialize(io)
	if _returnCodeErr != nil {
		return errors.Wrap(_returnCodeErr, "Error serializing 'returnCode' field")
	}

	// Enum field (transportSize)
	transportSize := CastDataTransportSize(m.TransportSize)
	_transportSizeErr := transportSize.Serialize(io)
	if _transportSizeErr != nil {
		return errors.Wrap(_transportSizeErr, "Error serializing 'transportSize' field")
	}

	// Implicit Field (dataLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	dataLength := uint16(uint16(uint16(m.LengthInBytes())) - uint16(uint16(4)))
	_dataLengthErr := io.WriteUint16("dataLength", 16, (dataLength))
	if _dataLengthErr != nil {
		return errors.Wrap(_dataLengthErr, "Error serializing 'dataLength' field")
	}

	// Simple Field (szlId)
	_szlIdErr := m.SzlId.Serialize(io)
	if _szlIdErr != nil {
		return errors.Wrap(_szlIdErr, "Error serializing 'szlId' field")
	}

	// Simple Field (szlIndex)
	szlIndex := uint16(m.SzlIndex)
	_szlIndexErr := io.WriteUint16("szlIndex", 16, (szlIndex))
	if _szlIndexErr != nil {
		return errors.Wrap(_szlIndexErr, "Error serializing 'szlIndex' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	_typeSwitchErr := serializeChildFunction()
	if _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	io.PopContext("S7PayloadUserDataItem")
	return nil
}

func (m *S7PayloadUserDataItem) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	foundContent := false
	if start.Attr != nil && len(start.Attr) > 0 {
		switch start.Attr[0].Value {
		// S7PayloadUserDataItemCpuFunctionReadSzlRequest needs special treatment as it has no fields
		case "org.apache.plc4x.java.s7.readwrite.S7PayloadUserDataItemCpuFunctionReadSzlRequest":
			if m.Child == nil {
				m.Child = &S7PayloadUserDataItemCpuFunctionReadSzlRequest{
					Parent: m,
				}
			}
		}
	}
	for {
		token, err = d.Token()
		if err != nil {
			if err == io.EOF && foundContent {
				return nil
			}
			return err
		}
		switch token.(type) {
		case xml.StartElement:
			foundContent = true
			tok := token.(xml.StartElement)
			switch tok.Name.Local {
			case "returnCode":
				var data DataTransportErrorCode
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.ReturnCode = data
			case "transportSize":
				var data DataTransportSize
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.TransportSize = data
			case "szlId":
				var data SzlId
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.SzlId = &data
			case "szlIndex":
				var data uint16
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.SzlIndex = data
			default:
				attr := start.Attr
				if attr == nil || len(attr) <= 0 {
					// TODO: workaround for bug with nested lists
					attr = tok.Attr
				}
				if attr == nil || len(attr) <= 0 {
					panic("Couldn't determine class type for childs of S7PayloadUserDataItem")
				}
				switch attr[0].Value {
				case "org.apache.plc4x.java.s7.readwrite.S7PayloadUserDataItemCpuFunctionReadSzlRequest":
					var dt *S7PayloadUserDataItemCpuFunctionReadSzlRequest
					if m.Child != nil {
						dt = m.Child.(*S7PayloadUserDataItemCpuFunctionReadSzlRequest)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.s7.readwrite.S7PayloadUserDataItemCpuFunctionReadSzlResponse":
					var dt *S7PayloadUserDataItemCpuFunctionReadSzlResponse
					if m.Child != nil {
						dt = m.Child.(*S7PayloadUserDataItemCpuFunctionReadSzlResponse)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				}
			}
		}
	}
}

func (m *S7PayloadUserDataItem) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	className := reflect.TypeOf(m.Child).String()
	className = "org.apache.plc4x.java.s7.readwrite." + className[strings.LastIndex(className, ".")+1:]
	if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
		{Name: xml.Name{Local: "className"}, Value: className},
	}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.ReturnCode, xml.StartElement{Name: xml.Name{Local: "returnCode"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.TransportSize, xml.StartElement{Name: xml.Name{Local: "transportSize"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.SzlId, xml.StartElement{Name: xml.Name{Local: "szlId"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.SzlIndex, xml.StartElement{Name: xml.Name{Local: "szlIndex"}}); err != nil {
		return err
	}
	marshaller, ok := m.Child.(xml.Marshaler)
	if !ok {
		return errors.Errorf("child is not castable to Marshaler. Actual type %T", m.Child)
	}
	if err := marshaller.MarshalXML(e, start); err != nil {
		return err
	}
	if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
		return err
	}
	return nil
}

func (m S7PayloadUserDataItem) String() string {
	return string(m.Box("", 120))
}

func (m *S7PayloadUserDataItem) Box(name string, width int) utils.AsciiBox {
	return m.Child.Box(name, width)
}

func (m *S7PayloadUserDataItem) BoxParent(name string, width int, childBoxer func() []utils.AsciiBox) utils.AsciiBox {
	boxName := "S7PayloadUserDataItem"
	if name != "" {
		boxName += "/" + name
	}
	boxes := make([]utils.AsciiBox, 0)
	// Enum field (returnCode)
	returnCode := CastDataTransportErrorCode(m.ReturnCode)
	boxes = append(boxes, returnCode.Box("returnCode", -1))
	// Enum field (transportSize)
	transportSize := CastDataTransportSize(m.TransportSize)
	boxes = append(boxes, transportSize.Box("transportSize", -1))
	// Implicit Field (dataLength)
	dataLength := uint16(uint16(uint16(m.LengthInBytes())) - uint16(uint16(4)))
	// uint16 can be boxed as anything with the least amount of space
	boxes = append(boxes, utils.BoxAnything("DataLength", dataLength, -1))
	// Complex field (case complex)
	boxes = append(boxes, m.SzlId.Box("szlId", width-2))
	// Simple field (case simple)
	// uint16 can be boxed as anything with the least amount of space
	boxes = append(boxes, utils.BoxAnything("SzlIndex", m.SzlIndex, -1))
	// Switch field (Depending on the discriminator values, passes the boxing to a sub-type)
	boxes = append(boxes, childBoxer()...)
	return utils.BoxBox(boxName, utils.AlignBoxes(boxes, width-2), 0)
}
