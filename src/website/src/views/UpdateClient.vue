<template>
  <section>
    <div>
      <h1>Update Client Information</h1>
      <br>
      
      <el-row>
        <el-col :span="12">
          <el-form ref="form" :model="form" label-width="12.5em" label-position="left">
           <!-- CLIENT INFORMATION  --> 
          <h2 style="font-size:20px"> Client Information </h2>

            <el-form-item label="Client ID ">
              <el-input v-model="form.ClientID" :disabled="true"></el-input>
            </el-form-item>

            <el-form-item label="Client Name">
              <el-input v-model="form.ClientName"></el-input>
            </el-form-item>

            <el-form-item label="Client ABN">
              <el-input v-model="form.ClientABN"></el-input>
            </el-form-item>


            <el-form-item label="Client ACN">
              <el-input v-model="form.ClientACN"></el-input>
            </el-form-item>

            <el-form-item label="Location">
              <el-select v-model="form.ClientLocation" placeholder="Pick a location">
                <el-option v-for="option in options.locations" :key="option.ID" :label="option.Name" :value="option.ID"></el-option>
              </el-select>
            </el-form-item>

            <el-form-item label="Type">
              <el-select v-model="form.ClientType" placeholder="Pick a type">
                <el-option v-for="option in options.types" :key="option.ID" :label="option.Name" :value="option.ID"></el-option>
              </el-select>
            </el-form-item>

            <el-form-item label="Client Office Code">
              <el-select v-model="form.ClientOffice" placeholder="Pick an Office">
                <el-option v-for="option in options.offices" :key="option.ID" :label="option.Name" :value="option.ID"></el-option>
              </el-select>
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

          <h3 style="font-size:15px"> Street Address </h3>

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

          <h3 style="font-size:15px"> Postal Address </h3>

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
            <p v-if="errors.length">
              <b>Please correct the following error(s):</b>
              <ul>
               <li v-for="error in errors" v-bind:key="error">{{ error }}</li>
              </ul>
            </p>
   

            <!-- shane fix this -->
            <el-form-item>
              <el-button type="primary" @click="validate">Update Client</el-button>
            </el-form-item>
          </el-form>
        </el-col>
      </el-row>
    </div>
  </section>
</template>

<script>
import api from '@/api.conf';
// import Search from '@/views/Search';

export default {
  name: 'update-client',
  components: {
  },
  data() {
    return {
      state: {
        name: '',
        token: '',
      },
      errors: '',
      title: 'Update Client',
      content: '',
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
        ClientID: '',
        ClientName: '',
        ClientABN: '',
        ClientACN: '',
        ClientType: '',
        ClientLocation: '',
        ClientOffice: '',
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
      this.$router.replace('/updateclient/:id');
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
    validate() {
      this.errors = [];
      if (this.form.ClientName === '') {
        this.errors.push('Client Name Required');
      }
      if (this.form.ClientABN === '') {
        this.errors.push('Client ABN Required');
      }
      if (this.form.ClientACN === '') {
        this.errors.push('Client ACN Required');
      }
      if (this.form.ClientLocation === '') {
        this.errors.push('Client Location Required');
      }
      if (this.form.ClientType === '') {
        this.errors.push('Client Type Required');
      }
      if (this.form.ClientPhoneNumber === '') {
        this.errors.push('Client Phone Nunmber Required');
      }
      if (this.form.ClientEmail === '') {
        this.errors.push('Client Email Required');
      }
      if (this.form.ClientOffice === '') {
        this.errors.push('Client Office Code Required');
      }
      if (this.errors.length === 0) {
        //this.authenticate(); add this back in when we are done with testing
        this.redirecting();
        this.updatePage();
      }
//      this.redirecting('/NewClient');
    },
    redirecting() {
      fetch(api.addClient, {
        method: 'post',
        headers: {
          Accept: 'application/json',
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          ClientName: this.form.ClientName,
          ClientID: this.form.ClientID,
          ClientOfficeCode: this.form.ClientOffice,
          ClientABNNumber: this.form.ClientABN,
          ClientACNNumber: this.form.ClientACN,
          ClientTypeCode: this.form.ClientType,
          FirstName: this.form.ClientFirstName,
          LastName: this.form.ClientLastName,
          ClientLocationCode: this.form.ClientLocation,
          StreetAddress: this.form.ClientAddress,
          StreetSuburb: this.form.ClientSuburb,
          StreetPostcode: this.form.ClientPostcode,
          PhoneNumber: this.form.ClientPhoneNumber,
          FaxNumber: this.form.ClientFaxNumber,
          EMailAddress: this.form.ClientEmail,
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
          this.form.ClientID = data.ClientID;
          this.form.ClientName = data.ClientName;
          this.form.ClientABN = data.ClientABNNumber;
          this.form.ClientACN = data.ClientACNNumber;
          this.form.ClientType = data.ClientTypeCode;
          this.form.ClientLocation = data.ClientLocationCode;
          this.form.ClientOffice = data.ClientOfficeCode;
          this.form.ClientFirstName = data.FirstName;
          this.form.ClientLastName = data.LastName;
          this.form.ClientPhoneNumber = data.PhoneNumber;
          this.form.ClientFaxNumber = data.FaxNumber;
          this.form.ClientEmail = data.EMailAddress;
          this.form.ClientAddress = data.StreetAddress;
          this.form.ClientSuburb = data.StreetSuburb;
          this.form.ClientState = data.StreetState;
          this.form.ClientPostcode = data.StreetPostcode;
          this.form.ClientAddressPostal = data.PostalAddress;
          this.form.ClientSuburbPostal = data.PostalAddressSuburb;
          this.form.ClientStatePostal = data.PostalAddressState;
          this.form.ClientPostcodePostal = data.PostalAddressPostcode;
        });
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
      this.title = 'Update Client';
      this.projects = [];
      this.form = {
        ClientName: '',
        ClientABN: '',
        ClientACN: '',
        ClientType: '',
        ClientLocation: '',
        ClientOffice: '',
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
