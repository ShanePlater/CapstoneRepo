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
import NewProject from '@/views/NewProject';
import NewClient from '@/views/NewClient';
import RedbookSearch from '@/views/RedbookSearch';
import ProjectDetails from '@/views/ProjectDetails';
import ClientDetails from '@/views/ClientDetails';
import ProjectsSiteInspections from '@/views/ProjectsSiteInspections';
import SiteInspection from '@/views/SiteInspection';
import DocumentUpload from '@/views/DocumentUpload';
import ExistingClientCheck from '@/views/ExistingClientCheck';
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
      path: '/new-project',
      name: 'Add New Project',
      component: NewProject,
    },
    {
      path: '/new-client',
      name: 'Add New Client',
      component: NewClient,
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
      path: '/siteinspection/:id',
      name: 'Site Inspection',
      component: SiteInspection,
    },
    {
      path: '/siteinspections/:id',
      name: 'Project Site Inspections',
      component: ProjectsSiteInspections,
    },
    {
      path: '/document-upload',
      name: 'Upload New Document',
      component: DocumentUpload,
    },
    {
      path: '/ExistingClientCheck',
      name: 'Existing Client Check',
      component: ExistingClientCheck,
    }],
});
