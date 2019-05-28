import axios from 'axios'
import Flash from '~/modules/flash.js'

// Source: https://alligator.io/vuejs/rest-api-axios/

export const HTTP = axios.create({
  // baseURL: process.env.baseURL
  baseURL: 'http://localhost:8081'
})

export const HTTP2 = {
  logic: function (method, endpointURL, form, auth, successMessage, successFunc) {
    const f = new Flash()
    let success = false

    let accessToken = ''
    if (auth && auth.loggedIn) {
      accessToken = auth.accessToken
    }

    // Send a login request to the server.
    const headers = {
      headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer ' + accessToken,
      }
    }

    let h = {}
    if (form) {
      h = HTTP[method](endpointURL, form, headers)
    } else {
      h = HTTP[method](endpointURL, headers)
    }

    h.then(response => {
      if (response.data !== undefined) {
        f.success(successMessage)
        successFunc(response.data)
        success = true
      } else {
        f.failed('Token is not in the correct format.')
      }
    })
      .catch(err => {
        if (err.response === undefined) {
          f.warning(
            'There was an error reaching the server. Please try again later.' +
            err
          )
        } else if (err.response.data.message !== undefined) {
          f.warning(err.response.data.message)
        } else {
          f.warning('There was an error. Please try again later.' + err)
        }
      })
      .finally(() => {
        if (!success) {
        }
      })
  },
}