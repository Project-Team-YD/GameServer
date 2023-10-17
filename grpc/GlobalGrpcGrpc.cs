// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: protos/global_grpc.proto
// </auto-generated>
#pragma warning disable 0414, 1591, 8981, 0612
#region Designer generated code

using grpc = global::Grpc.Core;

namespace MainGrpcClient {
  public static partial class GlobalGRpcService
  {
    static readonly string __ServiceName = "main_grpc.GlobalGRpcService";

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static void __Helper_SerializeMessage(global::Google.Protobuf.IMessage message, grpc::SerializationContext context)
    {
      #if !GRPC_DISABLE_PROTOBUF_BUFFER_SERIALIZATION
      if (message is global::Google.Protobuf.IBufferMessage)
      {
        context.SetPayloadLength(message.CalculateSize());
        global::Google.Protobuf.MessageExtensions.WriteTo(message, context.GetBufferWriter());
        context.Complete();
        return;
      }
      #endif
      context.Complete(global::Google.Protobuf.MessageExtensions.ToByteArray(message));
    }

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static class __Helper_MessageCache<T>
    {
      public static readonly bool IsBufferMessage = global::System.Reflection.IntrospectionExtensions.GetTypeInfo(typeof(global::Google.Protobuf.IBufferMessage)).IsAssignableFrom(typeof(T));
    }

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static T __Helper_DeserializeMessage<T>(grpc::DeserializationContext context, global::Google.Protobuf.MessageParser<T> parser) where T : global::Google.Protobuf.IMessage<T>
    {
      #if !GRPC_DISABLE_PROTOBUF_BUFFER_SERIALIZATION
      if (__Helper_MessageCache<T>.IsBufferMessage)
      {
        return parser.ParseFrom(context.PayloadAsReadOnlySequence());
      }
      #endif
      return parser.ParseFrom(context.PayloadAsNewBuffer());
    }

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::MainGrpcClient.GlobalGrpcRequest> __Marshaller_main_grpc_GlobalGrpcRequest = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::MainGrpcClient.GlobalGrpcRequest.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::MainGrpcClient.GlobalGrpcResponse> __Marshaller_main_grpc_GlobalGrpcResponse = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::MainGrpcClient.GlobalGrpcResponse.Parser));

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Method<global::MainGrpcClient.GlobalGrpcRequest, global::MainGrpcClient.GlobalGrpcResponse> __Method_GlobalGRpc = new grpc::Method<global::MainGrpcClient.GlobalGrpcRequest, global::MainGrpcClient.GlobalGrpcResponse>(
        grpc::MethodType.Unary,
        __ServiceName,
        "GlobalGRpc",
        __Marshaller_main_grpc_GlobalGrpcRequest,
        __Marshaller_main_grpc_GlobalGrpcResponse);

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Method<global::MainGrpcClient.GlobalGrpcRequest, global::MainGrpcClient.GlobalGrpcResponse> __Method_GlobalGrpcStream = new grpc::Method<global::MainGrpcClient.GlobalGrpcRequest, global::MainGrpcClient.GlobalGrpcResponse>(
        grpc::MethodType.ClientStreaming,
        __ServiceName,
        "GlobalGrpcStream",
        __Marshaller_main_grpc_GlobalGrpcRequest,
        __Marshaller_main_grpc_GlobalGrpcResponse);

    /// <summary>Service descriptor</summary>
    public static global::Google.Protobuf.Reflection.ServiceDescriptor Descriptor
    {
      get { return global::MainGrpcClient.GlobalGrpcReflection.Descriptor.Services[0]; }
    }

    /// <summary>Base class for server-side implementations of GlobalGRpcService</summary>
    [grpc::BindServiceMethod(typeof(GlobalGRpcService), "BindService")]
    public abstract partial class GlobalGRpcServiceBase
    {
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::System.Threading.Tasks.Task<global::MainGrpcClient.GlobalGrpcResponse> GlobalGRpc(global::MainGrpcClient.GlobalGrpcRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::System.Threading.Tasks.Task<global::MainGrpcClient.GlobalGrpcResponse> GlobalGrpcStream(grpc::IAsyncStreamReader<global::MainGrpcClient.GlobalGrpcRequest> requestStream, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

    }

    /// <summary>Client for GlobalGRpcService</summary>
    public partial class GlobalGRpcServiceClient : grpc::ClientBase<GlobalGRpcServiceClient>
    {
      /// <summary>Creates a new client for GlobalGRpcService</summary>
      /// <param name="channel">The channel to use to make remote calls.</param>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public GlobalGRpcServiceClient(grpc::ChannelBase channel) : base(channel)
      {
      }
      /// <summary>Creates a new client for GlobalGRpcService that uses a custom <c>CallInvoker</c>.</summary>
      /// <param name="callInvoker">The callInvoker to use to make remote calls.</param>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public GlobalGRpcServiceClient(grpc::CallInvoker callInvoker) : base(callInvoker)
      {
      }
      /// <summary>Protected parameterless constructor to allow creation of test doubles.</summary>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      protected GlobalGRpcServiceClient() : base()
      {
      }
      /// <summary>Protected constructor to allow creation of configured clients.</summary>
      /// <param name="configuration">The client configuration.</param>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      protected GlobalGRpcServiceClient(ClientBaseConfiguration configuration) : base(configuration)
      {
      }

      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::MainGrpcClient.GlobalGrpcResponse GlobalGRpc(global::MainGrpcClient.GlobalGrpcRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return GlobalGRpc(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::MainGrpcClient.GlobalGrpcResponse GlobalGRpc(global::MainGrpcClient.GlobalGrpcRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_GlobalGRpc, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::MainGrpcClient.GlobalGrpcResponse> GlobalGRpcAsync(global::MainGrpcClient.GlobalGrpcRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return GlobalGRpcAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::MainGrpcClient.GlobalGrpcResponse> GlobalGRpcAsync(global::MainGrpcClient.GlobalGrpcRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_GlobalGRpc, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncClientStreamingCall<global::MainGrpcClient.GlobalGrpcRequest, global::MainGrpcClient.GlobalGrpcResponse> GlobalGrpcStream(grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return GlobalGrpcStream(new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncClientStreamingCall<global::MainGrpcClient.GlobalGrpcRequest, global::MainGrpcClient.GlobalGrpcResponse> GlobalGrpcStream(grpc::CallOptions options)
      {
        return CallInvoker.AsyncClientStreamingCall(__Method_GlobalGrpcStream, null, options);
      }
      /// <summary>Creates a new instance of client from given <c>ClientBaseConfiguration</c>.</summary>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      protected override GlobalGRpcServiceClient NewInstance(ClientBaseConfiguration configuration)
      {
        return new GlobalGRpcServiceClient(configuration);
      }
    }

    /// <summary>Creates service definition that can be registered with a server</summary>
    /// <param name="serviceImpl">An object implementing the server-side handling logic.</param>
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    public static grpc::ServerServiceDefinition BindService(GlobalGRpcServiceBase serviceImpl)
    {
      return grpc::ServerServiceDefinition.CreateBuilder()
          .AddMethod(__Method_GlobalGRpc, serviceImpl.GlobalGRpc)
          .AddMethod(__Method_GlobalGrpcStream, serviceImpl.GlobalGrpcStream).Build();
    }

    /// <summary>Register service method with a service binder with or without implementation. Useful when customizing the service binding logic.
    /// Note: this method is part of an experimental API that can change or be removed without any prior notice.</summary>
    /// <param name="serviceBinder">Service methods will be bound by calling <c>AddMethod</c> on this object.</param>
    /// <param name="serviceImpl">An object implementing the server-side handling logic.</param>
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    public static void BindService(grpc::ServiceBinderBase serviceBinder, GlobalGRpcServiceBase serviceImpl)
    {
      serviceBinder.AddMethod(__Method_GlobalGRpc, serviceImpl == null ? null : new grpc::UnaryServerMethod<global::MainGrpcClient.GlobalGrpcRequest, global::MainGrpcClient.GlobalGrpcResponse>(serviceImpl.GlobalGRpc));
      serviceBinder.AddMethod(__Method_GlobalGrpcStream, serviceImpl == null ? null : new grpc::ClientStreamingServerMethod<global::MainGrpcClient.GlobalGrpcRequest, global::MainGrpcClient.GlobalGrpcResponse>(serviceImpl.GlobalGrpcStream));
    }

  }
}
#endregion