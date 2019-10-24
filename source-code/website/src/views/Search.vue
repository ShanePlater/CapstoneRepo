<template>
  <div>
    <el-row>
      <h1>Search Results</h1>
    </el-row>
    <el-row type="flex" class="row-bg" v-if="table === undefined">
      <el-col :span="24">
        <el-tabs v-model="name" type="card" @tab-click="handleTabClick">
          <el-tab-pane label="Projects" name="projects">
            <project-table :table="projects.slice"></project-table>
          </el-tab-pane>
          <el-tab-pane label="Staff Profiles" name="users">
            <user-table :table="users.slice"></user-table>
          </el-tab-pane>
          <el-tab-pane label="Client Profiles" name="clients">
            <client-table :table="clients.slice"></client-table>
          </el-tab-pane>
          <el-tab-pane label="Documents" name="documents">
            <resource-table :table="documents.slice"></resource-table>
          </el-tab-pane>
        </el-tabs>
      </el-col>
    </el-row>
    <el-row v-else>
      <project-table :table="projects.slice"></project-table>
    </el-row>
    <br><br><br>
    <el-row type="flex" class="row-bg" align="middle" justify="center" v-if="totalPage != 0">
      <el-pagination layout="prev, pager, next" :total="totalPage" @current-change="updatePage">
      </el-pagination>
    </el-row>
  </div>
</template>

<script>
import api from '@/api.conf';
import ProjectTable from '@/components/ProjectTable';
import ClientTable from '@/components/ClientTable';
import UserTable from '@/components/UserTable';
import ResourceTable from '@/components/ResourceTable';

export default {
  name: 'search',
  components: {
    ProjectTable,
    ClientTable,
    UserTable,
    ResourceTable,
  },
  props: {
    table: Array,
  },
  data() {
    return {
      projects: {
        data: [],
        slice: [],
        page: 1,
      },
      users: {
        data: [],
        slice: [],
        page: 1,
      },
      clients: {
        data: [],
        slice: [],
        page: 1,
      },
      documents: {
        data: [],
        slice: [],
        page: 1,
      },
      page: 0,
      totalPage: 0,
      name: 'projects',
    };
  },
  created() {
    this.pullData();
  },
  watch: {
    '$route.params.keyword': 'pullData',
    page: 'updateSlice',
    table: 'initPropTable',
  },
  methods: {
    pullData() {
      if (this.table === undefined) {
        this.searchProjects();
        this.searchUsers();
        this.searchClients();
        this.searchResources();
      }
    },
    searchProjects() {
      fetch(api.searchProjects, {
        method: 'post',
        headers: {
          Accept: 'application/json',
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          Keyword: this.$route.params.keyword,
        }),
      }).then((response) => {
        response.json().then((data) => {
          if (data === null) {
            this.projects.page = 0;
            this.projects.data = [];
            this.projects.slice = [];
            if (this.name === 'projects') {
              this.updateTotalPage(this.projects.data.length);
              this.page = this.projects.page;
            }
            return;
          }
          this.projects.data = [];
          this.projects.page = 1;
          data.forEach((item) => {
            this.projects.data.push(item);
          });
          // this.projects.data.sort((a, b) => (a.ID.charCodeAt(0) - b.ID.charCodeAt(0)));
          this.projects.slice = this.generateNewSlice(this.projects);
          if (this.name === 'projects') {
            this.updateTotalPage(this.projects.data.length);
            this.page = this.projects.page;
          }
        });
      });
    },
    searchClients() {
      fetch(api.searchClients, {
        method: 'post',
        headers: {
          Accept: 'application/json',
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          Keyword: this.$route.params.keyword,
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
    searchUsers() {
      fetch(api.searchUsers, {
        method: 'post',
        headers: {
          Accept: 'application/json',
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          Keyword: this.$route.params.keyword,
        }),
      }).then((response) => {
        response.json().then((data) => {
          if (data === null) {
            this.users.page = 0;
            this.users.data = [];
            this.users.slice = [];
            if (this.name === 'users') {
              this.updateTotalPage(this.users.data.length);
              this.page = this.users.page;
            }
            return;
          }
          this.users.data = [];
          this.users.page = 1;
          data.forEach((item) => {
            this.users.data.push(item);
          });
          // this.users.data.sort((a, b) => (a.Name.charCodeAt(0) - b.Name.charCodeAt(0)));
          this.users.slice = this.generateNewSlice(this.users);
          if (this.name === 'users') {
            this.updateTotalPage(this.users.data.length);
            this.page = this.users.page;
          }
        });
      });
    },
    searchResources() {
      fetch(api.searchResources, {
        method: 'post',
        headers: {
          Accept: 'application/json',
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          Keyword: this.$route.params.keyword,
        }),
      }).then((response) => {
        response.json().then((data) => {
          if (data === null) {
            this.documents.page = 0;
            this.documents.data = [];
            this.documents.slice = [];
            if (this.name === 'documents') {
              this.updateTotalPage(this.documents.data.length);
              this.page = this.documents.page;
            }
            return;
          }
          this.documents.data = [];
          this.documents.page = 1;
          data.forEach((item) => {
            this.documents.data.push({
              AuthorizedBy: item.AuthorizedBy,
              AuthorizedDate: Intl.DateTimeFormat('en-AU').format(new Date(item.AuthorizedDate)),
              CategoryID: item.CategoryID,
              FileID: item.FileID,
              FileName: item.FileName,
              FileRevision: item.FileRevision.trim(),
              URL: item.URL,
            });
          });
          // this.documents.data.sort((a, b) =>
          // (a.FileName.charCodeAt(0) - b.FileName.charCodeAt(0)));
          this.documents.slice = this.generateNewSlice(this.documents);
          if (this.name === 'documents') {
            this.updateTotalPage(this.documents.data.length);
            this.page = this.documents.page;
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
      switch (this.name) {
        case 'projects':
          this.projects.page = this.page;
          this.projects.slice = this.generateNewSlice(this.projects);
          break;
        case 'clients':
          this.clients.page = this.page;
          this.clients.slice = this.generateNewSlice(this.clients);
          break;
        case 'users':
          this.users.page = this.page;
          this.users.slice = this.generateNewSlice(this.users);
          break;
        case 'documents':
          this.documents.page = this.page;
          this.documents.slice = this.generateNewSlice(this.documents);
          break;
        default:
      }
    },
    handleTabClick() {
      switch (this.name) {
        case 'projects':
          this.updateTotalPage(this.projects.data.length);
          this.page = this.projects.page;
          break;
        case 'clients':
          this.updateTotalPage(this.clients.data.length);
          this.page = this.clients.page;
          break;
        case 'users':
          this.updateTotalPage(this.users.data.length);
          this.page = this.users.page;
          break;
        case 'documents':
          this.updateTotalPage(this.documents.data.length);
          this.page = this.documents.page;
          break;
        default:
      }
    },
    updateTotalPage(num) {
      if (num === 0) {
        this.totalPage = 0;
        return;
      }
      this.totalPage = num;
    },
    initPropTable() {
      if (this.table.length === 0) {
        this.projects.page = 0;
        return;
      }
      this.projects.data = this.table;
      this.name = 'projects';
      this.projects.slice = this.generateNewSlice(this.projects);
      this.updateTotalPage(this.projects.data.length);
      this.page = this.projects.page;
    },
  },
};
</script>

<style scoped>

</style>
