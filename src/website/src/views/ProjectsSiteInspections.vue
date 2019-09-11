<template>
  <section>
    <el-row>
      <h1>Site Inspections for {{ projectNumber }}</h1>
    </el-row>
    <el-row>
      <p>{{ info }}</p>
      <inspections-table :table="table"></inspections-table>
    </el-row>
  </section>
</template>

<script>
import api from '@/api.conf';
import InspectionsTable from '@/components/InspectionsTable';

export default {
  name: 'projectsiteinspections',
  components: {
    InspectionsTable,
  },
  data() {
    return {
      table: [],
      projectNumber: this.$route.params.id,
    };
  },
  created() {
    this.getProjectsSiteInspectionsByProjectNumber();
  },
  methods: {
    getProjectsSiteInspectionsByProjectNumber() {
      fetch(api.getProjectsSiteInspectionsByProjectNumber, {
        method: 'post',
        headers: {
          Accept: 'application/json',
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          ID: this.$route.params.id,
        }),
      }).then((response) => {
        response.json().then((data) => {
          const convertedData = [];
          data.forEach((item) => {
            convertedData.push({
              ID: item.ID,
              InspectedBy: item.InspectedBy,
              InspectionDateTime: Intl.DateTimeFormat('en-AU').format(new Date(item.Date)),
              InspectionDetails: item.InspectionDetails,
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