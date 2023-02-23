import {
  connect,
  StringCodec,
} from "https://cdn.jsdelivr.net/npm/nats.ws@1.13.0/esm/nats.js";

// https://nats-io.github.io/nats.deno/interfaces/NatsConnection.html
//

console.log(`connecting`);
const sc = new StringCodec();
const nc = await connect({ servers: "ws://127.0.0.1:8443" });
console.log(`connected to ${nc.getServer()}`);
// this promise indicates the client closed
const done = nc.closed();
// do something with the connection
nc.publish("hello", sc.encode("hi"));
let ms = await nc.rtt();
console.log(`${ms} ms`);
//nc.drain();

let jsm = await nc.jetstreamManager();

for await (const si of jsm.streams.list()) {
  console.log(si);
}
jsm.streams.add({
  name: "currentGame",
  description: "Announce of games that are opening",
  subjects: ["currentGame"],
  retention: "limits",
  max_consumers: -1,
  max_msgs: -1,
  max_bytes: -1,
  max_age: 900000000000,
  max_msgs_per_subject: 1,
  max_msg_size: -1,
  discard: "old",
  storage: "memory",
  num_replicas: 1,
  duplicate_window: 120000000000,
  allow_direct: false,
  mirror_direct: false,
  sealed: false,
  deny_delete: false,
  deny_purge: false,
  allow_rollup_hdrs: false,
});
jsm.streams.add({
  name: "gamesData",
  description: "All things related to the games",
  subjects: ["game", "game.>"],
  retention: "limits",
  max_consumers: -1,
  max_msgs: 10000,
  max_bytes: -1,
  max_age: 10800000000000,
  max_msgs_per_subject: 1000,
  max_msg_size: -1,
  discard: "old",
  storage: "memory",
  num_replicas: 1,
  duplicate_window: 120000000000,
  allow_direct: false,
  mirror_direct: false,
  sealed: false,
  deny_delete: false,
  deny_purge: false,
  allow_rollup_hdrs: false,
});

let js = await nc.jetstream();

window.publish = function (subject, data) {
  js.publish(subject, sc.encode(data))
    .then((pa) => {
      console.debug(pa.stream, pa.seq);
    })
    .catch((err) => {
      console.warn(
        "publish sub:'%s', data:'%s' failed: %s",
        subject,
        data,
        err
      );
    });
};

window.streamHasMsgs = async function (streamname) {
  jsm.streams.info(streamname).then((si) => {
    StreamHasMessageCB(si.state.messages > 0);
  });
};

window.streamLastMsg = async function (streamname) {
  try {
    let si = await jsm.streams.info(streamname);
  } catch {
    StreamLastMsgCB("");
  }
  if (si.state.last_seq === 0) {
    StreamLastMsgCB("");
  }
  jsm.streams
    .getMessage(streamname, { seq: si.state.last_seq })
    .then((msg) => {
      StreamLastMsgCB(sc.decode(msg.data));
    })
    .catch((err) => {
      console.error(err);
    });
};

window.subscribe= function(subject,num){
    let sub = js.subscribe(subject)
    (async ()=>{
        for await (const msg of sub){
            SubCB(msg.subject,sc.decode(msg.data),num)
        }
    })()
}
