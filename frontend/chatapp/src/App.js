// App.js
import React from 'react';
import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom';
import { ThemeProvider, createTheme, CssBaseline, Box, IconButton, Typography } from '@mui/material';
import {
  ArrowBack as ArrowBackIcon,
  Home as HomeIcon
} from '@mui/icons-material';
import {ChatList, GroupChatList, ChatRoom, GroupChatRoom } from './components/ChatComponents';
import SignUp from './components/Signup';
import Login from './components/Login';
import { useState,useEffect } from 'react';
import { useNavigate, useLocation } from 'react-router-dom';
import {jwtDecode} from 'jwt-decode';
// Navigation component with back button and title
const NavigationHeader = ({ title }) => {
  const navigate = useNavigate();
  const location = useLocation();
  
  return location.pathname !== '/chats' && location.pathname !== '/login' && location.pathname !== '/signup' ? (
    <Box sx={{ 
      position: 'fixed', 
      top: 0, 
      left: 0, 
      zIndex: 1000,
      padding: 1,
      display: 'flex',
      alignItems: 'center',
      gap: 1
    }}>
      <IconButton onClick={() => navigate(-1)} color="primary">
        <ArrowBackIcon />
      </IconButton>
      <Typography variant="h6">{title}</Typography>
    </Box>
  ) : null;
};

// Main Navigation component
const MainNavigation = () => {
  const navigate = useNavigate();
  const location = useLocation();

  // Only show on main screens
  if (location.pathname === '/chats' || location.pathname === '/groups') {
    return (
      <Box sx={{
        position: 'fixed',
        bottom: 0,
        left: 0,
        right: 0,
        display: 'flex',
        justifyContent: 'center',
        padding: 2,
        backgroundColor: 'background.paper',
        borderTop: 1,
        borderColor: 'divider'
      }}>
        <IconButton onClick={() => navigate('/chats')} color={location.pathname === '/chats' ? 'primary' : 'default'}>
          <HomeIcon />
        </IconButton>
      </Box>
    );
  }
  return null;
};

// Create theme
const theme = createTheme({
  palette: {
    primary: {
      main: '#1976d2',
    },
    secondary: {
      main: '#dc004e',
    },
  },
});
const App = () => {
  // Simple auth state management (replace with proper auth system)
 // checkAuth()
  const [isAuthenticated, setIsAuthenticated] = useState(() => {
    const token = localStorage.getItem('authToken');
    if (token) {
      try {
        const decodedToken = jwtDecode(token);
        const currentTime = Date.now() / 1000;
        return decodedToken.exp > currentTime; // Check if token is valid
      } catch (error) {
        console.error('Invalid token:', error);
        localStorage.removeItem('authToken');
      }
    }
    return false;
  });


  useEffect(() => {
    const token = localStorage.getItem('authToken'); // or sessionStorage.getItem('authToken')
    
    if (token) {
      try {
        const decodedToken = jwtDecode(token); // Decode the JWT
        const currentTime = Date.now() / 1000; // Current time in seconds

        if (decodedToken.exp > currentTime) {
          // Token is valid
          setIsAuthenticated(true);
        } else {
          // Token is expired
          localStorage.removeItem('authToken'); // Clear expired token
          setIsAuthenticated(false);
        }
      } catch (error) {
        console.error("Invalid token:", error);
        localStorage.removeItem('authToken'); // Clear invalid token
        setIsAuthenticated(false);
      }
    } else {
      setIsAuthenticated(false);
    }
  }, []);

  // Protected Route component
  const ProtectedRoute = ({ children }) => {
    return isAuthenticated ? children : <Navigate to="/login" />;
  };
  console.log("isAuthenticatedByServer ",isAuthenticated)
  return (
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <Router>
        <Box sx={{ height: '100vh', bgcolor: 'background.default' }}>
          <Routes>
            {/* Public Routes */}
            <Route path="/login" element={<Login onLogin={() => setIsAuthenticated(true)} />} />
            <Route path="/signup" element={<SignUp onSignup={() => setIsAuthenticated(true)} />} />

            {/* Protected Routes */}
            <Route
              path="/chats"
              element={
                <ProtectedRoute>
                  <ChatList />
                </ProtectedRoute>
              }
            />
            <Route
              path="/groups"
              element={
                <ProtectedRoute>
                  <GroupChatList />
                </ProtectedRoute>
              }
            />
            <Route
              path="/chat/:id"
              element={
                <ProtectedRoute>
                  <ChatRoom />
                </ProtectedRoute>
              }
            />
            <Route
              path="/group/:id"
              element={
                <ProtectedRoute>
                  <GroupChatRoom />
                </ProtectedRoute>
              }
            />

            {/* Redirect root to chats if authenticated, otherwise to login */}
            <Route
              path="/"
              element={
                isAuthenticated ? (
                  <Navigate to="/chats" replace />
                ) : (
                  <Navigate to="/login" replace />
                )
              }
            />
          </Routes>

          {/* Global Navigation Components */}
          <NavigationHeader title="Chat App" />
          <MainNavigation />
        </Box>
      </Router>
    </ThemeProvider>
  );
};

export default App;
