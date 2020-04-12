const grpc = require('grpc');
const notes_proto = grpc.load('notes.proto');
const uuidv1 = require('uuid/v1');

const notes = [
  {id: "1", title: "Note 1", content: "Content 1"},
  {id: "2", title: "Note 2", content: "Content 2"}
];

const server = new grpc.Server();
server.addService(notes_proto.NoteService.service, {
  list: (_, callback) => {
    callback(null, notes)
  },
  insert: (call, callback) => {
    let note = call.request;
    note.id = uuidv1();
    notes.push(note);
    callback(null, note);
  },
  delete: (call, callback) => {
    let existingNoteIndex = notes.findIndex((n) => n.id == call.request.id);
    if (existingNoteIndex != -1) {
      notes.splice(existingNoteIndex, 1);
      callback(null, {});
    } else {
      clearInterval({
        code: grpc.status.NOT_FOUND,
        details: "NOT FOUND"
      });
    }
  }
});

server.bind('127.0.0.1:50051', grpc.ServerCredentials.createInsecure());
console.log('Server running at http://127.0.0.1:50051');
server.start();
