import Axios from 'axios';
import { getCurrHost, getCurrPort } from './url'

const getRequest = () => {
  return Axios.create({
    baseURL: `http://${getCurrHost()}:${getCurrPort()}`
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

export const prepareMovie = async (movieId) => {
  const { status, data } = await getRequest().get('/movies/prepare', {
    params: { movieId }
  })

  if (status !== 200) {
    return {}
  }

  return data
}