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
type BVLCOriginalUnicastNPDU struct {
	Npdu   *NPDU
	Parent *BVLC
}

// The corresponding interface
type IBVLCOriginalUnicastNPDU interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BVLCOriginalUnicastNPDU) BvlcFunction() uint8 {
	return 0x0A
}

func (m *BVLCOriginalUnicastNPDU) InitializeParent(parent *BVLC) {
}

func NewBVLCOriginalUnicastNPDU(npdu *NPDU) *BVLC {
	child := &BVLCOriginalUnicastNPDU{
		Npdu:   npdu,
		Parent: NewBVLC(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastBVLCOriginalUnicastNPDU(structType interface{}) *BVLCOriginalUnicastNPDU {
	castFunc := func(typ interface{}) *BVLCOriginalUnicastNPDU {
		if casted, ok := typ.(BVLCOriginalUnicastNPDU); ok {
			return &casted
		}
		if casted, ok := typ.(*BVLCOriginalUnicastNPDU); ok {
			return casted
		}
		if casted, ok := typ.(BVLC); ok {
			return CastBVLCOriginalUnicastNPDU(casted.Child)
		}
		if casted, ok := typ.(*BVLC); ok {
			return CastBVLCOriginalUnicastNPDU(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BVLCOriginalUnicastNPDU) GetTypeName() string {
	return "BVLCOriginalUnicastNPDU"
}

func (m *BVLCOriginalUnicastNPDU) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BVLCOriginalUnicastNPDU) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (npdu)
	lengthInBits += m.Npdu.LengthInBits()

	return lengthInBits
}

func (m *BVLCOriginalUnicastNPDU) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BVLCOriginalUnicastNPDUParse(io utils.ReadBuffer, bvlcLength uint16) (*BVLC, error) {

	// Simple Field (npdu)
	npdu, _npduErr := NPDUParse(io, uint16(bvlcLength)-uint16(uint16(4)))
	if _npduErr != nil {
		return nil, errors.Wrap(_npduErr, "Error parsing 'npdu' field")
	}

	// Create a partially initialized instance
	_child := &BVLCOriginalUnicastNPDU{
		Npdu:   npdu,
		Parent: &BVLC{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *BVLCOriginalUnicastNPDU) Serialize(io utils.WriteBuffer) error {
	ser := func() error {
		io.PushContext("BVLCOriginalUnicastNPDU")

		// Simple Field (npdu)
		_npduErr := m.Npdu.Serialize(io)
		if _npduErr != nil {
			return errors.Wrap(_npduErr, "Error serializing 'npdu' field")
		}

		io.PopContext("BVLCOriginalUnicastNPDU")
		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *BVLCOriginalUnicastNPDU) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "npdu":
				var data NPDU
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.Npdu = &data
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

func (m *BVLCOriginalUnicastNPDU) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.Npdu, xml.StartElement{Name: xml.Name{Local: "npdu"}}); err != nil {
		return err
	}
	return nil
}

func (m BVLCOriginalUnicastNPDU) String() string {
	return string(m.Box("", 120))
}

func (m BVLCOriginalUnicastNPDU) Box(name string, width int) utils.AsciiBox {
	boxName := "BVLCOriginalUnicastNPDU"
	if name != "" {
		boxName += "/" + name
	}
	childBoxer := func() []utils.AsciiBox {
		boxes := make([]utils.AsciiBox, 0)
		// Complex field (case complex)
		boxes = append(boxes, m.Npdu.Box("npdu", width-2))
		return boxes
	}
	return m.Parent.BoxParent(boxName, width, childBoxer)
}
