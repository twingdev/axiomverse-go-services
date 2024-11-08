package og_master

import (
	"github.com/twingdev/axiomverse-go/core/consensus"
	"github.com/twingdev/axiomverse-go/core/networking"
	"log"
)

type OGMasterNode struct {
	Consensus *consensus.PoBPC
	NATS      *networking.NATSClient
}

func NewOGMasterNode(natsClient *NATSClient) *OGMasterNode {
	return &OGMasterNode{
		Consensus: consensus.Init(3), // Require 3 validators to pass
		NATS:      natsClient,
	}
}

func (node *OGMasterNode) ProposeData(data []byte) {
	proposalID := node.Consensus.Propose(data)
	log.Printf("OG Master Node proposed data with ID: %s", proposalID)
}
