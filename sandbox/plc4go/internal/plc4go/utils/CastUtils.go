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
package utils

func Int8ToUint8(input []int8) []uint8 {
    output := make([]uint8, len(input))
    for i, _val := range input {
        output[i] = uint8(_val)
    }
    return output
}

func Int8ToByte(input []int8) []byte {
    output := make([]byte, len(input))
    for i, _val := range input {
        output[i] = byte(_val)
    }
    return output
}