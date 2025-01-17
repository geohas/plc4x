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
	"fmt"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"io"
	"strconv"
)

// Code generated by build-utils. DO NOT EDIT.

// Constant values.
const BACnetTagWithContent_OPENTAG uint8 = 0x2e
const BACnetTagWithContent_CLOSINGTAG uint8 = 0x2f

// The data-structure of this message
type BACnetTagWithContent struct {
	TypeOrTagNumber    uint8
	ContextSpecificTag uint8
	LengthValueType    uint8
	ExtTagNumber       *uint8
	ExtLength          *uint8
	PropertyIdentifier []uint8
	Value              *BACnetTag
}

// The corresponding interface
type IBACnetTagWithContent interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

func NewBACnetTagWithContent(typeOrTagNumber uint8, contextSpecificTag uint8, lengthValueType uint8, extTagNumber *uint8, extLength *uint8, propertyIdentifier []uint8, value *BACnetTag) *BACnetTagWithContent {
	return &BACnetTagWithContent{TypeOrTagNumber: typeOrTagNumber, ContextSpecificTag: contextSpecificTag, LengthValueType: lengthValueType, ExtTagNumber: extTagNumber, ExtLength: extLength, PropertyIdentifier: propertyIdentifier, Value: value}
}

func CastBACnetTagWithContent(structType interface{}) *BACnetTagWithContent {
	castFunc := func(typ interface{}) *BACnetTagWithContent {
		if casted, ok := typ.(BACnetTagWithContent); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetTagWithContent); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetTagWithContent) GetTypeName() string {
	return "BACnetTagWithContent"
}

func (m *BACnetTagWithContent) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetTagWithContent) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (typeOrTagNumber)
	lengthInBits += 4

	// Simple field (contextSpecificTag)
	lengthInBits += 1

	// Simple field (lengthValueType)
	lengthInBits += 3

	// Optional Field (extTagNumber)
	if m.ExtTagNumber != nil {
		lengthInBits += 8
	}

	// Optional Field (extLength)
	if m.ExtLength != nil {
		lengthInBits += 8
	}

	// Array field
	if len(m.PropertyIdentifier) > 0 {
		lengthInBits += 8 * uint16(len(m.PropertyIdentifier))
	}

	// Const Field (openTag)
	lengthInBits += 8

	// Simple field (value)
	lengthInBits += m.Value.LengthInBits()

	// Const Field (closingTag)
	lengthInBits += 8

	return lengthInBits
}

func (m *BACnetTagWithContent) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetTagWithContentParse(io utils.ReadBuffer) (*BACnetTagWithContent, error) {

	// Simple Field (typeOrTagNumber)
	typeOrTagNumber, _typeOrTagNumberErr := io.ReadUint8(4)
	if _typeOrTagNumberErr != nil {
		return nil, errors.Wrap(_typeOrTagNumberErr, "Error parsing 'typeOrTagNumber' field")
	}

	// Simple Field (contextSpecificTag)
	contextSpecificTag, _contextSpecificTagErr := io.ReadUint8(1)
	if _contextSpecificTagErr != nil {
		return nil, errors.Wrap(_contextSpecificTagErr, "Error parsing 'contextSpecificTag' field")
	}

	// Simple Field (lengthValueType)
	lengthValueType, _lengthValueTypeErr := io.ReadUint8(3)
	if _lengthValueTypeErr != nil {
		return nil, errors.Wrap(_lengthValueTypeErr, "Error parsing 'lengthValueType' field")
	}

	// Optional Field (extTagNumber) (Can be skipped, if a given expression evaluates to false)
	var extTagNumber *uint8 = nil
	if bool((typeOrTagNumber) == (15)) {
		_val, _err := io.ReadUint8(8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'extTagNumber' field")
		}
		extTagNumber = &_val
	}

	// Optional Field (extLength) (Can be skipped, if a given expression evaluates to false)
	var extLength *uint8 = nil
	if bool((lengthValueType) == (5)) {
		_val, _err := io.ReadUint8(8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'extLength' field")
		}
		extLength = &_val
	}

	// Array field (propertyIdentifier)
	// Length array
	propertyIdentifier := make([]uint8, 0)
	_propertyIdentifierLength := utils.InlineIf(bool(bool((lengthValueType) == (5))), func() uint16 { return uint16((*extLength)) }, func() uint16 { return uint16(lengthValueType) })
	_propertyIdentifierEndPos := io.GetPos() + uint16(_propertyIdentifierLength)
	for io.GetPos() < _propertyIdentifierEndPos {
		_item, _err := io.ReadUint8(8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'propertyIdentifier' field")
		}
		propertyIdentifier = append(propertyIdentifier, _item)
	}

	// Const Field (openTag)
	openTag, _openTagErr := io.ReadUint8(8)
	if _openTagErr != nil {
		return nil, errors.Wrap(_openTagErr, "Error parsing 'openTag' field")
	}
	if openTag != BACnetTagWithContent_OPENTAG {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", BACnetTagWithContent_OPENTAG) + " but got " + fmt.Sprintf("%d", openTag))
	}

	// Simple Field (value)
	value, _valueErr := BACnetTagParse(io)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
	}

	// Const Field (closingTag)
	closingTag, _closingTagErr := io.ReadUint8(8)
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field")
	}
	if closingTag != BACnetTagWithContent_CLOSINGTAG {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", BACnetTagWithContent_CLOSINGTAG) + " but got " + fmt.Sprintf("%d", closingTag))
	}

	// Create the instance
	return NewBACnetTagWithContent(typeOrTagNumber, contextSpecificTag, lengthValueType, extTagNumber, extLength, propertyIdentifier, value), nil
}

