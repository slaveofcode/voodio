import Home from '@/pages/Home'
import Error404 from '@/pages/errors/404'

const routes = [
  {
    path: '/',
    name: 'home',
    component: Home
  },
  {
    path: '/movie/:id',
    name: 'movie-detail',
    component: () => import('@/pages/MovieDetail'),
  },
  {
    path: '/404',
    name: '404',
    component: Error404,
    // Allows props to be passed to the 404 page through route
    // params, such as `resource` to define what wasn't found.
    props: true,
  },
  // Redirect any unmatched routes to the 404 page. This may
  // require some server configuration to work in production:
  // https://router.vuejs.org/en/essentials/history-mode.html#example-server-configurations
  {
    path: '*',
    redirect: '404',
  },
]

export default routes
