package encoding_test

import (
	"encoding/hex"
	"testing"

	"github.com/jb0n/gobacnet/encoding"
	bactype "github.com/jb0n/gobacnet/types"
)

func TestParseObjectListMessage(t *testing.T) {
	t.Parallel()

	// this is a raw response from a Delta Controls DSC_1616E object-list request
	hb := "810a011d010030020e0c020000641e294c4ec400000001c402000064c402800001c403c00001c403c00002c403c00003c403c00004c403c00005c403c00006c403c00007c403c00008c403c00009c423800001c42c000015c42d400064c443800001c445800001c425800001c425800002c425800003c425800004c425800005c425800006c425800007c425800008c425800009c42580000ac424c00001c424c00002c424c00003c424c00004c424c00005c424c00006c424c00007c424c00008c424c00009c424c0000ac424c0000bc424c0000cc424c0000dc424c0000ec424c0000fc424c00010c424c00011c424c00012c424c00013c424c00014c426000001c442800001c448800001c444000001c44a800001c44dc000014f1f" //nolint
	buf, err := hex.DecodeString(hb)
	if err != nil {
		t.Fatal("can't parse hex bytes of obj list message")
	}
	dec := encoding.NewDecoder(buf)
	var header bactype.BVLC
	err = dec.BVLC(&header)
	if err != nil {
		t.Fatal("can't parse BVLC of obj list message")
	}

	var npdu bactype.NPDU
	err = dec.NPDU(&npdu)
	if err != nil {
		t.Fatal("can't parse NPDU of obj list message")
	}

	var apdu bactype.APDU
	err = dec.APDU(&apdu)
	if err != nil {
		t.Fatal("can't parse APDU of obj list message")
	}

	var rp bactype.ReadMultipleProperty
	err = dec.ReadMultiplePropertyAck(&rp)
	if err != nil {
		t.Fatal("can't parse ACK of obj list message")
	}

	const expectedLen = 53
	if len(rp.Objects) != expectedLen {
		t.Fatalf("expected to get %d objects from test messge, but got %d instead", expectedLen, len(rp.Objects))
	}
}
