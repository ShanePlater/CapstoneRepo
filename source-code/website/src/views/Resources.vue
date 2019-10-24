<template>
  <section>
    <h1>Resource</h1>
    <el-row>
      <p>{{ info }}</p>
      <collapse :tables="tables"></collapse>
    </el-row>
  </section>
</template>

<script>
import Collapse from '@/components/Collapse';
import api from '@/api.conf';

export default {
  name: 'resources',
  components: {
    Collapse,
  },
  data() {
    return {
      info: 'Download the most commonly used company documents and forms here.',
      tables: [],
    };
  },
  created() {
  // this.getResources(JSON.stringify({
    // CategoryID: 1,
    // }), 'Quality Assurance');
    this.getResources(JSON.stringify({
      CategoryID: 2,
    }), 'Human Resources');
    this.getResources(JSON.stringify({
      CategoryID: 4,
    }), 'Company Forms');
    this.getResources(JSON.stringify({
      CategoryID: 5,
    }), 'Insurance Certificates');
    this.getResources(JSON.stringify({
      CategoryID: 6,
    }), 'Information Technology');
    this.getResources(JSON.stringify({
      CategoryID: 7,
    }), 'OHS & QA Certificates');
    this.getResources(JSON.stringify({
      CategoryID: 8,
    }), 'Operationss');
    this.getResources(JSON.stringify({
      CategoryID: 9,
    }), 'Company Policies');
    this.getResources(JSON.stringify({
      CategoryID: 11,
    }), 'General Documents');
    this.getResources(JSON.stringify({
      CategoryID: 12,
    }), 'Work Instructions');
  },
  methods: {
    getResources(req, name) {
      fetch(api.getResources, {
        method: 'post',
        headers: {
          Accept: 'application/json',
          'Content-Type': 'application/json',
        },
        body: req,
      }).then((response) => {
        response.json().then((data) => {
          if (data === null) {
            return;
          }
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
          this.tables.push({
            Name: name,
            Data: convertedData,
          });
          this.tables.sort((a, b) => (a.Name.charCodeAt(0) - b.Name.charCodeAt(0)));
        });
      });
    },
  },
};
</script>

<style scoped>

</style>
