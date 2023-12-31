// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.4
// source: album.proto

package album

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

type Album struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AlbumId     string          `protobuf:"bytes,1,opt,name=albumId,proto3" json:"albumId,omitempty"`
	SpotifyUri  string          `protobuf:"bytes,2,opt,name=spotifyUri,proto3" json:"spotifyUri,omitempty"`
	SpotifyUrl  string          `protobuf:"bytes,3,opt,name=spotifyUrl,proto3" json:"spotifyUrl,omitempty"`
	Name        string          `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Artists     []*SimpleArtist `protobuf:"bytes,5,rep,name=artists,proto3" json:"artists,omitempty"`
	Tracks      *TrackPage      `protobuf:"bytes,6,opt,name=tracks,proto3" json:"tracks,omitempty"`
	CoverUrl    string          `protobuf:"bytes,7,opt,name=coverUrl,proto3" json:"coverUrl,omitempty"`
	ReleaseDate string          `protobuf:"bytes,8,opt,name=releaseDate,proto3" json:"releaseDate,omitempty"`
	Genres      []string        `protobuf:"bytes,9,rep,name=genres,proto3" json:"genres,omitempty"`
}

func (x *Album) Reset() {
	*x = Album{}
	if protoimpl.UnsafeEnabled {
		mi := &file_album_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Album) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Album) ProtoMessage() {}

func (x *Album) ProtoReflect() protoreflect.Message {
	mi := &file_album_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Album.ProtoReflect.Descriptor instead.
func (*Album) Descriptor() ([]byte, []int) {
	return file_album_proto_rawDescGZIP(), []int{0}
}

func (x *Album) GetAlbumId() string {
	if x != nil {
		return x.AlbumId
	}
	return ""
}

func (x *Album) GetSpotifyUri() string {
	if x != nil {
		return x.SpotifyUri
	}
	return ""
}

func (x *Album) GetSpotifyUrl() string {
	if x != nil {
		return x.SpotifyUrl
	}
	return ""
}

func (x *Album) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Album) GetArtists() []*SimpleArtist {
	if x != nil {
		return x.Artists
	}
	return nil
}

func (x *Album) GetTracks() *TrackPage {
	if x != nil {
		return x.Tracks
	}
	return nil
}

func (x *Album) GetCoverUrl() string {
	if x != nil {
		return x.CoverUrl
	}
	return ""
}

func (x *Album) GetReleaseDate() string {
	if x != nil {
		return x.ReleaseDate
	}
	return ""
}

func (x *Album) GetGenres() []string {
	if x != nil {
		return x.Genres
	}
	return nil
}

type SimpleAlbum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AlbumId    string          `protobuf:"bytes,1,opt,name=albumId,proto3" json:"albumId,omitempty"`
	SpotifyUri string          `protobuf:"bytes,2,opt,name=spotifyUri,proto3" json:"spotifyUri,omitempty"`
	SpotifyUrl string          `protobuf:"bytes,3,opt,name=spotifyUrl,proto3" json:"spotifyUrl,omitempty"`
	Name       string          `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Artists    []*SimpleArtist `protobuf:"bytes,5,rep,name=artists,proto3" json:"artists,omitempty"`
	CoverUrl   string          `protobuf:"bytes,6,opt,name=coverUrl,proto3" json:"coverUrl,omitempty"`
	Genres     []string        `protobuf:"bytes,7,rep,name=genres,proto3" json:"genres,omitempty"`
}

func (x *SimpleAlbum) Reset() {
	*x = SimpleAlbum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_album_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimpleAlbum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleAlbum) ProtoMessage() {}

func (x *SimpleAlbum) ProtoReflect() protoreflect.Message {
	mi := &file_album_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleAlbum.ProtoReflect.Descriptor instead.
func (*SimpleAlbum) Descriptor() ([]byte, []int) {
	return file_album_proto_rawDescGZIP(), []int{1}
}

func (x *SimpleAlbum) GetAlbumId() string {
	if x != nil {
		return x.AlbumId
	}
	return ""
}

func (x *SimpleAlbum) GetSpotifyUri() string {
	if x != nil {
		return x.SpotifyUri
	}
	return ""
}

func (x *SimpleAlbum) GetSpotifyUrl() string {
	if x != nil {
		return x.SpotifyUrl
	}
	return ""
}

func (x *SimpleAlbum) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SimpleAlbum) GetArtists() []*SimpleArtist {
	if x != nil {
		return x.Artists
	}
	return nil
}

func (x *SimpleAlbum) GetCoverUrl() string {
	if x != nil {
		return x.CoverUrl
	}
	return ""
}

