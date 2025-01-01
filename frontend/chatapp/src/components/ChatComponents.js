import React, { useState } from 'react';
import {
  TextField,
  Button,
  Container,
  Paper,
  Typography,
  List,
  ListItem,
  ListItemText,
  ListItemAvatar,
  Avatar,
  Divider,
  Box,
  IconButton,
  AppBar,
  Toolbar,
  Card,
  CardContent,
} from '@mui/material';
import {
  Send as SendIcon,
  Person as PersonIcon,
  Group as GroupIcon,
} from '@mui/icons-material';

// SignUp Component
const SignUp = () => {
  const [formData, setFormData] = useState({
    username: '',
    email: '',
    password: ''
  });

  const handleSubmit = (e) => {
    e.preventDefault();
    console.log("formdata",formData)
    // Handle signup logic
  };

  return (
    <Container maxWidth="sm">
      <Box sx={{ mt: 8, display: 'flex', flexDirection: 'column', alignItems: 'center' }}>
        <Paper elevation={3} sx={{ p: 4, width: '100%' }}>
          <Typography component="h1" variant="h5" align="center" gutterBottom>
            Sign Up
          </Typography>
          <Box component="form" onSubmit={handleSubmit} sx={{ mt: 2 }}>
            <TextField
              margin="normal"
              required
              fullWidth
              label="Username"
              name="username"
              autoComplete="username"
              value={formData.username}
              onChange={(e) => setFormData({ ...formData, username: e.target.value })}
            />
            <TextField
              margin="normal"
              required
              fullWidth
              label="Email Address"
              name="email"
              autoComplete="email"
              value={formData.email}
              onChange={(e) => setFormData({ ...formData, email: e.target.value })}
            />
            <TextField
              margin="normal"
              required
              fullWidth
              label="Password"
              type="password"
              name="password"
              autoComplete="new-password"
              value={formData.password}
              onChange={(e) => setFormData({ ...formData, password: e.target.value })}
            />
            <Button
              type="submit"
              fullWidth
              variant="contained"
              sx={{ mt: 3, mb: 2 }}
            >
              Register
            </Button>
          </Box>
        </Paper>
      </Box>
    </Container>
  );
};

// Login Component
const Login = () => {
  const [formData, setFormData] = useState({
    email: '',
    password: ''
  });

  const handleSubmit = (e) => {
    e.preventDefault();
    console.log("formdata",formData)
    // Handle login logic
  };

  return (
    <Container maxWidth="sm">
      <Box sx={{ mt: 8, display: 'flex', flexDirection: 'column', alignItems: 'center' }}>
        <Paper elevation={3} sx={{ p: 4, width: '100%' }}>
          <Typography component="h1" variant="h5" align="center" gutterBottom>
            Login
          </Typography>
          <Box component="form" onSubmit={handleSubmit} sx={{ mt: 2 }}>
            <TextField
              margin="normal"
              required
              fullWidth
              label="Email Address"
              name="email"
              autoComplete="email"
              value={formData.email}
              onChange={(e) => setFormData({ ...formData, email: e.target.value })}
            />
            <TextField
              margin="normal"
              required
              fullWidth
              label="Password"
              type="password"
              name="password"
              autoComplete="current-password"
              value={formData.password}
              onChange={(e) => setFormData({ ...formData, password: e.target.value })}
            />
            <Button
              type="submit"
              fullWidth
              variant="contained"
              sx={{ mt: 3, mb: 2 }}
            >
              Login
            </Button>
          </Box>
        </Paper>
      </Box>
    </Container>
  );
};

// ChatList Component
const ChatList = () => {
  const chats = [
    { id: 1, name: 'John Doe', lastMessage: 'Hey, how are you?', time: '10:30 AM' },
    { id: 2, name: 'Jane Smith', lastMessage: 'See you tomorrow!', time: '9:45 AM' },
  ];

  return (
    <Paper sx={{ maxWidth: 600, mx: 'auto' }}>
      <AppBar position="static" color="default">
        <Toolbar>
          <Typography variant="h6">Chats</Typography>
        </Toolbar>
      </AppBar>
      <List sx={{ bgcolor: 'background.paper' }}>
        {chats.map((chat) => (
          <React.Fragment key={chat.id}>
            <ListItem button>
              <ListItemAvatar>
                <Avatar>
                  <PersonIcon />
                </Avatar>
              </ListItemAvatar>
              <ListItemText
                primary={chat.name}
                secondary={
                  <React.Fragment>
                    <Typography
                      component="span"
                      variant="body2"
                      color="text.primary"
                    >
                      {chat.lastMessage}
                    </Typography>
                    {` — ${chat.time}`}
                  </React.Fragment>
                }
              />
            </ListItem>
            <Divider variant="inset" component="li" />
          </React.Fragment>
        ))}
      </List>
    </Paper>
  );
};