func (m *BACnetTagWithContent) Serialize(io utils.WriteBuffer) error {
	io.PushContext("BACnetTagWithContent")

	// Simple Field (typeOrTagNumber)
	typeOrTagNumber := uint8(m.TypeOrTagNumber)
	_typeOrTagNumberErr := io.WriteUint8("typeOrTagNumber", 4, (typeOrTagNumber))
	if _typeOrTagNumberErr != nil {
		return errors.Wrap(_typeOrTagNumberErr, "Error serializing 'typeOrTagNumber' field")
	}

	// Simple Field (contextSpecificTag)
	contextSpecificTag := uint8(m.ContextSpecificTag)
	_contextSpecificTagErr := io.WriteUint8("contextSpecificTag", 1, (contextSpecificTag))
	if _contextSpecificTagErr != nil {
		return errors.Wrap(_contextSpecificTagErr, "Error serializing 'contextSpecificTag' field")
	}

	// Simple Field (lengthValueType)
	lengthValueType := uint8(m.LengthValueType)
	_lengthValueTypeErr := io.WriteUint8("lengthValueType", 3, (lengthValueType))
	if _lengthValueTypeErr != nil {
		return errors.Wrap(_lengthValueTypeErr, "Error serializing 'lengthValueType' field")
	}

	// Optional Field (extTagNumber) (Can be skipped, if the value is null)
	var extTagNumber *uint8 = nil
	if m.ExtTagNumber != nil {
		extTagNumber = m.ExtTagNumber
		_extTagNumberErr := io.WriteUint8("extTagNumber", 8, *(extTagNumber))
		if _extTagNumberErr != nil {
			return errors.Wrap(_extTagNumberErr, "Error serializing 'extTagNumber' field")
		}
	}

	// Optional Field (extLength) (Can be skipped, if the value is null)
	var extLength *uint8 = nil
	if m.ExtLength != nil {
		extLength = m.ExtLength
		_extLengthErr := io.WriteUint8("extLength", 8, *(extLength))
		if _extLengthErr != nil {
			return errors.Wrap(_extLengthErr, "Error serializing 'extLength' field")
		}
	}

	// Array Field (propertyIdentifier)
	if m.PropertyIdentifier != nil {
		for _, _element := range m.PropertyIdentifier {
			_elementErr := io.WriteUint8("", 8, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'propertyIdentifier' field")
			}
		}
	}

	// Const Field (openTag)
	_openTagErr := io.WriteUint8("openTag", 8, 0x2e)
	if _openTagErr != nil {
		return errors.Wrap(_openTagErr, "Error serializing 'openTag' field")
	}

	// Simple Field (value)
	_valueErr := m.Value.Serialize(io)
	if _valueErr != nil {
		return errors.Wrap(_valueErr, "Error serializing 'value' field")
	}

	// Const Field (closingTag)
	_closingTagErr := io.WriteUint8("closingTag", 8, 0x2f)
	if _closingTagErr != nil {
		return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
	}

	io.PopContext("BACnetTagWithContent")
	return nil
}

