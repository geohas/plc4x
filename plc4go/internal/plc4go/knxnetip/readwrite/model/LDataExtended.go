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
type LDataExtended struct {
	GroupAddress        bool
	HopCount            uint8
	ExtendedFrameFormat uint8
	SourceAddress       *KnxAddress
	DestinationAddress  []int8
	Apdu                *Apdu
	Parent              *LDataFrame
}

// The corresponding interface
type ILDataExtended interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *LDataExtended) NotAckFrame() bool {
	return true
}

func (m *LDataExtended) Polling() bool {
	return false
}

func (m *LDataExtended) InitializeParent(parent *LDataFrame, frameType bool, notRepeated bool, priority CEMIPriority, acknowledgeRequested bool, errorFlag bool) {
	m.Parent.FrameType = frameType
	m.Parent.NotRepeated = notRepeated
	m.Parent.Priority = priority
	m.Parent.AcknowledgeRequested = acknowledgeRequested
	m.Parent.ErrorFlag = errorFlag
}

func NewLDataExtended(groupAddress bool, hopCount uint8, extendedFrameFormat uint8, sourceAddress *KnxAddress, destinationAddress []int8, apdu *Apdu, frameType bool, notRepeated bool, priority CEMIPriority, acknowledgeRequested bool, errorFlag bool) *LDataFrame {
	child := &LDataExtended{
		GroupAddress:        groupAddress,
		HopCount:            hopCount,
		ExtendedFrameFormat: extendedFrameFormat,
		SourceAddress:       sourceAddress,
		DestinationAddress:  destinationAddress,
		Apdu:                apdu,
		Parent:              NewLDataFrame(frameType, notRepeated, priority, acknowledgeRequested, errorFlag),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastLDataExtended(structType interface{}) *LDataExtended {
	castFunc := func(typ interface{}) *LDataExtended {
		if casted, ok := typ.(LDataExtended); ok {
			return &casted
		}
		if casted, ok := typ.(*LDataExtended); ok {
			return casted
		}
		if casted, ok := typ.(LDataFrame); ok {
			return CastLDataExtended(casted.Child)
		}
		if casted, ok := typ.(*LDataFrame); ok {
			return CastLDataExtended(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *LDataExtended) GetTypeName() string {
	return "LDataExtended"
}

func (m *LDataExtended) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *LDataExtended) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (groupAddress)
	lengthInBits += 1

	// Simple field (hopCount)
	lengthInBits += 3

	// Simple field (extendedFrameFormat)
	lengthInBits += 4

	// Simple field (sourceAddress)
	lengthInBits += m.SourceAddress.LengthInBits()

	// Array field
	if len(m.DestinationAddress) > 0 {
		lengthInBits += 8 * uint16(len(m.DestinationAddress))
	}

	// Implicit Field (dataLength)
	lengthInBits += 8

	// Simple field (apdu)
	lengthInBits += m.Apdu.LengthInBits()

	return lengthInBits
}

func (m *LDataExtended) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func LDataExtendedParse(io utils.ReadBuffer) (*LDataFrame, error) {

	// Simple Field (groupAddress)
	groupAddress, _groupAddressErr := io.ReadBit()
	if _groupAddressErr != nil {
		return nil, errors.Wrap(_groupAddressErr, "Error parsing 'groupAddress' field")
	}

	// Simple Field (hopCount)
	hopCount, _hopCountErr := io.ReadUint8(3)
	if _hopCountErr != nil {
		return nil, errors.Wrap(_hopCountErr, "Error parsing 'hopCount' field")
	}

	// Simple Field (extendedFrameFormat)
	extendedFrameFormat, _extendedFrameFormatErr := io.ReadUint8(4)
	if _extendedFrameFormatErr != nil {
		return nil, errors.Wrap(_extendedFrameFormatErr, "Error parsing 'extendedFrameFormat' field")
	}

	// Simple Field (sourceAddress)
	sourceAddress, _sourceAddressErr := KnxAddressParse(io)
	if _sourceAddressErr != nil {
		return nil, errors.Wrap(_sourceAddressErr, "Error parsing 'sourceAddress' field")
	}

	// Array field (destinationAddress)
	// Count array
	destinationAddress := make([]int8, uint16(2))
	for curItem := uint16(0); curItem < uint16(uint16(2)); curItem++ {
		_item, _err := io.ReadInt8(8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'destinationAddress' field")
		}
		destinationAddress[curItem] = _item
	}

	// Implicit Field (dataLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	dataLength, _dataLengthErr := io.ReadUint8(8)
	_ = dataLength
	if _dataLengthErr != nil {
		return nil, errors.Wrap(_dataLengthErr, "Error parsing 'dataLength' field")
	}

	// Simple Field (apdu)
	apdu, _apduErr := ApduParse(io, dataLength)
	if _apduErr != nil {
		return nil, errors.Wrap(_apduErr, "Error parsing 'apdu' field")
	}

	// Create a partially initialized instance
	_child := &LDataExtended{
		GroupAddress:        groupAddress,
		HopCount:            hopCount,
		ExtendedFrameFormat: extendedFrameFormat,
		SourceAddress:       sourceAddress,
		DestinationAddress:  destinationAddress,
		Apdu:                apdu,
		Parent:              &LDataFrame{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *LDataExtended) Serialize(io utils.WriteBuffer) error {
	ser := func() error {
		io.PushContext("LDataExtended")

		// Simple Field (groupAddress)
		groupAddress := bool(m.GroupAddress)
		_groupAddressErr := io.WriteBit("groupAddress", (groupAddress))
		if _groupAddressErr != nil {
			return errors.Wrap(_groupAddressErr, "Error serializing 'groupAddress' field")
		}

		// Simple Field (hopCount)
		hopCount := uint8(m.HopCount)
		_hopCountErr := io.WriteUint8("hopCount", 3, (hopCount))
		if _hopCountErr != nil {
			return errors.Wrap(_hopCountErr, "Error serializing 'hopCount' field")
		}

		// Simple Field (extendedFrameFormat)
		extendedFrameFormat := uint8(m.ExtendedFrameFormat)
		_extendedFrameFormatErr := io.WriteUint8("extendedFrameFormat", 4, (extendedFrameFormat))
		if _extendedFrameFormatErr != nil {
			return errors.Wrap(_extendedFrameFormatErr, "Error serializing 'extendedFrameFormat' field")
		}

		// Simple Field (sourceAddress)
		_sourceAddressErr := m.SourceAddress.Serialize(io)
		if _sourceAddressErr != nil {
			return errors.Wrap(_sourceAddressErr, "Error serializing 'sourceAddress' field")
		}

		// Array Field (destinationAddress)
		if m.DestinationAddress != nil {
			for _, _element := range m.DestinationAddress {
				_elementErr := io.WriteInt8("", 8, _element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'destinationAddress' field")
				}
			}
		}

		// Implicit Field (dataLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		dataLength := uint8(uint8(m.Apdu.LengthInBytes()) - uint8(uint8(1)))
		_dataLengthErr := io.WriteUint8("dataLength", 8, (dataLength))
		if _dataLengthErr != nil {
			return errors.Wrap(_dataLengthErr, "Error serializing 'dataLength' field")
		}

		// Simple Field (apdu)
		_apduErr := m.Apdu.Serialize(io)
		if _apduErr != nil {
			return errors.Wrap(_apduErr, "Error serializing 'apdu' field")
		}

		io.PopContext("LDataExtended")
		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *LDataExtended) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "groupAddress":
				var data bool
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.GroupAddress = data
			case "hopCount":
				var data uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.HopCount = data
			case "extendedFrameFormat":
				var data uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.ExtendedFrameFormat = data
			case "sourceAddress":
				var data KnxAddress
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.SourceAddress = &data
			case "destinationAddress":
				var _encoded string
				if err := d.DecodeElement(&_encoded, &tok); err != nil {
					return err
				}
				_decoded, err := hex.DecodeString(_encoded)
				_len := len(_decoded)
				if err != nil {
					return err
				}
				m.DestinationAddress = utils.ByteArrayToInt8Array(_decoded[0:_len])
			case "apdu":
				var dt *Apdu
				if err := d.DecodeElement(&dt, &tok); err != nil {
					if err == io.EOF {
						continue
					}
					return err
				}
				m.Apdu = dt
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

func (m *LDataExtended) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.GroupAddress, xml.StartElement{Name: xml.Name{Local: "groupAddress"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.HopCount, xml.StartElement{Name: xml.Name{Local: "hopCount"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.ExtendedFrameFormat, xml.StartElement{Name: xml.Name{Local: "extendedFrameFormat"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.SourceAddress, xml.StartElement{Name: xml.Name{Local: "sourceAddress"}}); err != nil {
		return err
	}
	_encodedDestinationAddress := hex.EncodeToString(utils.Int8ArrayToByteArray(m.DestinationAddress))
	_encodedDestinationAddress = strings.ToUpper(_encodedDestinationAddress)
	if err := e.EncodeElement(_encodedDestinationAddress, xml.StartElement{Name: xml.Name{Local: "destinationAddress"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.Apdu, xml.StartElement{Name: xml.Name{Local: "apdu"}}); err != nil {
		return err
	}
	return nil
}

func (m LDataExtended) String() string {
	return string(m.Box("", 120))
}

func (m LDataExtended) Box(name string, width int) utils.AsciiBox {
	boxName := "LDataExtended"
	if name != "" {
		boxName += "/" + name
	}
	childBoxer := func() []utils.AsciiBox {
		boxes := make([]utils.AsciiBox, 0)
		// Simple field (case simple)
		// bool can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("GroupAddress", m.GroupAddress, -1))
		// Simple field (case simple)
		// uint8 can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("HopCount", m.HopCount, -1))
		// Simple field (case simple)
		// uint8 can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("ExtendedFrameFormat", m.ExtendedFrameFormat, -1))
		// Complex field (case complex)
		boxes = append(boxes, m.SourceAddress.Box("sourceAddress", width-2))
		// Array Field (destinationAddress)
		if m.DestinationAddress != nil {
			// Simple array base type int8 will be rendered one by one
			arrayBoxes := make([]utils.AsciiBox, 0)
			for _, _element := range m.DestinationAddress {
				arrayBoxes = append(arrayBoxes, utils.BoxAnything("", _element, width-2))
			}
			boxes = append(boxes, utils.BoxBox("DestinationAddress", utils.AlignBoxes(arrayBoxes, width-4), 0))
		}
		// Implicit Field (dataLength)
		dataLength := uint8(uint8(m.Apdu.LengthInBytes()) - uint8(uint8(1)))
		// uint8 can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("DataLength", dataLength, -1))
		// Complex field (case complex)
		boxes = append(boxes, m.Apdu.Box("apdu", width-2))
		return boxes
	}
	return m.Parent.BoxParent(boxName, width, childBoxer)
}
