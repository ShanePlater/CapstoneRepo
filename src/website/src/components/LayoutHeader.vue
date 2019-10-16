<template>
  <section>
    <el-row :gutter="20" type="flex" justify="space-between" class="row-bg" align="bottom">
      <el-col :span="6">
        <img src="../assets/logo.png">
      </el-col>
      <el-col :span="6" :offset="12">
        <login></login>
      </el-col>
    </el-row>
    <el-row type="flex" align="middle" class="grid-bg" justify="space-around">
      <el-col :span="19">
        <el-menu :default-active="activeIndex" mode="horizontal" :router="true">
          <el-menu-item v-for="item in items" :key="item.name" :index="item.path" id="menu-item" v-if="!item.path.includes(':')">
            {{ item.name }}
          </el-menu-item>
        </el-menu>
      </el-col>
      <el-col :span="4">
        <form @submit.prevent="search">
          <!-- <el-autocomplete v-model="input" :fetch-suggestions="fetchSuggestions" placeholder="Internal Search..." @select="searchSelected" icon="search" :on-icon-click="search" :trigger-on-focus="false"></el-autocomplete> -->
          <el-input v-model="input" placeholder="Internal Search..." :on-icon-click="search" icon="search"></el-input>
        </form>
      </el-col>
    </el-row>
  </section>
</template>

<script>
import router from '@/router';
import api from '@/api.conf';
import Login from '@/components/Login';

export default {
  name: 'layout-header',
  components: {
    Login,
  },
  data() {
    return {
      activeIndex: this.$route.path,
      items: router.options.routes,
      input: '',
    };
  },
  methods: {
    search() {
      if (this.input.trim() === '') {
        return;
      }
      this.$router.push(`/search/${this.input}`);
      this.input = '';
    },
    fetchSuggestions(input, cb) {
      if (this.input === '') {
        return;
      }
      fetch(api.fetchSuggestions, {
        method: 'post',
        headers: {
          Accept: 'application/json',
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          Keyword: input,
        }),
      }).then((response) => {
        response.json().then((data) => {
          if (data === null) {
            return;
          }
          const res = [];
          data.forEach((i) => {
            res.push({ value: i });
          });
          cb(res);
        });
      });
    },
    searchSelected(item) {
      this.input = '';
      fetch(api.getSuggestionItemURL, {
        method: 'post',
        headers: {
          Accept: 'application/json',
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          Keyword: item.value,
        }),
      }).then((response) => {
        response.json().then((data) => {
          if (data.URL === '') {
            this.$router.push(`/search/${item.value}`);
            return;
          }
          if (data.URL.includes('static')) {
            window.open(data.URL);
            return;
          }
          this.$router.push(data.URL);
        });
      });
    },
  },
};
</script>

<style scoped>
.row-bg {
  padding: 1em 0;
}

.profile {
  width: 3em;
  height: 3em;
  border-radius: 50%;
}

.grid-bg {
  background: #EFF2F7;
}
</style>
