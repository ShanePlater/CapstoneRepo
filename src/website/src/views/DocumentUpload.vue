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
                :multiple="false" 
                :auto-upload="false" 
                :before-upload="handleData"               
                :on-success="handleSuccess" 
                :on-remove="handleRemove"
                :on-progress="handleProgress"                                     
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
              <el-input v-model="form.authorizedBy"></el-input>
            </el-form-item> 
            <el-form-item label="Authorised Date:">
                <!-- single date pick dont need range-->
              <el-date-picker v-model="form.authorizedDate" type="text" placeholder="Pick a date" :picker-options="datePicker">
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
      title: 'Upload New Document',     
      options: {
        categories: [],
      },
      datePicker: {
      },
      form: {
        friendlyFileName: '',
        fileRevision: '',
        authorizedBy: '',
        authorizedDate: '1999-01-29 00:00:00',
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
  },
  watch: {
    '$route.query.res': 'updatePage',
  },
  methods: { // still throwing error in g.multipartform()
    redirecting() {      
        this.$refs.upload.submit();
    },
    validate() {
      this.errors = [];
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
      this.redirecting('/DocumentUpload'); // if their are errors reset the page and do nothing
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
            default:
          }
        });
      });
    },
    updatePage() {
      // this is called if the page refreshes upon upload
      if (this.$route.query.res === 'true') {
        this.title = 'Upload New Document';
        return;
      }
      this.title = 'Upload New Document';      
      this.form = {       
        friendlyFileName: '',
        fileRevision: '',
        authorizedBy: '',
        authorizedDate: '',
        categoryID: '',        
      };
    },
    handleSuccess(response, file){      
      this.updatePage();      
    },

    handleExceed() {
      this.errors.push('Only one file can be uplaoded at a time');    
      
    },
    handleRemove(file) {

    },
    handleData(file) {
          
    },
    handleProgress(event, file) {

    },
    doNothing() {

    },
  },
};
</script>
<style scoped>
</style>
