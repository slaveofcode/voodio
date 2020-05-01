<template>
  <Layout>
    <Header title="Voodio Media Server" />
    <div class="flex flex-row align-top flex-wrap justify-evenly">
      <div class="movie-cover border border-solid border-gray-600 mb-10 text-center overflow-hidden relative" v-for="(movie, idx) in movies" :key="idx">
        <router-link :to="{ name: movie.isGroupDir ? 'movie-group-detail' : 'movie-detail', params: { id: movie.ID }, query: { tmdbId: (movie.details ? movie.details.id : 0) } }" class="text-xl">
          <img v-if="movie.details" class="bg-cover" :src="parseCover(movie.details.poster_path)" />
          <img v-if="!movie.details" class="bg-cover" src="../assets/movie.svg" />
          <div class="movie-title absolute bottom-0 text-center w-full h-10 pt-1">
            <span v-if="movie.isInPrepare || movie.isPrepared" class="text-green-500">{{ movie.cleanDirName }}</span>
            <span v-if="!(movie.isInPrepare || movie.isPrepared)">{{ movie.cleanDirName }}</span>
          </div>
        </router-link>
      </div>
    </div>
  </Layout>
</template>

<style lang="scss" scoped>
.movie-cover {
  width: 250px;
  height: 370px;
  display: flex;
  flex-direction: column;
  justify-content: center;
}
.movie-title {
  background: rgba(0,0,0, .6);
}
</style>

<script>
import { mapGetters } from 'vuex';
import Layout from '@/layouts/Main';
import Header from '@/components/Header';
import { getMovies } from '@/utils/voodio_request';
import tmdbApi from '@/utils/tmdb_request'; 

export default {
  components: {
    Layout,
    Header
  },
  created() {
    if (!this.tmdbApiKey) {
      this.$store.dispatch('tmdb/fetchTMDBApi').then(() => {
        this.prepareMovieList()
      })
    } else {
      this.prepareMovieList()
    }
  },
  data() {
    return {
      movies: [],
      totalMovies: 0
    };
  },
  computed: {
    ...mapGetters({
      tmdbApiKey: 'tmdb/tmdb_api_key',
    }),
  },
  methods: {
    prepareMovieList() {
      const { searchFirstByPopularityMovie } = tmdbApi(this.tmdbApiKey)
      getMovies().then(({ movies, count }) => {
        this.movies = movies
        this.totalMovies = count

        for (const m of this.movies) {
          searchFirstByPopularityMovie(m.cleanBaseName).then((mov) => {
            if (!mov) {
              // search again with another keyword
              return searchFirstByPopularityMovie(m.cleanDirName).then((movv) => {
                m.details = movv;
                this.$forceUpdate()
              })
            }
            m.details = mov;
            this.$forceUpdate()
          })
        }
      })
    },
    parseCover(coverFile) {
      return `https://image.tmdb.org/t/p/w300${coverFile}`
    }
  }
};
</script>