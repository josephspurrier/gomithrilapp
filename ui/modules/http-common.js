import axios from 'axios'

// Source: https://alligator.io/vuejs/rest-api-axios/

export const HTTP = axios.create({
  // baseURL: process.env.baseURL
  baseURL: 'http://localhost:8081'
  /* headers: {
      Authorization: 'Bearer {token}'
  } */
})
