<template>
  <div id="app">
    <layout-header></layout-header>
    <transition name="el-fade-in-linear">
      <router-view style="min-height:35em;"></router-view>
    </transition>
    <layout-footer style="margin-top:10%;"></layout-footer>
  </div>
</template>

<script>
import LayoutHeader from '@/components/LayoutHeader';
import LayoutFooter from '@/components/LayoutFooter';

export default {
  name: 'app',
  components: {
    LayoutHeader,
    LayoutFooter,
  },
};
</script>

<style>
#app {
  font-family: "Helvetica Neue", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  margin: auto;
  width: 73em;
}

h1,
h2,
h3,
h4,
h5,
h6,
small,
strong {
  color: #324057;
}

p {
  color: #1F2D3D;
}

.btn {
  color: #1F2D3D
}

.btn:hover {
  color: #20A0FF
}
</style>
