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

export function getApiKeys(done, fail) {
  incSpinCount()

  return fetch(`${BASE_URL}/api/v1/apikeys`)
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

export function register(apiKey, done, fail) {
  incSpinCount()

  fetch(`${BASE_URL}/api/v1/apikeys/${apiKey}`, {
    method: 'POST',
  })
    .then((res) => {
      return checkInitialResponse(res)
    })
    .then((_) => {
      if (done) {
        done()
      }
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
