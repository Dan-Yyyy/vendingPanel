import * as axios from 'axios';

const instanse = axios.create({
  withCredentials: false,
  baseURL: 'http://192.168.0.122:8000/',
});

export default instanse;