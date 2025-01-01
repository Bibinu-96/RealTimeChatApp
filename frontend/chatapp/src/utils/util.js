export const validatePhoneNumber = (phoneNumber) => {
    // Check if the input is a valid number (digits only) and has exactly 10 digits
    const regex = /^\d{10}$/;
    return regex.test(phoneNumber); // Returns true if valid, false otherwise
  };

  export const isValidEmail = (email) => {
    const regex = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;
    return regex.test(email);
  };