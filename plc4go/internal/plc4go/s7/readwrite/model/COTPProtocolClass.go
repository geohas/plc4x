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
	"io"
)

// Code generated by build-utils. DO NOT EDIT.

type COTPProtocolClass int8

type ICOTPProtocolClass interface {
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

const (
	COTPProtocolClass_CLASS_0 COTPProtocolClass = 0x00
	COTPProtocolClass_CLASS_1 COTPProtocolClass = 0x10
	COTPProtocolClass_CLASS_2 COTPProtocolClass = 0x20
	COTPProtocolClass_CLASS_3 COTPProtocolClass = 0x30
	COTPProtocolClass_CLASS_4 COTPProtocolClass = 0x40
)

var COTPProtocolClassValues []COTPProtocolClass

func init() {
	COTPProtocolClassValues = []COTPProtocolClass{
		COTPProtocolClass_CLASS_0,
		COTPProtocolClass_CLASS_1,
		COTPProtocolClass_CLASS_2,
		COTPProtocolClass_CLASS_3,
		COTPProtocolClass_CLASS_4,
	}
}

func COTPProtocolClassByValue(value int8) COTPProtocolClass {
	switch value {
	case 0x00:
		return COTPProtocolClass_CLASS_0
	case 0x10:
		return COTPProtocolClass_CLASS_1
	case 0x20:
		return COTPProtocolClass_CLASS_2
	case 0x30:
		return COTPProtocolClass_CLASS_3
	case 0x40:
		return COTPProtocolClass_CLASS_4
	}
	return 0
}

func COTPProtocolClassByName(value string) COTPProtocolClass {
	switch value {
	case "CLASS_0":
		return COTPProtocolClass_CLASS_0
	case "CLASS_1":
		return COTPProtocolClass_CLASS_1
	case "CLASS_2":
		return COTPProtocolClass_CLASS_2
	case "CLASS_3":
		return COTPProtocolClass_CLASS_3
	case "CLASS_4":
		return COTPProtocolClass_CLASS_4
	}
	return 0
}

func CastCOTPProtocolClass(structType interface{}) COTPProtocolClass {
	castFunc := func(typ interface{}) COTPProtocolClass {
		if sCOTPProtocolClass, ok := typ.(COTPProtocolClass); ok {
			return sCOTPProtocolClass
		}
		return 0
	}
	return castFunc(structType)
}

func (m COTPProtocolClass) LengthInBits() uint16 {
	return 8
}

func (m COTPProtocolClass) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func COTPProtocolClassParse(io utils.ReadBuffer) (COTPProtocolClass, error) {
	val, err := io.ReadInt8(8)
	if err != nil {
		return 0, nil
	}
	return COTPProtocolClassByValue(val), nil
}

func (e COTPProtocolClass) Serialize(io utils.WriteBuffer) error {
	err := io.WriteInt8("COTPProtocolClass", 8, int8(e))
	return err
}

func (m *COTPProtocolClass) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	for {
		token, err = d.Token()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		switch token.(type) {
		case xml.CharData:
			tok := token.(xml.CharData)
			*m = COTPProtocolClassByName(string(tok))
		}
	}
}

func (m COTPProtocolClass) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.String(), start); err != nil {
		return err
	}
	return nil
}

func (e COTPProtocolClass) name() string {
	switch e {
	case COTPProtocolClass_CLASS_0:
		return "CLASS_0"
	case COTPProtocolClass_CLASS_1:
		return "CLASS_1"
	case COTPProtocolClass_CLASS_2:
		return "CLASS_2"
	case COTPProtocolClass_CLASS_3:
		return "CLASS_3"
	case COTPProtocolClass_CLASS_4:
		return "CLASS_4"
	}
	return ""
}

func (e COTPProtocolClass) String() string {
	return e.name()
}

func (m COTPProtocolClass) Box(s string, i int) utils.AsciiBox {
	boxName := "COTPProtocolClass"
	if s != "" {
		boxName += "/" + s
	}
	return utils.BoxString(boxName, fmt.Sprintf("%#0*x %s", 2, int8(m), m.name()), -1)
}
