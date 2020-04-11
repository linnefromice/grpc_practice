const PROTO_PATH = __dirname + '/employee.proto';

const grpc = require('grpc');
const protoLoader = require('@grpc/proto-loader');
const _ = require('lodash');

let packageDefinition = protoLoader.loadSync(
  PROTO_PATH,
  {
    keepCase: true,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true
  }
);
let employee_proto = grpc.loadPackageDefinition(packageDefinition).employee;

let employees = [
  {
    id: 1,
    email: "abcd@abcd.com",
    firstName: "First1",
    lastName: "Last1"
  },
  {
    id: 2,
    email: "xyz@xyz.com",
    firstName: "First2",
    lastName: "Last2"
  },
  {
    id: 3,
    email: "temp@temp.com",
    firstName: "First3",
    lastName: "Last3"
  },
];

function getDetails(call, callback) {
  callback(null, {
    message: _.find(employees, { id: call.request.id })
  });
}

function main() {
  let server = new grpc.Server();
  server.addService(employee_proto.Employee.service, {getDetails: getDetails});
  server.bind('0.0.0.0:4500', grpc.ServerCredentials.createInsecure());
  server.start();
}

main();