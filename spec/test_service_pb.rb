# frozen_string_literal: true
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: test_service.proto

require 'google/protobuf'

require 'google/api/annotations_pb'
require 'google/protobuf/struct_pb'
require 'google/protobuf/timestamp_pb'


descriptor_data = "\n\x12test_service.proto\x12\x08testdata\x1a\x1cgoogle/api/annotations.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\xdc\x02\n\x0bTestRequest\x12\x0f\n\x07test_id\x18\x01 \x01(\t\x12\x0e\n\x06\x66oobar\x18\x02 \x01(\t\x12\x17\n\x0frepeated_string\x18\x03 \x03(\t\x12\'\n\nsub_record\x18\x04 \x01(\x0b\x32\x13.testdata.SubRecord\x12*\n\rsecond_record\x18\x05 \x01(\x0b\x32\x13.testdata.SubRecord\x12-\n\x0cstruct_field\x18\x06 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x33\n\x0ftimestamp_field\x18\x07 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12.\n\nlist_value\x18\x08 \x01(\x0b\x32\x1a.google.protobuf.ListValue\x12*\n\nbare_value\x18\t \x01(\x0b\x32\x16.google.protobuf.Value\"/\n\tSubRecord\x12\x0e\n\x06sub_id\x18\x01 \x01(\t\x12\x12\n\nanother_id\x18\x02 \x01(\t\"7\n\x0cTestResponse\x12\x10\n\x08some_int\x18\x01 \x01(\x05\x12\x15\n\rfull_response\x18\x02 \x01(\t2\xdf\x02\n\tMyService\x12T\n\x04Test\x12\x15.testdata.TestRequest\x1a\x16.testdata.TestResponse\"\x1d\x82\xd3\xe4\x93\x02\x17\x12\x15/test/{foobar=blah/*}\x12U\n\x05Test2\x12\x15.testdata.TestRequest\x1a\x16.testdata.TestResponse\"\x1d\x82\xd3\xe4\x93\x02\x17\"\x06/test2:\rsecond_record\x12Z\n\x05Test3\x12\x15.testdata.TestRequest\x1a\x16.testdata.TestResponse\"\"\x82\xd3\xe4\x93\x02\x1c\"\x1a/test3/{sub_record.sub_id}\x12I\n\x05Test4\x12\x15.testdata.TestRequest\x1a\x16.testdata.TestResponse\"\x11\x82\xd3\xe4\x93\x02\x0b\"\x06/test4:\x01*b\x06proto3"

pool = Google::Protobuf::DescriptorPool.generated_pool

begin
  pool.add_serialized_file(descriptor_data)
rescue TypeError
  # Compatibility code: will be removed in the next major version.
  require 'google/protobuf/descriptor_pb'
  parsed = Google::Protobuf::FileDescriptorProto.decode(descriptor_data)
  parsed.clear_dependency
  serialized = parsed.class.encode(parsed)
  file = pool.add_serialized_file(serialized)
  warn "Warning: Protobuf detected an import path issue while loading generated file #{__FILE__}"
  imports = [
    ["google.protobuf.Struct", "google/protobuf/struct.proto"],
    ["google.protobuf.Timestamp", "google/protobuf/timestamp.proto"],
  ]
  imports.each do |type_name, expected_filename|
    import_file = pool.lookup(type_name).file_descriptor
    if import_file.name != expected_filename
      warn "- #{file.name} imports #{expected_filename}, but that import was loaded as #{import_file.name}"
    end
  end
  warn "Each proto file must use a consistent fully-qualified name."
  warn "This will become an error in the next major version."
end

module Testdata
  TestRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("testdata.TestRequest").msgclass
  SubRecord = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("testdata.SubRecord").msgclass
  TestResponse = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("testdata.TestResponse").msgclass
end
