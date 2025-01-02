import React, { useState } from 'react';
import {
    TextField,
    Button,
    Container,
    Paper,
    Typography,
    Box,
  } from '@mui/material';
  import httpService from "../service/HttpService";
  import AlertDialog from './AlertDialog';
  import { useSnackbar } from 'notistack';
  import { useNavigate } from 'react-router-dom';
  import {isValidEmail} from "../utils/util"
  let alertContent=""
  // Login Component
  const Login = (props) => {
    const [formData, setFormData] = useState({
      email: '',
      password: ''
    });
     const [showErrorDialog ,setErrorDialogStatus]= useState(false)
      const { enqueueSnackbar } = useSnackbar();
      const navigate = useNavigate();
    
    
      const showSnackbar = (msg,variant) => {
        enqueueSnackbar(msg, { variant });
      };
  
    const handleSubmit = async (e) => {
      e.preventDefault();
      console.log("formdata",formData)
      // Handle login logic
      if(! isValidEmail(formData.email)){
              alertContent="Invalid EmailId"
              setErrorDialogStatus(true)
              return
          }

          const response = await httpService.post("login", 
            formData);
          
            if (response.status === 200) {
              console.log("Token:", response.data);
              showSnackbar('User logged in successfully!', 'success');
              localStorage.setItem('authToken',response.data?.token)
              try{
                props.onLogin();
                navigate('/chats');
               
              }catch (error) {
                console.error('Navigation error:', error);
              } 
            } else {
              console.error("Error login:", response);
              alertContent=response.error?.error
              setErrorDialogStatus(true)
            }
    };
    const handleClose=()=>{
        setErrorDialogStatus(false)
      }
      const goToSignup=()=>{
        navigate('/signup'); 
      }
      const alertDialog = showErrorDialog?<AlertDialog handleClose={handleClose} content={alertContent}/>:null
  
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
            <Button
                onClick={goToSignup}
                fullWidth
                variant="contained"
                color="error"
                sx={{ mt: 3, mb: 2 }}
              >
                SignUp
              </Button>
            {alertDialog}
          </Paper>
        </Box>
      </Container>
    );
  };
  export default Login