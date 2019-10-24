<template>
  <el-card class="box-card">
    <div slot="header">
      <span class="title">
        {{ item.Title }}
      </span>
      <el-button-group style="float:right;">
        <el-button type="primary" icon="edit"></el-button>
      </el-button-group>
    </div>
    <div v-html="item.Text"></div>
    <time class="time">Posted on: {{ Intl.DateTimeFormat('en-AU').format(new Date(item.Date)) }}</time>
    <time class="time">Posted by: {{ item.Author }}</time>
  </el-card>
</template>

<script>
export default {
  name: 'anncmnt-card',
  props: {
    item: Object,
  },
};
</script>

<style scoped>
.title {
  line-height: 1.875em;
  font-size: 1.5em;
  color: #20A0FF;
}

.time {
  font-size: 0.812em;
  color: #999;
}
</style>
