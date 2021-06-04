<template>
<div id="navbar">
  <b-navbar toggleable="lg" type="light" variant="light">
    <b-navbar-brand to="/">
        <img src="../assets/logo.svg" width="32px" class="align-top">
        Personal<b>Blog</b>
    </b-navbar-brand>

    <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>

    <b-collapse id="nav-collapse" is-nav>
      <b-navbar-nav class="ml-auto">
          <b-nav-item to="/home">首页</b-nav-item>
          <b-nav-item to="/stories">文章</b-nav-item>
          <b-nav-item to="/categories">分类</b-nav-item>
          <b-nav-item to="">留言</b-nav-item>
          <b-nav-item to="/about">关于</b-nav-item>
          <b-nav-item to="/login">登录</b-nav-item>
      </b-navbar-nav>
    </b-collapse>
  </b-navbar>
</div>
</template>

<script>
export default {
  name: 'navbar',
}
</script>

<style>

</style>