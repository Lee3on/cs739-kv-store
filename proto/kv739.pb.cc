// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: kv739.proto

#include "kv739.pb.h"

#include <algorithm>
#include "google/protobuf/io/coded_stream.h"
#include "google/protobuf/extension_set.h"
#include "google/protobuf/wire_format_lite.h"
#include "google/protobuf/descriptor.h"
#include "google/protobuf/generated_message_reflection.h"
#include "google/protobuf/reflection_ops.h"
#include "google/protobuf/wire_format.h"
// @@protoc_insertion_point(includes)

// Must be included last.
#include "google/protobuf/port_def.inc"
PROTOBUF_PRAGMA_INIT_SEG
namespace _pb = ::PROTOBUF_NAMESPACE_ID;
namespace _pbi = ::PROTOBUF_NAMESPACE_ID::internal;
namespace kv739 {
template <typename>
PROTOBUF_CONSTEXPR GetRequest::GetRequest(
    ::_pbi::ConstantInitialized): _impl_{
    /*decltype(_impl_.key_)*/ {
    &::_pbi::fixed_address_empty_string, ::_pbi::ConstantInitialized {}
  }

  , /*decltype(_impl_._cached_size_)*/{}} {}
struct GetRequestDefaultTypeInternal {
  PROTOBUF_CONSTEXPR GetRequestDefaultTypeInternal() : _instance(::_pbi::ConstantInitialized{}) {}
  ~GetRequestDefaultTypeInternal() {}
  union {
    GetRequest _instance;
  };
};

PROTOBUF_ATTRIBUTE_NO_DESTROY PROTOBUF_CONSTINIT
    PROTOBUF_ATTRIBUTE_INIT_PRIORITY1 GetRequestDefaultTypeInternal _GetRequest_default_instance_;
template <typename>
PROTOBUF_CONSTEXPR GetResponse::GetResponse(
    ::_pbi::ConstantInitialized): _impl_{
    /*decltype(_impl_.value_)*/ {
    &::_pbi::fixed_address_empty_string, ::_pbi::ConstantInitialized {}
  }

  , /*decltype(_impl_.status_)*/ 0

  , /*decltype(_impl_._cached_size_)*/{}} {}
struct GetResponseDefaultTypeInternal {
  PROTOBUF_CONSTEXPR GetResponseDefaultTypeInternal() : _instance(::_pbi::ConstantInitialized{}) {}
  ~GetResponseDefaultTypeInternal() {}
  union {
    GetResponse _instance;
  };
};

PROTOBUF_ATTRIBUTE_NO_DESTROY PROTOBUF_CONSTINIT
    PROTOBUF_ATTRIBUTE_INIT_PRIORITY1 GetResponseDefaultTypeInternal _GetResponse_default_instance_;
template <typename>
PROTOBUF_CONSTEXPR PutRequest::PutRequest(
    ::_pbi::ConstantInitialized): _impl_{
    /*decltype(_impl_.key_)*/ {
    &::_pbi::fixed_address_empty_string, ::_pbi::ConstantInitialized {}
  }

  , /*decltype(_impl_.value_)*/ {
    &::_pbi::fixed_address_empty_string, ::_pbi::ConstantInitialized {}
  }

  , /*decltype(_impl_._cached_size_)*/{}} {}
struct PutRequestDefaultTypeInternal {
  PROTOBUF_CONSTEXPR PutRequestDefaultTypeInternal() : _instance(::_pbi::ConstantInitialized{}) {}
  ~PutRequestDefaultTypeInternal() {}
  union {
    PutRequest _instance;
  };
};

PROTOBUF_ATTRIBUTE_NO_DESTROY PROTOBUF_CONSTINIT
    PROTOBUF_ATTRIBUTE_INIT_PRIORITY1 PutRequestDefaultTypeInternal _PutRequest_default_instance_;
template <typename>
PROTOBUF_CONSTEXPR PutResponse::PutResponse(
    ::_pbi::ConstantInitialized): _impl_{
    /*decltype(_impl_.old_value_)*/ {
    &::_pbi::fixed_address_empty_string, ::_pbi::ConstantInitialized {}
  }

  , /*decltype(_impl_.status_)*/ 0

  , /*decltype(_impl_._cached_size_)*/{}} {}
struct PutResponseDefaultTypeInternal {
  PROTOBUF_CONSTEXPR PutResponseDefaultTypeInternal() : _instance(::_pbi::ConstantInitialized{}) {}
  ~PutResponseDefaultTypeInternal() {}
  union {
    PutResponse _instance;
  };
};

PROTOBUF_ATTRIBUTE_NO_DESTROY PROTOBUF_CONSTINIT
    PROTOBUF_ATTRIBUTE_INIT_PRIORITY1 PutResponseDefaultTypeInternal _PutResponse_default_instance_;
}  // namespace kv739
static ::_pb::Metadata file_level_metadata_kv739_2eproto[4];
static constexpr const ::_pb::EnumDescriptor**
    file_level_enum_descriptors_kv739_2eproto = nullptr;
static constexpr const ::_pb::ServiceDescriptor**
    file_level_service_descriptors_kv739_2eproto = nullptr;
