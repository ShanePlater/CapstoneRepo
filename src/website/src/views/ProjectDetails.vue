<template>
  <section>
    <el-row>
      <h1>Project: {{ content.ProjectName }}</h1>
    </el-row>
    <el-row>
      <h3>Project Information</h3>
      <el-button type="primary" @click="gotoupdate(content.ProjectNumber)">Update Project Information</el-button>
      <hr>
      <p><strong>Project ID:</strong> {{ content.ProjectNumber }}</p>
      <p><strong>Project Name:</strong> {{ content.ProjectName }}</p>
      <p><strong>Client Name:</strong> {{ clientcontent.ClientName }}    <el-button type="primary" @click="gotoclient(clientcontent.ClientID)">View Client</el-button></p>
      <p><strong>Location:</strong> {{ locationame }}</p>
      <p><strong>Address:</strong> {{ content.ProjectAddress }}</p>
      <p><strong>Suburb:</strong> {{ content.ProjectSuburb }}</p>
      <p><strong>Type:</strong> {{ typename }}</p>
      <p><strong>Status:</strong> {{ statusname }}</p>
      <p><strong>Start date:</strong> {{ startDate }}</p>
      <p><strong>End date:</strong> {{ endDate }}</p>
    </el-row>
    <el-row>
      <h3>Client Representative Contact Information</h3>
      <hr>
      <p><strong>Name:</strong> {{ content.ClientRepName }}</p>
      <p><strong>Telephone:</strong> {{ content.ClientRepTelephone }}</p>
      <p><strong>Mobile:</strong> {{ content.ClientRepMobile }}</p>
      <p><strong>E-mail Address:</strong> {{ content.ClientRepEmailAddress }}</p>
    </el-row>
    <el-row>
      <h3>Internal Information</h3>
      <hr>
      <p><strong>Division:</strong> {{ content.Division }}</p>
      <p><strong>Project Director:</strong> {{ content.ProjectDirector }}</p>
      <p><strong>Project Manager:</strong> {{ content.ProjectManager }}</p>
      <p><strong>Project Value $:</strong> {{ content.ProjectValue }}</p>
      <p><strong>Description:</strong> {{ content.ProjectDescription }}</p>
    </el-row>
    <br>
    <el-row>
      <h3>Site Inspections for {{ content.ProjectNumber }}</h3>
      <inspections-table :table="inspections.slice"></inspections-table>
    </el-row>
    <br>
    <el-row type="flex" class="row-bg" align="middle" justify="center" v-if="totalPage != 0">
      <el-pagination layout="prev, pager, next" :total="totalPage" @current-change="updatePage">
      </el-pagination>
    </el-row>
    <br>
    <el-row>
      <el-button type="primary" @click="gotonewinspection(content.ProjectNumber)">Add New Site Inspection</el-button>
    </el-row>
  </section>
</template>

<script>
import api from '@/api.conf';
import InspectionsTable from '@/components/InspectionsTable';
import router from '@/router';

export default {
  name: 'project-details',
  components: {
    InspectionsTable,
  },
  data() {
    return {
      content: '',
      clientcontent: '',
      clientURL: '',
      startDate: '',
      endDate: '',
      locationame: '',
      typename: '',
      statusname: '',
      options: {
        locations: [],
        types: [],
        statuss: [],
        divisions: [],
        offices: [],
      },
      inspections: {
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
    this.getOptions(api.getOptionLocations);
    this.getOptions(api.getOptionTypes);
    this.getOptions(api.getOptionStatuss);
    this.getOptions(api.getOptionDivisions);
    this.getOptions(api.getOptionOffices);
    this.setlocation();
    this.settype();
    this.setstatus();
  },
  watch: {
    '$route.params.id': 'pullData',
    page: 'updateSlice',
  },
  methods: {
    gotoupdate(ProjectNumber) {
      this.$router.push(`/updateproject/${ProjectNumber}`);
    },
    setlocation() {
      this.options.locations.forEach((option) => {
        console.log(option);
        console.log("Test");
        if (this.conent.ProjectLocationCode === "NA") {
          this.locationame = "N/A";
        }
        if (this.conent.ProjectLocationCode === "") {
          this.locationame = "N/A";
        }
        if(this.content.ProjectLocationCode === option.ID) {
          this.locationame = option.Name;
        }
      });
      this.locationame = "Test";
    },
    settype() {
      console.log('were here and were queer');
      for(option in this.options.types) {
        console.log(option.ID);
        console.log('fuck')
        if(this.content.ProjectTypeCode === option.ID) {
          this.typename = option.Name;
        } else {
          this.typename = 'broke but not totally fucked';
        }
      }
    },
    setstatus() {
      var option;
      for(option in this.options.statuss) {
        if(this.content.ProjectStatusCode === option.ID) {
          this.statusname = option.Name;
        }
      }
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
    pullData() {
      this.pullProjectDetails();
      this.searchInspections();
    },
    gotonewinspection(projectnumberid) {
      this.$router.push(`/siteinspection/${projectnumberid}`);
    },
    gotoclient(clientid) {
      this.$router.push(`/client/${clientid}`);
    },
    pullProjectDetails() {
      fetch(api.getProject, {
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
          this.pullClientDetails(data.ClientID);
          if (this.content.ProjectStartDate !== '') {
            this.startDate = Intl.DateTimeFormat('en-AU').format(new Date(this.content.ProjectStartDate));
          }
          if (this.content.ProjectEndDate !== '') {
            this.endDate = Intl.DateTimeFormat('en-AU').format(new Date(this.content.ProjectEndDate));
          }
        });
      });
    },
    pullClientDetails(clientID) {
      fetch(api.getClient, {
        method: 'post',
        headers: {
          Accept: 'application/json',
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          ID: clientID,
        }),
      }).then((response) => {
        response.json().then((data) => {
          this.clientcontent = data;
        });
      });
    },
    searchInspections() {
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
          if (data === null) {
            this.inspections.page = 0;
            this.inspections.data = [];
            this.inspections.slice = [];
            this.updateTotalPage(this.inspections.data.length);
            this.page = this.inspections.page;
            return;
          }
          this.inspections.data = [];
          this.inspections.page = 1;
          data.forEach((item) => {
            this.inspections.data.push({
              ID: item.ID,
              InspectedBy: item.InspectedBy,
              InspectionDateTime: Intl.DateTimeFormat('en-AU').format(new Date(item.Date)),
              InspectionDetails: item.InspectionDetails,
            });
          });
          // this.projects.data.sort((a, b) => (a.ID.charCodeAt(0) - b.ID.charCodeAt(0)));
          this.inspections.slice = this.generateNewSlice(this.inspections);
          this.updateTotalPage(this.inspections.data.length);
          this.page = this.inspections.page;
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
      this.inspections.page = this.page;
      this.inspections.slice = this.generateNewSlice(this.inspections);
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

</style>