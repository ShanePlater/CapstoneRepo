<template>
  <section>
      <h1>New Site Inspection Record</h1>
          <br>
          <el-row>
            <el-col :span="12">
              <el-form ref="form" :model="form" label-width="12.5em" label-position="left">

                <el-form-item label="Project Number:">
                  <el-input v-model="form.ProjectNumber"></el-input>
                </el-form-item>

                <el-form-item label="Inspection Number:">
                  <el-input v-model="form.InspectionNumber"></el-input>
                </el-form-item>

                <el-form-item label="Inspected By:">
                  <el-input v-model="form.InspectedBy"></el-input>
                </el-form-item>       

                <el-form-item label="Inspection Date:">
                  <el-date-picker v-model="form.datePeriod" type="daterange" placeholder="Pick a range" :picker-options="datePicker">
                  </el-date-picker>
                </el-form-item>

                <el-form-item label="Details:">
                  <el-input v-model="form.Details" type="textarea"></el-input>
                </el-form-item>
              </el-form>
            </el-col>
          </el-row>
  </section>
</template>
<script>

export default {
  name: 'siteinspection',
  components: {},
  data() {
    return {
      title: 'New Client',
      projects: [],
      options: {
        locations: [],
        types: [],
        statuss: [],
        divisions: [],
        offices: [],
      },
      datePicker: {
        shortcuts: [{
          text: 'Last week',
          onClick(picker) {
            const end = new Date();
            const start = new Date();
            start.setTime(start.getTime() - (3600 * 1000 * 24 * 7));
            picker.$emit('pick', [start, end]);
          },
        }, {
          text: 'Last month',
          onClick(picker) {
            const end = new Date();
            const start = new Date();
            start.setTime(start.getTime() - (3600 * 1000 * 24 * 30));
            picker.$emit('pick', [start, end]);
          },
        }, {
          text: 'Last 3 months',
          onClick(picker) {
            const end = new Date();
            const start = new Date();
            start.setTime(start.getTime() - (3600 * 1000 * 24 * 90));
            picker.$emit('pick', [start, end]);
          },
        }],
      },
      form: {
        ProjectNumber: '',
        InspectionNumber: '',
        InspectedBy: '',
        InspectionDate: '',
        Details: '',
      },
      isSearching: true,
    };
  },
  methods: {
    contactHelpdesk() {
      window.location.href = 'mailto:Helpdesk@lar.net.au';
    },
  },
};
</script>

<style scoped>

</style>
