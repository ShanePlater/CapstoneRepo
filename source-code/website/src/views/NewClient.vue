<template>
  <section>
    <div v-if="title === 'New Client'">
      <h1>{{ title }}</h1>
      <br>
      <el-row>
        <el-col :span="12">
          <el-form ref="form" :model="form" label-width="12.5em" label-position="left">


           <!-- CLIENT INFORMATION  --> 
          <h2 style="font-size:20px"> Client Information </h2>
          
          <!-- client information -->


            <el-form-item label="Client Name">
              <el-input v-model="form.ClientName"></el-input>
            </el-form-item>

            <el-form-item label="Client ABN">
              <el-input v-model="form.ClientABN"></el-input>
            </el-form-item>


            <el-form-item label="Client ACN">
              <el-input v-model="form.ClientACN"></el-input>
            </el-form-item>


            <el-form-item label="Client Type">
              <el-input v-model="form.ClientType"></el-input>
            </el-form-item>


            <el-form-item label="Client Location">
              <el-input v-model="form.ClientLocation"></el-input>
            </el-form-item>

            <br>


             <!-- CLIENT CONTACT INFORMATION  --> 
          <h3 style="font-size:20px"> Contact Information </h3>
          
          <!-- client contact information -->

            <el-form-item label="First Name">
              <el-input v-model="form.ClientFirstName"></el-input>
            </el-form-item>

            
            <el-form-item label="Last Name">
              <el-input v-model="form.ClientLastName"></el-input>
            </el-form-item>


            <el-form-item label="Client Phone Number ">
              <el-input v-model="form.ClientPhoneNumber"></el-input>
            </el-form-item>


            <el-form-item label="Client Fax Number">
              <el-input v-model="form.ClientFaxNumber"></el-input>
            </el-form-item>

            <el-form-item label="Client E-Mail Address">
              <el-input v-model="form.ClientEmail"></el-input>
            </el-form-item>

          <br>


           <!-- ADDRESSES  --> 
          <h3 style="font-size:20px"> Internal Information </h3>
          
          <!-- address information -->

          <h3 style="font-size:10px"> Street Address </h3>

            <el-form-item label="Address">
              <el-input v-model="form.ClientAddress"></el-input>
            </el-form-item>

            <el-form-item label="Suburb">
              <el-input v-model="form.ClientSuburb"></el-input>
            </el-form-item>

             <el-form-item label="State">
              <el-input v-model="form.ClientState"></el-input>
            </el-form-item>

            <el-form-item label="Postcode">
              <el-input v-model="form.ClientPostcode"></el-input>
            </el-form-item>


         <!-- postal address -->

          <h3 style="font-size:10px"> Postal Address </h3>

            <el-form-item label="Address">
              <el-input v-model="form.ClientAddressPostal"></el-input>
            </el-form-item>

            <el-form-item label="Suburb">
              <el-input v-model="form.ClientSuburbPostal"></el-input>
            </el-form-item>

            <el-form-item label="State">
              <el-input v-model="form.ClientStatePostal"></el-input>
            </el-form-item>

            <el-form-item label="Postcode">
              <el-input v-model="form.ClientPostcodePostal"></el-input>
            </el-form-item>
            

            <!-- shane fix this -->
            <el-form-item>
              <el-button type="primary" @click="search">Add New Client</el-button>
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
import Search from '@/views/Search';

export default {
  name: 'new-client',
  components: {
    Search,
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
