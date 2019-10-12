<template>
  <section>
    <div>
      <h1>Test Hello</h1>
      <br>
      <el-row>
        <el-col :span="12">
          <el-form ref="form" :model="form" label-width="12.5em" label-position="left">
            <el-form-item label="Client Name">
              <el-input v-model="form.ClientName" value="test"></el-input>
            </el-form-item>

           <!-- CLIENT INFORMATION  --> 
          <h2 style="font-size:20px"> Client Information </h2>
          
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
  name: 'new-client',
  components: {

  },
  data() {
    return {
      title: 'New Client',
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
        ClientName: '',
      },
      isSearching: true,
    };
  },
  created() {
    if (this.$route.query.res === 'true') {
      this.$router.replace('/NewClient');
    }
    this.pullClientDetails();
//    this.getOptions(api.getOptionLocations);
//    this.getOptions(api.getOptionTypes);
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
      };
    },
  },
};
</script>

<style scoped>

</style>
