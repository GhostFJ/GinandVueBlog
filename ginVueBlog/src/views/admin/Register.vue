<template>
  <div class="register">
     <b-nav>
      <b-nav-item active to="/login">
        <b-icon icon="reply-fill"></b-icon>
        返回登录
      </b-nav-item>
    </b-nav>
    <b-row class="mt-5">
      <b-col
        md="8"
        offset-md="2"
        lg="6"
        offset-lg="3"
      >
        <b-card title="注册" header="用户注册" border-variant="primary" header-bg-variant="primary" header-text-variant="white" align="center">
          <b-form>
            <b-form-group label="用户名" description="Let us know your name.">
              <b-input-group>
                <b-input-group-prepend is-text>
                  <b-icon icon="person-fill"></b-icon>
                </b-input-group-prepend>
                <b-form-input
                  v-model="$v.user.name.$model"
                  type="text"
                  placeholder="请输入您的名称"
                  :state="validateState('name')"
                ></b-form-input>
              </b-input-group>
              <b-form-invalid-feedback :state="validateState('name')">
                用户名长度不符合要求(2~20字符)
              </b-form-invalid-feedback>
            </b-form-group>

            <b-form-group label="密码" description="Please set your password.">
              <b-input-group>
                <b-input-group-prepend is-text>
                  <b-icon icon="lock-fill"></b-icon>
                </b-input-group-prepend>
                <b-form-input
                  v-model="$v.user.password.$model"
                  type="password"
                  placeholder="请输入密码"
                  :state="validateState('password')"
                ></b-form-input>
              </b-input-group>
              <b-form-invalid-feedback :state="validateState('password')">
                密码长度至少为6位
              </b-form-invalid-feedback>

            </b-form-group>
            <b-form-group label="确认密码" description="Please confirm your password.">
              <b-input-group>
                <b-input-group-prepend is-text>
                  <b-icon icon="shield-lock-fill"></b-icon>
                </b-input-group-prepend>
                <b-form-input
                  v-model="$v.user.confirmPwd.$model"
                  type="password"
                  placeholder="请确认密码）"
                  :state="validateState('confirmPwd')"
                ></b-form-input>
                </b-input-group>
              <b-form-invalid-feedback :state="validateState('confirmPwd')">
                两次密码输入不同
              </b-form-invalid-feedback>
            </b-form-group>

            <b-form-group>
              <b-input-group >
                <b-col sm="3"><b-img  @click="refreshCode" fluid thumbnail :src="captcha_url"></b-img></b-col>
                <b-col sm="9">
                <b-form-input
                size="lg"
                  v-model="user.captcha_code"
                  type="text"
                  placeholder="请输入验证码"
                ></b-form-input>
                </b-col>
              </b-input-group>
            </b-form-group>

            <b-form-group>
              <b-button
                @click="register"
                variant="outline-primary"
                block
              >注册</b-button>
            </b-form-group>
          </b-form>
        </b-card>
      </b-col>
    </b-row>
    <!-- <Footer/> -->
  </div>
</template>
<script>
import { required, minLength, maxLength, sameAs } from 'vuelidate/lib/validators';
import { mapActions } from 'vuex';
import request from '@/utils/request'

export default {
  data() {
    return {
      user: {
        name: '',
        password: '',
        confirmPwd: '',
        captcha_code: '',
      },
      captcha_url: 'url',
      captcha_id: '',
    };
  },
  components: {
  },
  validations: {
    user: {
      name: {
        required,
        minLength: minLength(2),
        maxLength: maxLength(20),
      },
      password: {
        required,
        minLength: minLength(6),
      },
      confirmPwd: {
        required,
        sameAsPassword: sameAs('password')
      },
      captcha_code: {
        required
      }
    },
  },
  methods: {
    ...mapActions('userModule', { userRegister: 'register' }),
    validateState(name) {
      // 这里是es6 解构赋值
      const { $dirty, $error } = this.$v.user[name];
      return $dirty ? !$error : null;
    },
    register() {
      // 验证数据
      this.$v.user.$touch();
      if (this.$v.user.$anyError) {
        return;
      }
      // 请求
      const params = {
        captcha_code: this.user.captcha_code,
        captcha_id: this.captcha_id,
        password: this.user.password,
        password_again: this.user.confirmPwd,
        username: this.user.name
      }
      this.userRegister(params).then(() => {
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
      this.captcha_id = res.data.data.captchaId
      this.captcha_url = "http://localhost:8000/api/v1/pub/captcha?captchaId=" + this.captcha_id
    }
  },
  created() {
    this.refreshCode()
  },
};
</script>

<style scoped>
.col-sm-3 {
  padding: 0;
}
.col-sm-9 {
  padding: 0;
}
img {
  height: 50px ;
  width: 120px;
}
.register {
  height: 100%;
  background-image: url('../../assets/img/register.jpg');
  background-size:cover;
}
.row {
  margin: 0 !important;
}
</style>
