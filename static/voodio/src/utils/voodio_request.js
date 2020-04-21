import Axios from 'axios';

const getRequest = () => {
  return Axios.create({
    baseURL: 'http://192.168.8.102:1818'
  })
}

export const getMovies = async () => {
  const { status, data } = await getRequest().get('/movies')
  if (status !== 200) {
    return []
  }

  return data
}

export const getMovieDetail = async (movieId) => {
  const { status, data } = await getRequest().get('/movies/detail', {
    params: { movieId }
  })

  if (status !== 200) {
    return {}
  }

  return data
}