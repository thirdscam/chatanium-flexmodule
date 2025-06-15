package module

import (
	"context"

	"github.com/bwmarrin/discordgo"
	plugin "github.com/hashicorp/go-plugin"
	proto_common "github.com/thirdscam/chatanium-flexmodule/proto"
	proto "github.com/thirdscam/chatanium-flexmodule/proto/discord-v1"
	shared "github.com/thirdscam/chatanium-flexmodule/shared/discord-v1"
	"github.com/thirdscam/chatanium-flexmodule/shared/discord-v1/convert/buf2struct"
	"github.com/thirdscam/chatanium-flexmodule/shared/discord-v1/convert/struct2buf"
)

// `module/server.go` implements the gRPC server for receiving from the runtime.
//
// This part works on the module-side and is the gRPC server implementation for the runtime.
type GRPCServer struct {
	proto.UnimplementedHookServer
	Impl   shared.Hook // Hook functions to be called from runtime (module developers must implement this!)
	broker *plugin.GRPCBroker
	helper shared.Helper // Helper service provided by runtime
}

// OnInit is called when the discord plugin is initialized.
func (m *GRPCServer) OnInit(ctx context.Context, req *proto_common.Empty) (*proto.InitResponse, error) {
	// Pass helper service to hook implementation
	resp := m.Impl.OnInit(m.helper)

	interactions := make([]*proto.ApplicationCommand, 0)
	for _, v := range resp.Interactions {
		interactions = append(interactions, struct2buf.ApplicationCommmand(v))
	}

	return &proto.InitResponse{
		Interactions: interactions,
	}, nil
}

// OnCreateChatMessage is called when a message is created from the runtime.
func (m *GRPCServer) OnCreateMessage(ctx context.Context, req *proto.Message) (*proto_common.Empty, error) {
	// Convert the protobuf message to a discordgo.Message struct
	m.Impl.OnCreateChatMessage(buf2struct.Message(req))

	// Hook function is not required to return anything to the client (runtime)
	return &proto_common.Empty{}, nil
}

// OnCreateInteraction is called when an interaction is created from the runtime.
func (m *GRPCServer) OnCreateInteraction(ctx context.Context, req *proto.Interaction) (*proto_common.Empty, error) {
	// Convert the protobuf message to a discordgo.Interaction struct
	m.Impl.OnCreateInteraction(buf2struct.Interaction(req))

	// Hook function is not required to return anything to the client (runtime)
	return &proto_common.Empty{}, nil
}

// OnEvent is called when an (discord) event is created from the runtime.
func (m *GRPCServer) OnEvent(ctx context.Context, req *proto.OnEventRequest) (*proto_common.Empty, error) {
	m.Impl.OnEvent(req.Event)

	// Hook function is not required to return anything to the client (runtime)
	return &proto_common.Empty{}, nil
}

// HookClient implements the Hook interface for calling runtime's hook functions from module
type HookClient struct {
	client proto.HookClient
	broker *plugin.GRPCBroker
}

// OnInit calls the runtime's OnInit hook function
func (h *HookClient) OnInit(helper shared.Helper) shared.InitResponse {
	resp, err := h.client.OnInit(context.Background(), &proto_common.Empty{})
	if err != nil {
		// Return empty response on error
		return shared.InitResponse{
			Interactions: []*discordgo.ApplicationCommand{},
		}
	}

	interactions := make([]*discordgo.ApplicationCommand, 0, len(resp.Interactions))
	for _, interaction := range resp.Interactions {
		interactions = append(interactions, buf2struct.ApplicationCommand(interaction))
	}

	return shared.InitResponse{
		Interactions: interactions,
	}
}

// OnCreateChatMessage calls the runtime's OnCreateChatMessage hook function
func (h *HookClient) OnCreateChatMessage(message *discordgo.Message) error {
	_, err := h.client.OnCreateMessage(context.Background(), struct2buf.Message(message))
	return err
}

// OnCreateInteraction calls the runtime's OnCreateInteraction hook function
func (h *HookClient) OnCreateInteraction(interaction *discordgo.Interaction) error {
	_, err := h.client.OnCreateInteraction(context.Background(), struct2buf.Interaction(interaction))
	return err
}

// OnEvent calls the runtime's OnEvent hook function
func (h *HookClient) OnEvent(eventType string) error {
	_, err := h.client.OnEvent(context.Background(), &proto.OnEventRequest{Event: eventType})
	return err
}

var _ shared.Hook = &HookClient{}
