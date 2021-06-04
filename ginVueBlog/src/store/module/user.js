import storageService from '@/service/storageService';
import userService from '@/service/userService';

const userModule = {
  namespaced: true,
  state: {
    token: storageService.get(storageService.USER_TOKEN),
    userInfo: storageService.get(storageService.USER_INFO) ? storageService.get(storageService.USER_INFO) : null, //eslint-disable-line
  },
  mutations: {
    SET_TOKEN(state, token) {
      // 更新本地缓存
      storageService.set(storageService.USER_TOKEN, token);
      // 更新state
      state.token = token;
    },

    SET_USERINFO(state, userInfo) {
      // 更新本地缓存
      storageService.set(storageService.USER_INFO, userInfo);
      // 更新state
      state.userInfo = userInfo;
    },
  },

  actions: {
    register(context, { username, password, password_again, captcha_id, captcha_code  }) {
      return new Promise((resolve, reject) => {
        // 请求
        userService.register({ username, password, password_again, captcha_id, captcha_code })
        .then(() => {
          // 登录，并获取Token
          return userService.login({ username, password, captcha_id, captcha_code })
        })
        .then((res) => {
          // 保存token
          context.commit('SET_TOKEN', res.data.data.token);
          return userService.info();
        })
        .then((res) => {
          // 保存用户信息
          context.commit('SET_USERINFO', res.data.data.username);
          resolve(res);
        }).catch((err) => {
          reject(err);
        });
      });
    },

    login(context, { username, password, captcha_id, captcha_code }) {
      return new Promise((resolve, reject) => {
        // 请求
        userService.login({ username, password, captcha_id, captcha_code }).then((res) => {
          // 保存token
          context.commit('SET_TOKEN', res.data.data.token);
          return userService.info();
        }).then((res) => {
          // 保存用户信息
          context.commit('SET_USERINFO', res.data.data.data.username);
          resolve(res);
        }).catch((err) => {
          reject(err);
        });
      });
    },

    // 前台登出，后台登出会删除token，另外实现
    logout({ commit }) {
      // 清除token
      commit('SET_TOKEN', '');
      storageService.set(storageService.USER_TOKEN, '');
      // 清除用户信息
      commit('SET_USERINFO', '');
      storageService.set(storageService.SET_USERINFO, '');

      window.location.reload();
    },
  },
};

export default userModule;
