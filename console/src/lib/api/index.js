import { spinCount } from '../stores'

function incSpinCount() {
  spinCount.update((n) => n + 1)
}

function decSpinCount() {
  spinCount.update((n) => n - 1)
}

function checkInitialResponse(res) {
  if (!res.ok) {
    return res.text().then((text) => {
      throw new Error(text)
    })
  }
  return res.json()
}

function getErrorMessage(err) {
  let msg = ''
  try {
    const errJson = JSON.parse(err.message)
    msg = errJson.error
  } catch (error) {
    msg = err
  }
  return msg
}

function callApi(url, options = {}, done, fail) {
  if (!options.method) {
    options.method = 'GET'
  }
  incSpinCount()

  return fetch(url, options)
    .then((res) => {
      return checkInitialResponse(res)
    })
    .then((data) => {
      if (done) {
        done(data)
      }
      return data
    })
    .catch((err) => {
      const msg = getErrorMessage(err)
      console.log(msg)
      if (fail) {
        fail(msg)
      }
    })
    .finally(decSpinCount)
}

function get(url, options = {}, done, fail) {
  return callApi(url, { ...options, method: 'GET' }, done, fail)
}

function post(url, options = {}, done, fail) {
  return callApi(url, { ...options, method: 'POST' }, done, fail)
}

function del(url, options = {}, done, fail) {
  return callApi(url, { ...options, method: 'DELETE' }, done, fail)
}

// api methods

export function getApiKeys(done, fail) {
  return get(`${BASE_URL}/api/v1/apikeys`, {}, done, fail)
}

export function getApiKeysUsage(done, fail) {
  return get(`${BASE_URL}/api/v1/apikeys/usage`, {}, done, fail)
}

export function registerKey(apiKey, done, fail) {
  return post(`${BASE_URL}/api/v1/apikeys/${apiKey}`, {}, done, fail)
}

export function deleteKey(apiKey, done, fail) {
  return del(`${BASE_URL}/api/v1/apikeys/${apiKey}`, {}, done, fail)
}

export function getTransactions(hours, done, fail) {
  return get(
    `${BASE_URL}/api/v1/transactions/?hours_back=${hours}`,
    {},
    done,
    fail
  )
}

export function writeTransactions(options, done, fail) {
  return post(`${BASE_URL}/api/v1/transactions`, options, done, fail)
}
