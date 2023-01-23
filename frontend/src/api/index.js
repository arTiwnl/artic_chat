let socket = new WebSocket("ws://localhost:8000/ws");

let connect = cb => {
     console.log("tentando conectar...");

     socket.onopen = () => {
          console.log("estamos conectados")
     }

     socket.onmessage = msg => {
          console.log(msg);
          cb(msg);
     }

     socket.onclose = event => {
          console.log("conexao falhou:", event);
     }
     socket.onerror = error => {
          console.log("erro na lib:", error)
     }
};

let sendMsg = msg => {
     console.log("enviando msg", msg);
     socket.send(msg)
}

export { connect, sendMsg };