const ::uint32_t TableStruct_kv739_2eproto::offsets[] PROTOBUF_SECTION_VARIABLE(
    protodesc_cold) = {
    ~0u,  // no _has_bits_
    PROTOBUF_FIELD_OFFSET(::kv739::GetRequest, _internal_metadata_),
    ~0u,  // no _extensions_
    ~0u,  // no _oneof_case_
    ~0u,  // no _weak_field_map_
    ~0u,  // no _inlined_string_donated_
    ~0u,  // no _split_
    ~0u,  // no sizeof(Split)
    PROTOBUF_FIELD_OFFSET(::kv739::GetRequest, _impl_.key_),
    ~0u,  // no _has_bits_
    PROTOBUF_FIELD_OFFSET(::kv739::GetResponse, _internal_metadata_),
    ~0u,  // no _extensions_
    ~0u,  // no _oneof_case_
    ~0u,  // no _weak_field_map_
    ~0u,  // no _inlined_string_donated_
    ~0u,  // no _split_
    ~0u,  // no sizeof(Split)
    PROTOBUF_FIELD_OFFSET(::kv739::GetResponse, _impl_.status_),
    PROTOBUF_FIELD_OFFSET(::kv739::GetResponse, _impl_.value_),
    ~0u,  // no _has_bits_
    PROTOBUF_FIELD_OFFSET(::kv739::PutRequest, _internal_metadata_),
    ~0u,  // no _extensions_
    ~0u,  // no _oneof_case_
    ~0u,  // no _weak_field_map_
    ~0u,  // no _inlined_string_donated_
    ~0u,  // no _split_
    ~0u,  // no sizeof(Split)
    PROTOBUF_FIELD_OFFSET(::kv739::PutRequest, _impl_.key_),
    PROTOBUF_FIELD_OFFSET(::kv739::PutRequest, _impl_.value_),
    ~0u,  // no _has_bits_
    PROTOBUF_FIELD_OFFSET(::kv739::PutResponse, _internal_metadata_),
    ~0u,  // no _extensions_
    ~0u,  // no _oneof_case_
    ~0u,  // no _weak_field_map_
    ~0u,  // no _inlined_string_donated_
    ~0u,  // no _split_
    ~0u,  // no sizeof(Split)
    PROTOBUF_FIELD_OFFSET(::kv739::PutResponse, _impl_.status_),
    PROTOBUF_FIELD_OFFSET(::kv739::PutResponse, _impl_.old_value_),
};

static const ::_pbi::MigrationSchema
    schemas[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
        { 0, -1, -1, sizeof(::kv739::GetRequest)},
        { 9, -1, -1, sizeof(::kv739::GetResponse)},
        { 19, -1, -1, sizeof(::kv739::PutRequest)},
        { 29, -1, -1, sizeof(::kv739::PutResponse)},
};

static const ::_pb::Message* const file_default_instances[] = {
    &::kv739::_GetRequest_default_instance_._instance,
    &::kv739::_GetResponse_default_instance_._instance,
    &::kv739::_PutRequest_default_instance_._instance,
    &::kv739::_PutResponse_default_instance_._instance,
};
const char descriptor_table_protodef_kv739_2eproto[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
    "\n\013kv739.proto\022\005kv739\"\031\n\nGetRequest\022\013\n\003ke"
    "y\030\001 \001(\t\",\n\013GetResponse\022\016\n\006status\030\001 \001(\005\022\r"
    "\n\005value\030\002 \001(\t\"(\n\nPutRequest\022\013\n\003key\030\001 \001(\t"
    "\022\r\n\005value\030\002 \001(\t\"0\n\013PutResponse\022\016\n\006status"
    "\030\001 \001(\005\022\021\n\told_value\030\002 \001(\t2l\n\016KVStoreServ"
    "ice\022,\n\003Get\022\021.kv739.GetRequest\032\022.kv739.Ge"
    "tResponse\022,\n\003Put\022\021.kv739.PutRequest\032\022.kv"
    "739.PutResponseB\"Z cs739-kv-store/proto/"
    "kv739;kv739b\006proto3"
};
static ::absl::once_flag descriptor_table_kv739_2eproto_once;
const ::_pbi::DescriptorTable descriptor_table_kv739_2eproto = {
    false,
    false,
    339,
    descriptor_table_protodef_kv739_2eproto,
    "kv739.proto",
    &descriptor_table_kv739_2eproto_once,
    nullptr,
    0,
    4,
    schemas,
    file_default_instances,
    TableStruct_kv739_2eproto::offsets,
    file_level_metadata_kv739_2eproto,
    file_level_enum_descriptors_kv739_2eproto,
    file_level_service_descriptors_kv739_2eproto,
};

