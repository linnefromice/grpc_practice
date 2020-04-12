const grpc = require('grpc');
const notes_proto = grpc.load('notes.proto');

const server = new grpc.Server();

server.bind('127.0.0.1:50051', grpc.ServerCredentials.createInsecure());
console.log('Server running at http://127.0.0.1:50051');
server.start();