// GroupChatList Component
const GroupChatList = () => {
  const groups = [
    { id: 1, name: 'Project Team', lastMessage: 'Meeting at 2 PM', time: '11:00 AM', members: 5 },
    { id: 2, name: 'Family Group', lastMessage: 'Dinner plans?', time: '10:15 AM', members: 8 },
  ];

  return (
    <Paper sx={{ maxWidth: 600, mx: 'auto' }}>
      <AppBar position="static" color="default">
        <Toolbar>
          <Typography variant="h6">Group Chats</Typography>
        </Toolbar>
      </AppBar>
      <List sx={{ bgcolor: 'background.paper' }}>
        {groups.map((group) => (
          <React.Fragment key={group.id}>
            <ListItem button>
              <ListItemAvatar>
                <Avatar>
                  <GroupIcon />
                </Avatar>
              </ListItemAvatar>
              <ListItemText
                primary={group.name}
                secondary={
                  <React.Fragment>
                    <Typography
                      component="span"
                      variant="body2"
                      color="text.primary"
                    >
                      {group.lastMessage}
                    </Typography>
                    {` — ${group.time} • ${group.members} members`}
                  </React.Fragment>
                }
              />
            </ListItem>
            <Divider variant="inset" component="li" />
          </React.Fragment>
        ))}
      </List>
    </Paper>
  );
};

// Individual Chat Component
const ChatRoom = () => {
  const [message, setMessage] = useState('');
  const messages = [
    { id: 1, sender: 'John', content: 'Hey!', time: '10:30 AM', isSentByMe: false },
    { id: 2, sender: 'Me', content: 'Hi John!', time: '10:31 AM', isSentByMe: true },
  ];

  const handleSend = (e) => {
    e.preventDefault();
    // Handle send message logic
    setMessage('');
  };

  return (
    <Box sx={{ height: '100vh', display: 'flex', flexDirection: 'column' }}>
      <AppBar position="static">
        <Toolbar>
          <Typography variant="h6">John Doe</Typography>
        </Toolbar>
      </AppBar>

      <Box sx={{ flexGrow: 1, overflow: 'auto', p: 2, backgroundColor: 'grey.100' }}>
        {messages.map((msg) => (
          <Box
            key={msg.id}
            sx={{
              display: 'flex',
              justifyContent: msg.isSentByMe ? 'flex-end' : 'flex-start',
              mb: 2
            }}
          >
            <Card sx={{ maxWidth: '70%', bgcolor: msg.isSentByMe ? 'primary.main' : 'background.paper' }}>
              <CardContent>
                <Typography variant="body1" sx={{ color: msg.isSentByMe ? 'white' : 'text.primary' }}>
                  {msg.content}
                </Typography>
                <Typography variant="caption" sx={{ color: msg.isSentByMe ? 'white' : 'text.secondary' }}>
                  {msg.time}
                </Typography>
              </CardContent>
            </Card>
          </Box>
        ))}
      </Box>

      <Paper sx={{ p: 2, borderTop: 1, borderColor: 'divider' }}>
        <Box component="form" onSubmit={handleSend} sx={{ display: 'flex', gap: 1 }}>
          <TextField
            fullWidth
            size="small"
            placeholder="Type a message..."
            value={message}
            onChange={(e) => setMessage(e.target.value)}
          />
          <IconButton color="primary" type="submit">
            <SendIcon />
          </IconButton>
        </Box>
      </Paper>
    </Box>
  );
};

// Group Chat Room Component
const GroupChatRoom = () => {
  const [message, setMessage] = useState('');
  const messages = [
    { id: 1, sender: 'John', content: 'Hello everyone!', time: '10:30 AM', isSentByMe: false },
    { id: 2, sender: 'Me', content: 'Hi team!', time: '10:31 AM', isSentByMe: true },
    { id: 3, sender: 'Alice', content: 'When is the meeting?', time: '10:32 AM', isSentByMe: false },
  ];

  const handleSend = (e) => {
    e.preventDefault();
    // Handle send message logic
    setMessage('');
  };

  return (
    <Box sx={{ height: '100vh', display: 'flex', flexDirection: 'column' }}>
      <AppBar position="static">
        <Toolbar>
          <Box>
            <Typography variant="h6">Project Team</Typography>
            <Typography variant="caption">8 members</Typography>
          </Box>
        </Toolbar>
      </AppBar>

      <Box sx={{ flexGrow: 1, overflow: 'auto', p: 2, backgroundColor: 'grey.100' }}>
        {messages.map((msg) => (
          <Box
            key={msg.id}
            sx={{
              display: 'flex',
              justifyContent: msg.isSentByMe ? 'flex-end' : 'flex-start',
              mb: 2
            }}
          >
            <Card sx={{ maxWidth: '70%', bgcolor: msg.isSentByMe ? 'primary.main' : 'background.paper' }}>
              <CardContent>
                {!msg.isSentByMe && (
                  <Typography variant="subtitle2" color="text.secondary">
                    {msg.sender}
                  </Typography>
                )}
                <Typography variant="body1" sx={{ color: msg.isSentByMe ? 'white' : 'text.primary' }}>
                  {msg.content}
                </Typography>
                <Typography variant="caption" sx={{ color: msg.isSentByMe ? 'white' : 'text.secondary' }}>
                  {msg.time}
                </Typography>
              </CardContent>
            </Card>
          </Box>
        ))}
      </Box>

      <Paper sx={{ p: 2, borderTop: 1, borderColor: 'divider' }}>
        <Box component="form" onSubmit={handleSend} sx={{ display: 'flex', gap: 1 }}>
          <TextField
            fullWidth
            size="small"
            placeholder="Type a message..."
            value={message}
            onChange={(e) => setMessage(e.target.value)}
          />
          <IconButton color="primary" type="submit">
            <SendIcon />
          </IconButton>
        </Box>
      </Paper>
    </Box>
  );
};

export { SignUp, Login, ChatList, GroupChatList, ChatRoom, GroupChatRoom };
