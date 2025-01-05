import { io } from "socket.io-client";

const socket = io("http://localhost:3000");

// Join the server with a userId
const userId = "user123"; // Replace with a dynamic user ID
socket.emit("join", { userId });

// Listen for incoming messages
socket.on("message", (data) => {
  console.log("New message:", data);
});

// Send a private message
const sendPrivateMessage = (toUserId, message) => {
  socket.emit("privateMessage", { toUserId, message });
};

// Join a group
const joinGroup = (groupName) => {
  socket.emit("joinGroup", { groupName });
};

// Send a group message
const sendGroupMessage = (groupName, message) => {
  socket.emit("groupMessage", { groupName, message });
};
