<template>
  <section>
    <div v-if="title === 'Redbook Search'">
      <h1>{{ title }}</h1>
      <br>
      <el-row>
        <el-col :span="12">
          <el-form ref="form" :model="form" label-width="12.5em" label-position="left">
            <el-form-item label="Client Name">
              <el-input v-model="form.clientName"></el-input>
            </el-form-item>
            <el-form-item label="Address">
              <el-input v-model="form.address"></el-input>
            </el-form-item>
            <el-form-item label="Suburb">
              <el-input v-model="form.suburb"></el-input>
            </el-form-item>
            <el-form-item label="Location">
              <el-select v-model="form.location" placeholder="Pick a location">
                <el-option v-for="option in options.locations" :key="option.ID" :label="option.Name" :value="option.ID"></el-option>
              </el-select>
            </el-form-item>
            <el-form-item label="Type">
              <el-select v-model="form.type" placeholder="Pick a type">
                <el-option v-for="option in options.types" :key="option.ID" :label="option.Name" :value="option.ID"></el-option>
              </el-select>
            </el-form-item>
            <el-form-item label="Status">
              <el-select v-model="form.status" placeholder="Pick a status">
                <el-option v-for="option in options.statuss" :key="option.ID" :label="option.Name" :value="option.ID"></el-option>
              </el-select>
            </el-form-item>
            <el-form-item label="Division">
              <el-select v-model="form.division" placeholder="Pick a division">
                <el-option v-for="option in options.divisions" :key="option.Name" :label="option.Name" :value="option.Name"></el-option>
              </el-select>
            </el-form-item>
            <el-form-item label="Office">
              <el-select v-model="form.office" placeholder="pick a office">
                <el-option v-for="option in options.offices" :key="option.ID" :label="option.Name" :value="option.ID"></el-option>
              </el-select>
            </el-form-item>
            <el-form-item label="Project Date Scope">
              <el-date-picker v-model="form.datePeriod" type="daterange" placeholder="Pick a range" :picker-options="datePicker">
              </el-date-picker>
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="search">Search it now</el-button>
            </el-form-item>
          </el-form>
        </el-col>
      </el-row>
    </div>
    <div v-else>
      <el-row v-loading="isSearching">
        <search :table="projects"></search>
      </el-row>
    </div>
  </section>
</template>

<script>
import api from '@/api.conf';
import Search from '@/views/Search';

export default {
  name: 'redbook-search',
  components: {
    Search,
  },
  data() {
    return {
      title: 'Redbook Search',
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
        clientName: '',
        address: '',
        suburb: '',
        location: '',
        type: '',
        status: '',
        division: '',
        office: '',
        datePeriod: '',
      },
      isSearching: true,
    };
  },
  created() {
    if (this.$route.query.res === 'true') {
      this.$router.replace('/redbook');
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
    search() {
      this.$router.replace(`${this.$route.fullPath}?res=true`);
      let start;
      let end;
      if (this.form.datePeriod === '') {
        start = '';
        end = '';
      } else {
        start = this.form.datePeriod[0].toJSON();
        end = this.form.datePeriod[1].toJSON();
      }
      fetch(api.searchRedbook, {
        method: 'post',
        headers: {
          Accept: 'application/json',
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          ClientName: this.form.clientName,
          Address: this.form.address,
          Suburb: this.form.suburb,
          Location: this.form.location,
          Type: this.form.type,
          Status: this.form.status,
          Division: this.form.division,
          Office: this.form.office,
          StartDate: start,
          EndDate: end,
        }),
      }).then((response) => {
        response.json().then((data) => {
          if (data === null) {
            this.isSearching = false;
            return;
          }
          data.forEach((item) => {
            this.projects.push(item);
          });
          // this.projects.sort((a, b) => (a.ID.charCodeAt(0) - b.ID.charCodeAt(0)));
          this.isSearching = false;
        });
      });
    },
    updatePage() {
      if (this.$route.query.res === 'true') {
        this.title = 'z z';
        return;
      }
      this.title = 'zzz';
      this.projects = [];
      this.form = {
        clientName: '',
        address: '',
        suburb: '',
        location: '',
        type: '',
        status: '',
        division: '',
        office: '',
        datePeriod: '',
      };
    },
  },
};
</script>

<style scoped>

</style>
