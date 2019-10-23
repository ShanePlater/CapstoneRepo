<template>
  <el-table :data="table" stripe style="width: 100%">
    <el-table-column prop="ID" label="Inspection ID" width="130">
    </el-table-column>
    <el-table-column prop="InspectedBy" label="Inspected By" width="150">
    </el-table-column>
    <el-table-column prop="InspectionDateTime" label="Inspection Date" width="160" :show-overflow-tooltip="true">
    </el-table-column>
    <el-table-column prop="InspectionDetails" label="Details" :show-overflow-tooltip="true">
    </el-table-column>
    <el-table-column>
      <template scope="scope">
        <el-button class="btn" size="small" type="text" icon="el-icon-document-add" @click="handleClick(scope.row)">Delete</el-button>
      </template>
    </el-table-column>
  </el-table>
</template>

<script>
import api from '@/api.conf';

export default {
  name: 'inspections-table',
  props: {
    table: Array,
  },
  created(){
    if (this.getCookie('name') !== '') {
      this.state = {
        name: this.getCookie('name'),
      };
    }
  },
  methods: {
    handleClick(row) {
      fetch(api.deleteSiteInspection, {
        method: 'post',
        headers: {
          Accept: 'application/json',
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          ID: row,
        }),
      });
    },
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
          this.handleClick();
        }
      });
    });
  },
};
</script>