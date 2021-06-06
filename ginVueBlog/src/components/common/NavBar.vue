<template>
  <div class="nav-bar">
    <div class="collapse">
      <a @click="collapse">
        <i v-if="!iscollapse" class="el-icon-s-fold"></i>
        <i v-else class="el-icon-s-unfold"></i>
      </a>
    </div>
    <h3 class="title">{{userInfo}}'s Blog</h3>
    <ul class="nav-bar-right">
      <li>
        <a @click="logout">
          <span class="icon-sign-in"></span>&nbsp;
          <span>退出</span>
        </a>
      </li>
    </ul>
  </div>
</template>

<script type="text/ecmascript-6">
import { mapState } from 'vuex'

export default {
  data() {
    return {
      iscollapse: false,
    }
  },
  computed: mapState({
    userInfo: (state) => state.userModule.userInfo,
  }),
  methods: {
    collapse() {
      this.iscollapse = !this.iscollapse
      this.$root.$emit('side-bar-collapse')
    },
    logout() {
      this.$api.auth.logout().then(data => {
        if (data.success) {
          this.$message({
            type: 'success',
            message: '登出成功!'
          })
          this.$router.push('/login')
        } else {
          this.$message({
            type: 'error',
            message: data.msg || '登出失败!'
          })
          this.$router.push('/login')
        }
      })
    }
  }
}
</script>

<style scoped>
.nav-bar {
  width: 100%;
  height: 100%;
  display: inline-block;
  background-color: rgb(206, 235, 240);
  line-height: 60px;
  text-align: center;
  box-shadow: 0 2px 3px hsla(0, 0%, 7%, 0.1), 0 0 0 1px hsla(0, 0%, 7%, 0.1);
}

.collapse {
  display: block;
  float: left;
  width: 64px;
}

.title {
  margin: 0;
  display: inline-block;
}

.nav-bar-right {
  list-style: none;
  float: right;
  margin: 0 15px 0 0;
  color: #262a2b;
}

.nav-bar-right a {
  color: #2f3435;
  cursor: pointer;
  text-decoration: none;
}

@media screen and (max-width: 600px) {
  .collapse {
    display: block;
  }
}
</style>
