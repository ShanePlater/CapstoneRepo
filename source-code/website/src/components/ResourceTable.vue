<template>
  <el-table :data="table" stripe style="width: 100%">
    <el-table-column prop="FileFriendlyName" label="Document name" width="420" :show-overflow-tooltip="true">
    </el-table-column>
    <el-table-column prop="AuthorizedDate" label="Authorized Date" align="center">
    </el-table-column>
    <el-table-column prop="AuthorizedBy" label="Authorized By" width="140" >
    </el-table-column>
    <el-table-column prop="FileRevision" label="Rev" width="100">
    </el-table-column>
    <el-table-column label="Operations">
      <template scope="scope">
        <el-button class="btn" size="small" type="text" icon="document" @click="downFile(scope.row)">Download</el-button>
        <el-button class="btn" size="small" type="text" icon="view" @click="prevFile(scope.row)">Preview</el-button>
      </template>
    </el-table-column>
  </el-table>
</template>

<script>
export default {
  name: 'resource-table',
  props: {
    table: Array,
  },
  methods: {
    downFile(row) {
      const dl = document.createElement('a');
      dl.setAttribute('href', row.URL);
      dl.setAttribute('download', row.FileName);
      dl.click();
    },
    prevFile(row) {
      window.open(row.URL);
    },
  },
};
</script>
<style scoped>

</style>
