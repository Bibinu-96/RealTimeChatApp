//SignUp Component
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
  import {validatePhoneNumber,isValidEmail} from "../utils/util"
  let alertContent=""
const SignUp = () => {
  const [formData, setFormData] = useState({
    username: '',
    email: '',
    password: '',
    phoneno:''
  });
  const [showErrorDialog ,setErrorDialogStatus]= useState(false)
  const { enqueueSnackbar } = useSnackbar();
  const navigate = useNavigate();


  const showSnackbar = (msg,variant) => {
    enqueueSnackbar(msg, { variant });
  };
  const handleSubmit = async (e) => {
    e.preventDefault();
   
    // Handle signup logic
    console.log("formdata ",formData)
    if(! isValidEmail(formData.email)){
        alertContent="Invalid EmailId"
        setErrorDialogStatus(true)
        return
    }
    if( !validatePhoneNumber(formData.phoneno)){
        alertContent="Invalid Phone no"
        setErrorDialogStatus(true)
        return

    }

    const response = await httpService.post("register", 
      formData);
    
      if (response.status === 201) {
        console.log("User created:", response.data);
        showSnackbar('User Created successfully!', 'success');
        navigate('/login'); 
      } else {
        console.error("Error creating user:", response);
        alertContent=response.error?.error
        setErrorDialogStatus(true)
      }

  };
  const handleClose=()=>{
    setErrorDialogStatus(false)
  }
  const alertDialog = showErrorDialog?<AlertDialog handleClose={handleClose} content={alertContent}/>:null
const goToLogin=()=>{
  navigate('/login'); 
}
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
            <TextField
              margin="normal"
              required
              fullWidth
              label="PhoneNo"
              type="number"
              name="phoneno"
              autoComplete="Phone-no"
              value={formData.phoneno}
              onChange={(e) => setFormData({ ...formData, phoneno: e.target.value })}
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
           <Button
            onClick={goToLogin}
            fullWidth
            variant="contained"
            color="error"
            sx={{ mt: 3, mb: 2 }}
            >
            Login
            </Button>
          {alertDialog}
        </Paper>

      </Box>
    </Container>
  );
};
export default SignUp