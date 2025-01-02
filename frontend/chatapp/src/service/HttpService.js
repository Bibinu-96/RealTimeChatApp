import axios from "axios";
import NProgress from "nprogress"; // Import NProgress
import "./nprogress.css"; // Import NProgress CSS

// Create an Axios instance
const axiosInstance = axios.create({
  baseURL: "http://localhost:8000/api/", // Replace with your API base URL
  headers: {
    "Content-Type": "application/json",
  },
});

// Add request interceptor to show progress bar
axiosInstance.interceptors.request.use(
  (config) => {
     // Retrieve the token
      const token = localStorage.getItem('authToken'); // or sessionStorage.getItem('authToken')
      if (token) {
        config.headers.Authorization = `Bearer ${token}`;
      }
    NProgress.start(); // Start progress bar on request
    return config;
  },
  (error) => {
    NProgress.done(); // Stop progress bar if request fails
    return Promise.reject(error);
  }
);

// Add response interceptor to hide progress bar
axiosInstance.interceptors.response.use(
  (response) => {
    NProgress.done(); // End progress bar when response is received
    return response;
  },
  (error) => {
    NProgress.done(); // End progress bar when response fails
    return Promise.reject(error);
  }
);

// HTTP Service
const httpService = {
  // GET method
  get: async (url, config = {}) => {
    try {
      const response = await axiosInstance.get(url, config);
      return { status: response.status, data: response.data };
    } catch (error) {
      return {
        status: error.response?.status || 500,
        error: error.response?.data || error.message,
      };
    }
  },

  // POST method
  post: async (url, data, config = {}) => {
    try {
      const response = await axiosInstance.post(url, data, config);
      return { status: response.status, data: response.data };
    } catch (error) {
      return {
        status: error.response?.status || 500,
        error: error.response?.data || error.message,
      };
    }
  },

  // DELETE method
  delete: async (url, config = {}) => {
    try {
      const response = await axiosInstance.delete(url, config);
      return { status: response.status, data: response.data };
    } catch (error) {
      return {
        status: error.response?.status || 500,
        error: error.response?.data || error.message,
      };
    }
  },
};

export default httpService;
