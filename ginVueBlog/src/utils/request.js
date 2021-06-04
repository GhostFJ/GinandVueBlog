import axios from 'axios';
import storageService from '../service/storageService';

// 实例初始化后不会动态随着token变化，需要处理，用到axios的interceptors
const service = axios.create({
  baseURL: "http://127.0.0.1:8000/api/v1/",
  timeout: 1000 * 5,
  // headers: { Authorization: `Bearer ${storageService.get(storageService.USER_TOKEN)}` },
});

// Add a request interceptor
service.interceptors.request.use((config) => {
  // 修改headers
  Object.assign(config.headers, { Authorization: `${storageService.get(storageService.USER_TOKEN)}` });
  return config;
}, (error) => {
  // Do something with request error
  return Promise.reject(error);
});

export default service;
