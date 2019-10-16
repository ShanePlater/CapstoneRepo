<!--
what this page needs:
  current file upload functionality is only for the helpdesk tickets 
  need to look at how i can either modify that to facilitate these files
  or replicate it for what is trying to be done

  need to find out how i can save the category path value if it cant be assigned when its selected

  this process might be split up into two seperate process
  1. upload file and get the file id
  2. upload 
  test everything
  https://laracasts.com/discuss/channels/vue/element-ui-data-attribute-not-getting-updated-on-submission
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
              <el-upload
                action="/api/v1/post/uploadResource" 
                ref="upload" 
                name="file"
                accept=".xlsx, .docx, .ppsx, .pdf"
                :multiple='false'               
                :auto-upload="false" 
                :on-success="handleSuccess"
                :on-error="handleError"
                :on-change="handleChange"                                                     
                :data="form"> <!-- change on-change to on-exceed -->
              <el-button slot="trigger" size="small" type="primary">Select file</el-button>
              </el-upload>
            </el-form-item>            
            <el-form-item label="Category:">
              <el-select v-model="form.categoryID" placeholder="Select A Category">
                <el-option v-for="option in options.categories" :key="option.ID" :label="option.Name" :value="option.ID"></el-option>
              </el-select>
            </el-form-item>
            <el-form-item label="Document Name:">
              <el-input v-model="form.friendlyFileName"></el-input>
            </el-form-item>
            <el-form-item label="File Revision:">
              <el-input v-model="form.fileRevision"></el-input>
            </el-form-item>
            <el-form-item label="Authorised By:">
              <el-select v-model="form.authorizedBy" placeholder="Select Personnel">
                <el-option v-for="option in options.users" :key="option.ID" :label="option.Name" :value="option.ID"></el-option>
              </el-select>
            </el-form-item> 
            <el-form-item label="Authorised Date:">
              <div class="demonstration">Value: {{form.authorizedDate}} </div>
                <!-- single date pick dont need range-->
              <el-date-picker v-model="form.authorizedDate" type="date" placeholder="Pick a date" format="yyyy/MM/dd" value-format='YYYY-MM-DD'>
              
              </el-date-picker> 
            </el-form-item>
            <br> 
            <!-- This calls the redirecting method, which collects form data and sends it via an API call -->
            <el-form-item>
              <el-button type="primary" @click="validate">Create</el-button>
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="doNothing">Cancel</el-button>
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
      uploadFile: '',
      title: 'Upload New Document',     
      options: {
        categories: [],
        users: [],
      },      
      form: {
        friendlyFileName: '',
        fileRevision: '',
        authorizedBy: '',
        authorizedDate: '',
        categoryID: '',
        url: '',
      },
    };
  },
  created() { 
    if (this.$route.query.res === 'true') {
      this.$router.replace('/DocumentUpload');
    }
    this.getOptions(api.getOptionCategories);
    this.getOptions(api.getOptionUsers);
  },
  watch: {
    '$route.query.res': 'updatePage',
  },
  methods: {
    redirecting() {
        var d = new Date(this.form.authorizedDate);
        var dateConv = d.toISOString();
        this.form.authorizedDate = dateConv;
       
        this.$refs.upload.submit();
    },
    validate() {
      this.errors = [];     
      if(this.uploadFile ===''){
        this.errors.push('File not attatched');
      }else{
        this.fileTypeCheck();
      }
      if (this.form.friendlyFileName === '') {
        this.errors.push('Document Name Required');
      }
      if (this.form.categoryID === '') {
        this.errors.push('Category not Selected');
      }
      if (this.form.fileRevision === '') {
        this.errors.push('File Revision Value Required');
      }
      if (this.form.authorizedBy === '') {
        this.errors.push('Autherized Personnel Selection Required');
      }
      if (this.form.authorizedDate === '') {
        this.errors.push('Authorization Date Required');
      }
      if (this.errors.length === 0) {          
        this.redirecting(); 
        this.updatePage();       
      }      
      this.updatePage();
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
            case api.getOptionCategories:
              this.options.categories = data;
              break;
            case api.getOptionUsers:
              this.options.users = data;
              break;
            default:
          }
        });
      });
    },
    updatePage() {
      // this is called if the page refreshes upon upload
      if (this.$route.query.res === 'true') {
        this.title = 'Upload New Document';
        this.form = {       
        friendlyFileName: '',
        fileRevision: '',
        authorizedBy: '',
        authorizedDate: '',
        categoryID: '',        
        };
        return;
      }
      this.title = 'Upload New Document';      
      
    },
    fileTypeCheck() {
      var path = require('path');
      //console.log(this.uploadFile)
      var extCheck =  path.extname(this.uploadFile);
      //console.log(extCheck);
      if(extCheck != '.xlsx' && extCheck != '.docx' && extCheck != '.ppsx' && extCheck != '.pdf'){
        this.errors.push('only .xlsx, .docx, .ppsx and .pdf files allowed'); 
        this.$refs.upload.abort();          
       
      }
    },
    handleChange(file) {
      this.uploadFile = file.name;
    },
    handleSuccess(response, file){            
      this.updatePage();      
    },    
    handleError(err, file) {
      this.errors.push(err);     
      this.updatePage();
    },
    doNothing() {

    },
  },
};
</script>
<style scoped>
</style>
