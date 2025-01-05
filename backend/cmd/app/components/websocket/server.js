const express = require('express');
const http = require('http');
const { Server } = require('socket.io');

const app = express();
const server = http.createServer(app);
const io = new Server(server);

// Store connected users (socket.id mapped to userId)
const users = {};

io.on('connection', (socket) => {
  console.log('A user connected:', socket.id);

  // Event: User joins with their userId
  socket.on('join', ({ userId }) => {
    users[userId] = socket.id;
    console.log(`${userId} has joined with socket ID ${socket.id}`);
  });

  // Event: Peer-to-peer message
  socket.on('privateMessage', ({ toUserId, message }) => {
    const toSocketId = users[toUserId];
    if (toSocketId) {
      io.to(toSocketId).emit('message', {
        from: socket.id,
        message,
        type: 'private',
      });
    } else {
      console.log(`User ${toUserId} is not connected.`);
    }
  });

  // Event: User joins a group (room)
  socket.on('joinGroup', ({ groupName }) => {
    socket.join(groupName);
    console.log(`Socket ID ${socket.id} joined group: ${groupName}`);
    io.to(groupName).emit('message', {
      from: 'Server',
      message: `A new user has joined the group: ${groupName}`,
      type: 'group',
    });
  });

  // Event: Group message
  socket.on('groupMessage', ({ groupName, message }) => {
    io.to(groupName).emit('message', {
      from: socket.id,
      message,
      type: 'group',
    });
  });

  // Event: User disconnects
  socket.on('disconnect', () => {
    console.log('A user disconnected:', socket.id);
    // Remove the user from the users list
    for (const userId in users) {
      if (users[userId] === socket.id) {
        delete users[userId];
        break;
      }
    }
  });
});

// Start the server
server.listen(3000, () => {
  console.log('Socket.IO server running on port 3000');
});
