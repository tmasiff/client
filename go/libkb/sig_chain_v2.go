package libkb

import (
	"github.com/keybase/go-codec/codec"
)

type SigchainTypesV2 int

const (
	SigchainTypeV2Eldest                      SigchainTypesV2 = 1
	SigchainTypeV2WebServiceBinding           SigchainTypesV2 = 2
	SigchainTypeV2Track                       SigchainTypesV2 = 3
	SigchainTypeV2Untrack                     SigchainTypesV2 = 4
	SigchainTypeV2Revoke                      SigchainTypesV2 = 5
	SigchainTypeV2Cryptocurrency              SigchainTypesV2 = 6
	SigchainTypeV2Announcement                SigchainTypesV2 = 7
	SigchainTypeV2Device                      SigchainTypesV2 = 8
	SigchainTypeV2WebServiceBindingWithRevoke SigchainTypesV2 = 9
	SigchainTypeV2CryptocurrencyWithRevoke    SigchainTypesV2 = 10
	SigchainTypeV2Sibkey                      SigchainTypesV2 = 11
	SigchainTypeV2Subkey                      SigchainTypesV2 = 12
	SigchainTypeV2PGPUpdate                   SigchainTypesV2 = 13
)

// OuterLinkV2 is the second version of Keybase sigchain signatures.
type OuterLinkV2 struct {
	_struct  bool   `codec:",toarray"`
	version  int    `codec:"version"`
	seqno    Seqno  `codec:"seqno"`
	prev     LinkID `codec:"prev"`
	body     []byte `codec:"body"`
	linkType int    `codec:"type"`
}

func (o OuterLinkV2) Encode() ([]byte, error) {
	var encoded []byte
	err := codec.NewEncoderBytes(&encoded, codecHandle()).Encode(o)
	return encoded, err
}
