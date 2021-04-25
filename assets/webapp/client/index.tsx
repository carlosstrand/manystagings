import axios from 'axios';

export const createFreshHeaders = () => {
  const headers = {
    'Content-Type': 'application/json',
  };
  const token = localStorage.getItem("token");
  if (token) {
    headers['Authorization'] = `Bearer ${token}`;
  }
  return headers;
};

export default axios.create({
  baseURL: '',
  timeout: 1000,
  headers: createFreshHeaders(),
});
