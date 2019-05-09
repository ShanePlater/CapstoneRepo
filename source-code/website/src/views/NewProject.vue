<template>
  <section>
    <div v-if="title === 'Enter Project Details'">
      <h1>{{ title }}</h1>
      <br>
     
      <el-row>
        <el-col :span="12">
          <el-form ref="form" :model="form" label-width="12.5em" label-position="left">
          <el-form-item label="Fields marked in bold are required.">
          </el-form-item> 

           <!-- PROJECT INFORMATION  --> 
          <h2 style="font-size:20px"> Project Information </h2>
          
          <!-- project information -->
            <el-form-item label="Client Name:">
              <el-input v-model="form.clientname"></el-input>
            </el-form-item>

            <el-form-item label="Project Number">
              <el-input v-model="form.projectnumber"></el-input>
            </el-form-item>            

            <el-form-item label="Project Name">
              <el-input v-model="form.projectname"></el-input>
            </el-form-item>


            <el-form-item label="Address:">
              <el-input v-model="form.projectaddress"></el-input>
            </el-form-item>


            <el-form-item label="Suburb:">
              <el-input v-model="form.projectsuburb"></el-input>
            </el-form-item>


            <el-form-item label="Location:">
              <el-input v-model="form.projectlocation"></el-input>
            </el-form-item>


            <el-form-item label="Type:">
              <el-input v-model="form.projecttypecode"></el-input>
            </el-form-item>


            <el-form-item label="Status:">
              <el-input v-model="form.projectstatuscode"></el-input>
            </el-form-item>


            <el-form-item label="Start and End Date:">
              <el-date-picker v-model="form.datePeriod" type="daterange" placeholder="Pick a range" :picker-options="datePicker">
              </el-date-picker>
            </el-form-item>

            <br>


             <!-- CLIENT REPRESENTATIVE CONTACT INFORMATION  --> 
          <h3 style="font-size:20px"> Client Representative Contact Information </h3>
          
          <!-- client rep. information -->

            <el-form-item label="Name:">
              <el-input v-model="form.clientrepname"></el-input>
            </el-form-item>

            
            <el-form-item label="Telephone:">
              <el-input v-model="form.clientrepworknum"></el-input>
            </el-form-item>


            <el-form-item label="Mobile: ">
              <el-input v-model="form.clientrepmobnum"></el-input>
            </el-form-item>


            <el-form-item label="Email Address:">
              <el-input v-model="form.clientrepemail"></el-input>
            </el-form-item>

          <br>

           <!-- INTERNAL INFORMATION  --> 
          <h3 style="font-size:20px"> Internal Information </h3>
          
          <!-- internal information -->

            <el-form-item label="Division:">
              <el-input v-model="form.division"></el-input>
            </el-form-item>

            <el-form-item label="Project Director:">
              <el-input v-model="form.projectdirector"></el-input>
            </el-form-item>         

             <el-form-item label="Project Manager">
              <el-input v-model="form.projectmanager"></el-input>
            </el-form-item>
            <br>


          <h4 style="font-size:20px"> Details </h4>
          <!-- Project Details -->
            <el-form-item label="Project Value: ($)">
              <el-input v-model="form.projectvalue"></el-input>
            </el-form-item>

            <el-form-item label="Project Description:">
              <el-input v-model="form.projectdescription"></el-input>
            </el-form-item>
            

            <!-- shane fix this -->
            <el-form-item>
              <el-button type="primary" @click="search">Add New Project</el-button>
            </el-form-item>
          </el-form>
        </el-col>
      </el-row>
    </div>    
  </section>
</template>

<script>
import api from '@/api.conf';
import Search from '@/views/Search';

export default {
  name: 'new-project',
  components: {
    Search,
  },
  data() {
    return {
      title: 'Enter Project Details',
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
        projectnumber: '',
        projectname: '',
        clientName: '',
        projectlocationcode: '',
        projectaddress: '',
        projectsuburb: '',
        projecttypecode: '',
        projectstatuscode: '',
        datePeriod: '',
        clientrepname: '',
        clientrepworknum: '',
        clientrepmobnum: '',
        clientrepemail: '',
        division: '',
        projectdirector: '',
        projectmanager: '',
        projectvalue: '',
        projectdescription: '',
      },
      isSearching: true,
    };
  },
  created() {
    if (this.$route.query.res === 'true') {
      this.$router.replace('/NewProject');
    }
    this.getOptions(api.getOptionLocations);
    this.getOptions(api.getOptionTypes);
    this.getOptions(api.getOptionStatuss);
    this.getOptions(api.getOptionDivisions);
    this.getOptions(api.getOptionOffices);
  },
  watch: {
    '$route.query.res': 'updatePage',
  },
  methods: {
    getOptions(method) {
      fetch(method, {
        method: 'get',
        headers: {
          Accept: 'application/json',
          'Content-Type': 'application/json',
        },
      }).then((response) => {
        response.json().then((data) => {
          switch (method) {
            case api.getOptionLocations:
              this.options.locations = data;
              break;
            case api.getOptionTypes:
              this.options.types = data;
              break;
            case api.getOptionStatuss:
              this.options.statuss = data;
              break;
            case api.getOptionDivisions:
              this.options.divisions = data;
              break;
            case api.getOptionOffices:
              this.options.offices = data;
              break;
            default:
          }
        });
      });
    },
    updatePage() {
      if (this.$route.query.res === 'true') {
        this.title = 'Search Result';
        return;
      }
      this.title = 'New Project';
      this.projects = [];
      this.form = {
        projectnumber: '',
        projectname: '',
        clientName: '',
        projectlocationcode: '',
        projectaddress: '',
        projectsuburb: '',
        projecttypecode: '',
        projectstatuscode: '',
        daterange: '',
        clientrepname: '',
        clientrepworknum: '',
        clientrepmobnum: '',
        clientrepemail: '',
        division: '',
        projectdirector: '',
        projectmanager: '',
        projectvalue: '',
        projectdescription: '',
      };
    },
  },
};
</script>

<style scoped>
</style>