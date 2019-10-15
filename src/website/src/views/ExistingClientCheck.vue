<template>
  <div>
    <h1>Does this project belong to a new or existing client?</h1>
    <br>
    <login> </login>
    <p>If your client does not already exist, please click the the link below to add a new client. If the client does exist, search it in the field below then click Continue
    <el-button type="primary" @click="redirecting">Add New Client</el-button>  
    <!-- Form to enter keyword -->
    <el-form ref="form" :model="form" label-width="12.5em" label-position="left">
      <el-form-item label="Client Name:">
        <el-input v-model="form.clientName"></el-input>
      </el-form-item>  
    </el-form>
    <el-button type="primary" @click="searchClients">Search Clients</el-button>
    <div>
        <el-row>
          <h1>Search Results</h1>
        </el-row>
        <el-row type="flex" class="row-bg" v-if="table === undefined">
          <el-col :span="24">
            <el-tabs v-model="name" type="card">
              <el-tab-pane label="Client Profiles" name="clients">
                <client-table :table="clients.slice"></client-table>
              </el-tab-pane>
            </el-tabs>
          </el-col>
        </el-row>
        <el-row v-else>
           <client-table :table="clients.slice"></client-table>
        </el-row>
        <br><br><br>
        <el-row type="flex" class="row-bg" align="middle" justify="center" v-if="totalPage != 0">
          <el-pagination layout="prev, pager, next" :total="totalPage" @current-change="updatePage">
          </el-pagination>
        </el-row>
    </div>
    <el-button type="primary" @click="toprojectinput(ClientPicker.value)">Continue to Project Input</el-button>
  </div>
</template>

<script>
import api from '@/api.conf';
import ClientTable from '@/components/ClientTable';
import Login from '@/components/Login';

export default {
  name: 'search',
  components: {
    ClientTable,
    Login,
  },
  props: {
    table: Array,
  },
  data() {
    return {
      clients: {
        data: [],
        slice: [],
        page: 1,
      },
      form: {
        clientName: '',
      },
    };
  },
  created() {
  },
  watch: {
    '$route.params.keyword': 'searchClients',
    page: 'updateSlice',
    table: 'initPropTable',
  },
  methods: {
    searchClients() {
      fetch(api.searchClients, {
        method: 'post',
        headers: {
          Accept: 'application/json',
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          Keyword: this.form.clientName,
        }),
      }).then((response) => {
        response.json().then((data) => {
          if (data === null) {
            this.clients.page = 0;
            this.clients.data = [];
            this.clients.slice = [];
            if (this.name === 'clients') {
              this.updateTotalPage(this.clients.data.length);
              this.page = this.clients.page;
            }
            return;
          }
          this.clients.data = [];
          this.clients.page = 1;
          data.forEach((item) => {
            this.clients.data.push(item);
          });
          // this.clients.data.sort((a, b) => (a.Name.charCodeAt(0) - b.Name.charCodeAt(0)));
          this.clients.slice = this.generateNewSlice(this.clients);
          if (this.name === 'clients') {
            this.updateTotalPage(this.clients.data.length);
            this.page = this.clients.page;
          }
        });
      });
    },
    updatePage(key) {
      this.page = key;
    },
    generateNewSlice(data) {
      return data.data.slice((data.page - 1) * 10, data.page * 10);
    },
    updateSlice() {
      this.clients.page = this.page;
      this.clients.slice = this.generateNewSlice(this.clients);
    },
    updateTotalPage(num) {
      if (num === 0) {
        this.totalPage = 0;
        return;
      }
      this.totalPage = num;
    },
    redirecting() {
      this.$router.push('/new-client');
    },
    toprojectinput(clientnumberID) {
      this.$router.push('/new-project/${clientnumberID}');
    },
    testADAuth(clientnumberID) {
            fetch(api.authenticateAD, {
        method: 'post',
        headers: {
          Accept: 'application/json',
          'Content-Type': 'application/json',
        },
        //  the start and end dates here are they hardcoded or
        //  are they just placeholders where the data gets taken from the datetime picker
        //  temp user & pw to make sure I can actually authenticate with AD
        body: JSON.stringify({
          Username: 'training',
          Password: 'Laptop4Paper3',
        }),
      });
    },
  },
};
</script>