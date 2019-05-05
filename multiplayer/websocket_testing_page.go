package main

import (
	"github.com/gin-gonic/gin"
)

func WebSocketTestPage(c *gin.Context) {
	// language=HTML // Enable "IntelliLang" plugin if syntax highlighting did not work.
	html := []byte(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>WebSocket test</title>
    <link rel="icon" href="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAABHNCSVQICAgIfAhkiAAAAAtJREFUCJljYAACAAAFAAFiVTKIAAAAAElFTkSuQmCC">
    <style>
html {
    min-height: 100%;
}
body {
    margin: 0;
    min-height: 100%;
    display: flex;
    flex-grow: 1;
    justify-content: center;
    align-items: center;
    background-color: #2e3436;
    font-family: "Arial", sans-serif;
    color: #babdb6;
    font-size: 1rem;
}

main {
    width: 40rem;
    display: flex;
    flex-direction: column;
}
button:hover {
    color: #ffffff;
}
.main-button {
    display: flex;
    cursor: pointer;
    margin: 0.2rem 0 0.2rem 0;
    background-color: #272c2d;
    border-radius: 0.5rem;
    color: inherit;
    border: none;
    font-size: inherit;
    padding: 0.5rem;
}
button::-moz-focus-inner {
    border: 0;
}
.dispatch-container {
    position: relative;
}
.hanging-button {
    position: absolute;
    top: 0.5rem;
    right: 1rem;
    background-color: #2e3436;
    font-size: inherit;
    border: none;
    border-radius: 0.5rem;
    cursor: pointer;
    padding: 0.2rem;
	color: inherit;
}
textarea {
    box-sizing: border-box;
    width: 100%;
    min-height: 5rem;
    resize: vertical;
    font-family: "DejaVu Sans Mono", monospace;
}
.decorated-text {
    margin: 0.2rem 0 0.2rem 0;
    background-color: #272c2d;
    border-radius: 0.5rem 0.5rem 0 0.5rem;
    color: inherit;
    border: none;
    font-size: inherit;
    padding: 0.5rem;
}
#incoming-data {
    min-height: 40rem;
}
    </style>
</head>
<body>
<main>
    <button id="connect" class="main-button">Connect</button>
    <button id="disconnect" class="main-button">Disconnect</button>
    <div class="dispatch-container">
        <textarea id="sent-data-0" class="decorated-text" placeholder="Отправляемые через websocket данные."></textarea>
        <button id="send-0" class="hanging-button">Send</button>
    </div>
    <textarea id="incoming-data" class="decorated-text" placeholder="Присылаемые сервером данные."></textarea>
</main>
<script defer type="application/javascript">
"use strict";

let socket;

document.querySelector("#connect").addEventListener("click", (event) => {
    // очистка поля вывода 
    document.querySelector("#incoming-data").value = "";
    // устанавливаем cookie со случайной строкой.
    let d = new Date();
    d.setDate(d.getDate()+1);
    document.cookie = 
        "session_id="+Math.round(Math.random()*2**32).toString()+"; "+
        "path=/; "+
        "expires="+d.toUTCString()+";"; 

    socket = new WebSocket("ws://"+location.host+"/game/start");

    socket.addEventListener("open", (event) => {
        document.querySelector("#incoming-data").value += "// Open socket\n";
    });

    socket.addEventListener("close", (event) => {
        document.querySelector("#incoming-data").value += 
        "// Close socket " + (event.wasClean ? "clean" : "suddenly") + ". " + 
        "Code: " + event.code + " cause: " + event.reason + "\n";
    });

    socket.addEventListener("message", (event) => {
        document.querySelector("#incoming-data").value += 
        event.data + "\n";
    });

    socket.addEventListener("error", (error) => {
        document.querySelector("#incoming-data").value += 
        "// Error: " + error.message + "\n";
    });
});

document.querySelector("#disconnect").addEventListener("click", (event) => {
    socket.close();
});

for (const i of ["0", "1", "2", "3", "4"]){
    document.querySelector("#send-"+i).addEventListener("click", (event) => {
        socket.send(document.querySelector("#sent-data-"+i).value);
    });
}
</script>
</body>
</html>
    `)
	c.Writer.Header().Add("Content-Type", "text/html;charset=UTF-8")
	c.Status(200)
	c.Writer.Write(html)
	return
}
