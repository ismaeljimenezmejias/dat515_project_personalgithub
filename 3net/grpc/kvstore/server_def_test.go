package kvstore

import (
	"context"
	"slices"
	"testing"

	pb "dat515/3net/grpc/proto"
	"dat515/internal/test"

	"google.golang.org/protobuf/proto"
)

func testRequestSequence(t *testing.T, scoreHandler func()) {
	for _, tc := range testSequences {
		if ok := t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					test.PrintStackTrace(t.Name(), r)
					t.Fail()
				}
			}()

			server := NewKeyValueServicesServer()
			for _, kvMsg := range tc.kvMsgs {
				switch kvMsg := kvMsg.(type) {
				case *insert:
					gotInsertResp, err := server.Insert(context.Background(), kvMsg.req)
					if err != nil {
						t.Error(err)
					}
					if gotInsertResp.GetSuccess() != kvMsg.wantResp.GetSuccess() {
						t.Errorf("Insert(%s, %s) = %t, want %t", kvMsg.req.GetKey(), kvMsg.req.GetValue(), gotInsertResp.GetSuccess(), kvMsg.wantResp.GetSuccess())
					}
				case *lookup:
					gotLookupResp, err := server.Lookup(context.Background(), kvMsg.req)
					if err != nil {
						t.Error(err)
					}
					if gotLookupResp.GetValue() != kvMsg.wantResp.GetValue() {
						t.Errorf("Lookup(%s) = %s, want %s", kvMsg.req.GetKey(), gotLookupResp.GetValue(), kvMsg.wantResp.GetValue())
					}
				case *keys:
					gotKeysResp, err := server.Keys(context.Background(), kvMsg.req)
					if err != nil {
						t.Error(err)
					}
					if !slices.Equal(gotKeysResp.GetKeys(), kvMsg.wantResp.GetKeys()) {
						t.Errorf("Keys() = %v, want %v", gotKeysResp.GetKeys(), kvMsg.wantResp.GetKeys())
					}
				}
			}
		}); ok {
			scoreHandler()
		}
	}
}

type insert struct {
	req      *pb.InsertRequest
	wantResp *pb.InsertResponse
}

type lookup struct {
	req      *pb.LookupRequest
	wantResp *pb.LookupResponse
}

type keys struct {
	req      *pb.KeysRequest
	wantResp *pb.KeysResponse
}

