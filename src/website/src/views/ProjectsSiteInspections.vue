<template>
  <section>
    <div v-if="title === 'New Site Inspection Record'">
      <h1>{{ title }}</h1>
      <br>
      <el-row>
        <el-col :span="12">
          <el-form ref="form" :model="form" label-width="12.5em" label-position="left">


           <!-- CLIENT INFORMATION  --> 
          <h2 style="font-size:20px"> Enter Site Inspection Record </h2>
          
          <!-- client information -->


            <el-form-item label="Project Number">
              <el-input v-model="form.ProjectNumber"></el-input>
            </el-form-item>

            <el-form-item label="Inspection Number">
              <el-input v-model="form.InspectionNumber"></el-input>
            </el-form-item>

            <el-form-item label="Inspected By">
              <el-input v-model="form.InspectedBy"></el-input>
            </el-form-item>       

            <el-form-item label="Inspection Date">
              <el-date-picker v-model="form.InspectionDate" type="date" placeholder="Pick a date">
              </el-date-picker>
            </el-form-item>

            <el-form-item label="Details">
              <el-input v-model="form.Details" type="textarea"></el-input>
            </el-form-item>
            

            <!-- shane fix this -->
            <el-form-item>
              <el-button type="primary" @click="redirecting">Add New Client</el-button>
            </el-form-item>
          </el-form>
        </el-col>
      </el-row>
    </div>
    <div v-else>
      <el-row v-loading="isSearching">
        <search :table="clients"></search>
      </el-row>
    </div>
  </section>
</template>

<script>
import api from '@/api.conf';
// import Search from '@/views/Search';

export default {
  name: 'new-client',
  components: {

  },
  data() {
    return {
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
        ClientName: '',
        ClientABN: '',
        ClientACN: '',
        ClientType: '',
        ClientLocation: '',
        ClientFirstName: '',
        ClientLastName: '',
        ClientPhoneNumber: '',
        ClientFaxNumber: '',
        ClientEmail: '',
        ClientAddress: '',
        ClientSuburb: '',
        ClientState: '',
        ClientPostcode: '',
        ClientAddressPostal: '',
        ClientSuburbPostal: '',
        ClientStatePostal: '',
        ClientPostcodePostal: '',


      },
      isSearching: true,
    };
  },
  created() {
    if (this.$route.query.res === 'true') {
      this.$router.replace('/NewClient');
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
    redirecting() {
      fetch(api.addClient, {
        method: 'post',
        headers: {
          Accept: 'application/json',
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          ClientName: this.form.ClientName,
          ClientOfficeCode: 'BNE',
          ClientABNNumber: this.form.ClientABN,
          ClientACNNumber: this.form.ClientACN,
          ClientTypeCode: this.form.projecttypecode,
          FirstName: this.form.ClientFirstName,
          LastName: this.form.ClientLastName,
          ClientLocationCode: this.form.projectlocationcode,
          StreetAddress: this.form.ClientAddress,
          StreetSuburb: this.form.ClientSuburb,
          StreetPostcode: this.form.ClientPostcode,
          PhoneNumber: this.form.ClientPhoneNumber,
          FaxNumber: this.form.ClientFaxNumber,
          EMailAddress: this.form.ClientEmail,
        }),
      });
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
      this.title = 'New Client';
      this.projects = [];
      this.form = {
        ClientName: '',
        ClientABN: '',
        ClientACN: '',
        ClientType: '',
        ClientLocation: '',
        ClientFirstName: '',
        ClientLastName: '',
        ClientPhoneNumber: '',
        ClientFaxNumber: '',
        ClientEmail: '',
        ClientAddress: '',
        ClientSuburb: '',
        ClientState: '',
        ClientPostcode: '',
        ClientAddressPostal: '',
        ClientSuburbPostal: '',
        ClientStatePostal: '',
        ClientPostcodePostal: '',

      };
    },
  },
};
</script>

<style scoped>

</style>
