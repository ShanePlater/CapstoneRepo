<!--
what this page needs:
  file chooser implemented and figure out what return values will be utilise
    - i dont think this needs any backend input due to only being the pointer for the file being chosen
  category field needs to be changed to a dropdown box which contains a list of categories
    - need a get request for the list of categories assuming they are held on a table in the db
  auth by field needs to be converted to a dropdown box which contains a list of people able to authorise documents
    - need a get request for the list of authorised people
  create button needs to take the document and upload a copy to a folder in the static directory
    - post request needs to take document information and store that in the database
    - i dont know what the naming convention nigel wants  for the document stored in the static directory.
    - i dont know how to tie this into the logged in user so the db knows who uploaded the file if necessary
  cancel button needs to return to previous or home page idk
    - this button is in the image that nigel sent so i assume he wants it included
  dateTimePicker needs to be a singular date instead of a range
    - 
  make sure all routers are connected and the page can be viewed
-->
<template>
  <section>
    <div v-if="title === 'Upload New Document'">
      <h1>{{ title }}</h1>
      <br>
     
      <el-row>
        <el-col :span="12">
          <el-form ref="form" :model="form" label-width="12.5em" label-position="left">
          <el-form-item label="Fields marked in bold are required.">
          </el-form-item>

           <!-- DOCUMENT INFORMATION  --> 
          <h2 style="font-size:20px"> Document Information </h2>
          
          <!-- document information -->
            <el-form-item label="Select File:">
             <!-- file chooser is inline with this label-->
            </el-form-item>
            <el-form-item>
              <el-button type="button" @click="filePicker">Choose File</el-button>
            </el-form-item>
            <el-form-item label="Category:">
              <!-- category selection in form of drop down box-->
            </el-form-item>            

            <el-form-item label="Document Name:">
              <el-input v-model="form.docName"></el-input>
            </el-form-item>


            <el-form-item label="File Revision:">
              <el-input v-model="form.fileRevision"></el-input>
            </el-form-item>


            <el-form-item label="Authorised By:">
              <el-input v-model="form.authBy"></el-input>
            </el-form-item>            

            <el-form-item label="Authorised Date:">
                <!-- single date pick dont need range-->
              <el-date-picker v-model="form.datePeriod" type="daterange" placeholder="Pick a range" :picker-options="datePicker">
              </el-date-picker>
            </el-form-item>
            <br>            

            <!-- This calls the redirecting method, which collects form data and sends it via an API call -->
            <el-form-item>
              <el-button type="primary" @click="validate">Create</el-button>
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="validate">Cancel</el-button>
            </el-form-item>
          </el-form>
            <p v-if="errors.length">
              <b>Please correct the following error(s):</b>
              <ul>
               <li v-for="error in errors" v-bind:key="error">{{ error }}</li>
              </ul>
            </p>
        </el-col>
      </el-row>
    </div>    
  </section>
</template>

<script>
import api from '@/api.conf';
export default {
  name: 'document-upload',
  components: {

  },
  data() {
    return {
      errors: [],
      title: 'Upload New Document',      
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
        docName: '',
        category: '',
        fileRevision: '',
        authBy: '',
        authDate: '1999-01-29 00:00:00',                
      },
    };
  },
  created() { // dont think i need this 
    if (this.$route.query.res === 'true') {
      this.$router.replace('/DocumentUpload');
    }
    /* maybe add a get options for category if it doesnt already exist
    this.getOptions(api.getOptionLocations);
    this.getOptions(api.getOptionTypes);
    this.getOptions(api.getOptionStatuss);
    this.getOptions(api.getOptionDivisions);
    this.getOptions(api.getOptionOffices);
    */
  },
  watch: {
    '$route.query.res': 'updatePage',
  },
  methods: {

    redirecting() {
      fetch(api.addProject, {
        method: 'post',
        headers: {
          Accept: 'application/json',
          'Content-Type': 'application/json',
        },        
        body: JSON.stringify({          
          docName: this.form.docName,
          category: this.form.catagory,
          fileRevision: this.form.fileRevision,
          authBy: this.form.authBy,
          authDate: this.form.datePeriod,  // this is probably wrong         
        }),
      });
    },
    validate() {
      this.errors = [];
      if (this.form.docName === '') {
        this.errors.push('Document Name Required');
      }
      if (this.form.catagory === '') {
        this.errors.push('Category not Selected');
      }
      if (this.form.fileRevision === '') {
        this.errors.push('File Revision Number Required');
      }  
      if (this.form.authBy === '') {
        this.errors.push('Autherised Personnel Selection Required');
      }     
      if (this.form.authDate === '1999-01-29 00:00:00') {
        this.errors.push('Authorisation Date Required');
      } 
      if (this.errors.length === 0) {
        this.redirecting();
        this.updatePage();
      }
      this.redirecting('/DocumentUpload');
    },
    /*
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
    */
    updatePage() {
      /*
      if (this.$route.query.res === 'true') {
        this.title = 'Upload New Document';
        return;
      }
      */
      this.title = 'Upload New Document';      
      this.form = {
       docName: '',
        category: '',
        fileRevision: '',
        authBy: '',
        authDate: '1999-01-29 00:00:00', 
      };
    },
  },
};
</script>

<style scoped>
</style>
