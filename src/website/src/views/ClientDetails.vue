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

          <el-form ref="form" :model="form" label-width="20px" label-position="left">
            <p><strong>Client Location:</strong> {{ content.ClientLocationCode }}</p>
            <el-form-item>
              <el-select v-model="form.ClientLocation" placeholder="Pick a location">
                <el-option v-for="option in options.locations" :key="option.ID" :label="option.Name" :value="option.ID" :disabled="true"></el-option>
              </el-select>
            </el-form-item>
            <p><strong>Client Type:</strong> {{ content.ClientTypeCode }}</p>
            <el-form-item>
              <el-select v-model="form.ClientType" placeholder="Pick a type">
                <el-option v-for="option in options.types" :key="option.ID" :label="option.Name" :value="option.ID" :disabled="true"></el-option>
              </el-select>
            </el-form-item>
            <p><strong>Office code:</strong> {{ content.ClientOfficeCode }}</p>
            <el-form-item>
              <el-select v-model="form.ClientOffice" placeholder="Pick an Office">
                <el-option v-for="option in options.offices" :key="option.ID" :label="option.Name" :value="option.ID" :disabled="true"></el-option>
              </el-select>
            </el-form-item>
          </el-form>
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
      options: {
        locations: [
        {"ID": "1", "Name":	"Brisbane, QLD"},
        {"ID": "4", "Name":	"North Coast, QLD"},
        {"ID": "5", "Name":	"South Coast, QLD"},
        {"ID": "6", "Name":	"International"},
        {"ID": "7", "Name":	"Other"},
        {"ID": "8", "Name":	"Melbourne CBD"},
        {"ID": "9", "Name":	"Melbourne Metro"},
        {"ID": "10", "Name":	"Country North, VIC"},
        {"ID": "11", "Name":	"Country East, VIC"},
        {"ID": "12", "Name":	"Country West, VIC"},
        {"ID": "13", "Name":	"East Coast, VIC"},
        {"ID": "14", "Name":	"West Coast, VIC"},
        {"ID": "15", "Name":	"Wide Bay, QLD"},
        {"ID": "16", "Name":	"North QLD"},
        {"ID": "17", "Name":	"NSW"},
        {"ID": "18", "Name":	"NT"},
        {"ID": "19", "Name":	"TAS"},
        {"ID": "20", "Name":	"SA"},
        {"ID": "21", "Name":	"WA"},
        {"ID": "22", "Name":	"NZ"},
        {"ID": "23", "Name":	"ACT"},
        {"ID": "24", "Name":	"Sydney Metro"},
        {"ID": "25", "Name":	"Central Queensland"}
        ],
        types: [
        {"ID": "AC", "Name": "Airport Corportation"},
        {"ID": "C", "Name": "Contracter"},
        {"ID": "CG", "Name": "Commonwealth Government"},
        {"ID": "CT", "Name": "Consultant"},
        {"ID": "D", "Name": "Developer"},
        {"ID": "GOC", "Name": "Government Corporation"},
        {"ID": "I", "Name": "Institution"},
        {"ID": "LG", "Name": "Local Government"},
        {"ID": "P", "Name": "Private"},
        {"ID": "PM", "Name": "Project Manager"},
        {"ID": "SG", "Name": "State Government"}
        ],
        statuss: [],
        divisions: [],
        offices: [],
      },
      form: {
        ClientType: '',
        ClientLocation: '',
        ClientOffice: '',
      },
      page: 0,
      totalPage: 0,
    };
  },
  created() {
    this.pullData();
    this.getOptions(api.getOptionOffices);
  },
  watch: {
    '$route.params.id': 'pullData',
    page: 'updateSlice',
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
            /*
            case api.getClientLocations:
              this.options.locations = data;
              break;
            case api.getOptionTypes:
              this.options.types = data;
              break;
              */
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
    gotoupdate(ClientID) {
      this.$router.push(`/updateclient/${ClientID}`);
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
          this.form.ClientType = data.ClientTypeCode;
          this.form.ClientLocation = data.ClientLocationCode;
          this.form.ClientOffice = data.ClientOfficeCode;
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