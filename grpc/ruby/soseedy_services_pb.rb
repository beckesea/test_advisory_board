# Generated by the protocol buffer compiler.  DO NOT EDIT!
# Source: soseedy.proto for package 'soseedy'

require 'grpc'
require 'soseedy_pb'

module Soseedy
  module SoSeedy
    class Service

      include GRPC::GenericService

      self.marshal_class_method = :encode
      self.unmarshal_class_method = :decode
      self.service_name = 'soseedy.SoSeedy'

      rpc :CreateTeacher, CreateTeacherRequest, Teacher
    end

    Stub = Service.rpc_stub_class
  end
end
