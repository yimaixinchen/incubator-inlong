/**
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 * <p>
 * http://www.apache.org/licenses/LICENSE-2.0
 * <p>
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package codec defines the encoding and decoding logic between TubeMQ.
// If the protocol of encoding and decoding is changed, only this package
// will need to be changed.
package codec

import (
	"bufio"
	"encoding/binary"
	"errors"
	"io"
)

const (
	RPCProtocolBeginToken uint32 = 0xFF7FF4FE
	RPCMaxBufferSize      uint32 = 8192
	frameHeadLen          uint32 = 8
	maxBufferSize         int    = 128 * 1024
	defaultMsgSize        int    = 4096
	dataLen               uint32 = 4
	listSizeLen           uint32 = 4
	serialNoLen           uint32 = 4
	beginTokenLen         uint32 = 4
)

// TransportResponse is the abstraction of the transport response.
type TransportResponse interface {
	// GetSerialNo returns the `serialNo` of the corresponding request.
	GetSerialNo() uint32
	// GetResponseBuf returns the body of the response.
	GetResponseBuf() []byte
}

// Decoder is the abstraction of the decoder which is used to decode the response.
type Decoder interface {
	// Decode will decode the response to frame head and body.
	Decode() (TransportResponse, error)
}

// TubeMQDecoder is the implementation of the decoder of response from TubeMQ.
type TubeMQDecoder struct {
	reader io.Reader
	msg    []byte
}

// New will return a default TubeMQDecoder.
func New(reader io.Reader) *TubeMQDecoder {
	bufferReader := bufio.NewReaderSize(reader, maxBufferSize)
	return &TubeMQDecoder{
		msg:    make([]byte, defaultMsgSize),
		reader: bufferReader,
	}
}

// Decode will decode the response from TubeMQ to TransportResponse according to
// the RPC protocol of TubeMQ.
func (t *TubeMQDecoder) Decode() (TransportResponse, error) {
	num, err := io.ReadFull(t.reader, t.msg[:frameHeadLen])
	if err != nil {
		return nil, err
	}
	if num != int(frameHeadLen) {
		return nil, errors.New("framer: read frame header num invalid")
	}
	token := binary.BigEndian.Uint32(t.msg[:beginTokenLen])
	if token != RPCProtocolBeginToken {
		return nil, errors.New("framer: read framer rpc protocol begin token not match")
	}
	num, err = io.ReadFull(t.reader, t.msg[frameHeadLen:frameHeadLen+listSizeLen])
	if num != int(listSizeLen) {
		return nil, errors.New("framer: read invalid list size num")
	}
	listSize := binary.BigEndian.Uint32(t.msg[frameHeadLen : frameHeadLen+listSizeLen])
	totalLen := int(frameHeadLen)
	size := make([]byte, 4)
	for i := 0; i < int(listSize); i++ {
		n, err := io.ReadFull(t.reader, size)
		if err != nil {
			return nil, err
		}
		if n != int(dataLen) {
			return nil, errors.New("framer: read invalid size")
		}

		s := int(binary.BigEndian.Uint32(size))
		if totalLen+s > len(t.msg) {
			data := t.msg[:totalLen]
			t.msg = make([]byte, totalLen+s)
			copy(t.msg, data[:])
		}

		num, err = io.ReadFull(t.reader, t.msg[totalLen:totalLen+s])
		if err != nil {
			return nil, err
		}
		if num != s {
			return nil, errors.New("framer: read invalid data")
		}
		totalLen += s
	}

	data := make([]byte, totalLen-int(frameHeadLen))
	copy(data, t.msg[frameHeadLen:totalLen])

	return TubeMQResponse{
		serialNo:    binary.BigEndian.Uint32(t.msg[beginTokenLen : beginTokenLen+serialNoLen]),
		responseBuf: data,
	}, nil
}

// TubeMQRequest is the implementation of TubeMQ request.
type TubeMQRequest struct {
	serialNo uint32
	req      []byte
}

// TubeMQResponse is the TubeMQ implementation of TransportResponse.
type TubeMQResponse struct {
	serialNo    uint32
	responseBuf []byte
}

// GetSerialNo will return the SerialNo of TubeMQResponse.
func (t TubeMQResponse) GetSerialNo() uint32 {
	return t.serialNo
}

// GetResponseBuf will return the body of TubeMQResponse.
func (t TubeMQResponse) GetResponseBuf() []byte {
	return t.responseBuf
}