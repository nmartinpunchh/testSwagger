# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: punchh/workflowapi/workflow_api.proto

require 'google/protobuf'

require 'punchh/workflow/workflow_pb'
Google::Protobuf::DescriptorPool.generated_pool.build do
  add_file("punchh/workflowapi/workflow_api.proto", :syntax => :proto3) do
    add_message "workflowapi.CreateWorkflowRequest" do
      optional :workflow, :message, 1, "workflow.Workflow"
    end
    add_message "workflowapi.CreateWorkflowResponse" do
      optional :id, :int64, 1
    end
    add_message "workflowapi.ReadWorkflowRequest" do
      optional :id, :int64, 1
    end
    add_message "workflowapi.ReadWorkflowResponse" do
      optional :workflow, :message, 1, "workflow.Workflow"
    end
  end
end

module Workflowapi
  CreateWorkflowRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("workflowapi.CreateWorkflowRequest").msgclass
  CreateWorkflowResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("workflowapi.CreateWorkflowResponse").msgclass
  ReadWorkflowRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("workflowapi.ReadWorkflowRequest").msgclass
  ReadWorkflowResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("workflowapi.ReadWorkflowResponse").msgclass
end
