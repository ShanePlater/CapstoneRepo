<template>
<el-menu mode="vertical">
  <el-menu-item-group title="What's New">
    <el-menu-item class="item" @click="open(item.URL)" v-for="item in recent" :key="item.ID" index="nil">
      <span :title="item.FileName"><i class="el-icon-document"></i>{{ item.FileName.substring(0, 16) }}</span>
    </el-menu-item>
  </el-menu-item-group>
  <el-menu-item-group title="Suggestion Box">
    <el-menu-item class="item" @click="open('mailto:SuggestionBox@lar.net.au')" index="nil">
      <i class="el-icon-message"></i> Make a suggestion
    </el-menu-item>
  </el-menu-item-group>
</el-menu>
</template>

<script>
export default {
  name: 'whats-new',
  props: {
    recent: Array,
  },
  methods: {
    open: (url) => {
      window.open(url);
    },
  },
};
</script>

<style scoped>
.item {
  color: #324057;
}
</style>
