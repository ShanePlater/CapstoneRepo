<template>
<el-row>
  <el-col :span="24">
    <div class="copyright">
      <small>&#169; {{ date }} Lambert & Rehbein</small>
    </div>
  </el-col>
</el-row>
</template>

<script>
export default {
  name: 'layout-footer',
  data() {
    return {
      date: new Date().getFullYear(),
    };
  },
};
</script>

<style scoped>
.copyright {
  border-top: 1px solid rgba(132, 146, 166, 1);
  padding-bottom: 1em;
}

small {
  color: #8492A6;
}
</style>
