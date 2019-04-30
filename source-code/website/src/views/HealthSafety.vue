<template>
  <section>
    <h1>Occupational Health & Safety</h1>
    <h3>Induction & Training</h3>
    <el-row>
      <resource-table :table="intro"></resource-table>
    </el-row>
    <h3>Downloads</h3>
    <el-row>
      <collapse :tables="table"></collapse>
    </el-row>
    <br>
    <el-button type="primary" @click="openEmployeeDetails">Emergency Contact Details</el-button>
    <br>
  </section>
</template>

<script>
import ResourceTable from '@/components/ResourceTable';
import Collapse from '@/components/Collapse';
import api from '@/api.conf';

export default {
  name: 'resources',
  components: {
    ResourceTable,
    Collapse,
  },
  data() {
    return {
      intro: [],
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
          CategoryID: 3,
        }),
      }).then((response) => {
        response.json().then((data) => {
          const convertedData = [];
          data.forEach((item) => {
            if (item.FileID.includes('OHS-T-')) {
              this.intro.push({
                AuthorizedBy: item.AuthorizedBy,
                AuthorizedDate: Intl.DateTimeFormat('en-AU').format(new Date(item.AuthorizedDate)),
                CategoryID: item.CategoryID,
                FileID: item.FileID,
                FileName: item.FileName,
                FileRevision: item.FileRevision.trim(),
                URL: item.URL,
              });
            } else {
              convertedData.push({
                AuthorizedBy: item.AuthorizedBy,
                AuthorizedDate: Intl.DateTimeFormat('en-AU').format(new Date(item.AuthorizedDate)),
                CategoryID: item.CategoryID,
                FileID: item.FileID,
                FileName: item.FileName,
                FileRevision: item.FileRevision.trim(),
                URL: item.URL,
              });
            }
          });
          let tempData = [];
          convertedData.forEach((item) => {
            if (item.FileID.includes('OHS-EEMP-')) {
              tempData.push(item);
            }
          });
          tempData.sort((a, b) => (a.FileName.charCodeAt(0) - b.FileName.charCodeAt(0)));
          this.table.push({
            Name: 'Evacuation Management Plan',
            Data: tempData,
          });
          tempData = [];
          convertedData.forEach((item) => {
            if (item.FileID.includes('OHS-F-00')) {
              tempData.push(item);
            }
          });
          tempData.sort((a, b) => (a.FileName.charCodeAt(0) - b.FileName.charCodeAt(0)));
          this.table.push({
            Name: 'Forms',
            Data: tempData,
          });
          tempData = [];
          convertedData.forEach((item) => {
            if (item.FileID.includes('OHS-O-')) {
              tempData.push(item);
            }
          });
          tempData.sort((a, b) => (a.FileName.charCodeAt(0) - b.FileName.charCodeAt(0)));
          this.table.push({
            Name: 'Manual',
            Data: tempData,
          });
          tempData = [];
          convertedData.forEach((item) => {
            if (item.FileID.includes('OHS-P-')) {
              tempData.push(item);
            }
          });
          tempData.sort((a, b) => (a.FileName.charCodeAt(0) - b.FileName.charCodeAt(0)));
          this.table.push({
            Name: 'Policy',
            Data: tempData,
          });
          tempData = [];
          convertedData.forEach((item) => {
            if (item.FileID.includes('OHS-RA-')) {
              tempData.push(item);
            }
          });
          tempData.sort((a, b) => (a.FileName.charCodeAt(0) - b.FileName.charCodeAt(0)));
          this.table.push({
            Name: 'Risk Assessment Forms',
            Data: tempData,
          });
          tempData = [];
          convertedData.forEach((item) => {
            if (item.FileID.includes('OHS-F-019')) {
              tempData.push(item);
            }
          });
          tempData.sort((a, b) => (a.FileName.charCodeAt(0) - b.FileName.charCodeAt(0)));
          this.table.push({
            Name: 'Safety in Design Document Downloads',
            Data: tempData,
          });
        });
      });
    },
    openEmployeeDetails() {
      window.open(`//${window.location.host}/static/hr/Copy of Employee Details.xlsx?v=4`);
    },
  },
};
</script>

<style scoped>

</style>
