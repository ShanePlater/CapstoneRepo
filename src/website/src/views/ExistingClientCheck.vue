<template>
  <div>
    <h1>Does this project belong to a new or existing client?</h1>
    <br>
    <p>If your client does not already exist, please click the the link below to add a new client. If the client does exist, search it in the field below then click Continue
    <el-button type="primary" @click="redirecting">Add New Client</el-button>
    <br>
    <el-select
      v-model="ClientPicker"
      multiple
      filterable
      remote
      reserve-keyword
      placeholder="Please enter a client"
      :remote-method="remoteMethod"
      :loading="loading">
      <el-option v-for="option in options.clients" :key="option.Name" :label="option.Name" :value="option.Name"></el-option>
    </el-select>

    <el-form ref="form" :model="form" label-width="12.5em" label-position="left">
      <el-form-item label="Division">
        <el-select v-model="form.division" placeholder="Pick a division">
          <el-option v-for="option in options.clients" :key="option.Name" :label="option.Name" :value="option.Name"></el-option>
        </el-select>
      </el-form-item>
    </el-form>

    <el-button type="primary" @click="toprojectinput(ClientPicker.value)">Continue to Project Input</el-button>
  </div>
</template>

<script>
import api from '@/api.conf';

export default {
  data() {
    return {
      options: {
        clients: [],
      },
      value: [],
      list: [],
      loading: false,
      states: [ // They need to be in JSON object form like { value: "clientID", label: "client name"}.
      ],
      form: {
        division: '',
      },
    };
  },
  created() {
    this.getClients();
  },
  /* mounted() {
    this.list = this.states.map(item => {
      return { value: item, label: item };
    });
  },
  */
  methods: {
    /*
    remoteMethod(query) {
      if (query !== '') {
        this.loading = true;
        setTimeout(() => {
          this.loading = false;
          this.options = this.list.filter(item => {
            return item.label.toLowerCase().indexOf(query.toLowerCase()) > -1;
          });
        }, 200);
      } else {
        this.options = [];
      }
    },
    */
    getClients() {
      fetch(api.getOptionClients, {
        method: 'get',
        headers: {
          Accept: 'application/json',
          'Content-Type': 'application/json',
        },
      }).then((reponse) => {
        reponse.json().then((data) => {
          this.options.clients = data;
        });
      });
    },
    redirecting() {
      this.$router.push(`/new-client`);
    },
    toprojectinput(clientnumberID) {
      this.$router.push(`/new-project/${clientnumberID}`);
    },
  },
};
</script>