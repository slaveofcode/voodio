import Axios from 'axios';

const getRequest = () => {
  return Axios.create({
    baseURL: process.env.VUE_APP_TMDB_API,
  })
}

const searchMovie = (request, defaultParams) => async (keyword) => {
  const { data } = await request.get('/search/movie', {
    params: {
      ...defaultParams,
      query: keyword
    }
  })

  return data
}

const searchFirstByPopularityMovie = (request, defaultParams) => async (keyword) => {
  const { results } = await searchMovie(request, defaultParams)(keyword)
  let highestPops = 0;
  let movie = null;

  for (const item of results) {
    if (item.popularity > highestPops) {
      highestPops = item.popularity;
      movie = item;
    }
  }

  return movie;
}

const getDetailMovieById = (request, defaultParams) => async (id) => {
  const { data } = await request.get(`/movie/${id}`, {
    params: {
      ...defaultParams,
    }
  })

  return data
}

export default (api_key) => {
  const request = getRequest()
  const apiKeyParams = {
    api_key,
  }
  return {
    searchMovie: searchMovie(request, apiKeyParams),
    searchFirstByPopularityMovie: searchFirstByPopularityMovie(request, apiKeyParams),
    getDetailMovieById: getDetailMovieById(request, apiKeyParams),
  }
}