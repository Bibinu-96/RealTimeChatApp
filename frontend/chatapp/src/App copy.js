// Import necessary modules
import React, { useState } from 'react';
import { BrowserRouter as Router, Routes, Route, Link } from 'react-router-dom';
import { AppBar, Toolbar, Typography, Button, TextField, Container, Box} from '@mui/material';
import './App.css';

function App() {
  return (
    <Router>
      <div className="App">
        <AppBar position="static">
          <Toolbar>
            <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
              ChatApp
            </Typography>
            <Button color="inherit" component={Link} to="/">Register</Button>
            {/* <Button color="inherit" component={Link} to="/register">Register</Button> */}
            <Button color="inherit" component={Link} to="/login">Login</Button>
          </Toolbar>
        </AppBar>
        <Routes>
          <Route path="/" element={<Register />} />
          {/* <Route path="/register" element={<Register />} /> */}
          <Route path="/login" element={<Login />} />
        </Routes>
      </div>
    </Router>
  );
}

function Home() {
  return (
    <Container className="full-height-container">
      <Box textAlign="center" mt={5}>
        <Typography variant="h3">Welcome to the App</Typography>
        <Typography variant="body1" mt={2}>Navigate to Register or Login to get started!</Typography>
      </Box>
    </Container>
  );
}

function Register() {
  const [formData, setFormData] = useState({
    username: '',
    email: '',
    password: '',
  });

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setFormData({ ...formData, [name]: value });
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    console.log('Registration Data:', formData);
    alert('Registration Successful!');
  };

  return (
    <Container className="full-height-container">
      <Box mt={5} p={3} boxShadow={3} borderRadius={2} textAlign="center" bgcolor="white">
        <Typography variant="h4" mb={2}>Register</Typography>
        <form onSubmit={handleSubmit}>
          <TextField
            fullWidth
            label="Username"
            name="username"
            value={formData.username}
            onChange={handleInputChange}
            margin="normal"
            required
          />
          <TextField
            fullWidth
            label="Email"
            name="email"
            value={formData.email}
            onChange={handleInputChange}
            margin="normal"
            required
          />
          <TextField
            fullWidth
            label="Password"
            name="password"
            type="password"
            value={formData.password}
            onChange={handleInputChange}
            margin="normal"
            required
          />
          <Button variant="contained" color="primary" type="submit" fullWidth>Register</Button>
        </form>
      </Box>
    </Container>
  );
}

function Login() {
  const [loginData, setLoginData] = useState({
    email: '',
    password: '',
  });

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setLoginData({ ...loginData, [name]: value });
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    console.log('Login Data:', loginData);
    alert('Login Successful!');
  };

  return (
    <Container className="full-height-container">
      <Box mt={5} p={3} boxShadow={3} borderRadius={2} textAlign="center" bgcolor="white">
        <Typography variant="h4" mb={2}>Login</Typography>
        <form onSubmit={handleSubmit}>
          <TextField
            fullWidth
            label="Email"
            name="email"
            value={loginData.email}
            onChange={handleInputChange}
            margin="normal"
            required
          />
          <TextField
            fullWidth
            label="Password"
            name="password"
            type="password"
            value={loginData.password}
            onChange={handleInputChange}
            margin="normal"
            required
          />
          <Button variant="contained" color="primary" type="submit" fullWidth>Login</Button>
        </form>
      </Box>
    </Container>
  );
}

export default App;
