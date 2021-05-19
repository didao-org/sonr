package bind

import (
	"context"

	"github.com/getsentry/sentry-go"
	"github.com/sonr-io/core/internal/crypto"
	sc "github.com/sonr-io/core/pkg/client"
	md "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

// @ Return URLLink
func GetURLLink(url string) []byte {
	// Create Link
	link := md.NewURLLink(url)

	// Marshal
	bytes, err := proto.Marshal(link)
	if err != nil {
		return nil
	}
	return bytes
}

// @ Gets User from Storj
func GetUser(data []byte) []byte {
	// Unmarshal Request
	request := &md.StorjRequest{}
	proto.Unmarshal(data, request)

	// Get User from Uplink
	user, err := sc.GetUser(context.Background(), request.StorjApiKey, request.GetUserID())
	if err != nil {
		sentry.CaptureException(err)
		return nil
	}

	// Marshal
	bytes, err := proto.Marshal(user)
	if err != nil {
		return nil
	}
	return bytes
}

// @ Puts User into Storj
func PutUser(data []byte) bool {
	// Unmarshal Request
	request := &md.StorjRequest{}
	proto.Unmarshal(data, request)

	// Put User
	err := sc.PutUser(context.Background(), request.StorjApiKey, request.GetUser())
	if err != nil {
		sentry.CaptureException(err)
		return false
	}
	return true

}

// @ Join/Create/Leave Remote Group
func (mn *Node) Remote(data []byte) []byte {
	if mn.isReady() {
		// Generate Word List
		_, wordList, serr := crypto.RandomWords("english", 3)
		if serr != nil {
			mn.handleError(serr)
			return nil
		}
		// Create Remote Request and Join Lobby
		remote := md.NewRemote(wordList)

		// Join Lobby
		tm, serr := mn.client.JoinLobby(remote, true)
		if serr != nil {
			mn.handleError(serr)
			return nil
		}

		// Set Topic
		mn.topics[remote.Topic] = tm

		// Marshal
		data, err := proto.Marshal(remote)
		if err != nil {
			mn.handleError(md.NewError(err, md.ErrorMessage_MARSHAL))
			return nil
		}
		return data
	}
	return nil
}

// @ Update proximity/direction and Notify Lobby
func (mn *Node) Update(data []byte) {
	if mn.isReady() {
		// Initialize from Request
		update := &md.UpdateRequest{}
		if err := proto.Unmarshal(data, update); err != nil {
			mn.handleError(md.NewError(err, md.ErrorMessage_UNMARSHAL))
			return
		}

		// Update Peer
		mn.user.Update(update)

		// Notify Local Lobby
		err := mn.client.Update(mn.local)
		if err != nil {
			mn.handleError(err)
			return
		}
	}
}

// @ Invite Processes Data and Sends Invite to Peer
func (mn *Node) Invite(data []byte) {
	if mn.isReady() {
		// Update Status
		mn.setStatus(md.Status_PENDING)

		// Initialize from Request
		req := &md.AuthInvite{}
		if err := proto.Unmarshal(data, req); err != nil {
			mn.handleError(md.NewError(err, md.ErrorMessage_UNMARSHAL))
			return
		}

		// @ 2. Check Transfer Type
		if req.Payload == md.Payload_CONTACT || req.Payload == md.Payload_FLAT_CONTACT {
			err := mn.client.InviteContact(req, mn.local, mn.user.Contact)
			if err != nil {
				mn.handleError(err)
				return
			}
		} else if req.Payload == md.Payload_URL {
			err := mn.client.InviteLink(req, mn.local)
			if err != nil {
				mn.handleError(err)
				return
			}
		} else {
			// Invite With file
			err := mn.client.InviteFile(req, mn.local)
			if err != nil {
				mn.handleError(err)
				return
			}
		}
	}
}

// @ Respond to an Invite with Decision
func (mn *Node) Respond(data []byte) {
	if mn.isReady() {
		// Initialize from Request
		req := &md.AuthReply{}
		if err := proto.Unmarshal(data, req); err != nil {
			mn.handleError(md.NewError(err, md.ErrorMessage_UNMARSHAL))
			return
		}

		mn.client.Respond(req, mn.local)
		// Update Status
		if req.Decision {
			mn.setStatus(md.Status_INPROGRESS)
		} else {
			mn.setStatus(md.Status_AVAILABLE)
		}
	}
}
