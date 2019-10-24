<template>
  <el-table :data="table" stripe style="width: 100%">
    <el-table-column prop="ID" label="Number" width="90" :show-overflow-tooltip="true">
    </el-table-column>
    <el-table-column prop="Name" label="Project Name" width="220" :show-overflow-tooltip="true"  v-if="table[0].Company !== ''">
    </el-table-column>
    <el-table-column prop="Name" label="Project Name" width="330" :show-overflow-tooltip="true" v-else>
    </el-table-column>
    <el-table-column prop="Company" label="Client Company" :show-overflow-tooltip="true" v-if="table[0].Company !== ''">
    </el-table-column>
    <el-table-column prop="ClientName" label="Client Name" :show-overflow-tooltip="true">
    </el-table-column>
    <el-table-column prop="Address" label="Address" :show-overflow-tooltip="true">
    </el-table-column>
    <el-table-column prop="Division" label="Division" width="200" :show-overflow-tooltip="true">
    </el-table-column>
    <el-table-column width="90">
      <template scope="scope">
        <el-button class="btn" size="small" type="text" icon="view" @click="handleClick(scope.row)">View</el-button>
      </template>
    </el-table-column>
  </el-table>
</template>

<script>
export default {
  name: 'project-table',
  props: {
    table: Array,
  },
  methods: {
    handleClick(row) {
      this.$router.push(`/project/${row.ID}`);
    },
  },
};
</script>
