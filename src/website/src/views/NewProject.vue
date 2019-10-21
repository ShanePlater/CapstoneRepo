<template>
  <section>
    <div v-if="title === 'Enter Project Details'">
      <h1>{{ title }}</h1>
      <br>
      <el-row>
        <el-col :span="12">
          <el-form ref="form" :model="form" label-width="12.5em" label-position="left">
          <el-form-item label="Enter information below">
          </el-form-item>
           <!-- PROJECT INFORMATION  --> 
          <h2 style="font-size:20px"> Project Information </h2>

          <p><strong>Client Name:</strong> {{ content.ClientName }}</p>
          <!-- project information -->
            <el-form-item label="Client ID:">
              <el-input v-model="form.ClientID" :disabled="true"></el-input>
            </el-form-item>

            <el-form-item label="Project Number">
              <el-input v-model="form.projectnumber" :disabled="true"></el-input>
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

            <el-form-item label="Location">
              <el-select v-model="form.projectlocationcode" placeholder="Pick a location">
                <el-option v-for="option in options.locations" :key="option.ID" :label="option.Name" :value="option.ID"></el-option>
              </el-select>
            </el-form-item>

            <el-form-item label="Type">
              <el-select v-model="form.projecttypecode" placeholder="Pick a type">
                <el-option v-for="option in options.types" :key="option.ID" :label="option.Name" :value="option.ID"></el-option>
              </el-select>
            </el-form-item>

            <el-form-item label="Status">
              <el-select v-model="form.projectstatuscode" placeholder="Pick a status">
                <el-option v-for="option in options.statuss" :key="option.ID" :label="option.Name" :value="option.ID"></el-option>
              </el-select>
            </el-form-item>


            <el-form-item label="Start Date:">
              <el-date-picker v-model="form.StartDate" type="date" placeholder="Pick a range" format="yyyy/MM/dd" value-format="yyyy-MM-dd">
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
          <h3 style="font-size:20px"> Internal Information {{state.name}} </h3>
          
          <!-- internal information -->

            <el-form-item label="Division">
              <el-select v-model="form.division" placeholder="Pick a division">
                <el-option v-for="option in options.divisions" :key="option.Name" :label="option.Name" :value="option.Name"></el-option>
              </el-select>
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
            
            <el-form-item label="Archive Location:">
              <el-input v-model="form.archivelocation"></el-input>
            </el-form-item>
            <p v-if="errors.length">
              <b>Please correct the following error(s):</b>
              <ul>
               <li v-for="error in errors" v-bind:key="error">{{ error }}</li>
              </ul>
            </p>
            <!-- This calls the redirecting method, which collects form data and sends it via an API call -->
            <el-form-item>
              <el-button type="primary" @click="validate">Add New Project</el-button>
            </el-form-item>
          </el-form>

        </el-col>
      </el-row>
    </div>    
  </section>
</template>

<script>
import api from '@/api.conf';

export default {
  name: 'new-project',
  components: {
  },
  data() {
    return {
      state: {
        name: '',
        token: '',
      },
      content: '',
      errors: [],
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
        projectnumber: '0',
        projectname: '',
        ClientID: this.$route.params.id,
        projectlocationcode: '',
        projectaddress: '',
        projectsuburb: '',
        projecttypecode: '',
        projectstatuscode: '',
        StartDate: '',
        clientrepname: '',
        clientrepworknum: '',
        clientrepmobnum: '',
        clientrepemail: '',
        division: '',
        projectdirector: '',
        projectmanager: '',
        projectvalue: '',
        projectdescription: '',
        archivelocation: '',
      },
    };
  },
  created() {
    if (this.$route.query.res === 'true') {
      this.$router.replace('/NewProject');
    }
    this.pullClientDetails();
    this.getOptions(api.getOptionLocations);
    this.getOptions(api.getOptionTypes);
    this.getOptions(api.getOptionStatuss);
    this.getOptions(api.getOptionDivisions);
    this.getOptions(api.getOptionOffices);
    
    if (this.getCookie('name') !== '') {
      this.state = {
        name: this.getCookie('name'),
        token: this.getCookie('token'),
        session: this.getCookie('mysession'),
      };
    }
  },
  watch: {
    '$route.query.res': 'updatePage',
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
      fetch(api.addProject, {
        method: 'post',
        headers: {
          Accept: 'application/json',
          'Content-Type': 'application/json',
        },
        //  the start and end dates here are they hardcoded or
        //  are they just placeholders where the data gets taken from the datetime picker
        body: JSON.stringify({
          ID: this.form.projectnumber,
          Name: this.form.projectname,
          ClientID: this.form.ClientID,
          Address: this.form.projectaddress,
          Suburb: this.form.projectsuburb,
          Location: this.form.projectlocationcode,
          Type: this.form.projecttypecode,
          Status: this.form.projectstatuscode,
          StartDate: '1999-01-29 00:00:00',
          EndDate: '1999-01-29 00:00:00',
          CRName: this.form.clientrepname,
          CRPhone: this.form.clientrepworknum,
          CRMobile: this.form.clientrepmobnum,
          CREmail: this.form.clientrepemail,
          Division: this.form.division,
          Director: this.form.projectdirector,
          Manager: this.form.projectmanager,
          Value: this.form.projectvalue,
          Description: this.form.projectdescription,
          ArchiveLocation: this.form.archivelocation,
        }),
      });
    },
    pullClientDetails() {
      fetch(api.getClient, {
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
          this.content = data;
        });
      });
    },
    validate() {
      this.errors = [];
      if (this.form.ClientID === '') {
        this.errors.push('Client ID Required');
      }
      if (this.form.projectnumber === '') {
        this.errors.push('Project Number Required');
      }
      if (this.form.projectname === '') {
        this.errors.push('Project Name Required');
      }
      if (this.form.projectlocationcode === '') {
        this.errors.push('Project Location Required');
      }
      if (this.form.projecttypecode === '') {
        this.errors.push('Project Type Required');
      }
      if (this.form.projectstatuscode === '') {
        this.errors.push('Project Status Required');
      }
      if (this.form.division === '') {
        this.errors.push('Internal Division Required');
      }
      if (this.form.projectdirector === '') {
        this.errors.push('Internal Project Director  Required');
      }
      if (this.form.projectmanager === '') {
        this.errors.push('Internal Project Manager Required');
      }
      if (this.form.projectdescription === '') {
        this.errors.push('Project Description Required');
      }
      if (this.errors.length === 0) {
        //this.authenticate();
        this.redirecting();
        this.updatePage();
      }
      this.updatePage();
//      this.redirecting('/NewProject');
    },
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
        this.title = '';
        return;
      }
      this.title = 'New Project';
      this.projects = [];
      this.form = {
        projectnumber: '',
        projectname: '',
        ClientID: '',
        projectlocationcode: '',
        projectaddress: '',
        projectsuburb: '',
        projecttypecode: '',
        projectstatuscode: '',
        StartDate: '',
        clientrepname: '',
        clientrepworknum: '',
        clientrepmobnum: '',
        clientrepemail: '',
        division: '',
        projectdirector: '',
        projectmanager: '',
        projectvalue: '',
        projectdescription: '',
        archivelocation: '',
      };
    },
  },
};
</script>

<style scoped>
</style>