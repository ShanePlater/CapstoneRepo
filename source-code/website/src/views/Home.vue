<template>
  <section>
    <h1>News & Information</h1>
    <h3 v-if="!more">{{ date }}</h3>
    <el-row v-if="more">
      <el-date-picker v-model="datePeriod" type="daterange" placeholder="Pick a range" :picker-options="datePicker" :clearable="false" @change="updateNews">
      </el-date-picker>
    </el-row>
    <br v-if="more">
    <el-row v-for="item in newsSlice" :key="item.ID">
      <el-col :span="19">
        <news-card :item="item"></news-card>
        <br><br>
      </el-col>
      <el-col :span="4" style="float:right;" v-if="item == newsSlice[0]">
        <whats-new :recent="recent"></whats-new>
      </el-col>
    </el-row>
    <el-button type="primary" @click="more = !more; updatePage()" v-if="!more">
      Show More
    </el-button>
    <el-pagination layout="prev, pager, next" :total="news.length" @current-change="updatePage" v-else>
    </el-pagination>
  </section>
</template>

<script>
import NewsCard from '@/components/NewsCard';
import WhatsNew from '@/components/WhatsNew';
import api from '@/api.conf';

export default {
  name: 'home',
  components: {
    NewsCard,
    WhatsNew,
  },
  data() {
    return {
      datePeriod: '',
      more: false,
      page: 1,
      date: this.todayDate(),
      news: [],
      newsSlice: [],
      recent: [],
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
    };
  },
  created() {
    this.recentAPI();
    const start = new Date();
    start.setTime(start.getTime() - (3600 * 1000 * 24 * 365));
    this.datePeriod = [start, new Date()];
    this.updateNews();
  },
  watch: {
    page: 'updateSlice',
    news: 'updateSlice',
  },
  methods: {
    todayDate: () => {
      const date = new Date();
      const res = date.toString().split(' ');
      switch (date.getDay()) {
        case 0:
          this.date = 'Sunday';
          break;
        case 1:
          this.date = 'Monday';
          break;
        case 2:
          this.date = 'Tuesday';
          break;
        case 3:
          this.date = 'Wednesday';
          break;
        case 4:
          this.date = 'Thursday';
          break;
        case 5:
          this.date = 'Friday';
          break;
        case 6:
          this.date = 'Saturday';
          break;
        default:
      }
      return this.date.concat(', ', res[1], ' ', res[2], ' ', res[3]);
    },
    getNewsAPI(req) {
      fetch(api.getNews, {
        method: 'post',
        headers: {
          Accept: 'application/json',
          'Content-Type': 'application/json',
        },
        body: req,
      }).then((response) => {
        response.json().then((data) => {
          if (data === null) {
            this.$message({
              showClose: true,
              type: 'error',
              duration: 5000,
              message: `No News found between ${Intl.DateTimeFormat('en-AU').format(this.datePeriod[0])} and ${Intl.DateTimeFormat('en-AU').format(this.datePeriod[1])}`,
            });
            return;
          }
          this.news = data;
        });
      });
    },
    recentAPI() {
      fetch(api.recent, {
        method: 'post',
        headers: {
          Accept: 'application/json',
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          Length: 5,
        }),
      }).then((response) => {
        response.json().then((data) => {
          this.recent = data;
        });
      });
    },
    updatePage(key) {
      if (key !== undefined) {
        this.page = key;
      }
      window.scrollTo(0, 0);
    },
    updateSlice() {
      this.newsSlice = this.news.slice((this.page - 1) * 5, this.page * 5);
    },
    updateNews() {
      this.getNewsAPI(JSON.stringify({
        StartDate: this.datePeriod[0].toJSON(),
        EndDate: this.datePeriod[1].toJSON(),
      }));
    },
  },
};
</script>

<style scoped>
h3 {
  opacity: 0.70;
}
</style>
