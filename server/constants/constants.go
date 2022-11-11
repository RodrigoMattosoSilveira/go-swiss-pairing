package constants

const CertDir = "./cert/"
const ServerCertPem = CertDir + "server-cert.pem"
const ServerKeyPem = CertDir + "server-key.pem"

/**
Status codes and their use in gRPC
See https://grpc.github.io/grpc/core/md_doc_statuscodes.html for details
*/

// Continue - This interim response indicates that the client should continue the request or ignore the response if the
//request is already finished.
const Continue = 100

// SwitchingProtocols - This code is sent in response to an Upgrade request header from the client and indicates the
//protocol the server is switching to.
const SwitchingProtocols = 101

// Processing (WebDAV) - This code indicates that the server has received and is processing the request, but no response
//is available yet.
const Processing = 102

// EarlyHints - This status code is primarily intended to be used with the Link header, letting the user agent start
//preloading resources while the server prepares a response.
const EarlyHints = 103

/**
Status codes and their use in gRPC
See https://grpc.github.io/grpc/core/md_doc_statuscodes.html
for details
*/

const GRPC_STATUS_OK = 0
const GRPC_STATUS_CANCELLED = 1
const GRPC_STATUS_UNKNOWN = 2
const GRPC_STATUS_INVALID_ARGUMENT = 3
const GRPC_STATUS_DEADLINE_EXCEEDED = 4
const GRPC_STATUS_NOT_FOUND = 5
const GRPC_STATUS_ALREADY_EXISTS = 6
const GRPC_STATUS_PERMISSION_DENIED = 7
const GRPC_STATUS_RESOURCE_EXHAUSTED = 8
const GRPC_STATUS_FAILED_PRECONDITION = 9
const GRPC_STATUS_ABORTED = 10
const GRPC_STATUS_OUT_OF_RANGE = 11
const GRPC_STATUS_UNIMPLEMENTED = 12
const GRPC_STATUS_INTERNAL = 13
const GRPC_STATUS_UNAVAILABLE = 14
const GRPC_STATUS_DATA_LOSS = 15
const GRPC_STATUS_UNAUTHENTICATED = 16
