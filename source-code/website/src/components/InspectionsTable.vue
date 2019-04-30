<template>
  <el-table :data="table" stripe style="width: 100%">
    <el-table-column prop="ID" label="Inspection ID" width="130">
    </el-table-column>
    <el-table-column prop="InspectedBy" label="Inspected By" width="150">
    </el-table-column>
    <el-table-column prop="InspectionDateTime" label="Inspection Date" width="160" :show-overflow-tooltip="true">
    </el-table-column>
    <el-table-column prop="InspectionDetails" label="Details" :show-overflow-tooltip="true">
    </el-table-column>
  </el-table>
</template>

<script>
export default {
  name: 'inspections-table',
  props: {
    table: Array,
  },
};
</script>