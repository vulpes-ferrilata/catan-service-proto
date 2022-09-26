// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: protobuf/responses/game.proto

package responses

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Game struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID               string             `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Status           string             `protobuf:"bytes,2,opt,name=Status,proto3" json:"Status,omitempty"`
	Phase            string             `protobuf:"bytes,3,opt,name=Phase,proto3" json:"Phase,omitempty"`
	Turn             int32              `protobuf:"varint,4,opt,name=Turn,proto3" json:"Turn,omitempty"`
	ActivePlayer     *Player            `protobuf:"bytes,5,opt,name=ActivePlayer,proto3" json:"ActivePlayer,omitempty"`
	Players          []*Player          `protobuf:"bytes,6,rep,name=Players,proto3" json:"Players,omitempty"`
	Dices            []*Dice            `protobuf:"bytes,7,rep,name=Dices,proto3" json:"Dices,omitempty"`
	Achievements     []*Achievement     `protobuf:"bytes,8,rep,name=Achievements,proto3" json:"Achievements,omitempty"`
	ResourceCards    []*ResourceCard    `protobuf:"bytes,9,rep,name=ResourceCards,proto3" json:"ResourceCards,omitempty"`
	DevelopmentCards []*DevelopmentCard `protobuf:"bytes,10,rep,name=DevelopmentCards,proto3" json:"DevelopmentCards,omitempty"`
	Terrains         []*Terrain         `protobuf:"bytes,11,rep,name=Terrains,proto3" json:"Terrains,omitempty"`
	Lands            []*Land            `protobuf:"bytes,12,rep,name=Lands,proto3" json:"Lands,omitempty"`
	Paths            []*Path            `protobuf:"bytes,13,rep,name=Paths,proto3" json:"Paths,omitempty"`
}

func (x *Game) Reset() {
	*x = Game{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_responses_game_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Game) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Game) ProtoMessage() {}

func (x *Game) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_responses_game_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Game.ProtoReflect.Descriptor instead.
func (*Game) Descriptor() ([]byte, []int) {
	return file_protobuf_responses_game_proto_rawDescGZIP(), []int{0}
}

func (x *Game) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Game) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Game) GetPhase() string {
	if x != nil {
		return x.Phase
	}
	return ""
}

func (x *Game) GetTurn() int32 {
	if x != nil {
		return x.Turn
	}
	return 0
}

func (x *Game) GetActivePlayer() *Player {
	if x != nil {
		return x.ActivePlayer
	}
	return nil
}

func (x *Game) GetPlayers() []*Player {
	if x != nil {
		return x.Players
	}
	return nil
}

func (x *Game) GetDices() []*Dice {
	if x != nil {
		return x.Dices
	}
	return nil
}

func (x *Game) GetAchievements() []*Achievement {
	if x != nil {
		return x.Achievements
	}
	return nil
}

func (x *Game) GetResourceCards() []*ResourceCard {
	if x != nil {
		return x.ResourceCards
	}
	return nil
}

func (x *Game) GetDevelopmentCards() []*DevelopmentCard {
	if x != nil {
		return x.DevelopmentCards
	}
	return nil
}

func (x *Game) GetTerrains() []*Terrain {
	if x != nil {
		return x.Terrains
	}
	return nil
}

func (x *Game) GetLands() []*Land {
	if x != nil {
		return x.Lands
	}
	return nil
}

func (x *Game) GetPaths() []*Path {
	if x != nil {
		return x.Paths
	}
	return nil
}

var File_protobuf_responses_game_proto protoreflect.FileDescriptor

var file_protobuf_responses_game_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x73, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x09, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x1a, 0x1f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x2f, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x2f,
	0x64, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x2f, 0x61,
	0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x26, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x63, 0x61,
	0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x2f, 0x64, 0x65, 0x76,
	0x65, 0x6c, 0x6f, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x2f, 0x74, 0x65, 0x72, 0x72, 0x61, 0x69, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x2f, 0x6c, 0x61, 0x6e, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x2f, 0x70, 0x61, 0x74, 0x68, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xa4, 0x04, 0x0a, 0x04, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x68, 0x61, 0x73, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x50, 0x68, 0x61, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x75,
	0x72, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x54, 0x75, 0x72, 0x6e, 0x12, 0x35,
	0x0a, 0x0c, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73,
	0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x0c, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x07, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x73, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x07, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x73, 0x12, 0x25, 0x0a, 0x05, 0x44, 0x69, 0x63, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x2e, 0x44, 0x69,
	0x63, 0x65, 0x52, 0x05, 0x44, 0x69, 0x63, 0x65, 0x73, 0x12, 0x3a, 0x0a, 0x0c, 0x41, 0x63, 0x68,
	0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x2e, 0x41, 0x63, 0x68, 0x69,
	0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0c, 0x41, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x3d, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x43, 0x61, 0x72, 0x64, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x43, 0x61, 0x72, 0x64, 0x52, 0x0d, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43,
	0x61, 0x72, 0x64, 0x73, 0x12, 0x46, 0x0a, 0x10, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x6d,
	0x65, 0x6e, 0x74, 0x43, 0x61, 0x72, 0x64, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x76, 0x65, 0x6c,
	0x6f, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x72, 0x64, 0x52, 0x10, 0x44, 0x65, 0x76, 0x65,
	0x6c, 0x6f, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x72, 0x64, 0x73, 0x12, 0x2e, 0x0a, 0x08,
	0x54, 0x65, 0x72, 0x72, 0x61, 0x69, 0x6e, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x2e, 0x54, 0x65, 0x72, 0x72, 0x61,
	0x69, 0x6e, 0x52, 0x08, 0x54, 0x65, 0x72, 0x72, 0x61, 0x69, 0x6e, 0x73, 0x12, 0x25, 0x0a, 0x05,
	0x4c, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x2e, 0x4c, 0x61, 0x6e, 0x64, 0x52, 0x05, 0x4c, 0x61,
	0x6e, 0x64, 0x73, 0x12, 0x25, 0x0a, 0x05, 0x50, 0x61, 0x74, 0x68, 0x73, 0x18, 0x0d, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x2e, 0x50,
	0x61, 0x74, 0x68, 0x52, 0x05, 0x50, 0x61, 0x74, 0x68, 0x73, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x75, 0x6c, 0x70, 0x65, 0x73, 0x2d,
	0x66, 0x65, 0x72, 0x72, 0x69, 0x6c, 0x61, 0x74, 0x61, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6e, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62,
	0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_protobuf_responses_game_proto_rawDescOnce sync.Once
	file_protobuf_responses_game_proto_rawDescData = file_protobuf_responses_game_proto_rawDesc
)

func file_protobuf_responses_game_proto_rawDescGZIP() []byte {
	file_protobuf_responses_game_proto_rawDescOnce.Do(func() {
		file_protobuf_responses_game_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_responses_game_proto_rawDescData)
	})
	return file_protobuf_responses_game_proto_rawDescData
}

var file_protobuf_responses_game_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_protobuf_responses_game_proto_goTypes = []interface{}{
	(*Game)(nil),            // 0: responses.Game
	(*Player)(nil),          // 1: responses.Player
	(*Dice)(nil),            // 2: responses.Dice
	(*Achievement)(nil),     // 3: responses.Achievement
	(*ResourceCard)(nil),    // 4: responses.ResourceCard
	(*DevelopmentCard)(nil), // 5: responses.DevelopmentCard
	(*Terrain)(nil),         // 6: responses.Terrain
	(*Land)(nil),            // 7: responses.Land
	(*Path)(nil),            // 8: responses.Path
}
var file_protobuf_responses_game_proto_depIdxs = []int32{
	1, // 0: responses.Game.ActivePlayer:type_name -> responses.Player
	1, // 1: responses.Game.Players:type_name -> responses.Player
	2, // 2: responses.Game.Dices:type_name -> responses.Dice
	3, // 3: responses.Game.Achievements:type_name -> responses.Achievement
	4, // 4: responses.Game.ResourceCards:type_name -> responses.ResourceCard
	5, // 5: responses.Game.DevelopmentCards:type_name -> responses.DevelopmentCard
	6, // 6: responses.Game.Terrains:type_name -> responses.Terrain
	7, // 7: responses.Game.Lands:type_name -> responses.Land
	8, // 8: responses.Game.Paths:type_name -> responses.Path
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_protobuf_responses_game_proto_init() }
func file_protobuf_responses_game_proto_init() {
	if File_protobuf_responses_game_proto != nil {
		return
	}
	file_protobuf_responses_player_proto_init()
	file_protobuf_responses_dice_proto_init()
	file_protobuf_responses_achievement_proto_init()
	file_protobuf_responses_resource_card_proto_init()
	file_protobuf_responses_development_card_proto_init()
	file_protobuf_responses_terrain_proto_init()
	file_protobuf_responses_land_proto_init()
	file_protobuf_responses_path_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_protobuf_responses_game_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Game); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protobuf_responses_game_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protobuf_responses_game_proto_goTypes,
		DependencyIndexes: file_protobuf_responses_game_proto_depIdxs,
		MessageInfos:      file_protobuf_responses_game_proto_msgTypes,
	}.Build()
	File_protobuf_responses_game_proto = out.File
	file_protobuf_responses_game_proto_rawDesc = nil
	file_protobuf_responses_game_proto_goTypes = nil
	file_protobuf_responses_game_proto_depIdxs = nil
}
