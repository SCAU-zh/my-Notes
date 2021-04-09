# Grpc java Server使用梳理

### grpc depenencies引入

``` grade
ext {
	protobufVersion = '3.3.0'
	grpcVersion = '1.34.1'
	codingClientVersion = "20210204.1"
}

dependencies {
	implementation "net.coding.proto:common:20201109.2"
	implementation "io.grpc:grpc-netty:${grpcVersion}"
	implementation "io.grpc:grpc-protobuf:${grpcVersion}"
	implementation "io.grpc:grpc-stub:${grpcVersion}"
	implementation "io.grpc:grpc-services:${grpcVersion}"
	//implementation "net.coding.common:rpc-client:20201026.1"
	implementation "net.coding.common:rpc-client:${codingClientVersion}"
	implementation ("net.coding.grpc.client:platform:20210322.1"){
		exclude group: "com.google.gwt", module: "gwt-user"
	}
	implementation "io.github.lognet:grpc-spring-boot-starter:4.4.0"
	implementation "net.coding.proto:platform:20210301.2"
	implementation "net.coding.grpc.client:message:20210204.1"
	implementation "net.coding.proto:message:20210301.1"
	implementation 'me.dinowernli:java-grpc-prometheus:0.3.0'
	// https://mvnrepository.com/artifact/com.google.code.gson/gson
	implementation group: 'com.google.code.gson', name: 'gson', version: '2.8.6'
}
// protobuf生成插件
protobuf {
	protoc {
		artifact = "com.google.protobuf:protoc:${protobufVersion}"
	}
	generatedFilesBaseDir = "$projectDir/src/generated"
	clean {
		delete generatedFilesBaseDir
	}
	plugins {
		grpc {
			artifact = "io.grpc:protoc-gen-grpc-java:${grpcVersion}"
		}
	}
	generateProtoTasks {
		all()*.plugins {
			grpc {}
		}
	}
}
```

### .proto 编写

``` protobuf
syntax = "proto3";//表示使用proto3

package proto;
//引用其他的proto
import "code.proto";
import "google/protobuf/timestamp.proto";
//生成名称位置等
option java_generic_services = true;
option java_package = "net.coding.proto.invite";
option java_outer_classname = "inviteRecordProto";

//定义message
message RecordRequest {
  int32 inviterId = 1;
  int32 teamId = 2;
  google.protobuf.Timestamp inviteAt = 3;
}

message RecordResponse {
  proto.common.Code Code = 1;
  string msg = 2;
}

//定义service
service InviteRecordService {
  rpc createInviteRecord(RecordRequest) returns(RecordResponse);
}
```

### 使用proto plugins自动生成所需类

![image-20210409115932780](/Users/lzhonglin/Library/Application Support/typora-user-images/image-20210409115932780.png)

![image-20210409115949800](/Users/lzhonglin/Library/Application Support/typora-user-images/image-20210409115949800.png)

### 编写GrpcService类

``` java
package net.coding.operating.api;

import com.google.protobuf.Timestamp;
import io.grpc.stub.StreamObserver;
import lombok.extern.slf4j.Slf4j;
import net.coding.operating.api.entities.operating.Invite;
import net.coding.operating.api.repositories.operating.InviteRepository;
import net.coding.proto.invite.InviteRecordServiceGrpc;
import net.coding.proto.invite.inviteRecordProto.RecordRequest;
import net.coding.proto.invite.inviteRecordProto.RecordResponse;
import org.lognet.springboot.grpc.GRpcService;
import org.springframework.beans.factory.annotation.Autowired;
import proto.common.CodeProto;

/**
 * @title: InviteRecordGrpcService
 * @description: 邀请记录GrpcService
 * @Author LinZihong
 * @Date: 2021/4/8 10:51 上午
 * @Version 1.0
 */
@Slf4j
@GRpcService
public class InviteRecordGrpcService extends InviteRecordServiceGrpc.InviteRecordServiceImplBase {

  @Autowired
  private InviteRepository inviteRepository;

  @Override
  public void createInviteRecord(RecordRequest request,
      StreamObserver<RecordResponse> responseObserver) {
    Integer userId = request.getInviterId();
    Integer teamId = request.getTeamId();
    Timestamp inviteAt = request.getInviteAt();

    Invite invite = new Invite();
    invite.setInviteAt(new java.sql.Timestamp(inviteAt.getSeconds()));
    invite.setUserId(userId);
    invite.setTeamId(teamId);
    invite.setCreatedAt(new java.sql.Timestamp(System.currentTimeMillis()));
    invite.setUpdatedAt(new java.sql.Timestamp(System.currentTimeMillis()));
    //入库
    inviteRepository.save(invite);

    RecordResponse response = RecordResponse
        .newBuilder()
        .setCode(CodeProto.Code.SUCCESS)
        .build();
    responseObserver.onNext(response);
    responseObserver.onCompleted();
  }
}
```

* service需要实现xxxxxxServiceGrpc.xxxxxxxxServiceImplBase接口，并实现其中的方法
* @GrpcService注解将自动配置server，相关配置可在application.properties/yaml中配置

``` properties
# grpc
grpc.port=20153#监听端口
grpc.enable-reflection=true/#开启反射
grpc.enabled=false#是否开启server
```

