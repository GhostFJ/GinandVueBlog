<template>
  <div class="page login-page">
    <div class="container d-flex align-items-center">
      <div class="form-holder has-shdow">
        <div class="row">
          <!-- log area-->
          <div class="col-lg-6">
            <div class="info d-flex align-items-center">
              <div class="logo">
                <h1>欢迎登录</h1>
                <p>自信人生二百年，会当击水三千里</p>
              </div>
            </div>
          </div>
          <!-- form area -->
          <div class="col-lg-6 bg-white">
            <div class="form d-flex align-items-center">
              <div class="content">
                <form action="" method="post" class="form-validate" id="login-form">
                  <div class="form-group">
                    <b-input
                      id="username"
                      name="username"
                      required
                      data-msg="请输入用户名"
                      class="input-material"
                      placeholder="请输入用户名"
                      v-model="$v.user.username.$model"
                      :state="validateState('username')">
                    </b-input>
                    <b-form-invalid-feedback :state="validateState('username')">
                      用户名长度不符合要求(2~20字符)
                    </b-form-invalid-feedback>
                  </div>

                  <div class="form-group">
                    <b-input
                    type="password" 
                    name="password" 
                    required data-msg="请输入密码" 
                    class="input-material" 
                    placeholder="请输入密码"
                    v-model="$v.user.password.$model"
                    :state="validateState('password')"
                    ></b-input>
                    <b-form-invalid-feedback :state="validateState('password')">
                      密码必须大于等于 6 位
                    </b-form-invalid-feedback>
                  </div>

                  <div class="form-group">
                    <b-form-group>
                      <b-input-group>
                      <b-col sm="4">
                        <b-img  @click="refreshCode" fluid thumbnail :src="captcha_url"></b-img>
                      </b-col>
                      <b-input type="text" 
                      name="captcha_code" 
                      required data-msg="请输入验证码" 
                      class="input-material" 
                      placeholder="请输入验证码"
                      v-model="user.captcha_code"></b-input>
                    </b-input-group>
                    </b-form-group>
                  </div>
                  <b-button class="btn btn-info" id="btnLogin" @click="login">登录</b-button>
                  <div style="margin-top:-20px;">
                    <div class="custom-control custom-checkbox" style="float:right;">
                      <b-form-checkbox id="chkremenber">记住密码</b-form-checkbox>
                    </div>
                  </div>
                </form>
                <p style="clear:both;">
                  <small>还没账户？<b-link href="/register">立即注册</b-link></small>
                </p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { required, minLength, maxLength } from 'vuelidate/lib/validators';
import { mapActions } from 'vuex';
import request from '@/utils/request'

export default {
  name: 'Login',
  data() {
    return {
        user: {
        username: '',
        password: '',
        captcha_id: '',
        captcha_code: '',
      },
      captcha_url: 'url',
    }
  },
  validations: {
    user: {
      username: {
        required,
        minLength: minLength(2),
        maxLength: maxLength(20),
      },
      password: {
        required,
        minLength: minLength(6),
      },
      captcha_code: {
        required
      }
    },
  },
  methods: {
    ...mapActions('userModule', { userlogin: 'login' }),
    validateState(name) {
      // 这里是es6 解构赋值
      const { $dirty, $error } = this.$v.user[name];
      return $dirty ? !$error : null;
    },
    login() {
      // 验证数据
      this.$v.user.$touch();
      if (this.$v.user.$anyError) {
        return;
      }
      // 请求
      this.userlogin(this.user).then(() => {
        // 跳转主页
        this.$router.replace({ name: 'Admin' });
      }).catch((err) => {
        this.$bvToast.toast(err.response.data.msg, {
          title: '数据验证错误',
          variant: 'danger',
          solid: true,
        });
      });
    },
    async refreshCode() {
      const res = await request.get('pub/captchaid')
      if(res.data.code != 200) {
        return this.$message.error("获取验证码失败")
      }
      this.user.captcha_id = res.data.data.captchaId
      this.captcha_url = "http://localhost:8000/api/v1/pub/captcha?captchaId=" + this.user.captcha_id
    }
  },
  created() {                   
    this.refreshCode()
  },
}
</script>

<style scoped>
@import '../../assets/css/login.css';
img{
  width: 100px;
  height: 40px !important;
}
</style>