import Axios from 'axios';

const defaultParams = {
  api_key: 'd92ea85a7c28e02fe5fa3abd1b5c426e'
}

const getRequest = () => {
  return Axios.create({
    baseURL: 'https://api.themoviedb.org/3',
  })
}

export const searchMovie = async (keyword) => {
  const { data } = await getRequest().get('/search/movie', {
    params: {
      ...defaultParams,
      query: keyword
    }
  })

  return data
}

export const searchFirstByPopularityMovie = async (keyword) => {
  const { results } = await searchMovie(keyword)
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

export const getDetailMovieById = async (id) => {
  const { data } = await getRequest().get(`/movie/${id}`, {
    params: {
      ...defaultParams,
    }
  })

  return data
}