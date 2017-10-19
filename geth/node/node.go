package node

import (
	"sync"

	"github.com/status-im/status-go/geth/common/geth"
)

type node struct {
	geth.Node
	*sync.RWMutex
}

func newNode() *node {
	m := &sync.RWMutex{}
	return &node{RWMutex: m}
}
