package test

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"testing"
)
import "github.com/oceanbase/obkv-table-client-go/protocol"
import "github.com/oceanbase/obkv-table-client-go/util"

const RESERVED_SIZE = 2

type ListMetaValue struct {
	protocol.ObUniVersionHeader
	count_     int64
	left_idx_  int64
	right_idx_ int64
	ttl_       int64
	reserved_  [RESERVED_SIZE]int64
}

func (lmv *ListMetaValue) Decode(buffer *bytes.Buffer) {
	lmv.ObUniVersionHeader.Decode(buffer)
	lmv.count_ = util.DecodeVi64(buffer)
	lmv.left_idx_ = util.DecodeVi64(buffer)
	lmv.right_idx_ = util.DecodeVi64(buffer)
	lmv.ttl_ = util.DecodeVi64(buffer)
}

func (lmv *ListMetaValue) Print() {
	// 打印 ObUniVersionHeader（如果它有String()方法，否则您需要单独处理）
	fmt.Printf("ObUniVersionHeader: %v\n", lmv.ObUniVersionHeader.String())
	fmt.Printf("count: %v, left_idx: %v, right_idx: %v, ttl: %v\n", lmv.count_, lmv.left_idx_, lmv.right_idx_, lmv.ttl_)
}

func TestDecodeListMetaValue(t *testing.T) {
	encodMsg := "019B808080000B808080FBFFFFFFFFFF01000002808480809001E083AAD7F6F71F"
	// 使用hex包中的DecodeString函数把十六进制字符串解码成字节切片
	data, err := hex.DecodeString(encodMsg)
	if err != nil {
		fmt.Println("Error decoding string: ", err)
		return
	}

	// 使用bytes.NewBuffer或bytes.NewBuffer(data)将[]byte转换为bytes.Buffer
	buf := bytes.NewBuffer(data)
	var metaValue ListMetaValue
	metaValue.Decode(buf)
	metaValue.Print()
}
