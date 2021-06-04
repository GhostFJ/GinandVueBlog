<template>
  <div id="about">
      <b-container fluid class="header text-center">
        <b-avatar class="my-5" size="20em" icon="people-fill" variant="info"></b-avatar>
      </b-container>

      <b-container class="m-from-header">
        <h1 class="text-center text-muted display-4 mb-4">Hello, Everyone</h1>
        <h3 class="text-center">这是一个个人学习开发的博客网站，如果您有任何建议或想法，可以通过下面的发送和我联系.</h3>

        <h1 class="text-center text-muted display-4 my-5">Contact Me</h1>
        <b-form>
          <b-form-group>  
            <b-form-input type="email" placeholder="E-mail" required />
          </b-form-group>
          
          <b-form-group>
            <b-form-textarea placeholder="Message" rows="6" max-rows="6" />
          </b-form-group>

          <b-button class="mb-4" block>Send</b-button>

        </b-form>
      </b-container>
      

  </div>
</template>

<script>
export default {
  name: 'about',
}
</script>

<style scoped>
.header {
  background-color: gray;
  background-image: url('https://picsum.photos/1366/300/?image=3');
  background-repeat: no-repeat;
  background-size: cover;
  height: 18em;
}

.m-from-header {
  margin-top: 8em;
  animation: slideInFromBottom 0.3s;
}
</style>