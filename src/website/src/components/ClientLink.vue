<template>
  <el-table :data="table" stripe style="width: 100%">
    <el-table-column prop="ID" label="Client ID" width="100" :show-overflow-tooltip="false">
    </el-table-column>    
    <el-table-column prop="Name" label="Name" width="240" :show-overflow-tooltip="true">
    </el-table-column>
    <el-table-column prop="PhoneNumber" label="Phone Number" width="140" :show-overflow-tooltip="true">
    </el-table-column>
    <el-table-column prop="FaxNumber" label="Fax Number" width="140" :show-overflow-tooltip="true">
    </el-table-column>
    <el-table-column prop="Address" label="Address" width="450" :show-overflow-tooltip="true">
    </el-table-column>
    <el-table-column>
      <template scope="scope">
        <el-button class="btn" size="small" type="text" icon="el-icon-document-add" @click="handleClick(scope.row)">Add</el-button>
      </template>
    </el-table-column>
  </el-table>
</template>

<script>
export default {
  name: 'client-link',
  props: {
    table: Array,
  },
  methods: {
    handleClick(row) {
      this.$router.push(`/new-project/${row.ID}`);
    },
  },
};
</script>