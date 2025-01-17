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
	"encoding/hex"
	"encoding/xml"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"io"
	"strings"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type AdsReadResponse struct {
	Result ReturnCode
	Data   []int8
	Parent *AdsData
}

// The corresponding interface
type IAdsReadResponse interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *AdsReadResponse) CommandId() CommandId {
	return CommandId_ADS_READ
}

func (m *AdsReadResponse) Response() bool {
	return true
}

func (m *AdsReadResponse) InitializeParent(parent *AdsData) {
}

func NewAdsReadResponse(result ReturnCode, data []int8) *AdsData {
	child := &AdsReadResponse{
		Result: result,
		Data:   data,
		Parent: NewAdsData(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastAdsReadResponse(structType interface{}) *AdsReadResponse {
	castFunc := func(typ interface{}) *AdsReadResponse {
		if casted, ok := typ.(AdsReadResponse); ok {
			return &casted
		}
		if casted, ok := typ.(*AdsReadResponse); ok {
			return casted
		}
		if casted, ok := typ.(AdsData); ok {
			return CastAdsReadResponse(casted.Child)
		}
		if casted, ok := typ.(*AdsData); ok {
			return CastAdsReadResponse(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *AdsReadResponse) GetTypeName() string {
	return "AdsReadResponse"
}

func (m *AdsReadResponse) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *AdsReadResponse) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (result)
	lengthInBits += 32

	// Implicit Field (length)
	lengthInBits += 32

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *AdsReadResponse) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func AdsReadResponseParse(io utils.ReadBuffer) (*AdsData, error) {

	// Simple Field (result)
	result, _resultErr := ReturnCodeParse(io)
	if _resultErr != nil {
		return nil, errors.Wrap(_resultErr, "Error parsing 'result' field")
	}

	// Implicit Field (length) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	length, _lengthErr := io.ReadUint32(32)
	_ = length
	if _lengthErr != nil {
		return nil, errors.Wrap(_lengthErr, "Error parsing 'length' field")
	}

	// Array field (data)
	// Count array
	data := make([]int8, length)
	for curItem := uint16(0); curItem < uint16(length); curItem++ {
		_item, _err := io.ReadInt8(8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'data' field")
		}
		data[curItem] = _item
	}

	// Create a partially initialized instance
	_child := &AdsReadResponse{
		Result: result,
		Data:   data,
		Parent: &AdsData{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *AdsReadResponse) Serialize(io utils.WriteBuffer) error {
	ser := func() error {
		io.PushContext("AdsReadResponse")

		// Simple Field (result)
		_resultErr := m.Result.Serialize(io)
		if _resultErr != nil {
			return errors.Wrap(_resultErr, "Error serializing 'result' field")
		}

		// Implicit Field (length) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		length := uint32(uint32(len(m.Data)))
		_lengthErr := io.WriteUint32("length", 32, (length))
		if _lengthErr != nil {
			return errors.Wrap(_lengthErr, "Error serializing 'length' field")
		}

		// Array Field (data)
		if m.Data != nil {
			for _, _element := range m.Data {
				_elementErr := io.WriteInt8("", 8, _element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'data' field")
				}
			}
		}

		io.PopContext("AdsReadResponse")
		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *AdsReadResponse) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	foundContent := false
	token = start
	for {
		switch token.(type) {
		case xml.StartElement:
			foundContent = true
			tok := token.(xml.StartElement)
			switch tok.Name.Local {
			case "result":
				var data ReturnCode
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.Result = data
			case "data":
				var _encoded string
				if err := d.DecodeElement(&_encoded, &tok); err != nil {
					return err
				}
				_decoded, err := hex.DecodeString(_encoded)
				_len := len(_decoded)
				if err != nil {
					return err
				}
				m.Data = utils.ByteArrayToInt8Array(_decoded[0:_len])
			}
		}
		token, err = d.Token()
		if err != nil {
			if err == io.EOF && foundContent {
				return nil
			}
			return err
		}
	}
}

func (m *AdsReadResponse) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.Result, xml.StartElement{Name: xml.Name{Local: "result"}}); err != nil {
		return err
	}
	_encodedData := hex.EncodeToString(utils.Int8ArrayToByteArray(m.Data))
	_encodedData = strings.ToUpper(_encodedData)
	if err := e.EncodeElement(_encodedData, xml.StartElement{Name: xml.Name{Local: "data"}}); err != nil {
		return err
	}
	return nil
}

func (m AdsReadResponse) String() string {
	return string(m.Box("", 120))
}

func (m AdsReadResponse) Box(name string, width int) utils.AsciiBox {
	boxName := "AdsReadResponse"
	if name != "" {
		boxName += "/" + name
	}
	childBoxer := func() []utils.AsciiBox {
		boxes := make([]utils.AsciiBox, 0)
		// Complex field (case complex)
		boxes = append(boxes, m.Result.Box("result", width-2))
		// Implicit Field (length)
		length := uint32(uint32(len(m.Data)))
		// uint32 can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("Length", length, -1))
		// Array Field (data)
		if m.Data != nil {
			// Simple array base type int8 will be rendered one by one
			arrayBoxes := make([]utils.AsciiBox, 0)
			for _, _element := range m.Data {
				arrayBoxes = append(arrayBoxes, utils.BoxAnything("", _element, width-2))
			}
			boxes = append(boxes, utils.BoxBox("Data", utils.AlignBoxes(arrayBoxes, width-4), 0))
		}
		return boxes
	}
	return m.Parent.BoxParent(boxName, width, childBoxer)
}
