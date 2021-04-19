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
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type LDataInd struct {
	AdditionalInformationLength uint8
	AdditionalInformation       []*CEMIAdditionalInformation
	DataFrame                   *LDataFrame
	Parent                      *CEMI
}

// The corresponding interface
type ILDataInd interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *LDataInd) MessageCode() uint8 {
	return 0x29
}

func (m *LDataInd) InitializeParent(parent *CEMI) {
}

func NewLDataInd(additionalInformationLength uint8, additionalInformation []*CEMIAdditionalInformation, dataFrame *LDataFrame) *CEMI {
	child := &LDataInd{
		AdditionalInformationLength: additionalInformationLength,
		AdditionalInformation:       additionalInformation,
		DataFrame:                   dataFrame,
		Parent:                      NewCEMI(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastLDataInd(structType interface{}) *LDataInd {
	castFunc := func(typ interface{}) *LDataInd {
		if casted, ok := typ.(LDataInd); ok {
			return &casted
		}
		if casted, ok := typ.(*LDataInd); ok {
			return casted
		}
		if casted, ok := typ.(CEMI); ok {
			return CastLDataInd(casted.Child)
		}
		if casted, ok := typ.(*CEMI); ok {
			return CastLDataInd(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *LDataInd) GetTypeName() string {
	return "LDataInd"
}

func (m *LDataInd) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *LDataInd) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (additionalInformationLength)
	lengthInBits += 8

	// Array field
	if len(m.AdditionalInformation) > 0 {
		for _, element := range m.AdditionalInformation {
			lengthInBits += element.LengthInBits()
		}
	}

	// Simple field (dataFrame)
	lengthInBits += m.DataFrame.LengthInBits()

	return lengthInBits
}

func (m *LDataInd) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func LDataIndParse(io utils.ReadBuffer) (*CEMI, error) {

	// Simple Field (additionalInformationLength)
	additionalInformationLength, _additionalInformationLengthErr := io.ReadUint8(8)
	if _additionalInformationLengthErr != nil {
		return nil, errors.Wrap(_additionalInformationLengthErr, "Error parsing 'additionalInformationLength' field")
	}

	// Array field (additionalInformation)
	// Length array
	additionalInformation := make([]*CEMIAdditionalInformation, 0)
	_additionalInformationLength := additionalInformationLength
	_additionalInformationEndPos := io.GetPos() + uint16(_additionalInformationLength)
	for io.GetPos() < _additionalInformationEndPos {
		_item, _err := CEMIAdditionalInformationParse(io)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'additionalInformation' field")
		}
		additionalInformation = append(additionalInformation, _item)
	}

	// Simple Field (dataFrame)
	dataFrame, _dataFrameErr := LDataFrameParse(io)
	if _dataFrameErr != nil {
		return nil, errors.Wrap(_dataFrameErr, "Error parsing 'dataFrame' field")
	}

	// Create a partially initialized instance
	_child := &LDataInd{
		AdditionalInformationLength: additionalInformationLength,
		AdditionalInformation:       additionalInformation,
		DataFrame:                   dataFrame,
		Parent:                      &CEMI{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *LDataInd) Serialize(io utils.WriteBuffer) error {
	ser := func() error {
		io.PushContext("LDataInd")

		// Simple Field (additionalInformationLength)
		additionalInformationLength := uint8(m.AdditionalInformationLength)
		_additionalInformationLengthErr := io.WriteUint8("additionalInformationLength", 8, (additionalInformationLength))
		if _additionalInformationLengthErr != nil {
			return errors.Wrap(_additionalInformationLengthErr, "Error serializing 'additionalInformationLength' field")
		}

		// Array Field (additionalInformation)
		if m.AdditionalInformation != nil {
			for _, _element := range m.AdditionalInformation {
				_elementErr := _element.Serialize(io)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'additionalInformation' field")
				}
			}
		}

		// Simple Field (dataFrame)
		_dataFrameErr := m.DataFrame.Serialize(io)
		if _dataFrameErr != nil {
			return errors.Wrap(_dataFrameErr, "Error serializing 'dataFrame' field")
		}

		io.PopContext("LDataInd")
		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *LDataInd) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "additionalInformationLength":
				var data uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.AdditionalInformationLength = data
			case "additionalInformation":
			arrayLoop:
				for {
					token, err = d.Token()
					switch token.(type) {
					case xml.StartElement:
						tok := token.(xml.StartElement)
						var dt *CEMIAdditionalInformation
						if err := d.DecodeElement(&dt, &tok); err != nil {
							return err
						}
						m.AdditionalInformation = append(m.AdditionalInformation, dt)
					default:
						break arrayLoop
					}
				}
			case "dataFrame":
				var dt *LDataFrame
				if err := d.DecodeElement(&dt, &tok); err != nil {
					if err == io.EOF {
						continue
					}
					return err
				}
				m.DataFrame = dt
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

func (m *LDataInd) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.AdditionalInformationLength, xml.StartElement{Name: xml.Name{Local: "additionalInformationLength"}}); err != nil {
		return err
	}
	if err := e.EncodeToken(xml.StartElement{Name: xml.Name{Local: "additionalInformation"}}); err != nil {
		return err
	}
	for _, arrayElement := range m.AdditionalInformation {
		if err := e.EncodeElement(arrayElement, xml.StartElement{Name: xml.Name{Local: "additionalInformation"}}); err != nil {
			return err
		}
	}
	if err := e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "additionalInformation"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.DataFrame, xml.StartElement{Name: xml.Name{Local: "dataFrame"}}); err != nil {
		return err
	}
	return nil
}

func (m LDataInd) String() string {
	return string(m.Box("", 120))
}

func (m LDataInd) Box(name string, width int) utils.AsciiBox {
	boxName := "LDataInd"
	if name != "" {
		boxName += "/" + name
	}
	childBoxer := func() []utils.AsciiBox {
		boxes := make([]utils.AsciiBox, 0)
		// Simple field (case simple)
		// uint8 can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("AdditionalInformationLength", m.AdditionalInformationLength, -1))
		// Array Field (additionalInformation)
		if m.AdditionalInformation != nil {
			// Complex array base type
			arrayBoxes := make([]utils.AsciiBox, 0)
			for _, _element := range m.AdditionalInformation {
				arrayBoxes = append(arrayBoxes, utils.BoxAnything("", _element, width-2))
			}
			boxes = append(boxes, utils.BoxBox("AdditionalInformation", utils.AlignBoxes(arrayBoxes, width-4), 0))
		}
		// Complex field (case complex)
		boxes = append(boxes, m.DataFrame.Box("dataFrame", width-2))
		return boxes
	}
	return m.Parent.BoxParent(boxName, width, childBoxer)
}
