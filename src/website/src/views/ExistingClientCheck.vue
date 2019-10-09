<template>
  <div>
    <h1>Does this project belong to a new or existing client?</h1>
    <br>
    <p>If your client does not already exist, please click the the link below to add a new client. If the client does exist, search it in the field below then click Continue
    <el-button type="primary" @click="redirecting">Add New Client</el-button>
    <br>
    <el-select
      v-model="value"
      multiple
      filterable
      remote
      reserve-keyword
      placeholder="Please enter a client"
      :remote-method="remoteMethod"
      :loading="loading">
      <el-option v-for="item in options" :key="item.value" :label="item.label" :value="item.value"></el-option>
    </el-select>
    <br>

    <el-button type="primary" @click="toprojectinput(value)">Continue to Project Input</el-button>
  </div>
</template>

<script>
export default {
  data() {
    return {
      options: [],
      value: [],
      list: [],
      loading: false,
      states: [ // They need to be in JSON object form like { value: "clientID", label: "client name"}.
      ],
    };
  },
  mounted() {
    this.list = this.states.map(item => {
      return { value: item, label: item };
    });
  },
  methods: {
    remoteMethod(query) {
      if (query !== "") {
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
    getClients(){
       fetch(api.getClient, {
        method: 'get',
        headers: {
          Accept: 'application/json',
          'Content-Type': 'application/json',
        },
      }).then((reponse) => {
        reponse.json().then((data) => {
          this.options = data;
        });
      });
    },
    redirecting() {
      this.$router.push(`/new-client`);
    },
    toprojectinput(clientnumberID) {
      this.$router.push(`/new-project/${clientnumberID}`);
    },
  }
};
</script>