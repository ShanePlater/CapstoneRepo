<template>
  <section>
    <h1>Policies & Procedures</h1>
    <el-row>
      <p>{{ info }}</p>
      <resource-table :table="table"></resource-table>
    </el-row>
  </section>
</template>

<script>
import ResourceTable from '@/components/ResourceTable';
import api from '@/api.conf';

export default {
  name: 'policies-procedures',
  components: {
    ResourceTable,
  },
  data() {
    return {
      info: 'View up to date company polices here.',
      table: [],
    };
  },
  created() {
    this.getResources();
  },
  methods: {
    getResources() {
      fetch(api.getResources, {
        method: 'post',
        headers: {
          Accept: 'application/json',
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          CategoryID: 9,
        }),
      }).then((response) => {
        response.json().then((data) => {
          const convertedData = [];
          data.forEach((item) => {
            convertedData.push({
              AuthorizedBy: item.AuthorizedBy,
              AuthorizedDate: Intl.DateTimeFormat('en-AU').format(new Date(item.AuthorizedDate)),
              CategoryID: item.CategoryID,
              FileID: item.FileID,
              FileName: item.FileName,
              FileRevision: item.FileRevision.trim(),
              URL: item.URL,
            });
          });
          this.table = convertedData;
        });
      });
    },
  },
};
</script>

<style scoped>

</style>