func (x *SimpleAlbum) GetGenres() []string {
	if x != nil {
		return x.Genres
	}
	return nil
}

type SimpleArtist struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ArtistId   string `protobuf:"bytes,1,opt,name=artistId,proto3" json:"artistId,omitempty"`
	SpotifyUri string `protobuf:"bytes,3,opt,name=spotifyUri,proto3" json:"spotifyUri,omitempty"`
	SpotifyUrl string `protobuf:"bytes,4,opt,name=spotifyUrl,proto3" json:"spotifyUrl,omitempty"`
	Name       string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SimpleArtist) Reset() {
	*x = SimpleArtist{}
	if protoimpl.UnsafeEnabled {
		mi := &file_album_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimpleArtist) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleArtist) ProtoMessage() {}

func (x *SimpleArtist) ProtoReflect() protoreflect.Message {
	mi := &file_album_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleArtist.ProtoReflect.Descriptor instead.
func (*SimpleArtist) Descriptor() ([]byte, []int) {
	return file_album_proto_rawDescGZIP(), []int{2}
}

func (x *SimpleArtist) GetArtistId() string {
	if x != nil {
		return x.ArtistId
	}
	return ""
}

func (x *SimpleArtist) GetSpotifyUri() string {
	if x != nil {
		return x.SpotifyUri
	}
	return ""
}

func (x *SimpleArtist) GetSpotifyUrl() string {
	if x != nil {
		return x.SpotifyUrl
	}
	return ""
}

func (x *SimpleArtist) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Track struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TrackId     string `protobuf:"bytes,1,opt,name=trackId,proto3" json:"trackId,omitempty"`
	SpotifyUri  string `protobuf:"bytes,3,opt,name=spotifyUri,proto3" json:"spotifyUri,omitempty"`
	SpotifyUrl  string `protobuf:"bytes,4,opt,name=spotifyUrl,proto3" json:"spotifyUrl,omitempty"`
	Title       string `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
	DurationMs  int32  `protobuf:"varint,6,opt,name=durationMs,proto3" json:"durationMs,omitempty"`
	TrackNumber int32  `protobuf:"varint,7,opt,name=trackNumber,proto3" json:"trackNumber,omitempty"`
	PreviewUrl  string `protobuf:"bytes,8,opt,name=previewUrl,proto3" json:"previewUrl,omitempty"`
}

func (x *Track) Reset() {
	*x = Track{}
	if protoimpl.UnsafeEnabled {
		mi := &file_album_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Track) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Track) ProtoMessage() {}

func (x *Track) ProtoReflect() protoreflect.Message {
	mi := &file_album_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Track.ProtoReflect.Descriptor instead.
func (*Track) Descriptor() ([]byte, []int) {
	return file_album_proto_rawDescGZIP(), []int{3}
}

func (x *Track) GetTrackId() string {
	if x != nil {
		return x.TrackId
	}
	return ""
}

func (x *Track) GetSpotifyUri() string {
	if x != nil {
		return x.SpotifyUri
	}
	return ""
}

func (x *Track) GetSpotifyUrl() string {
	if x != nil {
		return x.SpotifyUrl
	}
	return ""
}

func (x *Track) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Track) GetDurationMs() int32 {
	if x != nil {
		return x.DurationMs
	}
	return 0
}

func (x *Track) GetTrackNumber() int32 {
	if x != nil {
		return x.TrackNumber
	}
	return 0
}

func (x *Track) GetPreviewUrl() string {
	if x != nil {
		return x.PreviewUrl
	}
	return ""
}

type TrackPage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tracks   []*Track `protobuf:"bytes,1,rep,name=tracks,proto3" json:"tracks,omitempty"`
	Limit    int32    `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset   int32    `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	Total    int32    `protobuf:"varint,4,opt,name=total,proto3" json:"total,omitempty"`
	Next     string   `protobuf:"bytes,5,opt,name=next,proto3" json:"next,omitempty"`
	Previous string   `protobuf:"bytes,6,opt,name=previous,proto3" json:"previous,omitempty"`
}

func (x *TrackPage) Reset() {
	*x = TrackPage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_album_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrackPage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrackPage) ProtoMessage() {}

func (x *TrackPage) ProtoReflect() protoreflect.Message {
	mi := &file_album_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrackPage.ProtoReflect.Descriptor instead.
func (*TrackPage) Descriptor() ([]byte, []int) {
	return file_album_proto_rawDescGZIP(), []int{4}
}

func (x *TrackPage) GetTracks() []*Track {
	if x != nil {
		return x.Tracks
	}
	return nil
}

func (x *TrackPage) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *TrackPage) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *TrackPage) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *TrackPage) GetNext() string {
	if x != nil {
		return x.Next
	}
	return ""
}

func (x *TrackPage) GetPrevious() string {
	if x != nil {
		return x.Previous
	}
	return ""
}

var File_album_proto protoreflect.FileDescriptor

var file_album_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61,
	0x6c, 0x62, 0x75, 0x6d, 0x22, 0xa4, 0x02, 0x0a, 0x05, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x70, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x55, 0x72, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x70,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x55, 0x72, 0x69, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x70, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x55, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x70,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x55, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x07,
	0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x61, 0x6c, 0x62, 0x75, 0x6d, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x41, 0x72, 0x74, 0x69,
	0x73, 0x74, 0x52, 0x07, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x73, 0x12, 0x28, 0x0a, 0x06, 0x74,
	0x72, 0x61, 0x63, 0x6b, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61, 0x6c,
	0x62, 0x75, 0x6d, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x50, 0x61, 0x67, 0x65, 0x52, 0x06, 0x74,
	0x72, 0x61, 0x63, 0x6b, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x55, 0x72,
	0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x55, 0x72,
	0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x44, 0x61, 0x74, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x18, 0x09, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x22, 0xde, 0x01, 0x0a, 0x0b,
	0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x6c, 0x62, 0x75, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x6c,
	0x62, 0x75, 0x6d, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x70, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x55, 0x72, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x70, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x55, 0x72, 0x69, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x70, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x55, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x70, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x55, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x61, 0x72, 0x74,
	0x69, 0x73, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x6c, 0x62,
	0x75, 0x6d, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x52,
	0x07, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x55, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x55, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x18, 0x07,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x22, 0x7e, 0x0a, 0x0c,
	0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x70, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x55, 0x72, 0x69, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x70,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x55, 0x72, 0x69, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x70, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x55, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x70,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x55, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xd9, 0x01, 0x0a,
	0x05, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x49, 0x64,
	0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x70, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x55, 0x72, 0x69, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x70, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x55, 0x72, 0x69,
	0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x70, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x55, 0x72, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x70, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x55, 0x72, 0x6c,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4d, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x74, 0x72, 0x61,
	0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x55, 0x72, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x55, 0x72, 0x6c, 0x22, 0xa5, 0x01, 0x0a, 0x09, 0x54, 0x72, 0x61,
	0x63, 0x6b, 0x50, 0x61, 0x67, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x2e, 0x54,
	0x72, 0x61, 0x63, 0x6b, 0x52, 0x06, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x65, 0x78, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x65, 0x78, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73,
	0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b,
	0x6e, 0x67, 0x6e, 0x6b, 0x67, 0x2f, 0x74, 0x75, 0x6e, 0x65, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x2f,
	0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x61, 0x6c, 0x62, 0x75,
	0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_album_proto_rawDescOnce sync.Once
	file_album_proto_rawDescData = file_album_proto_rawDesc
)

func file_album_proto_rawDescGZIP() []byte {
	file_album_proto_rawDescOnce.Do(func() {
		file_album_proto_rawDescData = protoimpl.X.CompressGZIP(file_album_proto_rawDescData)
	})
	return file_album_proto_rawDescData
}

var file_album_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_album_proto_goTypes = []interface{}{
	(*Album)(nil),        // 0: album.Album
	(*SimpleAlbum)(nil),  // 1: album.SimpleAlbum
	(*SimpleArtist)(nil), // 2: album.SimpleArtist
	(*Track)(nil),        // 3: album.Track
	(*TrackPage)(nil),    // 4: album.TrackPage
}
var file_album_proto_depIdxs = []int32{
	2, // 0: album.Album.artists:type_name -> album.SimpleArtist
	4, // 1: album.Album.tracks:type_name -> album.TrackPage
	2, // 2: album.SimpleAlbum.artists:type_name -> album.SimpleArtist
	3, // 3: album.TrackPage.tracks:type_name -> album.Track
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_album_proto_init() }
func file_album_proto_init() {
	if File_album_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_album_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Album); i {
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
		file_album_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimpleAlbum); i {
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
		file_album_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimpleArtist); i {
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
		file_album_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Track); i {
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
		file_album_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrackPage); i {
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
			RawDescriptor: file_album_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_album_proto_goTypes,
		DependencyIndexes: file_album_proto_depIdxs,
		MessageInfos:      file_album_proto_msgTypes,
	}.Build()
	File_album_proto = out.File
	file_album_proto_rawDesc = nil
	file_album_proto_goTypes = nil
	file_album_proto_depIdxs = nil
}
