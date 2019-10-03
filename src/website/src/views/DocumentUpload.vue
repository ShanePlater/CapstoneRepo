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
             <el-upload :auto-upload="false" :on-change="handleChange" :file-list="files"> <!-- change on-change to on-exceed -->
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
              <el-date-picker v-model="form.datePeriod" type="text" placeholder="Pick a date" :picker-options="datePicker">
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
      files: [],
      errors: [],
      title: 'Upload New Document',
      options: {
        categories: [],
      },
      datePicker: {
      },
      form: {
        fileName: '',
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
  methods: {
    redirecting() { 
      const formData = new FormData();   
      formData.append('files', this.fileList);             
      fetch(api.uploadResource, {
        method: 'post',
        headers: {
          Accept: 'application/json',      
          'content-type': 'multipart/form-data',
        },  
        body: JSON.stringify({  
          fileName: this.files[0].name,
          friendlyFileName: this.form.fileName,
          fileRevision: this.form.fileRevision,
          authorizedBy: this.form.authorizedBy,
          authorizedDate: this.form.datePeriod,
          categoryID: this.form.categoryID.value,
          url: this.form.url,
        }),       
      });     
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
        fileName: '',
        friendlyFileName: '',
        fileRevision: '',
        authorizedBy: '',
        authorizedDate: '',
        categoryID: '',
      };
    },
    handleChange(file, fileList) {
      this.files = fileList.splice(-1);
    },
    doNothing() {

    },
  },
};
</script>
<style scoped>
</style>
