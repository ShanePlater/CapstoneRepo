<template>
  <section>
    <el-collapse v-model="activeNames" v-for="table in tables" :key="tables.indexOf(table)">
      <el-collapse-item :name="tables.indexOf(table)">
        <template slot="title"><span class="table">{{ table.Name }}</span></template>
        <resource-table :table="table.Data"></resource-table>
      </el-collapse-item>
    </el-collapse>
  </section>
</template>

<script>
import ResourceTable from '@/components/ResourceTable';

export default {
  name: 'collapse',
  components: {
    ResourceTable,
  },
  props: {
    tables: Array,
  },
  data() {
    return {
      activeNames: [0],
    };
  },
};
</script>

<style scoped>
.table {
  font-size: 1.2em;
  font-weight: 700;
  color: #324057;
}

</style>