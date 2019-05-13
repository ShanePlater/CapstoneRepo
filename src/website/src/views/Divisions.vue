<template>
  <section>
    <h1>{{ title }}</h1>
    <el-select v-model="office" placeholder="Pick a office">
      <el-option v-for="option in options.offices" :key="option.ID" :label="option.Name" :value="option.ID"></el-option>
    </el-select>
    <el-select v-model="division" placeholder="Pick a division">
      <el-option v-for="option in options.divisions" :key="option.Name" :label="option.Name" :value="option.Name"></el-option>
    </el-select>
    <el-row v-if="office === '' || division === ''">
      <el-col :span="24">
        <div>
          <h1 class="content" align="center">Please select an option for both selectors to show division details.</h1>
        </div>
      </el-col>
    </el-row>
    <br><br>
    <el-row :gutter="24">
      <el-col :span="6" v-for="item in data" :key="item.Username">
        <el-card style="height:30em">
          <img :src="item.PhotoURL" class="image" width="914" height="1280">
          <div style="padding: 1em;">
            <span align="center" style="padding-left: 1em;">{{ item.Name }}</span>
            <div class="bottom">
                <ul class="fa-ul">
                  <li v-if=" item.PositionTitle !== ''">
                    <i class="fa-li fa fa-user fa-fw"></i> {{ item.PositionTitle }}</li>
                  <li v-if=" item.EmailAddress !== ''">
                    <i class="fa-li fa fa-envelope-o fa-fw"></i> {{ item.EmailAddress }}</li>
                  <li v-if=" item.MobileNumber !== ''">
                    <i class="fa-li fa fa-mobile fa-fw"></i> {{ item.MobileNumber }}</li>
                  <li v-if=" item.PhoneExtension !== ''">
                    <i class="fa-li fa fa-phone fa-fw"></i> {{ item.PhoneExtension }}</li>
                </ul>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </section>
</template>

<script>
import api from '@/api.conf';

export default {
  name: 'divisions',
  data() {
    return {
      title: 'Divisions',
      data: [],
      office: '',
      division: '',
      options: {
        offices: [],
        divisions: [],
      },
    };
  },
  created() {
    this.getOptions(api.getOptionOffices);
    this.getOptions(api.getOptionDivisions);
  },
  watch: {
    office: 'updateDivision',
    division: 'updateDivision',
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
            case api.getOptionOffices:
              this.options.offices = data;
              break;
            case api.getOptionDivisions:
              this.options.divisions = data;
              break;
            default:
          }
        });
      });
    },
    updateDivision() {
      if (this.office === '' || this.division === '') {
        return;
      }
      fetch(api.getDivisions, {
        method: 'post',
        headers: {
          Accept: 'application/json',
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          Office: this.office,
          Division: this.division,
        }),
      }).then((response) => {
        response.json().then((data) => {
          this.data = data;
        });
      });
    },
  },
};
</script>

<style scoped>
.content {
  padding-top: 20%;
  font-size: 2em;
  opacity: 0.5;
}

.bottom {
  opacity: 0.6;
}

.image {
  width: 80%;
  height: auto;
  display: block;
  margin: auto;
}

</style>
