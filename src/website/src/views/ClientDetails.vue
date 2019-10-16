<template>
  <section>
    <el-row>
      <h1>{{ content.ClientName }}</h1>
    </el-row>
    <el-row>
      <h3>Client Information</h3>
      <el-button type="primary" @click="gotoupdate(content.ClientID)">Edit Client</el-button>
      <hr>
      <p><strong>ID:</strong> {{ content.ClientID }}</p>
      <p><strong>Name:</strong> {{ content.ClientName }}</p>
      <p><strong>ABN:</strong> {{ content.ClientABNNumber }}</p>
      <p><strong>ACN:</strong> {{ content.ClientACNNumber }}</p>
      <p><strong>Client Type:</strong> {{ content.ClientTypeCode }}</p>
      <p><strong>Client Location:</strong> {{ content.ClientLocationCode }}</p>
      <p><strong>Office code:</strong> {{ content.ClientOfficeCode }}</p>
    </el-row>
    <el-row>
      <h3>Contact Information</h3>
      <hr>
      <p><strong>First Name:</strong> {{ content.FirstName }}</p>
      <p><strong>Last Name:</strong> {{ content.LastName }}</p>
      <p><strong>Mobile Number:</strong> {{ content.MobileNumber }}</p>
      <p><strong>Phone Number:</strong> {{ content.PhoneNumber }}</p>
      <p><strong>Fax Number:</strong> {{ content.FaxNumber }}</p>
      <p><strong>E-mail Address:</strong> {{ content.EMailAddress }}</p>
    </el-row>
    <el-row>
      <h3>Location</h3>
      <hr>
      <p><strong>Street Address:</strong> {{ content.StreetAddress }}</p>
      <p><strong>Suburb:</strong> {{ content.StreetSuburb }}</p>
      <p><strong>State:</strong> {{ content.StreetState }}</p>
      <p><strong>Postcode:</strong> {{ content.StreetPostcode }}</p>
    </el-row>
    <el-row>
      <h3>Postal Address</h3>
      <hr>
      <p><strong>Address:</strong> {{ content.PostalAddress }}</p>
      <p><strong>Suburb:</strong> {{ content.PostalAddressSuburb }}</p>
      <p><strong>State:</strong> {{ content.PostalAddressState }}</p>
      <p><strong>Postcode:</strong> {{ content.PostalAddressPostcode }}</p>
      <el-button type="primary" @click="authenticate">Authenticate Check</el-button>
    </el-row>
    <br>
    <el-row>
      <h3>Projects for {{ content.ClientName }}</h3>
      <project-table :table="projects.slice"></project-table>
    </el-row>
    <br>
    <el-row type="flex" class="row-bg" align="middle" justify="center" v-if="totalPage != 0">
      <el-pagination layout="prev, pager, next" :total="totalPage" @current-change="updatePage">
      </el-pagination>
    </el-row>
  </section>
</template>

<script>
import api from '@/api.conf';
import ProjectTable from '@/components/ProjectTable';

export default {
  name: 'client-details',
  components: {
    ProjectTable,
  },
  data() {
    return {
      content: '',
      projects: {
        data: [],
        slice: [],
        page: 1,
      },
      page: 0,
      totalPage: 0,
    };
  },
  created() {
    this.pullData();
  },
  watch: {
    '$route.params.id': 'pullData',
    page: 'updateSlice',
  },
  methods: {
    gotoupdate(ClientID) {
      this.$router.push(`/updateclient/${ClientID}`);
    },
    authenticate(){
        fetch(api.authRequired, {
        method: 'post',
        headers: {
          Accept: 'application/json',
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          Username: 'training',
        }),
      }).then((response) => {
        response.json().then((data) => {
          if(true){
            window.location.href = 'mailto:Helpdesk@lar.net.au';
          }
        });
      });
    },
    pullData() {
      this.pullClientDetails();
      this.searchProjects();
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
    searchProjects() {
      fetch(api.getProjectsByClientID, {
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
          if (data === null) {
            this.projects.page = 0;
            this.projects.data = [];
            this.projects.slice = [];
            this.updateTotalPage(this.projects.data.length);
            this.page = this.projects.page;
            return;
          }
          this.projects.data = [];
          this.projects.page = 1;
          data.forEach((item) => {
            this.projects.data.push(item);
          });
          // this.projects.data.sort((a, b) => (a.ID.charCodeAt(0) - b.ID.charCodeAt(0)));
          this.projects.slice = this.generateNewSlice(this.projects);
          this.updateTotalPage(this.projects.data.length);
          this.page = this.projects.page;
        });
      });
    },
    generateNewSlice(data) {
      return data.data.slice((data.page - 1) * 10, data.page * 10);
    },
    updatePage(key) {
      this.page = key;
    },
    updateSlice() {
      this.projects.page = this.page;
      this.projects.slice = this.generateNewSlice(this.projects);
    },
    updateTotalPage(num) {
      if (num === 0) {
        this.totalPage = 0;
        return;
      }
      this.totalPage = num;
    },
  },
};
</script>

<style scoped>
p {
  text-indent: 2em;
}
</style>