var testSequences = []struct {
	name   string
	kvMsgs []any
}{
	{name: "LookupOnEmpty", kvMsgs: []any{&lookup{req: pb.LookupRequest_builder{Key: proto.String("1")}.Build(), wantResp: pb.LookupResponse_builder{Value: proto.String("")}.Build()}}},
	{name: "LookupAfterInsert", kvMsgs: []any{
		&insert{req: pb.InsertRequest_builder{Key: proto.String("1"), Value: proto.String("one")}.Build(), wantResp: pb.InsertResponse_builder{Success: proto.Bool(true)}.Build()},
		&lookup{req: pb.LookupRequest_builder{Key: proto.String("1")}.Build(), wantResp: pb.LookupResponse_builder{Value: proto.String("one")}.Build()},
	}},
	{name: "LookupAfterInsertEmptyValue", kvMsgs: []any{
		&insert{req: pb.InsertRequest_builder{Key: proto.String("1"), Value: proto.String("")}.Build(), wantResp: pb.InsertResponse_builder{Success: proto.Bool(true)}.Build()},
		&lookup{req: pb.LookupRequest_builder{Key: proto.String("1")}.Build(), wantResp: pb.LookupResponse_builder{Value: proto.String("")}.Build()},
	}},
	{name: "LookupAfterInsertEmptyKey", kvMsgs: []any{
		&insert{req: pb.InsertRequest_builder{Key: proto.String(""), Value: proto.String("one")}.Build(), wantResp: pb.InsertResponse_builder{Success: proto.Bool(true)}.Build()},
		&lookup{req: pb.LookupRequest_builder{Key: proto.String("")}.Build(), wantResp: pb.LookupResponse_builder{Value: proto.String("one")}.Build()},
	}},
	{name: "LookupAfterInsertTwoKeys", kvMsgs: []any{
		&insert{req: pb.InsertRequest_builder{Key: proto.String("1"), Value: proto.String("one")}.Build(), wantResp: pb.InsertResponse_builder{Success: proto.Bool(true)}.Build()},
		&insert{req: pb.InsertRequest_builder{Key: proto.String("2"), Value: proto.String("two")}.Build(), wantResp: pb.InsertResponse_builder{Success: proto.Bool(true)}.Build()},
		&lookup{req: pb.LookupRequest_builder{Key: proto.String("2")}.Build(), wantResp: pb.LookupResponse_builder{Value: proto.String("two")}.Build()},
		&lookup{req: pb.LookupRequest_builder{Key: proto.String("1")}.Build(), wantResp: pb.LookupResponse_builder{Value: proto.String("one")}.Build()},
		&lookup{req: pb.LookupRequest_builder{Key: proto.String("1")}.Build(), wantResp: pb.LookupResponse_builder{Value: proto.String("one")}.Build()},
		&lookup{req: pb.LookupRequest_builder{Key: proto.String("2")}.Build(), wantResp: pb.LookupResponse_builder{Value: proto.String("two")}.Build()},
	}},
	{name: "LookupAfterInsertTwoKeysSameKey", kvMsgs: []any{
		&insert{req: pb.InsertRequest_builder{Key: proto.String("1"), Value: proto.String("one")}.Build(), wantResp: pb.InsertResponse_builder{Success: proto.Bool(true)}.Build()},
		&insert{req: pb.InsertRequest_builder{Key: proto.String("1"), Value: proto.String("one again")}.Build(), wantResp: pb.InsertResponse_builder{Success: proto.Bool(true)}.Build()},
		&lookup{req: pb.LookupRequest_builder{Key: proto.String("1")}.Build(), wantResp: pb.LookupResponse_builder{Value: proto.String("one again")}.Build()},
	}},
	{name: "LookupAfterInsertThreeKeys", kvMsgs: []any{
		&insert{req: pb.InsertRequest_builder{Key: proto.String("1"), Value: proto.String("one")}.Build(), wantResp: pb.InsertResponse_builder{Success: proto.Bool(true)}.Build()},
		&insert{req: pb.InsertRequest_builder{Key: proto.String("2"), Value: proto.String("two")}.Build(), wantResp: pb.InsertResponse_builder{Success: proto.Bool(true)}.Build()},
		&insert{req: pb.InsertRequest_builder{Key: proto.String("3"), Value: proto.String("three")}.Build(), wantResp: pb.InsertResponse_builder{Success: proto.Bool(true)}.Build()},
		&lookup{req: pb.LookupRequest_builder{Key: proto.String("3")}.Build(), wantResp: pb.LookupResponse_builder{Value: proto.String("three")}.Build()},
		&lookup{req: pb.LookupRequest_builder{Key: proto.String("2")}.Build(), wantResp: pb.LookupResponse_builder{Value: proto.String("two")}.Build()},
		&lookup{req: pb.LookupRequest_builder{Key: proto.String("1")}.Build(), wantResp: pb.LookupResponse_builder{Value: proto.String("one")}.Build()},
	}},
	{name: "LookupAfterInsertTenKeys", kvMsgs: []any{
		&insert{req: pb.InsertRequest_builder{Key: proto.String("1"), Value: proto.String("one")}.Build(), wantResp: pb.InsertResponse_builder{Success: proto.Bool(true)}.Build()},
		&insert{req: pb.InsertRequest_builder{Key: proto.String("2"), Value: proto.String("two")}.Build(), wantResp: pb.InsertResponse_builder{Success: proto.Bool(true)}.Build()},
		&insert{req: pb.InsertRequest_builder{Key: proto.String("3"), Value: proto.String("three")}.Build(), wantResp: pb.InsertResponse_builder{Success: proto.Bool(true)}.Build()},
		&insert{req: pb.InsertRequest_builder{Key: proto.String("4"), Value: proto.String("four")}.Build(), wantResp: pb.InsertResponse_builder{Success: proto.Bool(true)}.Build()},
		&insert{req: pb.InsertRequest_builder{Key: proto.String("5"), Value: proto.String("five")}.Build(), wantResp: pb.InsertResponse_builder{Success: proto.Bool(true)}.Build()},
		&insert{req: pb.InsertRequest_builder{Key: proto.String("6"), Value: proto.String("six")}.Build(), wantResp: pb.InsertResponse_builder{Success: proto.Bool(true)}.Build()},
		&insert{req: pb.InsertRequest_builder{Key: proto.String("7"), Value: proto.String("seven")}.Build(), wantResp: pb.InsertResponse_builder{Success: proto.Bool(true)}.Build()},
		&insert{req: pb.InsertRequest_builder{Key: proto.String("8"), Value: proto.String("eight")}.Build(), wantResp: pb.InsertResponse_builder{Success: proto.Bool(true)}.Build()},
		&insert{req: pb.InsertRequest_builder{Key: proto.String("9"), Value: proto.String("nine")}.Build(), wantResp: pb.InsertResponse_builder{Success: proto.Bool(true)}.Build()},
		&insert{req: pb.InsertRequest_builder{Key: proto.String("10"), Value: proto.String("ten")}.Build(), wantResp: pb.InsertResponse_builder{Success: proto.Bool(true)}.Build()},
		&lookup{req: pb.LookupRequest_builder{Key: proto.String("1")}.Build(), wantResp: pb.LookupResponse_builder{Value: proto.String("one")}.Build()},
		&lookup{req: pb.LookupRequest_builder{Key: proto.String("2")}.Build(), wantResp: pb.LookupResponse_builder{Value: proto.String("two")}.Build()},
		&lookup{req: pb.LookupRequest_builder{Key: proto.String("3")}.Build(), wantResp: pb.LookupResponse_builder{Value: proto.String("three")}.Build()},
		&lookup{req: pb.LookupRequest_builder{Key: proto.String("4")}.Build(), wantResp: pb.LookupResponse_builder{Value: proto.String("four")}.Build()},
		&lookup{req: pb.LookupRequest_builder{Key: proto.String("5")}.Build(), wantResp: pb.LookupResponse_builder{Value: proto.String("five")}.Build()},
		&lookup{req: pb.LookupRequest_builder{Key: proto.String("6")}.Build(), wantResp: pb.LookupResponse_builder{Value: proto.String("six")}.Build()},
		&lookup{req: pb.LookupRequest_builder{Key: proto.String("7")}.Build(), wantResp: pb.LookupResponse_builder{Value: proto.String("seven")}.Build()},
		&lookup{req: pb.LookupRequest_builder{Key: proto.String("8")}.Build(), wantResp: pb.LookupResponse_builder{Value: proto.String("eight")}.Build()},
		&lookup{req: pb.LookupRequest_builder{Key: proto.String("9")}.Build(), wantResp: pb.LookupResponse_builder{Value: proto.String("nine")}.Build()},
		&lookup{req: pb.LookupRequest_builder{Key: proto.String("10")}.Build(), wantResp: pb.LookupResponse_builder{Value: proto.String("ten")}.Build()},
	}},
	{name: "KeysOnEmpty", kvMsgs: []any{&keys{req: pb.KeysRequest_builder{}.Build(), wantResp: pb.KeysResponse_builder{Keys: []string{}}.Build()}}},
	{name: "KeysAfterInsert", kvMsgs: []any{
		&insert{req: pb.InsertRequest_builder{Key: proto.String("1"), Value: proto.String("one")}.Build(), wantResp: pb.InsertResponse_builder{Success: proto.Bool(true)}.Build()},
		&keys{req: pb.KeysRequest_builder{}.Build(), wantResp: pb.KeysResponse_builder{Keys: []string{"1"}}.Build()},
	}},
	{name: "KeysAfterInsertTwoKeys", kvMsgs: []any{
		&insert{req: pb.InsertRequest_builder{Key: proto.String("1"), Value: proto.String("one")}.Build(), wantResp: pb.InsertResponse_builder{Success: proto.Bool(true)}.Build()},
		&insert{req: pb.InsertRequest_builder{Key: proto.String("2"), Value: proto.String("two")}.Build(), wantResp: pb.InsertResponse_builder{Success: proto.Bool(true)}.Build()},
		&keys{req: pb.KeysRequest_builder{}.Build(), wantResp: pb.KeysResponse_builder{Keys: []string{"1", "2"}}.Build()},
	}},
	{name: "KeysAfterInsertTwoKeysSameKey", kvMsgs: []any{
		&insert{req: pb.InsertRequest_builder{Key: proto.String("1"), Value: proto.String("one")}.Build(), wantResp: pb.InsertResponse_builder{Success: proto.Bool(true)}.Build()},
		&insert{req: pb.InsertRequest_builder{Key: proto.String("1"), Value: proto.String("one again")}.Build(), wantResp: pb.InsertResponse_builder{Success: proto.Bool(true)}.Build()},
		&keys{req: pb.KeysRequest_builder{}.Build(), wantResp: pb.KeysResponse_builder{Keys: []string{"1"}}.Build()},
	}},
	{name: "KeysAfterInsertThreeKeys", kvMsgs: []any{
		&insert{req: pb.InsertRequest_builder{Key: proto.String("1"), Value: proto.String("one")}.Build(), wantResp: pb.InsertResponse_builder{Success: proto.Bool(true)}.Build()},
		&insert{req: pb.InsertRequest_builder{Key: proto.String("2"), Value: proto.String("two")}.Build(), wantResp: pb.InsertResponse_builder{Success: proto.Bool(true)}.Build()},
		&insert{req: pb.InsertRequest_builder{Key: proto.String("3"), Value: proto.String("three")}.Build(), wantResp: pb.InsertResponse_builder{Success: proto.Bool(true)}.Build()},
		&keys{req: pb.KeysRequest_builder{}.Build(), wantResp: pb.KeysResponse_builder{Keys: []string{"1", "2", "3"}}.Build()},
	}},
	{name: "KeysAfterInsertTenKeys", kvMsgs: []any{
		&insert{req: pb.InsertRequest_builder{Key: proto.String("1"), Value: proto.String("one")}.Build(), wantResp: pb.InsertResponse_builder{Success: proto.Bool(true)}.Build()},
		&insert{req: pb.InsertRequest_builder{Key: proto.String("2"), Value: proto.String("two")}.Build(), wantResp: pb.InsertResponse_builder{Success: proto.Bool(true)}.Build()},
		&insert{req: pb.InsertRequest_builder{Key: proto.String("3"), Value: proto.String("three")}.Build(), wantResp: pb.InsertResponse_builder{Success: proto.Bool(true)}.Build()},
		&insert{req: pb.InsertRequest_builder{Key: proto.String("4"), Value: proto.String("four")}.Build(), wantResp: pb.InsertResponse_builder{Success: proto.Bool(true)}.Build()},
		&insert{req: pb.InsertRequest_builder{Key: proto.String("5"), Value: proto.String("five")}.Build(), wantResp: pb.InsertResponse_builder{Success: proto.Bool(true)}.Build()},
		&insert{req: pb.InsertRequest_builder{Key: proto.String("6"), Value: proto.String("six")}.Build(), wantResp: pb.InsertResponse_builder{Success: proto.Bool(true)}.Build()},
		&insert{req: pb.InsertRequest_builder{Key: proto.String("7"), Value: proto.String("seven")}.Build(), wantResp: pb.InsertResponse_builder{Success: proto.Bool(true)}.Build()},
		&insert{req: pb.InsertRequest_builder{Key: proto.String("8"), Value: proto.String("eight")}.Build(), wantResp: pb.InsertResponse_builder{Success: proto.Bool(true)}.Build()},
		&insert{req: pb.InsertRequest_builder{Key: proto.String("9"), Value: proto.String("nine")}.Build(), wantResp: pb.InsertResponse_builder{Success: proto.Bool(true)}.Build()},
		&insert{req: pb.InsertRequest_builder{Key: proto.String("10"), Value: proto.String("ten")}.Build(), wantResp: pb.InsertResponse_builder{Success: proto.Bool(true)}.Build()},
		&keys{req: pb.KeysRequest_builder{}.Build(), wantResp: pb.KeysResponse_builder{Keys: []string{"1", "10", "2", "3", "4", "5", "6", "7", "8", "9"}}.Build()},
	}},
}
