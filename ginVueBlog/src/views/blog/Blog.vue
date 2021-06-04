<template>
  <div id="blog">
    <Navbar />
    <router-view />    
    <Footer />
  </div>
</template>

<script>
import Navbar from '@/components/Navbar'
import Footer from '@/components/Footer'

export default {
  name: 'Blog',
  components: {
    Navbar,
    Footer
  }
}
</script>

<style>

</style>