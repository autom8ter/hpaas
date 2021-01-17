# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: kdeploy.proto

require 'google/protobuf'

require 'google/protobuf/struct_pb'
require 'google/protobuf/timestamp_pb'
require 'google/protobuf/any_pb'
require 'google/protobuf/empty_pb'
Google::Protobuf::DescriptorPool.generated_pool.build do
  add_message "kdeploy.App" do
    optional :name, :string, 1
    optional :namespace, :string, 2
    optional :image, :string, 3
    repeated :args, :string, 5
    map :env, :string, :string, 6
    map :ports, :string, :uint32, 7
    optional :replicas, :uint32, 8
    optional :status, :message, 9, "kdeploy.AppStatus"
  end
  add_message "kdeploy.Task" do
    optional :name, :string, 1
    optional :namespace, :string, 2
    optional :image, :string, 3
    repeated :args, :string, 5
    map :env, :string, :string, 6
    optional :schedule, :string, 7
  end
  add_message "kdeploy.TaskConstructor" do
    optional :name, :string, 1
    optional :namespace, :string, 2
    optional :image, :string, 3
    repeated :args, :string, 5
    map :env, :string, :string, 6
    optional :schedule, :string, 7
  end
  add_message "kdeploy.TaskUpdate" do
    optional :name, :string, 1
    optional :namespace, :string, 2
    optional :image, :string, 3
    repeated :args, :string, 5
    map :env, :string, :string, 6
    optional :schedule, :string, 7
  end
  add_message "kdeploy.AppConstructor" do
    optional :name, :string, 1
    optional :namespace, :string, 2
    optional :image, :string, 3
    repeated :args, :string, 4
    map :env, :string, :string, 5
    map :ports, :string, :uint32, 6
    optional :replicas, :uint32, 7
  end
  add_message "kdeploy.AppUpdate" do
    optional :name, :string, 1
    optional :namespace, :string, 2
    optional :image, :string, 3
    repeated :args, :string, 4
    map :env, :string, :string, 5
    map :ports, :string, :uint32, 6
    optional :replicas, :uint32, 7
  end
  add_message "kdeploy.Ref" do
    optional :name, :string, 1
    optional :namespace, :string, 2
  end
  add_message "kdeploy.Replica" do
    optional :phase, :string, 1
    optional :condition, :string, 2
    optional :reason, :string, 3
  end
  add_message "kdeploy.AppStatus" do
    repeated :replicas, :message, 1, "kdeploy.Replica"
  end
  add_message "kdeploy.Log" do
    optional :message, :string, 1
  end
  add_message "kdeploy.Apps" do
    repeated :applications, :message, 1, "kdeploy.App"
  end
  add_message "kdeploy.Tasks" do
    repeated :tasks, :message, 1, "kdeploy.Task"
  end
  add_message "kdeploy.Namespace" do
    optional :namespace, :string, 1
  end
  add_message "kdeploy.Namespaces" do
    repeated :namespaces, :string, 1
  end
end

module Kdeploy
  App = Google::Protobuf::DescriptorPool.generated_pool.lookup("kdeploy.App").msgclass
  Task = Google::Protobuf::DescriptorPool.generated_pool.lookup("kdeploy.Task").msgclass
  TaskConstructor = Google::Protobuf::DescriptorPool.generated_pool.lookup("kdeploy.TaskConstructor").msgclass
  TaskUpdate = Google::Protobuf::DescriptorPool.generated_pool.lookup("kdeploy.TaskUpdate").msgclass
  AppConstructor = Google::Protobuf::DescriptorPool.generated_pool.lookup("kdeploy.AppConstructor").msgclass
  AppUpdate = Google::Protobuf::DescriptorPool.generated_pool.lookup("kdeploy.AppUpdate").msgclass
  Ref = Google::Protobuf::DescriptorPool.generated_pool.lookup("kdeploy.Ref").msgclass
  Replica = Google::Protobuf::DescriptorPool.generated_pool.lookup("kdeploy.Replica").msgclass
  AppStatus = Google::Protobuf::DescriptorPool.generated_pool.lookup("kdeploy.AppStatus").msgclass
  Log = Google::Protobuf::DescriptorPool.generated_pool.lookup("kdeploy.Log").msgclass
  Apps = Google::Protobuf::DescriptorPool.generated_pool.lookup("kdeploy.Apps").msgclass
  Tasks = Google::Protobuf::DescriptorPool.generated_pool.lookup("kdeploy.Tasks").msgclass
  Namespace = Google::Protobuf::DescriptorPool.generated_pool.lookup("kdeploy.Namespace").msgclass
  Namespaces = Google::Protobuf::DescriptorPool.generated_pool.lookup("kdeploy.Namespaces").msgclass
end