func (m *BACnetTagWithContent) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	foundContent := false
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
			case "typeOrTagNumber":
				var data uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.TypeOrTagNumber = data
			case "contextSpecificTag":
				var data uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.ContextSpecificTag = data
			case "lengthValueType":
				var data uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.LengthValueType = data
			case "extTagNumber":
				// When working with pointers we need to check for an empty element
				var dataString string
				if err := d.DecodeElement(&dataString, &tok); err != nil {
					return err
				}
				if dataString != "" {
					atoi, err := strconv.Atoi(dataString)
					if err != nil {
						return err
					}
					data := uint8(atoi)
					m.ExtTagNumber = &data
				}
			case "extLength":
				// When working with pointers we need to check for an empty element
				var dataString string
				if err := d.DecodeElement(&dataString, &tok); err != nil {
					return err
				}
				if dataString != "" {
					atoi, err := strconv.Atoi(dataString)
					if err != nil {
						return err
					}
					data := uint8(atoi)
					m.ExtLength = &data
				}
			case "propertyIdentifier":
				var data []uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.PropertyIdentifier = data
			case "value":
				var dt *BACnetTag
				if err := d.DecodeElement(&dt, &tok); err != nil {
					if err == io.EOF {
						continue
					}
					return err
				}
				m.Value = dt
			}
		}
	}
}

func (m *BACnetTagWithContent) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	className := "org.apache.plc4x.java.bacnetip.readwrite.BACnetTagWithContent"
	if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
		{Name: xml.Name{Local: "className"}, Value: className},
	}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.TypeOrTagNumber, xml.StartElement{Name: xml.Name{Local: "typeOrTagNumber"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.ContextSpecificTag, xml.StartElement{Name: xml.Name{Local: "contextSpecificTag"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.LengthValueType, xml.StartElement{Name: xml.Name{Local: "lengthValueType"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.ExtTagNumber, xml.StartElement{Name: xml.Name{Local: "extTagNumber"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.ExtLength, xml.StartElement{Name: xml.Name{Local: "extLength"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.PropertyIdentifier, xml.StartElement{Name: xml.Name{Local: "propertyIdentifier"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.Value, xml.StartElement{Name: xml.Name{Local: "value"}}); err != nil {
		return err
	}
	if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
		return err
	}
	return nil
}

func (m BACnetTagWithContent) String() string {
	return string(m.Box("", 120))
}

func (m BACnetTagWithContent) Box(name string, width int) utils.AsciiBox {
	boxName := "BACnetTagWithContent"
	if name != "" {
		boxName += "/" + name
	}
	boxes := make([]utils.AsciiBox, 0)
	// Simple field (case simple)
	// uint8 can be boxed as anything with the least amount of space
	boxes = append(boxes, utils.BoxAnything("TypeOrTagNumber", m.TypeOrTagNumber, -1))
	// Simple field (case simple)
	// uint8 can be boxed as anything with the least amount of space
	boxes = append(boxes, utils.BoxAnything("ContextSpecificTag", m.ContextSpecificTag, -1))
	// Simple field (case simple)
	// uint8 can be boxed as anything with the least amount of space
	boxes = append(boxes, utils.BoxAnything("LengthValueType", m.LengthValueType, -1))
	// Optional Field (extTagNumber) (Can be skipped, if the value is null)
	var extTagNumber *uint8 = nil
	if m.ExtTagNumber != nil {
		extTagNumber = m.ExtTagNumber
		// uint8 can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("ExtTagNumber", *(extTagNumber), -1))
	}
	// Optional Field (extLength) (Can be skipped, if the value is null)
	var extLength *uint8 = nil
	if m.ExtLength != nil {
		extLength = m.ExtLength
		// uint8 can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("ExtLength", *(extLength), -1))
	}
	// Array Field (propertyIdentifier)
	if m.PropertyIdentifier != nil {
		// Simple array base type uint8 will be hex dumped
		boxes = append(boxes, utils.BoxedDumpAnything("PropertyIdentifier", m.PropertyIdentifier))
		// Simple array base type uint8 will be rendered one by one
		arrayBoxes := make([]utils.AsciiBox, 0)
		for _, _element := range m.PropertyIdentifier {
			arrayBoxes = append(arrayBoxes, utils.BoxAnything("", _element, width-2))
		}
		boxes = append(boxes, utils.BoxBox("PropertyIdentifier", utils.AlignBoxes(arrayBoxes, width-4), 0))
	}
	// Const Field (openTag)
	boxes = append(boxes, utils.BoxAnything("OpenTag", uint8(0x2e), -1))
	// Complex field (case complex)
	boxes = append(boxes, m.Value.Box("value", width-2))
	// Const Field (closingTag)
	boxes = append(boxes, utils.BoxAnything("ClosingTag", uint8(0x2f), -1))
	return utils.BoxBox(boxName, utils.AlignBoxes(boxes, width-2), 0)
}
