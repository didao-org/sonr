// Copyright © 2020 AMIS Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package dkg

import (
	"io/ioutil"

	"github.com/golang/protobuf/proto"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/sonr-io/alice/crypto/tss/dkg"
	"github.com/sonr-io/alice/example/utils"
	"github.com/sonr-io/alice/message/types"
)

type service struct {
	config *DKGConfig
	pm     types.PeerManager

	dkg  *dkg.DKG
	done chan struct{}
}

func NewService(config *DKGConfig, pm types.PeerManager) (*service, error) {
	s := &service{
		config: config,
		pm:     pm,
		done:   make(chan struct{}),
	}

	// Create dkg
	d, err := dkg.NewDKG(utils.GetCurve(), pm, config.Threshold, config.Rank, s)
	if err != nil {
		return nil, err
	}
	s.dkg = d
	return s, nil
}

func (p *service) Handle(s network.Stream) {
	data := &dkg.Message{}
	buf, err := ioutil.ReadAll(s)
	if err != nil {
		return
	}
	s.Close()

	// unmarshal it
	err = proto.Unmarshal(buf, data)
	if err != nil {

		return
	}

	err = p.dkg.AddMessage(data)
	if err != nil {

		return
	}
}

func (p *service) Process() {
	// 1. Start a DKG process.
	p.dkg.Start()
	defer p.dkg.Stop()

	// 2. Wait the dkg is done or failed
	<-p.done
}

func (p *service) OnStateChanged(oldState types.MainState, newState types.MainState) {
	if newState == types.StateFailed {
		close(p.done)
		return
	} else if newState == types.StateDone {
		result, err := p.dkg.GetResult()
		if err == nil {
			writeDKGResult(p.pm.SelfID(), result)
		} else {
		}
		close(p.done)
		return
	}
}