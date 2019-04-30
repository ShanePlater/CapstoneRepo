<template>
  <el-table :data="table" stripe style="width: 100%">
    <el-table-column prop="Name" label="Name" width="150" :show-overflow-tooltip="true">
    </el-table-column>
    <el-table-column prop="PositionTitle" label="Position" width="250" :show-overflow-tooltip="true">
    </el-table-column>
    <el-table-column prop="PhoneNumber" label="Mobile Number" width="150" :show-overflow-tooltip="true">
    </el-table-column>
    <el-table-column prop="PhoneExtension" label="Phone Extension" width="160" :show-overflow-tooltip="true">
    </el-table-column>
    <el-table-column prop="UserDivision" label="Division" width="180" :show-overflow-tooltip="true">
    </el-table-column>
    <el-table-column prop="Email" label="Email" :show-overflow-tooltip="true">
    </el-table-column>
    <el-table-column>
      <template scope="scope">
        <el-button class="btn" size="small" type="text" icon="view" @click="handleClick(scope.row)">View</el-button>
      </template>
    </el-table-column>
  </el-table>
</template>

<script>
export default {
  name: 'user-table',
  props: {
    table: Array,
  },
  methods: {
    handleClick(row) {
      this.$route.push(`/user/${row.ID}`);
    },
  },
};
</script>
