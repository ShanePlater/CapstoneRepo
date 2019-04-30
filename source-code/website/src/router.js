import Vue from 'vue';
import Router from 'vue-router';
import Home from '@/views/Home';
import Resources from '@/views/Resources';
import PoliciesProcedures from '@/views/PoliciesProcedures';
import HealthSafety from '@/views/HealthSafety';
import Divisions from '@/views/Divisions';
import Quality from '@/views/Quality';
import HelpDesk from '@/views/HelpDesk';
import Search from '@/views/Search';
import RedbookSearch from '@/views/RedbookSearch';
import NewProject from '@/views/NewProject';
import ProjectDetails from '@/views/ProjectDetails';
import ClientDetails from '@/views/ClientDetails';
import ProjectsSiteInspections from '@/views/ProjectsSiteInspections';

Vue.use(Router);

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Home',
      component: Home,
    },
    {
      path: '/resources',
      name: 'Resources',
      component: Resources,
    },
    {
      path: '/policies&procedures',
      name: 'Policies & Procedures',
      component: PoliciesProcedures,
    },
    {
      path: '/divisions',
      name: 'Divisions',
      component: Divisions,
    },
    {
      path: '/health&safety',
      name: 'Health & Safety',
      component: HealthSafety,
    },
    {
      path: '/quality',
      name: 'Quality',
      component: Quality,
    },
    {
      path: '/help-desk',
      name: 'Help Desk',
      component: HelpDesk,
    },
    {
      path: '/redbook',
      name: 'Redbook',
      component: RedbookSearch,
    },
    {
      path: '/newproject',
      name: 'Add New Project',
      component: NewProject,
    },
    {
      path: '/search/:keyword',
      name: 'Search Result',
      component: Search,
    },
    {
      path: '/project/:id',
      name: 'Project Details',
      component: ProjectDetails,
    },
    {
      path: '/client/:id',
      name: 'Client Details',
      component: ClientDetails,
    },
    {
      path: '/siteinspections/:id',
      name: 'Project Site Inspections',
      component: ProjectsSiteInspections,
    }],
});
