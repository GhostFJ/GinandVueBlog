import request from '@/utils/request';

// 用户注册
const register = ({ username, password, password_again, captcha_id, captcha_code }) => {
  return request.post('reg', { username, password, password_again, captcha_id, captcha_code });
};

// 用户登录
const login = ({ username, password, captcha_id, captcha_code }) => {
  return request.post('auth', { username, password, captcha_id, captcha_code });
};

// 获取用户信息
const info = () => {
  return request.get('currentuser');
};

export default {
  register,
  login,
  info,
};