// This function exists to be marked as weak.
// It can significantly speed up compilation by breaking up LLVM's SCC
// in the .pb.cc translation units. Large translation units see a
// reduction of more than 35% of walltime for optimized builds. Without
// the weak attribute all the messages in the file, including all the
// vtables and everything they use become part of the same SCC through
// a cycle like:
// GetMetadata -> descriptor table -> default instances ->
//   vtables -> GetMetadata
// By adding a weak function here we break the connection from the
// individual vtables back into the descriptor table.
PROTOBUF_ATTRIBUTE_WEAK const ::_pbi::DescriptorTable* descriptor_table_kv739_2eproto_getter() {
  return &descriptor_table_kv739_2eproto;
}
// Force running AddDescriptors() at dynamic initialization time.
PROTOBUF_ATTRIBUTE_INIT_PRIORITY2
static ::_pbi::AddDescriptorsRunner dynamic_init_dummy_kv739_2eproto(&descriptor_table_kv739_2eproto);
namespace kv739 {
// ===================================================================

class GetRequest::_Internal {
 public:
};

GetRequest::GetRequest(::PROTOBUF_NAMESPACE_ID::Arena* arena)
  : ::PROTOBUF_NAMESPACE_ID::Message(arena) {
  SharedCtor(arena);
  // @@protoc_insertion_point(arena_constructor:kv739.GetRequest)
}
GetRequest::GetRequest(const GetRequest& from)
  : ::PROTOBUF_NAMESPACE_ID::Message() {
  GetRequest* const _this = this; (void)_this;
  new (&_impl_) Impl_{
      decltype(_impl_.key_) {}

    , /*decltype(_impl_._cached_size_)*/{}};

  _internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
  _impl_.key_.InitDefault();
  #ifdef PROTOBUF_FORCE_COPY_DEFAULT_STRING
        _impl_.key_.Set("", GetArenaForAllocation());
  #endif  // PROTOBUF_FORCE_COPY_DEFAULT_STRING
  if (!from._internal_key().empty()) {
    _this->_impl_.key_.Set(from._internal_key(), _this->GetArenaForAllocation());
  }
  // @@protoc_insertion_point(copy_constructor:kv739.GetRequest)
}

inline void GetRequest::SharedCtor(::_pb::Arena* arena) {
  (void)arena;
  new (&_impl_) Impl_{
      decltype(_impl_.key_) {}

    , /*decltype(_impl_._cached_size_)*/{}
  };
  _impl_.key_.InitDefault();
  #ifdef PROTOBUF_FORCE_COPY_DEFAULT_STRING
        _impl_.key_.Set("", GetArenaForAllocation());
  #endif  // PROTOBUF_FORCE_COPY_DEFAULT_STRING
}

GetRequest::~GetRequest() {
  // @@protoc_insertion_point(destructor:kv739.GetRequest)
  if (auto *arena = _internal_metadata_.DeleteReturnArena<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>()) {
  (void)arena;
    return;
  }
  SharedDtor();
}

inline void GetRequest::SharedDtor() {
  ABSL_DCHECK(GetArenaForAllocation() == nullptr);
  _impl_.key_.Destroy();
}

void GetRequest::SetCachedSize(int size) const {
  _impl_._cached_size_.Set(size);
}

void GetRequest::Clear() {
// @@protoc_insertion_point(message_clear_start:kv739.GetRequest)
  ::uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  _impl_.key_.ClearToEmpty();
  _internal_metadata_.Clear<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>();
}

const char* GetRequest::_InternalParse(const char* ptr, ::_pbi::ParseContext* ctx) {
#define CHK_(x) if (PROTOBUF_PREDICT_FALSE(!(x))) goto failure
  while (!ctx->Done(&ptr)) {
    ::uint32_t tag;
    ptr = ::_pbi::ReadTag(ptr, &tag);
    switch (tag >> 3) {
      // string key = 1;
      case 1:
        if (PROTOBUF_PREDICT_TRUE(static_cast<::uint8_t>(tag) == 10)) {
          auto str = _internal_mutable_key();
          ptr = ::_pbi::InlineGreedyStringParser(str, ptr, ctx);
          CHK_(ptr);
          CHK_(::_pbi::VerifyUTF8(str, "kv739.GetRequest.key"));
        } else {
          goto handle_unusual;
        }
        continue;
      default:
        goto handle_unusual;
    }  // switch
  handle_unusual:
    if ((tag == 0) || ((tag & 7) == 4)) {
      CHK_(ptr);
      ctx->SetLastTag(tag);
      goto message_done;
    }
    ptr = UnknownFieldParse(
        tag,
        _internal_metadata_.mutable_unknown_fields<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(),
        ptr, ctx);
    CHK_(ptr != nullptr);
  }  // while
message_done:
  return ptr;
failure:
  ptr = nullptr;
  goto message_done;
#undef CHK_
}

::uint8_t* GetRequest::_InternalSerialize(
    ::uint8_t* target, ::PROTOBUF_NAMESPACE_ID::io::EpsCopyOutputStream* stream) const {
  // @@protoc_insertion_point(serialize_to_array_start:kv739.GetRequest)
  ::uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  // string key = 1;
  if (!this->_internal_key().empty()) {
    const std::string& _s = this->_internal_key();
    ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::VerifyUtf8String(
        _s.data(), static_cast<int>(_s.length()), ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::SERIALIZE, "kv739.GetRequest.key");
    target = stream->WriteStringMaybeAliased(1, _s, target);
  }

  if (PROTOBUF_PREDICT_FALSE(_internal_metadata_.have_unknown_fields())) {
    target = ::_pbi::WireFormat::InternalSerializeUnknownFieldsToArray(
        _internal_metadata_.unknown_fields<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(::PROTOBUF_NAMESPACE_ID::UnknownFieldSet::default_instance), target, stream);
  }
  // @@protoc_insertion_point(serialize_to_array_end:kv739.GetRequest)
  return target;
}

::size_t GetRequest::ByteSizeLong() const {
// @@protoc_insertion_point(message_byte_size_start:kv739.GetRequest)
  ::size_t total_size = 0;

  ::uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  // string key = 1;
  if (!this->_internal_key().empty()) {
    total_size += 1 + ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::StringSize(
                                    this->_internal_key());
  }

  return MaybeComputeUnknownFieldsSize(total_size, &_impl_._cached_size_);
}

const ::PROTOBUF_NAMESPACE_ID::Message::ClassData GetRequest::_class_data_ = {
    ::PROTOBUF_NAMESPACE_ID::Message::CopyWithSourceCheck,
    GetRequest::MergeImpl
};
const ::PROTOBUF_NAMESPACE_ID::Message::ClassData*GetRequest::GetClassData() const { return &_class_data_; }


void GetRequest::MergeImpl(::PROTOBUF_NAMESPACE_ID::Message& to_msg, const ::PROTOBUF_NAMESPACE_ID::Message& from_msg) {
  auto* const _this = static_cast<GetRequest*>(&to_msg);
  auto& from = static_cast<const GetRequest&>(from_msg);
  // @@protoc_insertion_point(class_specific_merge_from_start:kv739.GetRequest)
  ABSL_DCHECK_NE(&from, _this);
  ::uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  if (!from._internal_key().empty()) {
    _this->_internal_set_key(from._internal_key());
  }
  _this->_internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
}

void GetRequest::CopyFrom(const GetRequest& from) {
// @@protoc_insertion_point(class_specific_copy_from_start:kv739.GetRequest)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

bool GetRequest::IsInitialized() const {
  return true;
}

void GetRequest::InternalSwap(GetRequest* other) {
  using std::swap;
  auto* lhs_arena = GetArenaForAllocation();
  auto* rhs_arena = other->GetArenaForAllocation();
  _internal_metadata_.InternalSwap(&other->_internal_metadata_);
  ::_pbi::ArenaStringPtr::InternalSwap(&_impl_.key_, lhs_arena,
                                       &other->_impl_.key_, rhs_arena);
}

::PROTOBUF_NAMESPACE_ID::Metadata GetRequest::GetMetadata() const {
  return ::_pbi::AssignDescriptors(
      &descriptor_table_kv739_2eproto_getter, &descriptor_table_kv739_2eproto_once,
      file_level_metadata_kv739_2eproto[0]);
}
// ===================================================================

class GetResponse::_Internal {
 public:
};

GetResponse::GetResponse(::PROTOBUF_NAMESPACE_ID::Arena* arena)
  : ::PROTOBUF_NAMESPACE_ID::Message(arena) {
  SharedCtor(arena);
  // @@protoc_insertion_point(arena_constructor:kv739.GetResponse)
}
GetResponse::GetResponse(const GetResponse& from)
  : ::PROTOBUF_NAMESPACE_ID::Message() {
  GetResponse* const _this = this; (void)_this;
  new (&_impl_) Impl_{
      decltype(_impl_.value_) {}

    , decltype(_impl_.status_) {}

    , /*decltype(_impl_._cached_size_)*/{}};

  _internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
  _impl_.value_.InitDefault();
  #ifdef PROTOBUF_FORCE_COPY_DEFAULT_STRING
        _impl_.value_.Set("", GetArenaForAllocation());
  #endif  // PROTOBUF_FORCE_COPY_DEFAULT_STRING
  if (!from._internal_value().empty()) {
    _this->_impl_.value_.Set(from._internal_value(), _this->GetArenaForAllocation());
  }
  _this->_impl_.status_ = from._impl_.status_;
  // @@protoc_insertion_point(copy_constructor:kv739.GetResponse)
}

inline void GetResponse::SharedCtor(::_pb::Arena* arena) {
  (void)arena;
  new (&_impl_) Impl_{
      decltype(_impl_.value_) {}

    , decltype(_impl_.status_) { 0 }

    , /*decltype(_impl_._cached_size_)*/{}
  };
  _impl_.value_.InitDefault();
  #ifdef PROTOBUF_FORCE_COPY_DEFAULT_STRING
        _impl_.value_.Set("", GetArenaForAllocation());
  #endif  // PROTOBUF_FORCE_COPY_DEFAULT_STRING
}

GetResponse::~GetResponse() {
  // @@protoc_insertion_point(destructor:kv739.GetResponse)
  if (auto *arena = _internal_metadata_.DeleteReturnArena<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>()) {
  (void)arena;
    return;
  }
  SharedDtor();
}

inline void GetResponse::SharedDtor() {
  ABSL_DCHECK(GetArenaForAllocation() == nullptr);
  _impl_.value_.Destroy();
}

void GetResponse::SetCachedSize(int size) const {
  _impl_._cached_size_.Set(size);
}

void GetResponse::Clear() {
// @@protoc_insertion_point(message_clear_start:kv739.GetResponse)
  ::uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  _impl_.value_.ClearToEmpty();
  _impl_.status_ = 0;
  _internal_metadata_.Clear<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>();
}

const char* GetResponse::_InternalParse(const char* ptr, ::_pbi::ParseContext* ctx) {
#define CHK_(x) if (PROTOBUF_PREDICT_FALSE(!(x))) goto failure
  while (!ctx->Done(&ptr)) {
    ::uint32_t tag;
    ptr = ::_pbi::ReadTag(ptr, &tag);
    switch (tag >> 3) {
      // int32 status = 1;
      case 1:
        if (PROTOBUF_PREDICT_TRUE(static_cast<::uint8_t>(tag) == 8)) {
          _impl_.status_ = ::PROTOBUF_NAMESPACE_ID::internal::ReadVarint32(&ptr);
          CHK_(ptr);
        } else {
          goto handle_unusual;
        }
        continue;
      // string value = 2;
      case 2:
        if (PROTOBUF_PREDICT_TRUE(static_cast<::uint8_t>(tag) == 18)) {
          auto str = _internal_mutable_value();
          ptr = ::_pbi::InlineGreedyStringParser(str, ptr, ctx);
          CHK_(ptr);
          CHK_(::_pbi::VerifyUTF8(str, "kv739.GetResponse.value"));
        } else {
          goto handle_unusual;
        }
        continue;
      default:
        goto handle_unusual;
    }  // switch
  handle_unusual:
    if ((tag == 0) || ((tag & 7) == 4)) {
      CHK_(ptr);
      ctx->SetLastTag(tag);
      goto message_done;
    }
    ptr = UnknownFieldParse(
        tag,
        _internal_metadata_.mutable_unknown_fields<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(),
        ptr, ctx);
    CHK_(ptr != nullptr);
  }  // while
message_done:
  return ptr;
failure:
  ptr = nullptr;
  goto message_done;
#undef CHK_
}

::uint8_t* GetResponse::_InternalSerialize(
    ::uint8_t* target, ::PROTOBUF_NAMESPACE_ID::io::EpsCopyOutputStream* stream) const {
  // @@protoc_insertion_point(serialize_to_array_start:kv739.GetResponse)
  ::uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  // int32 status = 1;
  if (this->_internal_status() != 0) {
    target = stream->EnsureSpace(target);
    target = ::_pbi::WireFormatLite::WriteInt32ToArray(
        1, this->_internal_status(), target);
  }

  // string value = 2;
  if (!this->_internal_value().empty()) {
    const std::string& _s = this->_internal_value();
    ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::VerifyUtf8String(
        _s.data(), static_cast<int>(_s.length()), ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::SERIALIZE, "kv739.GetResponse.value");
    target = stream->WriteStringMaybeAliased(2, _s, target);
  }

  if (PROTOBUF_PREDICT_FALSE(_internal_metadata_.have_unknown_fields())) {
    target = ::_pbi::WireFormat::InternalSerializeUnknownFieldsToArray(
        _internal_metadata_.unknown_fields<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(::PROTOBUF_NAMESPACE_ID::UnknownFieldSet::default_instance), target, stream);
  }
  // @@protoc_insertion_point(serialize_to_array_end:kv739.GetResponse)
  return target;
}

::size_t GetResponse::ByteSizeLong() const {
// @@protoc_insertion_point(message_byte_size_start:kv739.GetResponse)
  ::size_t total_size = 0;

  ::uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  // string value = 2;
  if (!this->_internal_value().empty()) {
    total_size += 1 + ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::StringSize(
                                    this->_internal_value());
  }

  // int32 status = 1;
  if (this->_internal_status() != 0) {
    total_size += ::_pbi::WireFormatLite::Int32SizePlusOne(
        this->_internal_status());
  }

  return MaybeComputeUnknownFieldsSize(total_size, &_impl_._cached_size_);
}

const ::PROTOBUF_NAMESPACE_ID::Message::ClassData GetResponse::_class_data_ = {
    ::PROTOBUF_NAMESPACE_ID::Message::CopyWithSourceCheck,
    GetResponse::MergeImpl
};
const ::PROTOBUF_NAMESPACE_ID::Message::ClassData*GetResponse::GetClassData() const { return &_class_data_; }


void GetResponse::MergeImpl(::PROTOBUF_NAMESPACE_ID::Message& to_msg, const ::PROTOBUF_NAMESPACE_ID::Message& from_msg) {
  auto* const _this = static_cast<GetResponse*>(&to_msg);
  auto& from = static_cast<const GetResponse&>(from_msg);
  // @@protoc_insertion_point(class_specific_merge_from_start:kv739.GetResponse)
  ABSL_DCHECK_NE(&from, _this);
  ::uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  if (!from._internal_value().empty()) {
    _this->_internal_set_value(from._internal_value());
  }
  if (from._internal_status() != 0) {
    _this->_internal_set_status(from._internal_status());
  }
  _this->_internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
}

void GetResponse::CopyFrom(const GetResponse& from) {
// @@protoc_insertion_point(class_specific_copy_from_start:kv739.GetResponse)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

bool GetResponse::IsInitialized() const {
  return true;
}

void GetResponse::InternalSwap(GetResponse* other) {
  using std::swap;
  auto* lhs_arena = GetArenaForAllocation();
  auto* rhs_arena = other->GetArenaForAllocation();
  _internal_metadata_.InternalSwap(&other->_internal_metadata_);
  ::_pbi::ArenaStringPtr::InternalSwap(&_impl_.value_, lhs_arena,
                                       &other->_impl_.value_, rhs_arena);

  swap(_impl_.status_, other->_impl_.status_);
}

::PROTOBUF_NAMESPACE_ID::Metadata GetResponse::GetMetadata() const {
  return ::_pbi::AssignDescriptors(
      &descriptor_table_kv739_2eproto_getter, &descriptor_table_kv739_2eproto_once,
      file_level_metadata_kv739_2eproto[1]);
}
// ===================================================================

class PutRequest::_Internal {
 public:
};

PutRequest::PutRequest(::PROTOBUF_NAMESPACE_ID::Arena* arena)
  : ::PROTOBUF_NAMESPACE_ID::Message(arena) {
  SharedCtor(arena);
  // @@protoc_insertion_point(arena_constructor:kv739.PutRequest)
}
PutRequest::PutRequest(const PutRequest& from)
  : ::PROTOBUF_NAMESPACE_ID::Message() {
  PutRequest* const _this = this; (void)_this;
  new (&_impl_) Impl_{
      decltype(_impl_.key_) {}

    , decltype(_impl_.value_) {}

    , /*decltype(_impl_._cached_size_)*/{}};

  _internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
  _impl_.key_.InitDefault();
  #ifdef PROTOBUF_FORCE_COPY_DEFAULT_STRING
        _impl_.key_.Set("", GetArenaForAllocation());
  #endif  // PROTOBUF_FORCE_COPY_DEFAULT_STRING
  if (!from._internal_key().empty()) {
    _this->_impl_.key_.Set(from._internal_key(), _this->GetArenaForAllocation());
  }
  _impl_.value_.InitDefault();
  #ifdef PROTOBUF_FORCE_COPY_DEFAULT_STRING
        _impl_.value_.Set("", GetArenaForAllocation());
  #endif  // PROTOBUF_FORCE_COPY_DEFAULT_STRING
  if (!from._internal_value().empty()) {
    _this->_impl_.value_.Set(from._internal_value(), _this->GetArenaForAllocation());
  }
  // @@protoc_insertion_point(copy_constructor:kv739.PutRequest)
}

inline void PutRequest::SharedCtor(::_pb::Arena* arena) {
  (void)arena;
  new (&_impl_) Impl_{
      decltype(_impl_.key_) {}

    , decltype(_impl_.value_) {}

    , /*decltype(_impl_._cached_size_)*/{}
  };
  _impl_.key_.InitDefault();
  #ifdef PROTOBUF_FORCE_COPY_DEFAULT_STRING
        _impl_.key_.Set("", GetArenaForAllocation());
  #endif  // PROTOBUF_FORCE_COPY_DEFAULT_STRING
  _impl_.value_.InitDefault();
  #ifdef PROTOBUF_FORCE_COPY_DEFAULT_STRING
        _impl_.value_.Set("", GetArenaForAllocation());
  #endif  // PROTOBUF_FORCE_COPY_DEFAULT_STRING
}

PutRequest::~PutRequest() {
  // @@protoc_insertion_point(destructor:kv739.PutRequest)
  if (auto *arena = _internal_metadata_.DeleteReturnArena<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>()) {
  (void)arena;
    return;
  }
  SharedDtor();
}

inline void PutRequest::SharedDtor() {
  ABSL_DCHECK(GetArenaForAllocation() == nullptr);
  _impl_.key_.Destroy();
  _impl_.value_.Destroy();
}

void PutRequest::SetCachedSize(int size) const {
  _impl_._cached_size_.Set(size);
}

void PutRequest::Clear() {
// @@protoc_insertion_point(message_clear_start:kv739.PutRequest)
  ::uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  _impl_.key_.ClearToEmpty();
  _impl_.value_.ClearToEmpty();
  _internal_metadata_.Clear<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>();
}

const char* PutRequest::_InternalParse(const char* ptr, ::_pbi::ParseContext* ctx) {
#define CHK_(x) if (PROTOBUF_PREDICT_FALSE(!(x))) goto failure
  while (!ctx->Done(&ptr)) {
    ::uint32_t tag;
    ptr = ::_pbi::ReadTag(ptr, &tag);
    switch (tag >> 3) {
      // string key = 1;
      case 1:
        if (PROTOBUF_PREDICT_TRUE(static_cast<::uint8_t>(tag) == 10)) {
          auto str = _internal_mutable_key();
          ptr = ::_pbi::InlineGreedyStringParser(str, ptr, ctx);
          CHK_(ptr);
          CHK_(::_pbi::VerifyUTF8(str, "kv739.PutRequest.key"));
        } else {
          goto handle_unusual;
        }
        continue;
      // string value = 2;
      case 2:
        if (PROTOBUF_PREDICT_TRUE(static_cast<::uint8_t>(tag) == 18)) {
          auto str = _internal_mutable_value();
          ptr = ::_pbi::InlineGreedyStringParser(str, ptr, ctx);
          CHK_(ptr);
          CHK_(::_pbi::VerifyUTF8(str, "kv739.PutRequest.value"));
        } else {
          goto handle_unusual;
        }
        continue;
      default:
        goto handle_unusual;
    }  // switch
  handle_unusual:
    if ((tag == 0) || ((tag & 7) == 4)) {
      CHK_(ptr);
      ctx->SetLastTag(tag);
      goto message_done;
    }
    ptr = UnknownFieldParse(
        tag,
        _internal_metadata_.mutable_unknown_fields<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(),
        ptr, ctx);
    CHK_(ptr != nullptr);
  }  // while
message_done:
  return ptr;
failure:
  ptr = nullptr;
  goto message_done;
#undef CHK_
}

::uint8_t* PutRequest::_InternalSerialize(
    ::uint8_t* target, ::PROTOBUF_NAMESPACE_ID::io::EpsCopyOutputStream* stream) const {
  // @@protoc_insertion_point(serialize_to_array_start:kv739.PutRequest)
  ::uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  // string key = 1;
  if (!this->_internal_key().empty()) {
    const std::string& _s = this->_internal_key();
    ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::VerifyUtf8String(
        _s.data(), static_cast<int>(_s.length()), ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::SERIALIZE, "kv739.PutRequest.key");
    target = stream->WriteStringMaybeAliased(1, _s, target);
  }

  // string value = 2;
  if (!this->_internal_value().empty()) {
    const std::string& _s = this->_internal_value();
    ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::VerifyUtf8String(
        _s.data(), static_cast<int>(_s.length()), ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::SERIALIZE, "kv739.PutRequest.value");
    target = stream->WriteStringMaybeAliased(2, _s, target);
  }

  if (PROTOBUF_PREDICT_FALSE(_internal_metadata_.have_unknown_fields())) {
    target = ::_pbi::WireFormat::InternalSerializeUnknownFieldsToArray(
        _internal_metadata_.unknown_fields<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(::PROTOBUF_NAMESPACE_ID::UnknownFieldSet::default_instance), target, stream);
  }
  // @@protoc_insertion_point(serialize_to_array_end:kv739.PutRequest)
  return target;
}

::size_t PutRequest::ByteSizeLong() const {
// @@protoc_insertion_point(message_byte_size_start:kv739.PutRequest)
  ::size_t total_size = 0;

  ::uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  // string key = 1;
  if (!this->_internal_key().empty()) {
    total_size += 1 + ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::StringSize(
                                    this->_internal_key());
  }

  // string value = 2;
  if (!this->_internal_value().empty()) {
    total_size += 1 + ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::StringSize(
                                    this->_internal_value());
  }

  return MaybeComputeUnknownFieldsSize(total_size, &_impl_._cached_size_);
}

const ::PROTOBUF_NAMESPACE_ID::Message::ClassData PutRequest::_class_data_ = {
    ::PROTOBUF_NAMESPACE_ID::Message::CopyWithSourceCheck,
    PutRequest::MergeImpl
};
const ::PROTOBUF_NAMESPACE_ID::Message::ClassData*PutRequest::GetClassData() const { return &_class_data_; }


void PutRequest::MergeImpl(::PROTOBUF_NAMESPACE_ID::Message& to_msg, const ::PROTOBUF_NAMESPACE_ID::Message& from_msg) {
  auto* const _this = static_cast<PutRequest*>(&to_msg);
  auto& from = static_cast<const PutRequest&>(from_msg);
  // @@protoc_insertion_point(class_specific_merge_from_start:kv739.PutRequest)
  ABSL_DCHECK_NE(&from, _this);
  ::uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  if (!from._internal_key().empty()) {
    _this->_internal_set_key(from._internal_key());
  }
  if (!from._internal_value().empty()) {
    _this->_internal_set_value(from._internal_value());
  }
  _this->_internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
}

void PutRequest::CopyFrom(const PutRequest& from) {
// @@protoc_insertion_point(class_specific_copy_from_start:kv739.PutRequest)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

bool PutRequest::IsInitialized() const {
  return true;
}

void PutRequest::InternalSwap(PutRequest* other) {
  using std::swap;
  auto* lhs_arena = GetArenaForAllocation();
  auto* rhs_arena = other->GetArenaForAllocation();
  _internal_metadata_.InternalSwap(&other->_internal_metadata_);
  ::_pbi::ArenaStringPtr::InternalSwap(&_impl_.key_, lhs_arena,
                                       &other->_impl_.key_, rhs_arena);
  ::_pbi::ArenaStringPtr::InternalSwap(&_impl_.value_, lhs_arena,
                                       &other->_impl_.value_, rhs_arena);
}

::PROTOBUF_NAMESPACE_ID::Metadata PutRequest::GetMetadata() const {
  return ::_pbi::AssignDescriptors(
      &descriptor_table_kv739_2eproto_getter, &descriptor_table_kv739_2eproto_once,
      file_level_metadata_kv739_2eproto[2]);
}
// ===================================================================

class PutResponse::_Internal {
 public:
};

PutResponse::PutResponse(::PROTOBUF_NAMESPACE_ID::Arena* arena)
  : ::PROTOBUF_NAMESPACE_ID::Message(arena) {
  SharedCtor(arena);
  // @@protoc_insertion_point(arena_constructor:kv739.PutResponse)
}
PutResponse::PutResponse(const PutResponse& from)
  : ::PROTOBUF_NAMESPACE_ID::Message() {
  PutResponse* const _this = this; (void)_this;
  new (&_impl_) Impl_{
      decltype(_impl_.old_value_) {}

    , decltype(_impl_.status_) {}

    , /*decltype(_impl_._cached_size_)*/{}};

  _internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
  _impl_.old_value_.InitDefault();
  #ifdef PROTOBUF_FORCE_COPY_DEFAULT_STRING
        _impl_.old_value_.Set("", GetArenaForAllocation());
  #endif  // PROTOBUF_FORCE_COPY_DEFAULT_STRING
  if (!from._internal_old_value().empty()) {
    _this->_impl_.old_value_.Set(from._internal_old_value(), _this->GetArenaForAllocation());
  }
  _this->_impl_.status_ = from._impl_.status_;
  // @@protoc_insertion_point(copy_constructor:kv739.PutResponse)
}

inline void PutResponse::SharedCtor(::_pb::Arena* arena) {
  (void)arena;
  new (&_impl_) Impl_{
      decltype(_impl_.old_value_) {}

    , decltype(_impl_.status_) { 0 }

    , /*decltype(_impl_._cached_size_)*/{}
  };
  _impl_.old_value_.InitDefault();
  #ifdef PROTOBUF_FORCE_COPY_DEFAULT_STRING
        _impl_.old_value_.Set("", GetArenaForAllocation());
  #endif  // PROTOBUF_FORCE_COPY_DEFAULT_STRING
}

PutResponse::~PutResponse() {
  // @@protoc_insertion_point(destructor:kv739.PutResponse)
  if (auto *arena = _internal_metadata_.DeleteReturnArena<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>()) {
  (void)arena;
    return;
  }
  SharedDtor();
}

inline void PutResponse::SharedDtor() {
  ABSL_DCHECK(GetArenaForAllocation() == nullptr);
  _impl_.old_value_.Destroy();
}

void PutResponse::SetCachedSize(int size) const {
  _impl_._cached_size_.Set(size);
}

void PutResponse::Clear() {
// @@protoc_insertion_point(message_clear_start:kv739.PutResponse)
  ::uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  _impl_.old_value_.ClearToEmpty();
  _impl_.status_ = 0;
  _internal_metadata_.Clear<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>();
}

const char* PutResponse::_InternalParse(const char* ptr, ::_pbi::ParseContext* ctx) {
#define CHK_(x) if (PROTOBUF_PREDICT_FALSE(!(x))) goto failure
  while (!ctx->Done(&ptr)) {
    ::uint32_t tag;
    ptr = ::_pbi::ReadTag(ptr, &tag);
    switch (tag >> 3) {
      // int32 status = 1;
      case 1:
        if (PROTOBUF_PREDICT_TRUE(static_cast<::uint8_t>(tag) == 8)) {
          _impl_.status_ = ::PROTOBUF_NAMESPACE_ID::internal::ReadVarint32(&ptr);
          CHK_(ptr);
        } else {
          goto handle_unusual;
        }
        continue;
      // string old_value = 2;
      case 2:
        if (PROTOBUF_PREDICT_TRUE(static_cast<::uint8_t>(tag) == 18)) {
          auto str = _internal_mutable_old_value();
          ptr = ::_pbi::InlineGreedyStringParser(str, ptr, ctx);
          CHK_(ptr);
          CHK_(::_pbi::VerifyUTF8(str, "kv739.PutResponse.old_value"));
        } else {
          goto handle_unusual;
        }
        continue;
      default:
        goto handle_unusual;
    }  // switch
  handle_unusual:
    if ((tag == 0) || ((tag & 7) == 4)) {
      CHK_(ptr);
      ctx->SetLastTag(tag);
      goto message_done;
    }
    ptr = UnknownFieldParse(
        tag,
        _internal_metadata_.mutable_unknown_fields<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(),
        ptr, ctx);
    CHK_(ptr != nullptr);
  }  // while
message_done:
  return ptr;
failure:
  ptr = nullptr;
  goto message_done;
#undef CHK_
}

::uint8_t* PutResponse::_InternalSerialize(
    ::uint8_t* target, ::PROTOBUF_NAMESPACE_ID::io::EpsCopyOutputStream* stream) const {
  // @@protoc_insertion_point(serialize_to_array_start:kv739.PutResponse)
  ::uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  // int32 status = 1;
  if (this->_internal_status() != 0) {
    target = stream->EnsureSpace(target);
    target = ::_pbi::WireFormatLite::WriteInt32ToArray(
        1, this->_internal_status(), target);
  }

  // string old_value = 2;
  if (!this->_internal_old_value().empty()) {
    const std::string& _s = this->_internal_old_value();
    ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::VerifyUtf8String(
        _s.data(), static_cast<int>(_s.length()), ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::SERIALIZE, "kv739.PutResponse.old_value");
    target = stream->WriteStringMaybeAliased(2, _s, target);
  }

  if (PROTOBUF_PREDICT_FALSE(_internal_metadata_.have_unknown_fields())) {
    target = ::_pbi::WireFormat::InternalSerializeUnknownFieldsToArray(
        _internal_metadata_.unknown_fields<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(::PROTOBUF_NAMESPACE_ID::UnknownFieldSet::default_instance), target, stream);
  }
  // @@protoc_insertion_point(serialize_to_array_end:kv739.PutResponse)
  return target;
}

::size_t PutResponse::ByteSizeLong() const {
// @@protoc_insertion_point(message_byte_size_start:kv739.PutResponse)
  ::size_t total_size = 0;

  ::uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  // string old_value = 2;
  if (!this->_internal_old_value().empty()) {
    total_size += 1 + ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::StringSize(
                                    this->_internal_old_value());
  }

  // int32 status = 1;
  if (this->_internal_status() != 0) {
    total_size += ::_pbi::WireFormatLite::Int32SizePlusOne(
        this->_internal_status());
  }

  return MaybeComputeUnknownFieldsSize(total_size, &_impl_._cached_size_);
}

const ::PROTOBUF_NAMESPACE_ID::Message::ClassData PutResponse::_class_data_ = {
    ::PROTOBUF_NAMESPACE_ID::Message::CopyWithSourceCheck,
    PutResponse::MergeImpl
};
const ::PROTOBUF_NAMESPACE_ID::Message::ClassData*PutResponse::GetClassData() const { return &_class_data_; }


void PutResponse::MergeImpl(::PROTOBUF_NAMESPACE_ID::Message& to_msg, const ::PROTOBUF_NAMESPACE_ID::Message& from_msg) {
  auto* const _this = static_cast<PutResponse*>(&to_msg);
  auto& from = static_cast<const PutResponse&>(from_msg);
  // @@protoc_insertion_point(class_specific_merge_from_start:kv739.PutResponse)
  ABSL_DCHECK_NE(&from, _this);
  ::uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  if (!from._internal_old_value().empty()) {
    _this->_internal_set_old_value(from._internal_old_value());
  }
  if (from._internal_status() != 0) {
    _this->_internal_set_status(from._internal_status());
  }
  _this->_internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
}

void PutResponse::CopyFrom(const PutResponse& from) {
// @@protoc_insertion_point(class_specific_copy_from_start:kv739.PutResponse)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

bool PutResponse::IsInitialized() const {
  return true;
}

void PutResponse::InternalSwap(PutResponse* other) {
  using std::swap;
  auto* lhs_arena = GetArenaForAllocation();
  auto* rhs_arena = other->GetArenaForAllocation();
  _internal_metadata_.InternalSwap(&other->_internal_metadata_);
  ::_pbi::ArenaStringPtr::InternalSwap(&_impl_.old_value_, lhs_arena,
                                       &other->_impl_.old_value_, rhs_arena);

  swap(_impl_.status_, other->_impl_.status_);
}

::PROTOBUF_NAMESPACE_ID::Metadata PutResponse::GetMetadata() const {
  return ::_pbi::AssignDescriptors(
      &descriptor_table_kv739_2eproto_getter, &descriptor_table_kv739_2eproto_once,
      file_level_metadata_kv739_2eproto[3]);
}
// @@protoc_insertion_point(namespace_scope)
}  // namespace kv739
PROTOBUF_NAMESPACE_OPEN
template<> PROTOBUF_NOINLINE ::kv739::GetRequest*
Arena::CreateMaybeMessage< ::kv739::GetRequest >(Arena* arena) {
  return Arena::CreateMessageInternal< ::kv739::GetRequest >(arena);
}
template<> PROTOBUF_NOINLINE ::kv739::GetResponse*
Arena::CreateMaybeMessage< ::kv739::GetResponse >(Arena* arena) {
  return Arena::CreateMessageInternal< ::kv739::GetResponse >(arena);
}
template<> PROTOBUF_NOINLINE ::kv739::PutRequest*
Arena::CreateMaybeMessage< ::kv739::PutRequest >(Arena* arena) {
  return Arena::CreateMessageInternal< ::kv739::PutRequest >(arena);
}
template<> PROTOBUF_NOINLINE ::kv739::PutResponse*
Arena::CreateMaybeMessage< ::kv739::PutResponse >(Arena* arena) {
  return Arena::CreateMessageInternal< ::kv739::PutResponse >(arena);
}
PROTOBUF_NAMESPACE_CLOSE
// @@protoc_insertion_point(global_scope)
#include "google/protobuf/port_undef.inc"