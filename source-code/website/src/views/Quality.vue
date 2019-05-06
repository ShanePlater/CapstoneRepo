<template>
  <section>
    <h1>Quality</h1>
    <h3>About the Quality Assurance System</h3>
    <p>Lambert & Rehbein recognise that a formal system is required to ensure that all staff produce a high standard of quality in performing their tasks. Accordingly, to ensure that quality is well managed throughout the organisation, Lambert &Rehbein established a Quality Management System based on the requirement of ISO9001:2008. Staff are required to implement the Quality System on all projects and pro-actively seek to improve the system in line with the overall goals and values on the Company.</p>
    <p>Lambert & Rehbein’s Quality Management System is certified by SAI Global complying with ISO 9001:2008</p>
    <resource-table :table="intro"></resource-table>
    <br>
    <h3>Using QMS Documents</h3>
    <p>All procedures and forms listed on this page are Controlled Copies. Once a QMA documents has been printed, it is no longer a controlled copy.</p>
    <br>
    <h3>Quality Manual Downloads</h3>
    <collapse :tables="table"></collapse>
    <br>
    <el-button type="primary" @click="openCorrectiveActionRegister">Corrective Action Register</el-button>
    <br><br>
    <el-button type="primary" @click="makeSuggestion">Make a suggestion</el-button>
    <br>
  </section>
</template>

<script>
import ResourceTable from '@/components/ResourceTable';
import Collapse from '@/components/Collapse';
import api from '@/api.conf';

export default {
  name: 'quality',
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
          CategoryID: 1,
        }),
      }).then((response) => {
        response.json().then((data) => {
          const convertedData = [];
          data.forEach((item) => {
            if (item.FileID === 'QA-T-001.ppsx') {
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
            if (item.FileID.includes('QA-M-')) {
              tempData.push(item);
            }
          });
          tempData.sort((a, b) => (a.FileName.charCodeAt(0) - b.FileName.charCodeAt(0)));
          this.table.push({
            Name: 'Management of QA System',
            Data: tempData,
          });
          tempData = [];
          convertedData.forEach((item) => {
            if (item.FileID.includes('QA-A-')) {
              tempData.push(item);
            }
          });
          tempData.sort((a, b) => (a.FileName.charCodeAt(0) - b.FileName.charCodeAt(0)));
          this.table.push({
            Name: 'Administration & Information Technology',
            Data: tempData,
          });
          tempData = [];
          convertedData.forEach((item) => {
            if (item.FileID.includes('QA-P-')) {
              tempData.push(item);
            }
          });
          tempData.sort((a, b) => (a.FileName.charCodeAt(0) - b.FileName.charCodeAt(0)));
          this.table.push({
            Name: 'Projects & Jobs',
            Data: tempData,
          });
          tempData = [];
          convertedData.forEach((item) => {
            if (item.FileID.includes('QA-O-')) {
              tempData.push(item);
            }
          });
          tempData.sort((a, b) => (a.FileName.charCodeAt(0) - b.FileName.charCodeAt(0)));
          this.table.push({
            Name: 'Operating Manuals',
            Data: tempData,
          });
          tempData = [];
          convertedData.forEach((item) => {
            if (item.FileID.includes('QA-F-')) {
              tempData.push(item);
            }
          });
          tempData.sort((a, b) => (a.FileName.charCodeAt(0) - b.FileName.charCodeAt(0)));
          this.table.push({
            Name: 'Forms',
            Data: tempData,
          });
        });
      });
    },
    openCorrectiveActionRegister() {
      window.open(`//${window.location.host}/static/quality-management-system/corrective-action-register/QA-F-005.pdf?v=4`);
    },
    makeSuggestion() {
      location.href = 'mailto:LARQualityAssurance@lar.net.au';
    },
  },
};
</script>

<style scoped>

</style>
