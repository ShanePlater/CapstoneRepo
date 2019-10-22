<template>
  <section>
      <h1>New Site Inspection Record</h1>
          <br>
          <el-row>
            <el-col :span="12">
              <el-form ref="form" :model="form" label-width="12.5em" label-position="left">

                <el-form-item label="Project Number:">
                  <el-input v-model="form.ProjectNumber" :disabled="true"></el-input>
                </el-form-item>

                <el-form-item label="Inspection Number:">
                  <el-input v-model="form.InspectionNumber"></el-input>
                </el-form-item>

                <el-form-item label="Inspected By:">
                  <el-input v-model="form.InspectedBy"></el-input>
                </el-form-item>       

            <el-form-item label="End Date:">
              <el-date-picker v-model="form.InspectionDate" type="date" placeholder="Pick a range" format="yyyy/MM/dd" value-format="yyyy-MM-dd">
              </el-date-picker>
            </el-form-item>          

                <el-form-item label="Details:">
                  <el-input v-model="form.Details" type="textarea"></el-input>
                </el-form-item>

                <el-form-item>
                  <el-button type="primary" @click="validate">Add Site Inspection</el-button>
                </el-form-item>

              </el-form>
            </el-col>
          </el-row>
  </section>
</template>
<script>

import api from '@/api.conf';

export default {
  name: 'siteinspection',
  components: {
  },
  data() {
    return {
      state: {
        name: '',
        token: '',
      },
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
        ProjectNumber: this.$route.params.id,
        InspectionNumber: '',
        InspectedBy: '',
        InspectionDate: '',
        Details: '',
      },
      isSearching: true,
    };
  },
  created() {
    if (this.getCookie('name') !== '') {
      this.state = {
        name: this.getCookie('name'),
        token: this.getCookie('token'),
        session: this.getCookie('mysession'),
      };
    }
  },
  methods: {
    getCookie(cname) {
      const name = `${cname}=`;
      let res = '';
      decodeURIComponent(document.cookie).split(';').forEach((ca) => {
        let a = ca;
        while (a.charAt(0) === ' ') {
          a = a.substring(1);
        }
        if (a.indexOf(name) === 0) {
          res = a.substring(name.length, a.length);
        }
      });
      return res;
    },
    authenticate() {
      fetch(api.authRequired, {
        method: 'post',
        headers: {
          Accept: 'application/json',
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          Username: this.state.name,
          Password: 'yeet',
        }),
      }).then((response) => {
        response.json().then((data) => {
          if (data.Username === 'Success') {
            this.redirecting();
          }
        });
      });
    },
    redirecting() {
      fetch(api.addSiteInspection, {
        method: 'post',
        headers: {
          Accept: 'application/json',
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          ProjectNumber: this.form.ProjectNumber,
          InspectionID: this.form.InspectionNumber,
          InspectedBy: this.form.InspectedBy,
          InspectionDateTime: this.form.InspectionDate,
          InspectionDetails: this.form.Details,
        }),
      });
    },
    validate() {
      this.errors = [];
      if (this.form.InspectionID === '') {
        this.errors.push('Inspection ID Required');
      }
      if (this.form.InspectedBy === '') {
        this.errors.push('Inspected By Required');
      }
      if (this.form.InspectionDate === '') {
        this.errors.push('Inspection Date Required');
      }
      if (this.form.InspectionDetails === '') {
        this.errors.push('Inspection Details Required');
      }
      if (this.errors.length === 0) {
        //this.authenticate();
        this.redirecting();
        //this.updatePage();
      }
//      this.redirecting('/NewProject');
    },
  },
};
</script>

<style scoped>

</style>
