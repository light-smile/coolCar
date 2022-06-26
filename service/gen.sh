
function genProto {
    DOMAIN=$1
     protoc -I=./${DOMAIN}/api/ --go_out=paths=source_relative:${DOMAIN}/api/gen/v1 --go-grpc_out=paths=source_relative:${DOMAIN}/api/gen/v1 ./${DOMAIN}/api/${DOMAIN}.proto
    protoc -I=./${DOMAIN}/api/ --grpc-gateway_out=paths=source_relative,grpc_api_configuration=${DOMAIN}/api/${DOMAIN}.yaml:${DOMAIN}/api/gen/v1 ./${DOMAIN}/api/${DOMAIN}.proto

}

genProto auth
genProto rental





